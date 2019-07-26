/*
Package sizer is the single source of truth for sizing all the elements
that make up the diagram. Not only how big things are, but also how far
apart they should be.

E.g. the coordinates for each lifeline title box,
the mark-space settings for dashed lines, arrow sizing, the margins or
padding required for each thing etc.

It is encapsulated in this dedicated package, to remove this responsibility
from the umli.diag package, so that umli.diag need only deal with the
algorithmic part of diagram creation.

*/
package sizer

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// Naming conventions:
// Consider the name: <XxxPadT>
// This should be read as the padding required by
// thing Xxxx at the top (T). Where T is from {LRTB}

// Sizer is the top level sizing component.
type Sizer struct {
	// Whole diagram scope
	DiagramPadT float64 // Margin above frame rectangle.
	DiagramPadB float64 // Margin below frame rectangle.

	// Outer frame and diagram title
	FramePadLR         float64 // Pushes frame slighly inside from edge of diagram.
	FrameInternalPadB  float64 // Pushes frame bottom below ends of lifelines.
	FrameTitleTextPadT float64 // Gives frame title text headroom.
	FrameTitleTextPadB float64
	FrameTitleTextPadL float64
	FrameTitleBoxWidth float64
	FrameTitleRectPadB float64 // Creates space below the frame title box.

	// Lifeline titles boxes
	IdealLifelineTitleBoxWidth float64 // Potentially moderated later.
	TitleBoxHeight             float64
	TitleBoxPadB               float64
	TitleBoxLabelPadB          float64
	TitleBoxLabelPadT          float64

	// Interaction lines
	InteractionLinePadB        float64
	InteractionLineTextPadB    float64
	SelfLoopHeight             float64
	InteractionLineLabelIndent float64

	// Arrows and dashes
	ArrowLen        float64
	ArrowHeight     float64
	DashLineDashLen float64
	DashLineDashGap float64

	// Activity boxes
	ActivityBoxWidth           float64
	ActivityBoxVerticalOverlap float64
	FinalizedActivityBoxesPadB float64
	MinLifelineSegLength       float64
}

// NewSizer provides a Sizer structure that has been initialised
// as is ready for use.
func NewSizer(diagramWidth int, fontHeight float64,
	lifelineStatements []*dslmodel.Statement) *Sizer {
	sizer := &Sizer{}

	// The requested font height is used as a datum reference,
	// and nearly everything is sized in proportion to this.

	fh := fontHeight

	// Whole diagram scope
	sizer.DiagramPadT = 1.0 * fh
	sizer.DiagramPadB = 1.0 * fh

	// Outer frame and diagram title
	sizer.FramePadLR = 0.5 * fh
	sizer.FrameInternalPadB = 1.0 * fh
	sizer.FrameTitleTextPadT = 0.5 * fh
	sizer.FrameTitleTextPadB = 1.0 * fh
	sizer.FrameTitleTextPadL = 1.0 * fh
	sizer.FrameTitleBoxWidth = 0.25 * float64(diagramWidth)
	sizer.FrameTitleRectPadB = 1.0 * fh

	// Lifeline title boxes
	sizer.TitleBoxLabelPadT = 0.25 * fh
	sizer.TitleBoxLabelPadB = 1.0 * fh
	sizer.IdealLifelineTitleBoxWidth = 15.0 * fh
	sizer.TitleBoxPadB = 1.5 * fh

	// Interaction lines
	sizer.ArrowLen = 1.5 * fh
	sizer.ArrowHeight = 0.4 * sizer.ArrowLen
	sizer.InteractionLinePadB = 0.5 * fh
	sizer.InteractionLineTextPadB = 0.5 * fh
	sizer.InteractionLineLabelIndent = sizer.ArrowLen + 1.0*fh
	sizer.SelfLoopHeight = 3.0 * fh

	// Dashes
	sizer.DashLineDashLen = 0.5 * fh
	sizer.DashLineDashGap = 0.25 * fh

	// Activity boxes
	sizer.ActivityBoxWidth = 1.5 * fh
	sizer.ActivityBoxVerticalOverlap = 0.5 * fh
	sizer.FinalizedActivityBoxesPadB = 1.0 * fh

	// Lifelines
	sizer.MinLifelineSegLength = 0.5 * fh

	return sizer
}
