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
type LifelineSegments struct {
	Segs []geom.Segment
}

/*
Assemble populates s by working out the segments required by finding out what gaps are
required to avoid NoGoZone(s), and to avoid lifeline's activity boxes, and reconciling
these needs, and making a few adjustments - for example to avoid ending
up with very tiny segments.
*/
func (s *LifelineSegments) Assemble(
	lifeline *dsl.Statement,
	topOfLifeline float64,
	bottomOfLifeline float64,
	minSegLen float64,
	noGoZones []interactions.NoGoZone,
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
