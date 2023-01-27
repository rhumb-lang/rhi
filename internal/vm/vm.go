package vm

import (
	"fmt"
	"io"
	"log"
	"os"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

var vmLogger = log.New(io.Discard, "", log.LstdFlags)

func init() {
	if os.Getenv("RHUMB_VM_DEBUG") == "1" {
		vmLogger = log.New(os.Stdout, "{VM} ", log.LstdFlags)
	}
}

type VirtualMachine struct {
	heap  []word.Word
	free  []bool
	stack []word.Word
	scope []map[string]uint64
	main  Chunk
}

var DEBUG_WIDTH int = 10

func incCheckNL(inc *int) {
	val := *inc
	val++
	if val%DEBUG_WIDTH == 0 {
		fmt.Println()
	}
	*inc = val
}

func (vm VirtualMachine) DebugHeap() {
	for i := 0; i < len(vm.heap); i++ {
		if i%DEBUG_WIDTH == 0 {
			fmt.Println()
		}
		if vm.free[i] {
			fmt.Printf("{%3v:           }", i)
		} else if vm.heap[i].IsRuneArrayMark() {
			j := i
			fmt.Printf("[%3v: %-9s ]", j, vm.heap[j].Debug())
			incCheckNL(&j)
			fmt.Printf("[%3v: %-9s ]", j, vm.heap[j].Debug())
			incCheckNL(&j)
			fmt.Printf("[%3v: %-9s ]", j, vm.heap[j].Debug())
			incCheckNL(&j)
			fmt.Printf("[%3v: %-9s ]", j, vm.heap[j].Debug())
			incCheckNL(&j)
			fmt.Printf("[%3v: ", j)
			ra := ReviveRuneArray(&vm, uint64(i))
			runes := ra.Runes(&vm)
			if len(runes) <= 2 {
				for _, r := range runes {
					fmt.Printf("'%s'", string(r))
					fmt.Print("  ")
				}
			} else {
				for i, r := range runes {
					fmt.Printf("'%s'", string(r))
					if i%2 == 1 {
						fmt.Print("  |")
						incCheckNL(&j)
						fmt.Print("|     ")
					} else {
						fmt.Print("  ")
					}
				}
				for range make([]int, (ra.Length(&vm) % 2)) {
					fmt.Print("     ")
				}

			}
			fmt.Print("]")
			i += int(ra.Size(&vm)) - 1
		} else {
			fmt.Printf("[%3v: %-9s ]", i, vm.heap[i].Debug())
		}
	}
	fmt.Println()
}

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)
	const init_heap_len int = 10
	vm.heap = make([]word.Word, 0, init_heap_len)
	vm.free = make([]bool, 0, init_heap_len)
	vm.ReAllocate(word.Word(word.SENTINEL))
	vm.stack = make([]word.Word, 0)
	vm.scope = make([]map[string]uint64, 0)
	vm.scope = append(vm.scope, make(map[string]uint64))
	vm.main = NewChunk(vm)
	return vm
}

// Traverses the free list until a sufficiently-sized chunk
// of free slots is available to place the word arguments.
//
// Returns the index of the final heap location.
func (vm *VirtualMachine) ReAllocate(ws ...word.Word) (uint64, error) {
	obj := *vm
	size := len(ws)
	if size == 0 {
		return 0, fmt.Errorf("no words provided")
	}
	var first, next int
	for first = 0; first < len(obj.free); first += next {
		next = 0
		if obj.free[first] {
			if size == 1 {
				vmLogger.Println("adding word to index:", first)
				obj.heap[first] = ws[0]
				obj.free[first] = false
				obj.DebugHeap()
				*vm = obj
				return uint64(first), nil
			} else {
				last := first + size
				for next = range obj.free[first+1 : last] {
					if !(obj.free[next]) {
						next++
						break
					} else if next == last {
						vmLogger.Println("appending words to index:", first)
						obj.allocInPlace(first, last, ws...)
						obj.DebugHeap()
						*vm = obj
						return uint64(first), nil
					}
				}
			}
		} else {
			next = 1
		}
	}
	// no available chunk in existing memory locations.
	first = len(obj.heap)
	vmLogger.Println("appending words to end:", first)
	obj = appendRhumb(obj, ws)
	*vm = obj
	return uint64(first), nil
}

