package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/code"
	obj "git.sr.ht/~madcapjake/grhumb/internal/object"
)

const MEM_SIZE = 5000
const STK_SIZE = 500
const SCP_SIZE = 50

type VirtualMachine struct {
	Memory [MEM_SIZE]obj.Any
	Stack  [STK_SIZE]obj.Any
	Scope  [SCP_SIZE]map[string]obj.Ref[obj.Any]
	Main   obj.Ref[RoutineMap]
	Top    struct {
		Memory int
		Stack  int
		Scope  int
	}
}

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)
	vm.Scope[0] = map[string]obj.Ref[obj.Any]{}
	legendRef := Allocate(vm, NewRoutineLegend())
	main := RoutineMap{vm, legendRef, []obj.Any{}}
	mainRef := Allocate(vm, main)
	vm.Main = mainRef
	return vm
}

func Get[T obj.Any](vm *VirtualMachine, ref obj.Ref[T]) T {
	return vm.Memory[ref.Value].(T)
}

func Set[T obj.Any, U obj.Any](
	vm *VirtualMachine,
	this obj.Ref[T],
	that U,
) obj.Ref[U] {
	vm.Memory[this.Value] = that
	return obj.Ref[U](this)
}

func Allocate[T obj.Any](vm *VirtualMachine, o T) obj.Ref[T] {
	if i := vm.Top.Memory + 1; i < MEM_SIZE {
		vm.Memory[i] = o
		vm.Top.Memory = i
		return obj.Ref[T]{i}
	} else {
		panic("out of memory")
	}
}

func PushStack[O obj.Any](vm *VirtualMachine, obj O) int {
	if i := vm.Top.Stack + 1; i < STK_SIZE {
		vm.Stack[i] = obj
		vm.Top.Stack = i
		return i
	} else {
		panic("out of stack")
	}
}

func PopStack(vm *VirtualMachine) obj.Any {
	i := vm.Top.Stack
	result := vm.Stack[i]
	vm.Top.Stack = i - 1
	return result
}

func EnterScope(vm *VirtualMachine) int {
	if i := vm.Top.Scope + 1; i < SCP_SIZE {
		vm.Scope[i] = map[string]obj.Ref[obj.Any]{}
		vm.Top.Scope = i
		return i
	} else {
		panic("out of scope")
	}
}

func Run(vm *VirtualMachine, codes []code.Any) (obj.Any, error) {
	scopeIndex := EnterScope(vm)
	for _, code := range codes {
		Execute(vm, scopeIndex, code)
	}
	return vm.Stack[vm.Top.Stack], nil
}

func Disassemble(vm *VirtualMachine, codes []code.Any) (obj.Any, error) {
	for _, code := range codes {
		fmt.Println(code.WHAT())
	}
	return vm.Stack[vm.Top.Stack], nil
}

func RunMain(vm *VirtualMachine) (obj.Any, error) {
	main := Get(vm, vm.Main)
	return main.Invoke()
}

func DisassembleMain(vm *VirtualMachine) (obj.Any, error) {
	main := Get(vm, vm.Main)
	for _, code := range main.Codes() {
		fmt.Println(code.WHAT())
	}
	return vm.Stack[vm.Top.Stack], nil
}

func Execute(vm *VirtualMachine, sIdx int, code code.Any) {
	fmt.Println(code.WHAT())
}

func WriteCode[C code.Any](
	vm *VirtualMachine,
	routineRef obj.Ref[RoutineMap],
	code C,
) {
	routine := Get(vm, routineRef)
	legend := Get(vm, routine.legend)
	codes := legend.Data
	legend.Data = append(codes, code)
	Set(vm, routine.legend, legend)
}
