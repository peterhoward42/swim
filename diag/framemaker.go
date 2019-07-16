package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
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
	return &frameMaker{creator: creator}
}

/*
initFrameAndMakeTitleBox is responsible capturing the Y coordinate at which
the diagram's frame rectangle should start, and then drawing the diagram title
in an enclosing rectangle just below it. Then advancing the tidemark
accordingly.
*/
func (fm *frameMaker) initFrameAndMakeTitleBox(statement *dslmodel.Statement) {
	c := fm.creator
	fm.frameTop = c.tideMark
	c.tideMark += c.sizer.FrameTitleTextPadT
	topOfTitleTextY := c.tideMark
	leftOfText := c.sizer.DiagPadL + c.sizer.FrameTitleTextPadL
	fm.creator.rowOfLabels(leftOfText, topOfTitleTextY, graphics.Left, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments)) * c.fontHeight
	c.tideMark += c.sizer.FrameTitleTextPadB
	rightOfFrameTitleBox := c.sizer.FrameTitleBoxWidth
	c.graphicsModel.Primitives.AddRect(c.sizer.DiagPadL, fm.frameTop, rightOfFrameTitleBox, c.tideMark)
	c.tideMark += c.sizer.FrameTitleRectPadB
}

/*
finalizeFrame claims a little space below the diagram vertical extent so far,
and draws the enclosing frame.
*/
func (fm *frameMaker) finalizeFrame() {
	c := fm.creator
	c.tideMark += c.sizer.FrameInternalPadB
	frameBottom := c.tideMark
	left := c.sizer.DiagPadL
	right := float64(c.width) - c.sizer.DiagPadL
	c.graphicsModel.Primitives.AddRect(left, fm.frameTop, right, frameBottom)
}
