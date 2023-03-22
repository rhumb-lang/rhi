package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

// type WordArray struct {
// 	Mark   word.Word
// 	Legend word.Word // Address
// 	Length word.Word // Integer
// 	Words  []word.Word
// }

const (
	word_lgd_offset uint64 = 1
	word_len_offset uint64 = 2
	word_arr_offset uint64 = 3
)

type WordArray struct {
	// address in vm's heap
	id uint64
}

func NewWordArray(
	vm *VirtualMachine,
	legAddr word.Word,
	words ...word.Word,
) WordArray {
	wordsLen := uint32(len(words))
	wordsSize := uint32(code_arr_offset) + wordsLen
	waWords := make([]word.Word, 0, wordsSize)
	waWords = append(waWords,
		/* Mark:   */ word.Word(word.MAIN_ARR),
		/* Legend: */ legAddr,
		/* Length: */ word.FromInt(wordsLen),
	)
	waWords = append(waWords, words...)

	loc, err := vm.ReAllocate(waWords...)
	if err != nil {
		panic("allocation failed")
	}
	return WordArray{uint64(loc)}
}

func ReviveWordArray(vm *VirtualMachine, addr word.Word) WordArray {
	i := addr.AsAddr()
	mark := vm.heap[i]
	if !(mark.IsMainArrayMark()) {
		panic("not a word array mark")
	}
	legend := vm.heap[i+word_lgd_offset]
	if !(legend.IsAddress()) {
		// fmt.Println(legend.Debug())
		panic("word array legend word is not an address")
	}
	length := vm.heap[i+word_len_offset]
	if !(length.IsInteger()) {
		panic("word array object length word is not an integer")
	}
	return WordArray{i}
}

func (wa WordArray) Find(vm *VirtualMachine, s string) (
	idx uint64,
	err error,
) {
	for i := range make([]int, wa.Length(vm)) {
		waVal := wa.Get(vm, i)
		if waVal.IsAddress() {
			heapVal := vm.heap[waVal.AsAddr()]
			if heapVal.IsRuneArrayMark() {
				heapRA := ReviveRuneArray(vm, waVal.AsAddr())
				strVal := heapRA.String(vm)
				if strVal == s {
					heapIdx := wa.id + word_arr_offset + uint64(i)
					idx = vm.heap[heapIdx].AsAddr()
					return
				}
			}
		}
	}
	err = fmt.Errorf("unable to find '%s'", s)
	return
}

func (wa WordArray) IndexOf(vm *VirtualMachine, x word.Word) (
	idx int,
	err error,
) {
	if x.IsAddress() {
		mkX := vm.heap[x.AsAddr()]
		if mkX.IsRuneArrayMark() {
			raX := ReviveRuneArray(vm, x.AsAddr())
			for i := range make([]int, wa.Length(vm)) {
				y := wa.Get(vm, i)
				if y.IsAddress() {
					mkY := vm.heap[y.AsAddr()]
					if mkY.IsRuneArrayMark() {
						ray := ReviveRuneArray(vm, y.AsAddr())
						if raX.String(vm) == ray.String(vm) {
							return i, nil
						}
					}
				}
			}
		}
	} else {
		for i := range make([]int, wa.Length(vm)) {
			if x.Equals(wa.Get(vm, i)) {
				return i, nil
			}
		}

	}
	return -1, fmt.Errorf("couldn't find word")
}

func (wa *WordArray) Append(vm *VirtualMachine, newWords ...word.Word) (uint64, error) {
	oldLen := uint32(wa.Length(vm))
	newLen := oldLen + uint32(len(newWords))
	id, err := vm.Allocate(int(wa.id), wa.Size(vm), newWords...)
	if err != nil {
		panic("word array append failed")
	} else {
		if id != wa.id {
			wa.id = id
		}
		wa.SetLength(vm, newLen)
		return id, err
	}
}

func (wa WordArray) Get(vm *VirtualMachine, i int) word.Word {
	waLen := wa.Length(vm)
	if i < 0 || i >= waLen {
		panic("index out of bounds")
	}
	return vm.heap[wa.id+word_arr_offset+uint64(i)]
}

func (wa WordArray) Legend(vm *VirtualMachine, i int) word.Word {
	return vm.heap[i]
}

func (wa WordArray) SetLength(vm *VirtualMachine, l uint32) {
	vm.heap[wa.id+word_len_offset] = word.FromInt(l)
}

func (wa WordArray) Length(vm *VirtualMachine) int {
	return int(vm.heap[wa.id+word_len_offset].AsInt())
}

func (wa WordArray) Size(vm *VirtualMachine) int {
	return wa.Length(vm) + 3
}
