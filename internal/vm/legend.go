package vm

import "git.sr.ht/~madcapjake/grhumb/internal/word"

// type Legend struct {
// 	Mark       Word
// 	MetaLegend Word
// }

// type BaseMapLegend struct {
// 	Legend
// 	TrashSweep Word
// 	Size       Word
// 	Length     Word
// 	ReqLink    Word // pointer, circular dependency list
// 	DepLink    Word // pointer, circular dependency list
// }

// type MainMapLegend struct {
// 	BaseMapLegend
// 	Field []Descriptor
// }

// type ListMapLegend struct {
// 	BaseMapLegend
// 	Items Word
// 	Field []Descriptor
// }

// type TextMapLegend struct {
// 	BaseMapLegend
// 	Runes Word
// 	Field []Descriptor
// }

// type FuncMapLegend struct {
// 	BaseMapLegend
// 	Chunk Word
// 	Field []Descriptor
// }

// type MetaMapLegend struct {
// 	BaseMapLegend
// 	Field []Descriptor
// }

type MainMapLegend []word.Word
type ListMapLegend []word.Word
type TextMapLegend []word.Word
type FuncMapLegend []word.Word
type MetaMapLegend []word.Word

const (
	base_lgd_offset uint64 = 1
	base_swp_offset uint64 = 2
	base_sze_offset uint64 = 3
	base_len_offset uint64 = 4
	base_req_offset uint64 = 5
	base_dep_offset uint64 = 6
	main_fld_offset uint64 = 7
	text_arr_offset uint64 = 8
	text_fld_offset uint64 = 9
	rout_chu_offset uint64 = 10
	rout_fld_offset uint64 = 11
)

func NewBaseMapLegend(
	mark word.Word,
	legAddr word.Word,
	length uint32,
	dWords []word.Word,
) []word.Word {
	words := make([]word.Word, 0, 8)
	words = append(words,
		/* Mark:    */ mark,
		/* Legend:  */ legAddr,
		/* Sweep:   */ word.FromAddress(0),
		/* Size:    */ word.FromInt(7+length),
		/* Length:  */ word.FromInt(length),
		/* ReqLink: */ word.FromAddress(0),
		/* DepLink: */ word.FromAddress(0),
	)
	words = append(words, dWords...)
	return words
}

func wordsFromDescriptors(
	d ...Descriptor,
) (buf []word.Word) {
	// buf = make([]word.Word, 0, len(d)*3)
	// for _, desc := range d {
	// 	buf = append(buf, d[descIndex]...)
	// }
	return
}

func NewMainMapLegend(
	legAddr word.Word,
	descs ...Descriptor,
) MainMapLegend {
	descCount := uint32(len(descs))
	dWords := wordsFromDescriptors(descs...)
	return NewBaseMapLegend(
		word.Word(word.MAIN_LGD),
		legAddr,
		descCount,
		dWords,
	)
}

func NewTextMapLegend(
	legAddr word.Word,
	descs ...Descriptor,
) TextMapLegend {
	descCount := uint32(len(descs))
	dWords := wordsFromDescriptors(descs...)
	return NewBaseMapLegend(
		word.Word(word.TEXT_LGD),
		legAddr,
		descCount,
		dWords,
	)
}

func NewRoutMapLegend(
	legAddr word.Word,
	descs ...Descriptor,
) FuncMapLegend {
	descCount := uint32(len(descs))
	dWords := wordsFromDescriptors(descs...)
	return NewBaseMapLegend(
		word.Word(word.ROUT_LGD),
		legAddr,
		descCount,
		dWords,
	)
}

func NewMetaMapLegend(descs ...Descriptor) MetaMapLegend {
	descCount := uint32(len(descs))
	dWords := wordsFromDescriptors(descs...)
	return NewBaseMapLegend(
		word.Word(word.META_LGD),
		word.FromAddress(0),
		descCount,
		dWords,
	)
}

type ArrayLegend struct {
	Mark       word.Word
	MetaLegend word.Word
	TrashSweep word.Word
	Field      []Descriptor
}

func NewArrayLegend(mark word.Word) ArrayLegend {
	return ArrayLegend{
		Mark:       mark,
		MetaLegend: word.FromAddress(0),
		TrashSweep: word.Empty(),
		Field:      make([]Descriptor, 0),
	}
}
