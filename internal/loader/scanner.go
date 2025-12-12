package loader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindShelfEntry returns the path to the +file.rh if it exists.
// Returns "", nil if no entry point (pure library).
// Returns error if multiple +files exist.
func FindShelfEntry(dir string) (string, []string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", nil, err
	}

	var entryPoint string
	var sources []string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if filepath.Ext(name) != ".rh" {
			continue
		}

		fullPath := filepath.Join(dir, name)
		if strings.HasPrefix(name, "+") {
			if entryPoint != "" {
				return "", nil, fmt.Errorf("multiple entry points found in %s", dir)
			}
			entryPoint = fullPath
		} else {
			sources = append(sources, fullPath)
		}
	}
	return entryPoint, sources, nil
}
