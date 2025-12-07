package main

import (
	"path/filepath"
	"testing"

	rhtesting "git.sr.ht/~madcapjake/rhi/internal/testing"
)

func TestRhumbFiles(t *testing.T) {
	files, err := filepath.Glob("tests/*.rhs")
	if err != nil {
		t.Fatalf("failed to find test files: %v", err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			rhtesting.RunRhumbTest(t, file)
		})
	}
}
