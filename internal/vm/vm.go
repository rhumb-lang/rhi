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
	panic("Expected number")
}