// Traverses the free list directly after the current
// used memory for this object. If there is sufficient free
// slots, then allocInPlace. Otherwise, call ReAllocate.
//
// Returns the index of the final heap location.
func (vm *VirtualMachine) Allocate(
	loc int,
	oldSize int,
	ws ...word.Word,
) (uint64, error) {
	var (
		obj       VirtualMachine = *vm
		newSize   int            = len(ws)
		lastOldId int            = loc + oldSize
	)
	if newSize == 0 {
		return 0, fmt.Errorf("no words provided")
	}
	var i int = lastOldId
	if i == len(obj.free) {
		vmLogger.Println("appending words to end:", i)
		obj = appendRhumb(obj, ws)
		*vm = obj
		return uint64(loc), nil
	}

	if obj.free[i] {
		if newSize == 1 {
			vmLogger.Println("appending word to index:", i)
			obj.heap[i] = ws[0]
			obj.free[i] = false
			obj.DebugHeap()
			*vm = obj
			return uint64(loc), nil
		} else {
			last := i + newSize - 1
			for i = i + 1; i <= last; i++ {
				if !(obj.free[i]) {
					break
				} else if i == last {
					first := lastOldId
					vmLogger.Println("appending words to index:", first)
					obj.allocInPlace(first, last, ws...)
					obj.DebugHeap()
					*vm = obj
					return uint64(loc), nil
				}
			}
		}
	}
	// Subsequent memory spots unavailable, search for any
	// available memory spot across heap
	totalWords := make([]word.Word, 0, oldSize+newSize)
	totalWords = append(totalWords, obj.heap[loc:lastOldId]...)
	totalWords = append(totalWords, ws...)

	for f := range obj.free[loc:lastOldId] {
		obj.free[loc+f] = true
	}

	id, err := obj.ReAllocate(totalWords...)
	*vm = obj
	return id, err
}

// Appends to heap and free slices, used in allocate and reallocate
func appendRhumb(vm VirtualMachine, ws []word.Word) VirtualMachine {
	for i := range ws {
		vm.heap = append(vm.heap, ws[i])
		vm.free = append(vm.free, false)
	}
	vm.DebugHeap()
	return vm
}

// Only run this if you are absolutely sure there is room
// to allocate in the heap. Overwriting memory is possible.
func (vm VirtualMachine) allocInPlace(x, y int, ws ...word.Word) {
	for hID, wID := x, 0; hID <= y; hID, wID = hID+1, wID+1 {
		vm.heap[hID], vm.free[hID] = ws[wID], false
	}
}

// func (vm VirtualMachine) UpdateAddresses(oldId, newId uint64) {
// 	for i, free := range vm.free {
// 		if !(free) {
// 			w := vm.heap[i]
// 			if w.IsAddress() && w.AsAddr() == oldId {
// 				vm.heap[i] = word.FromAddress(int(newId))
// 			}
// 		}
// 	}
// }

func (vm *VirtualMachine) WriteCodeToMain(
	line int,
	lit word.Word,
	codeFactory func(i uint64) []Code,
) {
	id, _ := vm.main.AddLiteral(vm, lit)
	codes := codeFactory(id)
	vm.main.WriteCode(vm, line, codes)
}

func (vm *VirtualMachine) Disassemble() {
	vm.main.Disassemble(vm)
}

func (vm *VirtualMachine) Execute() {
	vm.main.Execute(vm)
}

func logAddedToStack(stack []word.Word, txt string) {
	logStr := fmt.Sprintf("▏ %-7s ⇾ [", txt)
	for s := range stack {
		logStr = fmt.Sprint(logStr, " ")
		if s == len(stack)-1 {
			logStr = fmt.Sprint(logStr, "\033[1m", stack[s].Debug(), "\033[0m")
		} else {
			logStr = fmt.Sprint(logStr, stack[s].Debug())
		}
	}
	vmLogger.Println(logStr, " ]")
}

