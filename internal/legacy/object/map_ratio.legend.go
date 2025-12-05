package object

import (
	"arena"
	"reflect"
)

type RatioLegend struct {
	BaseLegend
	Data   float64
	Fields []Field
}

func (l RatioLegend) IsLegend() {}

func (l RatioLegend) WHAT() string { return "Legend[RatioMap]" }

func (x RatioLegend) Equals(y Legend) bool {
	if yMap, ok := y.(RatioLegend); ok {
		return x.Data == yMap.Data
	}
	return false
}

func (x RatioLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(RatioLegend); ok {
		if x.Data == yMap.Data {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewRatioLegend(a *arena.Arena, data float64) *RatioLegend {
	addr := arena.New[RatioLegend](a)
	*addr = RatioLegend{
		Data:   data,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}