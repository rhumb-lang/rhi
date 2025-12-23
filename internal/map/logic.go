package mapval

import "fmt"

// NewMap creates an empty map.
func NewMap() *Map {
	return &Map{
		Legend: NewLegend(),
		Fields: []Value{},
	}
}

func NewLegend() *Legend {
	return &Legend{
		Kind:   LegendMap,
		Fields: []FieldDesc{},
	}
}

// Get retrieves a value from the map by key.
// It follows the prototype chain if not found (Delegation).
func (m *Map) Get(key string) (Value, bool) {
	// 1. Local Lookup
	idx := m.Legend.FindIndex(key)
	if idx != -1 {
		return m.Fields[idx], true
	}
	
	// 2. Delegation (Look for subfields starting with @)
	// Iterate fields, if key starts with @, it's a parent.
	// Recursively search parents.
	// TODO: Optimization (cache parents, etc.)
	for i, desc := range m.Legend.Fields {
		if len(desc.Name) > 0 && desc.Name[0] == '@' {
			parentVal := m.Fields[i]
			if parentVal.Type == ValObject {
				if parentMap, ok := parentVal.Obj.(*Map); ok {
					if val, found := parentMap.Get(key); found {
						return val, true
					}
				}
			}
		}
	}
	
	return Value{}, false
}

// Set sets a value in the map.
// It triggers a Legend transition if the field is new.
func (m *Map) Set(key string, val Value, mutable bool) error {
	idx := m.Legend.FindIndex(key)

	if idx != -1 {
		// Update existing
		// Check mutability constraint
		desc := &m.Legend.Fields[idx]
		if desc.Kind == FieldImmutable {
			return fmt.Errorf("cannot assign value to immutable field '%s'", key)
		}
		
		// If the setter is immutable (.=), and field was mutable, freeze it.
		if !mutable {
			desc.Kind = FieldImmutable
		}

		m.Fields[idx] = val
		return nil
	}

	// New Field -> Transition Legend
	// For MVP, we just append to the current Legend (modifying it in place if it's unique, or COW).
	
	kind := FieldImmutable
	if mutable {
		kind = FieldMutable
	}

	m.Legend.Fields = append(m.Legend.Fields, FieldDesc{Name: key, Kind: kind})
	m.Fields = append(m.Fields, val)
	return nil
}

// FindIndex finds the index of a field in the legend.
func (l *Legend) FindIndex(key string) int {
	for i, f := range l.Fields {
		if f.Name == key {
			return i
		}
	}
	return -1
}
