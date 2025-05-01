package object

import (
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/color"
)

// Operators, variables and constants all start as label literals
// The bytecode decides how they will be utilized
type Label struct{ Value string }

func (l Label) IsObject() {}
func (l Label) WHAT() string {
	return fmt.Sprint("Label:", color.Green, "\"", l.Value, "\"", color.Reset)
}
func (x Label) Equals(y Any) bool {
	objY, ok := y.(Label)
	return ok && x.Value == objY.Value
}
func (key Label) HashBytes() []byte {
	return []byte(key.Value)
}
