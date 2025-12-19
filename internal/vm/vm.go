package vm

import (
	"fmt"

	"github.com/rhumb-lang/rhi/internal/config"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

const StackMax = 2048
const MaxFrames = 64 // Kept for legacy limits if needed, but not used for storage

// Loader defines the interface for loading libraries.
type Loader interface {
	Load(resolver, logicalPath string, constraint mapval.Value) (mapval.Value, error)
}

type State int

const (
	StateRunning State = iota
	StateSignaling
)

type VM struct {
	CurrentFrame *CallFrame

	Stack [StackMax]mapval.Value
	SP    int // Stack Pointer (points to empty slot)

	Config *config.Config

	Loader Loader

	// State of the VM (Running or Signaling)
	state State

	// The signal being bubbled up
	signalVal mapval.Value

	// Zombies: The "Parking Lot" for suspended frames.
	zombies []*CallFrame
}

// Result code for the VM interpretation
type Result int

const (
	Ok Result = iota
	CompileError
	RuntimeError
	Halt // Add Halt
)

func (r Result) String() string {
	switch r {
	case Ok:
		return "Ok"
	case CompileError:
		return "CompileError"
	case RuntimeError:
		return "RuntimeError"
	case Halt:
		return "Halt"
	default:
		return fmt.Sprintf("Unknown(%d)", r)
	}
}

func NewVM() *VM {
	return &VM{
		SP:           0,
		CurrentFrame: nil,
		Config:       config.DefaultConfig(),
		state:        StateRunning,
	}
}

func (vm *VM) currentFrame() *CallFrame {
	return vm.CurrentFrame
}

// Interpret executes the chunk.
func (vm *VM) Interpret(chunk *mapval.Chunk) (Result, error) {
	fn := &mapval.Function{
		Name:  "<script>",
		Chunk: chunk,
	}
	closure := &mapval.Closure{Fn: fn}

	vm.CurrentFrame = &CallFrame{
		Closure: closure,
		IP:      0,
		Base:    0,
		Parent:  nil,
	}
	vm.state = StateRunning

	return vm.run()
}

// CallAndReturn executes a standalone chunk/closure (used for libraries).
// It creates a temporary frame, runs until return, and gives back the value.
func (vm *VM) CallAndReturn(chunk *mapval.Chunk) (mapval.Value, error) {
	fn := &mapval.Function{
		Name:  "<library>",
		Chunk: chunk,
	}
	closure := &mapval.Closure{Fn: fn}

	// Push closure to stack so OP_RETURN can pop it (Base-1)
	vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})

	vm.CurrentFrame = &CallFrame{
		Closure: closure,
		IP:      0,
		Base:    vm.SP, // Locals start here (after closure)
		Parent:  vm.CurrentFrame,
	}

	res, err := vm.RunSynchronous()
	if err != nil {
		return mapval.NewEmpty(), err
	}
	if res != Ok && res != Halt {
		return mapval.NewEmpty(), fmt.Errorf("library execution failed: %s", res)
	}

	// result is on stack
	if vm.SP == 0 {
		return mapval.NewEmpty(), nil
	}
	return vm.pop(), nil
}

// Continue resumes execution from the given offset in the script frame.
func (vm *VM) Continue(offset int) (Result, error) {
	if vm.CurrentFrame == nil {
		return RuntimeError, fmt.Errorf("no active frame to continue")
	}
	vm.CurrentFrame.IP = offset
	return vm.run()
}

func (vm *VM) push(val mapval.Value) {
	if vm.SP >= StackMax {
		panic("Stack overflow")
	}
	vm.Stack[vm.SP] = val
	vm.SP++
}

func (vm *VM) pop() mapval.Value {
	if vm.SP == 0 {
		panic("Stack underflow")
	}
	vm.SP--
	return vm.Stack[vm.SP]
}

func (vm *VM) peek(distance int) mapval.Value {
	return vm.Stack[vm.SP-1-distance]
}

