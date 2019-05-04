package diag

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
)

// Scanner is capable of receiving a list of Statements, and capturing
// the list of graphical element drawing mandates that correspond to each.
type Scanner struct {
	Events map[*dslmodel.Statement][]EventType
}

// NewScanner produces a Scanner ready to use.
func NewScanner() *Scanner{
	return &Scanner{Events: map[*dslmodel.Statement][]EventType{}}
}

// Scan consumes the given list of Statements and stores the corresponding
// drawing-element mandates in its internal look up table.
func (s *Scanner) Scan(statements []*dslmodel.Statement) {
	for _, statement := range(statements) {
		events := EventsRequired[statement.Keyword]
		s.Events[statement] = events
	}
}