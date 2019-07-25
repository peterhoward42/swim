package diag

/*
This module contains code that keeps track of how much vertical space
each interaction line (and its label) has taken up  as it is drawn. It is
consulted later by the code that draws the lifelines, so that it can make
gaps in them to avoid overwriting (crossing) the interaction lines.
*/

import (
	"math"

	"github.com/peterhoward42/umli/dslmodel"
)

/*
Claim models a vertical space claimed by an interaction line or its
label.
*/
type claim struct {
	sourceLifeline *dslmodel.Statement
	destLifeline   *dslmodel.Statement
	extent         *segment
}

/*
InteractionLineZones hols information about the spaces taken up by
interaction lines and their labels, from the point of view knowing
where to make breaks in lifelines, so as not to clash with them.
*/
type InteractionLineZones struct {
	creator *Creator
	claims  []*claim
}

// NewInteractionLineZones provides a new InteractionLineZones
// ready to use.
func NewInteractionLineZones(creator *Creator) *InteractionLineZones {
	ilZones := &InteractionLineZones{
		creator: creator,
		claims:  []*claim{},
	}
	return ilZones
}

/*
RegisterSpaceClaim records the vertical space claimed by an interaction line,
or its label.
*/
func (ilz *InteractionLineZones) RegisterSpaceClaim(
	sourceLifeline *dslmodel.Statement,
	destLifeline *dslmodel.Statement,
	startY float64,
	endY float64) {
	seg := &segment{startY, endY}
	claim := &claim{sourceLifeline, destLifeline, seg}
	ilz.claims = append(ilz.claims, claim)
}

/*
gapsFor provides a list of Segment(s) that represent the gaps
that should be left in in a lifeline so as not to interfere with the
interaction lines that cross it.
*/
func (ilz *InteractionLineZones) gapsFor(lifeline *dslmodel.Statement) []*segment {
	gaps := []*segment{}
	for _, claim := range ilz.claims {
		if ilz.crosses(lifeline, claim.sourceLifeline, claim.destLifeline) {
			gaps = append(gaps, claim.extent)
		}
	}
	return gaps
}

/*
crosses returns true if a horizontal traversal between source and dest
crosses lifeline.
*/
func (ilz *InteractionLineZones) crosses(lifeline, source,
	dest *dslmodel.Statement) bool {
	x1 := ilz.creator.lifelineGeomH.CentreLine(source)
	x2 := ilz.creator.lifelineGeomH.CentreLine(dest)
	target := ilz.creator.lifelineGeomH.CentreLine(lifeline)
	left := math.Min(x1, x2)
	right := math.Max(x1, x2)
	return (left < target) && (right > target)
}
