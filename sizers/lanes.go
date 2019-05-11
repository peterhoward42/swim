package sizers

import (
	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
)

// Lanes holds sizing information for the lanes.
type Lanes struct {
	DiagramWidth       float64
	FontHeight         float64
	LaneStatements     []*dslmodel.Statement
	NumLanes           int
	TitleBoxWidth      float64
	TitleBoxPitch      float64
	TitleBoxHeight     float64
	TitleBoxHorizGap   float64
	TitleBoxLeftMargin float64
	Individual         InfoPerLane
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
func (l *Lanes) populateTitleBoxAttribs() {
	// The title boxes are all the same width and height.
	l.TitleBoxHeight = l.titleBoxHeight()
	// The gaps between them are a fixed proportion of their width.
	// The margins from the edge of the diagram is the same as this gap.
	n := float64(l.NumLanes)
	nMargins := 2.0
	nGaps := n - 1
	k := titleBoxSeparationK
	w := l.DiagramWidth / (k*(nMargins+nGaps) + n)
	l.TitleBoxWidth = w
	l.TitleBoxHorizGap = k * w
	l.TitleBoxPitch = w * (1 + k)
	l.TitleBoxLeftMargin = k * w
}

// titleBoxHeight calculates the height based on sufficient room for the
// title box with the most label lines.
func (l *Lanes) titleBoxHeight() float64 {
	labelLines := 0
	for _, s := range l.LaneStatements {
		n := len(s.LabelSegments)
		if n > labelLines {
			labelLines = n
		}
	}
	n := float64(labelLines)
	topBotMargins := 2.0 * titleBoxTextTopBotMarginK*l.FontHeight
	leading := (n-1) * titleBoxTextRowLeadingK * l.FontHeight
	lines := n * l.FontHeight
	ht := topBotMargins + leading + lines 
	return ht
}

func (l *Lanes) populateIndividualLaneInfo() {
	l.Individual = InfoPerLane{}
	for i, statement := range l.LaneStatements {
		centre := l.TitleBoxLeftMargin + 0.5*l.TitleBoxWidth +
			float64((i))*l.TitleBoxPitch
		left := centre - 0.5*l.TitleBoxWidth
		right := centre + 0.5*l.TitleBoxWidth
		laneInfo := &LaneInfo{left, centre, right}
		l.Individual[statement] = laneInfo
	}
}

// isolateLaneStatements isolates lane statements from a list.
func (l *Lanes) isolateLaneStatements(statements []*dslmodel.Statement) {
	for _, s := range statements {
		if s.Keyword == umli.Lane {
			l.LaneStatements = append(l.LaneStatements, s)
		}
	}
}
