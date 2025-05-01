package object

import (
	"arena"
	"reflect"
)

type VersionLegend struct {
	BaseLegend
	Data   int64 // TODO: update to multiple version segments
	Fields []Field
}

func (l VersionLegend) IsLegend() {}

func (l VersionLegend) WHAT() string { return "Legend[RoutineMap]" }

func (x VersionLegend) Equals(y Legend) bool {
	if yMap, ok := y.(VersionLegend); ok {
		return reflect.DeepEqual(x.Data, yMap.Data)
	}
	return false
}

func (x VersionLegend) DeepEquals(y Legend) bool {
	if yMap, ok := y.(VersionLegend); ok {
		if reflect.DeepEqual(x.Data, yMap.Data) {
			return reflect.DeepEqual(x.Fields, yMap.Fields)
		}
	}
	return false
}

func NewVersionLegend(a *arena.Arena, data int64) *VersionLegend {
	addr := arena.New[VersionLegend](a)
	*addr = VersionLegend{
		Data:   data,
		Fields: arena.MakeSlice[Field](a, 0, 2),
	}
	return addr
}
