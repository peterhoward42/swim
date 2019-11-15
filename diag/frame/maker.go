package frame

import (
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
Maker knows how to draw the diagram outer frame and title box into a
graphics.Primitives structure.
*/
type Maker struct {
	sizer      sizer.Sizer
	frameTop   float64
	fontHeight float64
	diagWidth  float64
	prims      *graphics.Primitives
}

// NewMaker provides a lifelineBoxes ready to use.
func NewMaker(
	s sizer.Sizer, fontHeight float64, diagWidth float64, prims *graphics.Primitives) *Maker {
	return &Maker{
		sizer:      s,
		diagWidth:  diagWidth,
		prims:      prims,
		fontHeight: fontHeight,
	}
}

/*
InitFrameAndMakeTitleBox is responsible capturing the Y coordinate at which
the diagram's frame rectangle should start, and then drawing the diagram title
in an enclosing rectangle just below it. Then advancing the tidemark
accordingly.
*/
func (fm *Maker) InitFrameAndMakeTitleBox(titleSegments []string,
	frameTop float64) (newTideMark float64) {
	fm.frameTop = frameTop
	tideMark := frameTop + fm.sizer.Get("FrameTitleTextPadT")
	topOfTitleTextY := tideMark
	leftOfBox := fm.sizer.Get("FramePadLR")
	leftOfText := leftOfBox + fm.sizer.Get("FrameTitleTextPadL")
	fm.prims.RowOfStrings(leftOfText, topOfTitleTextY,
		fm.fontHeight, graphics.Left, titleSegments)
	tideMark += float64(len(titleSegments)) * fm.fontHeight
	tideMark += fm.sizer.Get("FrameTitleTextPadB")
	rightOfBox := leftOfBox + fm.diagWidth*0.3 // Nowhere other good home for this constant.
	fm.prims.AddRect(leftOfBox, fm.frameTop, rightOfBox, tideMark)
	tideMark += fm.sizer.Get("FrameTitleRectPadB")
	return tideMark
}

/*
FinalizeFrame claims a little space below the diagram vertical extent so far,
and draws the enclosing frame. It is not responsible for reserving space,
below the frame - that is handled externally.
*/
func (fm *Maker) FinalizeFrame(currentTideMark float64) (newTideMark float64) {
	currentTideMark += fm.sizer.Get("FrameInternalPadB")
	frameBottom := currentTideMark
	left := fm.sizer.Get("FramePadLR")
	right := float64(fm.diagWidth) - fm.sizer.Get("FramePadLR")
	fm.prims.AddRect(left, fm.frameTop, right, frameBottom)
	return currentTideMark
}
