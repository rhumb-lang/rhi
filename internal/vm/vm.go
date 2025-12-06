package vm

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

const StackMax = 2048
const MaxFrames = 64

type VM struct {
	Frames [MaxFrames]CallFrame
	FrameCount int
	
	Stack [StackMax]mapval.Value
	SP    int // Stack Pointer (points to empty slot)
}

// Result code for the VM interpretation
type Result int

const (
	Ok Result = iota
	CompileError
	RuntimeError
)

func NewVM() *VM {
	return &VM{
		SP: 0,
		FrameCount: 0,
	}
}

func (vm *VM) currentFrame() *CallFrame {
	return &vm.Frames[vm.FrameCount-1]
}

// Interpret executes the chunk.
func (vm *VM) Interpret(chunk *mapval.Chunk) (Result, error) {
	// Wrap chunk in script function/closure
	fn := &mapval.Function{
		Name: "<script>",
		Chunk: chunk,
	}
	closure := &mapval.Closure{Fn: fn}
	
	vm.Frames[0] = CallFrame{
		Closure: closure,
		IP:      0,
		Base:    0,
	}
	vm.FrameCount = 1
	
	return vm.run()
}

func (vm *VM) push(val mapval.Value) {
	if vm.SP >= StackMax {
		panic("Stack overflow")
	}
	vm.Stack[vm.SP] = val
	vm.SP++
}

func (vm *VM) pop() mapval.Value {
	if vm.SP == 0 {
		panic("Stack underflow")
	}
	vm.SP--
	return vm.Stack[vm.SP]
}

// peek returns value at distance from top (0 is top)
func (vm *VM) peek(distance int) mapval.Value {
	return vm.Stack[vm.SP-1-distance]
}

func (vm *VM) readShort() int {
	frame := vm.currentFrame()
	chunk := frame.Closure.Fn.Chunk
	
	byte1 := chunk.Code[frame.IP]
	frame.IP++
	byte2 := chunk.Code[frame.IP]
	frame.IP++
	return (int(byte1) << 8) | int(byte2)
}

func (vm *VM) isFalsy(val mapval.Value) bool {
	return val.Type == mapval.ValEmpty || (val.Type == mapval.ValBoolean && val.Integer == 0)
}

func (vm *VM) isTruthy(val mapval.Value) bool {
	return !vm.isFalsy(val)
}

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
			val := vm.Stack[frame.Base + int(idx)]
			vm.push(val)

		case mapval.OP_STORE_LOC:
			idx := chunk.Code[frame.IP]
			frame.IP++
			val := vm.peek(0)
			vm.Stack[frame.Base + int(idx)] = val
			
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
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res > 0))

		case mapval.OP_LT:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res < 0))

		case mapval.OP_GTE:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res >= 0))

		case mapval.OP_LTE:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res <= 0))

		default:
			return RuntimeError, fmt.Errorf("unknown opcode: %d", instruction)
		}
	}
}

func asFloat(v mapval.Value) float64 {
	if v.Type == mapval.ValInteger {
		return float64(v.Integer)
	}
	if v.Type == mapval.ValFloat {
		return v.Float
	}
	// TODO: Handle other types that can coerce to float
	return 0.0 // Default or error
}

// isEqual compares two values for equality.
// Handles different types based on Rhumb semantics.
func isEqual(a, b mapval.Value) bool {
	if a.Type != b.Type {
		// Attempt coercion for numbers
		if (a.Type == mapval.ValInteger || a.Type == mapval.ValFloat) && (b.Type == mapval.ValInteger || b.Type == mapval.ValFloat) {
			return asFloat(a) == asFloat(b)
		}
		return false
	}
	
	switch a.Type {
	case mapval.ValInteger: return a.Integer == b.Integer
	case mapval.ValFloat: return a.Float == b.Float
	case mapval.ValBoolean: return a.Integer == b.Integer
	case mapval.ValEmpty: return true
	case mapval.ValText: return a.Str == b.Str
	// TODO: Keys, Dates, Versions, Objects
	default: return false
	}
}

// numericCompare compares two numeric values. Returns -1 if a<b, 0 if a==b, 1 if a>b.
func numericCompare(a, b mapval.Value) (int, error) {
	if (a.Type != mapval.ValInteger && a.Type != mapval.ValFloat) || (b.Type != mapval.ValInteger && b.Type != mapval.ValFloat) {
		return 0, fmt.Errorf("operands must be numbers for comparison")
	}

	fa := asFloat(a)
	fb := asFloat(b)

	if fa < fb { return -1, nil }
	if fa > fb { return 1, nil }
	return 0, nil
}