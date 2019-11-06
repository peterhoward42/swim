package lifeline

import (
	"fmt"

	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

// Finalizer knows how to draw lifelines including making the gaps required
// in them to avoid activity boxes and interaction line no go zones.
type Finalizer struct {
	lifelines     []*dsl.Statement
	spacer        *Spacing
	noGoZones     []nogozone.NoGoZone
	activityBoxes map[*dsl.Statement]*ActivityBoxes
	sizer         sizer.Sizer
}

// NewFinalizer provides a Finalizer ready to use.
func NewFinalizer(
	lifelines []*dsl.Statement,
	spacer *Spacing,
	noGoZones []nogozone.NoGoZone,
	activityBoxes map[*dsl.Statement]*ActivityBoxes,
	sizer sizer.Sizer) *Finalizer {
	return &Finalizer{
		lifelines:     lifelines,
		spacer:        spacer,
		noGoZones:     noGoZones,
		activityBoxes: activityBoxes,
		sizer:         sizer,
	}
}

// Finalize draws all the lifelines.
func (f *Finalizer) Finalize(
	top float64, bottom float64, minSegLen float64,
	primitives *graphics.Primitives) (updatedTidemark float64, err error) {
	for _, ll := range f.lifelines {
		if err := f.finalizeOne(ll, top, bottom, minSegLen, primitives); err != nil {
			return -1, fmt.Errorf("finalizeOne: %v", err)
		}
	}
	return bottom + f.sizer.Get("FrameInternalPadB"), nil
}

// finalizeOne draws one lifeline.
func (f *Finalizer) finalizeOne(
	lifeline *dsl.Statement, top float64, bottom float64,
	minSegLen float64, primitives *graphics.Primitives) error {
	activityBoxes := *f.activityBoxes[lifeline]
	lifelineSegments := LifelineSegments{}
	lifelineSegments.Assemble(lifeline, top, bottom, minSegLen, f.noGoZones, activityBoxes, f.lifelines)
	lifelineXCoords, err := f.spacer.CentreLine(lifeline)
	if err != nil {
		return fmt.Errorf("space.CentreLine: %v", err)
	}
	x := lifelineXCoords.Centre
	dashed := true
	for _, seg := range lifelineSegments.Segs {
		primitives.AddLine(x, seg.Start, x, seg.End, dashed)
	}
	return nil
}
