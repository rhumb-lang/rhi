package compiler

import (
	"os"
	"path/filepath"
	"testing"

	mapval "github.com/rhumb-lang/rhi/internal/map"
	"github.com/rhumb-lang/rhi/internal/vm"
)

func TestCompileShelf(t *testing.T) {
	// 1. Setup Temp Shelf
	tmpDir := t.TempDir()

	// File A: Defines x
	codeA := `x := 10`
	pathA := filepath.Join(tmpDir, "a.rh")
	os.WriteFile(pathA, []byte(codeA), 0644)

	// File B: Uses x, defines y
	codeB := `y := x ++ 5`
	pathB := filepath.Join(tmpDir, "b.rh")
	os.WriteFile(pathB, []byte(codeB), 0644)

	// Entry Point: Uses y, defines z
	codeEntry := `z := y ** 2`
	pathEntry := filepath.Join(tmpDir, "+entry.rh")
	os.WriteFile(pathEntry, []byte(codeEntry), 0644)

	// 2. Compile
	c := NewCompiler()
	chunk, err := c.CompileShelf([]string{pathA, pathB}, pathEntry)
	if err != nil {
		t.Fatalf("CompileShelf failed: %v", err)
	}

	// 3. Execute with VM
	v := vm.NewVM()
	v.Config.TraceStack = true // Enable tracing for debugging
	// CallAndReturn expects a chunk that returns a value (the Map)
	res, err := v.CallAndReturn(chunk)
	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	// 4. Assertions
	if res.Type != mapval.ValObject {
		t.Fatalf("Expected object result, got %s: %s", res.Type, res)
	}
	m, ok := res.Obj.(*mapval.Map)
	if !ok {
		t.Fatalf("Expected Map object, got %T", res.Obj)
	}

	// Helper to check map fields
	checkField := func(name string, expected int64) {
		// Linear search in Legend/Fields (simplified for test)
		found := false
		for i, field := range m.Legend.Fields {
			if field.Name == name {
				val := m.Fields[i]
				if val.Type != mapval.ValInteger || val.Integer != expected {
					t.Errorf("Field %s: expected %d, got %s", name, expected, val)
				}
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Field %s not found in export map", name)
		}
	}

	checkField("x", 10)
	checkField("y", 15) // 10 + 5
	checkField("z", 30) // 15 * 2
}
