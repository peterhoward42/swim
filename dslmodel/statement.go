package dslmodel

// Statement is an object that can represent any of the input line
// types in the DSL - and provides a superset of attributes for the data each must
// be qualified with.
type Statement struct {
	Keyword         string
	LaneName        string       // For <lane> statements.
	ReferencedLanes []*Statement // When lane operands are present
	LabelSegments   []string     // Each line of text in the label
}
