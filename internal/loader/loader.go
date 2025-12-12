package loader

import (
	"git.sr.ht/~madcapjake/rhi/internal/config"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

type LibraryLoader struct {
	Registry map[string]mapval.Value // Cache: AbsPath -> Result Map
	Sitemap  map[string]string       // Cache: "math@1.0" -> "/abs/path/..."
	Config   *config.Config
	VM       *vm.VM // Back-reference to execute code
}

func (l *LibraryLoader) Load(resolver, logicalPath, versionConstraint string) (mapval.Value, error) {
	// 1. Resolve Path
	//    If resolver == "-", look in ProjectRoot/src (or configured root).
	//    If resolver == "!", look in StdLib.
	//    Handle Version Matching (finding the right folder /0.1.0/ or /-/)

	physicalPath, err := l.resolvePath(resolver, logicalPath, versionConstraint)
	if err != nil {
		return mapval.NewEmpty(), err
	}

	// 2. Check Cache (Registry)
	if val, ok := l.Registry[physicalPath]; ok {
		return val, nil
	}

	// 3. Scan Directory
	sources, entryPoint, err := FindShelfEntry(physicalPath)

	// 4. Compile Shelf (See Phase 3)
	chunk, err := l.compileShelf(sources, entryPoint)

	// 5. Execute
	//    Use the VM helper to run the chunk and get the export map.
	result, err := l.VM.CallAndReturn(chunk)

	// 6. Cache & Return
	l.Registry[physicalPath] = result
	return result, nil
}
