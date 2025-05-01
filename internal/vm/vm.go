package vm

import (
	"arena"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	obj "git.sr.ht/~madcapjake/rhi/internal/object"
	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

const MEM_SIZE = 5000
const STK_SIZE = 500
const SCP_SIZE = 50

type VirtualMachine struct {
	Memory *arena.Arena               // arena where all objects live
	Frames *stack.Stack[*obj.Routine] // any currently open routines
}

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)
	vm.Memory = arena.NewArena()
	vm.Frames = stack.NewStack[*obj.Routine](vm.Memory)
	vm.Frames.Push(obj.NewRoutine(vm.Memory, vm.Frames, nil))
	return vm
}

func (vm *VirtualMachine) Run(routine *obj.Routine) (obj.Any, error) {
	return routine.Run()
}

func (vm *VirtualMachine) Disassemble(routine *obj.Routine) {
	routine.Disassemble()
}

func (vm *VirtualMachine) PushRoutine() {
	topFrame := *vm.Frames.Top()
	vm.Frames.Push(obj.NewRoutine(vm.Memory, vm.Frames, topFrame.Scope()))
}

func (vm *VirtualMachine) PopRoutine() *obj.Routine {
	return vm.Frames.Pop()
}

func (vm *VirtualMachine) RegisterObject(o obj.Any) int {
	topFrame := *vm.Frames.Top()
	return topFrame.Register(o)
}

func (vm *VirtualMachine) Write(c ...code.Any) {
	topFrame := *vm.Frames.Top()
	topFrame.Write(c...)
}

func (vm *VirtualMachine) RunMain() (obj.Any, error) {
	topFrame := *vm.Frames.Top()
	return vm.Run(topFrame)
}

func (vm *VirtualMachine) DisassembleMain() {
	topFrame := *vm.Frames.Top()
	vm.Disassemble(topFrame)
}
