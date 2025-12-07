package code

import (
	"fmt"

	"git.sr.ht/~madcapjake/rhi/internal/color"
)

type Any interface {
	WHAT() string
	GetLine() int
	GetColumn() int
	GetID() int
}

// Rhumb literals such as numbers, dates, texts, and labels.
//
// To store a value in a label, you must provide the label as a value
// to the assignment operators. The left-hand side of the asssignment
// operator is auto-referenced.
type Value struct{ line, column, id int }

func NewValue(l, c, id int) Value {
	return Value{l, c, id}
}
func (c Value) WHAT() string {
	return fmt.Sprint("code.Value ", color.Yellow, "#", c.id, color.Reset)
}
func (c Value) GetLine() int   { return c.line }
func (c Value) GetColumn() int { return c.column }
func (c Value) GetID() int     { return c.id }

type Local struct{ line, column, id int }

func NewLocal(l, c, id int) Local {
	return Local{l, c, id}
}
func (c Local) WHAT() string {
	return fmt.Sprint("code.Local ", color.Yellow, "#", c.id, color.Reset)
}
func (c Local) GetLine() int   { return c.line }
func (c Local) GetColumn() int { return c.column }
func (c Local) GetID() int     { return c.id }

type Inner struct{ line, column, id int }

func NewInner(l, c, id int) Inner {
	return Inner{l, c, id}
}
func (c Inner) WHAT() string {
	return fmt.Sprint("code.Inner ", color.Yellow, "#", c.id, color.Reset)
}
func (c Inner) GetLine() int   { return c.line }
func (c Inner) GetColumn() int { return c.column }
func (c Inner) GetID() int     { return c.id }

type Under struct{ line, column, id int }

func NewUnder(l, c, id int) Under {
	return Under{l, c, id}
}
func (c Under) WHAT() string {
	return fmt.Sprint("code.Under ", color.Yellow, "#", c.id, color.Reset)
}
func (c Under) GetLine() int   { return c.line }
func (c Under) GetColumn() int { return c.column }
func (c Under) GetID() int     { return c.id }

type Outer struct{ line, column, id int }

func NewOuter(l, c, id int) Outer {
	return Outer{l, c, id}
}
func (c Outer) WHAT() string {
	return fmt.Sprint("code.Outer ", color.Yellow, "#", c.id, color.Reset)
}
func (c Outer) GetLine() int   { return c.line }
func (c Outer) GetColumn() int { return c.column }
func (c Outer) GetID() int     { return c.id }

type Event struct{ line, column, id int }

func NewEvent(l, c, id int) Event {
	return Event{l, c, id}
}
func (c Event) WHAT() string {
	return fmt.Sprint("code.Event ", color.Yellow, "#", c.id, color.Reset)
}
func (c Event) GetLine() int   { return c.line }
func (c Event) GetColumn() int { return c.column }
func (c Event) GetID() int     { return c.id }

type Reply struct{ line, column, id int }

func NewReply(l, c, id int) Reply {
	return Reply{l, c, id}
}
func (c Reply) WHAT() string {
	return fmt.Sprint("code.Reply ", color.Yellow, "#", c.id, color.Reset)
}
func (c Reply) GetLine() int   { return c.line }
func (c Reply) GetColumn() int { return c.column }
func (c Reply) GetID() int     { return c.id }
