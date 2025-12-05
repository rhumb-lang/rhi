package object

import (
	"arena"
	"fmt"
	"reflect"

	"git.sr.ht/~madcapjake/rhi/internal/color"
)

// Operators, variables and constants all start as label literals
// The bytecode decides how they will be utilized
type Label struct{ Value []rune }

func NewLabel(a *arena.Arena, v string) *Label {
	value := arena.MakeSlice[rune](a, len(v), len(v))
	addr := arena.New[Label](a)
	*addr = Label{
		Value: value,
	}
	copy(addr.Value, []rune(v))
	return addr
}

func (l Label) IsObject() {}
func (l Label) WHAT() string {
	return fmt.Sprint("Label: ", color.Green, "\"", string(l.Value), "\"", color.Reset)
}

func (x Label) Equals(y Any) bool {
	if yMap, ok := y.(*Label); ok {
		for i, r := range x.Value {
			if r != yMap.Value[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (x Label) DeepEquals(y Any) bool {
	if yMap, ok := y.(*Label); ok {
		return reflect.DeepEqual(x.Value, yMap.Value)
	}
	return false
}

func (key Label) HashBytes() []byte {
	return []byte(string(key.Value))
}
