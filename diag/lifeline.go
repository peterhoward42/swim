package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

/*
This module contains code that knows how to create lifelines, including
where to break them in order not to overwrite activity boxes and
interaction lines that they cross.
*/

// Lifelines todo
type Lifelines struct {
	creator *Creator
}

func NewLifelines(creator *Creator) *Lifelines {
	lifelines := Lifelines{}
	lifelines.creator = creator
	return &lifelines
}

func (ll *Lifelines) ProduceLifelines() {
	for _, lifelineStatement := range ll.creator.lifelineStatements {
		segs := ll.produceOne(lifelineStatement)
        for i := 0; i < len(segs); i++ {
            ll.graphicsModel.Primitives.AddLine(x, seg.topY, x, seg.botY, true)
        }
	}
}

// topBot models a pair of Y coordinates (top and bottom) for either
// a line segment, or a gap.
type topBot struct {
	topY        float64
	botY        float64
}

/*
produceOne works out the set of dashed line segments that are required
to represent one lifeline - accomodating the gaps needed where
the lifeline activity boxes live, or interaction lines that cross this
lifeline.
*/
func (ll *Lifelines) produceOne(lifeline *dslmodel.Statement) (
    segments []*topBot){

    // Acquire and combine the (ordered) gap requirements - between which
    // line segments should exist.
    activityBoxGaps := ll.creator.boxStates.activityBoxExtents(lifeline)
    crossingLifelineGaps := ll.creator.interactionLineSpaceClaims(lifeline)

    pretendPreGap := &topBot{stuff}
    pretendPostGap := &topBot{stuff}

    allGaps := []*topBot{pretendPreGap}
    allGaps = append(allGaps, activityBoxGaps...)
    allGaps = append(allGaps, crossingLifelineGaps...)
    allGaps = append(allGaps, pretendPostGap)

    sortedGaps := sort(allGaps, pred)

    // Make a segment in between each pair of gaps.
    
    segments = []*topBot{}
    for i := 0; i < len(sortedGaps); i++ {
        top := sortedGaps[i].botY
        bot := sortedGaps[i+1].topY
        segments = append(segments, &topBot{top, bot}
    }
    return segments
}

