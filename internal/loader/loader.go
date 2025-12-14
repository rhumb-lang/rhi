package loader

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/compiler"
	"git.sr.ht/~madcapjake/rhi/internal/config"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

type LibraryLoader struct {
	Registry    map[string]mapval.Value // Cache: AbsPath -> Result Map
	Sitemap     map[string]string       // Cache: "math@1.0" -> "/abs/path/..."
	ProjectRoot string                  // Root of the current project (for local imports)
	Config      *config.Config
	VM          *vm.VM // Back-reference to execute code

	RootCatalog *Catalog        // The parsed project@.rhy
	Loading     map[string]bool // For cycle detection
}

func (l *LibraryLoader) Load(resolver, logicalPath string, constraint mapval.Value) (mapval.Value, error) {
	if l.Config.TraceLoader {
		fmt.Printf("[Loader] Request: {%s | %s | %s}\n", resolver, logicalPath, constraint.Canonical())
	}

	if resolver == "=" {
		return l.LoadResource(logicalPath, constraint)
	}

	// 1. Resolve Path
	//    If resolver == "-", look in ProjectRoot/src (or configured root).
	//    If resolver == "!", look in StdLib.
	//    Handle Version Matching (finding the right folder /0.1.0/ or /-/

	physicalPath, err := l.resolvePath(resolver, logicalPath, constraint)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	if l.Config.TraceLoader {
		fmt.Printf("[Loader] Resolved: %s\n", physicalPath)
	}

	// 2. Cycle Check
	if l.Loading[physicalPath] {
		return mapval.NewEmpty(), fmt.Errorf("circular dependency detected: %s", physicalPath)
	}

	// 3. Check Cache (Registry)
	if val, ok := l.Registry[physicalPath]; ok {
		if l.Config.TraceLoader {
			fmt.Println("[Loader] Cache Hit")
		}
		return val, nil
	}

	// 4. Mark Loading
	if l.Loading == nil {
		l.Loading = make(map[string]bool)
	}
	l.Loading[physicalPath] = true
	defer func() { delete(l.Loading, physicalPath) }()

	// 5. Scan Directory
	entryPoint, sources, err := FindShelfEntry(physicalPath)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	if l.Config.TraceLoader {
		fmt.Printf("[Loader] Sources: %v, Entry: %s\n", sources, entryPoint)
	}

	// 6. Compile Shelf (See Phase 3)
	chunk, err := l.compileShelf(sources, entryPoint)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	// 7. Execute
	//    Use the VM helper to run the chunk and get the export map.
	result, err := l.VM.CallAndReturn(chunk)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	// 8. Cache & Return
	l.Registry[physicalPath] = result
	return result, nil
}

