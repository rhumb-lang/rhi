package vm

import (
	"fmt"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) opPost() error {
	// Stub
	vm.readByte() // Consume index
	argc := int(vm.readByte())
	for i := 0; i < argc; i++ {
		vm.pop() // Pop arg
	}
	vm.pop() // Pop receiver
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: Signal Posted")
	}
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opInject() error {
	// Stub
	vm.readByte() // Consume index
	argc := int(vm.readByte())
	for i := 0; i < argc; i++ {
		vm.pop() // Pop arg
	}
	vm.pop() // Pop receiver
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: Reply Injected")
	}
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opWrite() error {
	// Stub
	vm.readByte() // Consume index
	argc := int(vm.readByte())
	for i := 0; i < argc; i++ {
		vm.pop() // Pop arg
	}
	vm.pop() // Pop receiver
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: Proclamation Written")
	}
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opSubscribe() error {
	// Stub
	// Subscribe takes logic?
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: Subscribed")
	}
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opNewRealm() {
	// Stub
	// vm.push(Realm)
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: New Realm")
	}
	vm.push(mapval.NewEmpty())
}
