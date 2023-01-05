package vm

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

var caLogger = log.New(io.Discard, "", log.LstdFlags)

func init() {
	if os.Getenv("RHUMB_CODE_ARRAY_DEBUG") == "1" {
		caLogger = log.New(os.Stdout, "{CA} ", log.LstdFlags)
	}
}

type CodeArray struct {
	Mark      Word
	Legend    Word // Address
	Length    Word
	CodeWords []Word // Words are packed 8-elem Codes
}

func NewCodeArray(words ...Word) CodeArray {
	return CodeArray{
		Mark:      Word(CODE_ARR),
		Legend:    Word(VAL_ADDR),
		Length:    WordFromInt(uint32(len(words))),
		CodeWords: words,
	}
}

func (ca *CodeArray) SetCodes(cs ...Code) {
	origLen := ca.Length.AsInt()
	newLen := uint32(len(cs))
	byteSize := uint32(8)
	bs := make([]byte, 0, byteSize)
	currWordVacancy := ca.currentVacancy()
	if currWordVacancy {
		ca.getCodesFromWord(ca.currentIndex(), &bs)
	}
	for c := uint32(0); c < newLen; c++ {
		bs = append(bs, byte(cs[c]))
		if (c+1)%byteSize == 0 {
			ca.CodeWords = append(
				ca.CodeWords,
				Word(binary.BigEndian.Uint64(bs)),
			)
			bs = make([]byte, 0, byteSize)
		}
	}
	if len(bs) > 0 {
		caLogger.Println(bs)
		for range bs[len(bs):byteSize] {
			bs = append(bs, 0x0)
		}
		newCodeWord := Word(binary.BigEndian.Uint64(bs))
		if currWordVacancy {
			ca.CodeWords[ca.currentIndex()] = newCodeWord
		} else {
			ca.CodeWords = append(ca.CodeWords, newCodeWord)
		}
	}
	ca.Length = WordFromInt(origLen + newLen)
}

func (ca *CodeArray) vacancy() uint32 {
	caLen := ca.Length.AsInt()
	modByte := caLen % 8
	return modByte
}

func (ca *CodeArray) currentVacancy() bool {
	return ca.vacancy() != 0
}

func (ca *CodeArray) currentIndex() int {
	if ca.currentVacancy() {
		return len(ca.CodeWords) - 1
	} else {
		return len(ca.CodeWords)
	}
}

func (ca *CodeArray) getCodesFromWord(wi int, b *[]byte) {
	var (
		buf  []byte = *b
		word Word   = ca.CodeWords[wi]
	)
	for offset := 56; offset >= 0; offset -= 8 {
		i := uint64(word) >> offset
		subByte := byte(i)
		if subByte == 0 {
			break
		}
		buf = append(buf, subByte)
	}
	*b = buf
}

func (ca *CodeArray) GetCodes() []byte {
	cwLen := len(ca.CodeWords)
	buf := make([]byte, 0, cwLen*8)
	for wordIndex := 0; wordIndex < cwLen; wordIndex++ {
		ca.getCodesFromWord(wordIndex, &buf)
	}
	return buf
}

func (ca *CodeArray) Len() int {
	return int(ca.Length.AsInt())
}
