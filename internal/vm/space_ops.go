package vm

import (
	"fmt"
)

func (vm *VM) opPost() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Signal Posted")
	return nil
}

func (vm *VM) opInject() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Reply Injected")
	return nil
}

func (vm *VM) opWrite() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Proclamation Written")
	return nil
}

func (vm *VM) opSubscribe() error {
	// Stub
	// Subscribe takes logic?
	fmt.Println("DEBUG: Subscribed")
	return nil
}

func (vm *VM) opNewRealm() {
	// Stub
	// vm.push(Realm)
}
