## 6\. Proposed Base Library

Here's the current draft of our upcoming base library.

In Rhumb, the base library is not a global namespace but a set of
**Resolvable Artifacts**. To use them, you must import them using the `!`
resolver.

```rhumb
math := {!|🧮|-} % each base library has an emoji identifier
area .= c ** math\π
```

The emoji are not directly available in your project files because Rhumb doesn't
bring anything extra into the global namespace without explicitly telling it to
do so but you could choose to label them with emoji if you like:

```rhumb
🧮 := {!|🧮|-}
area .= c ** 🧮\π
```

### 6\.1 Base Library Examples

Base libraries all use the signal `#***(code; msg; &data)` for bubbling up
non-panic errors.

#### 🐚 Shell (UI & TTY)
**Emoji:** Spiral Shell (`U+1F41A`)
**Domain:** Human-Computer Interaction, Terminal UX, ANSI codes.

*   **Top Level:**
    *   `🐚\ask("Prompt")`: Interactive input.
    *   `🐚\secret("Prompt")`: Masked input (passwords).
    *   `🐚\size()`: Returns terminal dimensions `[w; h]`.
*   **`🐚\🎨` Palette (Color):**
    *   `red()`, `blue()`, `bold()`, `rainbow()`: Text styling.
    *   `strip()`: Removes formatting for logs.
*   **`🐚\📍` Cursor:**
    *   `move(x; y)`, `up(n)`, `clear()`: Absolute positioning for TUI apps.
*   **`🐚\🍱` Widgets:**
    *   `spin(func)`: Async loading spinner.
    *   `bar(current; total)`: Progress bar.

#### 🖥️ System (OS & Kernel)
**Emoji:** Desktop Computer (`U+1F5A5`)
**Domain:** Operating System, Environment, Hardware.

*   **Top Level:**
    *   `🖥️\args`: CLI arguments.
    *   `🖥️\env` / `set-env()`: Environment variables.
    *   `🖥️\exit(code)`: Terminate process.
    *   `🖥️\pid`: Process ID.
*   **`🖥️\🚀` Launcher (Exec):**
    *   `run(cmd)`: Blocking execution.
    *   `spawn(cmd)`: Background execution.
    *   `find(tool)`: Look up binary in `$PATH`.
*   **`🖥️\🚦` Signals:**
    *   `listen(SIGINT)`: Handle Ctrl+C.
*   **`🖥️\🫆` Info:**
    *   `os`, `arch`, `cpus`: Hardware specs.

#### 🧮 Math
**Emoji:** Abacus (`U+1F9EE`)
**Domain:** Advanced calculation, Randomness, Constants.

*   **Top Level:**
    *   `🧮\pi`, `🧮\e`: Constants.
    *   `🧮\abs()`, `🧮\min()`, `🧮\max()`: Basic utilities.
*   **`🧮\📐` Trig:**
    *   `sin()`, `cos()`, `tan()`, `atan2()`: Geometry functions.
*   **`🧮\🎲` Random:**
    *   `int(min; max)`: Random integer.
    *   `float()`: Random 0.0-1.0.
    *   `shuffle(list)`: Randomizes an array order.
    *   `seed(val)`: Deterministic seeding.

#### 🔡 Text
**Emoji:** Input Latin Uppercase (`U+1F520`)
**Domain:** String manipulation, Regex, Parsing.

*   **Top Level:**
    *   `🔡\split(str; delim)`: Breaks strings into arrays.
    *   `🔡\join(arr; delim)`: Combines arrays into strings.
    *   `🔡\trim(str)`: Whitespace cleanup.
*   **`🔡\🔍` Regex:**
    *   `match(pattern; str)`: Boolean check.
    *   `find(pattern; str)`: Returns matches.
    *   `replace(pattern; str; new)`: Substitution.
*   **`🔡\✂️` Format:**
    *   `pad-left()`, `pad-right()`: Alignment helpers.

#### 📦 Data & Encoding
**Emoji:** Package (`U+1F4E6`)
**Domain:** Serialization, Formats, Hashing.

*   **`📦\📜` JSON:**
    *   `stringify(obj)`: Object to JSON string.
    *   `parse(str)`: JSON string to Object.
