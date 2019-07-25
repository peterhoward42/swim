package sizer

// This module defines constants that govern
// sizing decisions. They are all ratios (mostly) w.r.t. font height.

// Naming conventions:
// - ends with <K> means coeffient / ratio
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

const (

	// To explain the first one as an example...
	// the padding at top of the entire diagram should be the same as
	// one rendered string. I.e. 1.0 font height.
	diagramPadTK = 1.0
	diagramPadBK = 1.0
	diagPadLK    = 1.0

	// Frame and diagram title
	framePadLRK         = 0.5
	frameInternalPadBK  = 1.0
	frameTitleTextPadTK = 0.5
	frameTitleTextPadBK = 1.0
	frameTitleTextPadLK = 1.0

	frameTitleBoxWidthK = 0.25 // Of diagram width
	frameTitleRectPadBK = 3.0

	// Lifelines
	ideallifelineTitleBoxWidthK = 15.0
	titleBoxPadRK               = 0.25 // w.r.t title box width
	titleBoxTextPadTK           = 0.25 // Between top of title box and first line txt
	titleBoxTextPadBK           = 0.75 // Between text and the bottom of title box
	titleBoxPadBK               = 1.5  // Below entire title box
	minLifelineSegLengthK       = 0.5

	// Full and Dashed interaction lines
	dashLineDashLenK           = 0.5
	dashLineDashGapK           = 0.25
	selfLoopHeightK            = 3.0
	interactionLinePadBK       = 0.5 // Allows for half arrow height!
	interactionLineTextPadBK   = 0.5 // Between the text and its line
	interactionLineLabelIndent = 1.0 // Offset from arrow

	// Arrows
	arrowLenK        = 1.5
	arrowAspectRatio = 0.4 // Width of arrow head w.r.t. length

	// Lifeline activity boxes
	activityBoxWidthK           = 1.5
	activityBoxVerticalOverlapK = 0.5
	finalizedActivityBoxesPadB  = 1.0
)
