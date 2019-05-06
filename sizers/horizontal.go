package sizers

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
	umli "github.com/peterhoward42/umlinteraction"

)

// Horizontal is the master owner of all horizontal sizing and spacing.
type Horizontal struct {
	diagWidth   int
	fontHeight  float64
	statements  []*dslmodel.Statement
	laneStatements []*dslmodel.Statement
	numLanes int

	LaneCentres LaneCentres
	LaneTitleBoxes LaneTitleBoxes
}

// NewHorizontal creates a Horizontal ready to use.
func NewHorizontal(diagWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Horizontal {
	h := &Horizontal{}
	h.diagWidth = diagWidth
	h.fontHeight = fontHeight
	h.statements = statements
	h.laneStatements = h.collectLaneStatements()
	h.numLanes = len(h.laneStatements)

	h.LaneCentres = h.calcLaneCentres()
	h.LaneTitleBoxes = h.calcLaneTitleBoxes()
	return h
}

// LaneCentres maps a Lane's DSL Statement to its center coordinate.
type LaneCentres = map[*dslmodel.Statement]int

// LaneTitleBoxes maps a Lane's DSL Statement to its geometry.
type LaneTitleBoxes = map[*dslmodel.Statement]*LaneTitleBox

// LaneTitleBox holds the left and right coordinate for a Lane
// title box.
type LaneTitleBox struct {
	Left  int
	Right int
}

func (h *Horizontal) calcLaneCentres() (
	titleBoxWidth int, LaneCentres LaneCentres) {
	// All lane title boxes are the same width <titleBoxWidth>.
	// These boxes are separated by a gutter of width <B>.
	// titleBoxWidth and <B> are in proportion: B = K * titleBoxWidth
	// The margin between the left and right-most title boxes and the diagram
	// edge is also <B>.
	const K float64 = 0.25
	N := float64(h.numLanes)
	boxWidth = float64(h.diagWidth) / (N + K * (N + 1))
	B := K * boxWidth
	lanePitch := boxWidth + B
	leftMargin := B
	lc := LaneCentres{}
	for i, laneStatement := range h.laneStatements {
		centre := leftMargin + float64(i) * lanePitch
		lc[laneStatement] = int(centre)
	}
	return int(boxWidth), lc
}

// Isolate the lane statements in a list.
func (h *Horizontal) collectLaneStatements() []*dslmodel.Statement {
	ls := []*dslmodel.Statement{}
	for _, s := range h.statements {
		if s.Keyword == umli.Lane {
			ls = append(ls, s)
		}
	}
	return ls
}

func (h *Horizontal) calcLaneTitleBoxes() LaneTitleBoxes {
	boxes := LaneTitleBoxes{}
	for _ s := h.laneStatements {
		centre := h.LaneCentres[s]
		tb := LaneTitleBox{}
		tb.Left = centre - 0.5 * h.titleBoxWidth
		tb.Left = centre - 0.5 * h.titleBoxWidth
		boxes[s] = tb
	}
	return boxes
}

