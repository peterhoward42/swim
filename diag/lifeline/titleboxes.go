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
}

/*
Make works out the graphics primitives needed to represent all the lifeline
title boxes, and adds them to prims.
*/
func (tbx TitleBoxes) Make(
	currentTideMark float64,
	lifelines []*dsl.Statement,
	sizer sizer.Sizer,
	spacer *Spacing,
	prims *graphics.Primitives) (newTideMark float64, err error) {

	labelHeight := tbx.HeightRequiredForLabels(sizer, lifelines)
	for _, life := range lifelines {
		err := tbx.MakeOne(life, currentTideMark, labelHeight,
			sizer, spacer, prims)
		if err != nil {
			return -1.0, fmt.Errorf("MakeOne: %v", err)
		}
	}
	return -1.0, nil
}

/*
MakeOne works out the graphics primitives needed to represent the title box
for lifeline and adds them to prims.
*/
func (tbx TitleBoxes) MakeOne(
	lifeline *dsl.Statement,
	topOfBox float64,
	labelHeight float64,
	sizer sizer.Sizer,
	spacing *Spacing,
	prims *graphics.Primitives) error {

	titleBoxXCoords, err := spacing.CentreLine(lifeline)
	if err != nil {
		return fmt.Errorf("spacing.CentreLine: %v", err)
	}
	bottom := topOfBox + labelHeight + sizer.Get("TitleBoxLabelPadT") + sizer.Get("TitleBoxLabelPadB")
	prims.AddRect(titleBoxXCoords.Left, topOfBox, titleBoxXCoords.Right, bottom)

	_ = titleBoxXCoords
	return nil
}

/*
HeightRequiredForLabels provides the vertical space required for the lifeline
with the most label segments.
*/
func (tbx TitleBoxes) HeightRequiredForLabels(
	sizer sizer.Sizer, lifelines []*dsl.Statement) float64 {
	var maxN int
	for _, s := range lifelines {
		if len(s.LabelSegments) > maxN {
			maxN = len(s.LabelSegments)
		}
	}
	return float64(maxN) * sizer.Get("FontHt")
}

/*
// Height provides the height of the title box.
func (tb *TitleBox) Height(sizer sizer.Sizer) float64 {
	n := len(tb.lifeline.LabelSegments)
	topMargin := sizer.Get("TitleBoxLabelPadT")
	botMargin := sizer.Get("TitleBoxLabelPadB")
	ht := topMargin + botMargin + float64(n)*sizer.Get("FontHeight")
	return ht
}
*/
