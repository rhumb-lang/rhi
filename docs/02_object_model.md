## 2\.The Map (Object) Model

Rhumb uses a high-performance **Prototype** system inspired by the [Self programming language](https://selflanguage.org).

### 2\.1 Terminology

  * **Map:** What is called an "object" in other languages. It acts as a container for data and behavior.
  * **Field:** What is called a "slot" in other languages. Every field has a **Name** and a **Value**.
  * **Name:** The identifier for a field.
      * Can be: **Label**, **Number**, **Text**, **Key**, **Date**.
      * **Cannot** be a Map.
  * **Label:** An **unquoted identifier** (e.g., `x`, `count`).
      * Labels are used in the Left-Hand Side (LHS) of:
          * Assignment (`x .= 1`)
          * Map Field Operations (`map\x`)
          * Selector Rules (`x .. f`)
          * Space Operations (`space#signal`)
  * **Routine:** A sequence of expressions enclosed in `()`. Evaluated immediately.
  * **Subroutine:** A **Reference** to a routine `<()>`. Stored executable logic (Code).
  * **Submap:** A **Reference** to a Map literal `<[]>`. Represents a **Schema** or **Prototype Shape**.
      * **As Constructor:** When invoked, it creates a new Map with fields populated by arguments.
      * **As Parameters:** When used with `->`, it defines the local variable bindings for a function.

### 2\.2 The Universal Legend

We separate **State** (Map) from **Schema** (Legend). This allows thousands of
objects with the same structure to share a single description, saving memory and
enabling fast property lookups.

  * **Immutability:** The Legend stores whether a field was defined with `.`
    (Immutable) or `:` (Mutable).
  * **Transitions:** The Legend acts as a node in a state machine. Adding a
    field triggers a transition to a new or existing Legend.

<!-- end list -->

```go
type FieldKind uint8
const (
    FieldImmutable FieldKind = iota // Defined via '.'
    FieldMutable                    // Defined via ':'
)

type TransitionDesc struct {
    FieldName string
    Next      *Legend
}

type Legend struct {
    Kind        LegendType // Map, String, Dictionary
    Fields      []FieldDesc // Linear scan (Map Mode)
    Lookup      map[string]int // Hash map (Dictionary Mode)
    Transitions []TransitionDesc // Hidden Class transition tree
}
```

#### The Transition Algorithm

When a field `new_field` is added to a Map currently using `Legend A`:

1.  **Check:** The VM inspects `Legend A.Transitions` to see if a transition for
    `new_field` already exists.
2.  **Hit (Reuse):** If found, it points to `Legend B`. The Map swaps its Legend
    pointer to `Legend B` and extends its storage array. No new allocation
    occurs for the schema.
3.  **Miss (Branch):** If not found:
      * The VM allocates a new `Legend B` (Clone of A + `new_field`).
      * It records a new transition in `Legend A`: `new_field` $\rightarrow$ `Legend B`.
      * The Map swaps to `Legend B`.
4.  **Convergence:** If 1,000 empty maps all execute `m\x := 1; m\y := 2`, they
    will all end up sharing the **exact same** `Legend C` (via the path `Empty`
    $\rightarrow$ `Legend A (has x)` $\rightarrow$ `Legend C (has x, y)`).

#### Dictionary Mode (Bailout)

If the number of fields exceeds a defined threshold (e.g., 32), or if fields are
added/removed in a non-deterministic order that causes excessive branching, the
VM hydrates the `Lookup` map and detaches from the transition tree ('Dictionary
Mode') for O(1) access at the cost of memory.

### 2\.3 The Value Struct (Primitives)

Primitives are stack-allocated via a discriminated union. All bit-packed types utilize the `Integer` (int64) slot to avoid heap allocation.

  * **Integers:** Stored as standard `int64`.
  * **Floats:** Stored as `float64` (standard IEEE 754).
  * **Text:** Uses Go's native `string` (pointer + length).
  * **Range:** A Lazy Iterator struct `start|end` (stored as Object reference or packed if small).

#### Arbitrary Precision (Heap Allocated)

  * **Decimal:** High-precision fixed-point numbers backed by **`apd.Decimal`**.
      * **Syntax:** Distinguished from standard floats by a **Leading Zero**.
          * `01.5` $\rightarrow$ Decimal 1.5
          * `00.5` $\rightarrow$ Decimal 0.5 (Two leading zeros required if integer part is 0).
      * **Precision Control:** The `.-` suffix captures explicit ones-place precision.
          * `01.-` $\rightarrow$ Decimal 1 (Precision: Ones place).
          * `00.-` $\rightarrow$ Decimal 0 (Precision: Ones place).

#### Bit-Packed Primitives (Using `int64` slot)

  * **DateTime:** Stored as **Unix Milliseconds** (Absolute).
      * **Concept:** Represents a specific point in time.
      * **Range:** \~292 million years from 1970.
      * **Syntax:**
          * **Full:** `2025/01/01@12:00:00`
          * **Date-Only:** `2025/01/01` (Time defaults to midnight).
          * **Time-Only:** `00:05:00` (Date defaults to Unix Epoch: 1970-01-01).
  * **Duration:** Stored as **Milliseconds** (Relative).
      * **Concept:** Represents a magnitude of time (Vector).
      * **Syntax:** There is no distinct literal for Duration. Instead, use the **Unary Plus (`+`)** or **Unary Minus (`-`)** operator on a DateTime literal to cast it to a Duration.
          * **Time Vectors:** `+00:05:00` (5 Minutes).
          * **Calendar Vectors:** `+0001/00/00` (1 Year), `+0000/01/00` (1 Month), `+0000/00/01` (1 Day).
          * **Composite Vectors:** `+0000/00/01 @ 12:00:00` (1 Day and 12 Hours).
          * **Negative Vectors:** `-00:05:00` (Negative 5 Minutes).
      * **Arithmetic:**
          * `Date ++ Duration` $\rightarrow$ `Date` (Time Travel).
          * `Date -- Date` $\rightarrow$ `Duration` (Difference).
          * `Date ++ Date` $\rightarrow$ **Error** (Nonsensical).
  * **Version:** Stored as **Packed SemVer**.
      * **Bits 63-48:** Major (16 bits)
      * **Bits 47-32:** Minor (16 bits)
      * **Bits 31-0:** Patch (32 bits)
      * **Suffixes:** Prerelease and Build metadata are stored as auxiliary text pointers.
  * **Key:** Stored as **Interned Global ID**.
      * Comparison (`k1 == k2`) is a fast integer check. Keys are never garbage collected during the process lifetime.

### 2\.4 The Empty Value (`___`)

The concept of `nil` or `null` is represented by **`___` (Triple Underscore)**.

  * **Behavior:** Represents the absence of a value.
  * **Semantics:** Any label not yet defined is considered `empty` (`___`) by default.
  * **Logic:** `___` is falsy in boolean expressions.
  * **Retraction:** Assigning `___` to a Proclamation (`realm$state(___)`)
    removes the state from the Tuplespace.

#### Summary of the Retraction Mechanism
* **User Action:** `realm$topic(___)`
* **VM Action:**
    1.  Deletes the `$topic` tuple from `realm.Storage`.
    2.  Finds all threads subscribed to `$topic`.
    3.  Sends `$empty` signal to those threads.
* **Thread Action:** Executes `empty .. log("Deleted")` block and terminates.

### 2\.5 The Panic Value (`***`)

The concept of a Fatal Error or Crash is represented by **`***` (Triple Asterisk)**.

  * **Behavior (Standalone):** If the interpreter executes `***` as a standalone
    expression, it triggers a **VM Panic** (Crash/Exit).
  * **Behavior (Signal):** It is the standard label for Error Signals.
      * Syntax: `#***(code; message; data)`
      * Use Case: `{!|sys}` routines emit this signal on failure.
  * **Handling:** Users can trap errors by adding a selector pattern for it.
    ```rhumb
    risky-op {
        #***(c; m; d) .. log("Error $c: $m")
    }
    ```

### 2\.6 Privacy & Encapsulation

Rhumb uses **Capability-based Privacy** via **Keys** (`` ` ``).

  * **Public Fields:** Defined with Text/Label names. Accessible by anyone.
  * **Private Fields:** Defined with Key names. Accessible only by scopes
    holding the Key object.
  * **Reflection Safety:** The All Fields operator `[*]` **ignores** Key fields.
    It only returns Text/Label names.

### 2\.7 Hybrid Storage (Fields vs. Elements)

Maps serve as both "Objects" (named fields) and "Lists" (positional elements).

  * **Unified Mechanism:** Internally, positional elements are treated as Fields
    where the **Name** is a **Number**.
  * **Indexing:** Positional elements use **1-based** indexing. The index `0` is
    reserved to represent the aggregate of all positional elements.
  * **Separation of Concern (Operators):**
      * **`[*]` (All Fields):** Returns a list of **Text** labels only (keys).
        It ignores Keys (`` ` ``) and Numbers.
      * **`[0]` (All Positional):** Returns a new Map containing only the fields
        with **Number** names (elements).
  * **Iteration (`<>`):** By default, the Foreach operator iterates over
    **Positional Elements** (1..N). To iterate over fields, you must explicitly
    apply `[*]` first (e.g., `map[*] <> key -> ...`).

### 2\.8 Slurp & Spread Semantics (`&`)

The Ampersand operator (`&`) performs dual roles contextually, allowing flexible
list manipulation.

#### 1\. Spread (Construction)

When used inside a Map Literal `[...]`, it "spreads" the contents of a list into
the outer structure.

  * `list .= [2; 3]`
  * `full .= [1; &list; 4]` $\rightarrow$ `[1; 2; 3; 4]`

#### 2\. Slurp (Destructuring)

When used inside a Destructuring assignment `^=` or a Parameter Submap
`<[...]>`, it "slurps" remaining positional elements into a new List. Rhumb
supports **Maximum Slurping Power**, meaning the `&` can be placed anywhere.

  * **Tail Slurp:** `[.first; .&rest] ^= [1; 2; 3]` $\rightarrow$ `first=1`, `rest=[2; 3]`
  * **Head Slurp:** `[.&initial; .last] ^= [1; 2; 3]` $\rightarrow$ `initial=[1; 2]`, `last=3`
  * **Middle Slurp:** `[.top; .&mid; .bot] ^= [1; 2; 3; 4]` $\rightarrow$ `top=1`, `mid=[2; 3]`, `bot=4`

### 2\.9 Canonical String Representation

To support the "Blessed Output" testing strategy (`%=`), Rhumb defines a strict
canonical string representation for all values. This ensures that complex
objects can be asserted against a stable text format.

#### Primitives & References

  * **Numbers:** Printed in their simplest decimal form (e.g., `10`, `3.14`).
  * **Text:** Printed wrapped in single quotes (e.g., `'hello'`).
  * **Empty:** Printed as `___`.
  * **Subroutine References:** Printed with angle brackets surrounding the label
    (e.g., `<foo>`, `<print>`).

#### DateTime & Durations

  * **DateTime:** Printed using **Zero-Suppression** to match literal syntax.
      * **Midnight Suppression:** If the time is `00:00:00.000`, the time part is hidden.
          * `2025/01/01@00:00:00` $\equiv$ `2025/01/01`
      * **Epoch Suppression:** If the date is `1970/01/01`, the date part is hidden.
          * `1970/01/01@12:30:00` $\equiv$ `12:30:00`
      * **Milliseconds:** If 0, they are hidden.
      * **Zero Value:** If the value is exactly `0` (Epoch at Midnight), it prints as the Date: `1970/01/01`.
  * **Duration:** Printed using **Waterfall Normalization**.
      * **Algorithm:** The total milliseconds are distributed into units starting from the largest (Year) down to the smallest.
          * **1 Year** = 365 Days
          * **1 Month** = 30 Days
          * **1 Day** = 24 Hours
      * **Zero-Suppression:**
          * **Date Part:** If Years, Months, and Days are all 0, the date component (`YYYY/MM/DD`) is hidden.
          * **Time Part:** If Hours, Minutes, and Seconds are all 0, the time component (`HH:MM:SS`) is hidden.
          * **Milliseconds:** If 0, they are hidden (`.000` is removed).
      * **Examples:**
          * `366 Days` $\equiv$ `+0001/00/01` (Date only)
          * `1 Day, 1 Hour` $\equiv$ `+0000/00/01 @ 01:00:00` (Mixed)
          * `500ms` $\equiv$ `+00:00:00.500` (Time only)


#### Maps & Objects

Maps are printed as a bracketed list `[...]`. The order and format of fields are
governed by the following rules:

1.  **Positional First:** All positional elements (indices `1..N`) are printed
    first, displaying their **Values**.
2.  **Named Fields Append:** Non-positional fields are appended after the last
    positional element in **Insertion Order**.
3.  **Name-Only Display:** For non-positional fields, **only the Field Name** is
    shown (not the value).
4.  **Mutability Prefix:** Named fields are prefixed to indicate their
    mutability:
      * **Immutable (`.`):** e.g., `.x`
      * **Mutable (`:`):** e.g., `:count`
      * **Immutable Subfield (`.@`):** e.g., `.@base`
      * **Mutable Subfield (`:@`):** e.g., `:@traits`
5.  **Visibility Filter:**
      * **Text/Labels:** Shown.
      * **Dates/Versions:** Shown.
      * **Keys (`` ` ``):** **Hidden.** Private capability keys are never
        included in the string representation.

**Example:**

```rhumb
data .= [
  10; 20                % Positional 1, 2
  x :: 100              % Named 'x' (Mutable)
  config .. [a: 1]      % Named 'config' (Immutable)
  @logger .. <Log>      % Subfield 'logger' (Immutable)
  `secret :: "hidden"   % Key (Private)
]

% String Representation:
data %= "[10; 20; :x; .config; .@logger]"
```