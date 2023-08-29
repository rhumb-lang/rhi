package vm

import (
	"arena"
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/code"
	obj "git.sr.ht/~madcapjake/grhumb/internal/object"
)

const MEM_SIZE = 5000
const STK_SIZE = 500
const SCP_SIZE = 50

type VirtualMachine struct {
	Memory *arena.Arena
	Blocks []obj.Block
	Stacks []Stack
	Scopes []Scope
	Maps   []*obj.Map
}

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)
	vm.Memory = arena.NewArena()
	vm.Blocks = arena.MakeSlice[obj.Block](vm.Memory, 0, 5)
	vm.Blocks = append(vm.Blocks, obj.NewBlock(vm.Memory))
	vm.Stacks = arena.MakeSlice[Stack](vm.Memory, 0, 5)
	vm.Stacks = append(vm.Stacks, NewStack(vm.Memory))
	vm.Scopes = arena.MakeSlice[Scope](vm.Memory, 0, 5)
	vm.Scopes = append(vm.Scopes, NewScope(vm.Memory))
	vm.Maps = arena.MakeSlice[*obj.Map](vm.Memory, 0, 5)
	return vm
}

func (vm *VirtualMachine) TopBlock() *obj.Block {
	return &vm.Blocks[len(vm.Blocks)-1]
}

func (vm *VirtualMachine) TopStack() *Stack {
	return &vm.Stacks[len(vm.Stacks)-1]
}

func (vm *VirtualMachine) TopScope() *Scope {
	return &vm.Scopes[len(vm.Scopes)-1]
}

func (vm *VirtualMachine) TopMap() *obj.Map {
	length := len(vm.Maps)
	if length > 0 {
		return vm.Maps[length-1]
	} else {
		return nil
	}
}

func (vm *VirtualMachine) PushCode(c obj.Code) {
	vm.TopBlock().Write(c)
}

func (vm *VirtualMachine) RegisterObject(o obj.Any) int {
	return vm.TopBlock().Register(o)
}

func (vm *VirtualMachine) PushObject(o obj.Any) {
	vm.TopStack().Push(o)
}

func (vm *VirtualMachine) PopObject() obj.Any {
	return vm.TopStack().Pop()
}

func (vm *VirtualMachine) PushBlock() int {
	vm.Blocks = append(vm.Blocks, obj.NewBlock(vm.Memory))
	return len(vm.Blocks) - 1
}

func (vm *VirtualMachine) PopBlock() obj.Block {
	last := len(vm.Blocks) - 1
	block := vm.Blocks[last]
	vm.Blocks = vm.Blocks[:last]
	return block
}

func (vm *VirtualMachine) PushStack() int {
	vm.Stacks = append(vm.Stacks, NewStack(vm.Memory))
	return len(vm.Stacks) - 1
}

func (vm *VirtualMachine) PopStack() int {
	last := len(vm.Stacks) - 1
	vm.Stacks = vm.Stacks[:last]
	return last
}

func (vm *VirtualMachine) PushScope() int {
	vm.Scopes = append(vm.Scopes, NewScope(vm.Memory))
	return len(vm.Scopes) - 1
}

func (vm *VirtualMachine) PopScope() int {
	last := len(vm.Scopes) - 1
	vm.Scopes = vm.Scopes[:last]
	return last
}

func (vm *VirtualMachine) PushMap(m obj.Map) int {
	vm.Maps = append(vm.Maps, &m)
	return len(vm.Maps) - 1
}

func (vm *VirtualMachine) PopMap() *obj.Map {
	last := len(vm.Maps) - 1
	result := vm.Maps[last]
	vm.Maps = vm.Maps[:last]
	return result
}

func (vm *VirtualMachine) Run(block *obj.Block) (obj.Any, error) {
	vm.PushStack()
	vm.PushScope()
	for _, code := range block.Codes {
		val := block.Values[code.GetID()]
		vm.Execute(code, val)
	}
	result := vm.PopObject()
	vm.PopScope()
	vm.PopStack()
	return result, nil
}

func (vm *VirtualMachine) Disassemble(block *obj.Block) (obj.Any, error) {
	vm.PushStack()
	vm.PushScope()
	for _, code := range block.Codes {
		val := block.Values[code.GetID()]
		fmt.Println("Code =", code.WHAT(), "\tValue =", val.WHAT())
	}
	result := vm.PopObject()
	vm.PopScope()
	vm.PopStack()
	return result, nil
}

func (vm *VirtualMachine) RunMain() (obj.Any, error) {
	return vm.Run(vm.TopBlock())
}

func (vm *VirtualMachine) DisassembleMain() (obj.Any, error) {
	return vm.Disassemble(vm.TopBlock())
}

func (vm *VirtualMachine) Execute(c code.Any, v obj.Any) {
	switch c.(type) {
	case code.Value:
		vm.ExecuteValue(v)
	case code.Local:
		vm.ExecuteLocal(v)
	case code.Inner:
		vm.ExecuteInner(v)
	case code.Under:
		vm.ExecuteUnder(v)
	case code.Outer:
		vm.ExecuteOuter(v)
	case code.Event:
		vm.ExecuteEvent(v)
	case code.Reply:
		vm.ExecuteReply(v)
	}
}

func (vm *VirtualMachine) ExecuteValue(value obj.Any) {
	if topMap := vm.TopMap(); topMap != nil {
		switch mapVal := (*topMap).(type) {
		case obj.ListMap:
			mapVal.Append(value)
		default:
			panic("not yet implemented")
		}
	} else {
		vm.TopStack().Push(value)
	}
}
func (vm *VirtualMachine) ExecuteLocal(value obj.Any) {
	var stack *Stack = vm.TopStack()
	switch value := value.(type) {
	case obj.LabelSymbol:
		vm.TopScope().Get(value)
	case obj.OperatorSymbol:
		switch value.Value {
		case "[":
			vm.PushMap(*obj.NewListMap(vm.Memory))
		case "]":
			vm.TopStack().Push(*vm.PopMap())
		case "->":
			block, blockOk := stack.Pop().(obj.Block)
			pList, paramOk := stack.Pop().(obj.ListMap)
			if blockOk && paramOk {
				routine := obj.
					NewRoutineMap(vm.Memory, block).
					AsFunction(pList)
				stack.Push(routine)
			}
		case ".=":
			assignee := stack.Pop()
			if label, labelOk := stack.Pop().(obj.LabelSymbol); labelOk {
				vm.TopScope().Set(label, assignee)
			} else {
				panic("cannot assign to a non-label value")
			}
		default:
			panic("not yet implemented")
		}
	default:
		panic("not yet implemented")
	}
}
func (vm *VirtualMachine) ExecuteInner(value obj.Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
func (vm *VirtualMachine) ExecuteUnder(value obj.Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
func (vm *VirtualMachine) ExecuteOuter(value obj.Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
func (vm *VirtualMachine) ExecuteEvent(value obj.Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
func (vm *VirtualMachine) ExecuteReply(value obj.Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
