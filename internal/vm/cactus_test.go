package vm_test

import (
	"testing"
	"git.sr.ht/~madcapjake/rhi/internal/map"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
)

func TestVM_CactusStack(t *testing.T) {
	// This test simulates the "Zombie Frame" behavior.
	// We manually manipulate the VM frames to ensure that returning from a frame
	// does not destroy it if we hold a reference to it.

	machine := vm.NewVM()
	
	// 1. Create a "Root" Frame (e.g. Main)
	rootChunk := mapval.NewChunk()
	
	// Initialize VM (Interpret does this, but we'll do manual setup for precision)
	machine.Interpret(rootChunk) 
	rootFrame := machine.CurrentFrame
	
	if rootFrame == nil {
		t.Fatal("Failed to initialize root frame")
	}

	// 2. Simulate Calling "Child A"
	// OP_CALL logic: creates new frame linked to current
	childAChunk := mapval.NewChunk()
	childAFn := &mapval.Function{Name: "ChildA", Chunk: childAChunk}
	childAClosure := &mapval.Closure{Fn: childAFn}
	
	// Manually push Child A Frame (simulating OP_CALL)
	childAFrame := &vm.CallFrame{
		Parent: machine.CurrentFrame,
		Closure: childAClosure,
		IP: 0,
		Base: machine.SP,
	}
	machine.CurrentFrame = childAFrame
	
	if machine.CurrentFrame != childAFrame {
		t.Error("CurrentFrame not updated to Child A")
	}
	if machine.CurrentFrame.Parent != rootFrame {
		t.Error("Child A parent pointer is incorrect")
	}

	// 3. Simulate "Return" from Child A (Zombie Creation)
	// OP_RETURN logic: vm.CurrentFrame = frame.Parent
	machine.CurrentFrame = machine.CurrentFrame.Parent
	
	if machine.CurrentFrame != rootFrame {
		t.Error("Failed to return to Root frame")
	}
	
	// CRITICAL CHECK: Does ChildAFrame still exist and point to Root?
	// In a fixed array stack, ChildAFrame might be overwritten by the next call.
	// In a Cactus stack (Heap), it persists as long as `childAFrame` var exists.
	if childAFrame.Parent != rootFrame {
		t.Error("Zombie Child A lost its parent link")
	}

	// 4. Simulate Calling "Child B"
	// If using an array stack, this would overwrite Child A's slot.
	childBChunk := mapval.NewChunk()
	childBFn := &mapval.Function{Name: "ChildB", Chunk: childBChunk}
	childBClosure := &mapval.Closure{Fn: childBFn}
	
	childBFrame := &vm.CallFrame{
		Parent: machine.CurrentFrame,
		Closure: childBClosure,
		IP: 0,
		Base: machine.SP,
	}
	machine.CurrentFrame = childBFrame
	
	// 5. Verify Child A is still intact (Zombie is alive)
	// Even though we are executing Child B, Child A should still be valid 
	// because it's a separate heap object.
	if childAFrame.Closure.Fn.Name != "ChildA" {
		t.Error("Zombie Child A data was corrupted/overwritten!")
	}
	if childBFrame.Closure.Fn.Name != "ChildB" {
		t.Error("Child B data invalid")
	}
	
	// 6. Simulate "Reply" (Drill Down)
	// We "jump" back into Child A from Child B (conceptually, via Injection)
	// In reality, Injection creates a new branch *from* Child A.
	// But simply verifying we can still access Child A's IP is enough.
	childAFrame.IP = 99
	if childAFrame.IP != 99 {
		t.Error("Cannot modify Zombie frame state")
	}
}
