package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

func NewRoutMap(virtualMachine *VirtualMachine) Map {
	legend := NewRoutLegend(virtualMachine)
	words := make([]word.Word, 0, map_fld_offset)
	words = append(words,
		/* Mark:   */ word.Word(word.ROUT_MAP),
		/* Legend: */ word.Word(legend.Id),
	)
	loc, err := virtualMachine.ReAllocate(words...)
	if err != nil {
		panic("allocation failed")
	}
	return Map{Id: loc}
}

func ReviveRoutMap(virtualMachine *VirtualMachine, addr word.Word) Map {
	i := addr.AsAddr()
	CheckMapWords(
		virtualMachine,
		func(m word.Word) bool { return m.IsRoutMapMark() },
		i,
	)
	return Map{Id: i}
}
