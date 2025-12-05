package mapval

// FieldKind distinguishes between immutable (.) and mutable (:) fields.
type FieldKind uint8

const (
	FieldImmutable FieldKind = iota // Defined via '.'
	FieldMutable                    // Defined via ':'
)

// FieldDesc describes a single field in a Legend.
type FieldDesc struct {
	Name string    // The label/key of the field
	Kind FieldKind // Mutable or Immutable
}

// TransitionDesc represents a transition in the hidden class tree.
type TransitionDesc struct {
	Key       string // The field name that triggers this transition
	NewLegend *Legend
}

// LegendType distinguishes between Map (Linear), Dictionary (Hash), etc.
type LegendType uint8

const (
	LegendMap        LegendType = iota // Linear scan
	LegendDictionary                   // Hash map
)

// Legend represents the Schema of a Map (Hidden Class).
type Legend struct {
	Kind        LegendType        // Map, Dictionary
	Fields      []FieldDesc       // Linear scan (Map Mode)
	Lookup      map[string]int    // Hash map (Dictionary Mode) for O(1) offset lookup
	Transitions []TransitionDesc  // Hidden Class transition tree
}

// NewLegend creates an empty Legend.
func NewLegend() *Legend {
	return &Legend{
		Kind:   LegendMap,
		Fields: make([]FieldDesc, 0),
	}
}
