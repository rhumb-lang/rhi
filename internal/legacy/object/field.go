package object

import (
	"arena"
	"fmt"
)

type Field struct {
	Name     Hashable
	Demands  Reference[Any]
	Supplies Reference[Any]
	Mutable  bool
	Value    Any
}

func (f Field) Equals(key Hashable) bool {
	return f.Name.Equals(key)
}

// func NewField(a *arena.Arena, n Hashable, m bool, v Any) *Field {
// 	addr := arena.New[Field](a)
// 	*addr = Field{Name: n, Mutable: m, Value: v}
// 	return addr
// }

type FieldNotFoundError struct{ string }

func (e FieldNotFoundError) Error() string {
	return fmt.Sprint("Unable to find field: ", e.string)
}

func fieldsGet(fields []Field, elems []Any, key Hashable) (Any, error) {
	for _, field := range fields {
		if field.Equals(key) {
			if field.Mutable {
				if position, ok := field.Value.(Number); ok {
					return elems[position.legend.Data], nil
				} else {
					return nil, fmt.Errorf("mutable field value is not whole number")
				}
			} else {
				return field.Value, nil
			}
		}
	}
	return nil, &FieldNotFoundError{key.WHAT()}
}

func fieldsSetMutable(memory *arena.Arena, fields *[]Field, elems *[]Any, key Hashable, val Any) (Any, error) {
	for _, field := range *fields {
		if field.Equals(key) {
			if field.Mutable {
				if position, ok := field.Value.(Number); ok {
					(*elems)[position.legend.Data] = val
					return val, nil
				} else {
					return nil, fmt.Errorf("mutable field value is not whole number")
				}
			} else {
				return nil, fmt.Errorf("the requested field is not mutable")
			}
		}
	}
	var position Number = *NewNumber(memory, int64(len(*elems)))
	*fields = append(*fields, Field{
		Name:    key,
		Mutable: true,
		Value:   position,
	})
	*elems = append(*elems, val)
	return val, nil
}

func fieldsSetImmutable(fields *[]Field, elems *[]Any, key Hashable, val Any) (Any, error) {
	for i, field := range *fields {
		if field.Equals(key) {
			if field.Mutable {
				if position, ok := field.Value.(Number); ok {
					(*fields)[i].Mutable = false
					(*elems)[position.legend.Data] = nil // TODO: convert to garbage collection
					(*fields)[i].Value = val
					return val, nil
				} else {
					return nil, fmt.Errorf("mutable field value is not whole number")
				}
			} else {
				return nil, fmt.Errorf("the requested field is not mutable")
			}
		}
	}
	*fields = append(*fields, Field{
		Name:    key,
		Mutable: false,
		Value:   val,
	})
	return val, nil
}
