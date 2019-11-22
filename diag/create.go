package diag

import (
	"fmt"

	"github.com/peterhoward42/umli/diag/frame"
	"github.com/peterhoward42/umli/diag/interactions"
	"github.com/peterhoward42/umli/diag/lifeline"
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
It orchestrates the creation process which accumulates the graphics
primitives required in its graphicsModel and then returns that model.
*/
func (c *Creator) Create(dslModel dsl.Model) (*graphics.Model, error) {
	// We need to establish two fundamental sizing drivers, and seek the
	// the help of a sizer.Sizer that is initialised with these, before we do
	// much else.
	width, fontHeight := DrivingDimensions{}.WidthAndFontHeight(dslModel)
	sizer := sizer.NewCompleteSizer(fontHeight)

	// Initialise the graphics model that will be populated with lines, text,
	// arrows etc.
	graphicsModel := graphics.NewModel(
		width, fontHeight,
		sizer.Get("DashLineDashLen"),
		sizer.Get("DashLineDashGap"))
	prims := graphicsModel.Primitives

	// Delegate to a specialised object to take responsibility for the graphics
	// of the overall outer frame and title box.
	frameMaker := frame.NewMaker(sizer, fontHeight, width, prims)
	tideMark := frameMaker.InitFrameAndMakeTitleBox(dslModel.Title(),
		sizer.Get("DiagramPadT"))

	// Seek help from another sizing/spacing component - this time, one that is
	// knows how to spread lifelines across the diagram width-wise.
	lifelines := dslModel.LifelineStatements()
	lifelineSpacing := lifeline.NewSpacing(sizer, fontHeight, width, lifelines)

	// Still focussing on graphics that are conceptually anchored to the top
	// of the diagram, we can delegate to a component that knows how to make
	// the title boxes at the top of each lifeline.
	titleBoxes := lifeline.NewTitleBoxes(sizer, lifelineSpacing, lifelines, fontHeight)
	tideMark, bottomOfTitleBoxes, err := titleBoxes.Make(tideMark, prims)
	if err != nil {
		return nil, fmt.Errorf("titleBoxes.Make: %v", err)
	}

	// Now we're going to make the graphics for all the interaction lines,
	// and "work down the page" as we do so.
	// It requires us to prepare some helper components.

	// A set of components that keep track of where activity boxes should be
	// started and stopped on each lifeline.
	boxes := map[*dsl.Statement]*lifeline.BoxTracker{}
	for _, ll := range lifelines {
		boxes[ll] = lifeline.NewBoxTracker()
	}

	// Now construct the component that makes the interaction lines and their
	// labels and arrows.
	d := interactions.NewMakerDependencies(
		fontHeight, lifelineSpacing, sizer, boxes)
	interactionsMaker := interactions.NewMaker(d, graphicsModel)

	// And mandate it to do so.
	tideMark, noGoZones, err := interactionsMaker.ScanInteractionStatements(
		tideMark, dslModel.Statements())
	if err != nil {
		return nil, fmt.Errorf("interactionsMaker.ScanInteractionStatements: %v", err)
	}

	// Now we know how far south the diagram has grown, we can terminate and draw,
	// any activity boxes that have not been closed explicity with a stop command.
	for _, ll := range lifelines {
		boxes := boxes[ll]
		if err := boxes.TerminateAt(tideMark); err != nil {
			return nil, fmt.Errorf("boxes.TerminateAt: %v", err)
		}
		lifeCoords, err := lifelineSpacing.CentreLine(ll)
		if err != nil {
			return nil, fmt.Errorf("lifelineSpacing.CentreLine: %v", err)
		}
		lifeline.NewBoxDrawer(*boxes, lifeCoords.Centre,
			sizer.Get("ActivityBoxWidth")).Draw(prims)
	}

	tideMark += sizer.Get("FinalizedActivityBoxesPadB")

	// Draw the lifelines from top to bottom, leaving gaps where there are
	// activity boxes, or NoGoZone(s) in the way.
	lifelineFinalizer := lifeline.NewFinalizer(
		lifelines, lifelineSpacing, noGoZones, boxes, sizer)
	minSegLen := sizer.Get("MinLifelineSegLength")
	err = lifelineFinalizer.Finalize(
		bottomOfTitleBoxes, tideMark, minSegLen, graphicsModel.Primitives)
	if err != nil {
		return nil, fmt.Errorf("lifelineFinalizer.Finalize: %v", err)
	}

	tideMark += sizer.Get("LifelinePadB")

	// Finish up by drawing the frame's enclosing rectangle.
	tideMark = frameMaker.FinalizeFrame(tideMark)

	// Tell the graphicsModel what its resultant height is.
	tideMark += sizer.Get("DiagramPadB")
	graphicsModel.Height = tideMark

	return graphicsModel, nil
}
