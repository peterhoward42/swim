package lifeline

import (
	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
)

/*
LifelineSegments represents the separate line segments that should be drawn to
represent a lifeline, once the gaps that are needed in it  have been taken into
account.
*/
type LifelineSegments struct {
	Segs []geom.Segment
}

/*
Assemble populates a LifelineSegments by working out the segments
required taking into account
both the gaps that are required to avoid NoGoZone(s), and the gaps required to
not interfere with the lifeline's activity boxes. It also makes a few
adjustments - for example to avoid ending up with very tiny segments.
*/
func (s *LifelineSegments) Assemble(
	lifeline *dsl.Statement,
	topOfLifeline float64,
	bottomOfLifeline float64,
	minSegLen float64,
	noGoZones []nogozone.NoGoZone,
	activityBoxes ActivityBoxes,
	allLifelines []*dsl.Statement) {

	gaps := Gaps{}
	gaps.PopulateFromNoGoZones(noGoZones, lifeline, allLifelines)
	gapsForActivityBoxes := activityBoxes.AsSegments()
	gaps.Items = append(gaps.Items, gapsForActivityBoxes...)
	geom.SortSegments(gaps.Items)
	mergedGaps := geom.MergeSegments(gaps.Items)

	prev := topOfLifeline
	var segs []geom.Segment
	for _, gap := range mergedGaps {
		seg := geom.Segment{prev, gap.Start}
		if seg.Length() >= minSegLen {
			segs = append(segs, seg)
		}
		prev = gap.End
	}
	finalSeg := geom.Segment{prev, bottomOfLifeline}
	if finalSeg.Length() >= minSegLen {
		segs = append(segs, finalSeg)
	}
	s.Segs = segs
}
