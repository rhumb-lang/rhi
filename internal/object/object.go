package object

import "fmt"

type Any interface {
	IsObject()
	WHAT() string
}

type Ref[O Any] struct {
	Value int
}

func (r Ref[any]) IsObject() {}

type WholeNumber struct{ Value int64 }

func (w WholeNumber) IsObject() {}
func (w WholeNumber) WHAT() string {
	return fmt.Sprint("WholeNumber: ", w.Value)
}

type RationalNumber struct{ Value float64 }

func (r RationalNumber) IsObject() {}
func (r RationalNumber) WHAT() string {
	return fmt.Sprint("RationalNumber: ", r.Value)
}

type VersionNumber struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
}

func (v VersionNumber) IsObject() {}
func (v VersionNumber) WHAT() string {
	return fmt.Sprint("VersionNumber: ", v.Major, v.Minor, v.Patch, "-", v.PreRelease, "+", v.Build)
}

type DateNumber struct{ Value int }

func (d DateNumber) IsObject() {}
func (d DateNumber) WHAT() string {
	return fmt.Sprint("DateNumber: ", d.Value)

}

type RuneSymbol struct{ Value rune }

func (r RuneSymbol) IsObject() {}
func (r RuneSymbol) WHAT() string {
	return fmt.Sprint("RuneSymbol: ", r.Value)

}

type KeySymbol struct {
	Value int
	Name  string
}

func (k KeySymbol) IsObject() {}
func (k KeySymbol) WHAT() string {
	return fmt.Sprint("KeySymbol: \"", k.Name, "\"")
}

type TruthSymbol struct{ Value bool }

func (t TruthSymbol) IsObject() {}
func (t TruthSymbol) WHAT() string {
	return fmt.Sprint("TruthSymbol: ", t.Value)
}

type LabelSymbol struct{ Value string }

func (l LabelSymbol) IsObject() {}
func (l LabelSymbol) WHAT() string {
	return fmt.Sprint("LabelSymbol: \"", l.Value, "\"")
}
