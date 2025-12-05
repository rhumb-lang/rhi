package object

import (
	"arena"
	"reflect"
)

type List struct {
	legend *ListLegend
	Values []Any
}

func NewList(a *arena.Arena) *List {
	arrAddr := arena.New[[]Any](a)
	*arrAddr = []Any{}
	addr := arena.New[List](a)
	*addr = List{NewListLegend(a), *arrAddr}
	return addr
}

func (o List) IsObject() {}

func (o List) WHAT() string { return "ListMap" }

func (x List) Equals(y Any) bool {
	if yMap, ok := y.(List); ok {
		if x.legend == yMap.legend {
			return reflect.DeepEqual(x.Values, yMap.Values)
		}
	}
	return false
}

func (o List) Length() int {
	return len(o.legend.Data)
}

func (o List) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o List) GetAt(index int) Any {
	return o.legend.Data[index]
}

func (o *List) Append(val Any) {
	o.legend.Data = append(o.legend.Data, val)
}

func (o *List) SetAt(index int, val Any) {
	o.legend.Data[index] = val
}

func (o *List) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}
