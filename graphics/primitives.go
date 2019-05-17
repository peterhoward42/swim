package graphics

// Point represents one point (X,Y)
type Point struct {
	X float64
	Y float64
}

// Line represents a line, optionally dashed.
type Line struct {
	P1     *Point
	P2     *Point
	Dashed bool // vs. Full
}

// FilledPoly represents a filled arrow head.
type FilledPoly struct {
	Vertices []*Point // Do not repeat first point as last point.
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
	Anchor      *Point
	HJust       string
	VJust       string
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
	line := &Line{&Point{x1, y1}, &Point{x2, y2}, dashed}
	p.Lines = append(p.Lines, line)
}

// AddFilledPoly adds the given filled polygon to the Primitive's store.
func (p *Primitives) AddFilledPoly(vertices []*Point) {
	poly := &FilledPoly{vertices}
	p.FilledPolys = append(p.FilledPolys, poly)
}

// AddLabel adds a Label to the Primitive's Lable store.
func (p *Primitives) AddLabel(linesOfText []string, x float64, y float64,
	hJust string, vJust string) {
	label := &Label{linesOfText, &Point{x, y}, hJust, vJust}
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
