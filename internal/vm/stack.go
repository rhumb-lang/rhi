package vm

import (
	"arena"

	"git.sr.ht/~madcapjake/grhumb/internal/object"
)

type Stack struct {
	values []object.Any
}

func NewStack(a *arena.Arena) Stack {
	return Stack{values: arena.MakeSlice[object.Any](a, 0, 100)}
}

func (s *Stack) Push(o object.Any) {
	s.values = append(s.values, o)
}

func (s *Stack) Pop() (value object.Any) {
	last := len(s.values) - 1
	if last >= 0 {
		value = s.values[last]
		s.values = s.values[:last]
	}
	return
}
