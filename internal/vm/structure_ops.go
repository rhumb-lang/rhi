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
	// Check operands for Subscription pattern (Map <> Selector)
	// rhs := vm.peek(0)
	lhs := vm.peek(1)
	
	if lhs.Type == mapval.ValObject {
		if _, ok := lhs.Obj.(*mapval.Map); ok {
			// Subscription!
			vm.opSubscribe()
			return
		}
	}

	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opAccessNested() {
	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opMatchStruct() {
	// Stub for OP_MATCH_STRUCT
	// Check if subject matches structure.
	// For now, always false or check basic types?
	// We need operands to know what to match.
	// Assume it consumes operands or has args?
	// Compiler emitted it with 1 arg (index).
	
	idx := vm.readByte()
	// Get constant?
	// This opcode was a placeholder in the compiler for Signal name match.
	// But we are moving to OP_MATCH_TUPLE.
	// So just consume stack?
	// Stack: [Subject]
	// If we are here, compiler generated it.
	// But we are replacing it with OP_MATCH_TUPLE in compiler.
	// So this might not be reached if we update compiler correctly.
	_ = idx
	vm.push(mapval.NewBoolean(false))
}

func (vm *VM) opSelect() {
	// OP_SELECT placeholder
	// Usually pushes True to start match block?
}

func (vm *VM) opMatchTuple() {
	kind := mapval.TupleKind(vm.readByte())
	
	// Read Topic Index (2 bytes in new design, but readByte is 1 byte)
	// Architecture says 2 bytes for TopicIndex.
	// vm.readShort()
	topicIdx := vm.readShort()
	
	// Peek Subject
	subject := vm.peek(0)
	
	match := false
	if subject.Type == mapval.ValObject {
		if t, ok := subject.Obj.(*mapval.Tuple); ok {
			if t.Kind == kind {
				// Check Topic
				frame := vm.currentFrame()
				expectedTopic := frame.Closure.Fn.Chunk.Constants[topicIdx].Str
				if t.Topic == expectedTopic {
					match = true
				}
			}
		}
	}
	
	vm.push(mapval.NewBoolean(match))
}
