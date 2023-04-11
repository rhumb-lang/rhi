package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

const (
	chunk_ca_offset uint64 = 1
	chunk_wa_offset uint64 = 2
)

type RhumbChunk struct {
	// address in vm's heap
	id uint64

	// temp line tracking
	line int
}

func NewChunk(vm *VirtualMachine) RhumbChunk {
	var (
		ch [3]word.Word = [3]word.Word{}
		ca CodeArray    = NewCodeArray(vm, word.FromAddress(0))
		wa WordArray    = NewWordArray(vm, word.FromAddress(0))
	)
	ch[0] = word.Word(word.CMP_UNIT)
	ch[1] = word.FromAddress(ca.id)
	ch[2] = word.FromAddress(wa.id)
	id, err := vm.ReAllocate(ch[:]...)
	if err != nil {
		panic("failed to reallocate onto heap")
	}
	return RhumbChunk{id, 0}
}

func ReviveChunk(vm *VirtualMachine, addr word.Word) RhumbChunk {
	i := addr.AsAddr()
	mark := vm.Heap[i]
	if !(mark.IsCodeArrayMark()) {
		panic("not a chunk mark")
	}
	return RhumbChunk{i, 0}
}

func (ch *RhumbChunk) WriteCode(vm *VirtualMachine, line int, codes []Code) {

	instructions := ch.ReviveInstrs(vm)
	if line > ch.line {
		instructions.SetCodes(vm, true, codes...)
		ch.line = line
	} else {
		instructions.SetCodes(vm, false, codes...)
	}
	if instructions.id != ch.Instructions(vm).AsAddr() {
		ch.SetInstructionsAddress(vm, instructions.id)
	}
}

func (ch RhumbChunk) SetInstructionsAddress(vm *VirtualMachine, addr uint64) {
	vm.Heap[ch.id+chunk_ca_offset] = word.FromAddress(addr)
}

func (ch RhumbChunk) SetLiteralsAddress(vm *VirtualMachine, addr uint64) {
	vm.Heap[ch.id+chunk_wa_offset] = word.FromAddress(addr)
}

func (ch *RhumbChunk) AddLiteral(vm *VirtualMachine, lit word.Word) (uint64, error) {
	literals := ch.ReviveLits(vm)
	existingIndex, err := literals.IndexOf(vm, lit)
	if err != nil {
		waID, err := literals.Append(vm, lit)
		if err != nil {
			panic("literals append failed")
		}
		if waID != ch.Literals(vm).AsAddr() {
			ch.SetLiteralsAddress(vm, waID)
		}
		newLastID := uint64(literals.Length(vm)) - 1
		return newLastID, err
	} else {
		return uint64(existingIndex), nil
	}
}

func (ch RhumbChunk) Instructions(vm *VirtualMachine) word.Word {
	return vm.Heap[ch.id+chunk_ca_offset]
}

func (ch RhumbChunk) ReviveInstrs(vm *VirtualMachine) CodeArray {
	return ReviveCodeArray(vm, ch.Instructions(vm))
}

func (ch RhumbChunk) Literals(vm *VirtualMachine) word.Word {
	return vm.Heap[ch.id+chunk_wa_offset]
}

func (ch RhumbChunk) ReviveLits(vm *VirtualMachine) WordArray {
	return ReviveWordArray(vm, ch.Literals(vm).AsAddr())
}

func (ch RhumbChunk) Execute(vm *VirtualMachine) {
	var (
		instructions        = ch.ReviveInstrs(vm)
		cwLen        int    = int(instructions.Size(vm) - code_arr_offset)
		idx          uint64 = 0
		buf          []byte = make([]byte, 0, cwLen*8)
	)
	for wordIndex := 0; wordIndex < cwLen; wordIndex++ {
		w := vm.Heap[instructions.id+code_arr_offset+uint64(wordIndex)]
		if w.IsSentinel() {
			ch.line++
			continue
		}
		instructions.getCodesFromWord(vm, w, &buf)
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

func (ch RhumbChunk) execTagIndex(vm *VirtualMachine, tag byte, idx int) {
	var (
		literals WordArray = ch.ReviveLits(vm)
		lit      word.Word
	)
	switch tag {
	case TAG_VALUE_LITERAL:
		lit = literals.Get(vm, idx)
		fmt.Println("Executing literal tag:", lit.Debug())
		vm.AddLiteralToStack(lit)
	case TAG_LOCAL_REQUEST:
		lit = literals.Get(vm, idx)
		fmt.Println("Executing local tag:", lit.Debug())
		vm.SubmitLocalRequest(lit)
	case TAG_INNER_REQUEST:
		lit = literals.Get(vm, idx)
		fmt.Println("Executing inner tag:", lit.Debug())
		vm.SubmitInnerRequest(lit)
	case TAG_UNDER_REQUEST:
		lit = literals.Get(vm, idx)
		fmt.Println("Executing under tag:", lit.Debug())
		vm.SubmitUnderRequest(lit)
	case TAG_OUTER_REQUEST:
		lit = literals.Get(vm, idx)
		fmt.Println("Executing outer tag:", lit.Debug())
		vm.SubmitOuterRequest(lit)
	// case TAG_EVENT_REQUEST:
	// 	vm.SubmitEventRequest(literals.Get(idx))
	// case TAG_REPLY_REQUEST:
	// 	vm.SubmitReplyRequest(literals.Get(idx))
	default:
		panic("not a valid tag")
	}
}

func (ch RhumbChunk) Disassemble(vm *VirtualMachine) {
	fmt.Println("============= Chunk =============")
	var line int
	instructions := ch.ReviveInstrs(vm)
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

func (ch RhumbChunk) DisassembleCode(vm *VirtualMachine, currentLine, currentOffset int, bufPtr *[]byte) (int, int) {
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
			fmt.Printf("%s %d %v\n",
				Code(buf[o]).String(),
				idx,
				ch.ReviveLits(vm).Get(vm, int(idx)).Debug(),
			)
			return l, o + 1
		}
	}
	return recurse(currentLine, currentOffset, 0)
}
