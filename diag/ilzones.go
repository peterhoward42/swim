package diag

import (
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
	claims []*Claim
}

// NewInteractionLineZones provides a new InteractionLineZones
// ready to use.
func NewInteractionLineZones() *InteractionLineZones {
	ilZones := &InteractionLineZones{
		claims: []*Claim{},
	}
	return ilZones
}

/*
RegisterSpaceClaim records the vertical space claimed by an interaction line,
or its label.
*/
func (ilz *InteractionLineZones) RegisterSpaceClaim(
	sourceLifeline *dslmodel.Statement, destLifeline *dslmodel.Statement,
	startY float64, endY float64) {
	seg := &segment{startY, endY}
	claim := &Claim{sourceLifeline, destLifeline, seg}
	ilz.claims = append(ilz.claims, claim)
}