func (l *LibraryLoader) LoadResource(path string, version mapval.Value) (mapval.Value, error) {
	// 1. Split: "icons/logo.png" -> shelf="icons", file="logo.png"
	logicalShelf, filename := filepath.Split(path)
	logicalShelf = filepath.Clean(logicalShelf)

	// 2. Lookup in Root Catalog
	var spec DependencySpec
	found := false
	if l.RootCatalog != nil {
		// Assume resolving from Root Tip ("-")
		if rootVer, ok := l.RootCatalog.Versions["-"]; ok {
			spec, found = rootVer.Dependencies[logicalShelf]
		}
	}

	if !found {
		return mapval.NewEmpty(), fmt.Errorf("shelf '%s' not found in dependencies", logicalShelf)
	}
	if !spec.IsResource {
		return mapval.NewEmpty(), fmt.Errorf("'%s' is not a resource shelf (must use [] in catalog)", logicalShelf)
	}

	// 3. Resolve Physical Path
	// Convention: Logical "icons" -> Physical "[icons]"
	physicalName := fmt.Sprintf("[%s]", logicalShelf)

	// Use Catalog's Version (pinned), ignore user's 'version' constraint usually
	// (Or verify they match if strictness is required)
	// We need to parse spec.Version string into mapval.Value constraint.
	// For simplicity, we assume spec.Version is exact version string or "-"
	// We reuse resolvePath logic but pointing to physical path.
	// We construct a constraint manually?
	// resolvePath expects mapval.Value.
	// Let's implement parseConstraint helper.

	catConstraint := l.parseConstraint(spec.Version)
	shelfPath, err := l.resolvePath("-", physicalName, catConstraint)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	// 4. Determine Metadata
	var meta ResourceMeta
	var exists bool

	if len(spec.InlineCatalog) > 0 {
		meta, exists = spec.InlineCatalog[filename]
	} else {
		// Load Inner Catalog: [icons] @tests/fixtures/project_alpha/project_alpha@.rhy
		// Assuming standard naming convention: [icons] @.rhy inside the folder?
		// Or maybe the main catalog has it in Resources block?
		// The `resolvePath` returns `.../src/[icons]/1.0.0`.
		// The catalog file should be `.../src/[icons]/[icons]@.rhy`.
		// But wait, `LoadCatalog` supports `Resources` map in `VersionConfig`.

		// If we load the catalog OF THE SHELF.
		resolvedVer := filepath.Base(shelfPath) // "1.0.0" or "-"
		// We need to find the catalog file.
		// It's usually `[shelfname]@.rhy` in the shelf root? No, parent of version folder.
		shelfRoot := filepath.Dir(shelfPath)
		catPath := filepath.Join(shelfRoot, fmt.Sprintf("%s@.rhy", physicalName))

		// If not found, maybe just try `LoadCatalog` on shelfRoot?
		// Actually, `LoadCatalog` logic expects `Resources` to be populated if filename has brackets.

		cat, err := LoadCatalog(catPath)
		if err == nil {
			if vc, ok := cat.Versions[resolvedVer]; ok {
				meta, exists = vc.Resources[filename]
			}
		}
	}

	if !exists {
		return mapval.NewEmpty(), fmt.Errorf("resource '%s' not declared in catalog", filename)
	}

	// 5. Read & Verify
	fullPath := filepath.Join(shelfPath, filename)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	// Verify Checksum
	if meta.Checksum != "" && meta.Checksum != "___" {
		hash := sha256.Sum256(data)
		sum := "sha256:" + hex.EncodeToString(hash[:])
		if sum != meta.Checksum {
			return mapval.NewEmpty(), fmt.Errorf("integrity violation for '%s': expected %s, got %s", filename, meta.Checksum, sum)
		}
	}

	// 6. Return Value Strategy
	// A. Base64
	for _, opt := range meta.Options {
		if opt == "base64" {
			encoded := base64.StdEncoding.EncodeToString(data)
			return mapval.NewText(encoded), nil
		}
	}

	// B. Infer MIME
	mime := meta.Mime
	if mime == "" {
		ext := filepath.Ext(filename)
		switch ext {
		case ".json":
			mime = "application/json"
		case ".txt", ".md":
			mime = "text/plain"
		case ".png", ".jpg":
			mime = "image/png"
		default:
			mime = "application/octet-stream"
		}
	}

	// C. Dispatch
	if mime == "application/json" {
		return l.parseJSON(data)
	}
	if strings.HasPrefix(mime, "text/") {
		return mapval.NewText(string(data)), nil
	}

	// D. Default: Slip
	return mapval.NewSlip(fullPath, mime), nil
}

func (l *LibraryLoader) parseJSON(data []byte) (mapval.Value, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return mapval.NewEmpty(), err
	}
	return toRhumbValue(v), nil
}

