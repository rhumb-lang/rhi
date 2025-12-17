# Changelog

All notable changes to this project will be documented in this file.

## [v0.4.0] - 2025-12-16 "Zombie Zone"

The asynchronous core is now fully operational. The focus is shifting to hardening the base library and optimizing the VM.

### 🚀 Features

- **Zombie Frames**: Implemented re-entrant continuation. Functions can now suspend execution (`#signal`), be moved to a "parking lot" (Zombie List), and be revived (`^reply`) later with new data without losing their state.
- **Async Primitives**:
    - **Signals (`#`)**: Fully implemented bubbling up the call stack to find monitors.
    - **Replies (`^`)**: Implemented drilling down into the Zombie list to resume suspended frames.
    - **Anonymous Replies**: Support for `^(val)` syntax (Yield behavior).
- **Effect Monitoring**: Enhanced compilation of `expression { selector }`. Complex expressions (like Block or Map literals) are now wrapped in Thunks to ensuring signals emitted *during* construction are correctly caught by the attached selector.
- **Tooling & CLI**:
    - **Frame Tracing**: Added `-trace-frame` flag for detailed execution debugging.
    - **Silent Mode**: Added `-silent` flag to suppress runtime warnings.
    - **Integrity**: Added `-sha256` flag for catalog anchor management.
- **Project Structure**:
    - **Module Rename**: Migrated Go module to `github.com/rhumb-lang/rhi`.
    - **Documentation**: Added comprehensive guides on Concurrency, Project Structure, and the Unified Activity Model.
- **Resolvers**:
    - **Resource (`=`)**: Support for importing resources from Resource Shelves.
    - **Native (`_`)**: Support for Go function interface.

### 🛠 Improvements

- **Module Rename**: Migrated Go module to `github.com/rhumb-lang/rhi`.
- **Documentation**: Added comprehensive guides on Concurrency, Project Structure, and the Unified Activity Model.

---

## [v0.3.0] - 2025-11-20 "Library League"

The core language semantics are stable. The focus is now shifting from internal VM mechanics to the Base library and tooling.

### 🚀 Features

- **Dates/Durations**: A literal type for working with dates and times (`2025/01/01`, `12:00:00`)
- **Versions**: Builtin semver as a literal type (`1.0.0`, `1.2.-`)
- **Decimals**: Builtin support for full precision decimal math (using `apd`)
- **Resolvers**: Implementation of the import system (`{ resolver | path | version }`).
    - **Standard Resolver (`!`)**: Imports from `RHUMB_LIB` (e.g. `{! | math | -}`)
    - **Local Resolver (`-`)**: Imports from project source (e.g. `{- | libs/calc | -}`)
- **Catalogs**: Support for `project@.rhy` metadata and dependency aliasing
- **Cycle Detection**: Runtime detection of circular dependencies
- **Loader Tracing**: Added `-trace-loader` flag
