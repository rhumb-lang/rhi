package vm

import (
	"encoding/binary"
	"unicode/utf8"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

const (
	rune_lgd_offset uint64 = 1
	rune_sze_offset uint64 = 2
	rune_len_offset uint64 = 3
	rune_arr_offset uint64 = 4
)

type RuneArray struct {
	// address in vm's heap
	id uint64
}

func RuneWords(str string) (words []word.Word) {
	var (
		rn         rune
		sz, length int
	)
	for id := 0; id < len(str); id += sz {
		rn, sz = utf8.DecodeRuneInString(str[id:])
		words = append(words, word.FromRune(rn))
		length++
	}
	return
}

func NewRuneArray(
	vm *VirtualMachine,
	legAddr word.Word,
	runes ...rune,
) RuneArray {
	var (
		runesLen  int32       = int32(len(runes))
		runesRem  int32       = int32(runesLen % 2)
		runesCap  int32       = int32(code_arr_offset) + (runesLen / 2) + runesRem
		raWords   []word.Word = make([]word.Word, 0, runesCap)
		runeBytes []byte
		wordBytes []byte
	)
	raWords = append(raWords,
		/* Mark:   */ word.Word(word.RUNE_ARR),
		/* Legend: */ legAddr,
		/* Size:   */ word.FromInt(runesCap),
		/* Length: */ word.FromInt(runesLen),
	)
	wordBytes = make([]byte, 8)
	for i, r := range runes {
		runeBytes = make([]byte, 4)
		binary.LittleEndian.PutUint32(runeBytes, uint32(r))
		if i%2 == 0 {
			copy(wordBytes, runeBytes)
			if i == len(runes)-1 {
				newWord := word.Word(binary.BigEndian.Uint64(wordBytes))
				raWords = append(raWords, word.Word(newWord))
			}
		} else {
			copy(wordBytes[4:], runeBytes)
			newWord := word.Word(binary.BigEndian.Uint64(wordBytes))
			raWords = append(raWords, newWord)
			wordBytes = make([]byte, 8)
		}

	}

	loc, err := vm.ReAllocate(raWords...)
	if err != nil {
		panic("allocation failed")
	}
	return RuneArray{uint64(loc)}
}

func ReviveRuneArray(vm *VirtualMachine, addr uint64) RuneArray {
	mark := vm.Heap[addr]
	if !(mark.IsRuneArrayMark()) {
		panic("not a rune array mark")
	}
	legend := vm.Heap[addr+rune_lgd_offset]
	if !(legend.IsAddress()) {
		panic("rune array legend word is not an address")
	}
	size := vm.Heap[addr+rune_sze_offset]
	if !(size.IsInteger()) {
		panic("rune array object size word is not an integer")
	}
	return RuneArray{addr}
}

func (ra RuneArray) Id() uint64 { return ra.id }

func (ra RuneArray) Legend(vm *VirtualMachine) uint64 {
	return vm.Heap[ra.id+rune_lgd_offset].AsAddr()
}
func (ra RuneArray) Size(vm *VirtualMachine) uint64 {
	return uint64(vm.Heap[ra.id+rune_sze_offset].AsInt())
}
func (ra RuneArray) Length(vm *VirtualMachine) int32 {
	return vm.Heap[ra.id+rune_len_offset].AsInt()
}
func (ra RuneArray) Runes(vm *VirtualMachine) []rune {
	raLen := uint64(ra.Length(vm))
	buf := make([]rune, 0, raLen)
	start := ra.id + rune_arr_offset
	for i := range vm.Heap[start : ra.id+ra.Size(vm)] {
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, uint64(vm.Heap[start+uint64(i)]))
		buf = append(buf, rune(binary.LittleEndian.Uint32(bytes[:4])))
		secondRune := rune(binary.LittleEndian.Uint32(bytes[4:]))
		if secondRune != 0 {
			buf = append(buf, secondRune)
		}
	}
	return buf
}

func (ra RuneArray) String(vm *VirtualMachine) string {
	return string(ra.Runes(vm))
}

func (ra RuneArray) Rune(vm *VirtualMachine, i int32) rune {
	addr := i / 2
	offset := i % 2
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(vm.Heap[ra.id+uint64(addr)]))
	if offset == 0 {
		return rune(binary.LittleEndian.Uint32(bytes[:4]))
	} else {
		return rune(binary.LittleEndian.Uint32(bytes[4:]))
	}
}

func (ra RuneArray) SetSize(vm *VirtualMachine, s int32) {
	vm.Heap[ra.id+rune_sze_offset] = word.FromInt(s)
}

func (ra RuneArray) SetLength(vm *VirtualMachine, l int32) {
	vm.Heap[ra.id+rune_len_offset] = word.FromInt(l)
}

func (ra RuneArray) SetRune(vm *VirtualMachine, i int32, r rune) {
	var (
		addr      uint64 = uint64(i) / 2
		id        uint64 = ra.id + addr
		offset    int32  = i % 2
		runeBytes []byte = make([]byte, 4)
		wordBytes []byte = make([]byte, 8)
	)
	binary.LittleEndian.PutUint32(runeBytes, uint32(r))
	binary.LittleEndian.PutUint64(wordBytes, uint64(vm.Heap[id]))
	if offset == 0 {
		copy(wordBytes, runeBytes)
	} else {
		for i, rb := range runeBytes {
			wordBytes[i+4] = rb
		}
	}
	vm.Heap[id] = word.Word(binary.LittleEndian.Uint64(wordBytes))
}
