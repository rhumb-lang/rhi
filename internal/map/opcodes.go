package mapval

import "fmt"

type OpCode byte

const (
	// --- BANK 0: Control Flow & Selector ---
	OP_HALT         OpCode = iota
	OP_JUMP                // JUMP <offset>
	OP_DUP                 // Duplicate top of stack
	OP_POP                 // Pop top of stack
	OP_CALL                // CALL <arg_count>
	OP_RETURN              // RETURN
	OP_MAKE_FN             // MAKE_FN <const_idx>
	OP_SELECT              // SELECT (Start Pattern Match)
	OP_MATCH_STRUCT        // Match structural pattern

	// --- BANK 1: Lexical & Scope ---
	OP_LOAD_CONST    // LOAD_CONST <index>
	OP_LOAD_LOC      // LOAD_LOC <slot>
	OP_STORE_LOC     // STORE_LOC <slot>
	OP_STORE_LOC_IMMUT // STORE_LOC_IMMUT <slot> (Freeze)
	OP_LOAD_UPVALUE  // LOAD_UPVALUE <index>
	OP_STORE_UPVALUE // STORE_UPVALUE <index>
	OP_LOAD_STATIC   // LOAD_STATIC <index> (Module level)
	OP_MATCH_BIND    // Bind match to var
	OP_RESOLVE       // Library Resolution

	// --- BANK 2: Map & Inheritance ---
	OP_SEND        // SEND <key_const_index> (Get Field)
	OP_SET_FIELD   // SET_FIELD <key_const_index> <flags> (Set Field)
	OP_SELF        // Push 'self' (!)
	OP_LOAD_PARENT // Push 'parent' (@)
	OP_MAKE_MAP    // Create empty map

	// --- BANK 3: Space & Concurrency ---
	OP_POST        // POST #signal
	OP_INJECT      // INJECT ^reply
	OP_WRITE       // WRITE $proclamation
	OP_SUBSCRIBE   // SUBSCRIBE <>
	OP_NEW_REALM   // NEW_REALM <flags>
	OP_MONITOR     // MONITOR (Attach selector to call)
	OP_MATCH_TUPLE // MATCH_TUPLE <kind> <topic_idx>

	// --- NATIVE INTRINSICS (Operators) ---

	// Function
	OP_LET_FN
	OP_BIND_FN
	OP_REBIND
	OP_REF_FN
	OP_CURRY

	// Math
	OP_ADD
	OP_SUB
	OP_MULT
	OP_POW
	OP_DIV_FLOAT
	OP_DIV_INT
	OP_MOD
	OP_SCI_NOT
	OP_ROOT
	OP_DEV // Deviation (+-)

	// Logic
	OP_EQ
	OP_NEQ
	OP_GT
	OP_LT
	OP_GTE
	OP_LTE
	OP_AND
	OP_OR
	OP_NOT // Logical Negate (Prefix)

	// Map & Structure
	OP_RANGE
	OP_HAS_SUBFIELD
	OP_NOT_HAS_SUB
	OP_HAS_FIELD
	OP_NOT_HAS_FLD
	OP_TEMP_SUBFIELD
	OP_CONCAT
	OP_ACCESS_NESTED

	// Control Flow (Operator Forms)
	OP_ASSIGN_IMM
	OP_ASSIGN_MUT
	OP_DESTRUCT
	OP_JUMP_IF_FALSE
	OP_JUMP_IF_TRUE
	OP_WHILE
	OP_FOREACH
	OP_PIPE
	OP_COALESCE
	OP_MATCH_CONS
	OP_MATCH_PEEK

	// Field Operators (Postfix)
	OP_APPEND
	OP_UNSHIFT
	OP_LENGTH
	OP_IS_EMPTY
	OP_IS_MAP
	OP_ALL_SUB
	OP_ALL_FIELDS
	OP_ALL_POS
	OP_FREEZE
	OP_COPY
	OP_COERCE_DATE
	OP_GET_PARAMS
	OP_GET_CTOR
	OP_GET_BASE
	OP_COERCE_NUM
	OP_NUM_NEG
	OP_COERCE_BOOL
	OP_BOOL_NEG
	OP_SPREAD
	OP_COERCE_KEY

	// Testing
	OP_ASSERT_EQ
	OP_INSPECT
)

