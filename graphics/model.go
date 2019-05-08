package graphics

// Model is the topl-level model.
type Model struct {
	Width      int
	Height    int
	FontHeight float64
	Primitives *Primitives
}

// NewModel instantiates a Model and initializes it ready to use.
func NewModel(width int, fontHeight float64) *Model {
	const defaultHeight = 2000 // Can be changed or set any time.
	return &Model{width, defaultHeight, fontHeight, NewPrimitives()}
}
