package vm

import (
	"fmt"

	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
)

// setField sets a field on a map.
func (vm *VM) setField(target, key, val mapval.Value, mutable bool) (mapval.Value, error) {
	if target.Type != mapval.ValObject {
		return mapval.NewEmpty(), fmt.Errorf("target is not a map")
	}

	m, ok := target.Obj.(*mapval.Map)
	if !ok {
		return mapval.NewEmpty(), fmt.Errorf("target is not a map")
	}

	// Convert key to string
	var keyStr string
	switch key.Type {
	case mapval.ValText:
		keyStr = key.Str
	case mapval.ValInteger: // Key ID?
		keyStr = fmt.Sprintf("%d", key.Integer) // Temporary
	default:
		return mapval.NewEmpty(), fmt.Errorf("invalid key type: %v", key.Type)
	}

	// 1. Check if field exists in Legend
	for i, fieldDesc := range m.Legend.Fields {
		if fieldDesc.Name == keyStr {
			// Found. Check mutability
			if fieldDesc.Kind == mapval.FieldImmutable {
				return mapval.NewEmpty(), fmt.Errorf("field '%s' is immutable", keyStr)
			}

			m.Fields[i] = val
			return val, nil
		}
	}

	// 2. Not found. Append.
	// TODO: Implement proper Hidden Class transitions.
	kind := mapval.FieldImmutable
	if mutable {
		kind = mapval.FieldMutable
	}

	m.Legend.Fields = append(m.Legend.Fields, mapval.FieldDesc{Name: keyStr, Kind: kind})
	m.Fields = append(m.Fields, val)

	return val, nil
}
