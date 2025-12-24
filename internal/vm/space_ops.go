package vm

import (
	"fmt"
	"os"

	mapval "github.com/rhumb-lang/rhi/internal/map"
)

func (vm *VM) raiseSignal(topic string, signalVal mapval.Value) {
	// Logic matches opPost: Suspend current frame, check monitor, or bubble.

	currentFrame := vm.CurrentFrame
	currentFrame.WaitingSignal = topic

	// Move to Zombies
	vm.zombies = append(vm.zombies, currentFrame)

	// CHECK MONITOR ON CURRENT FRAME
	if currentFrame.Monitor != nil {
		// Push Monitor Closure & Signal
		vm.push(mapval.Value{Type: mapval.ValObject, Obj: currentFrame.Monitor})
		vm.push(signalVal)

		// Debug
		if vm.Config.TraceSpace {
			fmt.Fprintf(os.Stderr, "Raising Signal to Monitor: %s (Type %d)\n", signalVal, signalVal.Type)
			fmt.Fprintf(os.Stderr, "Stack at Raise: %v\n", vm.Stack[:vm.SP])
		}

		// Start Monitor Frame

		monitorFrame := &CallFrame{
			Parent:  currentFrame.Parent,
			Closure: currentFrame.Monitor,
			IP:      0,
			Base:    vm.SP - 1, // 1 Arg (Signal)
		}

		vm.CurrentFrame = monitorFrame
		vm.state = StateRunning // Handle immediately
		return
	}

	// No Monitor on current frame, bubble up from Parent
	// STACK UNWINDING: Restore SP to discard locals/args of the suspended frame
	targetSP := currentFrame.Base
	if targetSP > 0 {
		targetSP-- // Pop the Closure at Base-1
	}
	vm.SP = targetSP

	vm.CurrentFrame = currentFrame.Parent
	vm.signalVal = signalVal
	vm.state = StateSignaling
}

func (vm *VM) opMonitor() error {
	argc := int(vm.readByte())

	selectorVal := vm.pop()
	selector, ok := selectorVal.Obj.(*mapval.Closure)
	if !ok {
		return fmt.Errorf("monitor must be a selector closure")
	}

	if argc > 0 {
		// Monitored Call Mode
		// Stack: [Func, Arg1...ArgN]. Top is SP.
		// Func is at SP - argc - 1.
		funcVal := vm.Stack[vm.SP-argc-1]

		if funcVal.Type == mapval.ValObject {
			if closure, ok := funcVal.Obj.(*mapval.Closure); ok {
				newFrame := &CallFrame{
					Parent:  vm.CurrentFrame,
					Closure: closure,
					IP:      0,
					Base:    vm.SP - argc,
					Monitor: selector,
				}
				vm.CurrentFrame = newFrame
				return nil
			}
		}
		return fmt.Errorf("monitored target must be a closure")
	}

	// Dispatch Mode / Vassal Mode (argc == 0)
	// Stack: [Target]. Selector popped.
	targetVal := vm.peek(0)

	// If Target is a Closure, attach as Monitor (Vassal Mode)
	if targetVal.Type == mapval.ValObject {
		if closure, ok := targetVal.Obj.(*mapval.Closure); ok {
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

func (vm *VM) handleSignal() (bool, error) {
	// Bubbling Logic for Async Signals
	// We search up the stack (vm.CurrentFrame -> Parent -> ...)
	// for a frame with a Monitor (Selector) that can handle this signal.

	curr := vm.CurrentFrame
	for curr != nil {
		if curr.Monitor != nil {
			if vm.Config.TraceSpace {
				fmt.Printf("TRACE: Bubbled to Monitor in Frame %p\n", curr)
			}

			// Found Monitor
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: curr.Monitor})
			vm.push(vm.signalVal)

			// Create new Frame for Monitor
			monitorFrame := &CallFrame{
				Parent:  vm.CurrentFrame,
				Closure: curr.Monitor,
				IP:      0,
				Base:    vm.SP - 1, // 1 Arg (Signal)
			}

			vm.CurrentFrame = monitorFrame
			return true, nil
		}
		curr = curr.Parent
	}

	return false, nil
}

func (vm *VM) opPost() error {
	// Corresponds to OP_SIGNAL / Suspend
	idx := vm.readByte()
	argc := int(vm.readByte())

	// Pop Args
	args := make([]mapval.Value, argc)
	for i := argc - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	// Pop Receiver (Implicitly pushed by compiler, ignored for now in global signal model)
	vm.pop()

	// Get Signal Name
	frame := vm.currentFrame()
	name := frame.Closure.Fn.Chunk.Constants[idx].Str

	// Create Signal Value
	vm.signalVal = mapval.NewSignal(name, frame, args)

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: Signal Posted (Suspend): #%s\n", name)
	}

	// Suspend the Current Frame (The "Zombie" Logic)
	currentFrame := vm.CurrentFrame
	currentFrame.WaitingSignal = name

	// Move to Zombies
	vm.zombies = append(vm.zombies, currentFrame)

	// CHECK MONITOR ON CURRENT FRAME (The one being suspended)
	if currentFrame.Monitor != nil {
		// Push Monitor Closure & Signal
		vm.push(mapval.Value{Type: mapval.ValObject, Obj: currentFrame.Monitor})
		vm.push(vm.signalVal)

		// Start Monitor Frame
		monitorFrame := &CallFrame{
			Parent:  currentFrame.Parent, // Handler is sibling of suspended frame (child of caller)
			Closure: currentFrame.Monitor,
			IP:      0,
			Base:    vm.SP - 1, // 1 Arg (Signal)
		}

		vm.CurrentFrame = monitorFrame
		vm.state = StateRunning // Handle immediately
		return nil
	}

	// No Monitor on current frame, bubble up from Parent
	// Update Stack (Remove current from active chain)
	vm.CurrentFrame = currentFrame.Parent

	// Trigger Signal Propagation
	vm.state = StateSignaling

	return nil
}

