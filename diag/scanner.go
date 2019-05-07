package diag

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
)

// EventsForStatements provides a lookup table for the graphical element
// drawing events that are required for a set of DSL statements.
type EventsForStatements = map[*dslmodel.Statement][]EventType

// Scanner is capable of receiving a list of Statements, and producing
// the list of graphical element drawing mandates that correspond to each.
type Scanner struct {
}

// NewScanner produces a Scanner ready to use.
func NewScanner() *Scanner {
	return &Scanner{}
}

// Scan consumes the given list of Statements and captures the corresponding
// drawing-element mandates.
func (s *Scanner) Scan(statements []*dslmodel.Statement) EventsForStatements {
	eventsLookup := EventsForStatements{}
	for _, statement := range statements {
		eventsForStatement := EventsRequired[statement.Keyword]
		eventsLookup[statement] = eventsForStatement
	}
	return eventsLookup
}
