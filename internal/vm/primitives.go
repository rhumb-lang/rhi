package vm

import (
	"fmt"
	"math"

	"github.com/cockroachdb/apd/v3"
	mapval "github.com/rhumb-lang/rhi/internal/map"
)

// Precision context for Decimal operations
var decimalCtx = apd.BaseContext.WithPrecision(20)

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
	index := int(idx)

	// Check immutability
	if frame.LocalsFrozen != nil {
		if frozen, ok := frame.LocalsFrozen[index]; ok && *frozen {
			// Emit #*** signal
			msg := fmt.Sprintf("cannot assign value to immutable label (local %d)", index)
			sig := mapval.NewErrorSignal(0, msg, mapval.NewEmpty())
			vm.raiseSignal("***", sig)
			return
		}
	}

	val := vm.peek(0)
	vm.Stack[frame.Base+index] = val
}

func (vm *VM) opStoreLocImmut() {
	idx := vm.readByte()
	frame := vm.currentFrame()
	index := int(idx)

	if frame.LocalsFrozen == nil {
		frame.LocalsFrozen = make(map[int]*bool)
	}

	// Check immutability (cannot re-assign even with .= if already frozen)
	frozen, ok := frame.LocalsFrozen[index]
	if ok && *frozen {
		msg := fmt.Sprintf("cannot assign value to immutable label (local %d)", index)
		sig := mapval.NewErrorSignal(0, msg, mapval.NewEmpty())
		vm.signalVal = sig
		vm.state = StateSignaling
		return
	}

	if !ok {
		// Allocate new bool on heap
		heapBool := new(bool)
		*heapBool = true
		frame.LocalsFrozen[index] = heapBool
	} else {
		*frozen = true
	}

	val := vm.peek(0)
	vm.Stack[frame.Base+index] = val
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

	upvalue := closure.Upvalues[idx]

	fmt.Printf("DEBUG: StoreUpvalue idx %d Frozen ptr: %p\n", idx, upvalue.Frozen)
	fmt.Printf("DEBUG: Frame Base: %d, SP: %d\n", frame.Base, vm.SP)
	if upvalue.Frozen != nil {
		fmt.Printf("DEBUG: ... Value is: %v\n", *upvalue.Frozen)
	}

	// Check immutability
	if upvalue.Frozen != nil && *upvalue.Frozen {
		// Emit #*** signal instead of panic
		msg := "cannot assign value to immutable label"
		sig := mapval.NewErrorSignal(0, msg, mapval.NewEmpty())
		vm.raiseSignal("***", sig)
		return
	}

	val := vm.peek(0)

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

	// --- Path A: Numeric (Non-Time) ---
	// If both are numbers, use the Numeric Hierarchy
	if isNumeric(a) && isNumeric(b) {
		res, err := performNumericOp("+", a, b)
		if err != nil {
			return err
		}
		vm.push(res)
		return nil
	}

	// --- Path B: Time Algebra ---
	isTimeA := a.Type == mapval.ValDateTime || a.Type == mapval.ValDuration
	isTimeB := b.Type == mapval.ValDateTime || b.Type == mapval.ValDuration

	// 1. Coerce Scalars to Duration if interacting with Time
	var msA, msB int64

	// Normalize A
	if a.Type == mapval.ValDuration {
		msA = a.Integer
	} else if a.Type == mapval.ValDateTime {
		msA = a.Integer
	} else if okA, ms := scalarToDuration(a); okA {
		msA = ms
	}

	// Normalize B
	if b.Type == mapval.ValDuration {
		msB = b.Integer
	} else if b.Type == mapval.ValDateTime {
		msB = b.Integer
	} else if okB, ms := scalarToDuration(b); okB {
		msB = ms
	}

	// 2. Perform Time Addition
	// Case: Date + Duration (or Duration + Date) -> Date
	if (a.Type == mapval.ValDateTime && (b.Type == mapval.ValDuration || isNumeric(b))) ||
		(b.Type == mapval.ValDateTime && (a.Type == mapval.ValDuration || isNumeric(a))) {
		vm.push(mapval.Value{Type: mapval.ValDateTime, Integer: msA + msB})
		return nil
	}

	// Case: Duration + Duration -> Duration
	// (Scalar + Duration is treated as Duration + Duration)
	if (a.Type == mapval.ValDuration || isNumeric(a)) &&
		(b.Type == mapval.ValDuration || isNumeric(b)) {
		// Ensure at least one was actually a time type to trigger this path
		if isTimeA || isTimeB {
			vm.push(mapval.Value{Type: mapval.ValDuration, Integer: msA + msB})
			return nil
		}
	}

	// Case: Date + Date -> Error
	if a.Type == mapval.ValDateTime && b.Type == mapval.ValDateTime {
		return fmt.Errorf("cannot add two dates")
	}

	return fmt.Errorf("invalid operands for addition: %s and %s", a, b)
}

