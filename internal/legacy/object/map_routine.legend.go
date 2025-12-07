package object

import (
	"arena"
	"reflect"

	"git.sr.ht/~madcapjake/rhi/internal/stack"
)

type RoutineLegend struct {
	BaseLegend
	Data   *Block
	Fields []Field
}

func (l RoutineLegend) IsLegend() {}

func (l RoutineLegend) WHAT() string { return "Legend[RoutineMap]" }

func (x RoutineLegend) Equals(y Legend) bool {
	if yMap, ok := y.(RoutineLegend); ok {
		if x.Data == yMap.Data {
			return true
		}
	}
	return false
}

func (x RoutineLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(RoutineLegend); ok {
		if x.Data == yMap.Data {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewRoutineLegend(a *arena.Arena, f *stack.Stack[*Routine], p *Scope) *RoutineLegend {
	addr := arena.New[RoutineLegend](a)
	*addr = RoutineLegend{
		Data:   NewBlock(a, f, p),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
