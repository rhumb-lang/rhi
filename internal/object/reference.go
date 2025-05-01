package object

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type Reference[O Any] struct {
	Address int64
}

func (o Reference[Any]) IsObject() {}
func (o Reference[O]) WHAT() string {
	var rt [0]O
	return fmt.Sprint("Reference:", reflect.TypeOf(rt).Elem())
}
func (x Reference[O]) Equals(y Any) bool {
	objY, ok := y.(Reference[Any])
	return ok && x.Address == objY.Address
}
func (key Reference[Any]) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, key.Address)
	return buf[:n]
}
