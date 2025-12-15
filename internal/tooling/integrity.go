package tooling

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~madcapjake/rhi/internal/loader"
	"gopkg.in/yaml.v3"
)

// EnsureIntegrity scans the catalog, verifies hashes, and optionally fixes '___'.
// Returns error on mismatch or if '___' is found when autoFix is false.
func EnsureIntegrity(catalogPath, projectRoot string, autoFix bool) error {
	data, err := os.ReadFile(catalogPath)
	if os.IsNotExist(err) {
		return nil // No catalog, no integrity to enforce
	}
	if err != nil {
		return err
	}

	var node yaml.Node
	if err := yaml.Unmarshal(data, &node); err != nil {
		return fmt.Errorf("catalog parse error: %v", err)
	}

	modified := false
	catalogDir := filepath.Dir(catalogPath)

	// Recursive walker for the YAML AST
	var walk func(n *yaml.Node) error
	walk = func(n *yaml.Node) error {
		if n.Kind == yaml.MappingNode {
			for i := 0; i < len(n.Content); i += 2 {
				key := n.Content[i]
				val := n.Content[i+1]

				// We look for values that look like resource metadata: contains "___" or "sha256:"
				if val.Kind == yaml.ScalarNode && (strings.Contains(val.Value, "sha256:") || strings.Contains(val.Value, "___")) {
					// Assume Key is the filename (Resource)
					// TODO: Support Dependency resolution if needed, but Resources are relative to catalog.
					targetPath := filepath.Join(catalogDir, key.Value)

					// Calculate Real Hash
					actualHash, err := loader.HashShelf(targetPath)
					if err != nil {
						// If file doesn't exist, maybe it's not a simple resource or path is wrong.
						// Warn or Error?
						// For now, return error to be safe.
						return fmt.Errorf("hashing failed for '%s' (referenced in catalog): %v", key.Value, err)
					}

					// Parse existing hash from value string
					parts := strings.Fields(val.Value)
					var existingHash string
					var hashIdx int = -1

					for idx, p := range parts {
						if p == "___" || strings.HasPrefix(p, "sha256:") {
							existingHash = p
							hashIdx = idx
							break
						}
					}

					if existingHash == "___" {
						if autoFix {
							// Replace ___ with actualHash in the value string
							// We need to preserve other parts of the string
							parts[hashIdx] = actualHash
							val.Value = strings.Join(parts, " ")
							modified = true
						} else {
							return fmt.Errorf("anchor is not a match for dependency '%s' (line %d, col %d): expected '___' actual '%s'", key.Value, val.Line, val.Column, actualHash)
						}
					} else if existingHash != "" {
						if existingHash != actualHash {
							return fmt.Errorf("anchor is not a match for dependency '%s' (line %d, col %d): expected '%s' actual '%s'", key.Value, val.Line, val.Column, existingHash, actualHash)
						}
					}
				}
			}
		}
		
		// Recurse
		for _, child := range n.Content {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}

	if err := walk(&node); err != nil {
		return err
	}

	if modified && autoFix {
		f, err := os.Create(catalogPath)
		if err != nil {
			return err
		}
		defer f.Close()
		enc := yaml.NewEncoder(f)
		enc.SetIndent(2)
		return enc.Encode(&node)
	}

	return nil
}