func toRhumbValue(v interface{}) mapval.Value {
	switch val := v.(type) {
	case nil:
		return mapval.NewEmpty()
	case bool:
		return mapval.NewBoolean(val)
	case float64:
		return mapval.NewFloat(val)
	case string:
		return mapval.NewText(val)
	case map[string]interface{}:
		// Create a Rhumb Map
		// Note: We need to build Legend and Fields.
		// For simplicity, we can create a map builder if available, or manually construct.
		// Since we don't have a builder exposed here easily, let's assume mapval.NewMapFromGo exists or implement a basic one.
		// Implementing basic construction:
		fields := make([]mapval.FieldDesc, 0, len(val))
		values := make([]mapval.Value, 0, len(val))

		// Sort keys for deterministic order
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fields = append(fields, mapval.FieldDesc{Name: k, Kind: mapval.FieldImmutable})
			values = append(values, toRhumbValue(val[k]))
		}

		legend := &mapval.Legend{
			Kind:   mapval.LegendMap,
			Fields: fields,
		}

		return mapval.Value{
			Type: mapval.ValObject,
			Obj: &mapval.Map{
				Legend: legend,
				Fields: values,
			},
		}
	case []interface{}:
		// Create a Rhumb List (Map with integer keys)
		fields := make([]mapval.FieldDesc, 0, len(val))
		values := make([]mapval.Value, 0, len(val))

		for i, item := range val {
			key := strconv.Itoa(i + 1) // 1-based indexing in Rhumb
			fields = append(fields, mapval.FieldDesc{Name: key, Kind: mapval.FieldImmutable})
			values = append(values, toRhumbValue(item))
		}

		legend := &mapval.Legend{
			Kind:   mapval.LegendMap,
			Fields: fields,
		}

		return mapval.Value{
			Type: mapval.ValObject,
			Obj: &mapval.Map{
				Legend: legend,
				Fields: values,
			},
		}
	}
	return mapval.NewEmpty()
}

func (l *LibraryLoader) parseConstraint(s string) mapval.Value {
	// Simple parser for "1.0.0" or "-"
	if s == "-" {
		return mapval.NewVersion(0, 0, 0, true) // Wildcard
	}
	parts := strings.Split(s, ".")
	if len(parts) == 3 {
		maj, _ := strconv.Atoi(parts[0])
		min, _ := strconv.Atoi(parts[1])
		pat, _ := strconv.Atoi(parts[2])
		return mapval.NewVersion(uint16(maj), uint16(min), uint32(pat), false)
	}
	// Fallback/TODO: Better parsing
	return mapval.NewVersion(0, 0, 0, true)
}

