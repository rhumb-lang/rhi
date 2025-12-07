package runner

import (
	"git.sr.ht/~madcapjake/rhi/internal/generator"
	"git.sr.ht/~madcapjake/rhi/internal/parser"
	"git.sr.ht/~madcapjake/rhi/internal/vm"
	"github.com/antlr4-go/antlr/v4"
)

func RunString(code string, v *vm.VirtualMachine) (any, error) {
	input := antlr.NewInputStream(code)
	lexer := parser.NewRhumbLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRhumbParser(stream)
	p.AddErrorListener(new(generator.RhumbErrorListener))
	p.BuildParseTrees = true
	tree := p.Expressions()
	visitor := generator.NewRhumbVisitor(v)
	if _, err := visitor.Visit(tree); err != nil {
		return nil, err
	}

	return v.RunMain()
}
