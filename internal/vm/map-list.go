package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

func NewListMap(virtualMachine *VirtualMachine) Map {
	legend := NewListLegend(virtualMachine)
	words := make([]word.Word, 0, map_fld_offset)
	words = append(words,
		/* Mark:   */ word.Word(word.LIST_MAP),
		/* Legend: */ word.FromAddress(legend.Id),
	)
	if loc, err := virtualMachine.ReAllocate(words...); err != nil {
		panic("allocation failed")
	} else {
		return Map{Id: loc}
	}
}

func ReviveListMap(virtualMachine *VirtualMachine, addr word.Word) Map {
	i := addr.AsAddr()
	CheckMapWords(
		virtualMachine,
		func(m word.Word) bool { return m.IsListMapMark() },
		i,
	)
	return Map{Id: i}
}
