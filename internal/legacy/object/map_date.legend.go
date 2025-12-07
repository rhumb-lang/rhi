package object

import (
	"arena"
	"reflect"
	"time"
)

type DateLegend struct {
	BaseLegend
	Data   time.Time
	Fields []Field
}

func (l DateLegend) IsLegend() {}

func (l DateLegend) WHAT() string { return "Legend[DateMap]" }

func (x DateLegend) Equals(y Legend) bool {
	if yMap, ok := y.(DateLegend); ok {
		return x.Data == yMap.Data
	}
	return false
}

func (x DateLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(DateLegend); ok {
		if x.Data == yMap.Data {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewDateLegend(a *arena.Arena, data time.Time) *DateLegend {
	addr := arena.New[DateLegend](a)
	*addr = DateLegend{
		Data:   data,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}