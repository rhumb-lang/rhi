package object

import (
	"arena"
	"reflect"
)

type ScopeLegend struct {
	BaseLegend
	Data   *Scope
	Fields []Field
}

func (l ScopeLegend) IsLegend() {}

func (l ScopeLegend) WHAT() string { return "Legend[RoutineMap]" }

func (x ScopeLegend) Equals(y Legend) bool {
	if yMap, ok := y.(ScopeLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x ScopeLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(ScopeLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewScopeLegend(a *arena.Arena, parent *Scope) *ScopeLegend {
	addr := arena.New[ScopeLegend](a)
	*addr = ScopeLegend{
		Data:   parent,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
