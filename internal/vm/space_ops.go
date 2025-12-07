package vm

import (
	"fmt"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) opPost() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Signal Posted")
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opInject() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Reply Injected")
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opWrite() error {
	// Stub
	vm.readByte() // Consume index
	vm.pop()      // Pop receiver
	fmt.Println("DEBUG: Proclamation Written")
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opSubscribe() error {
	// Stub
	// Subscribe takes logic?
	fmt.Println("DEBUG: Subscribed")
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opNewRealm() {
	// Stub
	// vm.push(Realm)
	fmt.Println("DEBUG: New Realm")
	vm.push(mapval.NewEmpty())
}
