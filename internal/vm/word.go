package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type Word uint64

const TAG_SLT_NAN uint64 = 0x7F_FC_00_00_00_00_00_00

const TAG_ADDRESS uint64 = 0x80_00_00_00_00_00_00_00

const TAG_LIT_NAN uint64 = 0x7F_FC_00_00_00_00_00_00
const TAG_LIT_TRU uint64 = 0x7F_FC_40_00_00_00_00_00
const TAG_LIT_FLS uint64 = 0x7F_FC_80_00_00_00_00_00
const TAG_LIT_ETY uint64 = 0x7F_FC_C0_00_00_00_00_00

const TAG_LIT_INT uint64 = 0x7F_FD_00_00_00_00_00_00
const TAG_LIT_RNE uint64 = 0x7F_FD_40_00_00_00_00_00
const TAG_LIT_SYM uint64 = 0x7F_FD_80_00_00_00_00_00

const TAG_OBJ_MAP uint64 = 0x7F_FE_00_00_00_00_00_00
const TAG_OBJ_ARR uint64 = 0x7F_FE_40_00_00_00_00_00
const TAG_OBJ_ROU uint64 = 0x7F_FE_80_00_00_00_00_00
const TAG_OBJ_CHU uint64 = 0x7F_FE_C0_00_00_00_00_00

const TAG_LEG_MAP uint64 = 0x7F_FF_00_00_00_00_00_00
const TAG_LEG_ARR uint64 = 0x7F_FF_10_00_00_00_00_00
const TAG_LEG_ROU uint64 = 0x7F_FF_C0_00_00_00_00_00

const TAG_MRK_STL uint64 = 0x7F_FD_C0_00_00_00_00_00

const TAG_ALLMARK uint64 = TAG_LEG_ROU

const TAG_SWP_WHT uint64 = 0x00_00_00_00_00_00_00_00
const TAG_SWP_GRY uint64 = 0x00_00_00_00_00_00_00_01
const TAG_SWP_BLK uint64 = 0x00_00_00_00_00_00_00_02
const TAG_SWPMASK uint64 = 0x00_00_00_00_00_00_00_03

func NewWord(a any) Word {
	switch a := a.(type) {
	case float64:
		return WordFromFloat(a)
	case float32:
		return WordFromFloat(float64(a))
	case bool:
		return WordFromBool(a)
	case int:
		return WordFromInt(uint32(a))
	case int64:
		return WordFromInt(uint32(a))
	case rune:
		return WordFromRune(a)
	default:
		return EmptyWord()
	}
}

func EmptyWord() Word { return Word(TAG_LIT_ETY) }

func WordFromFloat(f float64) Word {
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
	return Word(binary.LittleEndian.Uint64(bytes[:]))
}

func WordFromBool(b bool) Word {
	if b {
		return Word(TAG_LIT_TRU)
	} else {
		return Word(TAG_LIT_FLS)
	}
}
func WordFromInt(i uint32) Word { return Word(TAG_LIT_INT | uint64(i)) }
func WordFromRune(r rune) Word  { return Word(TAG_LIT_RNE | uint64(r)) }
func WordFromSym(a int) Word    { return Word(TAG_LIT_SYM | uint64(a)) }
func WordFromAddr(a int) Word   { return Word(TAG_ADDRESS | uint64(a)) }

func (w Word) IsNAN() bool   { return uint64(w) == TAG_LIT_NAN }
func (w Word) IsTrue() bool  { return uint64(w) == TAG_LIT_TRU }
func (w Word) IsFalse() bool { return uint64(w) == TAG_LIT_FLS }
func (w Word) IsBool() bool  { return w.IsTrue() || w.IsFalse() }
func (w Word) IsEmpty() bool { return uint64(w) == TAG_LIT_ETY }
func (w Word) IsInt() bool   { return uint64(w)&TAG_LIT_INT == TAG_LIT_INT }
func (w Word) IsRune() bool  { return uint64(w)&TAG_LIT_RNE == TAG_LIT_RNE }
func (w Word) IsSym() bool   { return uint64(w)&TAG_LIT_SYM == TAG_LIT_SYM }
func (w Word) IsFloat() bool { return uint64(w)&TAG_SLT_NAN != TAG_SLT_NAN }
func (w Word) IsMark() bool  { return uint64(w)&TAG_OBJ_MAP == TAG_OBJ_MAP }
func (w Word) IsAddr() bool  { return uint64(w)&TAG_ADDRESS == TAG_ADDRESS }

func (w Word) isMrkT(t uint64) bool { return uint64(w)&TAG_ALLMARK == t }
func (w Word) IsMap() bool          { return w.isMrkT(TAG_OBJ_MAP) }
func (w Word) IsArray() bool        { return w.isMrkT(TAG_OBJ_ARR) }
func (w Word) IsRoutine() bool      { return w.isMrkT(TAG_OBJ_ROU) }
func (w Word) IsChunk() bool        { return w.isMrkT(TAG_OBJ_CHU) }

func (w Word) Debug() any {
	if w.IsNAN() {
		return "NAN"
	} else if w.IsTrue() {
		return "TRUE"
	} else if w.IsFalse() {
		return "FALSE"
	} else if w.IsEmpty() {
		return "EMPTY"
	} else if w.IsInt() {
		return fmt.Sprint("{INT: ", w.AsInt(), "}")
	} else if w.IsRune() {
		return fmt.Sprint("{RUNE: ", w.AsRune(), "}")
	} else if w.IsSym() {
		return "SYMBOL" // FIXME: Implement symbols
	} else if w.IsAddr() { // must trigger before IsFloat for unknown reasons
		return fmt.Sprint("{ADDR: ", w.AsAddr(), "}")
	} else if w.IsFloat() {
		return fmt.Sprint("{FLOAT: ", w.AsFloat(), "}")
	} else if w.IsMap() {
		return "MAP"
	} else if w.IsArray() {
		return "ARRAY"
	} else if w.IsRoutine() {
		return "ROUTINE"
	} else if w.IsChunk() {
		return "CHUNK"
	} else {
		return nil
	}
}

func (w Word) isSwpT(t uint64) bool {
	return uint64(w)&(TAG_OBJ_MAP|TAG_SWPMASK) == t
}
func (w Word) IsGrey() bool  { return w.isSwpT(TAG_OBJ_MAP | TAG_SWP_GRY) }
func (w Word) IsBlack() bool { return w.isSwpT(TAG_OBJ_MAP | TAG_SWP_BLK) }

func (w Word) AsBool() bool {
	switch uint64(w) {
	case TAG_LIT_TRU:
		return true
	case TAG_LIT_FLS, TAG_LIT_ETY, TAG_LIT_NAN:
		return false
	default:
		return w.IsAddr()
	}

}
func (w Word) AsFloat() float64 {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, w); err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	bits := binary.LittleEndian.Uint64(buf.Bytes())
	float := math.Float64frombits(bits)
	return float
}
func (w Word) AsInt() uint32  { return uint32(uint64(w) & ^TAG_LIT_INT) }
func (w Word) AsRune() rune   { return rune(uint64(w) & ^TAG_LIT_RNE) }
func (w Word) AsSym() uint64  { return uint64(w) & ^TAG_LIT_SYM }
func (w Word) AsAddr() uint64 { return uint64(w) & ^TAG_ADDRESS }
