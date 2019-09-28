package diag

import (
	"github.com/peterhoward42/umli/dsl"
)

/*
This module contains code that knows how to create lifelines, including
where to break them in order to avoid activity boxes and
interaction lines that cross them..

Where things are named *seg* this is short for segment.
*/

// Lifelines exposes the API for creating lifelineMaker.
type lifelineMaker struct {
	creator              *Creator
	titleBoxTopAndBottom *segment // Set once known
}

// NewLifelines creates a Lifelines, ready to use.
func newLifelineMaker(creator *Creator) *lifelineMaker {
	lifelines := lifelineMaker{}
	lifelines.creator = creator
	return &lifelines
}

// titleBoxHeight calculates the height based on sufficient room for the
// title box with the most label lines.
func (ll *lifelineMaker) titleBoxHeight() float64 {
	maxNLabelLines := 0
	for _, s := range ll.creator.model.LifelineStatements() {
		n := len(s.LabelSegments)
		if n > maxNLabelLines {
			maxNLabelLines = n
		}
	}
	topMargin := ll.creator.sizer.TitleBoxLabelPadT
	botMargin := ll.creator.sizer.TitleBoxLabelPadB
	ht := topMargin + botMargin + float64(maxNLabelLines)*ll.creator.fontHeight
	return ht
}

// ProduceLifelines works out the dashed line segments that should be created
// to render all the lifelines, including leaving gaps where there are activity
// boxes and interaction lines that cross a lifeline.
func (ll *lifelineMaker) produceLifelines() {
	for _, lifelineStatement := range ll.creator.model.LifelineStatements() {
		lineSegments := ll.segmentsForLifeline(lifelineStatement)
		x := ll.creator.lifelineGeomH.CentreLine(lifelineStatement)
		for i := 0; i < len(lineSegments); i++ {
			seg := lineSegments[i]
			dashed := true
			ll.creator.graphicsModel.Primitives.AddLine(
				x, seg.start, x, seg.end, dashed)
		}
	}
}

/*
segmentsForLifeline works out the set of dashed line segments that are required
to represent one lifeline - accomodating the gaps needed where
the lifeline activity boxes live, or interaction lines that cross this
lifeline.
*/
func (ll *lifelineMaker) segmentsForLifeline(lifeline *dsl.Statement) (
	lineSegments []*segment) {

	// Acquire and combine the (ordered) gap requirements - between which
	// line segments should exist.

	pretendPreGap := &segment{-1, ll.titleBoxTopAndBottom.end}
	activityBoxGaps := ll.creator.activityBoxes[lifeline].boxExtentsAsSegments()
	crossingInteractionLineGaps := ll.creator.ilZones.gapsFor(lifeline)
	lifelinesBottom := ll.creator.tideMark
	pretendPostGap := &segment{lifelinesBottom, lifelinesBottom + 1}

	gaps := []*segment{pretendPreGap}
	gaps = append(gaps, activityBoxGaps...)
	gaps = append(gaps, crossingInteractionLineGaps...)
	gaps = append(gaps, pretendPostGap)

	sortSegments(gaps)
	gaps = mergeSegments(gaps)

	// Make a line segment in between each pair of gaps.
	// (Omitting any that are too small to be sensible to render)

	lineSegments = []*segment{}
	for i := 0; i < len(gaps)-1; i++ {
		top := gaps[i].end
		bot := gaps[i+1].start
		length := bot - top
		if length < ll.creator.sizer.MinLifelineSegLength {
			continue
		}
		lineSegments = append(lineSegments, &segment{top, bot})
	}
	return lineSegments
}
