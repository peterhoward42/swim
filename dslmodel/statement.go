package dslmodel

// Statement is an object that can represent any of the input line
// types in the DSL - and provides a superset of attributes for the data each must
// be qualified with.
type Statement struct {

	DSLText		string // The originating line of text in the DSL
	LineNo int // The line number in the originating DSL
	Keyword         string
	LaneName        string       // For <lane> statements.
	ReferencedLanes []string // When lane operands are present
	LabelSegments   []string     // Each line of text in the label
}

// NewStatement instantiates a Statement, ready to use.
func NewStatement(line string, lineNo int, keyWord string,
	lanesReferenced []string, labelIndividualLines []string) *Statement{
	return &Statement{
		DSLText :line, 
		LineNo : lineNo,
		Keyword : keyWord,
		ReferencedLanes: lanesReferenced,
		LabelSegments: labelIndividualLines,
	}
}
