package sizers

// This module defines constants that govern sizing decisions.
// They are all ratios with names ending in <K>.
// Vertical ratios are w.r.t. font height.

const diagramTopMarginK = 1.0

// Lane title boxes.
// They are all the same (implied) width, equi-spaced according to
// these settings.
const (
	titleBoxSeparationK       = 0.25 // w.r.t title box width
	titleBoxTextRowLeadingK   = 0.25 // Leading as in typography
	titleBoxTextTopBotMarginK = 0.5  // Between title box and first line txt
)
