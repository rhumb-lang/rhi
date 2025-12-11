## 4\. Concurrency & The Tuplespace

Rhumb replaces the traditional Call Stack with a **Hierarchical Tuplespace** based on the *Syndicated Actor Model (SAM)*. This system unifies Concurrency, Event Handling, and State Management into a single spatial metaphor.

### 4\.1 Conceptual Model

  * **The Ether:** Every executing process (Green Thread) possesses a **Local Space**.
  * **Hierarchy:** Spaces are arranged in a tree. Every Space has a reference to its **Parent Space** (who spawned it).
  * **Zombie Frames:** When a process yields or pauses (e.g., waiting for a signal), its Stack Frame is not destroyed. It remains in memory as a "Zombie," allowing deeper frames (Children) to "Drill Down" and reply to it later.

### 4\.2 Operator Taxonomy

The syntax distinguishes between ephemeral events and persistent state to prevent race conditions.

| Feature          | Symbol  | Type  | Direction        | Semantics                                                                                                       |
|:-----------------|:--------|:------|:-----------------|:----------------------------------------------------------------------------------------------------------------|
| **Signal**       | **`#`** | Event | **Up** (Bubble)  | **Request.** Pauses execution. Bubbles up. Resumes with result of `^Reply` (or `___` if unhandled).             |
| **Reply**        | **`^`** | Event | **Down** (Drill) | **Response.** Targets a paused Zombie Frame. Injects a payload and resumes that frame.                          |
| **Proclamation** | **`$`** | State | **Static** (Pin) | **Persistent.** Sticks to the Local Space. To **Retract** (delete), assert the Empty Value: `realm$state(___)`. |

### 4\.3 Realm Syntax & Lifecycle

Realms are reified Spaces that can be assigned to variables.

  * **Child Realm `<$>`**: Creates a standard Space.
      * *Behavior:* Signals uncaught in this realm bubble up to the creator's current space.
      * *Use Case:* Workers, sub-components.
  * **Detached Realm `<|>`**: Creates a Sandboxed Space.
      * *Behavior:* `Parent` is set to `World` (Root). Signals hitting the ceiling are discarded/logged.
      * *Use Case:* Top-level servers, sandboxed plugins.
  * **Opcode:** `NEW_REALM <flags>`
      * Flag 0: Child (`Parent = CurrentSpace`)
      * Flag 1: Detached (`Parent = World`)

### 4\.4 The "Helium Balloon" Algorithm (Signal Propagation)

To prevent memory leaks in long-running servers, Signals (`#`) are active agents of transport.

1.  **Post & Pause:** Instruction `POST` suspends the `CurrentFrame` (marking it Zombie).
2.  **Bubble:** The signal bubbles up looking for a listener.
3.  **Outcome A (Replied):** A listener traps it and executes `INJECT ^val`. The Zombie resumes with `val` on the stack.
4.  **Outcome B (Unhandled):** The signal hits `World`. The Zombie resumes with `___` on the stack.
5.  **Garbage Collection:** If the Signal reaches `World` (Root) and is still uncaught, it is **discarded**. This ensures that "fire-and-forget" events do not accumulate in memory.

### 4\.5 The "Zombie Walk" Algorithm (Reply Injection)

Replies (`^`) allow a helper process to inject data back into a paused requestor.

1.  **Trigger:** Instruction `INJECT ^ack` is executed in a shallow frame (e.g., an error handler).
2.  **Scan:** The VM retrieves the **Stack Trace** of the current process.
3.  **Descend:** It iterates *forward* (deeper) into the stack, checking each paused "Zombie Frame."
4.  **Match:** It checks if the Zombie Frame has a `TRAP_REPLY` table entry matching `^ack`.
5.  **Resume:** If found, the VM transfers execution control **back** to that Zombie Frame's Instruction Pointer (IP), passing the data arguments.

### 4\.6 Reactive Realms & Subscriptions

Realms are reified Spaces that can be assigned to variables.

* **Realms as Maps:** A Realm is fundamentally a **Map**.
    * It has a **Legend** and supports the standard **Prototype Model**.
    * You can access fields (`realm\config`), delegate to subfields (`realm@parent`), and attach methods directly to the Realm object.
    * *Distinction:* While it acts as a Map for storage/lookup, it *also* possesses the `Space` interface for Concurrency (`#`, `$`, `<>`).
* **Child Realm `<$>`**: Creates a standard Space. Signals uncaught in this realm bubble up to the creator's current space.
* **Detached Realm `<|>`**: Creates a Sandboxed Space. `Parent` is set to `World` (Root). Signals hitting the ceiling are discarded/logged.
* **Opcode:** `NEW_REALM <flags>` (Flag 0: Child, Flag 1: Detached)

The syntax `realm <> [ pattern ] -> { body }` acts as a generic lifecycle manager.

  * **Initialization:** The `SUBSCRIBE` opcode registers a daemon on the target
    Space.
  * **Arrival (Spawn):** When a Tuple (Signal/Proclamation) matches `pattern`, a
    new Green Thread is spawned.
      * **Implicit Pinning:** Variables defined in `pattern` (e.g., `who`) act
        as **Filters**. The thread only wakes for tuples matching that specific
        value.
  * **Departure (Teardown):** When a Proclamation is updated to **`___`
    (Empty)**, it is considered **Retracted**. The VM removes the tuple from
    storage and injects a special `$empty` signal into any active subscriber
    threads to trigger their cleanup handlers.

### 4\.7 Vassals (Facets & Attenuation)

Vassals (`<{}>`) are **Bi-Directional Proxies** used to secure a Realm. They
enforce the *Principle of Least Privilege*.

**Syntax:**

