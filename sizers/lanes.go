package sizers

import (
	umli "github.com/peterhoward42/umlinteraction"
	"github.com/peterhoward42/umlinteraction/dslmodel"
)

// Lanes holds sizing information for the lanes.
type Lanes struct {
	DiagramWidth       float64
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
	lanes.populateLaneStatements(statements)
	lanes.NumLanes = len(lanes.LaneStatements)
	lanes.populateTitleBoxAttribs()
	lanes.populateIndividualLaneInfo()

	return lanes
}

// populateTitleBoxAttribs works out the values for the TitleBoxXXX attributes.
func (l *Lanes) populateTitleBoxAttribs() {
	// The title boxes are all the same width.
	// The gaps between them are a fixed proportion of their width.
	// The margins from the edge of the diagram is the same as this gap.
	const gapProportion float64 = 0.25
	N := float64(l.NumLanes)
	l.TitleBoxWidth = l.DiagramWidth / (N + gapProportion*(N+1))
	l.TitleBoxHorizGap = gapProportion * l.TitleBoxWidth
	l.TitleBoxPitch = l.TitleBoxWidth + l.TitleBoxHorizGap
	l.TitleBoxLeftMargin = l.TitleBoxHorizGap
}

func (l *Lanes) populateIndividualLaneInfo() {
	l.Individual = InfoPerLane{}
	for i, statement := range l.LaneStatements {
		centre := l.TitleBoxLeftMargin + float64(i)*l.TitleBoxPitch
		left := centre - 0.5*l.TitleBoxWidth
		right := centre + 0.5*l.TitleBoxWidth
		laneInfo := &LaneInfo{left, centre, right}
		l.Individual[statement] = laneInfo
	}
}

// populateLaneStatements isolates lane statements from a list.
func (l *Lanes) populateLaneStatements(statements []*dslmodel.Statement) {
	for _, s := range statements {
		if s.Keyword == umli.Lane {
			l.LaneStatements = append(l.LaneStatements, s)
		}
	}
}
