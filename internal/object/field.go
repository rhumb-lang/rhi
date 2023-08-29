package object

type Field struct {
	Name     Hashable
	Demands  Ref[Map]
	Supplies Ref[Map]
	Mutable  bool
	Value    Any
}

func (f Field) Equals(key Hashable) bool {
	return f.Name.Equals(key)
}
