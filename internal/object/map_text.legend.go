package object

import (
	"arena"
	"reflect"
)

type TextLegend struct {
	BaseLegend
	Data   []rune
	Fields []Field
}

func (l TextLegend) IsLegend() {}

func (l TextLegend) WHAT() string { return "Legend[RoutineMap]" }

func (x TextLegend) Equals(y Legend) bool {
	if yMap, ok := y.(TextLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x TextLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(TextLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewTextLegend(a *arena.Arena, data string) *TextLegend {
	addr := arena.New[TextLegend](a)
	*addr = TextLegend{
		Data:   arena.MakeSlice[rune](a, len(data), len(data)),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	copy(addr.Data, []rune(data))
	return addr
}
