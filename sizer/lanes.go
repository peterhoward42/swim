package sizers

// This module provides the *Lifeline* and *LifelineInfo* types, which
// encapsulatesizing data for lifelines.

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
)

// Naming conventions:
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

// Lifelines holds sizing information for the lifelines collectively.
// It delegates to LifelineInfo instances for lifeline-specific data.
type Lifelines struct {
	DiagramWidth       float64
	FontHeight         float64
	LifelineStatements []*dslmodel.Statement
	NumLifelines       int
	TitleBoxWidth      float64
	TitleBoxPitch      float64
	TitleBoxHeight     float64
	TitleBoxLabelPadB  float64
	TitleBoxPadR       float64 // Holds title boxes apart
	FirstTitleBoxPadL  float64 // Positions leftmost title box
	TitleBoxPadB       float64 // Below title box as a whole
	SelfLoopWidth      float64
	ActivityBoxWidth   float64
	Individual         InfoPerLifeline
}

// InfoPerLifeline provides information about individual lifelines, keyed on
// the DSL Lifeline statement.
type InfoPerLifeline map[*dslmodel.Statement]*LifelineInfo

// LifelineInfo carries information about one Lifeline.
type LifelineInfo struct {
	TitleBoxLeft     float64
	Centre           float64
	TitleBoxRight    float64
	SelfLoopRight    float64
	SelfLoopCentre   float64
	ActivityBoxLeft  float64
	ActivityBoxRight float64
}

// NewLifelines provides a lifelines structure that has been initialised
// as is ready for use.
func NewLifelines(diagramWidth int, fontHeight float64,
	lifelineStatements []*dslmodel.Statement) *Lifelines {
	lifelines := &Lifelines{}
	lifelines.DiagramWidth = float64(diagramWidth)
	lifelines.FontHeight = fontHeight
	lifelines.LifelineStatements = lifelineStatements
	lifelines.NumLifelines = len(lifelines.LifelineStatements)
	lifelines.populateTitleBoxAttribs()
	lifelines.SelfLoopWidth = 0.5 * lifelines.TitleBoxWidth // see how this works out
	lifelines.ActivityBoxWidth = activityBoxWidthK * fontHeight
	lifelines.populateIndividualLifelineInfo()

	return lifelines
}

// InteractionLineEndPoints works out the x coordinates for an interaction
// line between two given lifelines.
func (lifelines *Lifelines) InteractionLineEndPoints(
	sourceLifeline, destLifeline *dslmodel.Statement) (x1, x2 float64) {
	sourceLifelineSiz := lifelines.Individual[sourceLifeline]
	destLifelineSiz := lifelines.Individual[destLifeline]
	if destLifelineSiz.Centre > sourceLifelineSiz.Centre {
		x1 = sourceLifelineSiz.ActivityBoxRight
		x2 = destLifelineSiz.ActivityBoxLeft
	} else {
		x1 = sourceLifelineSiz.ActivityBoxLeft
		x2 = destLifelineSiz.ActivityBoxRight
	}
	return
}

// InteractionLabelPosition works out the position and justification
// that should be used for an interaction line's label.
func (lifelines *Lifelines) InteractionLabelPosition(
	sourceLifeline, destLifeline *dslmodel.Statement, padding float64) (
	x float64, horizJustification graphics.Justification) {
	sourceLifelineSiz := lifelines.Individual[sourceLifeline]
	destLifelineSiz := lifelines.Individual[destLifeline]
	if destLifelineSiz.Centre > sourceLifelineSiz.Centre {
		x = destLifelineSiz.ActivityBoxLeft - padding
		horizJustification = graphics.Right
	} else {
		x = destLifelineSiz.ActivityBoxRight + padding
		horizJustification = graphics.Left
	}
	return
}

// populateTitleBoxAttribs works out the values for the TitleBoxXXX attributes.
func (lifelines *Lifelines) populateTitleBoxAttribs() {
	// The title boxes are all the same width and height.
	lifelines.TitleBoxHeight = lifelines.titleBoxHeight()
	lifelines.TitleBoxLabelPadB = titleBoxTextPadBK * lifelines.FontHeight
	// The horizontal gaps between them are a fixed proportion of their width.
	// The margins from the edge of the diagram is the same as this gap.
	n := float64(lifelines.NumLifelines)
	nMargins := 2.0
	nGaps := n - 1
	k := titleBoxPadRK
	w := lifelines.DiagramWidth / (k*(nMargins+nGaps) + n)
	lifelines.TitleBoxWidth = w
	lifelines.TitleBoxPadR = k * w
	lifelines.TitleBoxPitch = w * (1 + k)
	lifelines.FirstTitleBoxPadL = k * w
	lifelines.TitleBoxPadB = titleBoxPadBK * lifelines.FontHeight
}

// titleBoxHeight calculates the height based on sufficient room for the
// title box with the most label lines.
func (lifelines *Lifelines) titleBoxHeight() float64 {
	maxNLabelLines := 0
	for _, s := range lifelines.LifelineStatements {
		n := len(s.LabelSegments)
		if n > maxNLabelLines {
			maxNLabelLines = n
		}
	}
	topMargin := titleBoxTextPadTK * lifelines.FontHeight
	botMargin := titleBoxTextPadBK * lifelines.FontHeight
	ht := topMargin + botMargin + float64(maxNLabelLines)*lifelines.FontHeight
	return ht
}

// populateIndividualLifelineInfo sets attributes for things like the
// left, right and centre of the lifeline title box.
func (lifelines *Lifelines) populateIndividualLifelineInfo() {
	lifelines.Individual = InfoPerLifeline{}
	for lifelineNumber, statement := range lifelines.LifelineStatements {
		centre := lifelines.FirstTitleBoxPadL + 0.5*lifelines.TitleBoxWidth +
			float64((lifelineNumber))*lifelines.TitleBoxPitch
		left := centre - 0.5*lifelines.TitleBoxWidth
		right := centre + 0.5*lifelines.TitleBoxWidth
		selfLoopRight := centre + lifelines.SelfLoopWidth
		selfLoopCentre := 0.5 * (centre + selfLoopRight)
		activityBoxLeft := centre - 0.5*lifelines.ActivityBoxWidth
		activityBoxRight := centre + 0.5*lifelines.ActivityBoxWidth
		lifelineInfo := &LifelineInfo{left, centre, right, selfLoopRight,
			selfLoopCentre, activityBoxLeft, activityBoxRight}
		lifelines.Individual[statement] = lifelineInfo
	}
}
