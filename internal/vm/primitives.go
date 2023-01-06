package vm

import (
	"fmt"
)

/* Phase 1
 * _+/_
 * _-/_
 */

func locateScopeLabel(
	scopes []map[string]int, lbl string,
) (
	idx int, ok bool,
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

func (vm *VirtualMachine) popStack() (popped Word) {
	idx := len(vm.stack) - 1
	popped = vm.stack[idx]
	vm.stack = vm.stack[:idx]
	return
}

func (vm *VirtualMachine) gatherTwoInts() (val1, val2 uint32) {
	stackVal2 := vm.popStack()
	stackVal1 := vm.popStack()
	if stackVal1.IsAddress() {
		val1 = vm.heap[stackVal1.AsAddr()].AsInt()
	} else if stackVal1.IsInteger() {
		val1 = stackVal1.AsInt()
	}
	if stackVal2.IsAddress() {
		val2 = vm.heap[stackVal2.AsAddr()].AsInt()
	} else if stackVal2.IsInteger() {
		val2 = stackVal2.AsInt()
	}
	return
}

func (vm *VirtualMachine) assignLabel() {
	valWord := vm.popStack()
	addrWord := vm.popStack()
	vm.heap[addrWord.AsAddr()] = valWord
}

func (vm *VirtualMachine) addTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.stack = append(vm.stack, WordFromInt(val1+val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " + ", val2))
}

func (vm *VirtualMachine) subTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.stack = append(vm.stack, WordFromInt(val1-val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " - ", val2))
}

func (vm *VirtualMachine) mulTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.stack = append(vm.stack, WordFromInt(val1*val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " x ", val2))
}

func (vm *VirtualMachine) divTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.stack = append(vm.stack, WordFromInt(val1/val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " / ", val2))
}

func (vm *VirtualMachine) expTwoInts() {
	val1, val2 := vm.gatherTwoInts()
	vm.stack = append(vm.stack, WordFromInt(expBySquaring(val1, val2)))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " ^ ", val2))
}

// New scope and add a sentinel to the stack
func (vm *VirtualMachine) beginRoutine() {
	vm.scope = append(vm.scope, make(map[string]int))
	vm.stack = append(vm.stack, Word(SENTINEL))
}

// Same as routine
func (vm *VirtualMachine) beginMap() {
	vm.beginRoutine()
}

func (vm *VirtualMachine) unwindToSentinel() error {
	for back := len(vm.stack); back > 0; back-- {
		if vm.stack[back].IsSentinel() {
			vm.stack = vm.stack[:back-1]
			return nil
		}
	}
	return fmt.Errorf("no sentinel found")
}

// Replace all sub stack values with one final result
func (vm *VirtualMachine) endRoutine() {
	stackLen := len(vm.stack)
	last := vm.stack[stackLen-1]

	vm.scope = vm.scope[:len(vm.scope)-1]

	if err := vm.unwindToSentinel(); err != nil {
		panic(err)
	}

	vm.stack = append(vm.stack, last)
}

// Delete all sub stack values and turn scope into map
func (vm *VirtualMachine) endMap() {
	scopeCount := len(vm.scope)
	// mapScope := vm.scope[scopeCount-1]

	vm.scope = vm.scope[:scopeCount-1]

	if err := vm.unwindToSentinel(); err != nil {
		panic(err)
	}

	// mapLeg := NewMapLegend(mapScope)
	// legAddr := vm.heap[len(vm.heap)-1]
	// vm.heap = append(vm.heap, mapLeg.AsWords())
	// mapObj := NewMapObject(legAddr, mapScope)
	// vm.heap = append(vm.heap)
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
