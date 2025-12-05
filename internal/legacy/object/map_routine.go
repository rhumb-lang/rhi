package object

import (
	"arena"
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

func (o *Routine) initParameters(mem *arena.Arena) {
	paramsKey := NewLabel(mem, "$0")
	if _, err := o.Scope().Get(paramsKey); err == nil {
		panic("$0 was already initialized")
	} else {
		paramList := NewList(mem)
		o.Scope().Set(mem, paramsKey, paramList, false)
	}

}

func (o *Routine) setParam(mem *arena.Arena, param Any) *Routine {
	paramsKey := NewLabel(mem, "$0")
	if paramsValue, err := o.Scope().Get(paramsKey); err == nil {
		if paramsList, ok := paramsValue.(*List); ok {
			paramsList.Append(param)
		} else {
			panic("$O does not contain a parameter list")
		}
	} else {
		panic("$0 has not been initialized")
	}
	return o
}

func (o *Routine) SetParameters(mem *arena.Arena, params List) *Routine {
	o.initParameters(mem)
	for _, param := range params.legend.Data {
		if key, ok := param.(Hashable); ok {
			empty := NewEmptyMap(mem)
			o.Scope().Set(mem, key, empty, true)
			o.setParam(mem, key)
		} else {
			panic("not a hashable parameter")
		}
	}
	return o
}

func (o *Routine) SetArguments(mem *arena.Arena, args *List) *Routine {
	paramsKey := NewLabel(mem, "$0")
	if params, err := o.Scope().Get(paramsKey); err == nil {
		if paramsList, ok := params.(*List); ok {
			for i, arg := range args.legend.Data {
				if paramLabel, ok := paramsList.GetAt(i).(Hashable); ok {
					o.Scope().Set(mem, paramLabel, arg, true)
				}
				// paramsList.SetAt(i, arg)
			}
		} else {
			panic("$0 is not a ListMap")
		}
	} else {
		panic("$0 is not found")
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

func (o *Routine) Disassemble(frame int) {
	o.legend.Data.Disassemble(frame)
}

// func (o *Routine) Push(n Any) {
// 	o.legend.Data.TopStack().Push(n)
// }

// func (o *Routine) Pop() Any {
// 	return o.legend.Data.TopStack().Pop()
// }

// func (o Routine) Invoke(args List) {
// 	// TODO: assign each argument to their corresponding parameter
// 	o.legend.Data.
// 	o.legend.Data.Run()
// }

func NewRoutine(a *arena.Arena, f *stack.Stack[*Routine], p *Scope) *Routine {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Routine](a)
	*addr = Routine{
		legend: NewRoutineLegend(a, f, p),
		Values: values,
	}
	return addr
}
