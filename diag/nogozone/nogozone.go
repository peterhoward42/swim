package nogozone

import (
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
)

/*
NoGoZone models the space that a (horizontal) interaction line and its label
occupies.
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
