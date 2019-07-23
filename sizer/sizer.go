/*
Package sizer is the single source of truth for sizing all the elements
that make up the diagram. Not only how big things are, but also how far
apart they should be.

E.g. the coordinates for each lifeline title box,
the mark-space settings for dashed lines, arrow sizing, the margins or
padding required for each thing etc.

It is encapsulated in this dedicated package, to remove this responsibility
from the umli.diag package, so that umli.diag can deal only with the
algorithmic part of diagram creation.

It provides the top level *Sizer* type, along with some subordinate types
it delegates to. For example: sizing.Lifelines.
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
	DiagramPadT float64
	DiagramPadB float64
	DiagPadL    float64

	// Outer frame and diagram title
	FrameInternalPadB  float64
	FrameTitleTextPadT float64
	FrameTitleTextPadB float64
	FrameTitleTextPadL float64
	FrameTitleBoxWidth float64
	FrameTitleRectPadB float64

	// Lifeline titles boxes
	IdealLifelineTitleBoxWidth float64
	TitleBoxHeight             float64
	TitleBoxPadB               float64
	TitleBoxLabelPadB          float64
	TitleBoxLabelPadT          float64

	// Interaction lines
	InteractionLinePadB        float64
	InteractionLineTextPadB    float64
	InteractionLineLabelIndent float64
	SelfLoopHeight             float64

	// Arrows and dashes
	ArrowLen        float64
	ArrowHeight     float64
	DashLineDashLen float64
	DashLineDashGap float64

	// Activity boxes
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
	// and nearly everything is sized in proportion to this, using
	// settings from settings.go.

	// These settings values typically end with the letter
	// K (e.g. diagramPadTK) to indicate they are proportion coefficients.

	fh := fontHeight

	// Whole diagram scope
	sizer.DiagramPadT = diagramPadTK * fh
	sizer.DiagramPadB = diagramPadBK * fh
	sizer.DiagPadL = diagPadLK * fh

	// Outer frame and diagram title
	sizer.FrameInternalPadB = frameInternalPadBK * fh
	sizer.FrameTitleTextPadT = frameTitleTextPadTK * fh
	sizer.FrameTitleTextPadB = frameTitleTextPadBK * fh
	sizer.FrameTitleTextPadL = frameTitleTextPadLK * fh
	sizer.FrameTitleBoxWidth = frameTitleBoxWidthK * float64(diagramWidth)
	sizer.FrameTitleRectPadB = frameTitleRectPadBK * fh

	// Lifeline title boxes
	sizer.TitleBoxLabelPadT = titleBoxTextPadTK * fh
	sizer.TitleBoxLabelPadB = titleBoxTextPadBK * fh
	sizer.IdealLifelineTitleBoxWidth = ideallifelineTitleBoxWidthK * fh
	sizer.TitleBoxPadB = titleBoxPadBK * fh

	// Interaction lines
	sizer.ArrowLen = arrowLenK * fh
	sizer.ArrowHeight = arrowAspectRatio * sizer.ArrowLen
	sizer.InteractionLinePadB = interactionLinePadBK * fh
	sizer.InteractionLineTextPadB = interactionLineTextPadBK * fh
	sizer.InteractionLineLabelIndent = sizer.ArrowLen +
		interactionLineLabelIndent*fh
	sizer.SelfLoopHeight = selfLoopHeightK * fh

	// Dashes
	sizer.DashLineDashLen = dashLineDashLenK * fh
	sizer.DashLineDashGap = dashLineDashGapK * fh

	// Activity boxes
	sizer.ActivityBoxVerticalOverlap = activityBoxVerticalOverlapK * fh
	sizer.FinalizedActivityBoxesPadB = finalizedActivityBoxesPadB * fh

	// Lifelines
	sizer.MinLifelineSegLength = minLifelineSegLengthK * fh

	return sizer
}
