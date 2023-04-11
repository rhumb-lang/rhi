package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

const (
	lgd_lgd_offset uint64 = 1
	lgd_swp_offset uint64 = 2
	lgd_sze_offset uint64 = 3
	lgd_len_offset uint64 = 4
	lgd_req_offset uint64 = 5
	lgd_dep_offset uint64 = 6
	lgd_dat_offset uint64 = 7
	lgd_fld_offset uint64 = 8
)

type Legend struct{ Id uint64 }

func CheckBaseLegendWords(
	vm *VirtualMachine,
	isCorrectKindOf func(m word.Word) bool,
	i uint64,
) {
	if !(isCorrectKindOf(vm.Heap[i])) {
		panic("not correct kind of legend mark")
	}
	if !(vm.Heap[i+lgd_lgd_offset].IsAddress()) {
		panic("meta legend word is not an address")
	}
	if !(vm.Heap[i+lgd_swp_offset].IsAddress()) {
		panic("gc sweep word is not an address")
	}
	if !(vm.Heap[i+lgd_sze_offset].IsInteger()) {
		panic("object size word is not an integer")
	}
	if !(vm.Heap[i+lgd_len_offset].IsInteger()) {
		panic("object length word is not an integer")
	}
	if !(vm.Heap[i+lgd_req_offset].IsAddress()) {
		panic("requirement link word is not an address")
	}
	if !(vm.Heap[i+lgd_dep_offset].IsAddress()) {
		panic("dependency link word is not an address")
	}
}

func (l Legend) NewDescriptor(
	vm *VirtualMachine,
	kind uint64,
	name RuneArray,
	offset int,
) Legend {
	oldLen := l.Length(vm)
	oldSze := l.Size(vm)
	descWords := []word.Word{
		word.Word(kind),
		word.FromAddress(name.id),
		word.FromAddress(0), // TODO: Req link
		word.FromAddress(0), // TODO: Dep link
		word.FromInt(int32(offset)),
	}
	loc, _ := vm.Allocate(int(l.Id), int(l.Size(vm)), descWords...)
	if loc != l.Id {
		l.Id = loc
	}
	l.SetLength(vm, oldLen+1)
	l.SetSize(vm, oldSze+5)
	return l
}

// Returns the location of the name provided
//
// TODO: Update to return word which is either address or constant
func (l Legend) Get(vm *VirtualMachine, name RuneArray) (
	idx int,
	err error,
) {
	// traverse the 5-word descriptors
	for i := range make([]int, l.Length(vm)*5) {
		if i%5 == 0 {
			descMark := vm.Heap[l.Id+lgd_fld_offset+uint64(i)]
			desc := ReviveDescriptor(vm, descMark)
			if desc.Name(vm) == name.String(vm) {
				return int(desc.Data(vm).AsInt()), nil
			}
		}
	}

	return -1, fmt.Errorf("couldn't find word")
}

func (l Legend) Length(vm *VirtualMachine) int32 {
	return vm.Heap[l.Id+lgd_len_offset].AsInt()
}

func (l Legend) SetLength(vm *VirtualMachine, i int32) {
	vm.Heap[l.Id+lgd_len_offset] = word.FromInt(i)
}

func (l Legend) Size(vm *VirtualMachine) int32 {
	return vm.Heap[l.Id+lgd_sze_offset].AsInt()
}

func (l Legend) SetSize(vm *VirtualMachine, i int32) {
	vm.Heap[l.Id+lgd_sze_offset] = word.FromInt(i)
}

func (l Legend) Data(vm *VirtualMachine) uint64 {
	return vm.Heap[l.Id+lgd_dat_offset].AsAddr()
}

func (l Legend) Append(vm *VirtualMachine, newWords ...word.Word) {
	mark := vm.Heap[l.Id]
	if mark.IsListLegendMark() {
		data := ReviveWordArray(vm, l.Data(vm))
		data.Append(vm, newWords...)
		// } else if mark.IsTextLegendMark() {
		// 	data := ReviveRuneArray(vm, l.Data(vm))
		// 	data.Append(vm, newWords...)
	} else {
		panic("no other append type implemented at this time")
	}
}
