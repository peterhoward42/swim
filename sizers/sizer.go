package sizers

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// Naming conventions:
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

// Sizer is the top level sizing component.
type Sizer struct {
	DiagPadT                float64
	Lanes                   *Lanes
	InteractionLineTextPadB float64
}

// NewSizer provides a Sizer structure that has been initialised
// as is ready for use.
func NewSizer(diagramWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Sizer {
	sizer := &Sizer{}
	sizer.DiagPadT = diagramPadTK * fontHeight
	sizer.Lanes = NewLanes(diagramWidth, fontHeight, statements)
	sizer.InteractionLineTextPadB =
		interactionLineTextPadBK * fontHeight
	return sizer
}