func locateScopeLabel(
	scopes []map[string]uint64, lbl string,
) (
	idx uint64, ok bool,
) {
	topScope := len(scopes) - 1
	for i := topScope; i >= 0; i-- {
		idx, ok = scopes[i][lbl]
		if ok {
			return
		}
	}
	return
}

// func (vm *VirtualMachine) TextMapFromString(s string) word.Word {
// 	var first word.Word
// 	var rn rune
// 	var sz, runeCount int
// 	for si := 0; si < len(s); si += sz {
// 		rn, sz = utf8.DecodeRuneInString(s[si:])
// 		vm.heap = append(vm.heap, word.FromRune(rn))
// 		if si == 0 {
// 			first = word.FromAddress(len(vm.heap) - 1)
// 		}
// 		runeCount += 1
// 	}
// 	last := int(first.AsInt()) + runeCount
// 	legend := TextMapLegend{
// 		BaseMapLegend: BaseMapLegend{
// 			Legend: Legend{
// 				Mark:       word.Word(word.TEXT_LGD),
// 				MetaLegend: word.FromAddress(0),
// 			},
// 			TrashSweep: Empty(),
// 			Length:     word.FromInt(0),
// 			Count:      word.FromInt(0),
// 			ReqLink:    word.FromAddress(0),
// 			DepLink:    word.FromAddress(0),
// 		},
// 		Runes: NewRuneArray(first, vm.heap[first.AsInt():last]...),
// 	}
// }

func (vm *VirtualMachine) AddLiteralToStack(literal word.Word) {
	vm.stack = append(vm.stack, literal)
	logAddedToStack(vm.stack, literal.Debug())
}

// Currently just for lexically traversing the scope
func (vm *VirtualMachine) SubmitLocalRequest(addr word.Word) {
	target := vm.heap[addr.AsAddr()]
	if target.IsRuneArrayMark() {
		label := ReviveRuneArray(vm, addr.AsAddr()).String(vm)
		idx, ok := locateScopeLabel(vm.scope, label)
		if ok {
			// TODO: Invoke address, skip addrRef
			vm.stack = append(vm.stack, word.FromAddress(idx))
			logAddedToStack(vm.stack, label)
		} else {
			vm.scope[len(vm.scope)-1][label] = addr.AsAddr()
			vm.stack = append(vm.stack, addr)
			logAddedToStack(vm.stack, label)
		}
	} else {
		vm.stack = append(vm.stack, target)
		logAddedToStack(vm.stack, target.Debug())
	}
}

// Used for traversing maps and legends
func (vm *VirtualMachine) SubmitInnerRequest(label word.Word) {

}

// Used for traversing submaps and legends
func (vm *VirtualMachine) SubmitUnderRequest(label word.Word) {
}

// Used for traversing primitives and compilations
func (vm *VirtualMachine) SubmitOuterRequest(label word.Word) {
	// FIXME: locate text
	lits := vm.main.ReviveLits(vm)
	addr, err := lits.IndexOf(vm, label)
	if err != nil {
		panic("unable to find word for outer request")
	}
	refId := lits.id + word_arr_offset + uint64(addr)
	refAddr := vm.heap[refId]
	ref := vm.heap[refAddr.AsAddr()]
	if !(ref.IsRuneArrayMark()) {
		panic("outer request submitted with non-ra value")
	}
	text := ReviveRuneArray(vm, refAddr.AsAddr()).String(vm)
	fmt.Println(text)
	switch text {
	case ".=", ":=":
		vm.assignLabel()
	case "++":
		vm.addTwoInts()
	case "--":
		vm.subTwoInts()
	case "**":
		vm.mulTwoInts()
	case "//":
		vm.divTwoInts()
	case "^^":
		vm.expTwoInts()
	case "[[":
		vm.beginMap()
	case "]]":
		vm.endMap()
	case "((":
		vm.beginRoutine()
	case "))":
		vm.endRoutine()

	default:
		panic("Not a valid outer operator")
	}
}

// // Used for signalling across Rhumb
// func (vm *VirtualMachine) SubmitEventRequest(ir InstrRef) {}

// // Used for replying to signals across Rhumb
// func (vm *VirtualMachine) SubmitReplyRequest(ir InstrRef) {}
