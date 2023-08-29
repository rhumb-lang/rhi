package object

import (
	"arena"
)

type Block struct {
	Values []Any
	Codes  []Code
}

func NewBlock(a *arena.Arena) Block {
	return Block{
		Values: arena.MakeSlice[Any](a, 0, 5),
		Codes:  arena.MakeSlice[Code](a, 0, 50),
	}
}

func (block Block) IsObject() {}

func (block Block) WHAT() string { return "Block: ..." }

func (x Block) Equals(y Any) bool { return false }

func (block *Block) Register(o Any) int {
	for i, existing := range block.Values {
		if o.Equals(existing) {
			return i
		}
	}
	block.Values = append(block.Values, o)
	return len(block.Values) - 1
}

func (block *Block) Write(c Code) {
	block.Codes = append(block.Codes, c)
}
