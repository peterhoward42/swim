package dsl

import (
	"github.com/peterhoward42/umli"
)

/*
Model encapsulates a machine readable, and structured model of a
parsed DSL script, in terms of a sequence of Statement(s).
*/
type Model struct {
	statements []*Statement
}

// Statements provides an order list of the statements held.
func (m *Model) Statements() []*Statement {
	return m.statements
}

/*
Append adds a new statement to the list of those already held.
*/
func (m *Model) Append(s *Statement) {
	m.statements = append(m.statements, s)
}

/*
AddLifelineLetters alters all the titles for all lifeline statements, by
inserting a new label segment at the beginning of the existing label segments,
comprising the lifeline's name (letter).
*/
func (m *Model) AddLifelineLetters() {
	for _, s := range m.LifelineStatements() {
		addition := []string{"", s.LifelineName}
		s.LabelSegments = append(s.LabelSegments, addition...)
	}
}

/*
LifelineStatements provides the subset of statements held that are
*life* statements - in the order in which they appear in the script.
*/
func (m *Model) LifelineStatements() []*Statement {
	var statements []*Statement
	for _, s := range m.statements {
		if s.Keyword == umli.Life {
			statements = append(statements, s)
		}
	}
	return statements
}

/*
LifelineStatementByName finds among the lifeline statements held, the one
with the given lifeline name. (Or returns nil)
*/
func (m *Model) LifelineStatementByName(name string) (s *Statement, ok bool) {
	for _, s := range m.LifelineStatements() {
		if s.LifelineName == name {
			return s, true
		}
	}
	return nil, false
}

/*
FirstStatementOfType provides the first statement in those held that has
the given keyword (if ok). (Use types such as umli.Full as the parameter),
*/
func (m *Model) FirstStatementOfType(statementType string) (
	s *Statement, ok bool) {
	for _, s := range m.statements {
		if s.Keyword == statementType {
			return s, true
		}
	}
	return nil, false
}

/*
LifelineIsKnown returns true if the model contains a lifeline with the
given name.
*/
func (m *Model) LifelineIsKnown(name string) bool {
	_, ok := m.LifelineStatementByName(name)
	return ok
}

// SizeFromTextStatement provides the text size specified by the first
// textsize statement, (or returns ok false if there isn't one).
func (m *Model) SizeFromTextStatement() (sz float64, ok bool) {
	s, ok := m.FirstStatementOfType(umli.TextSize)
	if !ok {
		return 0, false
	}
	return s.TextSize, true
}

/*
Title provides the title specified in a title statement or
[]string{"Title Unspecified"}
*/
func (m *Model) Title() []string {
	s, ok := m.FirstStatementOfType(umli.Title)
	if !ok {
		return []string{"Title Unspecified"}
	}
	return s.LabelSegments
}

// LifelineLettersSupressed returns true if there is an explict don't-show
// lifeline letters statement
func (m *Model) LifelineLettersSupressed() bool {
	s, ok := m.FirstStatementOfType(umli.ShowLetters)
	if !ok {
		return false
	}
	return !s.ShowLetters
}
