// Package dslmodel provides the Statement type, which makes it possible
// to compose a structured representation of a DSL script by using a list
// of Statement(s). It aims to decouple the downstream diagram building
// system from the Parser and to provide it with a clean and fully
// validated input model.
package dslmodel

// Statement is an object that can represent any of the input line
// types in the DSL - and provides a superset of attributes for the data each
// must be qualified with.
type Statement struct {
	Keyword         string
	LaneName        string       // For <lane> statements.
	ReferencedLanes []*Statement // When lane operands are present
	LabelSegments   []string     // Each line of text called for in the label
}

// NewStatement instantiates a Statement, ready to use.
func NewStatement() *Statement {

	return &Statement{
		ReferencedLanes: []*Statement{},
		LabelSegments:   []string{},
	}
}
