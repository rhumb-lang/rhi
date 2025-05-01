package object

import (
	"arena"
	"reflect"
)

type BooleanLegend struct {
	BaseLegend
	Data   bool
	Fields []Field
}

func (t BooleanLegend) IsLegend() {}

func (t BooleanLegend) WHAT() string { return "Literal[Label]" }

func (x BooleanLegend) Equals(y Legend) bool {
	if yMap, ok := y.(BooleanLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x BooleanLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(BooleanLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewBooleanLegend(a *arena.Arena, data bool) *BooleanLegend {
	addr := arena.New[BooleanLegend](a)
	*addr = BooleanLegend{
		Data: data,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}