package sizers

// This module defines constants that govern sizing decisions.
// They are all ratios with names ending in <K>.
// Vertical ratios are w.r.t. font height.

const diagramTopMarginK = 1.0

// Lane title boxes.
// They are all the same (implied) width, equi-spaced according to
// these settings.
const (
	titleBoxSeparationK    = 0.25 // w.r.t title box width
	titleBoxTextTopMarginK = 0.25 // Between top of title box and first line txt
	titleBoxTextBotMarginK = 0.75 // Between bot title box and last line txt

)
