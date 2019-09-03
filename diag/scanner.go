package diag

/*
This module contains code that is capable of consuming an ordered list
of DSL statements - and producing the required sequence of drawing events (or
mandates) to make the diagram. It's called a scanner, because the process is
conceptually concerned with working it way down the page incrementally as it 
produces things - according to the space taken up by each graphics event.
*/

import (
	"github.com/peterhoward42/umli/dsl"
)

// eventsForStatements provides a lookup table for the graphical element
// drawing events that are required for a set of DSL statements.
type eventsForStatements = map[*dsl.Statement][]eventType

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
func (s *scanner) Scan(statements []*dsl.Statement) eventsForStatements {
	eventsLookup := eventsForStatements{}
	for _, statement := range statements {
		eventsForStatement := EventsRequired[statement.Keyword]
		eventsLookup[statement] = eventsForStatement
	}
	return eventsLookup
}
