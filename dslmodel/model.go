/*
Package dslmodel encapsulates a programmatic, higher-level model of a DSL script.
It aims to provide something more convenient for the diagram building
stages to consume as input than the text form, and to decouple the diagram
building system from the DSL Parser.
*/
package dslmodel

import (
"regexp"
	umli "github.com/peterhoward42/umlinteraction"
)

// Model is the primary model class for the package.
// Representing a DSL script as an ordered sequence of Statement(s).
type Model struct {
	Statements []*Statement
	// This map provides a lookup for the Statements that define lanes. Keyed on
	// the lane letter.
	LaneLookup map[string]*Statement
}


var singleLetterRe = regexp.MustCompile(`[A-Z]`)

// NewModel constructs a usable Model.
func NewModel(statements []*Statement) (*Model, error) {
	model := &Model {statements, map[string]*Statement{}}
	// Populate the lookup table of lane letters to the corresponding
	// Statement.
	for _, statement := range model.Statements {
		// The parser has made the statement's ReferencedLanes single character
		// strings already, but we must check there is only one, and that it
		// is a capital letter.
		if statement.Keyword == umli.Lane {
			if len(statement.ReferencedLanes) != 1 {
				return nil, umli.DSLError(statement.DSLText, statement.LineNo, 
				"Lane name must a be a single capital letter.")
			}
			laneName := statement.ReferencedLanes[0]
			if !singleLetterRe.MatchString(laneName) {
				 return nil, umli.DSLError(statement.DSLText, statement.LineNo, 
					"Lane name must a be a single capital letter.")
			}
			model.LaneLookup[laneName] = statement
		}
	}
	return model, nil
}

