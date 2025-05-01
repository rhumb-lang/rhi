package object

import (
	"arena"
	"reflect"

	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

type SelectorMapLegend struct {
	BaseLegend
	Data   *Block
	Fields []Field
}

func (s SelectorMapLegend) IsLegend() {}

func (s SelectorMapLegend) WHAT() string { return "Legend[SelectorMap]" }

func (x SelectorMapLegend) Equals(y Legend) bool {
	if yMap, ok := y.(SelectorMapLegend); ok {
		if x.Data == yMap.Data {
			return true
		}
	}
	return false
}

func (x SelectorMapLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(SelectorMapLegend); ok {
		if x.Data == yMap.Data {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewSelectorLegend(a *arena.Arena, fs *stack.Stack[*Routine], p *Scope) *SelectorMapLegend {
	addr := arena.New[SelectorMapLegend](a)
	*addr = SelectorMapLegend{
		Data:   NewBlock(a, fs, p),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
