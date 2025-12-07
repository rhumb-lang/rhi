package object

type Legend interface {
	IsLegend()
	WHAT() string
	Equals(Legend) bool
	DeepEquals(Legend) bool
}

type BaseLegend struct {
	Link     *Reference[Any]
	Demands  *Reference[Any]
	Supplies *Reference[Any]
}
