package vm

import "git.sr.ht/~madcapjake/grhumb/internal/word"

//	type Descriptor struct {
//		Mark Word // immutable, mutable, submap
//		Name Word // address to TextMap
//      Req  Word // requirement linked list
//      Dep  Word // dependency linked list
//		Data Word // constant, field offset, or
//	}

const (
	desc_nme_offset uint64 = 1
	desc_req_offset uint64 = 2
	desc_dep_offset uint64 = 3
	desc_dat_offset uint64 = 4
)

type Descriptor struct {
	Id uint64
}

func ReviveDescriptor(vm *VirtualMachine, addr word.Word) Descriptor {
	i := addr.AsAddr()
	mark := vm.Heap[i]
	if !(mark.IsDescMark()) {
		panic("not a descriptor mark")
	}
	nameAddr := vm.Heap[i+desc_nme_offset]
	if !(nameAddr.IsAddress()) {
		panic("desciptor name word is not an address")
	}
	if !(vm.Heap[nameAddr].IsRuneArrayMark()) {
		panic("descriptor name address does not point to rune array")
	}

	return Descriptor{i}
}

func (d Descriptor) Name(vm *VirtualMachine) string {
	nameAddr := vm.Heap[d.Id+desc_nme_offset]
	if !(nameAddr.IsAddress()) {
		panic("desciptor name word is not an address")
	}
	addr := nameAddr.AsAddr()
	pointedAt := vm.Heap[addr]
	if !(pointedAt.IsRuneArrayMark()) {
		panic("word provided does not point to a rune array")
	}
	name := ReviveRuneArray(vm, addr)
	return name.String(vm)
}

func (d Descriptor) Data(vm *VirtualMachine) word.Word {
	return vm.Heap[d.Id+desc_dat_offset]
}