func (vm *VM) opSub() error {
	b := vm.pop()
	a := vm.pop()

	// --- Path A: Numeric ---
	if isNumeric(a) && isNumeric(b) {
		res, err := performNumericOp("-", a, b)
		if err != nil {
			return err
		}
		vm.push(res)
		return nil
	}

	// --- Path B: Time Algebra ---

	// Normalize to Milliseconds (same logic as Add)
	var msA, msB int64
	if a.Type == mapval.ValDateTime || a.Type == mapval.ValDuration {
		msA = a.Integer
	} else if okA, ms := scalarToDuration(a); okA {
		msA = ms
	}

	if b.Type == mapval.ValDateTime || b.Type == mapval.ValDuration {
		msB = b.Integer
	} else if okB, ms := scalarToDuration(b); okB {
		msB = ms
	}

	// Case: Date - Duration (or Scalar) -> Date
	if a.Type == mapval.ValDateTime && (b.Type == mapval.ValDuration || isNumeric(b)) {
		vm.push(mapval.Value{Type: mapval.ValDateTime, Integer: msA - msB})
		return nil
	}

	// Case: Date - Date -> Duration
	if a.Type == mapval.ValDateTime && b.Type == mapval.ValDateTime {
		vm.push(mapval.Value{Type: mapval.ValDuration, Integer: msA - msB})
		return nil
	}

	// Case: Duration - Duration -> Duration
	if (a.Type == mapval.ValDuration || isNumeric(a)) &&
		(b.Type == mapval.ValDuration || isNumeric(b)) {
		if a.Type == mapval.ValDuration || b.Type == mapval.ValDuration {
			vm.push(mapval.Value{Type: mapval.ValDuration, Integer: msA - msB})
			return nil
		}
	}

	// Case: Scalar - Date -> Error
	if isNumeric(a) && b.Type == mapval.ValDateTime {
		return fmt.Errorf("cannot subtract date from scalar")
	}

	return fmt.Errorf("invalid operands for subtraction")
}

func (vm *VM) opCoerceNum() error {
	val := vm.pop()
	if val.Type == mapval.ValDateTime {
		// +Date -> Duration
		vm.push(mapval.Value{Type: mapval.ValDuration, Integer: val.Integer})
		return nil
	}
	if val.Type == mapval.ValDuration || val.Type == mapval.ValInteger || val.Type == mapval.ValFloat || val.Type == mapval.ValDecimal {
		vm.push(val)
		return nil
	}
	return fmt.Errorf("cannot coerce type %d to number", val.Type)
}

func (vm *VM) opNumNeg() error {
	val := vm.pop()
	if val.Type == mapval.ValDateTime {
		// -Date -> Negative Duration
		vm.push(mapval.Value{Type: mapval.ValDuration, Integer: -val.Integer})
		return nil
	}
	if val.Type == mapval.ValDuration {
		vm.push(mapval.Value{Type: mapval.ValDuration, Integer: -val.Integer})
		return nil
	}

	if val.Type == mapval.ValInteger {
		vm.push(mapval.NewInt(-val.Integer))
		return nil
	}
	if val.Type == mapval.ValFloat {
		vm.push(mapval.NewFloat(-val.Float))
		return nil
	}
	if val.Type == mapval.ValDecimal {
		d := toDecimal(val)
		res := new(apd.Decimal)
		res.Neg(d)
		vm.push(mapval.Value{Type: mapval.ValDecimal, Obj: &mapval.Decimal{Raw: res}})
		return nil
	}

	return fmt.Errorf("cannot negate type %d", val.Type)
}

func (vm *VM) opMult() error {
	b := vm.pop()
	a := vm.pop()

	if isNumeric(a) && isNumeric(b) {
		res, err := performNumericOp("*", a, b)
		if err != nil {
			return err
		}
		vm.push(res)
		return nil
	}
	return fmt.Errorf("invalid operands for multiplication")
}

func (vm *VM) opDivFloat() error {
	b := vm.pop()
	a := vm.pop()

	if isNumeric(a) && isNumeric(b) {
		res, err := performNumericOp("/", a, b)
		if err != nil {
			return err
		}
		vm.push(res)
		return nil
	}
	return fmt.Errorf("invalid operands for division")
}

