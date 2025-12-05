package code

const (
	CB_HEIGHT int = 50
	CB_WEIGHT int = 5000
)

type CodeBuilder struct {
	height, length int
	Scope          [CB_HEIGHT][CB_WEIGHT]Any
}

func NewCodeBuilder() CodeBuilder {
	return CodeBuilder{
		height: 0,
		length: 0,
		Scope:  [50][5000]Any{},
	}
}

func (cb CodeBuilder) scopeWith(code Any) [CB_HEIGHT][CB_WEIGHT]Any {
	cb.Scope[cb.height][cb.length] = code
	return cb.Scope
}

func (cb CodeBuilder) Write(code Any) CodeBuilder {
	return CodeBuilder{
		height: cb.height,
		length: cb.length + 1,
		Scope:  cb.scopeWith(code),
	}
}

func (cb CodeBuilder) Enter() CodeBuilder {
	return CodeBuilder{
		height: cb.height + 1,
		length: cb.length,
		Scope:  cb.Scope,
	}
}

func (cb CodeBuilder) Exit() CodeBuilder {
	return CodeBuilder{
		height: cb.height - 1,
		length: cb.length,
		Scope:  cb.Scope,
	}
}
