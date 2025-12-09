package vm

import (
	"fmt"
	"math"
	"git.sr.ht/~madcapjake/rhi/internal/map"
)

// --- Stack Ops ---

func (vm *VM) opLoadConst() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	constant := frame.Closure.Fn.Chunk.Constants[idx]
	vm.push(constant)
}

func (vm *VM) opLoadLoc() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	val := vm.Stack[frame.Base+int(idx)]
	vm.push(val)
}

func (vm *VM) opStoreLoc() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	val := vm.peek(0)
	vm.Stack[frame.Base+int(idx)] = val
}

func (vm *VM) opDup() {
	val := vm.peek(0)
	vm.push(val)
}

func (vm *VM) opPop() {
	vm.pop()
}

func (vm *VM) opLoadUpvalue() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	closure := frame.Closure
	
	if int(idx) >= len(closure.Upvalues) {
		panic(fmt.Sprintf("upvalue index %d out of bounds (len %d)", idx, len(closure.Upvalues)))
	}
	
	upvalue := closure.Upvalues[idx]
	if upvalue.Location != nil {
		vm.push(*upvalue.Location)
	} else {
		vm.push(upvalue.Closed)
	}
}

func (vm *VM) opStoreUpvalue() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	closure := frame.Closure
	
	if int(idx) >= len(closure.Upvalues) {
		panic(fmt.Sprintf("upvalue index %d out of bounds (len %d)", idx, len(closure.Upvalues)))
	}
	
	val := vm.peek(0)
	upvalue := closure.Upvalues[idx]
	
	if upvalue.Location != nil {
		*upvalue.Location = val
	} else {
		upvalue.Closed = val
	}
}

// --- Math Ops ---

func (vm *VM) opAdd() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		vm.push(mapval.NewInt(a.Integer + b.Integer))
	} else if a.Type == mapval.ValFloat || b.Type == mapval.ValFloat {
		fa := asFloat(a)
		fb := asFloat(b)
		vm.push(mapval.NewFloat(fa + fb))
	} else if a.Type == mapval.ValObject && b.Type == mapval.ValObject {
		mapA, okA := a.Obj.(*mapval.Map)
		mapB, okB := b.Obj.(*mapval.Map)
		if okA && okB {
			newMap := mapval.NewMap()
			newMap.Fields = append(newMap.Fields, mapA.Fields...)
			newMap.Fields = append(newMap.Fields, mapB.Fields...)
			vm.push(mapval.Value{Type: mapval.ValObject, Obj: newMap})
		} else {
			return fmt.Errorf("operands must be numbers or maps for ADD")
		}
	} else {
		return fmt.Errorf("operands must be numbers or maps for ADD")
	}
	return nil
}

func (vm *VM) opSub() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		vm.push(mapval.NewInt(a.Integer - b.Integer))
	} else {
		vm.push(mapval.NewFloat(asFloat(a) - asFloat(b)))
	}
	return nil
}

func (vm *VM) opMult() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		vm.push(mapval.NewInt(a.Integer * b.Integer))
	} else {
		vm.push(mapval.NewFloat(asFloat(a) * asFloat(b)))
	}
	return nil
}

func (vm *VM) opDivFloat() error {
	b := vm.pop()
	a := vm.pop()
	vm.push(mapval.NewFloat(asFloat(a) / asFloat(b)))
	return nil
}

func (vm *VM) opDivInt() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		if b.Integer == 0 { return fmt.Errorf("division by zero") }
		vm.push(mapval.NewInt(a.Integer / b.Integer))
	} else {
		vm.push(mapval.NewInt(int64(asFloat(a) / asFloat(b))))
	}
	return nil
}

func (vm *VM) opMod() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		if b.Integer == 0 { return fmt.Errorf("division by zero") }
		vm.push(mapval.NewInt(a.Integer % b.Integer))
	} else {
		vm.push(mapval.NewFloat(math.Mod(asFloat(a), asFloat(b))))
	}
	return nil
}

func (vm *VM) opPow() error {
	b := vm.pop()
	a := vm.pop()
	vm.push(mapval.NewFloat(math.Pow(asFloat(a), asFloat(b))))
	return nil
}

