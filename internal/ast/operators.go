package ast

import "strings"

// Operator types
type OpType int

const (
	// Math
	OpAdd      OpType = iota // ++ (Add / Concat)
	OpSub                    // -- (Subtract)
	OpMult                   // ** (Multiply)
	OpPow                    // ^ ^ (Power)
	OpDivFloat               // // (Float Division)
	OpDivInt                 // +/ (Integer Division)
	OpMod                    // -/ (Modulo)
	OpSciNot                 // *^ (Scientific Notation)
	OpRoot                   // ^/ (Root)
	OpDev                    // +- (Deviation)

	// Logic
	OpEq  // == (Equal)
	OpNeq // ~~ (Not Equal)
	OpGt  // >> (Greater Than)
	OpLt  // << (Less Than)
	OpGte // >= (Greater Than or Equal)
	OpLte // <= (Less Than or Equal)
	OpAnd // /\ (AND)
	OpOr  // \/ (OR)

	// Structure
	OpRange     // | (Range)
	OpHasSub    // =@ (Has Subfield)
	OpNotHasSub // ~@ (Not Has Subfield)
	OpHasFld    // =\ (Has Field)
	OpNotHasFld // ~\ (Not Has Field)
	OpTempSub   // @@ (Temporary Subfield)
	OpConcat    // && (Concatenate)
	OpNested    // \\ (Nested Access)

	// Access / Postfix
	OpIndex     // [] (Index / Access)
	OpLength    // [#] (Length)
	OpIsEmpty   // [?] (Is Empty)
	OpAllSub    // [@] (All Subfields)
	OpAllFld    // [*] (All Public Fields)
	OpAllPos    // [0] (All Positional Elements)
	OpFreeze    // [.] (Freeze / Make Immutable)
	OpCopy      // [:] (Copy)
	OpToDate    // [/] (To Date)
	OpGetParams // [$] (Get Parameters)
	OpGetCtor   // [^] (Get Constructor)
	OpGetBase   // [!] (Get Base)
	OpToNum     // [+] (To Number)
	OpNegNum    // [-] (Negate Number)
	OpToBool    // [=] (To Boolean)
	OpNegBool   // [~] (Negate Boolean)
	OpSpread    // [&] (Spread)
	OpToKey     // [`] (To Key)
	OpAppend    // >[] (Append)
	OpUnshift   // [<] (Unshift)

	// Control/Flow
	OpAssignImm // .= (Immutable Assignment)
	OpAssignMut // := (Mutable Assignment)
	OpDestruct  // ^= (Destructuring Assignment)
	OpIfTrue    // => (If True)
	OpIfFalse   // ~> (If False)
	OpWhile     // |> (While Loop)
	OpForeach   // <> (Foreach Loop)
	OpPipe      // || (Pipe)
	OpCoalesce  // ?? (Coalesce / Default)
	OpMatchCons // .. (Match & Consume)
	OpMatchPeek // :: (Match & Peek)

	// Function
	OpMakeFn // -> (Make Function)
	OpBindFn // !> (Bind Function)
	OpLetFn  // +> (Let Function / IIFE)
	OpRebind // !! (Rebind)
)

// BinaryExpression represents "Left Op Right".
type BinaryExpression struct {
	Left  Expression
	Op    OpType
	Right Expression
}

func (b *BinaryExpression) expressionNode() {}
func (b *BinaryExpression) String() string {
	left := "NIL"
	if b.Left != nil {
		left = b.Left.String()
	}
	right := "NIL"
	if b.Right != nil {
		right = b.Right.String()
	}
	return "(" + left + " " + opToString(b.Op) + " " + right + ")"
}

// UnaryExpression represents "Op Expression" (Prefix).
type UnaryExpression struct {
	Op   OpType
	Expr Expression
}

func (u *UnaryExpression) expressionNode() {}
func (u *UnaryExpression) String() string {
	return "(" + opToString(u.Op) + u.Expr.String() + ")"
}

