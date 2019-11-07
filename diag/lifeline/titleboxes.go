package lifeline

import (
	"fmt"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
TitleBoxes knows how to draw the lifeline title boxes
*/
type TitleBoxes struct {
	sizer      sizer.Sizer
	spacer     *Spacing
	lifelines  []*dsl.Statement
	fontHeight float64
}

// NewTitleBoxes creates a TitleBoxes ready to use.
func NewTitleBoxes(sizer sizer.Sizer, lifelineSpacing *Spacing, lifelines []*dsl.Statement,
	fontHeight float64) *TitleBoxes {
	return &TitleBoxes{
		sizer:      sizer,
		spacer:     lifelineSpacing,
		lifelines:  lifelines,
		fontHeight: fontHeight,
	}
}

/*
Make works out the graphics primitives needed to represent all the lifeline
title boxes, and adds them to prims.
*/
func (tbx TitleBoxes) Make(
	currentTideMark float64,
	prims *graphics.Primitives) (newTideMark float64, bottomOfBoxes float64, err error) {

	totalHeight, forLabelsHeight := tbx.Height()
	for _, life := range tbx.lifelines {
		err := tbx.MakeOne(life, currentTideMark, totalHeight, forLabelsHeight,
			prims)
		if err != nil {
			return -1.0, -1.0, fmt.Errorf("MakeOne: %v", err)
		}
	}
	bottomOfBoxes = currentTideMark + totalHeight
	newTideMark = bottomOfBoxes + tbx.sizer.Get("TitleBoxPadB")
	return newTideMark, bottomOfBoxes, nil
}

/*
MakeOne works out the graphics primitives needed to represent the title box
for lifeline and adds them to prims.
*/
func (tbx TitleBoxes) MakeOne(
	lifeline *dsl.Statement,
	topOfBox float64,
	totalHeight float64,
	labelHeight float64,
	prims *graphics.Primitives) error {

	titleBoxXCoords, err := tbx.spacer.CentreLine(lifeline)
	if err != nil {
		return fmt.Errorf("spacing.CentreLine: %v", err)
	}

	// Make the rectangle.
	bottom := topOfBox + totalHeight
	prims.AddRect(titleBoxXCoords.Left, topOfBox, titleBoxXCoords.Right, bottom)

	// Make the strings.
	topRowOfTextY := bottom - tbx.sizer.Get("TitleBoxLabelPadB") - labelHeight
	prims.RowOfStrings(titleBoxXCoords.Centre, topRowOfTextY,
		tbx.fontHeight, graphics.Centre, lifeline.LabelSegments)

	return nil
}

/*
Height provides the height required for the titlebox, based on the lifeline
with the most label segments.
*/
func (tbx TitleBoxes) Height() (overallHeight, forLabels float64) {
	var maxN int
	for _, s := range tbx.lifelines {
		if len(s.LabelSegments) > maxN {
			maxN = len(s.LabelSegments)
		}
	}
	forLabels = float64(maxN) * tbx.fontHeight
	overallHeight = forLabels + tbx.sizer.Get("TitleBoxLabelPadT") + tbx.sizer.Get("TitleBoxLabelPadB")
	return overallHeight, forLabels
}
