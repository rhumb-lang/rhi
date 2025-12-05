package vm

import (
	"fmt"
	
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

const StackMax = 2048

type VM struct {
	Chunk *Chunk
	IP    int // Instruction Pointer
	
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
	}
}

// Interpret executes the chunk.
func (vm *VM) Interpret(chunk *Chunk) (Result, error) {
	vm.Chunk = chunk
	vm.IP = 0
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
	byte1 := vm.Chunk.Code[vm.IP]
	vm.IP++
	byte2 := vm.Chunk.Code[vm.IP]
	vm.IP++
	return (int(byte1) << 8) | int(byte2)
}

func (vm *VM) isFalsy(val mapval.Value) bool {
	return val.Type == mapval.ValEmpty || (val.Type == mapval.ValBoolean && val.Integer == 0)
}

func (vm *VM) isTruthy(val mapval.Value) bool {
	return !vm.isFalsy(val)
}

func (vm *VM) run() (Result, error) {
	for {
		// Fetch
		if vm.IP >= len(vm.Chunk.Code) {
			return RuntimeError, fmt.Errorf("IP out of bounds")
		}
		
		instruction := OpCode(vm.Chunk.Code[vm.IP])
		vm.IP++
		
		// Decode & Execute
		switch instruction {
		case OP_HALT:
			return Ok, nil
			
		case OP_LOAD_CONST:
			// Read 1 byte operand (index)
			idx := vm.Chunk.Code[vm.IP]
			vm.IP++
			constant := vm.Chunk.Constants[idx]
			vm.push(constant)
			
		case OP_LOAD_LOC:
			idx := vm.Chunk.Code[vm.IP]
			vm.IP++
			val := vm.Stack[idx]
			vm.push(val)

		case OP_STORE_LOC:
			idx := vm.Chunk.Code[vm.IP]
			vm.IP++
			val := vm.peek(0)
			vm.Stack[idx] = val
			
		case OP_JUMP:
			offset := vm.readShort()
			vm.IP += offset

		case OP_DUP:
			val := vm.peek(0)
			vm.push(val)

		case OP_POP:
			vm.pop()

		case OP_IF_TRUE:
			offset := vm.readShort()
			val := vm.pop()
			if vm.isFalsy(val) {
				vm.IP += offset
			}

		case OP_IF_FALSE:
			offset := vm.readShort()
			val := vm.pop()
			if vm.isTruthy(val) {
				vm.IP += offset
			}
			
		case OP_ADD:
			b := vm.pop()
			a := vm.pop()
			// Check types. For MVP assume Int.
			// TODO: Full type coercion logic
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer + b.Integer))
			} else if a.Type == mapval.ValFloat || b.Type == mapval.ValFloat {
				// Promote to float
				fa := asFloat(a)
				fb := asFloat(b)
				vm.push(mapval.NewFloat(fa + fb))
			} else {
				// String concat?
				return RuntimeError, fmt.Errorf("operands must be numbers for ADD")
			}
			
		case OP_SUB:
			b := vm.pop()
			a := vm.pop()
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer - b.Integer))
			} else {
				vm.push(mapval.NewFloat(asFloat(a) - asFloat(b)))
			}

		case OP_MULT:
			b := vm.pop()
			a := vm.pop()
			if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
				vm.push(mapval.NewInt(a.Integer * b.Integer))
			} else {
				vm.push(mapval.NewFloat(asFloat(a) * asFloat(b)))
			}
			
		case OP_DIV_FLOAT:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewFloat(asFloat(a) / asFloat(b)))

		case OP_EQ:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewBoolean(isEqual(a, b)))
		
		case OP_NEQ:
			b := vm.pop()
			a := vm.pop()
			vm.push(mapval.NewBoolean(!isEqual(a, b)))

		case OP_GT:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res > 0))

		case OP_LT:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res < 0))

		case OP_GTE:
			b := vm.pop()
			a := vm.pop()
			res, err := numericCompare(a, b)
			if err != nil { return RuntimeError, err }
			vm.push(mapval.NewBoolean(res >= 0))

		case OP_LTE:
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
