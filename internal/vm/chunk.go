package vm

import (
	"fmt"
)

type RefType uint8

const (
	RefInt RefType = iota
	RefFloat
	RefRune
	RefBool
	RefAddr
	RefLabel
)

type InstrRef struct {
	kind  RefType
	text  string
	value interface{}
}

func NewInstrRef(kind RefType, text string, value interface{}) InstrRef {
	return InstrRef{kind, text, value}
}

func (ir InstrRef) equals(other InstrRef) bool {
	return ir.text == other.text
}

type Chunk struct {
	instructions CodeArray

	// FIXME: Find a less memory intensive line tracking system
	lines []int

	// FIXME: Convert to array of refs & literals for true compilation
	prestack []InstrRef
}

func NewChunk(codes, refs []Word) Chunk {
	return Chunk{
		instructions: NewCodeArray(codes...),
		prestack:     make([]InstrRef, len(codes)/2),
		lines:        make([]int, len(codes)),
	}
}

func (ch *Chunk) WriteCode(line int, codes []Code) {
	for c := range codes {
		ch.instructions.SetCodes(codes[c])
		ch.lines = append(ch.lines, line)
	}
}

func (ch *Chunk) AddReference(ref InstrRef) uint32 {
	for r := range ch.prestack {
		if ch.prestack[r].equals(ref) {
			return uint32(r)
		}
	}
	newId := len(ch.prestack)
	ch.prestack = append(ch.prestack, ref)
	return uint32(newId)
}

func (ch Chunk) CodeWords() []Word {
	return ch.instructions.CodeWords
}

func (ch Chunk) CodeWordsCount() int {
	return len(ch.instructions.CodeWords)
}

func (ch Chunk) CodesFromWord(wordIndex int, buf *[]byte) {
	ch.instructions.getCodesFromWord(wordIndex, buf)
}

func (ch Chunk) Execute(vm *VirtualMachine) {
	var (
		cwLen int    = len(ch.instructions.CodeWords)
		idx   int    = 0
		buf   []byte = make([]byte, 0, cwLen*8)
	)
	for wordIndex := 0; wordIndex < cwLen; wordIndex++ {
		ch.instructions.getCodesFromWord(wordIndex, &buf)
		for b := range buf {
			if buf[b] == 0 {
				break
			}
			code := Code(buf[b])
			idx += int(code.Index())
			if code.IsIndexExtension() {
				continue
			} else {
				ch.execTagIndex(vm, code.Tag(), idx)
				idx = 0
			}
		}
		buf = make([]byte, 0, cwLen*8)
	}
}

func (ch Chunk) execTagIndex(vm *VirtualMachine, b byte, idx int) {
	code := Code(b)
	tag := code.Tag()
	switch tag {
	case TAG_VALUE_LITERAL:
		vm.AddValueToStack(ch.prestack[idx])
	case TAG_LOCAL_REQUEST:
		vm.SubmitLocalRequest(ch.prestack[idx])
	case TAG_INNER_REQUEST:
		vm.SubmitInnerRequest(ch.prestack[idx])
	case TAG_UNDER_REQUEST:
		vm.SubmitUnderRequest(ch.prestack[idx])
	case TAG_OUTER_REQUEST:
		vm.SubmitOuterRequest(ch.prestack[idx])
	case TAG_EVENT_REQUEST:
		vm.SubmitEventRequest(ch.prestack[idx])
	case TAG_REPLY_REQUEST:
		vm.SubmitReplyRequest(ch.prestack[idx])
	}
}

func (ch Chunk) Disassemble() {
	fmt.Println("============= Chunk =============")
	var line int
	buf := ch.instructions.GetCodes()
	for offset := 0; offset < ch.instructions.Len(); {
		line, offset = ch.DisassembleCode(line, offset, &buf)
	}
}

func (ch Chunk) DisassembleCode(currentLine, currentOffset int, bufPtr *[]byte) (int, int) {
	buf := *bufPtr
	fmt.Printf("%04d ", currentOffset)
	if currentOffset > 0 &&
		ch.lines[currentOffset] == ch.lines[currentOffset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", ch.lines[currentOffset])
	}

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
				ch.prestack[idx].text,
			)
			return l, o + 1
		}
	}
	return recurse(currentLine, currentOffset, 0)
}
