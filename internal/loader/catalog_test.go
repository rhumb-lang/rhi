package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadCatalog(t *testing.T) {
	// 1. Create a temp directory and file
	tmpDir := t.TempDir()
	catalogPath := filepath.Join(tmpDir, "test@.rhy")

	content := `
👤: "Jake Russo"
🪪: "MIT"
📦: "https://github.com/example/repo"
🏷️:
  - "game"
  - "physics"
📝: "A test catalog"
📂: "src"

-:
  <-: 0.1.0

0.1.0:
  core: 1.0.0
  math: 2.0.0!
  assets: [1.0.0]
`
	if err := os.WriteFile(catalogPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write catalog file: %v", err)
	}

	// 2. Load the catalog
	catalog, err := LoadCatalog(catalogPath)
	if err != nil {
		t.Fatalf("LoadCatalog failed: %v", err)
	}

	// 3. Assert Metadata
	if catalog.Author != "Jake Russo" {
		t.Errorf("expected author 'Jake Russo', got '%s'", catalog.Author)
	}
	if catalog.License != "MIT" {
		t.Errorf("expected license 'MIT', got '%s'", catalog.License)
	}
	if catalog.SourceRoot != "src" {
		t.Errorf("expected source root 'src', got '%s'", catalog.SourceRoot)
	}
	if len(catalog.Keywords) != 2 {
		t.Errorf("expected 2 keywords, got %d", len(catalog.Keywords))
	}

	// 4. Assert Versions
	if len(catalog.Versions) != 2 {
		t.Errorf("expected 2 versions, got %d", len(catalog.Versions))
	}

	// Check "0.1.0"
	v010, ok := catalog.Versions["0.1.0"]
	if !ok {
		t.Fatal("version 0.1.0 not found")
	}
	if spec, ok := v010.Dependencies["core"]; !ok || spec.Version != "1.0.0" {
		t.Errorf("expected core dependency '1.0.0', got '%v'", spec)
	}
    // Check Resource Dependency
    if spec, ok := v010.Dependencies["assets"]; !ok || !spec.IsResource {
        t.Errorf("expected assets to be a resource dependency")
    }

	// Check "-" (Flattening)
	vDash, ok := catalog.Versions["-"]
	if !ok {
		t.Fatal("version - not found")
	}
	if vDash.Alias != "0.1.0" {
		t.Errorf("expected alias '0.1.0', got '%s'", vDash.Alias)
	}
    // Verify Inheritance
    if spec, ok := vDash.Dependencies["core"]; !ok || spec.Version != "1.0.0" {
        t.Errorf("expected inherited core dependency '1.0.0' in tip, got '%v'", spec)
    }
}
