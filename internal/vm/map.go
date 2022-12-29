package vm

type Map struct {
	Mark   Word
	Legend Word // Address
	Field  []Word
}

func NewMap() Map {
	return Map{
		Word(TAG_OBJ_MAP),
		Word(TAG_ADDRESS), // FIXME: Empty Address
		make([]Word, 0),
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

type LegendFieldDescriptor struct {
	Name Word
	Type Word // immutable, mutable, subfield
	Data Word // constant, field offset, or
}
