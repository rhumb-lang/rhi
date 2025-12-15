package vm

import (
	"fmt"

	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (vm *VM) opMonitor() error {
	// Stack: [Target, Selector]
	selectorVal := vm.pop()
	targetVal := vm.peek(0) // Peek target

	selector, ok := selectorVal.Obj.(*mapval.Closure)
	if !ok {
		return fmt.Errorf("monitor must be a selector closure")
	}

	// If Target is a Closure, attach as Monitor
	if targetVal.Type == mapval.ValObject {
		if closure, ok := targetVal.Obj.(*mapval.Closure); ok {
			// Create new frame for the target closure
			newFrame := &CallFrame{
				Parent:  vm.CurrentFrame,
				Closure: closure,
				IP:      0,
				Base:    vm.SP, // Base points to first arg (or empty space if 0 args). Base-1 is Target.
				Monitor: selector,
			}

			vm.CurrentFrame = newFrame
			return nil
		}
	}

	// Dispatch Mode: Selector(Target)
	// Target is already on stack (peaked).
	// We need to setup a call to Selector with Target as arg.
	// But opMonitor expects [Target, Selector] on stack (before pops).
	// We popped Selector. Target is at Top.
	// Stack: [..., Target].
	// We want to call Selector(Target).
	// Calling convention: [Closure, Arg].
	// So push Selector (Closure).
	// Swap? [Target, Selector] -> [Selector, Target].
	// Target is at Top.
	// vm.push(selectorVal).
	// Stack: [Target, Selector].
	// Swap.
	// But wait, `opCall` expects [Closure, Args].
	// Stack: [Closure, Arg].
	// Current Stack: [Target].
	// Need: [Selector, Target].
	// So pop Target, push Selector, push Target.

	vm.pop() // Pop Target
	vm.push(selectorVal)
	vm.push(targetVal)

	// Setup Call Frame
	newFrame := &CallFrame{
		Parent:  vm.CurrentFrame,
		Closure: selector,
		IP:      0,
		Base:    vm.SP - 1, // 1 Arg (Target)
	}

	vm.CurrentFrame = newFrame
	return nil
}

func (vm *VM) opPost() error {
	idx := vm.readByte()
	argc := int(vm.readByte())

	args := make([]mapval.Value, argc)
	for i := argc - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	receiver := vm.pop()

	// Get Signal Name
	frame := vm.currentFrame()
	name := frame.Closure.Fn.Chunk.Constants[idx].Str

	// Create Signal
	sigVal := mapval.NewSignal(name, frame, args)

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: Signal Posted: #%s to %s\n", name, receiver)
	}

	// PHASE 1: Check Receiver Listeners (The Realm itself)
	if receiver.Type == mapval.ValObject {
		if m, ok := receiver.Obj.(*mapval.Map); ok {
			for _, listenerVal := range m.Listeners {
				if listenerVal.Type == mapval.ValObject {
					if closure, ok := listenerVal.Obj.(*mapval.Closure); ok {

						// --- Execute Listener (Level 1) ---
						vm.push(listenerVal)
						vm.push(sigVal)

						newFrame := &CallFrame{
							Parent:  vm.CurrentFrame,
							Closure: closure,
							IP:      0,
							Base:    vm.SP - 1,
						}

						vm.CurrentFrame = newFrame
						res, err := vm.RunSynchronous()
						if err != nil {
							return err
						}
						if res != Ok {
							continue
						} // Should not happen usually

						// --- Check Result ---
						result := vm.peek(0)

						// CHAINING: If result is a Closure (Selector), execute it!
						if result.Type == mapval.ValObject {
							if nextClosure, ok := result.Obj.(*mapval.Closure); ok {
								if vm.Config.TraceSpace {
									fmt.Println("TRACE: Listener returned Closure -> Executing Chain")
								}
								vm.pop() // Pop the closure (result of L1)

								// Setup Call for L2 (Selector)
								vm.push(result) // Push Selector
								vm.push(sigVal) // Push Signal

								chainFrame := &CallFrame{
									Parent:  vm.CurrentFrame,
									Closure: nextClosure,
									IP:      0,
									Base:    vm.SP - 1,
								}

								vm.CurrentFrame = chainFrame
								_, err = vm.RunSynchronous()
								if err != nil {
									return err
								}

								// Update result to the output of the chain
								result = vm.peek(0)
							}
						}

						// Final Decision: Consume or Bubble?
						if result.Type != mapval.ValEmpty {
							if vm.Config.TraceSpace {
								fmt.Printf("TRACE: Signal Consumed by Receiver: %s\n", result)
							}
							return nil // Consumed
						}

						// Empty result = Peek (::) or Fallthrough -> Continue Bubbling
						vm.pop()
					}
				}
			}
		}
	}

	// PHASE 2: Bubble Up (Check Monitors in Parent Chain)
	curr := vm.CurrentFrame
	for curr != nil {
		if curr.Monitor != nil {
			if vm.Config.TraceSpace {
				fmt.Printf("TRACE: Bubbled to Monitor in Frame %p\n", curr)
			}

			vm.push(mapval.Value{Type: mapval.ValObject, Obj: curr.Monitor})
			vm.push(sigVal)

			monitorFrame := &CallFrame{
				Parent:  vm.CurrentFrame,
				Closure: curr.Monitor,
				IP:      0,
				Base:    vm.SP - 1,
			}

			vm.CurrentFrame = monitorFrame
			res, err := vm.RunSynchronous()

			if err != nil {
				return err
			}

			if res == Ok {
				result := vm.peek(0)
				if result.Type != mapval.ValEmpty {
					if vm.Config.TraceSpace {
						fmt.Printf("TRACE: Signal Consumed by Monitor: %s\n", result)
					}
					return nil
				}
				vm.pop()
			}
		}
		curr = curr.Parent
	}

	// PHASE 3: Unhandled
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: Signal Unhandled (Dropped)")
	}
	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opInject() error {
	// opInject is called inside the listener.
	// It should "resume" the sender.
	// In our sync model, it just returns the value.

	idx := vm.readByte()
	argc := int(vm.readByte())

	if vm.Config.TraceSpace {
		fmt.Printf("DEBUG: opInject argc=%d SP=%d\n", argc, vm.SP)
		for i := 0; i < vm.SP; i++ {
			fmt.Printf("Stack[%d]: %s\n", i, vm.Stack[i])
		}
	}

	args := make([]mapval.Value, argc)
	for i := argc - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	receiver := vm.pop() // The Signal object

	if vm.Config.TraceSpace {
		frame := vm.currentFrame()
		name := frame.Closure.Fn.Chunk.Constants[idx].Str
		fmt.Printf("TRACE: Reply %s Injected to %s (Type: %d)\n", name, receiver, receiver.Type)
	}

	// Determine Payload
	var payload mapval.Value
	switch argc {
	case 0:
		payload = mapval.NewEmpty()
	case 1:
		payload = args[0]
	default:
		// Bundle into Map (List)
		m := mapval.NewMap()
		// Populate Fields and Legend
		m.Fields = args

		// Map logic: Positional elements have numeric names ("1", "2", etc.)
		// Always create a new Legend for this specific List structure
		m.Legend = &mapval.Legend{
			Kind:   mapval.LegendMap,
			Fields: make([]mapval.FieldDesc, len(args)),
		}

		for i := 0; i < len(args); i++ {
			// Name is "1"-based index string
			name := fmt.Sprintf("%d", i+1)
			m.Legend.Fields[i] = mapval.FieldDesc{
				Name: name,
				Kind: mapval.FieldMutable, // or Immutable? Args usually mutable in list?
			}
		}

		payload = mapval.Value{Type: mapval.ValObject, Obj: m}
	}

	vm.push(payload)
	return nil
}

func (vm *VM) opWrite() error {
	// Stub
	vm.readByte() // Consume index
	argc := int(vm.readByte())
	for range argc {
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
	// Stack: [Map, Selector]
	selector := vm.pop()
	receiver := vm.pop()

	if receiver.Type == mapval.ValObject {
		if m, ok := receiver.Obj.(*mapval.Map); ok {
			m.Listeners = append(m.Listeners, selector)
			if vm.Config.TraceSpace {
				fmt.Printf("TRACE: Subscribed listener to map\n")
			}
		} else {
			return fmt.Errorf("can only subscribe to maps")
		}
	} else {
		return fmt.Errorf("can only subscribe to maps")
	}

	vm.push(mapval.NewEmpty())
	return nil
}

func (vm *VM) opNewRealm() {
	// Consume the flags byte emitted by the compiler!
	_ = vm.readByte()

	// vm.push(Realm)
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: New Realm")
	}
	// Realm is just a Map for now (with space semantics handled by operators)
	vm.push(mapval.Value{Type: mapval.ValObject, Obj: mapval.NewMap()})
}
