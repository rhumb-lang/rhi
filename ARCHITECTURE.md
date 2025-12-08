# ARCHITECTURE.md

**Project:** Rhumb Programming Language
**Architecture:** Image-Based, Prototype-Oriented, Reactive Tuplespace
**Host Implementation:** Go (Golang)
**UI Strategy:** Wayland-Native (Gio/OpenGL), Stacked Window Manager

-----

## 1\. Core Philosophy

Rhumb is designed to induce a new "mindset" in the developer.

  * **Alien Syntax:** Heavily symbol-based to break muscle memory.
  * **Code is Data:** The canonical representation is an Abstract Syntax Tree (AST) stored in binary/serialized formats.
  * **Localized View:** The IDE projects code into the user's native language (e.g., `yes` vs `oui`), but underlying logic is language-agnostic.
  * **Three-Path Lookup:** Variables resolve via Lexical (Static), Inheritance (Map), or Dynamic (Space) paths.

-----

## 2\. Directory Structure & Code Organization

The codebase follows a "Small File" philosophy with strict separation between Runtime Logic and UI Presentation.

```text
/rhumb
├── go.mod                # Dependencies
├── Makefile              # Build scripts (ANTLR, Tests)
├── /cmd
│   ├── /rhi              # Rhumb Interpreter - CLI Runner / REPL
│   └── /rhide            # Rhumb IDE - Graphical Environment (Wayland)
├── /api                  # Bridge interfaces between UI and VM
├── /internal             # Core Logic (Private)
│   ├── /grammar          # ANTLR .g4 source & generated Go files
│   ├── /ast              # Abstract Syntax Tree definitions
│   ├── /compiler         # Source -> Bytecode transformation
│   ├── /visitor          # ANTLR Visitor Implementation (Bytecode Gen)
│   ├── /map              # Universal Legend & Primitives
│   ├── /vm               # Bytecode Interpreter
│   ├── /space            # Concurrency & Tuplespace
│   └── /storage          # Twin-File IO (.rnode/.rlabel)
├── /ui                   # Desktop Environment Logic
│   ├── /wm               # Window Manager (Stacking/Tabs logic)
│   ├── /views            # The "Apps" (Library, Book, Session)
│   └── /editor           # Code Projection Engine
└── /test                 # End-to-End Tests
    ├── /vm_spec          # .rnode files for opcode verification
    └── /scenarios        # Full system integration tests
```

### 2.1 Core Logic Sub-Folders (`/internal`)

  * **`/internal/grammar`**: Holds the ANTLR4 definitions.
  * **`/internal/visitor`**: **The Code Generator.**
      * `base_visitor.go`: Setup and common helpers.
      * `visit_selector.go`: Generates `SELECT`, `JUMP`.
      * `visit_lexical.go`: Generates `LOAD_LOC`, `STORE_LOC`.
      * `visit_map.go`: Generates `SEND`, `SELF`, `DELEGATE`.
      * `visit_space.go`: Generates `POST`, `SUBSCRIBE`, `NEW_REALM`.
      * `visit_math.go`: Generates Native Intrinsics (`+/`, `//`, `>>`).
      * `visit_proclamation.go`: Generates `WRITE` ($) logic.
      * `visit_function.go`: Generates `MAKE_FN`, `BIND_FN`, `LET_FN`, `CURRY_FN`.
  * **`/internal/map`**: Implements the "Universal Legend" and Primitives.
  * **`/internal/vm`**: The Bytecode Interpreter.
  * **`/internal/space`**: The Concurrency Model.
  * **`/internal/storage`**: The "Twin-File" System.

### 2.2 UI Sub-Folders (`/ui`)

  * **`/ui/wm` (Window Manager)**: Handles the "Physics" of the desktop.
  * **`/ui/views` (The Content)**: The actual tools (Library, Book, Tools, Session).
  * **`/ui/editor` (The Code Projection)**: Rendering AST to text and handling input.

-----

## 3\. The Three Lookup Algorithms

Rhumb resolves symbols using three distinct algorithms depending on the context of the operator used.

### 3.1 Lexical Lookup & Assignment

