package vm

import "git.sr.ht/~madcapjake/grhumb/internal/word"

//	type Descriptor struct {
//		Mark Word // immutable, mutable, submap
//		Name Word // address to TextMap
//		Data Word // constant, field offset, or
//	}

const (
	desc_name_offset uint64 = 1
	desc_data_offset uint64 = 2
)

type Descriptor struct {
	vm *VirtualMachine
	id uint64
	at []word.Word
}

func NewDescriptor(
	vm *VirtualMachine,
	name word.Word, // string address
	kind uint64,
	value word.Word,
) Descriptor {
	descWords := make([]word.Word, 0, 2)
	descWords = append(descWords, word.Word(kind))
	descWords = append(descWords, name)
	descWords = append(descWords, value)
	loc, _ := vm.ReAllocate(descWords...)
	return Descriptor{vm, loc, vm.heap[loc : loc+2]}
}

func ReviveDescriptor(vm *VirtualMachine, addr word.Word) Descriptor {
	i := addr.AsAddr()
	mark := vm.heap[i]
	if !(mark.IsDescMark()) {
		panic("not a descriptor mark")
	}
	name := vm.heap[i+desc_name_offset]
	if !(name.IsAddress()) {
		panic("desciptor name word is not an address")
	}
	return Descriptor{vm, i, vm.heap[i : i+2]}
}
