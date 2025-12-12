package loader

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestFindShelfEntry(t *testing.T) {
	tmpDir := t.TempDir()

	// Create files
	files := []string{
		"helper.rh",
		"logic.rh",
		"+entry.rh",
		"README.md", // Should be ignored
		".hidden.rh", // Should be ignored by convention? logic currently doesn't ignore dotfiles explicitly but FindShelfEntry checks Ext.
	}

	for _, name := range files {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte(""), 0644); err != nil {
			t.Fatalf("failed to create file %s: %v", name, err)
		}
	}

	entry, sources, err := FindShelfEntry(tmpDir)
	if err != nil {
		t.Fatalf("FindShelfEntry failed: %v", err)
	}

	// Verify Entry Point
	expectedEntry := filepath.Join(tmpDir, "+entry.rh")
	if entry != expectedEntry {
		t.Errorf("expected entry point %s, got %s", expectedEntry, entry)
	}

	// Verify Sources
	// Should contain helper.rh and logic.rh. README.md ignored.
	// logic.rh, helper.rh
	// .hidden.rh is NOT ignored by current implementation logic (only checks .rh ext and IsDir).
	// Let's assume .hidden.rh is a source file for now unless logic changes.
	
	expectedSources := []string{
		filepath.Join(tmpDir, ".hidden.rh"),
		filepath.Join(tmpDir, "helper.rh"),
		filepath.Join(tmpDir, "logic.rh"),
	}
	
	sort.Strings(sources)
	sort.Strings(expectedSources)

	if len(sources) != len(expectedSources) {
		t.Errorf("expected %d sources, got %d", len(expectedSources), len(sources))
	}

	for i, s := range sources {
		if s != expectedSources[i] {
			t.Errorf("source mismatch at %d: expected %s, got %s", i, expectedSources[i], s)
		}
	}
}

func TestFindShelfEntry_NoEntry(t *testing.T) {
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "lib.rh"), []byte(""), 0644)

	entry, sources, err := FindShelfEntry(tmpDir)
	if err != nil {
		t.Fatalf("FindShelfEntry failed: %v", err)
	}

	if entry != "" {
		t.Errorf("expected empty entry point, got %s", entry)
	}

	if len(sources) != 1 {
		t.Errorf("expected 1 source, got %d", len(sources))
	}
}

func TestFindShelfEntry_MultipleEntries(t *testing.T) {
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "+a.rh"), []byte(""), 0644)
	os.WriteFile(filepath.Join(tmpDir, "+b.rh"), []byte(""), 0644)

	_, _, err := FindShelfEntry(tmpDir)
	if err == nil {
		t.Error("expected error for multiple entry points, got nil")
	}
}
