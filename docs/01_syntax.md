## 1\. Language Syntax

### 1\.1 Comments & Meta-Syntax

Rhumb uses the percent sign (`%`) for comments and meta-annotations.

| Symbol      | Name          | Syntax                            | Semantics                                                                                                                                                                                                                                                                                                                                                 |
|:------------|:--------------|:----------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **`%`**     | Line Comment  | `% text`                          | Ignored by Parser. Continues to end of line.                                                                                                                                                                                                                                                                                                              |
| **`%( %)`** | Block Comment | `%( text %)`                      | Ignored by Parser. Can be nested.                                                                                                                                                                                                                                                                                                                         |
| **`%=`**    | Assertion     | `x %= str [% optional test name]` | **Meta-Operator.** Ignored by Runtime. In Test Mode, asserts that the **String Representation** (`.String()`) of the expression matches the **Raw Text** provided on the right. No parsing or evaluation is performed on the right-hand side. Optionally has a test name when the assertion is separated from the end of line by another % comment symbol |
| **`%?`**    | Inspection    | `expr %?`                         | **Meta-Operator.** Ignored by Runtime. In Test Mode, it acts as the **"Bless" / "What is this?"** operator. It triggers an inspection log (`INSPECT: <value>`) to stdout, allowing the user to copy the output and "bless" a test by pasting it into a `%=` assertion.                                                                                    |

#### Implementation Note for `RhumbLexer.g4`

To support this, the `LineComment` rule consumes everything after `%`, including the `=` in `%=`.

```antlr
// RhumbLexer.g4
LineComment : '%' (~[\r\n] ~[\r\n]*)? -> channel(HIDDEN) ;
```

Since `%=` starts with `%`, the Lexer treats it as a comment and hides it from
the Parser. An IDE, however, can scan the hidden channel tokens to find `%=` and
run the assertions.

### 1\.2 Literals & Values

Symbols that represent fixed data or references.

| Symbol            | Name       | Syntax             | Meaning                                             |
|:------------------|:-----------|:-------------------|:----------------------------------------------------|
| **`___`**         | Empty      | `x .= ___`         | Empty/Nil value                                     |
| **`***`**         | Panic      | `x .= ***`         | Panic/Error value                                   |
| **`*`**           | Ignore     | `x .. *`           | Void Label                                          |
| **`<fn>`**        | Reference  | `f .= <g>`         | Subroutine Reference (Capture without executing)    |
| **`1.0.0`**       | Version    | `v .= 1.0.0`       | Semantic Version Literal                            |
| **`2024/01/01`**  | DateTime   | `t .= 2025/12/31`  | DateTime Literal (Absolute Point).                  |
| **`+00:00:00`**   | Duration   | `d .= +00:00:01`   | Time Duration (1 Second).                           |
| **`+0000/00/00`** | Duration   | `d .= +0001/00/00` | Calendar Duration (1 Year).                         |
| **`01.0`**        | Decimal    | `d .= 01.0`        | Arbitrary Precision Decimal (Requires leading zero) |
| **`.-`**          | DotDash    | `1.-`              | Wildcard Suffix (for Versions and Decimals)         |
| **`<$>`**         | Realm      | `r .= <$>`         | Child Realm Literal                                 |
| **`<\|>`**        | Realm      | `r .= <\|>`        | Detached Realm Literal                              |
| **`<()>`**        | Subroutine | `s .= <()>`        | Subroutine Literal                                  |
| **`<{}>`**        | Vassal     | `v .= <{}>`        | Vassal (Proxy) Literal                              |

#### 1\.2\.1 The Void Label (`*`)

The asterisk `*` is a reserved identifier representing **The Void**. It is used to match structure without binding data.

* **Behavior:** It matches any value but stores nothing.
* **Reading:** Evaluating `*` always returns **Empty** (it never holds a value).
* **Assignment:** Explicit assignment to `*` (e.g., `* := 10`) is a **Syntax
  Error**. It cannot be used as a storage target.
* **Usage:** It is exclusively used in **Patterns** and **Destructuring** to
  ignore specific arguments/elements or as the default pattern for
  **Selectors**.

#### 1\.2\.2 The Reference Routine (`<(...)>`)
The fundamental unit of execution in Rhumb is the Subroutine, denoted by angle brackets.

```rhumb
my_sub := <(foo(?1); bar(?2))>
my_sub('a'; 12)
```

Most users will use the Function operator `->`, Method operator `!>`, or
Immediate operator `+>` to create subroutines, but `<(...)>` represents the raw
code block without parameter binding.

