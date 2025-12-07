package object

import "arena"

type Key struct {
	legend *KeyLegend
	Values []Any
}

func NewKey(a *arena.Arena, value string) *Key {
	addr := arena.New[Key](a)
	values := arena.MakeSlice[Any](a, 0, 2)
	*addr = Key{
		legend: NewKeyLegend(a, value),
		Values: values,
	}
	return addr
}

func (o Key) IsObject() {}

func (o Key) WHAT() string { return "KeyMap" }

func (x Key) Equals(y Any) bool {
	if yMap, ok := y.(Key); ok {
		return x.legend.Equals(yMap.legend)
	}
	return false
}

func (x Key) DeepEquals(y Any) bool {
	if yMap, ok := y.(Key); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}

func (o Key) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *Key) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}
