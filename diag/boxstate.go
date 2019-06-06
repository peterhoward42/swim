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

// newBoxState provides a boxState ready to use.
func newBoxState(lifelineStatement *dslmodel.Statement) *boxState {
	return &boxState{
		inProgress: false,
	}
}

// Type allBoxStates maps lifeline statements to a
// boxState.
type allBoxStates map[*dslmodel.Statement]*boxState

// NewLifelineActivityBoxes creates a lifelineActivityBoxes ready to use.
func newAllBoxStates(
	lifelineStatements []*dslmodel.Statement) allBoxStates {
	boxes := allBoxStates{}
	for _, lifelineStatement := range lifelineStatements {
		boxes[lifelineStatement] = newBoxState(lifelineStatement)
	}
	return boxes
}
