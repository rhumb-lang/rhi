package object

import "arena"

type Text struct {
	legend *TextLegend
	Values []Any
}

func NewText(a *arena.Arena, value string) *Text {
	addr := arena.New[Text](a)
	values := arena.MakeSlice[Any](a, 0, 2)
	*addr = Text{
		legend: NewTextLegend(a, value),
		Values: values,
	}
	return addr
}

func (o Text) IsObject() {}

func (o Text) WHAT() string { return "TextMap" }

func (o Text) Length() int {
	return len(o.legend.Data)
}

func (o Text) At(index int) rune {
	return o.legend.Data[index]
}

func (o Text) String(index int) string {
	return string(o.legend.Data)
}

func (o Text) HashBytes() []byte {
	return []byte(string(o.legend.Data))
}

func (o Text) Get(key Hashable) (Any, error) {
	return fieldsGet(o.legend.Fields, o.Values, key)
}

func (o *Text) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}

func (x Text) Equals(y Any) bool {
	if yMap, ok := y.(Text); ok {
		return x.legend.Equals(yMap.legend)
	}
	return false
}

func (x Text) DeepEquals(y Any) bool {
	if yMap, ok := y.(Text); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}
