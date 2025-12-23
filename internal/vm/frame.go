package vm

import mapval "github.com/rhumb-lang/rhi/internal/map"

type CallFrame struct {
	Parent  *CallFrame // Link to the caller (Cactus Stack)
	Closure *mapval.Closure
	IP      int
	Base    int             // Stack index where this frame's locals start
	Monitor *mapval.Closure // Attached Selector for Space ops
	ArgCount int // Number of arguments explicitly provided by the caller

	// WaitingSignal is the topic this frame is waiting for a reply to.
	// Empty if running.
	WaitingSignal string

	// LocalsFrozen tracks which locals are immutable.
	// Uses map of pointers to allow stable addressing for Upvalues.
	// Index i corresponds to Base + i.
	LocalsFrozen map[int]*bool
}
