package vm

import mapval "github.com/rhumb-lang/rhi/internal/map"

type CallFrame struct {
	Parent  *CallFrame // Link to the caller (Cactus Stack)
	Closure *mapval.Closure
	IP      int
	Base    int             // Stack index where this frame's locals start
	Monitor *mapval.Closure // Attached Selector for Space ops
}