func (vm *VM) readShort() int {
	frame := vm.currentFrame()
	chunk := frame.Closure.Fn.Chunk
	byte1 := chunk.Code[frame.IP]
	frame.IP++
	byte2 := chunk.Code[frame.IP]
	frame.IP++
	val := uint16(byte1)<<8 | uint16(byte2)
	return int(int16(val))
}

func (vm *VM) readByte() byte {
	frame := vm.currentFrame()
	b := frame.Closure.Fn.Chunk.Code[frame.IP]
	frame.IP++
	return b
}

func (vm *VM) RunSynchronous() (Result, error) {
	startDepth := 0
	// Count frames
	for f := vm.CurrentFrame; f != nil; f = f.Parent {
		startDepth++
	}

	for {
		if vm.state == StateSignaling {
			handled, err := vm.handleSignal()
			if err != nil {
				return RuntimeError, err
			}
			if handled {
				vm.state = StateRunning
			} else {
				vm.state = StateRunning
				// Drop signal
			}
			continue
		}

		res, err := vm.Step()
		if err != nil {
			return RuntimeError, err
		}
		if res != Ok {
			return res, nil
		}

		// Check if we popped the frame we started with
		currentDepth := 0
		for f := vm.CurrentFrame; f != nil; f = f.Parent {
			currentDepth++
		}

		if currentDepth < startDepth {
			// Returned from the frame
			return Ok, nil
		}
	}
}

func (vm *VM) run() (Result, error) {
	for {
		if vm.state == StateSignaling {
			handled, err := vm.handleSignal()
			if err != nil {
				return RuntimeError, err
			}
			if handled {
				vm.state = StateRunning
			} else {
				// Signal Unhandled.
				// For now, we drop it.
				// In a real system, we might want to panic or log.
				vm.state = StateRunning
			}
			continue
		}

		res, err := vm.Step()
		if err != nil {
			return RuntimeError, err
		}
		if res != Ok {
			return res, nil
		}
	}
}

