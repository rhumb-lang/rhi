package vm

import (
	"fmt"

	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) run() (Result, error) {
	var frame *CallFrame
	var chunk *mapval.Chunk

	// Cache current frame
	frame = vm.currentFrame()
	chunk = frame.Closure.Fn.Chunk

	for {
		// Re-fetch frame context each iter (for now, simpler than handling call/ret)
		// Optimization: Only do this after CALL/RET/RETURN
		frame = vm.currentFrame()
		chunk = frame.Closure.Fn.Chunk

		// Fetch
		if frame.IP >= len(chunk.Code) {
			return RuntimeError, fmt.Errorf("IP out of bounds")
		}

		instruction := mapval.OpCode(chunk.Code[frame.IP])
		frame.IP++

		// Decode & Execute
		switch instruction {
		case mapval.OP_HALT:
			return Ok, nil

		case mapval.OP_LOAD_CONST:
			idx := chunk.Code[frame.IP]
			frame.IP++
			constant := chunk.Constants[idx]
			vm.push(constant)

		case mapval.OP_LOAD_LOC:
			idx := chunk.Code[frame.IP]
			frame.IP++
			// Local is relative to Frame Base
			val := vm.Stack[frame.Base+int(idx)]
			vm.push(val)

		case mapval.OP_STORE_LOC:
			idx := chunk.Code[frame.IP]
			frame.IP++
			val := vm.peek(0)
			vm.Stack[frame.Base+int(idx)] = val

		case mapval.OP_JUMP:
			offset := vm.readShort()
			frame.IP += offset

		case mapval.OP_DUP:
			val := vm.peek(0)
			vm.push(val)

		case mapval.OP_POP:
			vm.pop()

		case mapval.OP_IF_TRUE:
			offset := vm.readShort()
			val := vm.pop()
			if vm.isFalsy(val) {
				frame.IP += offset
			}

		case mapval.OP_IF_FALSE:
			offset := vm.readShort()
			val := vm.pop()
			if vm.isTruthy(val) {
				frame.IP += offset
			}

		case mapval.OP_MAKE_MAP:
			m := &mapval.Map{
				Legend: mapval.NewLegend(),
				Fields: make([]mapval.Value, 0),
			}
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: m})

		case mapval.OP_ASSIGN_IMM:
			val := vm.pop()
			key := vm.pop()
			target := vm.pop()

			res, err := vm.setField(target, key, val, false)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(res)

		case mapval.OP_ASSIGN_MUT:
			val := vm.pop()
			key := vm.pop()
			target := vm.pop()

			res, err := vm.setField(target, key, val, true)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(res)

		case mapval.OP_SEND:
			idx := chunk.Code[frame.IP]
			frame.IP++
			key := chunk.Constants[idx] // Key

			target := vm.pop() // Receiver

			if target.Type != mapval.ValObject {
				return RuntimeError, fmt.Errorf("target is not a map (SEND)")
			}
			m, ok := target.Obj.(*mapval.Map)
			if !ok {
				return RuntimeError, fmt.Errorf("target is not a map (SEND)")
			}

			keyStr := ""
			if key.Type == mapval.ValText {
				keyStr = key.Str
			} else {
				return RuntimeError, fmt.Errorf("invalid key type for SEND: %v", key.Type)
			}

			found := false
			for i, fd := range m.Legend.Fields {
				if fd.Name == keyStr {
					vm.push(m.Fields[i])
					found = true
					break
				}
			}

			if !found {
				vm.push(mapval.NewEmpty())
			}

		case mapval.OP_MAKE_FN:
			idx := chunk.Code[frame.IP]
			frame.IP++
			fnVal := chunk.Constants[idx]
			// Assume it's a Function object
			fn := fnVal.Obj.(*mapval.Function)
			closure := &mapval.Closure{Fn: fn}
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})

		case mapval.OP_CALL:
			argCount := int(chunk.Code[frame.IP])
			frame.IP++

			// Callee is below args
			calleeVal := vm.peek(argCount)
			if calleeVal.Type != mapval.ValObject {
				return RuntimeError, fmt.Errorf("can only call closures")
			}

			closure, ok := calleeVal.Obj.(*mapval.Closure)
			if !ok {
				return RuntimeError, fmt.Errorf("can only call closures")
			}

			if argCount != closure.Fn.Arity {
				return RuntimeError, fmt.Errorf("arity mismatch: expected %d, got %d", closure.Fn.Arity, argCount)
			}

			if vm.FrameCount >= MaxFrames {
				return RuntimeError, fmt.Errorf("stack overflow")
			}

			// Push new frame
			vm.Frames[vm.FrameCount] = CallFrame{
				Closure: closure,
				IP:      0,
				Base:    vm.SP - argCount - 1,
			}
			vm.FrameCount++
			// Loop will pick up new frame

		case mapval.OP_RETURN:
			result := vm.pop()

			vm.FrameCount--
			if vm.FrameCount == 0 {
				vm.pop() // Pop Main Script Closure?
				return Ok, nil
			}

			// Restore SP to before callee
			// frame is the frame we are returning FROM
			// vm.SP = frame.Base
			vm.SP = frame.Base
			vm.push(result)
			// Loop will pick up parent frame

		case mapval.OP_ADD:
			b := vm.pop()
			a := vm.pop()
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer + b.Integer))
			} else if a.Type == mapval.ValFloat || b.Type == mapval.ValFloat {
				fa := asFloat(a)
				fb := asFloat(b)
				vm.push(mapval.NewFloat(fa + fb))
			} else {
				return RuntimeError, fmt.Errorf("operands must be numbers for ADD")
			}

		case mapval.OP_SUB:
			b := vm.pop()
			a := vm.pop()
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer - b.Integer))
			} else {
				vm.push(mapval.NewFloat(asFloat(a) - asFloat(b)))
			}

		case mapval.OP_MULT:
			b := vm.pop()
			a := vm.pop()
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer * b.Integer))
			} else {
				vm.push(mapval.NewFloat(asFloat(a) * asFloat(b)))
			}

		case mapval.OP_DIV_FLOAT:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewFloat(asFloat(a) / asFloat(b)))

		case mapval.OP_EQ:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewBoolean(isEqual(a, b)))

		case mapval.OP_NEQ:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewBoolean(!isEqual(a, b)))

		case mapval.OP_GT:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(mapval.NewBoolean(res > 0))

		case mapval.OP_LT:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(mapval.NewBoolean(res < 0))

		case mapval.OP_GTE:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(mapval.NewBoolean(res >= 0))

		case mapval.OP_LTE:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil {
				return RuntimeError, err
			}
			vm.push(mapval.NewBoolean(res <= 0))

		default:
			return RuntimeError, fmt.Errorf("unknown opcode: %d", instruction)
		}
	}
}