func (vm *VM) opRoot() error {
	b := vm.pop() // y (root)
	a := vm.pop() // x (base)
	// x^(1/y)
	vm.push(mapval.NewFloat(math.Pow(asFloat(a), 1.0/asFloat(b))))
	return nil
}

func (vm *VM) opSciNot() error {
	b := vm.pop() // exponent
	a := vm.pop() // base
	// a * 10^b
	vm.push(mapval.NewFloat(asFloat(a) * math.Pow(10, asFloat(b))))
	return nil
}

func (vm *VM) opDev() error {
	// Deviation (+-)
	// Usually returns a range or deviation object?
	// For now, return float addition? Or not implemented?
	// Placeholder.
	vm.pop()
	vm.pop()
	vm.push(mapval.NewEmpty())
	return nil
}

// --- Logic Ops ---

func (vm *VM) opAnd() {
	b := vm.pop()
	a := vm.pop()
	res := vm.isTruthy(a) && vm.isTruthy(b)
	vm.push(mapval.NewBoolean(res))
}

func (vm *VM) opOr() {
	b := vm.pop()
	a := vm.pop()
	res := vm.isTruthy(a) || vm.isTruthy(b)
	vm.push(mapval.NewBoolean(res))
}

func (vm *VM) opNot() {
	val := vm.pop()
	vm.push(mapval.NewBoolean(vm.isFalsy(val)))
}

// --- Comparison Ops ---

func (vm *VM) opEq() {
	b := vm.pop()
	a := vm.pop()
	vm.push(mapval.NewBoolean(isEqual(a, b)))
}

func (vm *VM) opNeq() {
	b := vm.pop()
	a := vm.pop()
	vm.push(mapval.NewBoolean(!isEqual(a, b)))
}

func (vm *VM) opGt() error {
	b := vm.pop()
	a := vm.pop()
	res, err := numericCompare(a, b)
	if err != nil {
		return err
	}
	vm.push(mapval.NewBoolean(res > 0))
	return nil
}

func (vm *VM) opLt() error {
	b := vm.pop()
	a := vm.pop()
	res, err := numericCompare(a, b)
	if err != nil {
		return err
	}
	vm.push(mapval.NewBoolean(res < 0))
	return nil
}

func (vm *VM) opGte() error {
	b := vm.pop()
	a := vm.pop()
	res, err := numericCompare(a, b)
	if err != nil {
		return err
	}
	vm.push(mapval.NewBoolean(res >= 0))
	return nil
}

func (vm *VM) opLte() error {
	b := vm.pop()
	a := vm.pop()
	res, err := numericCompare(a, b)
	if err != nil {
		return err
	}
	vm.push(mapval.NewBoolean(res <= 0))
	return nil
}

// --- Flow Ops ---

func (vm *VM) opJump() {
	offset := vm.readShort()
	frame := vm.currentFrame()
	frame.IP += offset
}

func (vm *VM) opIfTrue() {
	offset := vm.readShort()
	val := vm.pop()
	if vm.isFalsy(val) {
		frame := vm.currentFrame()
		frame.IP += offset
	}
}

func (vm *VM) opIfFalse() {
	offset := vm.readShort()
	val := vm.pop()
	if vm.isTruthy(val) {
		frame := vm.currentFrame()
		frame.IP += offset
	}
}

// --- Function Ops ---

func (vm *VM) opMakeFn() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	fnVal := frame.Closure.Fn.Chunk.Constants[idx]
	fn := fnVal.Obj.(*mapval.Function)
	closure := &mapval.Closure{Fn: fn}
	
	closure.Upvalues = make([]*mapval.Upvalue, fn.UpvalueCount)
	for i := 0; i < fn.UpvalueCount; i++ {
		isLocal := vm.readByte()
		index := vm.readByte()
		
		if isLocal == 1 {
			closure.Upvalues[i] = vm.captureUpvalue(frame.Base + int(index))
		} else {
			closure.Upvalues[i] = frame.Closure.Upvalues[index]
		}
	}
	
	vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})
}

func (vm *VM) captureUpvalue(location int) *mapval.Upvalue {
	// TODO: Reuse open upvalues (list in VM/Frame?) to support aliasing.
	// For now, always create new (no aliasing support between closures yet).
	// Proper implementation requires keeping a linked list of open upvalues.
	
	val := &vm.Stack[location]
	return &mapval.Upvalue{Location: val}
}

