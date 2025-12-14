package loader_test

import (
	"os"
	"path/filepath"
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/config"
	"git.sr.ht/~madcapjake/rhi/internal/loader"
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func TestLoadResource(t *testing.T) {
	// Setup paths
	wd, _ := os.Getwd()
	// We are in internal/loader, so we need to go up to root
	rootDir := filepath.Join(wd, "../../")
	fixtureRoot := filepath.Join(rootDir, "tests/fixtures/project_alpha")

	// Load Catalog
	catalogPath := filepath.Join(fixtureRoot, "project_alpha@.rhy")
	cat, err := loader.LoadCatalog(catalogPath)
	if err != nil {
		t.Fatalf("Failed to load catalog: %v", err)
	}

	// Setup Loader
	l := &loader.LibraryLoader{
		Registry:    make(map[string]mapval.Value),
		ProjectRoot: fixtureRoot,
		Config:      config.DefaultConfig(),
		RootCatalog: cat,
		VM:          vm.NewVM(), // Dummy VM
	}

	// Test
	// art: ["1.0.0"] in catalog -> src/[art]/1.0.0
	// We requested "art/hero.png" from "-" resolver?
	// In loader.go LoadResource:
	// logicalShelf="art", filename="hero.png"
	// spec found in RootCatalog (0.1.0 version if we pretend we are in 0.1.0 or tip?)
	// Wait, LoadResource uses l.RootCatalog.Versions["-"].
	// My updated catalog has "-": { "<-": "0.1.0" }.
	// My Flatten logic in LoadCatalog should merge 0.1.0 deps into "-".

	// Constraint is ignored for resources mostly, but we pass something.
	res, err := l.LoadResource("art/hero.png", mapval.NewEmpty())
	if err != nil {
		t.Fatalf("LoadResource failed: %v", err)
	}

	if res.Type != mapval.ValObject {
		t.Errorf("Expected ValObject, got %v", res.Type)
	}
    if res.Obj.Type() != mapval.ObjTypeSlip {
        t.Errorf("Expected ObjTypeSlip, got %v", res.Obj.Type())
    }

	slip := res.Obj.(*mapval.Slip)
	if slip.MimeType != "image/png" {
		t.Errorf("Expected image/png, got %s", slip.MimeType)
	}

	expectedPath := filepath.Join(fixtureRoot, "src/[art]/1.0.0/hero.png")
	if slip.Path != expectedPath {
		t.Errorf("Expected path %s, got %s", expectedPath, slip.Path)
	}

	// Test Inline Merging (icons)
	// icons:
	//   - logo.png: image/png
	//   - config.json: text/plain
	res1, err := l.LoadResource("icons/logo.png", mapval.NewEmpty())
	if err != nil {
		t.Fatalf("Failed to load icons/logo.png: %v", err)
	}
	if res1.Type != mapval.ValObject {
		t.Errorf("Failed to load first item in inline list, expected Object (Slip), got %v", res1.Type)
	} else if res1.Obj.Type() != mapval.ObjTypeSlip {
		t.Errorf("Expected ObjTypeSlip for logo.png, got %v", res1.Obj.Type())
	} else {
		s := res1.Obj.(*mapval.Slip)
		if s.MimeType != "image/png" {
			t.Errorf("Expected image/png for logo.png, got %s", s.MimeType)
		}
	}

	res2, err := l.LoadResource("icons/config.json", mapval.NewEmpty())
	if err != nil {
		t.Fatalf("Failed to load icons/config.json: %v", err)
	}
	// config.json: text/plain -> Should be loaded as Text
	if res2.Type != mapval.ValText {
		t.Errorf("Failed to load second item in inline list, expected ValText, got %v", res2.Type)
	}
}
