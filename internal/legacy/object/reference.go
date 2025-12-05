package object

import (
	"arena"
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/color"
)

type Reference[O Any] struct {
	Address O
}

func NewReference[O Any](a *arena.Arena, obj O) *Reference[O] {
	addr := arena.New[Reference[O]](a)
	*addr = Reference[O]{
		Address: obj,
	}
	return addr
}

func (o Reference[O]) IsObject() {}
func (o Reference[O]) WHAT() string {
	// var rt [0]O
	// return fmt.Sprint("Reference: ", reflect.TypeOf(rt).Elem())
	return fmt.Sprint(
		"Reference",
		color.Red, "[", color.Reset,
		o.Address.WHAT(),
		color.Red, "]", color.Reset,
	)
}
func (x Reference[O]) Equals(y Any) bool {
	objY, ok := y.(Reference[O])
	return ok && x.Address.Equals(objY.Address)
}
