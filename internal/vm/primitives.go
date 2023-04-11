package vm

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/word"
)

/* Phase 1
 * _+/_
 * _-/_
 */

func expBySquaring(x, n uint32) uint32 {
	if n == 0 {
		return 1
	}
	var y uint32 = 1
	for n > 1 {
		if n%2 == 0 {
			x = x * x
			n = n / 2
		} else {
			y = x * y
			x = x * x
			n = (n - 1) / 2
		}
	}
	return x * y
}

func (vm *VirtualMachine) popStack() (popped word.Word) {
	idx := len(vm.Stack) - 1
	popped = vm.Stack[idx]
	vm.Stack = vm.Stack[:idx]
	return
}

func (vm *VirtualMachine) gatherTwoInts() (val1, val2 uint32) {
	stackVal2 := vm.popStack()
	stackVal1 := vm.popStack()
	if stackVal1.IsAddress() {
		val1 = vm.Heap[stackVal1.AsAddr()].AsInt()
	} else if stackVal1.IsInteger() {
		val1 = stackVal1.AsInt()
	}
	if stackVal2.IsAddress() {
		val2 = vm.Heap[stackVal2.AsAddr()].AsInt()
	} else if stackVal2.IsInteger() {
		val2 = stackVal2.AsInt()
	}
	return
}

func (vm *VirtualMachine) assignScopeLabel(mut bool) {
	valWord := vm.popStack()
	addrWord := vm.popStack()
	vm.Heap[addrWord.AsAddr()] = valWord
}

func (vm *VirtualMachine) assignMapLabel(mut bool) {
	valWord := vm.popStack()
	lblWord := vm.popStack()
	label := ReviveRuneArray(vm, lblWord.AsAddr())
	lastMap := len(vm.MapScope) - 1
	vm.MapScope[lastMap] = vm.MapScope[lastMap].Set(vm, label, valWord)
	logAddedToStack(vm.Stack, fmt.Sprint(label, " set to '", valWord))
}

func (vm *VirtualMachine) addTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.Stack = append(vm.Stack, word.FromInt(val1+val2))
	logAddedToStack(vm.Stack, fmt.Sprint(val1, " + ", val2))
}

func (vm *VirtualMachine) subTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.Stack = append(vm.Stack, word.FromInt(val1-val2))
	logAddedToStack(vm.Stack, fmt.Sprint(val1, " - ", val2))
}

func (vm *VirtualMachine) mulTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.Stack = append(vm.Stack, word.FromInt(val1*val2))
	logAddedToStack(vm.Stack, fmt.Sprint(val1, " x ", val2, "  "))
}

func (vm *VirtualMachine) divTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.Stack = append(vm.Stack, word.FromInt(val1/val2))
	logAddedToStack(vm.Stack, fmt.Sprint(val1, " / ", val2))
}

func (vm *VirtualMachine) expTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.Stack = append(vm.Stack, word.FromInt(expBySquaring(val1, val2)))
	logAddedToStack(vm.Stack, fmt.Sprint(val1, " ^ ", val2))
}

// New scope and add a sentinel to the stack
func (vm *VirtualMachine) beginRoutine() {
	vm.LexScope = append(vm.LexScope, make(map[string]uint64))
	vm.Stack = append(vm.Stack, word.Sentinel())
	logAddedToStack(vm.Stack, "'('")
}

// New CurrentMap
func (vm *VirtualMachine) beginMap() {
	vm.MapScope = append(vm.MapScope, NewListMap(vm))
	vm.Stack = append(vm.Stack, word.Sentinel())
	logAddedToStack(vm.Stack, "*Map")
}

func (vm *VirtualMachine) unwindToSentinel(each ...func(w word.Word)) error {
	for back := len(vm.Stack) - 1; back > 0; back-- {
		currentWord := vm.Stack[back]
		if currentWord.IsSentinel() {
			vm.Stack = vm.Stack[:back]
			return nil
		} else {
			for _, fn := range each {
				fn(currentWord)
			}
		}
	}
	return fmt.Errorf("no sentinel found")
}

// Replace all sub stack values with one final result
func (vm *VirtualMachine) endRoutine() {
	stackLen := len(vm.Stack)
	last := vm.Stack[stackLen-1]

	vm.LexScope = vm.LexScope[:len(vm.LexScope)-1]

	if err := vm.unwindToSentinel(); err != nil {
		panic(err)
	}

	vm.Stack = append(vm.Stack, last)
	logAddedToStack(vm.Stack, last.Debug())
}

// Delete all sub stack values and place map address on stack
func (vm *VirtualMachine) endMap() {
	var (
		last               int             = len(vm.MapScope) - 1
		currMap            Map             = vm.MapScope[last]
		unwoundVals        []word.Word     = make([]word.Word, 0)
		storePositionalVal func(word.Word) = func(w word.Word) {
			unwoundVals = append(unwoundVals, w)
		}
	)

	if last == 0 {
		vm.MapScope = nil
	} else {
		vm.MapScope = vm.MapScope[:last-1]
	}

	if err := vm.unwindToSentinel(storePositionalVal); err != nil {
		panic(err)
	}
	for i, j := 0, len(unwoundVals)-1; i < j; i, j = i+1, j-1 {
		unwoundVals[i], unwoundVals[j] = unwoundVals[j], unwoundVals[i]
	}
	currMap.Append(vm, unwoundVals...)

	vm.Stack = append(vm.Stack, word.FromAddress(currMap.Id))
	logAddedToStack(vm.Stack, fmt.Sprint("Map@", currMap.Id))
}

/* Phase 2
 * _.._
 * _::_s
 * _\_
 * _\\_
 * _@_
 * _@@_
 */

/* Phase 3
 * _==_
 * _!=_
 * _<<_
 * _<=_
 * _>>_
 * _>=_
 * _&&_
 * _||_
 */

/* Phase 4
 * _=>_
 * _!>_
 * _=~_
 * _!~_
 * _=@_
 * _!@_
 * _=^_
 * _~~_
 */

/* Phase 5
 * _<>_
 * _|>_
 * _->_
 * _(_
 * _[_
 */

/* Phase 6
 * _{}_
 * _~>_
 * _<_
 * _<(_
 * _<[_
 * _<{_
 */