### 1\.3 General Operators

Symbols used for object access, assignment, and function definition.

| Symbol   | Name       | Role        | Syntax      | Meaning                                     |
|:---------|:-----------|:------------|:------------|:--------------------------------------------|
| **`!`**  | Base       | Receiver    | `!\field`   | "My field" (Mutable/Immutable based on def) |
| **`@`**  | Parent     | Inheritance | `f@log`     | Delegate/Parent access                      |
| **`?`**  | Question   | Argument    | `?1`        | Raw subroutine argument access              |
| **`\`**  | Access     | Member      | `u\name`    | Field access operator                       |
| **`^=`** | Destruct   | Assign      | `[.a] ^= b` | Destructuring assignment                    |
| **`->`** | Arrow      | Function    | `[] -> ()`  | Function definition                         |
| **`!>`** | Bang-Arrow | Method      | `[] !> ()`  | Method definition (binds `!` to receiver)   |
| **`+>`** | Plus-Arrow | IIFE        | `[] +> ()`  | Define and execute immediately              |
| **`!!`** | Bang-Bang  | Bind        | `f !! obj`  | Rebind function to new object               |

#### 1\.2\.2 The Argument Accessor (`?`)
The prefix operator `?` is used to access the raw arguments passed to the current subroutine.

* **`?1` ... `?N`**: Accesses the Nth argument (1-based index).
* **`?0`**: Accesses the full list of arguments as a Map/Tuple.
* **Behavior**: Accessing an index that was not passed evaluates to `___` (Empty).


### 1\.4 Concurrency Symbols

| Symbol     | Name        | Role             | Syntax         | Meaning                                                        |
|:-----------|:------------|:-----------------|:---------------|:---------------------------------------------------------------|
| **`#`**    | Hash        | Signal           | `obj#click`    | Emit Event (Bubble Up)                                         |
| **`#()`**  | Hash-Paren  | Anonymous Signal | `#(foo)`       | Equivalent to `#""(val)`. Acts as a **Return** statement       |
| **`^`**    | Caret       | Reply            | `worker^ack`   | Inject Event (Drill Down)                                      |
| **`^()`**  | Caret-Paren | Anonymous Reply  | `^(bar)`       | Equivalent to `^""(val)`. Acts as a **Resume/Yield** statement |
| **`$`**    | Dollar      | Proclamation     | `obj$status`   | Set State (Persistent)                                         |
| **`<>`**   | Diamond     | Subscription     | `obj <> [...]` | React to changes                                               |
| **`<$>`**  | Realm       | Literal          | `r .= <$>`     | Create Child Realm                                             |
| **`<\|>`** | Realm       | Literal          | `r .= <\|>`    | Create Detached Realm                                          |
| **`<{}>`** | Vassal      | Literal          | `v .= <{...}>` | Create Vassal (Attenuation Facet)                              |

_**Note:** Unlike Object literals, a Signal literal is an executable statement. It immediately interrupts the current flow. You cannot write `x := #signal`._

### 1\.5 Math & Logic Operators

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

### 1\.6 Map & Structure Operators

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

### 1\.7 Control & Assignment

| Symbol     | Name   | Role   | Syntax     | Meaning          |
|:-----------|:-------|:-------|:-----------|:-----------------|
| **`=>`**   | Arrow  | Flow   | `x => y`   | If True Then     |
| **`~>`**   | Squig  | Flow   | `x ~> y`   | If False Then    |
| **`\|>`**  | Pipe   | Flow   | `x \|> y`  | While Loop       |
| **`\|\|`** | Pipe   | Flow   | `x \|\| y` | Functional Pipe  |
| **`??`**   | Ques   | Flow   | `x ?? y`   | Default/Coalesce |

### 1\.8 Contextual Operators

| Operator | Context          | Meaning                   | Behavior                                                 |
|----------|------------------|---------------------------|----------------------------------------------------------|
| **`.=`** | Routine / Map    | **Immutable Declaration** | Creates or updates a label/field **and freezes it**.     |
| **`:=`** | Routine / Map    | **Mutable Declaration**   | Creates or updates a label/field. Allows future changes. |
| **`..`** | Map Literal      | **Immutable Field**       | `[ key .. val ]`                                         |
| **`::`** | Map Literal      | **Mutable Field**         | `[ key :: val ]`                                         |
| **`..`** | Selector Literal | **Case & Break**          | `{ pattern .. routine }`                                 |
| **`::`** | Selector Literal | **Case & Fallthrough**    | `{ pattern :: routine }`                                 |