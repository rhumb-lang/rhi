package object

import (
	"arena"
	"encoding/binary"
	"fmt"
	"math"
)

type Ratio struct {
	legend *RatioLegend
	Values []Any
}

func NewRatio(a *arena.Arena, val float64) *Ratio {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Ratio](a)
	*addr = Ratio{
		legend: NewRatioLegend(a, val),
		Values: values,
	}
	return addr
}

func (o Ratio) IsObject() {}
func (o Ratio) WHAT() string {
	return fmt.Sprint("Ratio: ", o.legend.Data)
}
func (x Ratio) Equals(y Any) bool {
	objY, ok := y.(Date)
	return ok && x.legend.Equals(objY.legend)
}
func (key Ratio) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, math.Float64bits(key.legend.Data))
	return buf[:n]
}

func (o Ratio) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *Ratio) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}
