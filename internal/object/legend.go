package object

import (
	"arena"
)

type Legend[T any] struct {
	Meta     *Legend[Legend[any]]
	Link     *Map
	Demands  *Map
	Supplies *Map
	Data     T
	Fields   []Field
}

func (l Legend[any]) IsObject()    {}
func (l Legend[any]) WHAT() string { return "Legend" }

func NewMapLegend(a *arena.Arena) *Legend[[]Any] {
	addr := arena.New[Legend[[]Any]](a)
	*addr = Legend[[]Any]{
		Data:   arena.MakeSlice[Any](a, 0, 2),
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}

func NewRoutineLegend(a *arena.Arena, b Block) *Legend[Block] {
	addr := arena.New[Legend[Block]](a)
	*addr = Legend[Block]{
		Data:   b,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}

func NewSelectorLegend(a *arena.Arena) *Legend[Block] {
	addr := arena.New[Legend[Block]](a)
	*addr = Legend[Block]{
		Data: NewBlock(a),
	}
	return addr
}

func NewTextLegend(a *arena.Arena) *Legend[string] {
	addr := arena.New[Legend[string]](a)
	*addr = Legend[string]{}
	return addr
}
