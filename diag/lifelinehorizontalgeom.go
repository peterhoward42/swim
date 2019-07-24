package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
LifelineHorizontalGeometry holds the knowledge about the horizontal geometry of lifelines.
For example, how to space them out across the page, and providing the left
and right edge coordinates for the activity boxes on each etc.

It makes all the title boxes the same width, and distributes these equally
across the width of the diagram. It uses the same gap (gutter) between these
boxes and as margins at the left and right edge of the diagram.
*/
type LifelineHorizontalGeometry struct {
	lifelineIndices  map[*dslmodel.Statement]int
	TitleBoxWidth    float64
	TitleBoxGutter   float64
	activityBoxWidth float64
}

// NewLifelineHorizontalGeometry provides a LifelineSpacing ready to use.
func NewLifelineHorizontalGeometry(diagWidth int, fontHt float64,
	lifelines []*dslmodel.Statement, idealLifelineTitleBoxWidth float64,
	activityBoxWidth float64) *LifelineHorizontalGeometry {

	sp := &LifelineHorizontalGeometry{}
	sp.activityBoxWidth = activityBoxWidth
	sp.lifelineIndices = map[*dslmodel.Statement]int{}
	for i, lifeline := range lifelines {
		sp.lifelineIndices[lifeline] = i
	}

	// Evaluate and store the two key driving dimensions - i.e. the width of
	// title boxes and the gap between them.

	// Start with idealised title box width.
	// I.e big enough to fit circa 3 short title words across.
	n := float64(len(lifelines))
	sp.TitleBoxWidth = idealLifelineTitleBoxWidth
	spaceAvail := float64(diagWidth) - sp.TitleBoxWidth*n
	nGuttersRequired := n + 1
	sp.TitleBoxGutter = spaceAvail / nGuttersRequired

	// But if that has that made the gutter too small, or even negative,
	// make the boxes less wide to preserve a minimum gutter equal to
	// one font height.
	if sp.TitleBoxGutter < fontHt {
		sp.TitleBoxGutter = fontHt
		sp.TitleBoxWidth = (float64(diagWidth) - ((n + 1) * sp.TitleBoxGutter)) / n
	}
	return sp
}

/*
CentreLine provides the X coordinate for the centreline of the Nth lifeline.
Zero-based index.
*/
func (sp *LifelineHorizontalGeometry) CentreLine(lifeline *dslmodel.Statement) float64 {
	n := float64(sp.lifelineIndices[lifeline])
	return (n+1)*sp.TitleBoxGutter + (n+0.5)*sp.TitleBoxWidth
}

/*
ActivityBoxXCoords provides the X coordinates for an activity box on a lifeline
*/
func (sp *LifelineHorizontalGeometry) ActivityBoxXCoords(lifeline *dslmodel.Statement,
	sizer *sizer.Sizer) (left, centre, right float64) {
	centre = sp.CentreLine(lifeline)
	left = centre - 0.5*sp.activityBoxWidth
	right = centre + 0.5*sp.activityBoxWidth
	return left, centre, right
}

/*
SelfInteractionLineXCoords provides the left and right coordinates for a *self*
interaction line.
*/
func (sp *LifelineHorizontalGeometry) SelfInteractionLineXCoords(lifeline *dslmodel.Statement,
	sizer *sizer.Sizer) (left, right float64) {
	_, abc, abr := sp.ActivityBoxXCoords(lifeline, sizer)
	left = abr
	right = abc + 0.5*sp.TitleBoxWidth
	return left, right
}

// InteractionLineEndPoints works out the x coordinates for an interaction
// line between two given lifelines.
func (sp *LifelineHorizontalGeometry) InteractionLineEndPoints(
	sourceLifeline, destLifeline *dslmodel.Statement,
	sizer *sizer.Sizer) (x1, x2 float64) {
	sourceC := sp.CentreLine(sourceLifeline)
	destC := sp.CentreLine(destLifeline)
	if sourceC > destC {
		x1 = sourceC - 0.5*sp.activityBoxWidth
		x2 = destC + 0.5*sp.activityBoxWidth
	} else {
		x1 = sourceC + 0.5*sp.activityBoxWidth
		x2 = destC - 0.5*sp.activityBoxWidth
	}
	return
}

// InteractionLabelPosition works out the position and justification
// that should be used for an interaction line's label.
func (sp *LifelineHorizontalGeometry) InteractionLabelPosition(
	sourceLifeline, destLifeline *dslmodel.Statement) (
	x float64, horizJustification graphics.Justification) {
	x = 0.5 * (sp.CentreLine(sourceLifeline) + sp.CentreLine(destLifeline))
	horizJustification = graphics.Centre
	return
}
