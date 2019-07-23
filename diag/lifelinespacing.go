package diag

import (
	"fmt"

	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
LifelineSpacing holds the knowledge about the horizontal geometry of lifelines.
For example, how to space them out across the page, and providing the left
and right edge coordinates for the activity boxes on each etc.

It makes all the title boxes the same width, and distributes these equally
across the width of the diagram. It uses the same gap (gutter) between these
boxes and as margins at the left and right edge of the diagram.
*/
type LifelineSpacing struct {
	lifelineIndices map[*dslmodel.Statement]int
	BoxWidth        float64
	Gutter          float64
}

// NewLifelineSpacing provides a LifelineSpacing ready to use.
func NewLifelineSpacing(diagWidth int, fontHt float64,
	lifelines []*dslmodel.Statement, idealLifelineTitleBoxWidth float64) *LifelineSpacing {

	sp := &LifelineSpacing{}
	sp.lifelineIndices = map[*dslmodel.Statement]int{}
	for i, lifeline := range lifelines {
		sp.lifelineIndices[lifeline] = i
	}

	// Evaluate and store the two key driving dimensions - i.e. the width of
	// title boxes and the gap between them.

	// Start with idealised title box width.
	// I.e big enough to fit circa 3 short title words across.
	n := float64(len(lifelines))
	sp.BoxWidth = idealLifelineTitleBoxWidth
	spaceAvail := float64(diagWidth) - sp.BoxWidth*n
	nGuttersRequired := n + 1
	sp.Gutter = spaceAvail / nGuttersRequired

	fmt.Printf("Ideal: %v, %v\n", sp.BoxWidth, sp.Gutter)

	// But if that has that made the gutter too small, or even negative,
	// make the boxes less wide to preserve a minimum gutter equal to
	// one font height.
	if sp.Gutter < fontHt {
		sp.Gutter = fontHt
		sp.BoxWidth = (float64(diagWidth) - ((n + 1) * sp.Gutter)) / n
	}
	return sp
}

/*
CentreLine provides the X coordinate for the centreline of the Nth lifeline.
Zero-based index.
*/
func (sp *LifelineSpacing) CentreLine(lifeline *dslmodel.Statement) float64 {
	n := float64(sp.lifelineIndices[lifeline])
	return (n+1)*sp.Gutter + (n+0.5)*sp.BoxWidth
}

/*
ActivityBoxXCoords provides the X coordinates for an activity box on a lifeline
*/
func (sp *LifelineSpacing) ActivityBoxXCoords(lifeline *dslmodel.Statement,
	sizer *sizer.Sizer) (left, centre, right float64) {
	centre = sp.CentreLine(lifeline)
	left = centre - 0.5*sp.BoxWidth
	right = centre + 0.5*sp.BoxWidth
	return left, centre, right
}

// InteractionLineEndPoints works out the x coordinates for an interaction
// line between two given lifelines.
func (sp *LifelineSpacing) InteractionLineEndPoints(
	sourceLifeline, destLifeline *dslmodel.Statement,
	sizer *sizer.Sizer) (x1, x2 float64) {
	sourceC := sp.CentreLine(sourceLifeline)
	destC := sp.CentreLine(destLifeline)
	if sourceC > destC {
		x1 = sourceC - 0.5*sp.BoxWidth
		x2 = destC + 0.5*sp.BoxWidth
	} else {
		x1 = sourceC + 0.5*sp.BoxWidth
		x2 = destC - 0.5*sp.BoxWidth
	}
	return
}

// InteractionLabelPosition works out the position and justification
// that should be used for an interaction line's label.
func (sp *LifelineSpacing) InteractionLabelPosition(
	sourceLifeline, destLifeline *dslmodel.Statement) (
	x float64, horizJustification graphics.Justification) {
	x = 0.5 * (sp.CentreLine(sourceLifeline) + sp.CentreLine(sourceLifeline))
	horizJustification = graphics.Centre
	return
}
