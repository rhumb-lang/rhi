package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type word uint64

const SIGN_BIT uint64 = 0x8000_0000_0000_0000
const QUIET_NAN uint64 = 0x7FFC_0000_0000_0000

const TAG_EMPTY = 0b0001
const TAG_FALSE = 0b0010
const TAG_TRUE = 0b0011
const TAG_NAN = 0b0100
const TAG_INF = 0b0101
const TAG_PI = 0b0110
const TAG_EULER = 0b0111
const TAG_MARK = 0b1000
const TAG_MARK_WHITE = 0b1001
const TAG_MARK_GREY = 0b1010
const TAG_MARK_BLACK = 0b1011

const EMPTY uint64 = QUIET_NAN | TAG_EMPTY
const FALSE uint64 = QUIET_NAN | TAG_FALSE
const TRUE uint64 = QUIET_NAN | TAG_TRUE
const NAN uint64 = QUIET_NAN | TAG_NAN
const INF uint64 = QUIET_NAN | TAG_INF
const PI uint64 = QUIET_NAN | TAG_PI
const EULER uint64 = QUIET_NAN | TAG_EULER

type Word interface {
	IsEmpty() bool
	IsBool() bool
	IsNAN() bool
	IsInf() bool
	IsPi() bool
	IsEuler() bool
	IsNumber() bool
	IsMapAddr() bool
	IsMarker() bool

	AsBool() bool
	AsNumber() float64
	AsMapAddr() uint64
}

func FromNumber(f float64) word {
	var bytes [8]byte
	fbits := math.Float64bits(f)
	bytes[0] = byte(fbits >> 56)
	bytes[1] = byte(fbits >> 48)
	bytes[2] = byte(fbits >> 40)
	bytes[3] = byte(fbits >> 32)
	bytes[4] = byte(fbits >> 24)
	bytes[5] = byte(fbits >> 16)
	bytes[6] = byte(fbits >> 8)
	bytes[7] = byte(fbits)
	return word(binary.BigEndian.Uint64(bytes[:]))
}

func FromBool(b bool) word {
	if b {
		return word(TRUE)
	} else {
		return word(FALSE)
	}
}

/*
*
 */
func FromMapAddr(a uint64) word {
	return word(a & ^(SIGN_BIT | QUIET_NAN))
}

func (w word) IsEmpty() bool  { return uint64(w) == EMPTY }
func (w word) IsBool() bool   { return (uint64(w) | 1) == TRUE }
func (w word) IsNAN() bool    { return uint64(w) == NAN }
func (w word) IsInf() bool    { return uint64(w) == INF }
func (w word) IsPi() bool     { return uint64(w) == PI }
func (w word) IsEuler() bool  { return uint64(w) == EULER }
func (w word) IsNumber() bool { return (uint64(w) & QUIET_NAN) != QUIET_NAN }
func (w word) IsMapAddr() bool {
	return (uint64(w) & (QUIET_NAN | SIGN_BIT)) == (QUIET_NAN | SIGN_BIT)
}
func (w word) IsMarker() bool {
	return (uint64(w) & (QUIET_NAN | TAG_MARK)) == (QUIET_NAN | TAG_MARK)
}

func (w word) AsBool() bool { return uint64(w) == TRUE }
func (w word) AsNumber() float64 {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, w); err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	bits := binary.BigEndian.Uint64(buf.Bytes())
	float := math.Float64frombits(bits)
	return float
}
func (w word) AsMapAddr() uint64 {
	return uint64(w) & ^(SIGN_BIT | QUIET_NAN)
}
