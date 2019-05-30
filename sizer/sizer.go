/*
Package sizers is the single source of truth for sizing all the elements
that make up the diagram. E.g. the coordinates for each lane title box,
the mark-space settings for dashed lines, arrow sizing, the margins or
padding required for each thing etc.

It is encapsulated in this dedicated package, to remove this responsibility
from the umli.diag package, so that umli.diag can deal only with the
algorithmic part of diagram creation.

It provides the top level *Sizer* type, along with some subordinate types
it delegates to. For example: sizing.Lanes.
*/
package sizers

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// Naming conventions:
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

// Sizer is the top level sizing component.
type Sizer struct {
	DiagramPadT             float64
	DiagramPadB             float64
	Lanes                   *Lanes
	InteractionLinePadB     float64
	InteractionLineTextPadB float64
	ArrowLen                float64
	ArrowHeight             float64
	DashLineDashLen         float64
	DashLineDashGap         float64
	SelfLoopHeight          float64
	ActivityBoxTopPadB      float64
}

// NewSizer provides a Sizer structure that has been initialised
// as is ready for use.
func NewSizer(diagramWidth int, fontHeight float64,
	lifelineStatements []*dslmodel.Statement) *Sizer {
	sizer := &Sizer{}
	sizer.DiagramPadT = diagramPadTK * fontHeight
	sizer.DiagramPadB = diagramPadBK * fontHeight
	sizer.Lanes = NewLanes(diagramWidth, fontHeight, lifelineStatements)
	sizer.InteractionLinePadB = interactionLinePadBK * fontHeight
	sizer.InteractionLineTextPadB = interactionLineTextPadBK * fontHeight
	sizer.ArrowLen = arrowLenK * fontHeight
	sizer.ArrowHeight = arrowAspectRatio * sizer.ArrowLen
	sizer.DashLineDashLen = dashLineDashLenK * fontHeight
	sizer.DashLineDashGap = dashLineDashGapK * fontHeight
	sizer.SelfLoopHeight = sizer.Lanes.SelfLoopWidth * selfLoopAspectRatio
	sizer.ActivityBoxTopPadB = activityBoxTopPadBK * fontHeight

	return sizer
}