func (op OpCode) String() string {
	switch op {
	case OP_HALT:
		return "OP_HALT"
	case OP_JUMP:
		return "OP_JUMP"
	case OP_DUP:
		return "OP_DUP"
	case OP_POP:
		return "OP_POP"
	case OP_CALL:
		return "OP_CALL"
	case OP_RETURN:
		return "OP_RETURN"
	case OP_MAKE_FN:
		return "OP_MAKE_FN"
	case OP_SELECT:
		return "OP_SELECT"
	case OP_MATCH_STRUCT:
		return "OP_MATCH_STRUCT"
	case OP_LOAD_CONST:
		return "OP_LOAD_CONST"
	case OP_LOAD_LOC:
		return "OP_LOAD_LOC"
	case OP_STORE_LOC:
		return "OP_STORE_LOC"
	case OP_STORE_LOC_IMMUT:
		return "OP_STORE_LOC_IMMUT"
	case OP_LOAD_UPVALUE:
		return "OP_LOAD_UPVALUE"
	case OP_STORE_UPVALUE:
		return "OP_STORE_UPVALUE"
	case OP_LOAD_STATIC:
		return "OP_LOAD_STATIC"
	case OP_MATCH_BIND:
		return "OP_MATCH_BIND"
	case OP_RESOLVE:
		return "OP_RESOLVE"
	case OP_SEND:
		return "OP_SEND"
	case OP_SET_FIELD:
		return "OP_SET_FIELD"
	case OP_SELF:
		return "OP_SELF"
	case OP_LOAD_PARENT:
		return "OP_LOAD_PARENT"
	case OP_MAKE_MAP:
		return "OP_MAKE_MAP"
	case OP_POST:
		return "OP_POST"
	case OP_INJECT:
		return "OP_INJECT"
	case OP_WRITE:
		return "OP_WRITE"
	case OP_SUBSCRIBE:
		return "OP_SUBSCRIBE"
	case OP_NEW_REALM:
		return "OP_NEW_REALM"
	case OP_MONITOR:
		return "OP_MONITOR"
	case OP_MATCH_TUPLE:
		return "OP_MATCH_TUPLE"
	case OP_LET_FN:
		return "OP_LET_FN"
	case OP_BIND_FN:
		return "OP_BIND_FN"
	case OP_REBIND:
		return "OP_REBIND"
	case OP_REF_FN:
		return "OP_REF_FN"
	case OP_CURRY:
		return "OP_CURRY"
	case OP_ADD:
		return "OP_ADD"
	case OP_SUB:
		return "OP_SUB"
	case OP_MULT:
		return "OP_MULT"
	case OP_POW:
		return "OP_POW"
	case OP_DIV_FLOAT:
		return "OP_DIV_FLOAT"
	case OP_DIV_INT:
		return "OP_DIV_INT"
	case OP_MOD:
		return "OP_MOD"
	case OP_SCI_NOT:
		return "OP_SCI_NOT"
	case OP_ROOT:
		return "OP_ROOT"
	case OP_DEV:
		return "OP_DEV"
	case OP_EQ:
		return "OP_EQ"
	case OP_NEQ:
		return "OP_NEQ"
	case OP_GT:
		return "OP_GT"
	case OP_LT:
		return "OP_LT"
	case OP_GTE:
		return "OP_GTE"
	case OP_LTE:
		return "OP_LTE"
	case OP_AND:
		return "OP_AND"
	case OP_OR:
		return "OP_OR"
	case OP_NOT:
		return "OP_NOT"
	case OP_RANGE:
		return "OP_RANGE"
	case OP_HAS_SUBFIELD:
		return "OP_HAS_SUBFIELD"
	case OP_NOT_HAS_SUB:
		return "OP_NOT_HAS_SUB"
	case OP_HAS_FIELD:
		return "OP_HAS_FIELD"
	case OP_NOT_HAS_FLD:
		return "OP_NOT_HAS_FLD"
	case OP_TEMP_SUBFIELD:
		return "OP_TEMP_SUBFIELD"
	case OP_CONCAT:
		return "OP_CONCAT"
	case OP_ACCESS_NESTED:
		return "OP_ACCESS_NESTED"
	case OP_ASSIGN_IMM:
		return "OP_ASSIGN_IMM"
	case OP_ASSIGN_MUT:
		return "OP_ASSIGN_MUT"
	case OP_DESTRUCT:
		return "OP_DESTRUCT"
	case OP_JUMP_IF_FALSE:
		return "OP_JUMP_IF_FALSE"
	case OP_JUMP_IF_TRUE:
		return "OP_JUMP_IF_TRUE"
	case OP_WHILE:
		return "OP_WHILE"
	case OP_FOREACH:
		return "OP_FOREACH"
	case OP_PIPE:
		return "OP_PIPE"
	case OP_COALESCE:
		return "OP_COALESCE"
	case OP_MATCH_CONS:
		return "OP_MATCH_CONS"
	case OP_MATCH_PEEK:
		return "OP_MATCH_PEEK"
	case OP_APPEND:
		return "OP_APPEND"
	case OP_UNSHIFT:
		return "OP_UNSHIFT"
	case OP_LENGTH:
		return "OP_LENGTH"
	case OP_IS_EMPTY:
		return "OP_IS_EMPTY"
	case OP_IS_MAP:
		return "OP_IS_MAP"
	case OP_ALL_SUB:
		return "OP_ALL_SUB"
	case OP_ALL_FIELDS:
		return "OP_ALL_FIELDS"
	case OP_ALL_POS:
		return "OP_ALL_POS"
	case OP_FREEZE:
		return "OP_FREEZE"
	case OP_COPY:
		return "OP_COPY"
	case OP_COERCE_DATE:
		return "OP_COERCE_DATE"
	case OP_GET_PARAMS:
		return "OP_GET_PARAMS"
	case OP_GET_CTOR:
		return "OP_GET_CTOR"
	case OP_GET_BASE:
		return "OP_GET_BASE"
	case OP_COERCE_NUM:
		return "OP_COERCE_NUM"
	case OP_NUM_NEG:
		return "OP_NUM_NEG"
	case OP_COERCE_BOOL:
		return "OP_COERCE_BOOL"
	case OP_BOOL_NEG:
		return "OP_BOOL_NEG"
	case OP_SPREAD:
		return "OP_SPREAD"
	case OP_COERCE_KEY:
		return "OP_COERCE_KEY"
	case OP_ASSERT_EQ:
		return "OP_ASSERT_EQ"
	case OP_INSPECT:
		return "OP_INSPECT"
	default:
		return fmt.Sprintf("OP_UNKNOWN(%d)", op)
	}
}
