package object

import (
	"arena"
	"fmt"
	"reflect"

	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

type SelectorMap struct {
	legend *SelectorMapLegend
	Values []Any
}

func (o SelectorMap) IsObject() {}

func (o SelectorMap) WHAT() string { return "SelectorMap" }

func (x SelectorMap) Equals(y Any) bool {
	if yMap, ok := y.(SelectorMap); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}

func (o SelectorMap) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *SelectorMap) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}

// func (r *SelectorMap) AsFunction(params ListMap) *SelectorMap {
// 	for _, param := range params.legend.Data {
// 		if key, ok := param.(Hashable); ok {
// 			fieldsSet(&r.legend.Fields, &r.Values, key,
// 				ParameterLiteral{Value: len(r.Values)})
// 			r.Values = append(r.Values, EmptyMapLiteral{})
// 		} else {
// 			panic("not a hashable parameter")
// 		}
// 	}
// 	return r
// }

func (r SelectorMap) Invoke(scopeIndex int) {
	for code := range r.legend.Data.Codes {
		fmt.Printf("code: %v\n", code)
	}
}

func NewSelectorMap(a *arena.Arena, fs *stack.Stack[*Routine], p *Scope) *SelectorMap {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[SelectorMap](a)
	*addr = SelectorMap{
		NewSelectorLegend(a, fs, p),
		values,
	}
	return addr
}
