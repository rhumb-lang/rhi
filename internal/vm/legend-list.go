package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

func NewListLegend(virtualMachine *VirtualMachine) Legend {
	words := make([]word.Word, 0, 8)
	init_size := int32(lgd_fld_offset)
	dataArray := NewWordArray(virtualMachine, word.FromAddress(0))
	words = append(words,
		/* Mark:    */ word.Word(word.LIST_LGD),
		/* Legend:  */ word.FromAddress(0), // TODO: Implement Meta Legend
		/* Sweep:   */ word.FromAddress(0),
		/* Size:    */ word.FromInt(init_size),
		/* Length:  */ word.FromInt(0),
		/* ReqLink: */ word.FromAddress(0),
		/* DepLink: */ word.FromAddress(0),
		/* Data:    */ word.FromAddress(dataArray.id),
	)
	loc, err := virtualMachine.ReAllocate(words...)
	if err != nil {
		panic("allocation failed")
	}
	return Legend{Id: loc}
}

func ReviveListLegend(virtualMachine *VirtualMachine, addr word.Word) Legend {
	i := addr.AsAddr()
	CheckBaseLegendWords(
		virtualMachine,
		func(m word.Word) bool { return m.IsListLegendMark() },
		i,
	)
	if !(virtualMachine.Heap[i+lgd_dat_offset].IsAddress()) {
		panic("list array word is not an address")
	}
	return Legend{Id: i}
}
