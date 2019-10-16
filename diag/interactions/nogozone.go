package interactions

import (
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
)

/*
NoGoZone models the space that a (horizontal) interaction line and its label
occupies, in order to advise crossing (vertical) lifelines where they must be
broken, so as to avoid clashing.
*/
type NoGoZone struct {
	Height           geom.Segment
	OneEndLifeline   *dsl.Statement
	OtherEndLifeline *dsl.Statement
}

// NewNoGoZone creates and initialises a NoGoZone
func NewNoGoZone(height geom.Segment, oneEndLifeline,
	otherEndLifelone *dsl.Statement) NoGoZone {
	return NoGoZone{height, oneEndLifeline, otherEndLifelone}
}
