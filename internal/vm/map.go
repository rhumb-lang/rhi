package vm

type Map struct {
	Mark   Word
	Legend Word // Address
	Field  []Word
}

func NewMap(count uint32, legAddr Word) Map {
	return Map{
		Word(MAIN_MAP),
		WordFromAddress(0),
		make([]Word, count),
	}
}

type MapLegend struct {
	Mark       Word
	MetaLegend Word
	TrashSweep Word
	Length     Word
	Count      Word
	PrevLegend Word // pointer, circular dependency list
	NextLegend Word // pointer, circular dependency list
	Field      []LegendFieldDescriptor
}

func NewMapLegend() MapLegend {
	return MapLegend{
		Word(MAIN_LGD),
		WordFromAddress(0),
		EmptyWord(),
		WordFromInt(0),
		WordFromInt(0),
		WordFromAddress(0),
		WordFromAddress(0),
		make([]LegendFieldDescriptor, 0),
	}
}

type RoutineLegend struct {
	Mark       Word
	MetaLegend Word // pointer to the toplevel MapLegend
	TrashSweep Word
	Length     Word
	Count      Word
	PrevLegend Word // pointer, circular dependency list
	NextLegend Word // pointer, circular dependency list
	Code       Word // pointer to CodeArray
	Field      []LegendFieldDescriptor
}

type ArrayLegend struct {
	Mark       Word
	MetaLegend Word
	TrashSweep Word
	Parent     LegendFieldDescriptor
}

// T

type LegendFieldDescriptor struct {
	Mark Word // immutable, mutable, subfield
	Name Word // address to TextMap
	Data Word // constant, field offset, or
}