*   **`📦\📊` CSV:**
    *   `read(string)`: Parses CSV content into list of maps.
*   **`📦\🧱` Base64:**
    *   `encode(bytes)`, `Decode(str)`.
*   **`📦\#️⃣` Hash:**
    *   `MD5()`, `SHA256()`: Checksums.

#### 📂 Filesystem
**Emoji:** File Folder (`U+1F4C2`)
**Domain:** Disk I/O, Paths, Directories.

*   **Top Level:**
    *   `📂\read(path)`: Returns file content.
    *   `📂\write(path; content)`: Overwrites file.
    *   `📂\append(path; content)`: Adds to file.
    *   `📂\delete(path)`: Removes file.
*   **`📂\🛣️` Path:**
    *   `join(a; b)`: Cross-platform path combining.
    *   `ext(path)`: Get file extension.
    *   `base(path)`: Get filename.
*   **`📂\🗄️` Dir:**
    *   `list(path)`: Get files in directory.
    *   `make(path)`: Create directory (mkdir -p).

#### 🌐 Network
**Emoji:** Globe with Meridians (`U+1F310`)
**Domain:** HTTP, Sockets, URLs.

*   **`🌐\🔗` URL:**
    *   `parse(str)`: Breaks URL into host, path, query.
    *   `query(map)`: Builds a query string.
*   **`🌐\🛰️` HTTP Client:**
    *   `get(url)`: Simple fetch.
    *   `post(url; body)`: Submit data.
    *   `fetch(req)`: Advanced request with headers.
*   **`🌐\🕸️` Server:**
    *   `listen(port; handler)`: Starts a web server.
    *   `serve-directory(path)`: Static file server.

#### 🕰️ Time
**Emoji:** Mantelpiece Clock (`U+1F570`)
**Domain:** Dates, Durations, Scheduling.

*   **Top Level:**
    *   `🕰️\now`: Current timestamp.
    *   `🕰️\sleep(ms)`: Pauses execution.
*   **`🕰️\📅` Calendar:**
    *   `parse("YYYY-MM-DD")`: String to Date object.
    *   `format(date; "Format")`: Date to String.
*   **`🕰️\⏱️` Stopwatch:**
    *   `start()`, `stop()`: High-precision timing for benchmarks.

### 6\.2 Standard Error Codes

All Base Libraries return errors as the signal `#***(code; msg; data)`. Rhumb adopts a hybrid coding standard:

* **Universal Codes (HTTP-aligned):** Used for common logical errors across all libraries.
* **Domain Codes (1000+):** Used for specific error conditions within a library.

| Range       | Domain               | Description                                       |
|:------------|:---------------------|:--------------------------------------------------|
| **400-499** | **Usage Error**      | You called the function incorrectly (User Fault). |
| **500-599** | **Runtime Error**    | The function failed to execute (System Fault).    |
| **1000+**   | **Library Specific** | Deep errors specific to `fs`, `net`, etc.         |

**Common Universal Codes:**

| Code    | Name            | Meaning                        | Example                  |
|:--------|:----------------|:-------------------------------|:-------------------------|
| **400** | **Bad Arg**     | Invalid type, range, or arity. | `math\sin("foo")`        |
| **401** | **Permission**  | capabilities denied by Vassal. | `fs\read` blocked        |
| **404** | **Not Found**   | Resource/Key missing.          | `fs\read("missing.txt")` |
| **408** | **Timeout**     | Operation took too long.       | `net\fetch`              |
| **409** | **Conflict**    | State prevents action.         | `fs\mkdir` exists        |
| **500** | **Native Fail** | Internal Go error / Panic.     | Memory allocation        |
| **501** | **Not Impl**    | Feature missing on this OS.    | Windows-only syscall     |
| **999** | **Panic**       | Unrecoverable Runtime Panic.   | `core\***`               |

### 6\.3 Core Functions

These functions are implicitly imported and available everywhere.

| Function  | Signature  | Description                                                                                                                                                  |
|:----------|:-----------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **`***`** | `***(msg)` | **Panic.** Immediately raises the `#***` signal with code `999`. Used to mark unreachable code or fatal errors. <br>`valid ?? ***("Invalid state")` |