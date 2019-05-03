/*
events.go provides the types that represent the drawing steps to be used when
building the interaction diagram.

The diagram building process will iterate over DSL input statements in sequence
and decide what new diagram drawing mandates are required arising from
each statement.

For example to start a box, or to drawn an arrow from one lane to another.

Types are used to represent these drawing mandates so that:
	- They can be parameterized and hold state
	- They can become a lingua-franca for collaborating types
	- They can provide a structured and de-coupled input to the downstream
	  rendering process
*/

package diag

import (
)

// EndBox is an instruction to close-off or terminate the drawing one of the
// thin vertical boxes that sits on a lane to emit and receive interactions.
type EndBox struct {
	whichBox *StartBox
}

// Interaction is an instruction to draw a horizontal line from one lane
// to another.
type Interaction struct {
	fromLane *LaneLine
	toLane *LaneLine
	dashed bool
}

// InteractionLabel is an instruction to create a multi-line label that belongs
// to an Interaction object. I.e. the horizontal line it must sit above.
type InteractionLabel struct {
	labelSegments []string
	interaction Interaction
}

// LaneLine is an instruction to draw the vertical dashed line that
// represents one lane.
type LaneLine struct {
	laneTitleBox *LaneTitleBox
}

// LaneTitleBox is an instruction to draw one of the lane title boxes that
// sit along the top of the diagram.
type LaneTitleBox struct {
	labelSegments []string
}

// SelfInteraction is an instruction to draw the series of lines required to
// represent an internal interaction on a lane. I.e. 3 sides of a rectangle.
type SelfInteraction struct {
	lane *LaneLine
}

// StartBox is an instruction to start drawing one of the thin vertical boxes
// that sits on a lane to emit and receive interactions.
type StartBox struct {
	lane *LaneLine
}
