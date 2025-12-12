package vm

import mapval "git.sr.ht/~madcapjake/rhi/internal/map"

type CallFrame struct {
	Parent  *CallFrame // Link to the caller (Cactus Stack)
	Closure *mapval.Closure
	IP      int
	Base    int             // Stack index where this frame's locals start
	Monitor *mapval.Closure // Attached Selector for Space ops
}
