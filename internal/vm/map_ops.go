package vm

import (
	"fmt"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

func (vm *VM) opMakeMap() {
	m := mapval.NewMap()
	vm.push(mapval.Value{Type: mapval.ValObject, Obj: m})
}

func (vm *VM) opSend() error {
	frame := vm.currentFrame()
	chunk := frame.Closure.Fn.Chunk
	
	idx := chunk.Code[frame.IP]
	frame.IP++
	keyVal := chunk.Constants[idx]
	key := keyVal.Str
	
	receiver := vm.pop()
	if receiver.Type != mapval.ValObject {
		return fmt.Errorf("receiver is not an object")
	}
	
	m, ok := receiver.Obj.(*mapval.Map)
	if ok {
		val, found := m.Get(key)
		if !found {
			vm.push(mapval.NewEmpty())
			return nil
		}
		vm.push(val)
		return nil
	}
	
	// Handle Tuple (as List)
	t, ok := receiver.Obj.(*mapval.Tuple)
	if ok {
		var idx int
		n, _ := fmt.Sscanf(key, "%d", &idx)
		if n == 1 && idx > 0 && idx <= len(t.Payload) {
			vm.push(t.Payload[idx-1])
			return nil
		}
		return fmt.Errorf("tuple index out of bounds or invalid: %s", key)
	}

	return fmt.Errorf("receiver is not a map or tuple")
}

func (vm *VM) opSetField() error {

	frame := vm.currentFrame()

	chunk := frame.Closure.Fn.Chunk

	

	idx := chunk.Code[frame.IP]

	frame.IP++

	flags := chunk.Code[frame.IP]

	frame.IP++

	

	keyVal := chunk.Constants[idx]

	key := keyVal.Str

	

	val := vm.pop()

	receiver := vm.pop()

	

	if receiver.Type != mapval.ValObject {

		return fmt.Errorf("receiver is not an object")

	}

	

	m, ok := receiver.Obj.(*mapval.Map)

	if !ok {

		return fmt.Errorf("receiver is not a map")

	}

	

		mutable := (flags & 1) == 1

	

	

	

		m.Set(key, val, mutable)

	

		

	

		vm.push(val)

	

	

	

		return nil

	

	}

	

	