func (vm *VM) opDivInt() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		if b.Integer == 0 {
			return fmt.Errorf("division by zero")
		}
		vm.push(mapval.NewInt(a.Integer / b.Integer))
	} else {
		// Promote to integer? or perform numeric op and truncate?
		// Usually // is int div.
		// If dealing with floats/decimals, we might want to floor.
		// For now, let's keep it simple or promote to float then floor.
		vm.push(mapval.NewInt(int64(asFloat(a) / asFloat(b))))
	}
	return nil
}

// performNumericOp handles +, -, *, / for pure numbers (Int, Float, Decimal)
// Returns result Value or error.
func performNumericOp(op string, a, b mapval.Value) (mapval.Value, error) {
	// 1. If either is Decimal -> Promote both to Decimal
	if a.Type == mapval.ValDecimal || b.Type == mapval.ValDecimal {
		dA := toDecimal(a)
		dB := toDecimal(b)
		res := new(apd.Decimal)

		var err error
		switch op {
		case "+":
			_, err = decimalCtx.Add(res, dA, dB)
		case "-":
			_, err = decimalCtx.Sub(res, dA, dB)
		case "*":
			_, err = decimalCtx.Mul(res, dA, dB)
		case "/":
			_, err = decimalCtx.Quo(res, dA, dB)
		}
		if err != nil {
			return mapval.NewEmpty(), err
		}
		return mapval.Value{Type: mapval.ValDecimal, Obj: &mapval.Decimal{Raw: res}}, nil
	}

	// 2. If either is Float -> Promote both to Float
	if a.Type == mapval.ValFloat || b.Type == mapval.ValFloat {
		fA := asFloat(a)
		fB := asFloat(b)
		switch op {
		case "+":
			return mapval.NewFloat(fA + fB), nil
		case "-":
			return mapval.NewFloat(fA - fB), nil
		case "*":
			return mapval.NewFloat(fA * fB), nil
		case "/":
			return mapval.NewFloat(fA / fB), nil
		}
	}

	// 3. Fallback: Both are Integers
	iA := a.Integer
	iB := b.Integer
	switch op {
	case "+":
		return mapval.NewInt(iA + iB), nil
	case "-":
		return mapval.NewInt(iA - iB), nil
	case "*":
		return mapval.NewInt(iA * iB), nil
	case "/":
		if iB == 0 {
			return mapval.NewEmpty(), fmt.Errorf("division by zero")
		}
		// opDivInt handles integer division
		if op == "/" {
			return mapval.NewFloat(float64(iA) / float64(iB)), nil
		}
		return mapval.NewInt(iA / iB), nil
	}

	return mapval.NewEmpty(), fmt.Errorf("unknown numeric operator %s", op)
}

// toDecimal converts Int/Float/Decimal to *apd.Decimal
func toDecimal(v mapval.Value) *apd.Decimal {
	if v.Type == mapval.ValDecimal {
		return v.Obj.(*mapval.Decimal).Raw
	}
	d := new(apd.Decimal)
	if v.Type == mapval.ValInteger {
		d.SetInt64(v.Integer)
	} else if v.Type == mapval.ValFloat {
		d.SetFloat64(v.Float)
	}
	return d
}

// scalarToDuration converts a numeric value to Milliseconds (int64)
// Returns (true, ms) if convertible, (false, 0) if not a number.
func scalarToDuration(v mapval.Value) (bool, int64) {
	switch v.Type {
	case mapval.ValInteger:
		// Integer -> Milliseconds (Direct)
		return true, v.Integer
	case mapval.ValFloat:
		// Float -> Seconds -> Milliseconds
		return true, int64(v.Float * 1000)
	case mapval.ValDecimal:
		// Decimal -> Seconds -> Milliseconds
		d := v.Obj.(*mapval.Decimal).Raw
		// Multiply by 1000
		thou := apd.New(1000, 0)
		res := new(apd.Decimal)
		decimalCtx.Mul(res, d, thou)
		// Convert to Int64
		i, _ := res.Int64()
		return true, i
	}
	return false, 0
}

func isNumeric(v mapval.Value) bool {
	return v.Type == mapval.ValInteger || v.Type == mapval.ValFloat || v.Type == mapval.ValDecimal
}

