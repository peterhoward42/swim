/*
Package graphics encapsulates the vector-graphics output of the umlinteraction package.
It provides the Diagram type, comprising little more than a set of Line
objects, and Text objects.

The aim is that these be easily renderable into diverse graphics formats.

The size units used are abstract pixels.
*/
package graphics

// Model is the topl-level model.
type Model struct {
	Width  int
	Height int
	Lines  []*Line
	Labels []*Label
}

// Line represents a line, optionally with an arrow at the X2, Y2 end.
type Line struct {
	X1, X2, Y1, Y2 int
	Arrow          bool
	Dashed         bool // or Full
}

// Justification is an enumerated type for vertical or horizontal justifications.
type Justification int

// Constants to define members of the Justification types.
const (
	Left Justification = iota
	Right
	Top
	Bottom
	Centre
)

func (j Justification) String() string {
	return [...]string{"Left", "Right", "Top", "Bottom", "Centre"}[j]
}

// Label encapsulates a (potentially) multi-line label in terms of a position,
// justification and its consituent lines of text.
type Label struct {
	LinesOfText []string
	// Anchor point about which the justifications are applied
	X, Y  int
	HJust Justification
	VJust Justification
}
