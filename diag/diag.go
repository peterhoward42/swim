package diag

import (
	"github.com/peterhoward42/umli/diag/frame"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
Creator is the type that provides the API and entry point for the diag package.
It provides the main Create method that produces a diagram.
*/
type Creator struct {
}

/*
NewCreator instantiates a Creator ready to use.
*/
func NewCreator() (*Creator, error) {
	return nil, nil
}

/*
Create is the main API method which work out what the diagram should look like.
It orchestrates a multi-pass creation process which accumulates the graphics
primitives required in its graphicsModel and then returns that model.
*/
func (c *Creator) Create(dslModel dsl.Model) (*graphics.Model, error) {
	width, fontHeight := DrivingDimensions{}.WidthAndFontHeight(dslModel)
	sizer := sizer.NewCompleteSizer(width, fontHeight)

	graphicsModel := graphics.NewModel(
		width, fontHeight,
		sizer.Get("DashLineDashLen"),
		sizer.Get("DashLineDashGap"))
	prims := graphicsModel.Primitives

	frameMaker := frame.NewMaker(sizer, prims)
	tideMark := frameMaker.InitFrameAndMakeTitleBox(dslModel.Title(),
		sizer.Get("DiagramPadT"))
	_ = tideMark

	return nil, nil
}
