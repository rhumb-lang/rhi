package object

import (
	"arena"
	"encoding/binary"
	"fmt"
)

type Version struct {
	legend *VersionLegend
	Values []Any
}

func (w Version) IsObject() {}

func (w Version) WHAT() string {
	return fmt.Sprint("Version: ", w.legend.Data)
}

func (x Version) Equals(y Any) bool {
	objY, ok := y.(Version)
	return ok && x.legend.Equals(objY.legend)
}

func (x Version) DeepEquals(y Any) bool {
	if yMap, ok := y.(Version); ok {
		return x.legend.DeepEquals(yMap.legend)
	}
	return false
}

func (key Version) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, key.legend.Data)
	return buf[:n]
}

func NewVersion(a *arena.Arena, val int64) *Version {
	values := arena.MakeSlice[Any](a, 0, 2)
	addr := arena.New[Version](a)
	*addr = Version{
		legend: NewVersionLegend(a, val),
		Values: values,
	}
	return addr
}
