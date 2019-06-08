package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// Type boxState keeps track of the in-progress state of a lifeline
// activity box (for one lifeline) during diagram creation.
type boxState struct {
	inProgress bool
	topY       float64
}

// newBoxStates provides a boxState ready to use.
func newBoxStates(lifelineStatement *dslmodel.Statement) *boxState {
	return &boxState{
		inProgress: false,
	}
}

// Type boxStates maps lifeline statements to a
// boxState.
type boxStates map[*dslmodel.Statement]*boxState

// NewLifelineActivityBoxes creates a lifelineActivityBoxes ready to use.
func newAllBoxStates(
	lifelineStatements []*dslmodel.Statement) boxStates {
	boxes := boxStates{}
	for _, lifelineStatement := range lifelineStatements {
		boxes[lifelineStatement] = newBoxStates(lifelineStatement)
	}
	return boxes
}
