package diag

import (
	"fmt"

	"github.com/peterhoward42/umli/dslmodel"
)

/*
BoxState keeps track of the Lifeline Activity Boxes during the
drawing element creation process as it works through the DSL statements
and works down the page.
*/
type BoxState struct {
	// Boxes are said to be in progress when the interaction that
	// implies their existence has been encountered, but no trigger
	// has yet been encountered to terminate it. The map is keyed
	// on the Lifeline Statements.
	boxesInProgress map[*dslmodel.Statement]bool
}

// NewBoxState provides a BoxState ready to use.
func NewBoxState(lifelineStatements []*dslmodel.Statement) *BoxState {
	b := &BoxState{
		boxesInProgress: map[*dslmodel.Statement]bool{},
	}
	// for debugging
	fmt.Printf("XXXX registering statements in BoxState:\n")
	for _, statement := range lifelineStatements {
		fmt.Printf("XXXX %p\n", statement)
		b.boxesInProgress[statement] = false
	}
	return b
}

// boxIsInProgress yields if a lifeline activity box is in progress
// for the given lifelineStatement
func (bs *BoxState) boxIsInProgress(
	lifelineStatement *dslmodel.Statement) bool {
	fmt.Printf("XXXX boxIsInProgress() test: %p\n", lifelineStatement)
	inProgress, ok := bs.boxesInProgress[lifelineStatement]
	if !ok {
		return false
	}
	return inProgress
}
