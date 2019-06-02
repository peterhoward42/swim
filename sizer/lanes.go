package sizers

// This module provides the *Lanes* and *LaneInfo* types, which encapsulate
// sizing data for lanes.

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
)

// Naming conventions:
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

// Lanes holds sizing information for the lanes collectively.
// It delegates to LaneInfo instances for lane-specific data.
type Lanes struct {
	DiagramWidth            float64
	FontHeight              float64
	LifelineStatements      []*dslmodel.Statement
	NumLanes                int
	TitleBoxWidth           float64
	TitleBoxPitch           float64
	TitleBoxHeight          float64
	TitleBoxBottomRowOfText float64 // Offset below top of title box.
	TitleBoxPadR            float64 // Holds title boxes apart
	FirstTitleBoxPadL       float64 // Positions leftmost title box
	TitleBoxPadB            float64 // Below title box as a whole
	SelfLoopWidth           float64
	ActivityBoxWidth        float64
	Individual              InfoPerLane
}

// InfoPerLane provides information about individual lanes, keyed on
// the DSL Lane statement.
type InfoPerLane map[*dslmodel.Statement]*LaneInfo

// LaneInfo carries information about one Lane.
type LaneInfo struct {
	TitleBoxLeft     float64
	Centre           float64
	TitleBoxRight    float64
	SelfLoopRight    float64
	SelfLoopCentre   float64
	ActivityBoxLeft  float64
	ActivityBoxRight float64
}

// NewLanes provides a Lanes structure that has been initialised
// as is ready for use.
func NewLanes(diagramWidth int, fontHeight float64,
	lifelineStatements []*dslmodel.Statement) *Lanes {
	lanes := &Lanes{}
	lanes.DiagramWidth = float64(diagramWidth)
	lanes.FontHeight = fontHeight
	lanes.LifelineStatements = lifelineStatements
	lanes.NumLanes = len(lanes.LifelineStatements)
	lanes.populateTitleBoxAttribs()
	lanes.SelfLoopWidth = 0.5 * lanes.TitleBoxWidth // see how this works out
	lanes.ActivityBoxWidth = activityBoxWidthK * fontHeight
	lanes.populateIndividualLaneInfo()

	return lanes
}

// InteractionLineEndPoints works out the x coordinates for an interaction
// line between two given lifelines.
func (lanes *Lanes) InteractionLineEndPoints(
	sourceLane, destLane *dslmodel.Statement) (x1, x2 float64) {
	sourceLaneSiz := lanes.Individual[sourceLane]
	destLaneSiz := lanes.Individual[destLane]
	if destLaneSiz.Centre > sourceLaneSiz.Centre {
		x1 = sourceLaneSiz.ActivityBoxRight
		x2 = destLaneSiz.ActivityBoxLeft
	} else {
		x1 = sourceLaneSiz.ActivityBoxLeft
		x2 = destLaneSiz.ActivityBoxRight
	}
	return
}

// InteractionLabelPosition works out the position and justification
// that should be used for an interaction line's label.
func (lanes *Lanes) InteractionLabelPosition(
    sourceLane, destLane *dslmodel.Statement, padding float64) (
    x float64, horizJustification graphics.Justification) {
	sourceLaneSiz := lanes.Individual[sourceLane]
	destLaneSiz := lanes.Individual[destLane]
	if destLaneSiz.Centre > sourceLaneSiz.Centre {
        x = destLaneSiz.ActivityBoxLeft - padding
        horizJustification = graphics.Right
    } else {
        x = destLaneSiz.ActivityBoxRight + padding
        horizJustification = graphics.Left
    }
    return
}

// populateTitleBoxAttribs works out the values for the TitleBoxXXX attributes.
func (lanes *Lanes) populateTitleBoxAttribs() {
	// The title boxes are all the same width and height.
	lanes.TitleBoxHeight = lanes.titleBoxHeight()
	lanes.TitleBoxBottomRowOfText = lanes.TitleBoxHeight -
		titleBoxTextPadBK*lanes.FontHeight
	// The horizontal gaps between them are a fixed proportion of their width.
	// The margins from the edge of the diagram is the same as this gap.
	n := float64(lanes.NumLanes)
	nMargins := 2.0
	nGaps := n - 1
	k := titleBoxPadRK
	w := lanes.DiagramWidth / (k*(nMargins+nGaps) + n)
	lanes.TitleBoxWidth = w
	lanes.TitleBoxPadR = k * w
	lanes.TitleBoxPitch = w * (1 + k)
	lanes.FirstTitleBoxPadL = k * w
	lanes.TitleBoxPadB = titleBoxPadBK * lanes.FontHeight
}

// titleBoxHeight calculates the height based on sufficient room for the
// title box with the most label lines.
func (lanes *Lanes) titleBoxHeight() float64 {
	maxNLabelLines := 0
	for _, s := range lanes.LifelineStatements {
		n := len(s.LabelSegments)
		if n > maxNLabelLines {
			maxNLabelLines = n
		}
	}
	topMargin := titleBoxTextPadTK * lanes.FontHeight
	botMargin := titleBoxTextPadBK * lanes.FontHeight
	ht := topMargin + botMargin + float64(maxNLabelLines)*lanes.FontHeight
	return ht
}

// populateIndividualLaneInfo sets attributes for things like the
// left, right and centre of the lane title box.
func (lanes *Lanes) populateIndividualLaneInfo() {
	lanes.Individual = InfoPerLane{}
	for laneNumber, statement := range lanes.LifelineStatements {
		centre := lanes.FirstTitleBoxPadL + 0.5*lanes.TitleBoxWidth +
			float64((laneNumber))*lanes.TitleBoxPitch
		left := centre - 0.5*lanes.TitleBoxWidth
		right := centre + 0.5*lanes.TitleBoxWidth
		selfLoopRight := centre + lanes.SelfLoopWidth
		selfLoopCentre := 0.5 * (centre + selfLoopRight)
		activityBoxLeft := centre - 0.5*lanes.ActivityBoxWidth
		activityBoxRight := centre + 0.5*lanes.ActivityBoxWidth
		laneInfo := &LaneInfo{left, centre, right, selfLoopRight,
			selfLoopCentre, activityBoxLeft, activityBoxRight}
		lanes.Individual[statement] = laneInfo
	}
}
