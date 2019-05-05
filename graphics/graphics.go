/*
Package graphics encapsulates the types that are needed to capture a diagram
in terms of low level primites (like Lines and Labels).

The aim is that these be easily renderable into diverse graphics
formats, or serialized into JSON or YAML.

The size units used are pixels.
*/
package graphics

// Line represents a line, optionally dashed, and optionally with an arrow at
// the (X2, Y2) end.
type Line struct {
	X1, X2, Y1, Y2 int
	Arrow          bool
	Dashed         bool // vs. Full
}

// The values that may be used in label justification.
const (
	Left   = "Left"
	Right  = "Right"
	Top    = "Top"
	Bottom = "Bottom"
	Centre = "Centre"
)

// Label encapsulates a (potentially) multi-line label in terms of a position,
// justification and its consituent lines of text.
type Label struct {
	LinesOfText []string
	// Anchor point about which the justifications are applied
	X, Y  int
	HJust string
	VJust string
}

// Primitives is a container for a set of Line(s) and a set of Label(s).
type Primitives struct {
	Lines  []*Line
	Labels []*Label
}

// NewPrimitives constructs a Primitives ready to use.
func NewPrimitives() *Primitives {
	return &Primitives{[]*Line{}, []*Label{}}
}

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

// Append adds the Primitives given to those already held in the model.
func (m *Model) Append(prims *Primitives) {
	m.Primitives.Lines = append(m.Primitives.Lines, prims.Lines...)
	m.Primitives.Labels = append(m.Primitives.Labels, prims.Labels...)
}
