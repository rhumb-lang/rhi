package object

import (
	"arena"
	"reflect"
)

type LabelLiteralLegend struct {
	BaseLegend
	Data   []rune
	Fields []Field
}

func (t LabelLiteralLegend) IsLegend() {}

func (t LabelLiteralLegend) WHAT() string { return "Literal[Label]" }

func (x LabelLiteralLegend) Equals(y Legend) bool {
	if yMap, ok := y.(LabelLiteralLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x LabelLiteralLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(LabelLiteralLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewLabelLiteralLegend(a *arena.Arena, data string) *LabelLiteralLegend {
	addr := arena.New[LabelLiteralLegend](a)
	*addr = LabelLiteralLegend{
		Data: arena.MakeSlice[rune](a, len(data), len(data)),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	copy(addr.Data, []rune(data))
	return addr
}