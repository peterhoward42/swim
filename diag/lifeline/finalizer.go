package lifeline

import (
	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
)

// Finalizer knows how to draw lifelines including making the gaps required
// in them to avoid activity boxes and the no-go-zones where they would conflict
// with interaction lines that cross them.
type Finalizer struct {
}

// NewFinalizer provides a Finalizer ready to use.
func NewFinalizer(
	lifelines []*dsl.Statement,
	noGoZones []nogozone.NoGoZone,
	activityBoxes map[*dsl.Statement]*ActivityBoxes) *Finalizer {
	return nil
}

// Finalize draws the lifelines.
func (f *Finalizer) Finalize(
	top float64, bottom float64, primitives *graphics.Primitives) error {
	return nil
}
