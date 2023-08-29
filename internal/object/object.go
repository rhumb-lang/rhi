package object

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type Any interface {
	IsObject()
	WHAT() string
	Equals(Any) bool
}

type Hashable interface {
	Any
	HashBytes() []byte
}

type Code interface {
	WHAT() string
	GetLine() int
	GetColumn() int
	GetID() int
}

type Ref[O Any] struct {
	Value int
}

func (r Ref[any]) IsObject() {}
func (r Ref[any]) WHAT() string {
	return fmt.Sprint("Reference: ", r.Value)
}
func (x Ref[any]) Equals(y Any) bool {
	objY, ok := y.(Ref[any])
	return ok && x.Value == objY.Value
}
func (key Ref[any]) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(key.Value))
	return buf[:n]
}

type Empty struct{}

func (e Empty) IsObject() {}
func (e Empty) WHAT() string {
	return "Empty"
}
func (x Empty) Equals(y Any) bool {
	_, ok := y.(Empty)
	return ok
}

type WholeNumber struct{ Value int64 }

func (w WholeNumber) IsObject() {}
func (w WholeNumber) WHAT() string {
	return fmt.Sprint("WholeNumber: ", w.Value)
}
func (x WholeNumber) Equals(y Any) bool {
	objY, ok := y.(WholeNumber)
	return ok && x.Value == objY.Value
}
func (key WholeNumber) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(key.Value))
	return buf[:n]
}

type RationalNumber struct{ Value float64 }

func (r RationalNumber) IsObject() {}
func (r RationalNumber) WHAT() string {
	return fmt.Sprint("RationalNumber: ", r.Value)
}
func (x RationalNumber) Equals(y Any) bool {
	objY, ok := y.(RationalNumber)
	return ok && x.Value == objY.Value
}

type DateNumber struct{ Value int }

func (d DateNumber) IsObject() {}
func (d DateNumber) WHAT() string {
	return fmt.Sprint("DateNumber: ", d.Value)
}
func (x DateNumber) Equals(y Any) bool {
	objY, ok := y.(DateNumber)
	return ok && x.Value == objY.Value
}
func (key DateNumber) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(key.Value))
	return buf[:n]
}

type RuneSymbol struct{ Value []rune }

func (r RuneSymbol) IsObject() {}
func (r RuneSymbol) WHAT() string {
	return fmt.Sprint("RuneSymbol: ", r.Value)
}
func (x RuneSymbol) Equals(y Any) bool {
	objY, ok := y.(RuneSymbol)
	return ok && reflect.DeepEqual(x.Value, objY.Value)
}
func (key RuneSymbol) HashBytes() []byte {
	return []byte(string(key.Value))
}

type KeySymbol struct {
	Value int
	Name  string
}

func (k KeySymbol) IsObject() {}
func (k KeySymbol) WHAT() string {
	return fmt.Sprint("KeySymbol: \"", k.Name, "\"")
}
func (x KeySymbol) Equals(y Any) bool {
	objY, ok := y.(KeySymbol)
	return ok && x.Value == objY.Value
}
func (key KeySymbol) HashBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(key.Value))
	return buf[:n]
}

type TruthSymbol struct{ Value bool }

func (t TruthSymbol) IsObject() {}
func (t TruthSymbol) WHAT() string {
	return fmt.Sprint("TruthSymbol: ", t.Value)
}
func (x TruthSymbol) Equals(y Any) bool {
	if objY, ok := y.(TruthSymbol); ok {
		return x.Value == objY.Value
	} else {
		return false
	}
}

type LabelSymbol struct{ Value string }

func (l LabelSymbol) IsObject() {}
func (l LabelSymbol) WHAT() string {
	return fmt.Sprint("LabelSymbol: \"", l.Value, "\"")
}
func (x LabelSymbol) Equals(y Any) bool {
	objY, ok := y.(LabelSymbol)
	return ok && x.Value == objY.Value
}
func (key LabelSymbol) HashBytes() []byte {
	return []byte(key.Value)
}

type ArgumentSymbol struct {
	Value int
	Name  string
}

func (a ArgumentSymbol) IsObject() {}
func (a ArgumentSymbol) WHAT() string {
	if len(a.Name) > 0 {
		return fmt.Sprint("InterpretedSymbol: $", a.Name)
	} else {
		return fmt.Sprint("InterpretedSymbol: $", a.Value)
	}
}
func (x ArgumentSymbol) Equals(y Any) bool {
	objY, ok := y.(ArgumentSymbol)
	if len(x.Name) > 0 {
		return ok && x.Name == objY.Name
	} else {
		return ok && x.Value == objY.Value
	}
}
func (key ArgumentSymbol) HashBytes() []byte {
	if len(key.Name) > 0 {
		return []byte(fmt.Sprint("$", key.Name))
	} else {
		buf := make([]byte, binary.MaxVarintLen64)
		buf[0] = '$'
		n := binary.PutVarint(buf[1:], int64(key.Value))
		return buf[:n]
	}
}
func (a ArgumentSymbol) IsNamed() bool {
	return len(a.Name) > 0
}

type TextSymbol struct{ Value string }

func (l TextSymbol) IsObject() {}
func (l TextSymbol) WHAT() string {
	return fmt.Sprint("TextSymbol: \"", l.Value, "\"")
}
func (x TextSymbol) Equals(y Any) bool {
	objY, ok := y.(TextSymbol)
	return ok && x.Value == objY.Value
}
func (key TextSymbol) HashBytes() []byte {
	return []byte(key.Value)
}

type OperatorSymbol struct{ Value string }

func (l OperatorSymbol) IsObject() {}
func (l OperatorSymbol) WHAT() string {
	return fmt.Sprint("OperatorSymbol: \"", l.Value, "\"")
}
func (x OperatorSymbol) Equals(y Any) bool {
	if objY, ok := y.(OperatorSymbol); ok {
		return x.Value == objY.Value
	} else {
		return false
	}
}
