package testing

import (
	"io/ioutil"
	"strings"
	"testing"

	"git.sr.ht/~madcapjake/rhi/internal/runner"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func RunRhumbTest(t *testing.T, path string) {
	t.Helper()

	code, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("test panicked: %v", r)
		}
	}()

	v := vm.New()
	_, err = runner.RunString(string(code), v)
	if err != nil {
		t.Fatalf("vm execution failed: %v", err)
	}

	assertionCount := strings.Count(string(code), "%=")
	if assertionCount != v.AssertionCount {
		t.Errorf("expected %d assertions, but found %d", assertionCount, v.AssertionCount)
	}

	if len(v.Assertions) > 0 {
		t.Errorf("%d of %d assertions failed:", len(v.Assertions), assertionCount)
		for _, err := range v.Assertions {
			t.Error(err)
		}
	}
}
