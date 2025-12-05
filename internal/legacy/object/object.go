package object

import "arena"

type Any interface {
	IsObject()
	WHAT() string
	Equals(Any) bool
}

type Hashable interface {
	Any
	HashBytes() []byte
}

type Map interface {
	Any
	Get() Any
	Set(*arena.Arena, Hashable, Any, bool)
}