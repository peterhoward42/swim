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