func (vm *VM) opInject() error {
	// Corresponds to OP_INJECT / Resume / Reply
	idx := vm.readByte()
	argc := int(vm.readByte())

	// Pop Payload (Args)
	args := make([]mapval.Value, argc)
	for i := argc - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	// Pop Receiver (The Signal we are replying to)
	receiver := vm.pop()

	// Resolve Topic
	frame := vm.currentFrame()
	targetTopic := frame.Closure.Fn.Chunk.Constants[idx].Str

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: Injecting Reply ^%s\n", targetTopic)
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
		m.Fields = args
		m.Legend = &mapval.Legend{
			Kind:   mapval.LegendMap,
			Fields: make([]mapval.FieldDesc, len(args)),
		}
		for i := 0; i < len(args); i++ {
			name := fmt.Sprintf("%d", i+1)
			m.Legend.Fields[i] = mapval.FieldDesc{
				Name: name,
				Kind: mapval.FieldMutable,
			}
		}
		payload = mapval.Value{Type: mapval.ValObject, Obj: m}
	}

	// Hunt for Zombie
	var zombie *CallFrame

	// Strategy 1: Direct Reply via Signal Source
	if receiver.Type == mapval.ValObject {
		if t, ok := receiver.Obj.(*mapval.Tuple); ok && t.Kind == mapval.TupleSignal {
			if sourceFrame, ok := t.Source.(*CallFrame); ok {
				// Verify this frame is actually a zombie (in the list)
				// This is O(N) but safe. Optimization: Flag frames as zombies?
				for i, z := range vm.zombies {
					if z == sourceFrame {
						zombie = z
						// Remove from list
						vm.zombies = append(vm.zombies[:i], vm.zombies[i+1:]...)
						break
					}
				}
			}
		}
	}

	// Strategy 2: Topic Scan (Fallback/Legacy)
	if zombie == nil {
		zombieIdx := -1
		for i := len(vm.zombies) - 1; i >= 0; i-- {
			if targetTopic == "" || vm.zombies[i].WaitingSignal == targetTopic {
				zombie = vm.zombies[i]
				zombieIdx = i
				break
			}
		}
		if zombie != nil {
			vm.zombies = append(vm.zombies[:zombieIdx], vm.zombies[zombieIdx+1:]...)
		}
	}

	if zombie == nil {
		if vm.Config.TraceSpace {
			fmt.Printf("TRACE: No zombie found for ^%s\n", targetTopic)
		}
		return nil // Silent ignore?
	}

	zombie.WaitingSignal = ""

	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: Reviving Zombie Frame %p\n", zombie)
	}

	// Push Payload to Global Stack (Where Zombie will pick it up)
	vm.push(payload)

	// Attach Zombie to Current Execution
	// The Zombie becomes the active frame, on top of the current frame (Handler).
	// We do NOT reparent it, so it returns to its original caller.
	vm.CurrentFrame = zombie

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
