package dsl

// Statement is an object that can represent any of the input line
// types in the DSL - and provides a superset of attributes required.
type Statement struct {
	Keyword             string       // E.g. "full|stop"
	LifelineName        string       // Only used for <life> statements.
	ReferencedLifelines []*Statement // When lifeline operands are present
	LabelSegments       []string     // Each line of text called for in the label
	TextSize            float64      // Only used for <textsize> statements.
	ShowLetters         bool         // Only used for <showletters> statements.
}

// NewStatement instantiates a Statement, ready to use.
func NewStatement() *Statement {
	return &Statement{
		ReferencedLifelines: []*Statement{},
		LabelSegments:       []string{},
	}
}
