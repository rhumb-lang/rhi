## 3\. The Three Lookup Algorithms

Rhumb resolves symbols using three distinct algorithms depending on the context of the operator used.

### 3\.1 Lexical Lookup & Assignment

Variables are resolved by searching the stack.

  * **Lookup Scope:** Current Frame $\to$ Closure (Upvalues) $\to$ Module Static.
  * **Assignment Logic (`.=` / `:=` / `^=`):**
    1.  **Search:** Traverse scopes to find an existing label.
    2.  **Hit:** If found, check current mutability.
          * If **Immutable**: Throw `WriteViolation`.
          * If **Mutable**: Update value. Apply new mutability flag (Freeze if `.=`).
    3\.  **Miss:** Create new label in **Current Frame** with specified mutability.
  * **Shadowing:** To shadow a variable (create a new local with the same name), use **Parameters** in a Function (`->`) or Immediate Function (`+>`). Parameters *always* create new labels in the new frame.

### 3\.2 Map Lookup (The Inheritance Path)

Resolves fields and methods on a specific receiver.

  * **Trigger:** Accessing a field via `\` (e.g., `point\x`) or `SEND`.
  * **Scope:** The Map and its Subfield Chain.
  * **Algorithm:**
    1.  **Local Lookup:** Query the receiver's **Legend** (Schema) for the field offset. If found, read from the Map's storage array.
    2.  **Delegation:** If not found, iterate through **Subfields** defined with the `@` prefix.
    3\.  **Recursion:** Traversal is Depth-First, Left-to-Right through the subfield chain.
    4.  **Base Binding:** During execution, `!` (Base) remains bound to the original receiver.

### 3\.3 Dynamic Lookup (The Space Path)

Resolves concurrency events and state.

  * **Trigger:** Using Space operators (`#`, `^`, `$`).
  * **Scope:** The Hierarchical Tuplespace.

#### Algorithm: Signal (`#`) - Bubble Up

1.  **Check:** Check current Space's Listeners (`<>`).
2.  **Ascend:** If not caught, move to `Parent` Space.
3\.  **Repeat:** Continue until `World` (Root). If uncaught, discard/GC.

#### Algorithm: Reply (`^`) - Drill Down (Zombie Walk)

1.  **Trigger:** A shallower frame (Caller) emits `^reply`.
2.  **Local Check:** Check the current frame's own local Reply Traps.
3\.  **Zombie Scan:** Retrieve the list of "Zombie Frames" (deeper frames that were paused/yielded to reach the current frame).
4.  **Descend:** Iterate *down* into the Zombies.
      * Check each Zombie Frame for a `TRAP_REPLY` matching the signal name.
5.  **Inject:** If a match is found, transfer execution control *back* into that Zombie Frame at the handler offset.

#### Algorithm: Proclamation (`$`) - Pin, Notify & Retract

1.  **Check Value:**
    * **If Value is Present:** Write/Update the tuple in Local Space storage.
    * **If Value is `___` (Empty):** **Remove** the tuple from Local Space storage (Retraction).
2.  **Notify:** Check Local Listeners (`<>`).
    * **On Update:** If the new value matches a pattern, spawn a handler immediately.
    * **On Retraction:** If the tuple was removed, inject the **`$empty`** signal into all active threads subscribed to that topic.
3\.  **Persist:** The state (or absence thereof) persists until the next Proclamation.


### 3\.4 Bytecode Architecture

Instruction Set is divided into four banks.

| Bank         | Purpose      | Examples                                                     |
|:-------------|:-------------|:-------------------------------------------------------------|
| **Selector** | Control Flow | `SELECT` (Pattern Match), `MATCH_STRUCT`, `JUMP`             |
| **Lexical**  | Static Scope | `LOAD_STATIC` (Module), `LOAD_LOC` (Local Var), `MATCH_BIND` |
| **Map**      | Inheritance  | `SEND`, `SELF`, `LOAD_PARENT` (`@`)                          |
| **Space**    | Concurrency  | `POST`, `INJECT`, `WRITE`, `SUBSCRIBE`                       |

