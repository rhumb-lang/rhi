package object

import (
	"arena"
	"encoding/binary"
	"fmt"
)

type Number struct {
	legend *NumberLegend
	Values []Any
}

func NewNumber(a *arena.Arena, val int64) *Number {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Number](a)
	*addr = Number{
		legend: NewNumberLegend(a, val),
		Values: values,
	}
	return addr
}

func (w Number) IsObject() {}
func (w Number) WHAT() string {
	return fmt.Sprint("Number: ", w.legend.Data)
}
func (x Number) Equals(y Any) bool {
	objY, ok := y.(Date)
	return ok && x.legend.Equals(objY.legend)
}
func (key Number) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, key.legend.Data)
	return buf[:n]
}

func (o Number) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *Number) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}
