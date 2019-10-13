package lifeline

import (
	"github.com/peterhoward42/umli/diag/interactions"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
)

/*
Segments represents the line segments that should be drawn to represent a
lifeline while avoiding drawing over the top of other things.
*/
type Segments struct {
	Items []geom.Segment
}

/*
Assemble works out the segments required by finding out what gaps are
required to avoid NoGoZone(s), and to avoid activity boxes, and reconciling
these needs, and making a few adjustments - for example to avoid ending
up with very tiny segments.
*/
func (s *Segments) Assemble(
	lifeline *dsl.Statement,
	topOfLifeline float64,
	bottomOfLifeline float64,
	minSegLen float64,
	noGoZones []interactions.NoGoZone,
	activityBoxes ActivityBoxes,
	allLifelines []*dsl.Statement) {

	gaps := Gaps{}
	gaps.PopulateFromNoGoZones(noGoZones, lifeline, allLifelines)
	gaps.Items := append(gaps.Items, activityBoxes.AsSegments()...)
	geom.SortSegments(gaps.Items)
	mergedGaps := geom.MergeSegments(gaps.Items)

	prev := topOfLifeline
	for _, gap := range mergedGaps {
		seg := geom.Segment{prev, gap.Start}
		if seg.Length() >= minSegLen {
			segs = append(segs. seg)
		}
		prev := gap.End
	}
	finalSeg := geom.Segment{prev, bottomOfLifeline}
	segs = append(segs, finalSeg)
	s.Items = segs
}
