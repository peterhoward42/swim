package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
)

// eventsForStatements provides a lookup table for the graphical element
// drawing events that are required for a set of DSL statements.
type eventsForStatements = map[*dslmodel.Statement][]eventType

// scanner is capable of receiving a list of Statements, and producing
// the list of graphical element drawing mandates that correspond to each.
type scanner struct {
}

// newScanner produces a Scanner ready to use.
func newScanner() *scanner {
	return &scanner{}
}

// Scan consumes the given list of Statements and captures the corresponding
// drawing-element mandates.
func (s *scanner) Scan(statements []*dslmodel.Statement) eventsForStatements {
	eventsLookup := eventsForStatements{}
	for _, statement := range statements {
		eventsForStatement := EventsRequired[statement.Keyword]
		eventsLookup[statement] = eventsForStatement
	}
	return eventsLookup
}
