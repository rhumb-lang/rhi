package object

import (
	"arena"
	"fmt"
	"reflect"
)

type Scope struct {
	legend *ScopeLegend
	Values []Any
}

func NewScope(a *arena.Arena, parent *Scope) *Scope {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Scope](a)
	*addr = Scope{
		legend: NewScopeLegend(a, parent),
		Values: values,
	}
	return addr
}

func (o Scope) IsObject() {}

func (o Scope) WHAT() string {
	return fmt.Sprint("Scope: ", o.legend.Data)
}

func isNil(x any) bool {
	defer func() { recover() }() // nolint:errcheck
	return x == nil || reflect.ValueOf(x).IsNil()
}

func (o Scope) Parent() (*Scope, bool) {
	if isNil(o.legend.Data) {
		return nil, false
	} else {
		return o.legend.Data, true
	}
}

func (o Scope) Get(key Hashable) (Any, error) {
	if val, err := fieldsGet(o.legend.Fields, o.Values, key); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			if next, ok := o.Parent(); ok {
				return next.Get(key)
			} else {
				return nil, err
			}
		default:
			return nil, err
		}
	} else {
		return val, nil
	}
}

func (o *Scope) Set(mem *arena.Arena, key Hashable, val Any, mutable bool) (Any, error) {
	if mutable {
		return fieldsSetMutable(mem, &o.legend.Fields, &o.Values, key, val)
	} else {
		return fieldsSetImmutable(&o.legend.Fields, &o.Values, key, val)
	}
}

func (x Scope) Equals(y Any) bool {
	objY, ok := y.(Scope)
	return ok && x.legend.Equals(objY.legend)
}

func (x Scope) DeepEquals(y Any) bool {
	if yMap, ok := y.(Scope); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}