Variables are resolved by searching the stack.

  * **Lookup Scope:** Current Frame $\to$ Closure (Upvalues) $\to$ Module Static.
  * **Assignment Logic (`.=` / `:=` / `^=`):**
    1.  **Search:** Traverse scopes to find an existing label.
    2.  **Hit:** If found, check current mutability.
          * If **Immutable**: Throw `WriteViolation`.
          * If **Mutable**: Update value. Apply new mutability flag (Freeze if `.=`).
    3.  **Miss:** Create new label in **Current Frame** with specified mutability.
  * **Shadowing:** To shadow a variable (create a new local with the same name), use **Parameters** in a Function (`->`) or Immediate Function (`+>`). Parameters *always* create new labels in the new frame.

### 3.2 Map Lookup (The Inheritance Path)

Resolves fields and methods on a specific receiver.

  * **Trigger:** Accessing a field via `\` (e.g., `point\x`) or `SEND`.
  * **Scope:** The Map and its Subfield Chain.
  * **Algorithm:**
    1.  **Local Lookup:** Query the receiver's **Legend** (Schema) for the field offset. If found, read from the Map's storage array.
    2.  **Delegation:** If not found, iterate through **Subfields** defined with the `@` prefix.
    3.  **Recursion:** Traversal is Depth-First, Left-to-Right through the subfield chain.
    4.  **Base Binding:** During execution, `!` (Base) remains bound to the original receiver.

### 3.3 Dynamic Lookup (The Space Path)

Resolves concurrency events and state.

  * **Trigger:** Using Space operators (`#`, `^`, `$`).
  * **Scope:** The Hierarchical Tuplespace.

#### Algorithm: Signal (`#`) - Bubble Up

1.  **Check:** Check current Space's Listeners (`<>`).
2.  **Ascend:** If not caught, move to `Parent` Space.
3.  **Repeat:** Continue until `World` (Root). If uncaught, discard/GC.

#### Algorithm: Reply (`^`) - Drill Down (Zombie Walk)

1.  **Trigger:** A shallower frame (Caller) emits `^reply`.
2.  **Local Check:** Check the current frame's own local Reply Traps.
3.  **Zombie Scan:** Retrieve the list of "Zombie Frames" (deeper frames that were paused/yielded to reach the current frame).
4.  **Descend:** Iterate *down* into the Zombies.
      * Check each Zombie Frame for a `TRAP_REPLY` matching the signal name.
5.  **Inject:** If a match is found, transfer execution control *back* into that Zombie Frame at the handler offset.

#### Algorithm: Proclamation (`$`) - Pin & Notify

1.  **Write:** Pin directly to Local Space storage.
2.  **Notify:** Check Local Listeners (`<>`). If any match the new state, spawn a handler immediately.
3.  **Persist:** Remains until explicit removal.

-----

## 4\. Storage & Source (The "Babel" Layer)

Rhumb uses a **Semantic Storage** model. Source code is never plain text.

### 4.1 The "Twin-File" Persistence

To support Git merging and Localization, every module consists of:

1.  **Logic Node (`.rnode`):** Contains the AST using immutable IDs.
      * **Logic:** Canonical OpCodes e.g., `@OP_ADD`.
      * **Values:** Booleans are stored as `1` or `0`.
2.  **Translation Map (`.rlabel`):** Maps IDs and Values to human-readable strings.
      * **Localization:** Supports full matrix (Afrikaans `ja` to Zulu `yebo`).
      * `en-US`: `$x9A` -\> "velocity", `1` -\> "yes"

### 4.2 Dependency Resolution

External dependencies use a **Resolver Protocol** in the file header.

  * **Syntax:** `{ RESOLVER | PATH | VERSION }`
  * **Standard Lib:** `{!|math|1.2.0}`
  * **git:** `{git|https://path.to/repo|0.1.0}`

-----

## 5\. The Map Model (Self-Style)

Rhumb uses a high-performance **Prototype** system.

### 5.1 Terminology

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

### 5.2 The Universal Legend

We separate **State** (Map) from **Schema** (Legend).

**Immutability:** The Legend stores whether a field was defined with:
  * `.` (Immutable), or
  * `:` (Mutable).

<!-- end list -->

