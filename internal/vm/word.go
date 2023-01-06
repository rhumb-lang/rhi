package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type Word uint64

const MASK_NAN uint64 = 0x7F_FC_00_00_00_00_00_00

const VAL_ADDR uint64 = 0xFF_FC_00_00_00_00_00_00

const MARK_ERR uint64 = 0x7F_FC_00_00_00_00_00_00

const ERR_MAIN uint64 = 0x7F_FC_00_10_00_00_00_00
const ERR_NUMB uint64 = 0x7F_FC_00_20_00_00_00_00

const VAL_TRUE uint64 = 0x7F_FC_40_00_00_00_00_00
const VAL_FALS uint64 = 0x7F_FC_80_00_00_00_00_00
const VAL_EMPT uint64 = 0x7F_FC_C0_00_00_00_00_00

const VAL_NUMB uint64 = 0x7F_FD_00_00_00_00_00_00
const VAL_RUNE uint64 = 0x7F_FD_40_00_00_00_00_00
const VAL_SYMB uint64 = 0x7F_FD_80_00_00_00_00_00
const VAL_DATE uint64 = 0x7F_FD_C0_00_00_00_00_00

const SENTINEL uint64 = 0x7F_FE_00_00_00_00_00_00

const MARK_MAP uint64 = 0x7F_FE_40_00_00_00_00_00

const MAIN_MAP uint64 = 0x7F_FE_40_10_00_00_00_00
const LIST_MAP uint64 = 0x7F_FE_40_20_00_00_00_00
const TEXT_MAP uint64 = 0x7F_FE_40_30_00_00_00_00
const FUNC_MAP uint64 = 0x7F_FE_40_40_00_00_00_00
const META_MAP uint64 = 0x7F_FE_40_F0_00_00_00_00

const MARK_ARR uint64 = 0x7F_FE_80_00_00_00_00_00

const MAIN_ARR uint64 = 0x7F_FE_80_20_00_00_00_00
const RUNE_ARR uint64 = 0x7F_FE_80_30_00_00_00_00
const CODE_ARR uint64 = 0x7F_FE_40_40_00_00_00_00
const META_ARR uint64 = 0x7F_FE_80_F0_00_00_00_00

const MARK_LGD uint64 = 0x7F_FE_C0_00_00_00_00_00

const MAIN_LGD uint64 = 0x7F_FE_C0_10_00_00_00_00
const LIST_LGD uint64 = 0x7F_FE_C0_20_00_00_00_00
const TEXT_LGD uint64 = 0x7F_FE_C0_30_00_00_00_00
const FUNC_LGD uint64 = 0x7F_FE_C0_40_00_00_00_00
const META_LGD uint64 = 0x7F_FE_C0_F0_00_00_00_00

const MARK_DES uint64 = 0x7F_FF_00_00_00_00_00_00

const DES_NUMB uint64 = 0x7F_FF_00_10_00_00_00_00
const DES_FLOA uint64 = 0x7F_FF_00_20_00_00_00_00
const DES_BOOL uint64 = 0x7F_FF_00_30_00_00_00_00
const DES_RUNE uint64 = 0x7F_FF_00_40_00_00_00_00
const DES_SYMB uint64 = 0x7F_FF_00_50_00_00_00_00
const DES_DATE uint64 = 0x7F_FF_00_60_00_00_00_00
const DES_ADDR uint64 = 0x7F_FF_00_70_00_00_00_00

const CMP_UNIT uint64 = 0x7F_FF_40_00_00_00_00_00

const TAG_IMMU uint64 = 0x00_00_08_00_00_00_00_00
const TAG_MUTA uint64 = 0x00_00_0C_00_00_00_00_00
const TAG_GREY uint64 = 0x00_00_02_00_00_00_00_00
const TAG_BLAK uint64 = 0x00_00_03_00_00_00_00_00

const MASK_ONE uint64 = 0x7F_FF_C0_00_00_00_00_00
const MASK_TWO uint64 = 0x7F_FF_C0_F0_00_00_00_00

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

func EmptyWord() Word      { return Word(VAL_EMPT) }
func NotANumberWord() Word { return Word(ERR_NUMB) }

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
		return Word(VAL_TRUE)
	} else {
		return Word(VAL_FALS)
	}
}
func WordFromInt(i uint32) Word  { return Word(VAL_NUMB | uint64(i)) }
func WordFromRune(r rune) Word   { return Word(VAL_RUNE | uint64(r)) }
func WordFromSym(a int) Word     { return Word(VAL_SYMB | uint64(a)) }
func WordFromAddress(a int) Word { return Word(VAL_ADDR | uint64(a)) }

func (w Word) isVal(v uint64) bool { return uint64(w)&MASK_ONE == v }

func (w Word) IsFloat() bool   { return uint64(w)&MASK_NAN != MASK_NAN }
func (w Word) IsAddress() bool { return uint64(w)&VAL_ADDR == VAL_ADDR }
func (w Word) IsTrue() bool    { return uint64(w) == VAL_TRUE }
func (w Word) IsFalse() bool   { return uint64(w) == VAL_FALS }
func (w Word) IsBool() bool    { return w.IsTrue() || w.IsFalse() }
func (w Word) IsEmpty() bool   { return uint64(w) == VAL_EMPT }
func (w Word) IsInteger() bool { return w.isVal(VAL_NUMB) }
func (w Word) IsRune() bool    { return w.isVal(VAL_RUNE) }
func (w Word) IsDate() bool    { return w.isVal(VAL_DATE) }
func (w Word) IsSym() bool     { return w.isVal(VAL_SYMB) }

