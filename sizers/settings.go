package sizers

// This module defines constants that govern sizing decisions.
// They are all ratios (mostly) w.r.t. font height.

// Naming conventions:
// - ends with <K> means coeffient / ratio
// - begins with the graphics entity if applies to
// - the fragment <PadT> should be read as paddingTop (where T is from {LRTB})

const diagramPadTK = 1.0

// Lane title boxes.
// They are all the same (implied) width, equi-spaced according to
// these settings.
const (
	// Lane Title Boxes
	titleBoxPadRK     = 0.25 // w.r.t title box width
	titleBoxTextPadTK = 0.25 // Between top of title box and first line txt
	titleBoxTextPadBK = 0.75 // Between text and the bottom of title box
	titleBoxPadBK     = 0.5  // Below entire title box

	// Full and Dashed interaction lines

	// Below the line itself - including allowance for the arrow height.
	interactionLinePadBK     = 0.5
	interactionLineTextPadBK = 0.5 // Below the text on an interaction line
	arrowLenK                = 1.5
	arrowAspectRatio         = 0.4 // Width of arrow head w.r.t. length
)
