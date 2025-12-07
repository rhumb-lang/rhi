package vm

import (
	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

const StackMax = 2048
const MaxFrames = 64

type VM struct {
	Frames     [MaxFrames]CallFrame
	FrameCount int

	Stack [StackMax]mapval.Value
	SP    int // Stack Pointer (points to empty slot)
	
	Debug bool // Enable instruction tracing
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
		SP:         0,
		FrameCount: 0,
		Debug:      false,
	}
}

func (vm *VM) currentFrame() *CallFrame {
	return &vm.Frames[vm.FrameCount-1]
}

// Interpret executes the chunk.
func (vm *VM) Interpret(chunk *mapval.Chunk) (Result, error) {
	// Wrap chunk in script function/closure
	fn := &mapval.Function{
		Name:  "<script>",
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