func (w Word) IsSentinel() bool { return uint64(w) == SENTINEL }

func (w Word) isMark(m uint64) bool { return uint64(w)&MASK_ONE == m }

func (w Word) IsNAN() bool { return w.isMark(ERR_NUMB) }

func (w Word) IsAnyMark() bool     { return uint64(w)&MASK_ONE > MARK_MAP }
func (w Word) IsMapMark() bool     { return w.isMark(MARK_MAP) }
func (w Word) IsMainMapMark() bool { return w.isMark(MAIN_MAP) }
func (w Word) IsListMapMark() bool { return w.isMark(LIST_MAP) }
func (w Word) IsTextMapMark() bool { return w.isMark(TEXT_MAP) }
func (w Word) IsFuncMapMark() bool { return w.isMark(FUNC_MAP) }
func (w Word) IsMetaMapMark() bool { return w.isMark(META_MAP) }

func (w Word) IsArrayMark() bool     { return w.isMark(MARK_ARR) }
func (w Word) IsMainArrayMark() bool { return w.isMark(MAIN_ARR) }
func (w Word) IsRuneArrayMark() bool { return w.isMark(RUNE_ARR) }
func (w Word) IsCodeArrayMark() bool { return w.isMark(CODE_ARR) }
func (w Word) IsMetaArrayMark() bool { return w.isMark(META_ARR) }

func (w Word) IsLegendMark() bool     { return w.isMark(MARK_LGD) }
func (w Word) IsMainLegendMark() bool { return w.isMark(MAIN_LGD) }
func (w Word) IsListLegendMark() bool { return w.isMark(LIST_LGD) }
func (w Word) IsTextLegendMark() bool { return w.isMark(TEXT_LGD) }
func (w Word) IsFuncLegendMark() bool { return w.isMark(FUNC_LGD) }
func (w Word) IsMetaLegendMark() bool { return w.isMark(META_LGD) }

func (w Word) IsDescMark() bool        { return w.isMark(MARK_DES) }
func (w Word) IsIntegerDescMark() bool { return w.isMark(DES_NUMB) }
func (w Word) IsFloatDescMark() bool   { return w.isMark(DES_FLOA) }
func (w Word) IsBoolDescMark() bool    { return w.isMark(DES_BOOL) }
func (w Word) IsRuneDescMark() bool    { return w.isMark(DES_RUNE) }
func (w Word) IsSymbolDescMark() bool  { return w.isMark(DES_SYMB) }
func (w Word) IsDateDescMark() bool    { return w.isMark(DES_DATE) }
func (w Word) IsAddressDescMark() bool { return w.isMark(DES_ADDR) }

func (w Word) IsChunk() bool { return w.isMark(CMP_UNIT) }

func (w Word) Debug() any {
	if w.IsNAN() {
		return "NAN"
	} else if w.IsTrue() {
		return "TRUE"
	} else if w.IsFalse() {
		return "FALSE"
	} else if w.IsEmpty() {
		return "EMPTY"
	} else if w.IsInteger() {
		return fmt.Sprint("{INT: ", w.AsInt(), "}")
	} else if w.IsRune() {
		return fmt.Sprint("{RUNE: ", w.AsRune(), "}")
	} else if w.IsSym() {
		return "SYMBOL" // FIXME: Implement symbols
	} else if w.IsAddress() { // must trigger before IsFloat for unknown reasons
		return fmt.Sprint("{ADDR: ", w.AsAddr(), "}")
	} else if w.IsFloat() {
		return fmt.Sprint("{FLOAT: ", w.AsFloat(), "}")
	} else if w.IsMapMark() {
		return "MAP"
	} else if w.IsArrayMark() {
		return "ARRAY"
	} else if w.IsFuncMapMark() {
		return "ROUTINE"
	} else if w.IsChunk() {
		return "CHUNK"
	} else {
		return nil
	}
}

// func (w Word) isSwpT(t uint64) bool {
// 	return uint64(w)&(TAG_OBJ_MAP|TAG_SWPMASK) == t
// }
// func (w Word) IsGrey() bool  { return w.isSwpT(TAG_OBJ_MAP | TAG_SWP_GRY) }
// func (w Word) IsBlack() bool { return w.isSwpT(TAG_OBJ_MAP | TAG_SWP_BLK) }

func (w Word) AsBool() bool {
	switch uint64(w) {
	case VAL_TRUE:
		return true
	case VAL_FALS, VAL_EMPT, ERR_NUMB:
		return false
	default:
		return w.IsAddress()
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
func (w Word) AsInt() uint32  { return uint32(uint64(w) & ^VAL_NUMB) }
func (w Word) AsRune() rune   { return rune(uint64(w) & ^VAL_RUNE) }
func (w Word) AsSym() uint64  { return uint64(w) & ^VAL_SYMB }
func (w Word) AsAddr() uint64 { return uint64(w) & ^VAL_ADDR }
