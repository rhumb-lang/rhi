package mapval

import "fmt"

type OpCode byte

const (
	// --- BANK 0: Control Flow & Selector ---
	OP_HALT OpCode = iota
	OP_JUMP             // JUMP <offset>
	OP_DUP              // Duplicate top of stack
	OP_POP              // Pop top of stack
	OP_CALL             // CALL <arg_count>
	OP_RETURN           // RETURN
	OP_MAKE_FN          // MAKE_FN <const_idx>
	OP_SELECT           // SELECT (Start Pattern Match)
	OP_MATCH_STRUCT     // Match structural pattern

	
	// --- BANK 1: Lexical & Scope ---
	OP_LOAD_CONST       // LOAD_CONST <index>
	OP_LOAD_LOC         // LOAD_LOC <slot>
	OP_STORE_LOC        // STORE_LOC <slot>
	OP_LOAD_UPVALUE     // LOAD_UPVALUE <index>
	OP_STORE_UPVALUE    // STORE_UPVALUE <index>
	OP_LOAD_STATIC      // LOAD_STATIC <index> (Module level)
	OP_MATCH_BIND       // Bind match to var

	// --- BANK 2: Map & Inheritance ---
	OP_SEND             // SEND <key_const_index> (Get Field)
	OP_SET_FIELD        // SET_FIELD <key_const_index> <flags> (Set Field)
	OP_SELF             // Push 'self' (!)
	OP_LOAD_PARENT      // Push 'parent' (@)
	OP_MAKE_MAP         // Create empty map

	// --- BANK 3: Space & Concurrency ---
	OP_POST             // POST #signal
	OP_INJECT           // INJECT ^reply
	OP_WRITE            // WRITE $proclamation
	OP_SUBSCRIBE        // SUBSCRIBE <>
	OP_NEW_REALM        // NEW_REALM <flags>
	OP_MONITOR          // MONITOR (Attach selector to call)
	OP_MATCH_TUPLE      // MATCH_TUPLE <kind> <topic_idx>

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
	case OP_HALT: return "OP_HALT"
    // ...
    case OP_ASSERT_EQ: return "OP_ASSERT_EQ"
    case OP_INSPECT: return "OP_INSPECT"
	default: return fmt.Sprintf("OP_UNKNOWN(%d)", op)
	}
}
