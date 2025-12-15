package loader

import (
	"testing"

	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func TestEngineConstraint(t *testing.T) {
	l := &LibraryLoader{
		CurrentRuntime: mapval.NewVersion(0, 4, 0, false),
	}

	// Case 1: Compatible Exact
	catOK := &Catalog{Engines: map[string]string{"rhi": "0.4.0"}}
	if err := l.validateCatalog(catOK, "ok"); err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	// Case 2: Compatible Wildcard Major
	catWild := &Catalog{Engines: map[string]string{"rhi": "0.4.-"}}
	if err := l.validateCatalog(catWild, "wild"); err != nil {
		t.Errorf("Expected success, got %v", err)
	}

	// Case 3: Incompatible Major
	catFailMaj := &Catalog{Engines: map[string]string{"rhi": "1.0.0"}}
	if err := l.validateCatalog(catFailMaj, "failMaj"); err == nil {
		t.Error("Expected error for major mismatch, got nil")
	}

	// Case 4: Incompatible Minor
	catFailMin := &Catalog{Engines: map[string]string{"rhi": "0.5.0"}}
	if err := l.validateCatalog(catFailMin, "failMin"); err == nil {
		t.Error("Expected error for minor mismatch, got nil")
	}
}
