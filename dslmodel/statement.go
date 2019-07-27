// Package dslmodel provides the Statement type, which makes it possible
// to compose a machine readable,**structured** representation of a DSL 
// script by using a list Statement(s).
// It aims to decouple the downstream diagram building
// system from the Parser and to provide it with a clean and fully
// validated input model.
package dslmodel

// Statement is an object that can represent any of the input line
// types in the DSL - and provides a superset of attributes required.
type Statement struct {
	Keyword             string       // E.g. "full|stop"
	LifelineName        string       // Only used for <life> statements.
	ReferencedLifelines []*Statement // When lifeline operands are present
	LabelSegments       []string     // Each line of text called for in the label
}

// NewStatement instantiates a Statement, ready to use.
func NewStatement() *Statement {
	return &Statement{
		ReferencedLifelines: []*Statement{},
		LabelSegments:       []string{},
	}
}
