package diag

import (
	"sort"

	"github.com/peterhoward42/umli/dslmodel"
)

/*
This module contains code that knows how to create lifelines, including
where to break them in order to avoid activity boxes and
interaction lines that cross lifelines.

Where things are named *seg* this is short for segment.
*/

// topBot models a pair of Y coordinates (top and bottom) for either
// a line segment, or a gap.
type topBot struct {
	topY float64
	botY float64
}

// gap and lineSeg are aliases for topBot which makes the particular usage
// clearer.
type gap topBot
type lineSeg topBot

// Lifelines todo
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
		lineSegments := ll.produceOne(lifelineStatement)
		x := ll.creator.sizer.Lifelines.Individual[lifelineStatement].Centre
		for i := 0; i < len(lineSegments); i++ {
			seg := lineSegments[i]
			ll.creator.graphicsModel.Primitives.AddLine(
				x, seg.topY, x, seg.botY, true)
		}
	}
}

/*
produceOne works out the set of dashed line segments that are required
to represent one lifeline - accomodating the gaps needed where
the lifeline activity boxes live, or interaction lines that cross this
lifeline.
*/
func (ll *Lifelines) produceOne(lifeline *dslmodel.Statement) (
	segments []*lineSeg) {

	// Acquire and combine the (ordered) gap requirements - between which
	// line segments should exist.

	// activityBoxGaps := ll.creator.boxStates.activityBoxExtents(lifeline)
	activityBoxGaps := []*gap{}

	//crossingLifelineGaps := ll.creator.interactionLineSpaceClaims(lifeline)
	crossingLifelineGaps := []*gap{}

	pretendPreGap := &gap{0, ll.creator.sizer.TitleBoxBottom()}
	lifelinesBottom := ll.creator.tideMark
	pretendPostGap := &gap{lifelinesBottom, lifelinesBottom + 1}

	allGaps := []*gap{pretendPreGap}
	allGaps = append(allGaps, activityBoxGaps...)
	allGaps = append(allGaps, crossingLifelineGaps...)
	allGaps = append(allGaps, pretendPostGap)

	sortGaps(allGaps)

	// Make a segment in between each pair of gaps.

	segments = []*lineSeg{}
	for i := 0; i < len(allGaps)-1; i++ {
		top := allGaps[i].botY
		bot := allGaps[i+1].topY
		segments = append(segments, &lineSeg{top, bot})
	}
	return segments
}

func sortGaps(segs []*gap) {
	sort.Slice(segs, func(i, j int) bool {
		return segs[i].topY < segs[j].topY
	})
}
