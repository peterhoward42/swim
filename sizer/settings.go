package sizers

// This module defines (package private) constants that govern
// sizing decisions. They are all ratios (mostly) w.r.t. font height.

// Naming conventions:
// - ends with <K> means coeffient / ratio
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

const (

	// Diagram as a whole.
	diagramPadTK = 1.0
	diagramPadBK = 1.0

	// Lane title boxes.
	titleBoxPadRK     = 0.25 // w.r.t title box width
	titleBoxTextPadTK = 0.25 // Between top of title box and first line txt
	titleBoxTextPadBK = 0.75 // Between text and the bottom of title box
	titleBoxPadBK     = 0.5  // Below entire title box

	// Full and Dashed interaction lines
	dashLineDashLenK = 0.5
	dashLineDashGapK = 0.25

	interactionLinePadBK     = 0.5 // Allows for half arrow height!
	interactionLineTextPadBK = 0.5 // Between the text and its line
	arrowLenK                = 1.5
	arrowAspectRatio         = 0.4 // Width of arrow head w.r.t. length

	// Self interactions
	selfLoopAspectRatio = 0.6 // Width ratio to height of the three lines

	// Lifeline activity boxes
	activityBoxWidthK   = 2.0
	activityBoxTopPadBK = 0.5
)
