package object

import (
	"arena"
	"reflect"
)

type EmptyLegend struct {
	BaseLegend
	Fields []Field
}

func (l EmptyLegend) IsLegend() {}

func (l EmptyLegend) WHAT() string { return "Legend[EmptyMap]" }

func (x EmptyLegend) Equals(y Legend) bool {
	if yMap, ok := y.(EmptyLegend); ok {
		return &x == &yMap
	}
	return false
}

func (x EmptyLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(EmptyLegend); ok {
		if &x == &yMap {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewEmptyLegend(a *arena.Arena) *EmptyLegend {
	addr := arena.New[EmptyLegend](a)
	*addr = EmptyLegend{
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
