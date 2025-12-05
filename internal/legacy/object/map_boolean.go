package object

import (
	"arena"
	"fmt"
)

var (
	TRUE  *Boolean
	FALSE *Boolean
)

func init() {
	a := arena.NewArena()
	TRUE = NewBooleanUnsafe(a, true)
	FALSE = NewBooleanUnsafe(a, false)
}

type Boolean struct {
	legend *BooleanLegend
	Values []Any
}

func (t Boolean) IsObject() {}

func (t Boolean) WHAT() string {
	return fmt.Sprint("BooleanNumber: ", t.legend.Data)
}

func (x Boolean) Equals(y Any) bool {
	if yMap, ok := y.(Boolean); ok {
		return x.legend.Equals(yMap.legend)
	}
	return false
}

func (x Boolean) DeepEquals(y Any) bool {
	if yMap, ok := y.(Boolean); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}

func NewBoolean(a *arena.Arena, val bool) *Boolean {
	if val {
		return TRUE
	}
	return FALSE
}

func NewBooleanUnsafe(a *arena.Arena, val bool) *Boolean {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Boolean](a)
	*addr = Boolean{
		legend: NewBooleanLegend(a, val),
		Values: values,
	}
	return addr
}