```go
type FieldKind uint8
const (
    FieldImmutable FieldKind = iota // Defined via '.'
    FieldMutable                    // Defined via ':'
)

type Legend struct {
    Kind        LegendType // Map, String, Dictionary
    Fields      []FieldDesc // Linear scan (Map Mode)
    Lookup      map[string]int // Hash map (Dictionary Mode)
    Transitions []TransitionDesc // Hidden Class transition tree
}
```

The `Lookup` map is `nil` by default. The VM performs a linear scan on `Fields`. If
the number of fields exceeds a defined threshold (e.g., 32), the VM hydrates the
`Lookup` map ('Dictionary Mode') for O(1) access. Transitions ensure that 'Fast
Mode' legends remain small.

### 5.3 The Value Struct (Primitives)

Primitives are stack-allocated via a discriminated union. All bit-packed types utilize the `Integer` (int64) slot to avoid heap allocation.

  * **Integers:** Stored as standard `int64`.
  * **Floats:** Stored as `float64` (separate slot).
  * **Text:** Uses Go's native `string` (pointer + length).
  * **Range:** A Lazy Iterator struct `start|end` (stored as Object reference or packed if small).

#### Bit-Packed Primitives (Using `int64` slot)

  * **Date:** Stored as **Unix Nanoseconds**.
      * `int64` range allows for \~292 years from 1970.
  * **Version:** Stored as **Packed SemVer**.
      * **Bits 63-48:** Major (16 bits, max 65,535)
      * **Bits 47-32:** Minor (16 bits, max 65,535)
      * **Bits 31-0:** Patch (32 bits, max 4,294,967,295)
  * **Key:** Stored as **Interned Global ID**.
      * When a Key `` `id `` is created, the VM checks a global symbol table.
      * If unique, it is assigned a monotonic `int64` ID.
      * Comparison (`k1 == k2`) is a fast integer check. Keys are never garbage collected during the process lifetime.

### 5.4 The Empty Value (`___`)

The concept of `nil` or `null` is represented by **`___` (Triple Underscore)**.

  * **Behavior:** Represents the absence of a value.
  * **Semantics:** Any label not yet defined is considered `empty` (`___`) by default.
  * **Logic:** `___` is falsy in boolean expressions.

### 5.5 Privacy & Encapsulation

Rhumb uses **Capability-based Privacy** via **Keys** (`` ` ``).

  * **Public Fields:** Defined with Text/Label names. Accessible by anyone.
  * **Private Fields:** Defined with Key names. Accessible only by scopes holding the Key object.
  * **Reflection Safety:** The All Fields operator `[*]` **ignores** Key fields. It only returns Text/Label names.

### 5.6 Hybrid Storage (Fields vs. Elements)

