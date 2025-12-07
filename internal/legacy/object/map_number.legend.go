package object

import (
	"arena"
	"reflect"
)

type NumberLegend struct {
	BaseLegend
	Data   int64
	Fields []Field
}

func (l NumberLegend) IsLegend() {}

func (l NumberLegend) WHAT() string { return "Legend[NumberMap]" }

func (x NumberLegend) Equals(y Legend) bool {
	if yMap, ok := y.(NumberLegend); ok {
		return x.Data == yMap.Data
	}
	return false
}

func (x NumberLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(NumberLegend); ok {
		if x.Data == yMap.Data {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewNumberLegend(a *arena.Arena, data int64) *NumberLegend {
	addr := arena.New[NumberLegend](a)
	*addr = NumberLegend{
		Data:   data,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}