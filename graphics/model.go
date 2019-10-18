package graphics

// Model is the topl-level graphics model.
type Model struct {
	Width           float64
	Height          float64 // Remains at zero-value at contstruction time.
	FontHeight      float64
	DashLineDashLen float64
	DashLineGapLen  float64
	Primitives      *Primitives
}

// NewModel instantiates a Model and initializes it ready to use.
func NewModel(width float64, fontHeight float64,
	dashLineDashLen float64, dashLineGapLen float64) *Model {
	model := &Model{}
	model.Width = width
	model.FontHeight = fontHeight
	model.DashLineDashLen = dashLineDashLen
	model.DashLineGapLen = dashLineGapLen
	model.Primitives = NewPrimitives()

	return model
}
