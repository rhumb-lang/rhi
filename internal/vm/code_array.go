package vm

import (
	"encoding/binary"
	"io"
	"log"
	"os"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

var caLogger = log.New(io.Discard, "", log.LstdFlags)

func init() {
	if os.Getenv("RHUMB_CODE_ARRAY_DEBUG") == "1" {
		caLogger = log.New(os.Stdout, "{CA} ", log.LstdFlags)
	}
}

// type CodeArray struct {
// 	Mark      word.Word
// 	Legend    word.Word // Address
//  Size      word.Word
// 	Length    word.Word
// 	CodeWords []word.Word // Words are packed 8-elem Codes
// }

const (
	code_lgd_offset uint64 = 1
	code_sze_offset uint64 = 2
	code_len_offset uint64 = 3
	code_arr_offset uint64 = 4
)

type CodeArray struct {
	// address in vm's heap
	id uint64
}

func NewCodeArray(
	vm *VirtualMachine,
	legAddr word.Word,
	words ...word.Word,
) CodeArray {
	wordsLen := uint32(len(words))
	wordsSize := uint32(code_arr_offset) + wordsLen
	caWords := make([]word.Word, 0, wordsSize)
	caWords = append(caWords,
		/* Mark:   */ word.Word(word.CODE_ARR),
		/* Legend: */ legAddr,
		/* Size:   */ word.FromInt(wordsSize),
		/* Length: */ word.FromInt(wordsLen),
	)
	caWords = append(caWords, words...)

	loc, err := vm.ReAllocate(caWords...)
	if err != nil {
		panic("allocation failed")
	}
	return CodeArray{uint64(loc)}
}

func ReviveCodeArray(vm *VirtualMachine, addr word.Word) CodeArray {
	i := addr.AsAddr()
	mark := vm.heap[i]
	if !(mark.IsCodeArrayMark()) {
		panic("not a code array mark")
	}
	legend := vm.heap[i+code_lgd_offset]
	if !(legend.IsAddress()) {
		panic("code array legend word is not an address")
	}
	size := vm.heap[i+code_sze_offset]
	if !(size.IsInteger()) {
		panic("code array object size word is not an integer")
	}
	return CodeArray{i}
}

// Given a possible new line and some new codes, what
// will the new total word size of this object?
func (ca CodeArray) NewSize(vm *VirtualMachine, newLine bool, cs ...Code) (size uint64) {
	var (
		origSize    uint64 = ca.Size(vm)
		newLength   int    = len(cs)
		freshLength int
		wholeWords  int
		remainder   int
	)
	if newLine {
		size = origSize + 1
		wholeWords = newLength / 8
		remainder = newLength % 8
		if remainder == 0 {
			size += uint64(wholeWords)
		} else {
			size += uint64(wholeWords + 1)
		}
	} else {
		size = origSize
		id := ca.id + code_arr_offset + uint64(ca.Size(vm)-1)
		word := vm.heap[id]
		remainder = ca.getWordCodeCount(vm, word)
		freshLength = newLength - remainder
		wholeWords = freshLength / 8
		remainder = freshLength % 8
		if remainder == 0 {
			size += uint64(wholeWords)
		} else {
			size += uint64(wholeWords + 1)
		}
	}
	return
}

func (ca *CodeArray) SetCodes(
	vm *VirtualMachine,
	newLine bool,
	cs ...Code,
) {
	var (
		origLen    uint64 = uint64(ca.Length(vm))
		origSz     uint32 = uint32(ca.Size(vm))
		newLen     uint32 = uint32(len(cs))
		byteSz     uint32 = 8
		codeID     int
		code       Code
		initSz     uint32
		bs         []byte      = make([]byte, 0, byteSz)
		ws         []word.Word = make([]word.Word, 0, newLen/8)
		lastWordID uint64      = ca.id + uint64(ca.Size(vm)) - 1
	)
	if newLine {
		ws = append(ws, word.Sentinel())
	} else {
		if ca.lastWordVacancy(vm) {
			ca.getCodesFromWord(vm, vm.heap[lastWordID], &bs)
			vm.free[lastWordID] = true
			origSz--
			initSz = uint32(len(bs))
			tempCS := make([]Code, len(cs))
			copy(tempCS, cs)
			for codeID, code = range tempCS {
				bs = append(bs, byte(code))
				cs = cs[codeID+1:] // shift from main cs var
				if initSz+uint32(codeID)+1 == byteSz {
					ws = append(ws,
						word.Word(binary.BigEndian.Uint64(bs)))
					bs = make([]byte, 0, byteSz)
					break
				}
			}

		}
	}
	for codeID, code = range cs {
		bs = append(bs, byte(code))
		if codeID+1 == int(byteSz) {
			ws = append(ws, word.Word(binary.BigEndian.Uint64(bs)))
			bs = make([]byte, 0, byteSz)
		}
	}

	if len(bs) > 0 {
		caLogger.Println(bs)
		for range bs[len(bs):byteSz] {
			bs = append(bs, 0x0)
		}
		bsWord := word.Word(binary.BigEndian.Uint64(bs))
		ws = append(ws, bsWord)
	}
	wordsLen := uint32(len(ws))
	finalLength := uint32(origLen) + newLen
	ca.SetSize(vm, origSz+wordsLen)
	ca.SetLength(vm, finalLength)
	newId, _ := vm.Allocate(
		int(ca.id),
		int(origSz),
		ws...,
	)
	if newId != ca.id {
		// vm.UpdateAddresses(ca.id, newId)
		ca.id = newId
	}
}

func (ca *CodeArray) lastWordVacancy(vm *VirtualMachine) bool {
	var (
		lastIndex uint64    = uint64(ca.Size(vm) - 1)
		word      word.Word = vm.heap[ca.id+lastIndex]
	)
	if word.IsSentinel() {
		return false
	}
	return ca.getWordCodeCount(vm, word) != 0
}

// func (ca *CodeArray) currentIndex(vm *VirtualMachine) int {
// 	if ca.lastWordVacancy(vm) {
// 		return int(ca.Length(vm)) - 1
// 	} else {
// 		return int(ca.Length(vm))
// 	}
// }

func (ca CodeArray) getCodesFromWord(vm *VirtualMachine, word word.Word, b *[]byte) {
	var buf []byte = *b
	if word.IsSentinel() {
		panic("cannot get codes from sentinel")
	}
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

func (ca CodeArray) getWordCodeCount(vm *VirtualMachine, word word.Word) (count int) {
	if word.IsSentinel() {
		panic("cannot get codes from sentinel")
	}
	for offset := 56; offset >= 0; offset -= 8 {
		i := uint64(word) >> offset
		subByte := byte(i)
		if subByte == 0 {
			break
		}
		count = count + 1
	}
	return
}

func (ca CodeArray) GetCodes(vm *VirtualMachine) []byte {
	var (
		cwLen uint64 = uint64(ca.Size(vm))
		buf   []byte = make([]byte, 0, cwLen*8)
	)
	for wIndex := code_arr_offset; wIndex < cwLen; wIndex++ {
		w := vm.heap[ca.id+wIndex]
		if !(w.IsSentinel()) {
			ca.getCodesFromWord(vm, w, &buf)
		}
	}
	return buf
}

func (ca *CodeArray) GetLine(vm *VirtualMachine, codeIndex int) (lines int) {
	codes, cwLen, cwSize := 0, ca.Length(vm), ca.Size(vm)
	if uint32(codeIndex) > cwLen {
		panic("index greater than length")
	}
	for i := uint64(0); i < cwSize && codes <= codeIndex; i++ {
		if vm.heap[ca.id+code_arr_offset+i].IsSentinel() {
			lines += 1
		} else {
			var (
				id   uint64    = ca.id + code_arr_offset + i
				word word.Word = vm.heap[id]
			)
			codes += ca.getWordCodeCount(vm, word)
		}
	}
	return
}

func (ca CodeArray) Legend(vm *VirtualMachine) uint64 {
	return vm.heap[ca.id+code_lgd_offset].AsAddr()
}
func (ca CodeArray) Size(vm *VirtualMachine) uint64 {
	return uint64(vm.heap[ca.id+code_sze_offset].AsInt())
}
func (ca CodeArray) Length(vm *VirtualMachine) uint32 {
	return vm.heap[ca.id+code_len_offset].AsInt()
}

func (ca *CodeArray) SetSize(vm *VirtualMachine, s uint32) {
	vm.heap[ca.id+code_sze_offset] = word.FromInt(s)
}

func (ca *CodeArray) SetLength(vm *VirtualMachine, l uint32) {
	vm.heap[ca.id+code_len_offset] = word.FromInt(l)
}
