package stack

import (
	"arena"
)

type Stack[O any] struct {
	objects []O
}

func NewStack[O any](a *arena.Arena) (stack *Stack[O]) {
	stack = arena.New[Stack[O]](a)
	*stack = Stack[O]{
		objects: arena.MakeSlice[O](a, 0, 100),
	}
	return
}

func (s *Stack[O]) Push(o O) {
	s.objects = append(s.objects, o)
}

func (s *Stack[O]) Top() (value *O) {
	last := len(s.objects) - 1
	if last >= 0 {
		value = &s.objects[last]
	}
	return
}

func (s *Stack[O]) Pop() (value O) {
	last := len(s.objects) - 1
	if last >= 0 {
		value = s.objects[last]
		s.objects = s.objects[:last]
	}
	return
}

func (s Stack[O]) Size() (sz int) {
	sz = len(s.objects)
	return
}
