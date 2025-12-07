package vm

import (
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) opCoalesce() {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValEmpty {
		vm.push(b)
	} else {
		vm.push(a)
	}
}

func (vm *VM) opConcat() {
	b := vm.pop()
	a := vm.pop()
	// String concat
	if a.Type == mapval.ValText && b.Type == mapval.ValText {
		vm.push(mapval.NewText(a.Str + b.Str))
	} else {
		// TODO: List concat
		// For now, error or empty
		vm.push(mapval.NewEmpty())
	}
}

func (vm *VM) opRange() {
	// a | b
	b := vm.pop()
	a := vm.pop()
	
	if a.Type != mapval.ValInteger || b.Type != mapval.ValInteger {
		// TODO: Handle non-integer ranges?
		vm.push(mapval.NewEmpty())
		return
	}
	
	r := &mapval.Range{Start: a.Integer, End: b.Integer}
	vm.push(mapval.Value{Type: mapval.ValRange, Obj: r}) // Use ValRange tag
}

func (vm *VM) opHasSub() {
	// a =@ b
	vm.pop() // b
	vm.pop() // a
	// Check if a has subfield b
	// Stub
	vm.push(mapval.NewBoolean(false))
}

func (vm *VM) opPipe() {
	// a || b -> b(a)
	// TODO: Implement functional pipe
	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opForeach() {
	// a <> b
	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opAccessNested() {
	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
}
