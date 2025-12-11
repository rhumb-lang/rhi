# Rhumb Architecture

**Project:** Rhumb Programming Language

**Architecture:** Image-Based, Prototype-Oriented, Reactive Tuplespace

**Host Implementation:** Go (Golang)

**UI Strategy:** Wayland-Native (Gio/OpenGL), Stacked Window Manager

-----

## Core Philosophy

Rhumb is designed to induce a new "mindset" in the developer.

  * **Alien Syntax:** Heavily symbol-based to break muscle memory.
  * **Code is Data:** The canonical representation is an Abstract Syntax Tree (AST) stored in binary/serialized formats.
  * **Localized View:** The IDE projects code into the user's native language (e.g., `yes` vs `oui`), but underlying logic is language-agnostic.
  * **Three-Path Lookup:** Variables resolve via Lexical (Static), Inheritance (Map), or Dynamic (Space) paths.

-----

## Directory Structure & Code Organization

The codebase follows a "Small File" philosophy with strict separation between Runtime Logic and UI Presentation.

```text
/rhumb
├── go.mod                # Dependencies
├── Makefile              # Build scripts (ANTLR, Tests)
├── /cmd
│   └── /rhi              # Rhumb Interpreter - CLI Runner / REPL
│   └── /rhide            # Rhumb IDE - Graphical Environment (Wayland)
├── /internal             # Core Logic (Private)
│   ├── /ast              # Abstract Syntax Tree definitions
│   ├── /cli              # CLI Entrypoint Helpers
│   ├── /color            # Terminal Output Styling
│   ├── /compiler         # Source -> Bytecode transformation
│   ├── /config           # Runtime Configuration (Flags/Env)
│   ├── /grammar          # ANTLR .g4 source & generated Go files
│   ├── /legacy           # Reference implementations (Pre-Architecture)
│   ├── /map              # Data Model & Opcodes (formerly 'object')
│   ├── /parser_util      # Syntax Error Formatting
│   ├── /testing          # Test Harness Utilities
│   ├── /visitor          # AST Builder (Parse Tree -> AST)
│   └── /vm               # Bytecode Interpreter & Cactus Stack
└── /tests                # Integration Scripts (*.rh)
```

### Core Logic Sub-Folders (`/internal`)

* **`/internal/grammar`**: Holds the ANTLR4 definitions.
    * `Rhumb.g4`: The master grammar file.
    * `generate.go` & `generate.sh`: Generation scripts.
    * *(Generated files: `parser.go`, `lexer.go`, `tokens`, etc.)*

* **`/internal/ast`**: Abstract Syntax Tree definitions.
    * `ast.go`: Base Node interfaces and Document root.
    * `literals.go`: Value nodes (Number, String, Key, Empty).
    * `fields.go`: Map field definitions.
    * `structures.go`: Complex structures (Maps, Selectors, Realms).
    * `operators.go`: Operator constants.

* **`/internal/compiler`**: **The Bytecode Compiler.** Converts AST to Chunks.
    * `compiler.go`: Main compiler driver.
    * `hoister.go`: Lexical scope analysis and variable hoisting.
    * `scope.go`: Symbol table management.
    * `expr.go`, `binary.go`, `chain.go`: Expression compilation.
    * `call.go`, `routine.go`: Function and call compilation.
    * `selector.go`, `map.go`: Structure compilation.

* **`/internal/visitor`**: **The AST Builder.** Converts Parse Tree (ANTLR) to AST.
    * `builder.go`: Main visitor entry point.
    * `chain.go`: "Fold Left" logic for chains (`a\b`, `realm#sig`).
    * `fields.go`: Map field parsing.
    * `ops.go`, `more_ops.go`: Operator parsing.

* **`/internal/map`**: **The Data Model.** (Formerly `object`).
    * `types.go`: Defines `Value`, `Object`, `Map`, `Legend`, `Function`, `Chunk`, `Upvalue`.
    * `opcodes.go`: Definition of all Bytecode Instructions.
    * `logic.go`: Boolean logic helpers.

* **`/internal/vm`**: **The Runtime Engine.**
    * `vm.go`: The main Fetch/Decode/Execute loop.
    * `frame.go`: Cactus Stack Frame implementation.
    * `primitives.go`: Native math and logic intrinsics.
    * `map_ops.go`: Object/Map manipulation instructions.
    * `space_ops.go`: Concurrency instructions (`POST`, `SUBSCRIBE`).
    * `structure_ops.go`: Flow control and pattern matching.

* **`/internal/config`**: Runtime configuration (Flags, Trace options).
    * `config.go`: `Config` struct.

* **`/internal/cli`**: Command-line interface logic.
    * `cli.go`: Entry point helpers.

* **`/internal/parser_util`**: Error handling utilities.
    * `errors.go`: Custom syntax error formatting.

## Documentation Index

### 📐 [Syntax & Grammar](01_syntax.md)
Detailed specification of the language grammar, symbols, literals (DateTime,
Duration, Version), and meta-syntax.

### 🧬 [The Map (Object) Model](02_object_model.md)
Explains the "Self"-style Prototype system, the Universal Legend (schema
optimization), and the memory layout of Values.

### ⚙️ [The Runtime (VM)](03_runtime.md)
Deep dive into the Cactus Stack, the Bytecode Compiler, and the Dynamic Type
Interaction matrix.

### 📡 [Concurrency & Space](04_concurrency.md)
Describes the Tuplespace model, Realm hierarchies, and the Signal/Reply
propagation algorithms.

### 📦 [Project System](05_project_system.md)
Covers the Library Resolution protocol, Catalogs (`.rhy`), and the "Tip" folder
structure for versioning.

### 🧪 [Testing](07_testing.md)
Details the multi-layered testing strategy, including Fuzzing, Concurrency
Stress Testing (Chaos Monkey), and the "Blessed Output" assertion syntax (%=).

### 🖥️ [rhIDE](08_rhide.md)
Specifications for the graphical environment, including the "Code Projector",
the spatial "Route/Desktop" metaphor, window management, and the
dual-persistence strategy.