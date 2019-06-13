package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

/*
This module contains code that knows how to create lifelines, including
where to break them in order to avoid activity boxes and
interaction lines that cross lifelines.

Where things are named *seg* this is short for segment.
*/

// Lifelines is the exposes the API for creating lifelines.
type Lifelines struct {
	creator *Creator
}

// NewLifelines creates a Lifelines, ready to use.
func NewLifelines(creator *Creator) *Lifelines {
	lifelines := Lifelines{}
	lifelines.creator = creator
	return &lifelines
}

// ProduceLifelines works out the dashed line segments that should be created
// to render all the lifelines, including leaving gaps where there are activity
// boxes and interaction lines that cross a lifeline.
func (ll *Lifelines) ProduceLifelines() {
	for _, lifelineStatement := range ll.creator.lifelineStatements {
		lineSegments := ll.produceOneLifeline(lifelineStatement)
		x := ll.creator.sizer.Lifelines.Individual[lifelineStatement].Centre
		for i := 0; i < len(lineSegments); i++ {
			seg := lineSegments[i]
			ll.creator.graphicsModel.Primitives.AddLine(
				x, seg.start, x, seg.end, true)
		}
	}
}

/*
produceOneLifeline works out the set of dashed line segments that are required
to represent one lifeline - accomodating the gaps needed where
the lifeline activity boxes live, or interaction lines that cross this
lifeline.
*/
func (ll *Lifelines) produceOneLifeline(lifeline *dslmodel.Statement) (
	lineSegments []*segment) {

	// Acquire and combine the (ordered) gap requirements - between which
	// line segments should exist.

	// activityBoxGaps := ll.creator.boxStates.activityBoxExtents(lifeline)
	activityBoxGaps := []*segment{}

	//crossingLifelineGaps := ll.creator.interactionLineSpaceClaims(lifeline)
	crossingLifelineGaps := []*segment{}

	pretendPreGap := &segment{0, ll.creator.sizer.TitleBoxBottom()}
	lifelinesBottom := ll.creator.tideMark
	pretendPostGap := &segment{lifelinesBottom, lifelinesBottom + 1}

	allGaps := []*segment{pretendPreGap}
	allGaps = append(allGaps, activityBoxGaps...)
	allGaps = append(allGaps, crossingLifelineGaps...)
	allGaps = append(allGaps, pretendPostGap)

	sortSegments(allGaps)
	allGaps = mergeSegments(allGaps)

	// Make a line segment in between each pair of gaps.

	lineSegments = []*segment{}
	for i := 0; i < len(allGaps)-1; i++ {
		top := allGaps[i].end
		bot := allGaps[i+1].start
		lineSegments = append(lineSegments, &segment{top, bot})
	}
	return lineSegments
}
