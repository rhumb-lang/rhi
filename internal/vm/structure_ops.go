package vm

import (
	"strconv"

	mapval "github.com/rhumb-lang/rhi/internal/map"
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
		return
	}

	// Map concat
	if a.Type == mapval.ValObject && b.Type == mapval.ValObject {
		mapA, okA := a.Obj.(*mapval.Map)
		mapB, okB := b.Obj.(*mapval.Map)
		if okA && okB {
			newMap := mapval.NewMap()

			// 1. Copy A
			newMap.Legend.Fields = append(newMap.Legend.Fields, mapA.Legend.Fields...)
			newMap.Fields = append(newMap.Fields, mapA.Fields...)

			// Calc max positional in A
			maxIdx := 0
			for _, f := range mapA.Legend.Fields {
				if idx, err := strconv.Atoi(f.Name); err == nil && idx > maxIdx {
					maxIdx = idx
				}
			}

			// 2. Append B (Renaming positionals)
			for i, desc := range mapB.Legend.Fields {
				newName := desc.Name
				if idx, err := strconv.Atoi(desc.Name); err == nil && idx > 0 {
					newName = strconv.Itoa(maxIdx + idx)
				}

				newMap.Legend.Fields = append(newMap.Legend.Fields, mapval.FieldDesc{
					Name: newName,
					Kind: desc.Kind,
				})
				newMap.Fields = append(newMap.Fields, mapB.Fields[i])
			}

			vm.push(mapval.Value{Type: mapval.ValObject, Obj: newMap})
			return
		}
	}

	// Fallback
	vm.push(mapval.NewEmpty())
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

func (vm *VM) opForeach() error {
	// a <> b
	rhs := vm.pop()
	lhs := vm.pop()

	// Case 1: Subscription (Map <> Selector)
	if lhs.Type == mapval.ValObject {
		if _, ok := lhs.Obj.(*mapval.Map); ok {
			// This is a subscription.
			// Re-push in order: [Map, Selector]
			vm.push(lhs)
			vm.push(rhs)
			return vm.opSubscribe()
		}
	}

	// Case 2: Range Iteration (Range <> Closure)
	if lhs.Type == mapval.ValRange && rhs.Type == mapval.ValObject {
		rng := lhs.Obj.(*mapval.Range)
		if closure, ok := rhs.Obj.(*mapval.Closure); ok {
			// Loop
			start := rng.Start
			end := rng.End
			step := int64(1)
			if start > end {
				step = -1
			}

			for i := start; ; i += step {
				// Execute Closure(i)

				// Push Closure (Function)
				vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})
				// Push Arg (i)
				vm.push(mapval.NewInt(i))

				// Setup Frame
				// We can reuse opCall logic?
				// opCall expects args on stack.
				// opCall sets up frame but DOES NOT RUN.
				// So we call opCall, then RunSynchronous.

				// opCall checks peek(argCount).
				// Arg count is 1.
				// But opCall reads argCount from Bytecode stream!
				// We are in Go code, we don't have an instruction stream for "CALL 1".
				// So we must manually setup frame.

				newFrame := &CallFrame{
					Parent:  vm.CurrentFrame,
					Closure: closure,
					IP:      0,
					Base:    vm.SP - 1, // 1 Arg (i)
				}

				vm.CurrentFrame = newFrame

				// Run
				res, err := vm.RunSynchronous()
				if err != nil {
					return err
				}
				if res != Ok {
					// Propagate Halt/Error?
					// If Halt, stop everything.
					return nil // Stop loop? Or return error?
				}

				// Pop result of closure
				vm.pop()

				if i == end {
					break
				}
			}

			vm.push(mapval.NewEmpty())
			return nil
		}
	}

	// Fallback
	vm.push(mapval.NewEmpty())
	return nil
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

	// Read Topic Index (2 bytes)
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