func opToString(op OpType) string {
	switch op {
	case OpAdd:
		return "++"
	case OpSub:
		return "--"
	case OpMult:
		return "**"
	case OpPow:
		return "^ ^"
	case OpDivFloat:
		return "//"
	case OpDivInt:
		return "+/"
	case OpMod:
		return "-/"
	case OpSciNot:
		return "*^"
	case OpRoot:
		return "^/"
	case OpDev:
		return "+-"
	case OpEq:
		return "=="
	case OpNeq:
		return "~~"
	case OpGt:
		return ">>"
	case OpLt:
		return "<<"
	case OpGte:
		return ">="
	case OpLte:
		return "<="
	case OpAnd:
		return "/\\"
	case OpOr:
		return "\\/"
	case OpRange:
		return "|"
	case OpHasSub:
		return "=@"
	case OpNotHasSub:
		return "~@"
	case OpHasFld:
		return "=\""
	case OpNotHasFld:
		return "~\""
	case OpTempSub:
		return "@@"
	case OpConcat:
		return "&&"
	case OpNested:
		return "\\\\"
	case OpIndex:
		return "[]"
	case OpLength:
		return "[#]"
	case OpIsEmpty:
		return "[?]"
	case OpAllSub:
		return "[@]"
	case OpAllFld:
		return "[* ]"
	case OpAllPos:
		return "[0]"
	case OpFreeze:
		return "[.]"
	case OpCopy:
		return "[: ]"
	case OpToDate:
		return "[/]"
	case OpGetParams:
		return "[$]"
	case OpGetCtor:
		return "[^]"
	case OpGetBase:
		return "[! ]"
	case OpToNum:
		return "[+]"
	case OpNegNum:
		return "[-]"
	case OpToBool:
		return "[=]"
	case OpNegBool:
		return "[~]"
	case OpSpread:
		return "[&]"
	case OpToKey:
		return "[`]"
	case OpAppend:
		return ">[ ]"
	case OpUnshift:
		return "[<]"
	case OpAssignImm:
		return ".="
	case OpAssignMut:
		return ":="
	case OpDestruct:
		return "^="
	case OpIfTrue:
		return "=>"
	case OpIfFalse:
		return "~>"
	case OpWhile:
		return "| >"
	case OpForeach:
		return "<>"
	case OpPipe:
		return "||"
	case OpCoalesce:
		return "??"
	case OpMatchCons:
		return ".."
	case OpMatchPeek:
		return "::"
	case OpMakeFn:
		return "->"
	case OpBindFn:
		return "!>"
	case OpLetFn:
		return "+>"
	case OpRebind:
		return "!!"
	default:
		return "?"
	}
}

// ChainOpType defines the type of access in a chain.
type ChainOpType int

const (
	ChainMember       ChainOpType = iota // \
	ChainSubfield                        // @
	ChainSignal                          // #
	ChainReply                           // ^
	ChainProclamation                    // $
)

// ChainStep represents one step in a chain expression (e.g., `\x`).
type ChainStep struct {
	Op    ChainOpType
	Ident string // The identifier (label)
}

// ChainExpression represents a chain of accesses like `user\name` or `worker#start`.
type ChainExpression struct {
	Base  Expression // The starting point (e.g., `user`)
	Steps []ChainStep
}

func (c *ChainExpression) expressionNode() {}
func (c *ChainExpression) String() string {
	var sb strings.Builder
	sb.WriteString(c.Base.String())
	for _, step := range c.Steps {
		switch step.Op {
		case ChainMember:
			sb.WriteString("\\")
		case ChainSubfield:
			sb.WriteString("@")
		case ChainSignal:
			sb.WriteString("#")
		case ChainReply:
			sb.WriteString("^")
		case ChainProclamation:
			sb.WriteString("$")
		}
		sb.WriteString(step.Ident)
	}
	return sb.String()
}
