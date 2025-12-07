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
	val := uint16(byte1)<<8 | uint16(byte2)
	return int(int16(val))
}

func (vm *VM) readByte() byte {
	frame := vm.currentFrame()
	b := frame.Closure.Fn.Chunk.Code[frame.IP]
	frame.IP++
	return b
}

func (vm *VM) run() (Result, error) {
	for {
		frame := vm.currentFrame()
		chunk := frame.Closure.Fn.Chunk
		
		if frame.IP >= len(chunk.Code) {
			return RuntimeError, fmt.Errorf("IP out of bounds")
		}
		
		instruction := mapval.OpCode(chunk.Code[frame.IP])
		frame.IP++
		
		var err error
		
		switch instruction {
		case mapval.OP_HALT: return Ok, nil
		
		// Stack
		case mapval.OP_LOAD_CONST: vm.opLoadConst()
		case mapval.OP_LOAD_LOC:   vm.opLoadLoc()
		case mapval.OP_STORE_LOC:  vm.opStoreLoc()
		case mapval.OP_DUP:        vm.opDup()
		case mapval.OP_POP:        vm.opPop()
		
		// Flow
		case mapval.OP_JUMP:       vm.opJump()
		case mapval.OP_IF_TRUE:    vm.opIfTrue()
		case mapval.OP_IF_FALSE:   vm.opIfFalse()
		case mapval.OP_CALL:       err = vm.opCall()
		case mapval.OP_RETURN:     
			res, e := vm.opReturn()
			if res != 0 { return Ok, nil } // Halt
			if e != nil { err = e }
		case mapval.OP_MAKE_FN:    vm.opMakeFn()
		
		// Map
		case mapval.OP_MAKE_MAP:   vm.opMakeMap()
		case mapval.OP_SEND:       err = vm.opSend()
		case mapval.OP_SET_FIELD:  err = vm.opSetField()
		
		// Space
		case mapval.OP_POST:       err = vm.opPost()
		case mapval.OP_INJECT:     err = vm.opInject()
		case mapval.OP_WRITE:      err = vm.opWrite()
		case mapval.OP_SUBSCRIBE:  err = vm.opSubscribe()
		case mapval.OP_NEW_REALM:  vm.opNewRealm()
		
		// Math
		case mapval.OP_ADD:        err = vm.opAdd()
		case mapval.OP_SUB:        err = vm.opSub()
		case mapval.OP_MULT:       err = vm.opMult()
		case mapval.OP_DIV_FLOAT:  err = vm.opDivFloat()
		case mapval.OP_DIV_INT:    err = vm.opDivInt()
		case mapval.OP_MOD:        err = vm.opMod()
		case mapval.OP_POW:        err = vm.opPow()
		case mapval.OP_ROOT:       err = vm.opRoot()
		case mapval.OP_SCI_NOT:    err = vm.opSciNot()
		case mapval.OP_DEV:        err = vm.opDev()
		
		// Structure
		case mapval.OP_COALESCE:   vm.opCoalesce()
		case mapval.OP_CONCAT:     vm.opConcat()
		case mapval.OP_RANGE:      vm.opRange()
		case mapval.OP_HAS_SUBFIELD: vm.opHasSub()
		case mapval.OP_ACCESS_NESTED: vm.opAccessNested()
		case mapval.OP_PIPE:       vm.opPipe()
		case mapval.OP_FOREACH:    vm.opForeach()
		
		// Logic
		case mapval.OP_AND:        vm.opAnd()
		case mapval.OP_OR:         vm.opOr()
		case mapval.OP_NOT:        vm.opNot()
		
		case mapval.OP_EQ:         vm.opEq()
		case mapval.OP_NEQ:        vm.opNeq()
		case mapval.OP_GT:         err = vm.opGt()
		case mapval.OP_LT:         err = vm.opLt()
		case mapval.OP_GTE:        err = vm.opGte()
		case mapval.OP_LTE:        err = vm.opLte()
		
		default:
			return RuntimeError, fmt.Errorf("unknown opcode: %d", instruction)
		}
		
		if err != nil {
			return RuntimeError, err
		}
	}
}
