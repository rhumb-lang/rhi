package code

import (
	"fmt"

	"git.sr.ht/~madcapjake/grhumb/internal/object"
)

type Any interface {
	WHAT() string
	GetLine() int
	SetLine(int)
	GetColumn() int
	SetColumn(int)
	GetTarget() object.Any
	SetTarget(object.Any)
}

type Value struct {
	line, column int
	target       object.Any
}

func NewValue(l int, c int, t object.Any) *Value {
	return &Value{l, c, t}
}
func (c Value) WHAT() string {
	return fmt.Sprint("code.Value: ", c.target.WHAT())
}
func (c Value) GetLine() int            { return c.line }
func (c *Value) SetLine(l int)          { c.line = l }
func (c Value) GetColumn() int          { return c.column }
func (c *Value) SetColumn(m int)        { c.column = m }
func (c Value) GetTarget() object.Any   { return c.target }
func (c *Value) SetTarget(t object.Any) { c.target = t }

type Local struct {
	line, column int
	target       object.Any
}

func NewLocal(l int, c int, t object.Any) *Local {
	return &Local{l, c, t}
}
func (c Local) WHAT() string {
	return fmt.Sprint("code.Local", c.target.WHAT())
}
func (c Local) GetLine() int            { return c.line }
func (c *Local) SetLine(l int)          { c.line = l }
func (c Local) GetColumn() int          { return c.column }
func (c *Local) SetColumn(m int)        { c.column = m }
func (c Local) GetTarget() object.Any   { return c.target }
func (c *Local) SetTarget(t object.Any) { c.target = t }

type Inner struct {
	line, column int
	target       object.Any
}

func NewInner(l int, c int, t object.Any) *Inner {
	return &Inner{l, c, t}
}
func (c Inner) WHAT() string {
	return fmt.Sprint("code.Inner", c.target.WHAT())
}
func (c Inner) GetLine() int            { return c.line }
func (c *Inner) SetLine(l int)          { c.line = l }
func (c Inner) GetColumn() int          { return c.column }
func (c *Inner) SetColumn(m int)        { c.column = m }
func (c Inner) GetTarget() object.Any   { return c.target }
func (c *Inner) SetTarget(t object.Any) { c.target = t }

type Under struct {
	line, column int
	target       object.Any
}

func NewUnder(l int, c int, t object.Any) *Under {
	return &Under{l, c, t}
}
func (c Under) WHAT() string {
	return fmt.Sprint("code.Under: ", c.target.WHAT())
}
func (c Under) GetLine() int            { return c.line }
func (c *Under) SetLine(l int)          { c.line = l }
func (c Under) GetColumn() int          { return c.column }
func (c *Under) SetColumn(m int)        { c.column = m }
func (c Under) GetTarget() object.Any   { return c.target }
func (c *Under) SetTarget(t object.Any) { c.target = t }

type Outer struct {
	line, column int
	target       object.Any
}

func NewOuter(l int, c int, t object.Any) *Outer {
	return &Outer{l, c, t}
}
func (c Outer) WHAT() string {
	return fmt.Sprint("code.Outer: ", c.target.WHAT())
}
func (c Outer) GetLine() int            { return c.line }
func (c *Outer) SetLine(l int)          { c.line = l }
func (c Outer) GetColumn() int          { return c.column }
func (c *Outer) SetColumn(m int)        { c.column = m }
func (c Outer) GetTarget() object.Any   { return c.target }
func (c *Outer) SetTarget(t object.Any) { c.target = t }

type Event struct {
	line, column int
	target       object.Any
}

func NewEvent(l int, c int, t object.Any) *Event {
	return &Event{l, c, t}
}
func (c Event) WHAT() string {
	return fmt.Sprint("code.Event: ", c.target.WHAT())
}
func (c Event) GetLine() int            { return c.line }
func (c *Event) SetLine(l int)          { c.line = l }
func (c Event) GetColumn() int          { return c.column }
func (c *Event) SetColumn(m int)        { c.column = m }
func (c Event) GetTarget() object.Any   { return c.target }
func (c *Event) SetTarget(t object.Any) { c.target = t }

type Reply struct {
	line, column int
	target       object.Any
}

func NewReply(l int, c int, t object.Any) *Reply {
	return &Reply{l, c, t}
}
func (c Reply) WHAT() string {
	return fmt.Sprint("code.Reply: ", c.target.WHAT())
}
func (c Reply) GetLine() int            { return c.line }
func (c *Reply) SetLine(l int)          { c.line = l }
func (c Reply) GetColumn() int          { return c.column }
func (c *Reply) SetColumn(m int)        { c.column = m }
func (c Reply) GetTarget() object.Any   { return c.target }
func (c *Reply) SetTarget(t object.Any) { c.target = t }
