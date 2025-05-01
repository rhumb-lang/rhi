package object

import (
	"arena"
	"reflect"
)

type ListLegend struct {
	BaseLegend
	Data   []Any
	Fields []Field
}

func (r ListLegend) IsLegend() {}

func (r ListLegend) WHAT() string { return "Legend[ListMap]" }

func (x ListLegend) Equals(y Legend) bool {
	if _, ok := y.(ListLegend); ok {
		return true
	}
	return false
}

func (x ListLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(ListLegend); ok {
		return reflect.DeepEqual(x.Fields, yMap.Fields)
	}
	return false
}

func NewListLegend(a *arena.Arena) *ListLegend {
	addr := arena.New[ListLegend](a)
	*addr = ListLegend{
		Data:   arena.MakeSlice[Any](a, 0, 2),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