func (vm *VM) opMod() error {
	b := vm.pop()
	a := vm.pop()
	if a.Type == mapval.ValInteger && b.Type == mapval.ValInteger {
		if b.Integer == 0 {
			return fmt.Errorf("division by zero")
		}
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
	// Use Equality (Wildcard aware for versions) for ==
	vm.push(mapval.NewBoolean(isEqual(a, b)))
}

func (vm *VM) opNeq() {
	b := vm.pop()
	a := vm.pop()
	// Use Strict inequality for ~~
	vm.push(mapval.NewBoolean(!isStrictEqual(a, b)))
}

func (vm *VM) opGt() error {
	b := vm.pop()
	a := vm.pop()

	if a.Type == mapval.ValVersion && b.Type == mapval.ValVersion {
		cmp := compareVersions(a, b)
		vm.push(mapval.NewBoolean(cmp == 1))
		return nil
	}

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

	if a.Type == mapval.ValVersion && b.Type == mapval.ValVersion {
		cmp := compareVersions(a, b)
		vm.push(mapval.NewBoolean(cmp == -1))
		return nil
	}

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

	if a.Type == mapval.ValVersion && b.Type == mapval.ValVersion {
		cmp := compareVersions(a, b)
		vm.push(mapval.NewBoolean(cmp == 1 || cmp == 0))
		return nil
	}

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

	if a.Type == mapval.ValVersion && b.Type == mapval.ValVersion {
		cmp := compareVersions(a, b)
		vm.push(mapval.NewBoolean(cmp == -1 || cmp == 0))
		return nil
	}

	res, err := numericCompare(a, b)
	if err != nil {
		return err
	}
	vm.push(mapval.NewBoolean(res <= 0))
	return nil
}

func asFloat(v mapval.Value) float64 {
	// ... (unchanged)
	if v.Type == mapval.ValInteger {
		return float64(v.Integer)
	}
	if v.Type == mapval.ValFloat {
		return v.Float
	}
	return 0.0
}

func isStrictEqual(a, b mapval.Value) bool {
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
	case mapval.ValVersion:
		// Strict check: Integer encoding and Suffix must match
		return a.Integer == b.Integer && a.Str == b.Str
	case mapval.ValObject:
		// Handle Map equality
		mapA, okA := a.Obj.(*mapval.Map)
		mapB, okB := b.Obj.(*mapval.Map)
		if okA && okB {
			if len(mapA.Fields) != len(mapB.Fields) {
				return false
			}
			for i := range mapA.Fields {
				if !isStrictEqual(mapA.Fields[i], mapB.Fields[i]) {
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

func isEqual(a, b mapval.Value) bool {
	if a.Type == mapval.ValVersion && b.Type == mapval.ValVersion {
		return compareVersions(a, b) == 0
	}
	return isStrictEqual(a, b)
}

func compareVersions(a, b mapval.Value) int {
	M1, m1, p1, w1 := a.VersionUnpack()
	M2, m2, p2, w2 := b.VersionUnpack()

	// Major
	if M1 > M2 {
		return 1
	}
	if M1 < M2 {
		return -1
	}

	// Minor
	// Check Wildcards: If either is wildcard AT THIS LEVEL, it's a match (0) for this and subsequent levels.
	// Wildcard bit is set.
	// How to check level? Use Sentinel values.
	// MaxUint16 = 65535
	// MaxUint32 = 4294967295

	wild1 := w1 && m1 == 0xFFFF
	wild2 := w2 && m2 == 0xFFFF

	if wild1 || wild2 {
		// If either is 1.- (Major Wildcard), then any Minor/Patch matches.
		// Base versions match (Major is same).
		return 0
	}

	if m1 > m2 {
		return 1
	}
	if m1 < m2 {
		return -1
	}

	// Patch
	wild1 = w1 && p1 == 0xFFFFFFFF
	wild2 = w2 && p2 == 0xFFFFFFFF

	if wild1 || wild2 {
		return 0
	}

	if p1 > p2 {
		return 1
	}
	if p1 < p2 {
		return -1
	}

	// Suffixes
	// If suffixes differ, return Incomparable? Or non-zero?
	// Tests:
	// v1 (1.0.0) vs v7 (1.0.0-pre) -> GT=no, LT=no, EQ=no.
	// This implies compareVersions returned something other than 1, -1, 0.
	// Let's use -2.

	if a.Str != b.Str {
		// If wildcards were active at patch level, we returned 0 already.
		// "Wildcards cannot be used in conjunction with suffixes".
		// But if comparing `1.-` vs `1.0.0-pre`?
		// `1.-` has `wild1` true at minor/patch. We returned 0.
		// So suffixes ignored if wildcard matched earlier.
		// But here, no wildcard matched.
		// Exact version match `1.0.0` vs `1.0.0`.
		// Suffixes differ.
		return -2
	}

	return 0
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
		index := int(vm.readByte())

		if isLocal == 1 {
			// Ensure we have a frozen pointer for this local
			if frame.LocalsFrozen == nil {
				frame.LocalsFrozen = make(map[int]*bool)
			}
			
			frozenPtr, ok := frame.LocalsFrozen[index]
			if !ok {
				// Allocate new bool on heap (default false/mutable)
				b := new(bool)
				*b = false
				frame.LocalsFrozen[index] = b
				frozenPtr = b
			}

			closure.Upvalues[i] = vm.captureUpvalue(frame.Base+index, frozenPtr)
		} else {
			closure.Upvalues[i] = frame.Closure.Upvalues[index]
		}
	}

	vm.push(mapval.Value{Type: mapval.ValObject, Obj: closure})
}

func (vm *VM) captureUpvalue(location int, frozen *bool) *mapval.Upvalue {
	val := &vm.Stack[location]
	// Note: Rhumb VM currently doesn't share open upvalues (cache).
	// This means multiple closures capturing the same var get different Upvalue structs.
	// But they point to the same Stack location and the same Frozen bool pointer.
	// So they share state correctly.
	return &mapval.Upvalue{Location: val, Frozen: frozen}
}

func (vm *VM) opCall() error {
	argCount := int(vm.readByte())

	calleeVal := vm.peek(argCount)
	if calleeVal.Type != mapval.ValObject {
		return fmt.Errorf("can only call closures, maps, or native functions")
	}

	if closure, ok := calleeVal.Obj.(*mapval.Closure); ok {
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

	if native, ok := calleeVal.Obj.(*mapval.NativeFunction); ok {
		// 1. Pop Args
		args := make([]mapval.Value, argCount)
		for i := argCount - 1; i >= 0; i-- {
			args[i] = vm.pop()
		}
		vm.pop() // Pop function

		// 2. Execute
		// Note: No 'err' return. Native code handles its own errors via Signals.
		result := native.Fn(args)

		// 3. Push Result
		vm.push(result)
		return nil
	}

	if m, ok := calleeVal.Obj.(*mapval.Map); ok {
		if argCount != 1 {
			return fmt.Errorf("map access requires exactly 1 argument")
		}
		keyVal := vm.peek(0)

		var keyStr string
		if keyVal.Type == mapval.ValText {
			keyStr = keyVal.Str
		} else if keyVal.Type == mapval.ValInteger {
			keyStr = fmt.Sprintf("%d", keyVal.Integer)
		} else {
			keyStr = keyVal.Canonical()
		}

		var result mapval.Value = mapval.NewEmpty()

		if m.Legend != nil {
			for i, f := range m.Legend.Fields {
				if f.Name == keyStr {
					result = m.Fields[i]
					break
				}
			}
		}

		// Pop arg and callee
		vm.pop() // arg
		vm.pop() // callee
		vm.push(result)
		return nil
	}

	return fmt.Errorf("can only call closures or maps")
}

func (vm *VM) opReturn() (int, error) {
	result := vm.pop()
	frame := vm.currentFrame() // Frame returning FROM

	vm.CurrentFrame = frame.Parent // Pop frame

	// Reset Stack Pointer
	// If Base > 0, it means there is a Closure (or Caller's stack) below.
	// Standard convention (OP_CALL): Base-1 is the Closure. We pop it.
	// Interpret convention: Base=0. No Closure on stack. We just reset to 0.
	targetSP := frame.Base
	if targetSP > 0 {
		targetSP--
	}

	vm.SP = targetSP
	vm.push(result)

	if vm.CurrentFrame == nil {
		return 1, nil // Done
	}
	return 0, nil
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

func (vm *VM) opInspect() {
	val := vm.pop()
	fmt.Printf("INSPECT: %s\n", val.Canonical())
}

func (vm *VM) opAssertEq() {
	nameVal := vm.pop()
	expected := vm.pop()
	actual := vm.pop()

	expectedStr := expected.Str

	actualStr := actual.Canonical()

	name := ""
	if nameVal.Type == mapval.ValText {
		name = nameVal.Str
	}

	if actualStr != expectedStr {
		msg := fmt.Sprintf("FAIL: Expected '%s', got '%s'", expectedStr, actualStr)
		if name != "" {
			fmt.Printf("%s (%s)\n", msg, name)
		} else {
			fmt.Printf("%s\n", msg)
		}
	} else {
		if name != "" {
			fmt.Printf("PASS: %s (%s)\n", actualStr, name)
		} else {
			fmt.Printf("PASS: %s\n", actualStr)
		}
	}
}
