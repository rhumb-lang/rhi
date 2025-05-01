package object

import (
	"arena"
	"fmt"
	"reflect"

	"git.sr.ht/~madcapjake/rhi/internal/code"
	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

type Routine struct {
	legend *RoutineLegend
	Values []Any
}

func (o Routine) IsObject() {}

func (o Routine) WHAT() string { return "Routine" }

func (x Routine) Equals(y Any) bool {
	if yMap, ok := y.(Routine); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}

func (o Routine) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *Routine) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}

func (o *Routine) SetParameters(mem *arena.Arena, params List) *Routine {
	for _, param := range params.legend.Data {
		if key, ok := param.(Hashable); ok {
			empty := NewEmptyMap(mem)
			fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, empty)
		} else {
			panic("not a hashable parameter")
		}
	}
	return o
}

func (o *Routine) SetArguments(mem *arena.Arena, args List) *Routine {
	for _, arg := range args.legend.Data {
		if key, ok := arg.(Hashable); ok {
			empty := NewEmptyMap(mem)
			fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, empty)
		} else {
			panic("not a hashable parameter")
		}
	}
	return o
}

func (o Routine) Scope() *Scope {
	return o.legend.Data.Scope
}

func (o *Routine) Register(n Any) int {
	return o.legend.Data.Register(n)
}

func (o *Routine) Write(c ...code.Any) {
	o.legend.Data.Write(c...)
}

func (o *Routine) Run() (Any, error) {
	return o.legend.Data.Run()
}

func (o *Routine) Disassemble() {
	o.legend.Data.Disassemble()
}

// func (o *Routine) Push(n Any) {
// 	o.legend.Data.TopStack().Push(n)
// }

// func (o *Routine) Pop() Any {
// 	return o.legend.Data.TopStack().Pop()
// }

func (o Routine) Invoke(args List) {

	for code := range o.legend.Data.Codes {
		fmt.Printf("code: %v\n", code)
	}
}

func NewRoutine(a *arena.Arena, f *stack.Stack[*Routine], p *Scope) *Routine {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Routine](a)
	*addr = Routine{
		legend: NewRoutineLegend(a, f, p),
		Values: values,
	}
	return addr
}