func (vm *VM) opCall() error {
	argCount := int(vm.readByte())

	calleeVal := vm.peek(argCount)
	if calleeVal.Type != mapval.ValObject {
		return fmt.Errorf("can only call closures")
	}

	closure, ok := calleeVal.Obj.(*mapval.Closure)
	if !ok {
		return fmt.Errorf("can only call closures")
	}

	if argCount != closure.Fn.Arity {
		return fmt.Errorf("arity mismatch: expected %d, got %d", closure.Fn.Arity, argCount)
	}

	// Cactus Stack: Allocate new frame on heap
	newFrame := &CallFrame{
		Parent:  vm.CurrentFrame,
		Closure: closure,
		IP:      0,
		Base:    vm.SP - argCount,
	}

	vm.CurrentFrame = newFrame
	
	return nil
}

func (vm *VM) opReturn() (int, error) {
	result := vm.pop()
	frame := vm.currentFrame() // Frame returning FROM

	vm.CurrentFrame = frame.Parent // Pop frame
	
	if vm.CurrentFrame == nil {
		vm.pop()      // Pop Main Script Closure
		return 1, nil // Done
	}

	vm.SP = frame.Base - 1
	vm.push(result)
	return 0, nil
}

func asFloat(v mapval.Value) float64 {
	if v.Type == mapval.ValInteger {
		return float64(v.Integer)
	}
	if v.Type == mapval.ValFloat {
		return v.Float
	}
	return 0.0
}

func isEqual(a, b mapval.Value) bool {
	if a.Type != b.Type {
		if (a.Type == mapval.ValInteger || a.Type == mapval.ValFloat) && (b.Type == mapval.ValInteger || b.Type == mapval.ValFloat) {
			return asFloat(a) == asFloat(b)
		}
		return false
	}
	switch a.Type {
	case mapval.ValInteger:
		return a.Integer == b.Integer
	case mapval.ValFloat:
		return a.Float == b.Float
	case mapval.ValBoolean:
		return a.Integer == b.Integer
	case mapval.ValEmpty:
		return true
	case mapval.ValText:
		return a.Str == b.Str
	case mapval.ValObject:
		// Handle Map equality
		mapA, okA := a.Obj.(*mapval.Map)
		mapB, okB := b.Obj.(*mapval.Map)
		if okA && okB {
			if len(mapA.Fields) != len(mapB.Fields) {
				return false
			}
			for i := range mapA.Fields {
				if !isEqual(mapA.Fields[i], mapB.Fields[i]) {
					return false
				}
			}
			return true
		}
		// Handle Tuple/Signal equality?
		tupA, okTA := a.Obj.(*mapval.Tuple)
		tupB, okTB := b.Obj.(*mapval.Tuple)
		if okTA && okTB {
			return tupA.Kind == tupB.Kind && tupA.Topic == tupB.Topic // And payload?
		}
		return a.Obj == b.Obj // Reference equality for others
	default:
		return false
	}
}

func (vm *VM) isFalsy(val mapval.Value) bool {
	if val.Type == mapval.ValEmpty {
		return true
	}
	if val.Type == mapval.ValBoolean && val.Integer == 0 {
		return true
	}
	if val.Type == mapval.ValInteger && val.Integer == 0 {
		return true
	}
	return false
}

func (vm *VM) isTruthy(val mapval.Value) bool {
	return !vm.isFalsy(val)
}

func numericCompare(a, b mapval.Value) (int, error) {
	if (a.Type != mapval.ValInteger && a.Type != mapval.ValFloat) || (b.Type != mapval.ValInteger && b.Type != mapval.ValFloat) {
		return 0, fmt.Errorf("operands must be numbers for comparison")
	}

	fa := asFloat(a)
	fb := asFloat(b)

	if fa < fb {
		return -1, nil
	}
	if fa > fb {
		return 1, nil
	}
	return 0, nil
}

// --- Testing Ops ---

func (vm *VM) opAssertEq() {
	expected := vm.pop()
	actual := vm.pop()
	
	if !isEqual(actual, expected) {
		panic(fmt.Sprintf("Assertion Failed: Expected %s, got %s", expected, actual))
	}
	fmt.Printf("PASS: %s\n", actual)
}
	