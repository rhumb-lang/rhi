## 8\. rhIDE (The Rhumb IDE)

> 🚧 Not Yet Implemented 🚧

The IDE is a **Spatial Integrated Development Environment (SIDE)** that behaves like a Window Manager / Operating System rather than a text editor.

### 8\.1 Metaphor: The Route & The Desktop

* **The Route:** The persistent state of the IDE session (Window positions, icon position, open tools, active debuggers). It acts as the "Image" of the project.
* **The Desktop:** An infinite 2D plane where artifacts live.
* **The Deskbar:** A persistent HUD (Head-Up Display) in the corner containing:
    * **Mini-Map:** Overview of the desktop with Click-to-Teleport.
    * **Open List:** Tracks all active windows and icons.
    * **System Menu:** Settings and Global Commands.

### 8\.2 Window Management (BeOS Style)
The Window Manager (`/ui/wm`) implements a **Hybrid Stacking/Tiling** model.

* **Primitives:**
    * **Icon:** A minimized artifact resting on the "floor" (Z-index 0).
    * **Window:** An active artifact floating above icons (Z-index 1+).
    * **Tab:** The handle of a window.
* **Interactions:**
    * **Stacking:** Dragging Window A's tab onto Window B's tab merges them into a **Tab Group**. Tabs can slide horizontally to rearrange titles and clicked to make one tab active over the others.
    * **Tiling:** Dragging Window A's border to Window B's border "glues" them together. Resizing the edge resizes both.
    * **Edge Snapping:** Dragging a window to the screen edge tiles it to that quadrant.
* **Lifecycle:**
    * **Minimize:** Double-click tab $\rightarrow$ Converts to Icon at original location.
    * **Close:** Remove from Desktop (available via Context Menu).
    * **Open:** Context Menu $\rightarrow$ Spawns Icon at nearest empty space. Double-clicking icons will expand into the appropriate window at nearest available space

### 8\.3 Artifact Visualization
Different artifacts render differently when expanded into Windows.

| Artifact    | Icon View | Window View         | Behavior                                                                                   |
|:------------|:----------|:--------------------|:-------------------------------------------------------------------------------------------|
| **Shelf**   | Folder    | **Sub-Desktop**     | Displays contained Books/Shelves as Icons in a grid. Acts as a spatial file manager.       |
| **Library** | Package   | **Sub-Desktop**     | Read-only view of external dependency contents.                                            |
| **Book**    | File      | **Code Projector**  | The Projection Editor (Source Code).                                                       |
| **Tool**    | (none)    | **Form Projector** | Certain actions in the Projection Editor can trigger sidebar tools (e.g., Docs, Debugging, Stack Visualizations, Tuplespace visualizations, frame info, etc.) |
| **Catalog** | Tag       | **Form Projector** | Form-based editor for `.rhy` YAML files.                                                   |
| **Shell**   | (none)    | **Terminal Projector** | Openable from context menu. Like a line-based code projection editor but with some batteries included. |
| **Voyage**  | Gear      | **Viewport**        | Renders the output of a running process (Game/App).                                        |

### 8\.4 Quality Of Life
* **Context Menu:** The primary mechanism for instantiating new objects
  (Libraries, Shelves, Books).
* **Hyper-Navigation:** Right-clicking a label in the Code Projector offers
  "Jump to Definition" which spawns/focuses the target Book's window, drawing a
  visual "cable" between the usage and the definition if both are visible.
* **Documentation:** Right-clicking a label in the Code Projector offers
  "Documentation" which spawns a small sidebar window that renders the
  documentation for a given subroutine or function.
* **Translations:** Right-clicking a label in the Code Projector offers
  "Translations" which spanws a small sidebar window that shows all translations
  currently saved for that label. Right clicking on the tab will show a context
  menu that allows switching the presented target language for the whole file.
  Any untranslated content is shown in the default language for that library.
* **Refactoring:** Drag-and-drop operations on the Desktop map to AST transformations:
    * Dragging code text from Book A to Book B moves the function.
    * Dragging Book A into Shelf B moves the file on disk and updates references.
* **Testing:** Right clicking on a book or shelf offers a context menu option
  "Run Tests" which will execute the file using the `-test` flag and print the
  test results in a sidebar window (errors can be clicked to navigate to the
  offending code),
* **Hot Reload:** If you have executed a book and it generated a viewport, you
  can right click on the viewport tab to be offered an option "Reload On Code
  Change" which will automatically re-execute the same command that generated
  that viewport whenever a piece of code that was used by that viewport has been
  modified and the changes committed to the route, replacing the viewport with
  the newly executed version.

### 8\.5 Persistence Strategy

Rhumb employs a **Dual-Persistence Model** to decouple the requirements of the machine (speed/stability) from the requirements of the human (readability/git).

#### A. Hot State (`.ri`)
* **Format:** **Fast Binary Snapshot**.
* **Content:** A serialized graph of the `VM` heap (objects, stack frames, tuplespace) and the `WM` state.
* **Role:** Crash recovery and Fast Resume.
* **Mechanism:**
    * **Serialization:** Uses **The Freezer (Snapshot Mode)**. This reuses the distributed computing logic to efficiently flatten pointers into IDs, but allows for the persistence of local resource metadata.
    * **On Launch:** The IDE loads this file to resume execution quickly (< 100ms).
    * **On Run:** The IDE snapshots to this file frequently (e.g., on focus loss).

#### B. Cold Storage (The Artifacts)
* **Format:** Text-based files (`.__.rh`, `.rhy`).
* **Role:** Source of Truth for Git.
* **Mechanism:** The **Babel Daemon** watches the `.ri` and decompiles it to text in the background.

#### C. The Babel Daemon (`rhi babel`)
To ensure the UI remains fluid, the IDE **never writes text files directly**.
1.  **The Loop:** The IDE writes the `.ri` (Hot State) to disk.
2.  **The Signal:** The IDE notifies the background **Babel Daemon**.
3.  **The Sync:** The Daemon reads the `.ri` (Read-Only) and "decompiles" the binary state into the canonical Cold Storage files (`.__.rh`, `.rhy`) in the background.
    * **Safety:** If the text generation crashes or lags, the user's running session is unaffected.
    * **Performance:** The "Save" operation is instant from the user's perspective.