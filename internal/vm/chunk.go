package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

const (
	chunk_ca_offset uint64 = 1
	chunk_wa_offset uint64 = 2
)

type Chunk struct {
	// address in vm's heap
	id uint64

	// temp line tracking
	line int
}

func NewChunk(vm *VirtualMachine) Chunk {
	var (
		ch [3]word.Word = [3]word.Word{}
		ca CodeArray    = NewCodeArray(vm, word.FromAddress(0))
		wa WordArray    = NewWordArray(vm, word.FromAddress(0))
	)
	ch[0] = word.Word(word.CMP_UNIT)
	ch[1] = word.FromAddress(int(ca.id))
	ch[2] = word.FromAddress(int(wa.id))
	id, err := vm.ReAllocate(ch[:]...)
	if err != nil {
		panic("failed to reallocate onto heap")
	}
	return Chunk{id, 0}
}

func ReviveChunk(vm *VirtualMachine, addr word.Word) Chunk {
	i := addr.AsAddr()
	mark := vm.heap[i]
	if !(mark.IsCodeArrayMark()) {
		panic("not a chunk mark")
	}
	return Chunk{i, 0}
}

func (ch *Chunk) WriteCode(vm *VirtualMachine, line int, codes []Code) {

	instructions := ch.Instructions(vm)
	if line > ch.line {
		instructions.SetCodes(vm, true, codes...)
		ch.line = line
	} else {
		instructions.SetCodes(vm, false, codes...)
	}
	caId := ch.id + chunk_ca_offset
	if instructions.id != vm.heap[caId].AsAddr() {
		vm.heap[caId] = word.FromAddress(int(instructions.id))
	}
}

func (ch *Chunk) AddLiteral(vm *VirtualMachine, lit word.Word) (uint64, error) {
	literals := ch.Literals(vm)
	existingIndex, err := literals.IndexOf(vm, lit)
	if err != nil {
		id, err := literals.Append(vm, lit)
		if err != nil {
			panic("literals append failed")
		}
		if id != vm.heap[ch.id+chunk_wa_offset].AsAddr() {
			vm.heap[ch.id+chunk_wa_offset] = word.FromAddress(int(id))
		}
		return id, err
	} else {
		return uint64(existingIndex), nil
	}
}

func (ch Chunk) Instructions(vm *VirtualMachine) CodeArray {
	return ReviveCodeArray(vm, vm.heap[ch.id+chunk_ca_offset])
}

func (ch Chunk) Literals(vm *VirtualMachine) WordArray {
	return ReviveWordArray(vm, vm.heap[ch.id+chunk_wa_offset])
}

func (ch Chunk) Execute(vm *VirtualMachine) {
	var (
		instructions        = ch.Instructions(vm)
		cwLen        int    = int(instructions.Length(vm))
		idx          uint64 = 0
		buf          []byte = make([]byte, 0, cwLen*8)
	)
	for wordIndex := 0; wordIndex < cwLen; wordIndex++ {
		instructions.getCodesFromWord(vm, uint64(wordIndex), &buf)
		for b := range buf {
			if buf[b] == 0 {
				break
			}
			code := Code(buf[b])
			idx += uint64(code.Index())
			if code.IsIndexExtension() {
				continue
			} else {
				ch.execTagIndex(vm, code.Tag(), int(idx))
				idx = 0
			}
		}
		buf = make([]byte, 0, cwLen*8)
	}
}

func (ch Chunk) execTagIndex(vm *VirtualMachine, b byte, idx int) {
	var (
		code     Code  = Code(b)
		tag      uint8 = code.Tag()
		literals       = ch.Literals(vm)
	)
	switch tag {
	case TAG_VALUE_LITERAL:
		vm.AddLiteralToStack(literals.Get(vm, idx))
	case TAG_LOCAL_REQUEST:
		vm.SubmitLocalRequest(literals.Get(vm, idx))
	case TAG_INNER_REQUEST:
		vm.SubmitInnerRequest(literals.Get(vm, idx))
	case TAG_UNDER_REQUEST:
		vm.SubmitUnderRequest(literals.Get(vm, idx))
	case TAG_OUTER_REQUEST:
		vm.SubmitOuterRequest(literals.Get(vm, idx))
		// case TAG_EVENT_REQUEST:
		// 	vm.SubmitEventRequest(literals.Get(idx))
		// case TAG_REPLY_REQUEST:
		// 	vm.SubmitReplyRequest(literals.Get(idx))
	}
}

func (ch Chunk) Disassemble(vm *VirtualMachine) {
	fmt.Println("============= Chunk =============")
	var line int
	instructions := ch.Instructions(vm)
	buf := instructions.GetCodes(vm)
	for offset := 0; offset < int(instructions.Length(vm)); {
		if line < instructions.GetLine(vm, offset) {
			line += 1
			fmt.Printf("%4d ", line)
		} else {
			fmt.Printf("   | ")
		}
		line, offset = ch.DisassembleCode(vm, line, offset, &buf)
	}
}

func (ch Chunk) DisassembleCode(vm *VirtualMachine, currentLine, currentOffset int, bufPtr *[]byte) (int, int) {
	buf := *bufPtr
	fmt.Printf("%04d ", currentOffset)
	var recurse func(l, o, i int) (int, int)
	recurse = func(l, o, i int) (int, int) {
		tag := Code(buf[o]).Tag()
		switch tag {
		case TAG_IDX_EXTENSION:
			extension := Code(buf[o]).Index()
			if extension > 0 {
				return recurse(l, o+1, i+int(Code(buf[o]).Index()))
			} else { // Zero extension is code sentinel
				return l, o
			}
		default:
			idx := uint32(i + int(Code(buf[o]).Index()))
			fmt.Printf("%s %d '%v'\n",
				Code(buf[o]).String(),
				idx,
				ch.Literals(vm).Get(vm, int(idx)).Debug(),
			)
			return l, o + 1
		}
	}
	return recurse(currentLine, currentOffset, 0)
}
