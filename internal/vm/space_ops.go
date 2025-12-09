package vm

import (
	"fmt"
	"git.sr.ht/~madcapjake/rhi/internal/map"
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
	// sig := sigVal.Obj // Not needed if we push sigVal directly
	
	if vm.Config.TraceSpace {
		fmt.Printf("TRACE: Signal Posted: %s from Frame %p\n", name, frame)
	}
	
	// Synchronous Dispatch (Mocking the Scheduler)
	// 1. Check Monitors in Parent Chain
	curr := vm.CurrentFrame
	for curr != nil {
		if curr.Monitor != nil {
			if vm.Config.TraceSpace {
				fmt.Printf("TRACE: Found Monitor in Frame %p\n", curr)
			}
			
			// Execute Monitor synchronously
			// We MUST satisfy opReturn calling convention: [Function, Args...]
			// Push Monitor Closure (Function)
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: curr.Monitor})
			
			// Push Signal (Argument to Selector)
			vm.push(sigVal)
			
			// Setup Frame for Monitor
			monitorFrame := &CallFrame{
				Parent:  vm.CurrentFrame,
				Closure: curr.Monitor,
				IP:      0,
				Base:    vm.SP - 1, // Points to SigVal. Base-1 is Monitor Closure.
			}
			
			// Run Monitor
			vm.CurrentFrame = monitorFrame
			
			res, err := vm.RunSynchronous()
			
			if err != nil {
				return err
			}
			
			if res == Ok {
				// Monitor returned. Top of stack is the result (Reply Value).
				return nil
			}
		}
		curr = curr.Parent
	}

	// 2. Check Receiver for Listeners (e.g. Realm)
	if receiver.Type == mapval.ValObject {
		if m, ok := receiver.Obj.(*mapval.Map); ok {
			for _, listenerVal := range m.Listeners {
				if listenerVal.Type == mapval.ValObject {
					if closure, ok := listenerVal.Obj.(*mapval.Closure); ok {
						// Execute Listener synchronously
						// Satisfy calling convention
						vm.push(listenerVal) // Push Listener Closure
						vm.push(sigVal)      // Push Arg
						
						newFrame := &CallFrame{
							Parent:  vm.CurrentFrame,
							Closure: closure,
							IP:      0,
							Base:    vm.SP - 1, 
						}
						
						vm.CurrentFrame = newFrame
						res, err := vm.RunSynchronous()
						if err != nil { return err }
						if res == Ok { return nil }
					}
				}
			}
		}
	}
	
	// No listener or no reply? Return Empty.
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
	if argc == 0 {
		payload = mapval.NewEmpty()
	} else if argc == 1 {
		payload = args[0]
	} else {
		// Bundle into Map (List)
		m := mapval.NewMap()
		// Populate Fields and Legend
		m.Fields = args
		
		// Map logic: Positional elements have numeric names ("1", "2", etc.)
		// Always create a new Legend for this specific List structure
		m.Legend = &mapval.Legend{
			Kind: mapval.LegendMap,
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
	// Stub
	// vm.push(Realm)
	if vm.Config.TraceSpace {
		fmt.Println("TRACE: New Realm")
	}
	vm.push(mapval.NewEmpty())
}
