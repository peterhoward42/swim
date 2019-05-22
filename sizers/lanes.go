package sizers

import (
	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
)

// Lanes holds sizing information for the lanes.
type Lanes struct {
	DiagramWidth            float64
	FontHeight              float64
	LaneStatements          []*dslmodel.Statement
	NumLanes                int
	TitleBoxWidth           float64
	TitleBoxPitch           float64
	TitleBoxHeight          float64
	TitleBoxBottomRowOfText float64 // Below top of title box.
	TitleBoxHorizGap        float64
	TitleBoxLeftMargin      float64
	Individual              InfoPerLane
}

// InfoPerLane provides information about individual lanes, keyed on
// the DSL Lane statement.
type InfoPerLane map[*dslmodel.Statement]*LaneInfo

// LaneInfo carries information about one Lane.
type LaneInfo struct {
	TitleBoxLeft  float64
	Centre        float64
	TitleBoxRight float64
}

// NewLanes provides a Lanes structure that has been initialised
// as is ready for use.
func NewLanes(diagramWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Lanes {
	lanes := &Lanes{}
	lanes.DiagramWidth = float64(diagramWidth)
	lanes.FontHeight = fontHeight
	lanes.isolateLaneStatements(statements)
	lanes.NumLanes = len(lanes.LaneStatements)
	lanes.populateTitleBoxAttribs()
	lanes.populateIndividualLaneInfo()

	return lanes
}

// populateTitleBoxAttribs works out the values for the TitleBoxXXX attributes.
func (lanes *Lanes) populateTitleBoxAttribs() {
	// The title boxes are all the same width and height.
	lanes.TitleBoxHeight = lanes.titleBoxHeight()
	lanes.TitleBoxBottomRowOfText = lanes.TitleBoxHeight -
		titleBoxTextBotMarginK*lanes.FontHeight
	// The horizontal gaps between them are a fixed proportion of their width.
	// The margins from the edge of the diagram is the same as this gap.
	n := float64(lanes.NumLanes)
	nMargins := 2.0
	nGaps := n - 1
	k := titleBoxSeparationK
	w := lanes.DiagramWidth / (k*(nMargins+nGaps) + n)
	lanes.TitleBoxWidth = w
	lanes.TitleBoxHorizGap = k * w
	lanes.TitleBoxPitch = w * (1 + k)
	lanes.TitleBoxLeftMargin = k * w
}

// titleBoxHeight calculates the height based on sufficient room for the
// title box with the most label lines.
func (lanes *Lanes) titleBoxHeight() float64 {
	maxNLabelLines := 0
	for _, s := range lanes.LaneStatements {
		n := len(s.LabelSegments)
		if n > maxNLabelLines {
			maxNLabelLines = n
		}
	}
	topMargin := titleBoxTextTopMarginK * lanes.FontHeight
	botMargin := titleBoxTextBotMarginK * lanes.FontHeight
	ht := topMargin + botMargin + float64(maxNLabelLines)*lanes.FontHeight
	return ht
}

// populateIndividualLaneInfo sets attributes for things like the
// left, right and centre of the lane title box.
func (lanes *Lanes) populateIndividualLaneInfo() {
	lanes.Individual = InfoPerLane{}
	for laneNumber, statement := range lanes.LaneStatements {
		centre := lanes.TitleBoxLeftMargin + 0.5*lanes.TitleBoxWidth +
			float64((laneNumber))*lanes.TitleBoxPitch
		left := centre - 0.5*lanes.TitleBoxWidth
		right := centre + 0.5*lanes.TitleBoxWidth
		laneInfo := &LaneInfo{left, centre, right}
		lanes.Individual[statement] = laneInfo
	}
}

// isolateLaneStatements isolates the lane statements in a DSL list.
func (lanes *Lanes) isolateLaneStatements(statements []*dslmodel.Statement) {
	for _, s := range statements {
		if s.Keyword == umli.Lane {
			lanes.LaneStatements = append(lanes.LaneStatements, s)
		}
	}
}
