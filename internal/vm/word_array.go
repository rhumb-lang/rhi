package vm

import "fmt"

type WordArray struct {
	Mark   Word
	Legend Word // Address
	Length Word
	Words  []Word // Words are normal single words
}

func NewWordArray(words ...Word) WordArray {
	return WordArray{
		Word(MAIN_ARR),
		Word(VAL_ADDR),
		WordFromInt(uint32(len(words))),
		words,
	}
}

func (wa *WordArray) IndexOf(w Word) (idx int, err error) {
	for i := range wa.Words {
		if wa.Words[i] == w {
			return i, nil
		}
	}
	return 0, fmt.Errorf("couldn't find word")
}

func (wa *WordArray) Add(w Word) {
	wa.Words = append(wa.Words, w)
	wa.Length = WordFromInt(uint32(len(wa.Words)))
}

func (wa *WordArray) Get(i uint32) Word {
	return wa.Words[i]
}

func (wa *WordArray) All() []Word {
	return wa.Words
}

func (wa *WordArray) Count() int {
	return int(wa.Length.AsInt())
}
