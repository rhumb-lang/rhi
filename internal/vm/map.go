package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/code"
	"git.sr.ht/~madcapjake/grhumb/internal/object"
)

type Map interface {
	object.Any
	IsMap()
}

type ListMap struct {
	home   *VirtualMachine
	legend object.Ref[Legend[[]object.Ref[object.Any]]]
	Values []interface{}
}

func (l ListMap) IsObject()    {}
func (l ListMap) isMap()       {}
func (l ListMap) WHAT() string { return "ListMap" }
func (l ListMap) At(index int) object.Ref[object.Any] {
	legend := Get[Legend[[]object.Ref[object.Any]]](l.home, l.legend)
	return legend.Data[index]
}

type RoutineMap struct {
	home   *VirtualMachine
	legend object.Ref[Legend[[]code.Any]]
	Values []object.Any
}

func (r RoutineMap) IsObject()    {}
func (r RoutineMap) isMap()       {}
func (r RoutineMap) WHAT() string { return "RoutineMap" }
func (r RoutineMap) Codes() []code.Any {
	legend := Get(r.home, r.legend)
	return legend.Data
}
func (r RoutineMap) Invoke() (object.Any, error) {
	legend := Get(r.home, r.legend)
	return Run(r.home, legend.Data)
}

type SelectorMap struct {
	home   *VirtualMachine
	legend object.Ref[Legend[[]code.Any]]
	Values []interface{}
}

func (s SelectorMap) IsObject()    {}
func (s SelectorMap) isMap()       {}
func (s SelectorMap) WHAT() string { return "SelectorMap" }

type TextMap struct {
	home   *VirtualMachine
	legend object.Ref[Legend[string]]
	Values []interface{}
}

func (t TextMap) IsObject()    {}
func (t TextMap) isMap()       {}
func (t TextMap) WHAT() string { return "TextMap" }
func (t TextMap) At(index int) rune {
	legend := Get(t.home, t.legend)
	for rIndex, r := range legend.Data {
		if rIndex == index {
			return r
		}
	}
	panic("index out of bounds")
}
func (t TextMap) String(index int) string {
	legend := Get(t.home, t.legend)
	return legend.Data
}
