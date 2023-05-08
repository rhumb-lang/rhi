package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/code"
	"git.sr.ht/~madcapjake/grhumb/internal/object"
)

type Legend[T any] struct {
	Meta     object.Ref[Legend[any]]
	Link     object.Ref[Map]
	Needs    object.Ref[Map]
	Provides object.Ref[Map]
	Data     T
	Fields   []object.Any
}

func (l Legend[any]) IsObject()    {}
func (l Legend[any]) WHAT() string { return "Legend" }

func NewRoutineLegend() Legend[[]code.Any] {
	return Legend[[]code.Any]{
		Data: []code.Any{},
	}
}