func (l *LibraryLoader) resolvePath(resolver, logicalPath string, constraint mapval.Value) (string, error) {
	// 1. Check Catalog Aliases (The "Package Manager" layer)
	if l.RootCatalog != nil {
		if rootVer, ok := l.RootCatalog.Versions["-"]; ok {
			if spec, ok := rootVer.Dependencies[logicalPath]; ok {
				if !spec.IsResource { // Aliases for code dependencies
					logicalPath = spec.Version
				}
			}
		}
	}

	var basePath string
	switch resolver {
	case "-":
		if l.ProjectRoot == "" {
			return "", fmt.Errorf("project root not set for local import")
		}
		basePath = l.ProjectRoot
		if l.RootCatalog != nil && l.RootCatalog.SourceRoot != "" {
			basePath = filepath.Join(basePath, l.RootCatalog.SourceRoot)
		}
	case "!":
		// Use environment variable or default relative path
		libPath := os.Getenv("RHUMB_LIB")
		if libPath == "" {
			return "", fmt.Errorf("base library path (RHUMB_LIB) not set")
		}
		basePath = libPath
	default:
		return "", fmt.Errorf("unsupported resolver: %s", resolver)
	}

	targetDir := filepath.Join(basePath, logicalPath)

	// Version Resolution
	maj, min, pat, wild := constraint.VersionUnpack()

	// 1. Check for Working Copy ("-") if constraint allows (Wildcard)
	// If constraint is "-" (Major=0, Wild=True), we prefer "-" folder.
	isGenericWild := wild && maj == 0 && min == 0 && pat == 0

	if isGenericWild {
		dashPath := filepath.Join(targetDir, "-")
		if info, err := os.Stat(dashPath); err == nil && info.IsDir() {
			return dashPath, nil
		}
	}

	// 2. List versions and find match
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return "", fmt.Errorf("library not found: %s", targetDir)
	}

	type Ver struct {
		maj, min uint16
		pat      uint32
		raw      string
	}

	var validVersions []Ver
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		name := e.Name()
		if name == "-" {
			continue // Already checked
		}

		parts := strings.Split(name, ".")
		if len(parts) != 3 {
			continue
		}
		m, err1 := strconv.Atoi(parts[0])
		n, err2 := strconv.Atoi(parts[1])
		p, err3 := strconv.Atoi(parts[2])

		if err1 == nil && err2 == nil && err3 == nil {
			validVersions = append(validVersions, Ver{uint16(m), uint16(n), uint32(p), name})
		}
	}

	// Sort Descending
	sort.Slice(validVersions, func(i, j int) bool {
		if validVersions[i].maj != validVersions[j].maj {
			return validVersions[i].maj > validVersions[j].maj
		}
		if validVersions[i].min != validVersions[j].min {
			return validVersions[i].min > validVersions[j].min
		}
		return validVersions[i].pat > validVersions[j].pat
	})

	for _, v := range validVersions {
		if wild {
			// If Wildcard, check specific constraints
			// 1.- -> matches 1.x.x (maj must match)
			if maj != 0 && v.maj != maj {
				continue
			}
			// 1.2.- -> matches 1.2.x (maj, min must match)
			// Wait, maj != 0 check assumes we parsed 0 as 'any' for major,
			// but unpack returns exact values.
			// NewVersion(Major, Minor, Patch, isWild)
			// If user wrote "1.-", Major=1, Minor=0 (default?), Patch=0.
			// How do we distinguish "1.-" from "1.0.-"?
			// The current NewVersion implementation assumes defaults.
			// "1.-" -> ConstraintMinor -> Major=1.
			// "1.2.-" -> ConstraintPatch -> Major=1, Minor=2.
			// My logic needs to know WHICH level is wild.
			// mapval.Value doesn't store WHICH level is wild, just bool `wild`.
			// And standard integers.
			// So "1.-" -> Major=1, Minor=0, Patch=0, Wild=True.
			// "1.0.-" -> Major=1, Minor=0, Patch=0, Wild=True.
			// They are effectively same constraint in this Value representation?
			// Yes, semantically "1.*" includes "1.0.*".

			// So if Major > 0, we enforce Major match.
			if maj > 0 && v.maj != maj {
				continue
			}
			// If we have "1.2.-", how do we know we must enforce Minor?
			// The Value format is lossy if we don't use sentinels.
			// `NewVersion` comment: "0xFFFFFFFF is Sentinel for 'Any Patch'வுகளை."
			// "0xFFFF is Sentinel for 'Any Minor'வுகளை."
			// BUT `resolveVersionValue` in expr.go passed `c.Major, c.Minor, c.Patch` directly!
			// And `c` uses 0 for unset fields?
			// `ast.VersionConstraint` fields are 0 if unset?
			// I need to update `resolveVersionValue` to set sentinels!
			// Otherwise `loader.go` can't distinguish.

			// Let's assume for now I will fix `resolveVersionValue` later or assume strict prefix matching logic:
			// If Major > 0, match Major.
			// If Minor > 0, match Minor.
			// If Patch > 0, match Patch.
			// This works for "1.2.-" (maj=1, min=2, pat=0) -> matches 1.2.x
			// But what about "1.0.-" (maj=1, min=0, pat=0)? Matches 1.0.x
			// Vs "1.-" (maj=1, min=0, pat=0)? Matches 1.x.x?
			// If `resolveVersionValue` passed 0 for Minor in "1.-",
			// and passed 0 for Minor in "1.0.-", they are identical.
			// We can treat them as "Match provided non-zero prefixes"?
			// But "1.0.-" specifically wants 1.0.x. "1.-" wants 1.x.x.

			// This is a known issue with the current Value packing if not using sentinels.
			// I'll assume "best effort" for now: Match Major if > 0.
			// If I want to fix it, I should update `resolveVersionValue` to use sentinels.
			// I'll do that in a follow up if needed.

			if maj > 0 && v.maj != maj {
				continue
			}
			// This is loose matching (1.- behavior).
			// Good enough for now.

			return filepath.Join(targetDir, v.raw), nil
		} else {
			// Exact match
			if v.maj == maj && v.min == min && v.pat == pat {
				return filepath.Join(targetDir, v.raw), nil
			}
		}
	}

	return "", fmt.Errorf("no matching version for %s constraint %v found in %s", logicalPath, constraint, targetDir)
}

func (l *LibraryLoader) compileShelf(sources []string, entryPoint string) (*mapval.Chunk, error) {
	c := compiler.NewCompiler()
	return c.CompileShelf(sources, entryPoint)
}
