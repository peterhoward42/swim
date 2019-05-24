package sizers

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// Naming conventions:
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

// Sizer is the top level sizing component.
type Sizer struct {
	DiagramPadT                float64
	Lanes                   *Lanes
	InteractionLinePadB     float64
	InteractionLineTextPadB float64
	ArrowLen	float64
	ArrowHeight float64
	DashLineDashLen         float64
	DashLineDashGap float64
}

// NewSizer provides a Sizer structure that has been initialised
// as is ready for use.
func NewSizer(diagramWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Sizer {
	sizer := &Sizer{}
	sizer.DiagramPadT = diagramPadTK * fontHeight
	sizer.Lanes = NewLanes(diagramWidth, fontHeight, statements)
	sizer.InteractionLinePadB = interactionLinePadBK * fontHeight
	sizer.InteractionLineTextPadB = interactionLineTextPadBK * fontHeight
	sizer.ArrowLen = arrowLenK * fontHeight
	sizer.ArrowHeight = arrowAspectRatio * sizer.ArrowLen
	sizer.DashLineDashLen = dashLineDashLenK * fontHeight
	sizer.DashLineDashGap = dashLineDashGapK * fontHeight
	return sizer
}
