package object

import (
	"arena"
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/color"
	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

type Block struct {
	memory *arena.Arena

	// bytecode
	Values []Any
	Codes  []code.Any

	// runtime
	Frames *stack.Stack[*Routine]
	Scope  *Scope
	Stacks *stack.Stack[*stack.Stack[Any]]
	Lists  *stack.Stack[*List]
}

func NewBlock(a *arena.Arena, fs *stack.Stack[*Routine], p *Scope) *Block {
	addr := arena.New[Block](a)
	*addr = Block{
		memory: a,
		Values: arena.MakeSlice[Any](a, 0, 5),
		Codes:  arena.MakeSlice[code.Any](a, 0, 50),
		Frames: fs,
		Scope:  NewScope(a, p),
		Stacks: stack.NewStack[*stack.Stack[Any]](a),
		Lists:  stack.NewStack[*List](a),
	}
	addr.Stacks.Push(stack.NewStack[Any](a))
	return addr
}

// func (block Block) IsObject() {}

// func (block Block) WHAT() string { return "Block: ..." }

// func (x Block) Equals(y Any) bool {
// 	switch yObj := y.(type) {
// 	case Block:
// 		if len(x.Codes) != len(yObj.Codes) {
// 			return false
// 		} else {
// 			for i, c := range x.Codes {
// 				if c != yObj.Codes[i] {
// 					return false
// 				}
// 			}
// 			return true
// 		}
// 	default:
// 		return false
// 	}
// }

func (b *Block) Register(o Any) int {
	for i, existing := range b.Values {
		if o.Equals(existing) {
			return i
		}
	}
	b.Values = append(b.Values, o)
	return len(b.Values) - 1
}

func (b *Block) Write(c ...code.Any) {
	b.Codes = append(b.Codes, c...)
}

func (b Block) TopList() (list *List) {
	list = *b.Lists.Top()
	return
}

func (b *Block) PushList(list *List) {
	b.Lists.Push(list)
}

func (b *Block) PopList() *List {
	return (b.Lists.Pop())
}

func (b *Block) TopStack() (stack *stack.Stack[Any]) {
	stack = *b.Stacks.Top()
	return
}

func (b *Block) PushStack(stack *stack.Stack[Any]) {
	b.Stacks.Push(stack)
}

func (b *Block) PopStack() *stack.Stack[Any] {
	return (b.Stacks.Pop())
}

func (b *Block) Run() (Any, error) {
	for _, code := range b.Codes {
		val := b.Values[code.GetID()]
		b.Execute(code, val)
	}
	result := b.TopStack().Pop()
	return result, nil
}

func scopeTabs(frameLevel int) string {
	tabs := []rune{}
	for range frameLevel {
		tabs = append(tabs, '\t')
	}
	return string(tabs)
}

func printDisassembly(frame, i int, code code.Any, val Any) {
	disText := fmt.Sprint(
		scopeTabs(frame),
		color.Gray, i, color.Reset, " ",
		color.Purple, "Code", color.Reset, " = ", code.WHAT(), " ",
		color.Blue, "Value", color.Reset, " = ", val.WHAT(),
	)
	fmt.Println(disText)
}

func (b Block) Disassemble(frame int) {
	for i, code := range b.Codes {
		val := b.Values[code.GetID()]
		if routineVal, ok := val.(*Routine); ok {
			printDisassembly(frame, i, code, routineVal)
			routineVal.Disassemble(frame + 1)
		} else {
			printDisassembly(frame, i, code, val)
		}
	}
}

func (b *Block) Execute(c code.Any, o Any) {
	switch c.(type) {
	case code.Value:
		b.ExecuteValue(o)
	case code.Local:
		b.ExecuteLocal(o)
	case code.Inner:
		b.ExecuteInner(o)
	case code.Under:
		b.ExecuteUnder(o)
	case code.Outer:
		b.ExecuteOuter(o)
	case code.Event:
		b.ExecuteEvent(o)
	case code.Reply:
		b.ExecuteReply(o)
	}
}

func (b *Block) ExecuteValue(obj Any) {
	b.TopStack().Push(obj)
}

func (b *Block) ExecuteLocal(value Any) {
	switch obj := value.(type) {
	case *Label:
		switch string(obj.Value) {
		case "_[[_":
			b.PushList(NewList(b.memory))
		case "_[>_":
			b.TopList().Append(b.TopStack().Pop())
		case "_]]_":
			b.TopStack().Push(b.PopList())
		case "_((_":
			b.PushStack(stack.NewStack[Any](b.memory))
		case "_))_":
			b.TopStack().Push(b.PopStack().Pop())
		case "_<[_":
			b.Frames.Push(NewRoutine(b.memory, b.Frames, b.Scope))
		case "_]>_":
			b.TopStack().Push(b.Frames.Pop())
		case "_->_":
			routine, routineOk := b.TopStack().Pop().(*Routine)
			if !routineOk {
				panic("top of stack is not a routine map")
			}
			list, listOk := b.TopStack().Pop().(*List)
			if !listOk {
				panic("top of stack is not a list map")
			}
			routine.SetParameters(b.memory, *list)
			b.TopStack().Push(routine)
		case "_.=_":
			assignee := b.TopStack().Pop()
			if label, labelOk := b.TopStack().Pop().(*Label); labelOk {
				b.Scope.Set(b.memory, label, assignee, false)
			} else {
				panic("cannot assign to a non-label value")
			}
		case "_:=_":
			assignee := b.TopStack().Pop()
			if label, labelOk := b.TopStack().Pop().(*Label); labelOk {
				b.Scope.Set(b.memory, label, assignee, true)
			} else {
				panic("cannot assign to a non-label value")
			}
		case "_**_":
			left := b.TopStack().Pop()
			right := b.TopStack().Pop()
			if leftNum, ok := left.(*Number); ok {
				result := NewNumber(b.memory, leftNum.Multiply(right))
				b.TopStack().Push(result)
			} else {
				panic("Left operand is not a number")
			}
		default:
			if val, err := b.Scope.Get(obj); err == nil {
				b.TopStack().Push(val)
			} else {
				panic("label not found")
			}
		}
	case *Reference[*Label]:
		b.TopStack().Push(obj)
	default:
		panic(fmt.Sprintf("Object of type '%s' not yet implemented", obj.WHAT()))
	}
}

func (b *Block) ExecuteInner(value Any) {
	switch obj := value.(type) {
	case *Label:
		switch string(obj.Value) {
		case "_[[_":
			b.PushList(NewList(b.memory))
		case "_[>_":
			b.TopList().Append(b.TopStack().Pop())
		case "_]]_":
			argList := b.PopList()
			switch routineOrReference := b.TopStack().Pop().(type) {
			case *Routine:
				b.Frames.Push(routineOrReference)
				routineOrReference.SetArguments(b.memory, argList)
				if result, err := routineOrReference.Run(); err == nil {
					b.TopStack().Push(result)
				} else {
					panic(err)
				}
				b.Frames.Pop()
			case *Reference[*Label]:
				labelRefObj, labelRefObjError := b.Scope.Get(routineOrReference.Address)
				if labelRefObjError != nil {
					panic("No referenced value found in scope")
				}
				if refRoutine, refRoutineOk := labelRefObj.(*Routine); refRoutineOk {
					b.Frames.Push(refRoutine)
					refRoutine.SetArguments(b.memory, argList)
					if result, err := refRoutine.Run(); err == nil {
						b.TopStack().Push(result)
					} else {
						panic(err)
					}
					b.Frames.Pop()
				} else {
					panic("The scoped value referenced is not a routine")
				}
			default:
				panic("No routine or reference found in stack to invoke")
			}

		// TODO: convert to field assigns
		// case ".=":
		// 	assignee := b.TopStack().Pop()
		// 	if label, labelOk := b.TopStack().Pop().(Label); labelOk {
		// 		b.Scope.Set(b.memory, label, assignee, false)
		// 	} else {
		// 		panic("cannot assign to a non-label value")
		// 	}
		// case ":=":
		// 	assignee := b.TopStack().Pop()
		// 	if label, labelOk := b.TopStack().Pop().(Label); labelOk {
		// 		b.Scope.Set(b.memory, label, assignee, true)
		// 	} else {
		// 		panic("cannot assign to a non-label value")
		// 	}
		default:
			b.Scope.Get(obj)
		}
	case *Reference[*Label]:
		b.TopStack().Push(obj)
	default:
		panic(fmt.Sprintf("Object of type '%s' not yet implemented", obj.WHAT()))
	}
}

func (b *Block) ExecuteUnder(value Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}

func (b *Block) ExecuteOuter(value Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}

func (b *Block) ExecuteEvent(value Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}

func (b *Block) ExecuteReply(value Any) {
	switch value.(type) {
	default:
		panic("not yet implemented")
	}
}
