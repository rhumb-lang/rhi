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

func (vm *VirtualMachine) popStack() (popped Word) {
	idx := len(vm.stack) - 1
	popped = vm.stack[idx]
	vm.stack = vm.stack[:idx]
	return
}

func (vm *VirtualMachine) assignLabel() {
	valWord := vm.popStack()
	addrWord := vm.popStack()
	vm.heap[addrWord.AsAddr()] = valWord
}

func (vm *VirtualMachine) addTwoInts() {
	val2 := vm.heap[vm.popStack().AsAddr()].AsInt()
	val1 := vm.heap[vm.popStack().AsAddr()].AsInt()
	vm.stack = append(vm.stack, WordFromInt(val1+val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " + ", val2))
}

func (vm *VirtualMachine) subTwoInts() {
	val2 := vm.heap[vm.popStack().AsAddr()].AsInt()
	val1 := vm.heap[vm.popStack().AsAddr()].AsInt()
	vm.stack = append(vm.stack, WordFromInt(val1-val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " - ", val2))
}

func (vm *VirtualMachine) mulTwoInts() {
	val2 := vm.heap[vm.popStack().AsAddr()].AsInt()
	val1 := vm.heap[vm.popStack().AsAddr()].AsInt()
	vm.stack = append(vm.stack, WordFromInt(val1*val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " x ", val2))
}

func (vm *VirtualMachine) divTwoInts() {
	val2 := vm.heap[vm.popStack().AsAddr()].AsInt()
	val1 := vm.heap[vm.popStack().AsAddr()].AsInt()
	vm.stack = append(vm.stack, WordFromInt(val1/val2))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " / ", val2))
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

func (vm *VirtualMachine) expTwoInts() {
	val2 := vm.heap[vm.popStack().AsAddr()].AsInt()
	val1 := vm.heap[vm.popStack().AsAddr()].AsInt()
	vm.stack = append(vm.stack, WordFromInt(expBySquaring(val1, val2)))
	logAddedToStack(vm.stack, fmt.Sprint(val1, " ^ ", val2))
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
