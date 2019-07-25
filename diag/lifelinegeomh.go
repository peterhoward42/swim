package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
lifelineGeomH holds the knowledge about the horizontal geometry
of lifelines. For example, how to space them out across the page, and
providing the left and right edge coordinates for the activity boxes on each
etc.

It makes all the title boxes the same width, and distributes these equally
across the width of the diagram. It uses the same gap (gutter) between these
boxes and as margins at the left and right edge of the diagram.
*/
type lifelineGeomH struct {
	sizer           *sizer.Sizer
	lifelineIndices map[*dslmodel.Statement]int
	diagWidth       int
	TitleBoxWidth   float64
	TitleBoxGutter  float64
}

// NewLifelineGeomH provides a lifelineGeomH ready to use.
func newLifelineGeomH(diagWidth int, fontHt float64, sizer *sizer.Sizer,
	lifelines []*dslmodel.Statement) *lifelineGeomH {

	g := &lifelineGeomH{
		lifelineIndices: map[*dslmodel.Statement]int{},
		diagWidth:       diagWidth,
		sizer:           sizer,
	}
	for i, lifeline := range lifelines {
		g.lifelineIndices[lifeline] = i
	}

	// Evaluate and store the two key driving dimensions - i.e. the width of
	// title boxes and the gap between them.

	// Start with idealised title box width.
	// I.e big enough to fit circa 3 short title words across.
	g.TitleBoxWidth = sizer.IdealLifelineTitleBoxWidth
	n := float64(len(lifelines))
	spaceAvail := float64(diagWidth) - g.TitleBoxWidth*n
	nGuttersRequired := n + 1
	g.TitleBoxGutter = spaceAvail / nGuttersRequired

	// But if that has that made the gutter too small, or even negative,
	// make the boxes less wide to preserve a minimum gutter equal to
	// one font height.
	if g.TitleBoxGutter < fontHt {
		g.TitleBoxGutter = fontHt
		g.TitleBoxWidth = (float64(diagWidth) - ((n + 1) *
			g.TitleBoxGutter)) / n
	}
	return g
}

/*
CentreLine provides the X coordinate for the centreline of the Nth lifeline.
Zero-based index.
*/
func (g *lifelineGeomH) CentreLine(lifeline *dslmodel.Statement) float64 {
	n := float64(g.lifelineIndices[lifeline])
	return (n+1)*g.TitleBoxGutter + (n+0.5)*g.TitleBoxWidth
}

/*
ActivityBoxXCoords provides the X coordinates for an activity box on a lifeline
*/
func (g *lifelineGeomH) ActivityBoxXCoords(
	lifeline *dslmodel.Statement) (left, centre, right float64) {
	centre = g.CentreLine(lifeline)
	left = centre - 0.5*g.sizer.ActivityBoxWidth
	right = centre + 0.5*g.sizer.ActivityBoxWidth
	return left, centre, right
}

/*
SelfInteractionLineXCoords provides the left and right coordinates for a *self*
interaction line.
*/
func (g *lifelineGeomH) SelfInteractionLineXCoords(
	lifeline *dslmodel.Statement) (left, right float64) {
	_, abc, abr := g.ActivityBoxXCoords(lifeline)
	left = abr
	right = abc + 0.5*g.TitleBoxWidth
	return left, right
}

// InteractionLineEndPoints works out the x coordinates for an interaction
// line between two given lifelines.
func (g *lifelineGeomH) InteractionLineEndPoints(
	sourceLifeline, destLifeline *dslmodel.Statement) (x1, x2 float64) {
	sourceC := g.CentreLine(sourceLifeline)
	destC := g.CentreLine(destLifeline)
	if sourceC > destC {
		x1 = sourceC - 0.5*g.sizer.ActivityBoxWidth
		x2 = destC + 0.5*g.sizer.ActivityBoxWidth
	} else {
		x1 = sourceC + 0.5*g.sizer.ActivityBoxWidth
		x2 = destC - 0.5*g.sizer.ActivityBoxWidth
	}
	return
}

// InteractionLabelPosition works out the position and justification
// that should be used for an interaction line's label.
func (g *lifelineGeomH) InteractionLabelPosition(
	sourceLifeline, destLifeline *dslmodel.Statement) (
	x float64, horizJustification graphics.Justification) {
	x = 0.5 * (g.CentreLine(sourceLifeline) + g.CentreLine(destLifeline))
	horizJustification = graphics.Centre
	return
}
