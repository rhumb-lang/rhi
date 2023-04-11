package vm

import (
	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

const (
	map_lgd_offset uint64 = 1
	map_fld_offset uint64 = 2
)

type Map struct{ Id uint64 }

func CheckMapWords(
	vm *VirtualMachine,
	isCorrectKindOf func(m word.Word) bool,
	i uint64,
) {
	if !(isCorrectKindOf(vm.Heap[i])) {
		panic("not correct kind of legend mark")
	}
	if !(vm.Heap[i+map_lgd_offset].IsAddress()) {
		panic("map legend word is not an address")
	}
}

func (m Map) Append(vm *VirtualMachine, newWords ...word.Word) Map {
	legend := m.ReviveLegend(vm)
	legend.Append(vm, newWords...)
	if legend.Id != m.Legend(vm) {
		m.SetLegend(vm, legend.Id)
	}
	return m
}

func (m Map) Get(vm *VirtualMachine, name RuneArray) word.Word {
	legend := m.ReviveLegend(vm)
	if i, err := legend.Get(vm, name); err != nil {
		return word.Empty()
	} else {
		return m.At(vm, i)
	}

}

func (m Map) Set(vm *VirtualMachine, name RuneArray, value word.Word) (
	result Map,
) {
	legend := m.ReviveLegend(vm)
	if i, err := legend.Get(vm, name); err != nil {
		// New label name
		result = m.Append(vm, value)
		legend = legend.NewDescriptor(
			vm,
			value.AsDesc(),
			name,
			int(result.Length(vm))-1,
		)
		if legend.Id != m.Legend(vm) {
			m.SetLegend(vm, legend.Id)
		}
	} else {
		// Existing label name, i = value offset
		vm.Heap[m.Locate(vm, i)] = value
		result = m
	}
	return
}

func (m Map) Locate(vm *VirtualMachine, i int) uint64 {
	return m.Id + lgd_dat_offset + uint64(i)
}

func (m Map) At(vm *VirtualMachine, i int) word.Word {
	mLen := m.Length(vm)
	if i < 0 || uint32(i) >= mLen {
		panic("index out of bounds")
	}
	return vm.Heap[m.Locate(vm, i)]
}

func (m Map) Legend(vm *VirtualMachine) uint64 {
	return vm.Heap[m.Id+map_lgd_offset].AsAddr()
}

func (m Map) SetLegend(vm *VirtualMachine, l uint64) {
	vm.Heap[m.Id+map_lgd_offset] = word.FromAddress(l)
}

func (m Map) Length(vm *VirtualMachine) uint32 {
	legend := m.ReviveLegend(vm)
	return legend.Length(vm)
}

func (m Map) SetLength(vm *VirtualMachine, i uint32) {
	legend := m.ReviveLegend(vm)
	legend.SetLength(vm, i)
}

func (m Map) Size(vm *VirtualMachine) uint32 {
	legend := m.ReviveLegend(vm)
	return legend.Size(vm)
}

func (m Map) SetSize(vm *VirtualMachine, i uint32) {
	legend := m.ReviveLegend(vm)
	legend.SetSize(vm, i)
}

func (m Map) ReviveLegend(vm *VirtualMachine) (legend Legend) {
	legendAddr := vm.Heap[m.Id+map_lgd_offset]
	switch vm.Heap[legendAddr.AsAddr()] {
	case word.Word(word.LIST_LGD):
		legend = ReviveListLegend(vm, legendAddr)
	case word.Word(word.TEXT_LGD):
		legend = ReviveTextLegend(vm, legendAddr)
	case word.Word(word.ROUT_LGD):
		legend = ReviveRoutLegend(vm, legendAddr)
	case word.Word(word.SELE_LGD):
		legend = ReviveSeleLegend(vm, legendAddr)
	case word.Word(word.META_LGD):
		legend = ReviveMetaLegend(vm, legendAddr)
	default:
		panic("Not a map legend")
	}
	return
}
