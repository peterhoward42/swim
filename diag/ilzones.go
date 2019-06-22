package diag

import (
	"math"

	"github.com/peterhoward42/umli/dslmodel"
)

/*
Claim models a vertical space claimed by an interaction line or its
label.
*/
type Claim struct {
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
	claims  []*Claim
}

// NewInteractionLineZones provides a new InteractionLineZones
// ready to use.
func NewInteractionLineZones(creator *Creator) *InteractionLineZones {
	ilZones := &InteractionLineZones{
		creator: creator,
		claims:  []*Claim{},
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
	claim := &Claim{sourceLifeline, destLifeline, seg}
	ilz.claims = append(ilz.claims, claim)
}

/*
GapsFor provides a list of Segment(s) that represent the gaps
that should be left in in a lifeline so as not to interfere with the
interaction lines that cross it.
*/
func (ilz *InteractionLineZones) GapsFor(lifeline *dslmodel.Statement) []*segment {
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
	x1 := ilz.creator.sizer.Lifelines.Individual[source].Centre
	x2 := ilz.creator.sizer.Lifelines.Individual[dest].Centre
	target := ilz.creator.sizer.Lifelines.Individual[lifeline].Centre
	left := math.Min(x1, x2)
	right := math.Max(x1, x2)
	return (left < target) && (right > target)
}
