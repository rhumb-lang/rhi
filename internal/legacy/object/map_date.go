package object

import (
	"arena"
	"encoding/binary"
	"fmt"
	"time"
)

type Date struct{
	legend *DateLegend
	Values []Any
}

func (w Date) IsObject() {}

func (w Date) WHAT() string {
	return fmt.Sprint("Date: ", w.legend.Data)
}

func (x Date) Equals(y Any) bool {
	objY, ok := y.(Date)
	return ok && x.legend.Equals(objY.legend)
}

func (x Date) DeepEquals(y Any) bool {
	if yMap, ok := y.(Date); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}

func (key Date) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, key.legend.Data.Unix())
	return buf[:n]
}


func NewDate(a *arena.Arena, val time.Time) *Date {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Date](a)
	*addr = Date{
		legend: NewDateLegend(a, val),
		Values: values,
	}
	return addr
}
