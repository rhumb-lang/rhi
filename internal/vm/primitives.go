package vm

import (
	"fmt"

	mapval "git.sr.ht/~madcapjake/rhi/internal/map"
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
	} else {
		return fmt.Errorf("operands must be numbers for ADD")
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
	vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})
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

	if vm.FrameCount >= MaxFrames {
		return fmt.Errorf("stack overflow")
	}

	vm.Frames[vm.FrameCount] = CallFrame{
		Closure: closure,
		IP:      0,
		Base:    vm.SP - argCount - 1,
	}
	vm.FrameCount++
	return nil
}

func (vm *VM) opReturn() (int, error) {
	result := vm.pop()
	frame := vm.currentFrame() // Frame returning FROM

	vm.FrameCount--
	if vm.FrameCount == 0 {
		vm.pop()      // Pop Main Script Closure
		return 1, nil // Done
	}

	vm.SP = frame.Base
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
	default:
		return false
	}
}

func (vm *VM) isFalsy(val mapval.Value) bool {
	return val.Type == mapval.ValEmpty || (val.Type == mapval.ValBoolean && val.Integer == 0)
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
