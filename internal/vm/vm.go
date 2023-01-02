package vm

import "fmt"

type VirtualMachine struct {
	heap  []Word
	free  []bool
	stack []Word
	scope []map[string]int
	main  Chunk
}

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)
	vm.heap = make([]Word, 0)
	vm.free = make([]bool, 0)
	vm.stack = make([]Word, 0)
	vm.scope = make([]map[string]int, 0)
	vm.scope = append(vm.scope, make(map[string]int))
	vm.main = NewChunk(nil, nil)
	return vm
}

func (vm *VirtualMachine) WriteCodeToMain(
	line int,
	ref InstrRef,
	codeFactory func(i uint32) []Code,
) {
	id := vm.main.AddReference(ref)
	codes := codeFactory(id)
	vm.main.WriteCode(line, codes)
}

func (vm *VirtualMachine) Disassemble() {
	vm.main.Disassemble()
}

func (vm *VirtualMachine) Execute() {
	vm.main.Execute(vm)
}

func logAddedToStack(stack []Word, txt string) {
	fmt.Print("Added ", txt, " to stack: [")
	for s := range stack {
		fmt.Print(" ")
		if s == len(stack)-1 {
			fmt.Print("\033[1m", stack[s].Debug(), "\033[0m")
		} else {
			fmt.Print(stack[s].Debug())
		}
	}
	fmt.Println(" ]")
}

func (vm *VirtualMachine) AddValueToStack(ir InstrRef) {
	vm.stack = append(vm.stack, NewWord(ir.value))
	logAddedToStack(vm.stack, ir.text)
}

// Currently just for traversing the scope outwardly
func (vm *VirtualMachine) SubmitLocalRequest(ir InstrRef) {
	idx, ok := locateScopeLabel(vm.scope, ir.text)
	if ok {
		// TODO: Invoke address, skip addrRef
		vm.stack = append(vm.stack, WordFromAddr(idx))
		logAddedToStack(vm.stack, ir.text)
		return
	}
	vm.heap = append(vm.heap, EmptyWord())
	idx = len(vm.heap) - 1
	vm.scope[len(vm.scope)-1][ir.text] = idx
	vm.stack = append(vm.stack, WordFromAddr(idx))
	logAddedToStack(vm.stack, ir.text)
}

// Used for traversing maps and legends
func (vm *VirtualMachine) SubmitInnerRequest(ir InstrRef) {

}

// Used for traversing submaps and legends
func (vm *VirtualMachine) SubmitUnderRequest(ir InstrRef) {
}

// Used for traversing primitives and compilations
func (vm *VirtualMachine) SubmitOuterRequest(ir InstrRef) {
	switch ir.text {
	case ".=", ":=":
		vm.assignLabel()
	case "++":
		vm.addTwoInts()
	case "--":
		vm.subTwoInts()
	case "**":
		vm.mulTwoInts()
	case "//":
		vm.divTwoInts()
	case "^^":
		vm.expTwoInts()
	case "[[":
		vm.beginMap()
	case "]]":
		vm.endMap()
	case "((":
		vm.beginRoutine()
	case "))":
		vm.endRoutine()

	default:
		panic("Not a valid outer operator")
	}
}

// Used for signalling across Rhumb
func (vm *VirtualMachine) SubmitEventRequest(ir InstrRef) {}

// Used for replying to signals across Rhumb
func (vm *VirtualMachine) SubmitReplyRequest(ir InstrRef) {}
