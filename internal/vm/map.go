package vm

import "git.sr.ht/~madcapjake/grhumb/internal/word"

// type RhumbMap struct {
// 	Mark   word.Word
// 	Legend word.Word
// 	Field  []word.Word
// }

type RhumbMap []word.Word

func NewMap(mark word.Word, legend word.Word, fields ...word.Word) RhumbMap {
	rmap := make([]word.Word, 0, 2+len(fields))
	rmap = append(rmap, mark, legend)
	rmap = append(rmap, fields...)
	return rmap
}

func NewMainMap(count uint32, legAddr word.Word) RhumbMap {
	return NewMap(word.Word(word.MAIN_MAP), legAddr)
}

func NewListMap(count uint32, legAddr word.Word) RhumbMap {
	return NewMap(word.Word(word.LIST_MAP), legAddr)
}

func NewTextMap(count uint32, legAddr word.Word) RhumbMap {
	return NewMap(word.Word(word.TEXT_MAP), legAddr)
}

func NewFuncMap(count uint32, legAddr word.Word) RhumbMap {
	return NewMap(word.Word(word.FUNC_MAP), legAddr)
}

func NewMetaMap(count uint32, legAddr word.Word) RhumbMap {
	return NewMap(word.Word(word.META_MAP), legAddr)
}
