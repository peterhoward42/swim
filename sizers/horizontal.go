package sizers

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
)

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

// Horizontal is the master owner of all horizontal sizing and spacing.
type Horizontal struct {
	diagWidth   int
	fontHeight  float64
	statements  []*dslmodel.Statement
	LaneCentres LaneCentres
}

// NewHorizontal creates a Horizontal ready to use.
func NewHorizontal(diagWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Horizontal {
	h := &Horizontal{}
	h.diagWidth = diagWidth
	h.fontHeight = fontHeight
	h.statements = statements
	h.LaneCentres = h.calcLaneCentres()
	return h
}

func (h *Horizontal) calcLaneCentres() LaneCentres {
	// basic pitch is based on diag width / nLanes
	// but lanes with a self need more room
	// draw picture and have const for self arrow side step as f(lane pitch)
	return nil
}
