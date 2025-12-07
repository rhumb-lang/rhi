package compiler

// Local represents a local variable on the stack.
type Local struct {
	Name  string
	Depth int // Scope depth (0 = global/module, 1 = function, etc.)
	IsUpvalue bool // Captured?
}

// CompilerScope tracks locals for a single function/script.
type CompilerScope struct {
	Locals    []Local
	ScopeDepth int
}

func NewCompilerScope() *CompilerScope {
	return &CompilerScope{
		Locals:    make([]Local, 0),
		ScopeDepth: 0,
	}
}

func (s *CompilerScope) resolveLocal(name string) int {
	// Search from top of stack down (shadowing)
	for i := len(s.Locals) - 1; i >= 0; i-- {
		if s.Locals[i].Name == name {
			return i
		}
	}
	return -1
}

func (s *CompilerScope) addLocal(name string) int {
	local := Local{
		Name:  name,
		Depth: s.ScopeDepth,
	}
	s.Locals = append(s.Locals, local)
	return len(s.Locals) - 1
}
