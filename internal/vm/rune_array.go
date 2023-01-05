package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RuneArray struct {
	Mark   Word
	Legend Word // Address
	Length Word
	Runes  []Word // Words are packed 2-elem runes
}

func NewRuneArray(addr Word, words ...Word) RuneArray {
	return RuneArray{
		Word(RUNE_ARR),
		addr,
		WordFromInt(uint32(len(words))),
		words,
	}
}

func (ra *RuneArray) SetRunes(rs []rune) {
	var buf bytes.Buffer
	ra.Runes = nil
	for r := range rs {
		_, err := buf.WriteRune(rs[r])
		if err != nil {
			fmt.Println("buf.WriteRune failed:", err)
		}
		if (r+1)%2 == 0 {
			ra.Runes = append(
				ra.Runes,
				Word(binary.LittleEndian.Uint64(buf.Bytes())),
			)
			buf.Reset()
		}
	}
	ra.Length = WordFromInt(uint32(len(ra.Runes)))
}

func (ra *RuneArray) GetRunes() []rune {
	b := make([]byte, 8)
	buf := new(bytes.Buffer)
	for word := range ra.Runes {
		binary.LittleEndian.PutUint64(b, uint64(word))
		buf.Write(b)
		b = nil
	}
	return []rune(buf.String())
}
