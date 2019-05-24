package graphics

// Point represents one point (X,Y)
type Point struct {
	X float64
	Y float64
}

// NewPoint creates a Point ready to use.
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Line represents a line, optionally dashed.
type Line struct {
	P1     *Point
	P2     *Point
	Dashed bool // vs. Full
}

// FilledPoly represents a filled polygon.
// Can be used for an arrow head.
type FilledPoly struct {
	Vertices []*Point // Do not repeat first point as last point.
}

// Justification is a typesafe string for text justifications
type Justification string

// The corresponding values for label justification.
const (
	Left   Justification = "Left"
	Right  Justification = "Right"
	Top    Justification = "Top"
	Bottom Justification = "Bottom"
	Centre Justification = "Centre"
)

// Label is just a (single line) string, the expected font height,
// and its position.
type Label struct {
	TheString  string
	FontHeight float64
	Anchor     *Point
	HJust      Justification
	VJust      Justification
}

// Primitives is a container for a set of Line(s) and a set of Label(s).
type Primitives struct {
	Lines       []*Line
	FilledPolys []*FilledPoly
	Labels      []*Label
}

// NewPrimitives constructs a Primitives ready to use.
func NewPrimitives() *Primitives {
	return &Primitives{[]*Line{}, []*FilledPoly{}, []*Label{}}
}

// AddLine adds the given line to the Primitive's line store.
func (p *Primitives) AddLine(
	x1 float64, y1 float64, x2 float64, y2 float64, dashed bool) {
	line := NewLine(Point(x1, y1), NewLine(x2, y2), dashed}
	p.Lines = append(p.Lines, line)
}

// AddFilledPoly adds the given filled polygon to the Primitive's store.
func (p *Primitives) AddFilledPoly(vertices []*Point) {
	poly := &FilledPoly{vertices}
	p.FilledPolys = append(p.FilledPolys, poly)
}

// AddLabel adds a Label to the Primitive's Lable store.
func (p *Primitives) AddLabel(theString string, fontHeight float64,
	x float64, y float64, hJust Justification, vJust Justification) {
	label := &Label{theString, fontHeight, &Point{x, y}, hJust, vJust}
	p.Labels = append(p.Labels, label)
}

// AddRect adds 4 lines to the Primitive's line store to represent
// the rectangle of the given opposite corners.
func (p *Primitives) AddRect(
	left float64, top float64, right float64, bot float64) {
	p.AddLine(left, top, right, top, false)
	p.AddLine(right, top, right, bot, false)
	p.AddLine(right, bot, left, bot, false)
	p.AddLine(left, bot, left, top, false)
}

// Add adds the Primitives given to those already held in the model.
func (p *Primitives) Add(newPrims *Primitives) {
	p.Lines = append(p.Lines, newPrims.Lines...)
	p.FilledPolys = append(p.FilledPolys, newPrims.FilledPolys...)
	p.Labels = append(p.Labels, newPrims.Labels...)
}
