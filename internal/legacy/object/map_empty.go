package object

import "arena"

type Empty struct {
	legend *EmptyLegend
	Values []Any
}

func (e Empty) IsObject() {}
func (e Empty) WHAT() string {
	return "Empty"
}
func (x Empty) Equals(y Any) bool {
	_, ok := y.(Empty)
	return ok
}

func NewEmptyMap(a *arena.Arena) *Empty {
	addr := arena.New[Empty](a)
	values := arena.MakeSlice[Any](a, 0, 2)
	*addr = Empty{
		legend: NewEmptyLegend(a),
		Values: values,
	}
	return addr
}
