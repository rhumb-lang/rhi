package object

import (
	"arena"
	"fmt"
	"reflect"
)

type Map interface {
	Any
	IsMap()
	Get(Hashable) (Any, error)
	Set(Hashable, Any)
}

func mapGet(fields []Field, elems []Any, key Hashable) (Any, error) {
	for _, field := range fields {
		if field.Equals(key) {
			switch value := field.Value.(type) {
			case ArgumentSymbol:
				if value.IsNamed() {
					panic("field values should not be named arguments")
				} else {
					return elems[value.Value], nil
				}
			default:
				return value, nil
			}
		}
	}
	return nil, fmt.Errorf("no value found")
}
func mapSet(fields *[]Field, elems *[]Any, key Hashable, val Any) {
	for _, field := range *fields {
		if field.Equals(key) {
			switch value := field.Value.(type) {
			case ArgumentSymbol:
				if value.IsNamed() {
					panic("field values should not be named arguments")
				} else {
					(*elems)[value.Value] = val
					return
				}
			default:
				field.Value = val
				return
			}
		}
	}
	*fields = append(*fields, Field{
		Name:    key,
		Mutable: true,
		Value:   val,
	})
}

type ListMap struct {
	legend *Legend[[]Any]
	Values []Any
}

func NewListMap(a *arena.Arena) *ListMap {
	arrAddr := arena.New[[]Any](a)
	*arrAddr = []Any{}
	addr := arena.New[ListMap](a)
	*addr = ListMap{NewMapLegend(a), *arrAddr}
	return addr
}

func (l ListMap) IsObject()    {}
func (l ListMap) WHAT() string { return "ListMap" }
func (x ListMap) Equals(y Any) bool {
	if yMap, ok := y.(ListMap); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}
func (l ListMap) IsMap() {}
func (l ListMap) Length() int {
	return len(l.legend.Data)
}
func (l ListMap) At(index int) *Any {
	return &l.legend.Data[index]
}
func (l ListMap) Append(val Any) {
	l.legend.Data = append(l.legend.Data, val)
}
func (l ListMap) Get(key Hashable) (Any, error) {
	return mapGet(l.legend.Fields, l.Values, key)
}
func (l ListMap) Set(key Hashable, val Any) {
	mapSet(&l.legend.Fields, &l.Values, key, val)
}

type RoutineMap struct {
	legend *Legend[Block]
	Values []Any
}

func NewRoutineMap(a *arena.Arena, block Block) *RoutineMap {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[RoutineMap](a)
	*addr = RoutineMap{
		NewRoutineLegend(a, block),
		values,
	}
	return addr
}

func (r RoutineMap) IsObject()    {}
func (r RoutineMap) WHAT() string { return "RoutineMap" }
func (x RoutineMap) Equals(y Any) bool {
	if yMap, ok := y.(RoutineMap); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}
func (r RoutineMap) IsMap() {}
func (r RoutineMap) Get(key Hashable) (Any, error) {
	return mapGet(r.legend.Fields, r.Values, key)
}
func (r *RoutineMap) Set(key Hashable, val Any) {
	mapSet(&r.legend.Fields, &r.Values, key, val)
}
func (r *RoutineMap) AsFunction(params ListMap) *RoutineMap {
	for _, param := range params.legend.Data {
		if key, ok := param.(Hashable); ok {
			mapSet(&r.legend.Fields, &r.Values, key,
				ArgumentSymbol{Value: len(r.Values)})
			r.Values = append(r.Values, Empty{})
		} else {
			panic("not a hashable parameter")
		}
	}
	return r
}
func (r RoutineMap) Invoke(scopeIndex int) {
	for code := range r.legend.Data.Codes {
		fmt.Printf("code: %v\n", code)
	}
}

// type SelectorMap struct {
// 	legend *Legend[Chunk]
// 	Values []object.Any
// }

// func NewSelectorMap() *SelectorMap {
// 	values := arena.New[[]object.Any](vm.Memory)
// 	*values = []object.Any{}
// 	addr := arena.New[RoutineMap](vm.Memory)
// 	*addr = RoutineMap{vm.NewRoutineLegend(), *values}
// 	return addr
// }

// func (s SelectorMap) IsObject()    {}
// func (s SelectorMap) IsMap()       {}
// func (s SelectorMap) WHAT() string { return "SelectorMap" }

type TextMap struct {
	legend *Legend[string]
	Values []Any
}

func (t TextMap) IsObject()    {}
func (t TextMap) WHAT() string { return "TextMap" }
func (t TextMap) IsMap()       {}
func (t TextMap) Length() int {
	return len(t.legend.Data)
}
func (t TextMap) At(index int) byte {
	return t.legend.Data[index]
}
func (t TextMap) String(index int) string {
	return t.legend.Data
}
func (x TextMap) Equals(y Any) bool {
	if yMap, ok := y.(TextMap); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}
func (t TextMap) Get(key Hashable) (Any, error) {
	return mapGet(t.legend.Fields, t.Values, key)
}
func (t TextMap) Set(key Hashable, val Any) {
	mapSet(&t.legend.Fields, &t.Values, key, val)
}
