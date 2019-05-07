/*
Package graphics encapsulates the types that are needed to capture a diagram
in terms of low level primites (like Lines and Labels).

The aim is that these be easily renderable into diverse graphics
formats, or serialized into JSON or YAML.

The size units used are pixels.
*/
package graphics


// Model is the topl-level model.
type Model struct {
	Width      int
	FontHeight float64
	Primitives *Primitives
}

// NewModel instantiates a Model and initializes it ready to use.
func NewModel(width int, fontHeight float64) *Model {
	return &Model{width, fontHeight, NewPrimitives()}
}
