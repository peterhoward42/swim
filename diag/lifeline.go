package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

/*
This module contains code that knows how to create lifelines, including
where to break them in order not to overwrite activity boxes and
int/eraction lines that they cross.
*/

// Lifelines todo
type Lifelines struct {
	creator *Creator
}

func NewLifelines(c *Creator) *Lifelines {
	lifelines := Lifelines{}
	lifelines.creator = c

	return &lifelines
}

func (ll *Lifelines) Create() {
	for _, lifelineStatement := range ll.creator.lifelineStatements {
		ll.createForOneLifeline(lifelineStatement)
	}
}

type lineSegment struct {
	topY        float64
	botY        float64
	stillInPlay bool
}

func (ll *Lifelines) createForOneLifeline(lifeline *dslmodel.Statement) {
	// Start with a continuous unbroken line from the title box down
	// to the bottom.
	top := ll.creator.sizer.DiagramPadT + ll.creator.sizer.Lifelines.TitleBoxHeight
	bot := ll.creator.tideMark
	segment := &lineSegment{top, bot, true}
	segments := []*lineSegment{segment}

	// Cut up the continuous line segment to make room for the lifeline's
	// activity boxes.
	segments = ll.makeRoomForActivityBoxes(lifeline, segments)

	// Cut up the segments further to make room for interactions lines,
	// and their labels that cross this lifeline.
	segments = ll.makeRoomForCrossingInteractionLines(lifeline, segments)

	// Post the segments (that remain in play) into the graphics models,
	// as dashed lines.

	x := ll.creator.sizer.Lifelines.Individual[lifeline].Centre
	for _, seg := range segments {
		if seg.stillInPlay {
			ll.creator.graphicsModel.Primitives.AddLine(
				x, seg.topY, x, seg.botY, true)
		}
	}
}

func (ll *Lifelines) makeRoomForActivityBoxes(
	lifeline *dslmodel.Statement, segments []*lineSegment) (
	updatedSegments []*lineSegment) {
	return nil
}

func (ll *Lifelines) makeRoomForCrossingInteractionLines(
	lifeline *dslmodel.Statement, segments []*lineSegment) (
	updatedSegments []*lineSegment) {
	return nil
}
