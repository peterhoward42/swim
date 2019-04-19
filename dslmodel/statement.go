package dslmodel

// Statement is an object that can represent any of the input line
// types - and provides attributes for the superset of data each must
// be qualified with.
type Statement struct {
	Keyword    string
	LaneName	string // For <lane> statements.
	FirstLane  *Statement // Lane operand
	SecondLane *Statement // Lane operand
    LabelLines []string // Each line of text in the label
}