### 3\.5 Native Intrinsics (Operator Mapping)

Operators are mapped to the following Native Opcodes. If the operand is a Map,
the VM attempts to find a matching **Hook Field** (surrounded by _).

#### Functions & Binding

| Operator      | Syntax   | Native Opcode | Semantics                                 |
|:--------------|:---------|:--------------|:------------------------------------------|
| **Function**  | `->`     | `OP_MAKE_FN`  | Create anonymous function                 |
| **Bind Fn**   | `!>`     | `OP_BIND_FN`  | Create method with bound `!` (Self)       |
| **Let Fn**    | `+>`     | `OP_LET_FN`   | Define and execute immediately (IIFE)     |
| **Bind**      | `!!`     | `OP_REBIND`   | Rebind an existing function to a new base |
| **Reference** | `<fn>`   | `OP_REF_FN`   | Fetch function without invoking           |
| **Curry**     | `<fn>()` | `OP_CURRY`    | Partial Application (Bind Args)           |

#### Math & Logic

| Operator    | Syntax | Native Opcode  | Hook Slot | Semantics                |
|:------------|:-------|:---------------|:----------|:-------------------------|
| **Add**     | `++`   | `OP_ADD`       | `_++_`    | Addition / Concatenation |
| **Sub**     | `--`   | `OP_SUB`       | `_--_`    | Subtraction              |
| **Mult**    | `**`   | `OP_MULT`      | `_**_`    | Multiplication           |
| **Exp**     | `^^`   | `OP_POW`       | `_^^_`    | Exponentiation           |
| **Dec Div** | `//`   | `OP_DIV_FLOAT` | `_//_`    | Floating Point Division  |
| **Int Div** | `+/`   | `OP_DIV_INT`   | `_+/_`    | Floor Division           |
| **Mod**     | `-/`   | `OP_MOD`       | `_-/_`    | Remainder                |
| **Sci Not** | `*^`   | `OP_SCI_NOT`   | `_*^_`    | Scientific Notation      |
| **Root**    | `^/`   | `OP_ROOT`      | `_^/_`    | Radication               |
| **Eq**      | `==`   | `OP_EQ`        | `_==_`    | Equality Check           |
| **Neq**     | `~~`   | `OP_NEQ`       | `_~~_`    | Inequality Check         |
| **GT**      | `>>`   | `OP_GT`        | `_>>_`    | Greater Than             |
| **LT**      | `<<`   | `OP_LT`        | `_<<_`    | Less Than                |
| **GTE**     | `>=`   | `OP_GTE`       | `_>=_`    | Greater Than or Equal    |
| **LTE**     | `<=`   | `OP_LTE`       | `_<=_`    | Less Than or Equal       |
| **And**     | `/\`   | `OP_AND`       | `_/\_`    | Logical Conjunction      |
| **Or**      | `\/`   | `OP_OR`        | `_\/_`    | Logical Disjunction      |


#### Map & Structure

| Operator     | Syntax | Native Opcode      | Semantics                                    |
|:-------------|:-------|:-------------------|:---------------------------------------------|
| **Range**    | `\|`   | `OP_RANGE`         | Create **generic** lazy pair which can become different things depending on context (1\|3 -\> [1;2;3]) |
| **Has Sub**  | `=@`   | `OP_HAS_SUBFIELD`  | Check for subfield presence                  |
| **Not Sub**  | `~@`   | `OP_NOT_HAS_SUB`   | Check for subfield absence                   |
| **Has Fld**  | `=\`   | `OP_HAS_FIELD`     | Check for field presence                     |
| **Not Fld**  | `~\`   | `OP_NOT_HAS_FLD`   | Check for field absence                      |
| **Temp Sub** | `@@`   | `OP_TEMP_SUBFIELD` | Temporary assignment scope                   |
| **Concat**   | `&&`   | `OP_CONCAT`        | Merge positional elements                    |
| **Nested**   | `\\`   | `OP_ACCESS_NESTED` | Deep search via wildcard                     |

#### Control Flow & Assignment

| Operator     | Syntax | Native Opcode   | Semantics                                                   |
|:-------------|:-------|:----------------|:------------------------------------------------------------|
| **Assgn Im** | `.=`   | `OP_ASSIGN_IMM` | Find/Create. Update Value. **Set Immutable.**               |
| **Assgn Mu** | `:=`   | `OP_ASSIGN_MUT` | Find/Create. Update Value. **Set Mutable.**                 |
| **Destruct** | `^=`   | `OP_DESTRUCT`   | Destructuring Assignment (Updates existing or creates new). |
| **If True**  | `=>`   | `OP_IF_TRUE`    | Execute if LHS is yes                                       |
| **If False** | `~>`   | `OP_IF_FALSE`   | Execute if LHS is no/empty                                  |
| **While**    | `\|>`  | `OP_WHILE`      | Loop LHS until no                                           |
| **Foreach**  | `<>`   | `OP_FOREACH`    | Iterate Positional Elements / Lifecycle                     |
| **Pipe**     | `\|\|` | `OP_PIPE`       | Functional Pipe                                             |
| **Default**  | `??`   | `OP_COALESCE`   | Return LHS unless empty                                     |
| **Match**    | `..`   | `OP_MATCH_CONS` | Select & Consume (Stop)                                     |
| **Peek**     | `::`   | `OP_MATCH_PEEK` | Select & Continue (Fallthrough)                             |

#### Space & Concurrency
All Space operations consume their operands and **must push a result** to maintain stack integrity.

| Operator        | Syntax     | Native Opcode    | Semantics                              | Stack Return               |
|:----------------|:-----------|:-----------------|:---------------------------------------|:---------------------------|
| **Signal**      | `#`        | `OP_POST`        | Emit & Suspend. Wait for Reply.        | **Reply Value** (or `___`) |
| **Reply**       | `^`        | `OP_INJECT`      | Resume Zombie Frame with Value.        | `___` (Empty)              |
| **Proclaim**    | `$`        | `OP_WRITE`       | Set State & Notify.                    | `___` (Empty)              |
| **Subscribe**   | `<>`       | `OP_SUBSCRIBE`   | Register Listener.                     | `___` (Empty)              |
| **Match Tuple** | (internal) | `OP_MATCH_TUPLE` | Efficiently checks Tuple Kind (#/^/$). | `yes` / `no`               |


#### Field Operators (Postfix `[]`)

Map Field Operators can be hooked by defining a field matching the operator name surrounded by `_`.

| Operator    | Syntax  | Native Opcode    | Hook Field | Semantics                 |
|:------------|:--------|:-----------------|:-----------|:--------------------------|
| **Append**  | `[>]`   | `OP_APPEND`      | `_>_`      | Add to end of list        |
| **Unshift** | `[<]`   | `OP_UNSHIFT`     | `_<_`      | Add to start of list      |
| **Length**  | `[#]`   | `OP_LENGTH`      | `_#_`      | Count elements            |
| **Empty?**  | `[?]`   | `OP_IS_EMPTY`    | `_?_`      | Check if empty            |
| **All Sub** | `[@]`   | `OP_ALL_SUB`     | `_@_`      | Get all subfields         |
| **All Fld** | `[*]`   | `OP_ALL_FIELDS`  | `_*_`      | Get all public field keys |
| **All Pos** | `[0]`   | `OP_ALL_POS`     | `_0_`      | Get all list items        |
| **Freeze**  | `[.]`   | `OP_FREEZE`      | `_._`      | Make immutable            |
| **Copy**    | `[:]`   | `OP_COPY`        | `_:_`      | Clone map                 |
| **Date**    | `[/]`   | `OP_COERCE_DATE` | `_/_`      | Coerce to Date            |
| **Params**  | `[$]`   | `OP_GET_PARAMS`  | `_$_`      | Get function parameters   |
| **Ctor**    | `[^]`   | `OP_GET_CTOR`    | `_^_`      | Get map constructor       |
| **Base**    | `[!]`   | `OP_GET_BASE`    | `_!_`      | Get bound base            |
| **Num**     | `[+]`   | `OP_COERCE_NUM`  | `_+_`      | Coerce to Number          |
| **Neg**     | `[-]`   | `OP_NUM_NEG`     | `_-_`      | Numerical Negate          |
| **Bool**    | `[=]`   | `OP_COERCE_BOOL` | `_=_`      | Coerce to Truth           |
| **Not**     | `[~]`   | `OP_BOOL_NEG`    | `_~_`      | Logical Negate            |
| **Spread**  | `[&]`   | `OP_SPREAD`      | `_&_`      | Slurp/Spread elements     |
| **Key**     | ``[`]`` | `OP_COERCE_KEY`  | ``_`_``    | Coerce to Key             |

### Code Example

Here is an example showing some of the operators and a hook field:

```rhumb
Point .= <(
  arg1 .= $1
  arg2 .= $2

  % this
  x := ___
  y := ___
  % is the same as
  [:x; :y] ^= [___; ___]

  % if arg1's constructor is Point
  arg1[^] == Point => (
    x := arg1\x
    y := arg1\y
  ) ~> ( % chaining => and ~> makes if/else patterns
    x := arg1 ?? 0
    y := arg2 ?? 0
  )
  % since x and y are local variables, they are only updatable through these methods
  !\set-x .= [x1] !> (x := x1)
  !\set-y .= [y1] !> (y := y1)
  % by default, #(!) is signalled if the block reaches the end without signalling
)>

% destructuring base to make interpolation cleaner
<Point>[!]\print1 .. [] !> ([x; y] ^= !; console\log("Point($x;$y)"))
% or just using interpolated base access expressions directly
<Point>[!]\print2 .. [] !> console\log("Point($(!\x);$(!\y))")
<Point>[!]@effects .= [
  % you can also use a base within a subfield by manually binding it
  print3 .. <([x; y] ^= !; console\log("Point($x;$y)"))> !! <Point>[!]

  % Operator Overloading (_==_ is the hook field label)
  _==_ .. [other] !> other\x == !\x /\ other\y == !\y
]

p1 := Point(10;15)
p2 := Point(p1)
p1 == p2 %= yes
```

### 3\.6 Selector Semantics

Selectors (`{...}`) behave differently based on the type of their Subject.

**1. Argument Supply Mode (Subject is Subroutine):** When the subject is an
anonymous subroutine `<(...)>`, the LHS of the selector acts as an argument
provider.

  * `1 .. f` : Supplies value `1` as an argument to function `f`.
  * `y .. _` : Assigns the Subject subroutine to variable `y`.

**2. Dispatch Mode (Subject is Function/Value):** When the subject is a standard
value or result, the selector acts as a switch/match block.

  * `1 .. f` : Compares Subject to `1`. If equal, executes `f`.
  * `x .. f` : Compares Subject to `x`. If equal, executes `f` (Pinning).
  * `y .. f` : Binds Subject to `y` and executes `f`.

**3. Structural Mode (Subject is Map/Tuple):** Used heavily in Concurrency
(`<>`). Matches the structure of the Subject against the LHS pattern.

  * `[x; 2] .. f` : Checks if Subject is a tuple where 2nd element is `2`. Binds 1st element to `x`.

**4. Attachment Mode (Subject is Execution Context):** When attached to a
function call (e.g., `func() { ... }`), the selector becomes a **Monitor** for
that specific activation (Frame Space). It subscribes to the lifecycle of the
call.

  * **Return:** The selector matches implicitly against the return value
    (unnamed signal).
  * **Signals (`#`):** Acts as a Trap for signals bubbling up from *inside* the
    function call. `  #err .. log ` catches errors.
  * **Replies (`^`):** Can inject replies back *down* into the running function.
    `^retry` sends data to a `TRAP_REPLY` inside.
  * **Proclamations (`$`):** Can react to state changes within the function's
    local space. If the function executes `$status("working")`, the attached
    selector can match `$status(s) .. log(s)`.


To implement the **Zombie Frame** behavior required for the Reply (`^`) system,
Rhumb cannot use a standard linear stack (like C or Java). It must use a
**Cactus Stack** (also known as a Spaghetti Stack or Parent-Pointer Tree).

This structure allows execution branches to fork, pause, and persist
independently, which is the foundation of the concurrency model.

#### 3\.6\.1 Vassals (Sub-Selectors)

When you surround a selector in `<...>` it becomes a vassal which is a special
kind of stored selector that we use for managing the boundaries of signals,
replies and proclamations
([§4.7](https://github.com/rhumb-lang/rhi/blob/main/docs/04_concurrency.md#47-vassals-facets--attenuation))
as well as for managing the boundaries of libraries
([§5.5.1](05_project_structure.md#551-initial-library-state)).

-----

### 3\.7 Memory Model: The Cactus Stack

To support **Zombie Frames** and **Resumable Replies**, the VM does not use a
contiguous block of memory for the stack. Instead, it uses a **Cactus Stack** (a
tree of linked frames allocated on the Heap).

#### 3\.7\.1 Structure

  * **Heap Allocation:** Every `CallFrame` is a struct allocated on the Go Heap.
  * **Parent Pointers:** Each frame holds a pointer to its **Caller** (`Parent`).
  * **The Tree:** Because multiple closures or concurrent processes can be
    spawned from the same context, a single Parent Frame may have multiple
    active Child Frames (branches), giving the stack a cactus-like shape.

<!-- end list -->

```go
type CallFrame struct {
    Parent   *CallFrame  // Link to the caller (Shallow)
    Closure  *Closure    // The code being executed
    IP       int         // Instruction Pointer
    Locals   []Value     // Local variables
    State    FrameState  // Active, Zombie, or Dead
}
```

#### 3\.7\.2 Lifecycle & Zombies

  * **Call (`OP_CALL`):** Creates a new Frame, links `Parent = CurrentFrame`, and sets `CurrentFrame = NewFrame`.
  * **Return (`OP_RETURN`):**
    1.  The VM marks the Current Frame as **Zombie** (Dormant).
    2.  It sets `CurrentFrame = CurrentFrame.Parent`.
    3.  **Crucial:** The returned frame is **not deallocated**. It remains reachable via any **Reply Traps** or **Closures** that captured it.
  * **Garbage Collection:** The Go GC handles memory. If a Zombie Frame is no longer referenced by any active Process, Listener, or Child, it is automatically swept.

#### 3\.7\.3 The "Drill Down" Mechanism

The Cactus Stack enables the **Reply (`^`)** operator to traverse "forward" into history.

1.  **Lookup:** When `^reply` is issued, the VM inspects the **Zombie List** associated with the current process.
2.  **Traversal:** It walks down the linked list of Zombie Frames that were "popped" to reach the current state.
3.  **Resume:** If a matching trap is found in a Zombie, the VM creates a new branch (Green Thread) resuming execution at that Zombie's IP, effectively "forking" the history.

### 3\.8 Subroutine Semantics

Rhumb defines specific roles for Logic units, separating the executable code from its interface and invocation state.

1.  **Subroutine `<(...)>`**: The **Fundamental Unit of Code Execution**.
      * It is a **Reference** to logic (Bytecode Chunk).
      * It is **Anonymous** and has no inherent parameter names (uses positional `$1`, `$2`).
      * **Invocation:** Accessing a label bound to a subroutine *executes* it immediately.
          * `foo` executes `foo`.
          * `foo()` executes `foo` with empty arguments (if any were missing).
2.  **Function `[...] -> (...)`**: A syntactic construct that combines a **Submap** (Parameters) with a **Subroutine**.
      * `->` is a **Constructor Operator**.
      * It converts the LHS `[]` into a Parameter Schema (Submap).
      * It converts the RHS `()` into a Subroutine.
      * **Invocation:** `foo(1)` binds `1` to the Submap labels, then executes the Subroutine.

#### 3\.8\.1 Loose Argument Policy

When a function is called (`foo(...)`):

  * **Missing Args:** If fewer arguments are provided than parameters defined in the submap, the remaining parameters are bound to **`___` (Empty)**.
  * **Extra Args:** If more arguments are provided than parameters, the extra values are ignored by the named binding but remain accessible via the variadic argument operator **`$`** (e.g., `$0` for all args, or `$N` for the Nth).

#### 3\.8\.2 Referencing & Currying

  * **Referencing (No Call):** The **Only Way** to pass a subroutine or function without executing it is to wrap it in Angle Brackets `<...>`.
      * `foo` $\rightarrow$ Executes.
      * `<foo>` $\rightarrow$ Pushes the Function Object onto the stack.
  * **Partial Application (Currying):** Syntax sugar for creating a new closure.
      * `<foo>(1)` $\rightarrow$ References `foo`, applies `1`, and returns a **New Function** (Closure) waiting for the remaining arguments. It does *not* execute.

### 3\.9 Dynamic Type Interactions

Rhumb enforces a strict, bi-directional type promotion strategy. The order of
operands does not change the **Result Type** (e.g., `Int ++ Float` and `Float ++
Int` both yield `Float`).

#### 3\.9\.1 The Numeric Hierarchy
When mixing standard numeric types, the result is promoted to the type with the greatest expressivity to prevent data loss.

1.  **Decimal** (Highest Precision)
2.  **Float** (Highest Range)
3.  **Integer** (Lowest)

*Rule:* `Lower + Higher` $\rightarrow$ `Higher`.

#### 3\.9\.2 The Interaction Matrix
The following table defines the result type for binary arithmetic (`++`, `--`).
* **Rows/Cols:** The input types.
* **Cells:** The output type.
* **Text:** Not included. Text operations use `&&` (Concatenate) and do not respond to math operators.

| `++` / `--` | **Integer** | **Float** | **Decimal** | **Duration** | **DateTime** |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Integer** | `Integer` | `Float` | `Decimal` | `Duration` | `DateTime`† |
| **Float** | `Float` | `Float` | `Decimal` | `Duration` | `DateTime`† |
| **Decimal** | `Decimal` | `Decimal` | `Decimal` | `Duration` | `DateTime`† |
| **Duration** | `Duration` | `Duration` | `Duration` | `Duration` | `DateTime` |
| **DateTime** | `DateTime`† | `DateTime`† | `DateTime`† | `DateTime` | `Duration`‡|

#### 3\.9\.3 Scalar Interpretation Rules
When mixing Time types with Scalars (Integer/Float/Decimal), the Scalar is implicitly converted to a **Duration** before the operation proceeds.

* **Integer:** Treated as **Milliseconds**.
    * `Time + 500` $\rightarrow$ `Time + 500ms`
* **Float / Decimal:** Treated as **Seconds**.
    * The integer part becomes Seconds.
    * The fractional part is converted to Milliseconds (Truncated precision).
    * `Time + 1.5` $\rightarrow$ `Time + 1500ms`

#### 3\.9\.4 Special Logic (†/‡)

**†: Date/Scalar Interaction**
  * **Addition (`++`):** Commutative. Shifts the Date by the scalar amount.
  * **Subtraction (`--`):**
      * `Date -- Scalar` $\rightarrow$ **DateTime** (Rewind).
      * `Scalar -- Date` $\rightarrow$ **Error** (Cannot subtract a Point from a Scalar).

**‡: Date/Date Interaction**
  * **Subtraction (`--`):** `Date - Date` $\rightarrow$ **Duration** (The magnitude between two points).
  * **Addition (`++`):** `Date + Date` $\rightarrow$ **Error** (Geometrically meaningless).

#### 3\.9\.5 Multiplication & Division (`**`, `//`, `+/`, `-/`)

These operators strictly follow the **Numeric Hierarchy**.
* `DateTime` cannot be multiplied or divided.
* **Duration Scaling:** `Duration` can be multiplied or divided by a **Numeric** (scalar) to scale the time vector.
    * `Duration ** Int` $\rightarrow$ `Duration`
    * `Duration // Float` $\rightarrow$ `Duration` (with ms truncation)