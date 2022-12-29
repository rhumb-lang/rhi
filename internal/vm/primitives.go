package vm

import "fmt"

/* Phase 1
 * _.=_
 * _:=_
 * _++_
 * _--_
 * _**_
 * _^^_
 * _//_
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
