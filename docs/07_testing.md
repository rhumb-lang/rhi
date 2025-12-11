## 7\. Testing

To ensure Rhumb is "ironclad" and production-ready, we employ a multi-layered
testing strategy ranging from unit logic to chaotic concurrency stress testing.

### 7\.1 Current Implementation (The Foundation)
The current codebase utilizes Go's native testing framework to verify individual
components.

* **Unit Tests (`*_test.go`):** Located alongside the code in `/internal`.
    * **Compiler Tests:** Verify that AST nodes translate to correct OpCodes
      (`internal/compiler/math_comprehensive_test.go`, `selector_test.go`).
    * **VM Tests:** Verify that OpCodes mutate the stack and heap correctly
      (`internal/vm/vm_test.go`).
    * **Cactus Stack Tests:** Specifically verify Zombie Frame persistence and
      parent-pointer traversal (`internal/vm/cactus_test.go`).
* **Integration Scripts (`.rh`):**
    * Located in `/tests` (e.g., `basics.rh`).
    * Executed via `main_test.go` which spins up a full VM instance, parses the
      script, and asserts the final stack state or standard output.

### 7\.2 Production Roadmap (The "Ironclad" Standard)
To move from "Working" to "Production Ready," the following testing layers must
be implemented.

#### A. Fuzz Testing (Security & Stability)
Since "Code is Data," the VM must be resilient to malformed bytecode and ASTs.
* **Parser Fuzzing:** Feed random byte streams to the ANTLR lexer/parser to
  ensure no panics occur on invalid syntax.
* **VM Fuzzing:** Feed random (valid) chunks of bytecode to the VM loop to
  ensure stack underflows/overflows are caught gracefully and do not crash the
  host process.

#### B. Concurrency Stress Testing (The "Chaos Monkey")
The Tuplespace model relies on complex locking (`sync.Map`, `sync.Mutex`). We
need to prove it is deadlock-free.
* **Race Detection:** Run all tests with `go test -race` in CI/CD.
* **"Thundering Herd" Simulators:** Spawn 10,000 Green Threads that all try to
  `SUBSCRIBE` and `POST` to the same Realm simultaneously to verify the "Helium
  Balloon" propagation logic doesn't drop signals or deadlock.

#### C. The "Rhumb Spec" Suite (Compliance)
A language-agnostic test suite defined in `.__.rh` files, similar to `test262`
for JavaScript.
* **Structure:** A folder of thousands of tiny files, each testing one edge case
  (e.g., `math-div-zero.__.rh`, `realm-retract-empty.__.rh`).
* **Metadata:** Each file contains a header defining the expected result or
  expected error code.
* **Role:** This allows future alternative VMs (e.g., a Rust or C++
  implementation) to verify they match the Rhumb specification.

#### D. Benchmark Regression (Performance)
* **Micro-benchmarks:** Measure the nanosecond cost of `OP_SEND` vs `OP_CALL` vs
  `OP_POST`.
* **Macro-benchmarks:** "Game Loop" scenarios (e.g., updating 1,000 entities in
  a Realm per frame) to track regression in the Garbage Collector or Dispatch
  logic.

#### E. UI/Projection Tests
* **Headless Rendering:** Verify that the Editor correctly projects AST $\to$
  Text without opening a window.
* **Input Fuzzing:** Simulate random keystrokes in the Editor logic to ensure
  the AST remains valid (or recovers gracefully) during partially-typed code.