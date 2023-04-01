package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

func NewSeleMap(virtualMachine *VirtualMachine) Map {
	legend := NewSeleLegend(virtualMachine)
	words := make([]word.Word, 0, map_fld_offset)
	words = append(words,
		/* Mark:   */ word.Word(word.SELE_MAP),
		/* Legend: */ word.Word(legend.Id),
	)
	loc, err := virtualMachine.ReAllocate(words...)
	if err != nil {
		panic("allocation failed")
	}
	return Map{Id: loc}
}

func ReviveSeleMap(virtualMachine *VirtualMachine, addr word.Word) Map {
	i := addr.AsAddr()
	CheckMapWords(
		virtualMachine,
		func(m word.Word) bool { return m.IsSeleMapMark() },
		i,
	)
	return Map{Id: i}
}
