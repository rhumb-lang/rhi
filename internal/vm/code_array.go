package vm

type Op int8

const (
	Return Op = iota
)

type CodeArray struct {
	at []uint8
}
