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
  assets: true
`
	if err := os.WriteFile(catalogPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write catalog file: %v", err)
	}

	// 2. Load the catalog
	manifest, err := LoadCatalog(catalogPath)
	if err != nil {
		t.Fatalf("LoadCatalog failed: %v", err)
	}

	// 3. Assert Metadata
	if manifest.Author != "Jake Russo" {
		t.Errorf("expected author 'Jake Russo', got '%s'", manifest.Author)
	}
	if manifest.License != "MIT" {
		t.Errorf("expected license 'MIT', got '%s'", manifest.License)
	}
	if manifest.SourceRoot != "src" {
		t.Errorf("expected source root 'src', got '%s'", manifest.SourceRoot)
	}
	if len(manifest.Keywords) != 2 {
		t.Errorf("expected 2 keywords, got %d", len(manifest.Keywords))
	}

	// 4. Assert Versions
	if len(manifest.Versions) != 2 {
		t.Errorf("expected 2 versions, got %d", len(manifest.Versions))
	}

	// Check "0.1.0"
	v010, ok := manifest.Versions["0.1.0"]
	if !ok {
		t.Fatal("version 0.1.0 not found")
	}
	if v010.Dependencies["core"] != "1.0.0" {
		t.Errorf("expected core dependency '1.0.0', got '%s'", v010.Dependencies["core"])
	}
	if !v010.IsResource {
		t.Error("expected IsResource to be true for 'assets: true' (parsed logic in LoadCatalog needs verification)")
	}

	// Check "-"
	vDash, ok := manifest.Versions["-"]
	if !ok {
		t.Fatal("version - not found")
	}
	if vDash.Alias != "0.1.0" {
		t.Errorf("expected alias '0.1.0', got '%s'", vDash.Alias)
	}
}