Maps serve as both "Objects" (named fields) and "Lists" (positional elements).

  * **Unified Mechanism:** Internally, positional elements are treated as Fields where the **Name** is a **Number**.
  * **Indexing:** Positional elements use **1-based** indexing. [cite_start]The index `0` is reserved to represent the aggregate of all positional elements[cite: 1097, 1200].
  * **Separation of Concern (Operators):**
      * **`[*]` (All Fields):** Returns a list of **Text** labels only (keys). It ignores Keys (`` ` ``) and Numbers.
      * **`[0]` (All Positional):** Returns a new Map containing only the fields with **Number** names (elements).
  * **Iteration (`<>`):** By default, the Foreach operator iterates over **Positional Elements** (1..N). To iterate over fields, you must explicitly apply `[*]` first (e.g., `map[*] <> key -> ...`).

-----

## 6\. The Virtual Machine (RhumbVM)

A custom Bytecode Interpreter written in Go.

### 6.1 Bytecode Architecture

Instruction Set is divided into four banks.

| Bank         | Purpose      | Examples                                                     |
|:-------------|:-------------|:-------------------------------------------------------------|
| **Selector** | Control Flow | `SELECT` (Pattern Match), `MATCH_STRUCT`, `JUMP`             |
| **Lexical**  | Static Scope | `LOAD_STATIC` (Module), `LOAD_LOC` (Local Var), `MATCH_BIND` |
| **Map**      | Inheritance  | `SEND`, `SELF`, `LOAD_PARENT` (`@`)                          |
| **Space**    | Concurrency  | `POST`, `INJECT`, `WRITE`, `SUBSCRIBE`                       |

### 6.2 Native Intrinsics (Operator Mapping)

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
| **Range**    | `\|`   | `OP_RANGE`         | Create inclusive **lazy** sequence (1\|3 -\> [1;2;3]) |
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

| Operator      | Syntax | Native Opcode  | Semantics                  | Stack Return  |
|:--------------|:-------|:---------------|:---------------------------|:--------------|
| **Signal**    | `#`    | `OP_POST`      | Emit Event (Bubble Up)     | `___` (Empty) |
| **Reply**     | `^`    | `OP_INJECT`    | Inject Event (Drill Down)  | `___` (Empty) |
| **Proclaim**  | `$`    | `OP_WRITE`     | Set State (Persistent)     | `___` (Empty) |
| **Subscribe** | `<>`   | `OP_SUBSCRIBE` | Register Reactive Listener | `___` (Empty) |

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

### 6.3 Selector Semantics

Selectors (`{...}`) behave differently based on the type of their Subject.

**1. Argument Supply Mode (Subject is Subroutine):**
When the subject is an anonymous subroutine `<(...)>`, the LHS of the selector acts as an argument provider.

  * `1 .. f` : Supplies value `1` as an argument to function `f`.
  * `y .. _` : Assigns the Subject subroutine to variable `y`.

**2. Dispatch Mode (Subject is Function/Value):**
When the subject is a standard value or result, the selector acts as a switch/match block.

  * `1 .. f` : Compares Subject to `1`. If equal, executes `f`.
  * `x .. f` : Compares Subject to `x`. If equal, executes `f` (Pinning).
  * `y .. f` : Binds Subject to `y` and executes `f`.

**3. Structural Mode (Subject is Map/Tuple):**
Used heavily in Concurrency (`<>`). Matches the structure of the Subject against the LHS pattern.

  * `[x; 2] .. f` : Checks if Subject is a tuple where 2nd element is `2`. Binds 1st element to `x`.

**4. Attachment Mode (Subject is Execution Context):**
When attached to a function call (e.g., `func() { ... }`), the selector becomes a **Monitor** for that specific activation (Frame Space). It subscribes to the lifecycle of the call.

  * **Return:** The selector matches implicitly against the return value (unnamed signal).
  * **Signals (`#`):** Acts as a Trap for signals bubbling up from *inside* the function call. `  #err .. log ` catches errors.
  * **Replies (`^`):** Can inject replies back *down* into the running function. `^retry` sends data to a `TRAP_REPLY` inside.
  * **Proclamations (`$`):** Can react to state changes within the function's local space. If the function executes `$status("working")`, the attached selector can match `$status(s) .. log(s)`.


To implement the **Zombie Frame** behavior required for your Reply (`^`) system, Rhumb cannot use a standard linear stack (like C or Java). It must use a **Cactus Stack** (also known as a Spaghetti Stack or Parent-Pointer Tree).

This structure allows execution branches to fork, pause, and persist independently, which is the foundation of your concurrency model.

Here is the detailed **Section 6.4** to add to your `ARCHITECTURE.md`.

-----

### 6.4 Memory Model: The Cactus Stack

To support **Zombie Frames** and **Resumable Replies**, the VM does not use a contiguous block of memory for the stack. Instead, it uses a **Cactus Stack** (a tree of linked frames allocated on the Heap).

#### A. Structure

  * **Heap Allocation:** Every `CallFrame` is a struct allocated on the Go Heap.
  * **Parent Pointers:** Each frame holds a pointer to its **Caller** (`Parent`).
  * **The Tree:** Because multiple closures or concurrent processes can be spawned from the same context, a single Parent Frame may have multiple active Child Frames (branches), giving the stack a cactus-like shape.

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

#### B. Lifecycle & Zombies

  * **Call (`OP_CALL`):** Creates a new Frame, links `Parent = CurrentFrame`, and sets `CurrentFrame = NewFrame`.
  * **Return (`OP_RETURN`):**
    1.  The VM marks the Current Frame as **Zombie** (Dormant).
    2.  It sets `CurrentFrame = CurrentFrame.Parent`.
    3.  **Crucial:** The returned frame is **not deallocated**. It remains reachable via any **Reply Traps** or **Closures** that captured it.
  * **Garbage Collection:** The Go GC handles memory. If a Zombie Frame is no longer referenced by any active Process, Listener, or Child, it is automatically swept.

#### C. The "Drill Down" Mechanism

The Cactus Stack enables the **Reply (`^`)** operator to traverse "forward" into history.

1.  **Lookup:** When `^reply` is issued, the VM inspects the **Zombie List** associated with the current process.
2.  **Traversal:** It walks down the linked list of Zombie Frames that were "popped" to reach the current state.
3.  **Resume:** If a matching trap is found in a Zombie, the VM creates a new branch (Green Thread) resuming execution at that Zombie's IP, effectively "forking" the history.

### 6.5 Function Semantics & Currying

Rhumb functions adhere to a **"Loose Argument"** policy with explicit support for Partial Application (Currying).

#### Loose Invocation: `foo(...)`

When a function is called directly:

  * **Missing Args:** If the function expects 3 args but gets 1, the remaining args are filled with **`___` (Empty)**.
  * **Extra Args:** If the function expects 1 arg but gets 3, the extras are ignored (but available via `$`).
  * **Result:** The function executes immediately.

#### Partial Application (Currying): `<foo>(...)`

When a function is wrapped in Angle Brackets `<...>`:

  * **Mechanism:** The VM creates a **New Closure**.
  * **Binding:** The supplied arguments are bound to the parameters of the *original* function.
  * **Result:** Returns a **New Function** expecting the *remaining* arguments. It does **not** execute the body.


-----

## 7\. Concurrency & The Tuplespace

Rhumb replaces the traditional Call Stack with a **Hierarchical Tuplespace** based on the *Syndicated Actor Model (SAM)*. This system unifies Concurrency, Event Handling, and State Management into a single spatial metaphor.

### 7.1 Conceptual Model

  * **The Ether:** Every executing process (Green Thread) possesses a **Local Space**.
  * **Hierarchy:** Spaces are arranged in a tree. Every Space has a reference to its **Parent Space** (who spawned it).
  * **Zombie Frames:** When a process yields or pauses (e.g., waiting for a signal), its Stack Frame is not destroyed. It remains in memory as a "Zombie," allowing deeper frames (Children) to "Drill Down" and reply to it later.

### 7.2 Operator Taxonomy

The syntax distinguishes between ephemeral events and persistent state to prevent race conditions. **All operators evaluate to ___ (Empty).**

| Feature          | Symbol   | Type  | Direction        | Semantics                                                                                                               |
|:-----------------|:---------|:------|:-----------------|:------------------------------------------------------------------------------------------------------------------------|
| **Signal**       | **`#`**  | Event | **Up** (Bubble)  | **Ephemeral.** "I am emitting this event now." Travels from Child $\to$ Parent. If not caught immediately, it vanishes. |
| **Reply**        | **`^`**  | Event | **Down** (Drill) | **Ephemeral.** "I am replying to this specific context." Travels from Parent $\to$ Zombie Frame.                        |
| **Proclamation** | **`$`**  | State | **Static** (Pin) | **Persistent.** "I am AT this state." Sticks to the Local Space until explicitly removed.                               |
| **Subscription** | **`<>`** | Logic | **N/A**          | **Reactive.** "Watch this Space." Spawns a handler when a matching Tuple appears.                                       |

### 7.3 Realm Syntax & Lifecycle

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

### 7.4 The "Helium Balloon" Algorithm (Signal Propagation)

To prevent memory leaks in long-running servers, Signals (`#`) are active agents of transport.

1.  **Post:** Instruction `POST #sig` is executed.
2.  **Check:** VM checks the **Current Space's** active listeners (`<>`).
3.  **Deliver:** If a listener matches, the Signal is handed off, and propagation stops.
4.  **Ascend:** If no listener matches, the Signal immediately moves to `Space.Parent`.
5.  **Garbage Collection:** If the Signal reaches `World` (Root) and is still uncaught, it is **discarded**. This ensures that "fire-and-forget" events do not accumulate in memory.

### 7.5 The "Zombie Walk" Algorithm (Reply Injection)

Replies (`^`) allow a helper process to inject data back into a paused requestor.

1.  **Trigger:** Instruction `INJECT ^ack` is executed in a shallow frame (e.g., an error handler).
2.  **Scan:** The VM retrieves the **Stack Trace** of the current process.
3.  **Descend:** It iterates *forward* (deeper) into the stack, checking each paused "Zombie Frame."
4.  **Match:** It checks if the Zombie Frame has a `TRAP_REPLY` table entry matching `^ack`.
5.  **Resume:** If found, the VM transfers execution control **back** to that Zombie Frame's Instruction Pointer (IP), passing the data arguments.

### 7.6 Reactive Realms & Subscriptions

The syntax `realm <> [ pattern ] -> { body }` acts as a generic lifecycle manager.

  * **Initialization:** The `SUBSCRIBE` opcode registers a daemon on the target Space.
  * **Arrival (Spawn):** When a Tuple (Signal/Proclamation) matches `pattern`, a new Green Thread is spawned.
      * **Implicit Pinning:** Variables defined in `pattern` (e.g., `who`) act as **Filters**. The thread only wakes for tuples matching that specific value.
  * **Departure (Teardown):** If a Proclamation is retracted (removed), the VM injects a special `$empty` signal into the thread.
      * The user handles this via `empty .. log("Left")`.

### 7.7 Vassals (Facets & Attenuation)

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

### 7.8 Go Implementation Strategy

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

### 7.9 Distributed Rhumb (Network Transparency)

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

#### B. Process Migration (The Freezer)

When a function is passed to a Networked Realm (e.g., `node -> task()`), the VM
performs **Transparent Migration**.

1.  **Freeze:** The `task` closure is paused.
2.  **Sanitize:** The VM verifies that the closure does *not* capture non-transferable resources (e.g., File Handles, Local Mutexes).
3.  **Serialize:** The Closure's **Bytecode** (Logic Node IDs) and **Upvalues** (Captured Data) are packed into a binary payload.
4.  **Transmit:** The payload is sent to the remote node.
5.  **Hydrate:** The remote node unpacks the payload, links the IDs to its local Libraries (via the Resolver Protocol), and resumes execution.

#### C. The Dependency Check

Before accepting a migrated process, the Remote Node validates the **Resolver
Headers**.

  * **Check:** "Do I have `{!|math|1.2.0}`?"
  * **Result:**
      * **Yes:** Accept and run.
      * **No:** Reject (or optionally request the missing library blob).

-----

## 8\. Testing Strategy

A pyramid approach to ensure language stability.

1.  **Unit Tests:** Go tests inside `/internal`.
2.  **VM Spec Tests:** `.rnode` script files in `/test/vm_spec`.
3.  **UI Logic Tests:** Headless tests in `/ui/wm`.
4.  **Integration Scenarios:** End-to-end tests simulating a user interaction.

-----

## 9\. Language Grammar & Symbols

The final reserved symbol table for the parser.

### General Symbols

| Symbol     | Name       | Role          | Syntax          | Meaning                                     |
|:-----------|:-----------|:--------------|:----------------|:--------------------------------------------|
| **`!`**    | Base       | Receiver      | `!\field`       | "My field" (Mutable/Immutable based on def) |
| **`@`**    | Parent     | Inheritance   | `!@console\log` | "My Parent named Console"                   |
| **`\`**    | Access     | Member        | `user\name`     | "Field 'name' of user"                      |
| **`___`**  | Empty      | Literal       | `x .= ___`      | Empty/Nil value                             |
| **`_`**    | Ignore     | Literal       | `x .. _`        | Ignore/Placeholder Label                    |
| **`^=`**   | Caret-Eq   | Destruct      | `[x; y] ^= pt`  | "Unpack Point into x, y"                    |
| **`..`**   | Dot-Dot    | Select (Stop) | `yes .. log()`  | "If match, consume & execute"               |
| **`::`**   | Col-Col    | Select (Peek) | `yes :: log()`  | "If match, execute & continue"              |
| **`->`**   | Arrow      | Function      | `[] -> ()`      | Function definition                         |
| **`!>`**   | Bang-Arrow | Function      | `[] !> ()`      | Create method with bound \!                 |
| **`+>`**   | Plus-Arrow | Function      | `[] +> ()`      | Define and execute immediately (IIFE)       |
| **`!!`**   | Bang-Bang  | Bind          | `f !! obj`      | Rebind function to object                   |
| **`<fn>`** | Ref        | Reference     | `fn2 .= <fn1>`  | Get function by reference (No call)         |


### Concurrency Symbols

| Symbol     | Name    | Role         | Syntax         | Meaning                           |
|:-----------|:--------|:-------------|:---------------|:----------------------------------|
| **`#`**    | Hash    | Signal       | `obj#click`    | "Emit Event (Bubble Up)"          |
| **`^`**    | Caret   | Reply        | `worker^ack`   | "Inject Event (Drill Down)"       |
| **`$`**    | Dollar  | Proclamation | `obj$status`   | "Set State (Persistent)"          |
| **`<>`**   | Diamond | Subscription | `obj <> [...]` | "React to changes"                |
| **`<$>`**  | Realm   | Literal      | `r .= <$>`     | Create Child Realm                |
| **`<\|>`** | Realm   | Literal      | `r .= <\|>`    | Create Detached Realm             |
| **`<{}>`** | Vassal  | Literal      | `v .= <{...}>` | Create Vassal (Attenuation Facet) |

### Math & Logic Operators

| Symbol   | Name     | Role  | Syntax   | Meaning             |
|:---------|:---------|:------|:---------|:--------------------|
| **`++`** | Plus     | Math  | `x ++ y` | Add / Concat        |
| **`--`** | Minus    | Math  | `x -- y` | Subtract            |
| **`**`** | Star     | Math  | `x ** y` | Multiply            |
| **`^^`** | Caret    | Math  | `x ^^ y` | Exponent            |
| **`//`** | Slash    | Math  | `x // y` | Decimal Divide      |
| **`+/`** | Plus-Sl  | Math  | `x +/ y` | Integer Divide      |
| **`-/`** | Min-Sl   | Math  | `x -/ y` | Modulo              |
| **`*^`** | Star-Car | Math  | `x *^ y` | Scientific Notation |
| **`^/`** | Car-Sl   | Math  | `x ^/ y` | Root                |
| **`==`** | Eq       | Logic | `x == y` | Equality            |
| **`~~`** | Tilde    | Logic | `x ~~ y` | Inequality          |
| **`>>`** | GT       | Logic | `x >> y` | Greater Than        |
| **`<<`** | LT       | Logic | `x << y` | Less Than           |
| **`>=`** | GTE      | Logic | `x >= y` | Greater/Equal       |
| **`<=`** | LTE      | Logic | `x <= y` | Less/Equal          |
| **`/\`** | Wedge    | Logic | `x /\ y` | AND                 |
| **`\/`** | Vee      | Logic | `x \/ y` | OR                  |

### Map & Structure Operators

| Symbol   | Name    | Role      | Syntax   | Meaning            |
|:---------|:--------|:----------|:---------|:-------------------|
| **`\|`** | Pipe    | Structure | `1\|5`   | Range (1 to 5)     |
| **`=@`** | Eq-At   | Structure | `x =@ y` | Has Subfield       |
| **`~@`** | Til-At  | Structure | `x ~@ y` | Not Has Subfield   |
| **`=\`** | Eq-Sl   | Structure | `x =\ y` | Has Field          |
| **`~\`** | Til-Sl  | Structure | `x ~\ y` | Not Has Field      |
| **`@@`** | At-At   | Structure | `x @@ y` | Temporary Subfield |
| **`&&`** | Amp-Amp | Structure | `x && y` | Concatenate Lists  |
| **`\\`** | Sl-Sl   | Structure | `x \\ y` | Nested Access      |

### Control & Assignment

| Symbol     | Name   | Role   | Syntax     | Meaning          |
|:-----------|:-------|:-------|:-----------|:-----------------|
| **`.=`**   | Dot-Eq | Assign | `x .= y`   | Immutable Assign |
| **`:=`**   | Col-Eq | Assign | `x := y`   | Mutable Assign   |
| **`=>`**   | Arrow  | Flow   | `x => y`   | If True Then     |
| **`~>`**   | Squig  | Flow   | `x ~> y`   | If False Then    |
| **`\|>`**  | Pipe   | Flow   | `x \|> y`  | While Loop       |
| **`\|\|`** | Pipe   | Flow   | `x \|\| y` | Functional Pipe  |
| **`??`**   | Ques   | Flow   | `x ?? y`   | Default/Coalesce |