```rhumb
ReadOnly .= <{
    #allow .. #allow         % Pass-through (Consume)
    .#allow                  % if no changes, use prefix
    #peek :: #peek           % Pass-through (Non-destructive)
    :#peek                   % if no changes, use prefix
    #secret .. empty         % Block (Drop)
    #shout(x) .. #whisper(x) % Transform
}>
```

**Semantics:**

  * **Take (`..`):** The Vassal **consumes** the event. It stops bubbling in the Raw Realm. (Use for Admin/Overrides).
  * **Peek (`::`):** The Vassal **copies** the event to the user but lets it continue bubbling in the Raw Realm. (Use for Logging/Monitoring).

For Proclamations (`$`), `..` (Take) hides the value from subsequent rules in the
Vassal but **does not** remove it from the underlying Realm. To delete a
Proclamation from the underlying Realm, the Vassal must explicitly emit an
`$empty` signal to the Target.

**Implementation:** The `Vassal` struct implements the `Space` interface but
delegates storage to a `Target`.

```go
type Vassal struct {
    Target *Space
    Rules  *Closure // The <{}> logic
}
```

### 4\.8 Go Implementation Strategy

#### The Space Struct

```go
type Space struct {
    Parent *Space
    
    // Active Listeners (Traps)
    // Key: Signal ID. Value: Channel to wake up the process.
    Listeners sync.Map 
    
    // Persistent State (Proclamations)
    // Key: Proclamation ID. Value: Data Tuple.
    Storage sync.Map
    
    // Thread-Safety
    Lock sync.Mutex
}
```

#### The Tuple Struct

```go
type Tuple struct {
    Kind    TupleType // Signal (#), Reply (^), Proclamation ($)
    Topic   string    // The ID (e.g., "says", "present")
    Payload []Value   // The arguments
}
```

#### Optimization: The "Fast Path"

  * **Synchronous Return:** Code returning a value (unnamed signal) compiles to `RETURN`, bypassing the Tuplespace entirely for speed.
  * **Asynchronous IPC:** Only named signals (`#error`, `#log`) trigger the Tuplespace logic, preserving the "90/10" performance rule.

### 4\.9 Distributed Rhumb (Network Transparency)

Rhumb treats the network as just another Space boundary. Concurrency primitives
(`#`, `$`, `->`) work identically across local cores and remote nodes.

#### A. Networked Realms

A Realm can be bound to a Transport Layer (TCP/WebSockets).

  * **Syntax:** Configuration via Proclamation.
    ```rhumb
    node .= <$>;
    node$connect("tcp://192.168.1.5:8080");
    ```
  * **Behavior:**
      * **Signal (`node#msg`):** Serializes `msg` and sends it over the socket.
      * **Proclamation (`node$state`):** Syncs state to the remote node (CRDT-like consistency).
      * **Subscription (`node <> ...`):** deserializes incoming packets into local Tuples.

#### B. The Freezer (Serialization Engine)

The VM employs a unified serialization engine called **The Freezer** to flatten
the object graph into a binary format. It operates in two modes:

1.  **Migration Mode (Network):**
    * **Scope:** Serializes a specific Closure or Process.
    * **Sanitization:** **Strict.** If the graph contains non-transferable local
      resources (File Handles, Window Pointers), the freeze **fails** with an
      error. This prevents "It works on my machine" bugs in distributed code.
2.  **Snapshot Mode (Disk):**
    * **Scope:** Serializes the entire VM Heap and Tuplespace.
    * **Sanitization:** **Permissive.** Local resources are serialized as
      **Rehydration Instructions** (e.g., "Open file X at offset Y"). On reload,
      the VM attempts to restore them; if impossible (file missing), the handle
      becomes `___` or an error.
#### C. The Dependency Check

Before accepting a migrated process, the Remote Node validates the **Resolver
Headers**.

  * **Check:** "Do I have `{!|math|1.2.0}`?"
  * **Result:**
      * **Yes:** Accept and run.
      * **No:** Reject (or optionally request the missing library blob).

### 4\.10 Syntactic Symmetries

The Concurrency primitives exhibit powerful symmetries that allow developers to
choose between **Inline Logic** (Subscriptions) and **Reusable Logic**
(Vassals).

#### A. The Monitor Symmetry (`<>` vs `{}`)

Listening to a Realm via an empty subscription is identical to attaching a
Selector directly to the Realm (Attachment Mode).

  * **Syntax A:** `realm <> [] -> { ... }`
      * *Meaning:* "Subscribe to `realm`. Filter nothing (`[]`). Execute block for every event."
  * **Syntax B:** `realm { ... }`
      * *Meaning:* "Attach monitor to `realm`. Trap all signals bubbling up."
  * **Equivalence:** `realm <> []` $\equiv$ `realm` (as a Monitor Source).

#### B. The Filter Symmetry (Vassals vs Patterns)

Applying a Vassal to a Realm is identical to Subscribing with a Pattern. Both
act as filters on the event stream.

  * **Scenario:** We want to listen only for `#sig`.
  * **Approach A (Inline Pattern):**
    ```rhumb
    realm <> [#sig] -> { ... }
    ```
      * *Mechanism:* The **Pattern** inside the subscription performs the filtering.
  * **Approach B (Applied Vassal):**
    ```rhumb
    v .= <{ .#sig }>   % Define Vassal (Allow #sig)
    v(realm) { ... }   % Apply Vassal -> Attach Monitor
    ```
      * *Mechanism:* The **Vassal** performs the filtering before the Monitor sees it.
  * **Conclusion:** **Submaps are ephemeral Vassals.** When you write `[#sig]`,
    you are defining a temporary attenuation strategy for that specific
    subscription.
