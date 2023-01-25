package vm

import (
	"encoding/binary"
	"unicode/utf8"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

// type RuneArray struct {
// 	Mark   word.Word
// 	Legend word.Word // Address
// 	Size   word.Word
// 	Length word.Word
// 	Runes  []word.Word // Words are packed 2-elem runes
// }

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
		runesLen  uint32      = uint32(len(runes))
		runesSize uint32      = uint32(code_arr_offset) + runesLen
		raWords   []word.Word = make([]word.Word, 0, runesSize)
		runeBytes []byte
		wordBytes []byte
		offset    uint32
	)
	raWords = append(raWords,
		/* Mark:   */ word.Word(word.RUNE_ARR),
		/* Legend: */ legAddr,
		/* Size:   */ word.FromInt(runesSize),
		/* Length: */ word.FromInt(runesLen),
	)
	wordBytes = make([]byte, 8)
	for i, r := range runes {
		offset = uint32(i) % 2
		runeBytes = make([]byte, 4)
		binary.LittleEndian.PutUint32(runeBytes, uint32(r))
		if offset == 0 {
			copy(wordBytes, runeBytes)
			if i == len(runes)-1 {
				raWords = append(raWords,
					word.Word(binary.LittleEndian.Uint64(wordBytes)))
			}
		} else {
			for i, rb := range runeBytes {
				wordBytes[i+4] = rb
			}
			raWords = append(raWords,
				word.Word(binary.LittleEndian.Uint64(wordBytes)))
		}

	}

	loc, err := vm.ReAllocate(raWords...)
	if err != nil {
		panic("allocation failed")
	}
	return RuneArray{uint64(loc)}
}

// func GetOrMakeLabelRA(vm *VirtualMachine, text string) RuneArray {

// }

func ReviveRuneArray(vm *VirtualMachine, addr word.Word) RuneArray {
	i := addr.AsAddr()
	mark := vm.heap[i]
	if !(mark.IsRuneArrayMark()) {
		panic("not a rune array mark")
	}
	legend := vm.heap[i+rune_lgd_offset]
	if !(legend.IsAddress()) {
		panic("rune array legend word is not an address")
	}
	size := vm.heap[i+rune_sze_offset]
	if !(size.IsInteger()) {
		panic("rune array object size word is not an integer")
	}
	return RuneArray{i}
}

func (ra RuneArray) Id() uint64 { return ra.id }

func (ra RuneArray) Legend(vm *VirtualMachine) uint64 {
	return vm.heap[ra.id+rune_lgd_offset].AsAddr()
}
func (ra RuneArray) Size(vm *VirtualMachine) uint64 {
	return uint64(vm.heap[ra.id+rune_sze_offset].AsInt())
}
func (ra RuneArray) Length(vm *VirtualMachine) uint32 {
	return vm.heap[ra.id+rune_len_offset].AsInt()
}
func (ra RuneArray) Runes(vm *VirtualMachine) []rune {
	buf := make([]rune, 0, ra.Length(vm))
	for i := range vm.heap[ra.id+rune_arr_offset:] {
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(vm.heap[ra.id+uint64(i)]))
		buf = append(buf, rune(binary.LittleEndian.Uint32(bytes[:4])))
		buf = append(buf, rune(binary.LittleEndian.Uint32(bytes[4:])))
	}
	return buf
}

func (ra RuneArray) String(vm *VirtualMachine) string {
	return string(ra.Runes(vm))
}

func (ra RuneArray) Rune(vm *VirtualMachine, i uint32) rune {
	addr := i / 2
	offset := i % 2
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(vm.heap[ra.id+uint64(addr)]))
	if offset == 0 {
		return rune(binary.LittleEndian.Uint32(bytes[:4]))
	} else {
		return rune(binary.LittleEndian.Uint32(bytes[4:]))
	}
}

func (ra RuneArray) SetSize(vm *VirtualMachine, s uint32) {
	vm.heap[ra.id+rune_sze_offset] = word.FromInt(s)
}

func (ra RuneArray) SetLength(vm *VirtualMachine, l uint32) {
	vm.heap[ra.id+rune_len_offset] = word.FromInt(l)
}

func (ra RuneArray) SetRune(vm *VirtualMachine, i uint32, r rune) {
	var (
		addr      uint64 = uint64(i) / 2
		id        uint64 = ra.id + addr
		offset    uint32 = i % 2
		runeBytes []byte = make([]byte, 4)
		wordBytes []byte = make([]byte, 8)
	)
	binary.LittleEndian.PutUint32(runeBytes, uint32(r))
	binary.LittleEndian.PutUint64(wordBytes, uint64(vm.heap[id]))
	if offset == 0 {
		copy(wordBytes, runeBytes)
	} else {
		for i, rb := range runeBytes {
			wordBytes[i+4] = rb
		}
	}
	vm.heap[id] = word.Word(binary.LittleEndian.Uint64(wordBytes))
}
