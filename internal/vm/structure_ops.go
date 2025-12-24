package vm

import (
	"fmt"
	"strconv"

	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (vm *VM) opLength() {
	val := vm.pop()
	res := int64(0)
	switch val.Type {
	case mapval.ValText:
		res = int64(len(val.Str))
	case mapval.ValObject:
		if m, ok := val.Obj.(*mapval.Map); ok {
			res = int64(len(m.Fields))
		} else if t, ok := val.Obj.(*mapval.Tuple); ok {
			res = int64(len(t.Payload))
		}
	}
	vm.push(mapval.NewInt(res))
}

func (vm *VM) opFreeze() {
	// [.] - Freeze target
	// For MVP, we can just return the value as is, or mark it immutable if possible.
	// Since opStoreLocImmut handles immutability at variable level, this might be for objects.
	// Placeholder: Identity
}

func (vm *VM) opCopy() {
	// [:] - Copy target
	// Should perform shallow copy of Map or List
	val := vm.pop()
	if val.Type == mapval.ValObject {
		if m, ok := val.Obj.(*mapval.Map); ok {
			newMap := mapval.NewMap()
			// Copy Legend
			newMap.Legend.Fields = make([]mapval.FieldDesc, len(m.Legend.Fields))
			copy(newMap.Legend.Fields, m.Legend.Fields)
			// Copy Fields
			newMap.Fields = make([]mapval.Value, len(m.Fields))
			copy(newMap.Fields, m.Fields)
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: newMap})
			return
		}
	}
	vm.push(val)
}

func (vm *VM) opToDate() {
	// [/] - To Date
	// Coerce Int/Float/String to Date
	val := vm.pop()
	res := mapval.NewEmpty()
	if val.Type == mapval.ValInteger {
		res = mapval.Value{Type: mapval.ValDateTime, Integer: val.Integer}
	}
	vm.push(res)
}

func (vm *VM) opGetCtor() {
	// [^] - Get Constructor
	// Placeholder
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opGetBase() {
	// [!] - Get Base
	// Placeholder
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opToBool() {
	val := vm.pop()
	vm.push(mapval.NewBoolean(vm.isTruthy(val)))
}

func (vm *VM) opBoolNeg() {
	val := vm.pop()
	vm.push(mapval.NewBoolean(vm.isFalsy(val)))
}

func (vm *VM) opSpread() {
	// [&] - Spread
	// Placeholder
	vm.pop()
	vm.push(mapval.NewEmpty())
}

func (vm *VM) opToKey() {
	// [`] - To Key
	// Placeholder
	vm.pop()
	vm.push(mapval.NewEmpty())
}

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
	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: opRange %d | %d\n", a.Integer, b.Integer)
	}
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

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: opForeach LHS=%s RHS=%s\n", lhs, rhs)
	}

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
				if vm.Config.TraceSpace {
					fmt.Printf("TRACE: opForeach Iteration i=%d\n", i)
				}
				// Execute Closure(i)

				// Push Closure (Function)
				vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})
				// Push Arg (i)
				vm.push(mapval.NewInt(i))

				// Setup Frame
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
					return nil // Stop loop
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

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: opMatchTuple Subject=%s Kind=%d\n", subject, kind)
	}

	match := false
	if subject.Type == mapval.ValObject {
		if t, ok := subject.Obj.(*mapval.Tuple); ok {
			if t.Kind == kind {
				// Check Topic
				frame := vm.currentFrame()
				expectedTopic := frame.Closure.Fn.Chunk.Constants[topicIdx].Str
				if vm.Config.TraceSpace {
					fmt.Printf("TRACE: opMatchTuple Topic=%s Expected=%s\n", t.Topic, expectedTopic)
				}
				if t.Topic == expectedTopic {
					match = true
				}
			} else {
				if vm.Config.TraceSpace {
					fmt.Printf("TRACE: opMatchTuple Kind Mismatch Got=%d Expected=%d\n", t.Kind, kind)
				}
			}
		} else {
			if vm.Config.TraceSpace {
				fmt.Printf("TRACE: opMatchTuple Subject Not Tuple (is %T)\n", subject.Obj)
			}
		}
	} else {
		if vm.Config.TraceSpace {
			fmt.Printf("TRACE: opMatchTuple Subject Not Object (is %s)\n", subject.Type)
		}
	}

	vm.push(mapval.NewBoolean(match))
}
