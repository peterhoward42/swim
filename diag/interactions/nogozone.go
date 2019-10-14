package interactions

import (
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
)

/*
NoGoZone models the space that an interaction line and its label occupies,
in order to advise crossing lifelines where they must be broken, so as to avoid
clashing with the interaction lines.
*/
type NoGoZone struct {
	Height      geom.Segment
	OneEndLifeline      *dsl.Statement
	OtherEndLifeline *dsl.Statement
}