func (vm *VM) Step() (Result, error) {

	frame := vm.currentFrame()

	chunk := frame.Closure.Fn.Chunk

	if frame.IP >= len(chunk.Code) {

		return RuntimeError, fmt.Errorf("IP out of bounds")

	}

	if vm.Config.TraceFrame {
		fmt.Printf("FRAME: %s IP:%04d SP:%d Base:%d", frame.Closure.Fn.Name, frame.IP, vm.SP, frame.Base)
		if frame.Monitor != nil {
			fmt.Print(" [Monitored]")
		}
		if frame.WaitingSignal != "" {
			fmt.Printf(" [Waiting: %s]", frame.WaitingSignal)
		}
		fmt.Println()
	}

	if vm.Config.TraceStack {

		fmt.Print("          ")

		for i := 0; i < vm.SP; i++ {

			fmt.Printf("[ %s ]", vm.Stack[i])

		}

		fmt.Println()

		fmt.Printf("%04d %s\n", frame.IP, mapval.OpCode(chunk.Code[frame.IP]))

	}

	instruction := mapval.OpCode(chunk.Code[frame.IP])

	frame.IP++

	var err error

	switch instruction {

	case mapval.OP_HALT:
		return Halt, nil

		// Stack

	case mapval.OP_LOAD_CONST:
		vm.opLoadConst()

	case mapval.OP_LOAD_LOC:
		vm.opLoadLoc()

	case mapval.OP_STORE_LOC:
		vm.opStoreLoc()

	case mapval.OP_STORE_LOC_IMMUT:
		vm.opStoreLocImmut()

	case mapval.OP_LOAD_UPVALUE:
		vm.opLoadUpvalue()

	case mapval.OP_STORE_UPVALUE:
		vm.opStoreUpvalue()

	case mapval.OP_DUP:
		vm.opDup()

	case mapval.OP_POP:
		vm.opPop()

	// Flow

	case mapval.OP_JUMP:
		vm.opJump()

	case mapval.OP_JUMP_IF_FALSE:
		vm.opIfTrue() // Jump if False (rename func later)

	case mapval.OP_JUMP_IF_TRUE:
		vm.opIfFalse() // Jump if True

	case mapval.OP_CALL:
		err = vm.opCall()

	case mapval.OP_RETURN:

		res, e := vm.opReturn()

		if res == 1 {
			return Halt, nil
		}

		if e != nil {
			err = e
		}

	case mapval.OP_MAKE_FN:
		vm.opMakeFn()

	// Map

	case mapval.OP_MAKE_MAP:
		vm.opMakeMap()

	case mapval.OP_LOAD_STATIC:
		// Not implemented yet
		return RuntimeError, fmt.Errorf("OP_LOAD_STATIC not implemented")

	case mapval.OP_MATCH_BIND:
		// Not implemented yet
		return RuntimeError, fmt.Errorf("OP_MATCH_BIND not implemented")

	case mapval.OP_RESOLVE:
		err = vm.opResolve()

	// --- BANK 2: Map & Inheritance ---

	case mapval.OP_SEND:
		err = vm.opSend()

	case mapval.OP_SET_FIELD:
		err = vm.opSetField()

	// Space

	case mapval.OP_POST:
		err = vm.opPost()

	case mapval.OP_INJECT:
		err = vm.opInject()

	case mapval.OP_WRITE:
		err = vm.opWrite()

	case mapval.OP_SUBSCRIBE:
		err = vm.opSubscribe()

	case mapval.OP_NEW_REALM:
		vm.opNewRealm()

	case mapval.OP_MONITOR:
		err = vm.opMonitor()

	// Math

	case mapval.OP_ADD:
		err = vm.opAdd()

	case mapval.OP_SUB:
		err = vm.opSub()

	case mapval.OP_MULT:
		err = vm.opMult()

	case mapval.OP_DIV_FLOAT:
		err = vm.opDivFloat()

	case mapval.OP_DIV_INT:
		err = vm.opDivInt()

	case mapval.OP_MOD:
		err = vm.opMod()

	case mapval.OP_POW:
		err = vm.opPow()

	case mapval.OP_ROOT:
		err = vm.opRoot()

	case mapval.OP_SCI_NOT:
		err = vm.opSciNot()

	case mapval.OP_DEV:
		err = vm.opDev()

	case mapval.OP_COERCE_NUM:
		err = vm.opCoerceNum()

	case mapval.OP_NUM_NEG:
		err = vm.opNumNeg()

	// Structure

	case mapval.OP_COALESCE:
		vm.opCoalesce()

	case mapval.OP_CONCAT:
		vm.opConcat()

	case mapval.OP_RANGE:
		vm.opRange()

	case mapval.OP_HAS_SUBFIELD:
		vm.opHasSub()

	case mapval.OP_ACCESS_NESTED:
		vm.opAccessNested()

	case mapval.OP_PIPE:
		vm.opPipe()

	case mapval.OP_FOREACH:
		vm.opForeach()

	// Logic

	case mapval.OP_AND:
		vm.opAnd()

	case mapval.OP_OR:
		vm.opOr()

	case mapval.OP_NOT:
		vm.opNot()

	case mapval.OP_EQ:
		vm.opEq()

	case mapval.OP_NEQ:
		vm.opNeq()

	case mapval.OP_GT:
		err = vm.opGt()

	case mapval.OP_LT:
		err = vm.opLt()

	case mapval.OP_GTE:
		err = vm.opGte()

	case mapval.OP_LTE:
		err = vm.opLte()

	// Testing

	case mapval.OP_ASSERT_EQ:
		vm.opAssertEq()

	case mapval.OP_INSPECT:
		vm.opInspect()

		// Selectors

	case mapval.OP_SELECT:
		vm.opSelect() // Not implemented?

	case mapval.OP_MATCH_STRUCT:
		vm.opMatchStruct()

	case mapval.OP_MATCH_TUPLE:
		vm.opMatchTuple()

	default:

		return RuntimeError, fmt.Errorf("unknown opcode: %d", instruction)

	}

	if err != nil {

		return RuntimeError, err

	}

	return Ok, nil

}

func (vm *VM) DumpStack() {
	fmt.Printf("STACK (SP=%d): %v\n", vm.SP, vm.Stack[:vm.SP])
}
