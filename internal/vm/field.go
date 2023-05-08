package vm

import "git.sr.ht/~madcapjake/grhumb/internal/object"

type Field struct {
	Name     string
	Needs    object.Ref[Map]
	Provides object.Ref[Map]
	Value    interface{}
}
