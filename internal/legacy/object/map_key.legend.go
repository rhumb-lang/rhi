package object

import (
	"arena"
	"math/rand"
	"reflect"
)

type KeyLegend struct {
	BaseLegend
	Data   int64
	Fields []Field
}

func (l KeyLegend) IsLegend() {}

func (l KeyLegend) WHAT() string { return "Legend[KeyMap]" }

func (x KeyLegend) Equals(y Legend) bool {
	if yMap, ok := y.(KeyLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x KeyLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(KeyLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewKeyLegend(a *arena.Arena, data string) *KeyLegend {
	num := int64(rand.Int())
	addr := arena.New[KeyLegend](a)
	*addr = KeyLegend{
		Data:   num,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	keyName := NewText(a, "_[']_")
	keyValue := NewText(a, data)
	addr.Fields = append(addr.Fields, Field{
		Name:    keyName,
		Mutable: false,
		Value:   keyValue,
	})
	return addr
}
