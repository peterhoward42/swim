package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

/*
This module contains code that owns state and processing logic in relation
to drawing the diagram outer frame and title box.
*/

type frameMaker struct {
	creator  *Creator
	frameTop float64
}

// newlifelineBoxes provides a lifelineBoxes ready to use.
func newFrameMaker(creator *Creator) *frameMaker {
	return &frameMaker{creator}
}

func (fm *frameMaker) produceFrameAndTitleBox(statement *dslmodel.Statement) {
	fm.frameTop = fm.creator.tideMark
	// advance tidemark to give title headroom inside box
	// make text
	// advance tidemark to give space twixt title text and title box
	// note bottom of box
	// capture frameleft and titlebox right from sizer
	// draw rect between frame top and bottom of box
	// advance tidemark to give title box some space below
}
