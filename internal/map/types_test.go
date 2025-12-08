package mapval

import (
	"testing"
)

func TestValue_Version(t *testing.T) {
	// Test Creation and Unpacking
	v1 := NewVersion(1, 2, 3)
	if v1.Type != ValVersion {
		t.Errorf("Expected Type ValVersion, got %v", v1.Type)
	}

	maj, min, pat := v1.VersionUnpack()
	if maj != 1 || min != 2 || pat != 3 {
		t.Errorf("Expected 1.2.3, got %d.%d.%d", maj, min, pat)
	}

	// Test Boundaries
	// Max uint16 for Major/Minor, Max uint32 for Patch
	vMax := NewVersion(65535, 65535, 4294967295)
	maj, min, pat = vMax.VersionUnpack()
	if maj != 65535 || min != 65535 || pat != 4294967295 {
		t.Errorf("Boundary check failed. Got %d.%d.%d", maj, min, pat)
	}

	// Test Equality
	vA := NewVersion(1, 0, 0)
	vB := NewVersion(1, 0, 0)
	vC := NewVersion(1, 0, 1)

	if vA != vB {
		t.Errorf("Expected vA == vB for identical versions")
	}
	if vA == vC {
		t.Errorf("Expected vA != vC for different versions")
	}
}

func TestValue_Key(t *testing.T) {
	// Test Creation
	k1 := NewKey(100)
	if k1.Type != ValKey {
		t.Errorf("Expected Type ValKey, got %v", k1.Type)
	}
	if k1.KeyID() != 100 {
		t.Errorf("Expected KeyID 100, got %d", k1.KeyID())
	}

	// Test Equality
	k1Copy := NewKey(100)
	k2 := NewKey(101)

	if k1 != k1Copy {
		t.Errorf("Expected k1 == k1Copy for same ID")
	}
	if k1 == k2 {
		t.Errorf("Expected k1 != k2 for different IDs")
	}

	// Test Zero/Negative ID (should be valid as raw int64)
	kZero := NewKey(0)
	if kZero.KeyID() != 0 {
		t.Errorf("Expected KeyID 0, got %d", kZero.KeyID())
	}
}
