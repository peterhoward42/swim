package parser

// ParsedLine is a model to capture the fields in a single line in the DSL input text.
type ParsedLine struct {
	FullText	string
	KeyWord       string
	Lanes         []string
	LabelSegments []string
}

// NewParsedLine constructs a ParsedLine ready to use.
func NewParsedLine(fullText string, keyWord string, lanes []string, labelSegments []string) *ParsedLine {
	return &ParsedLine{fullText, keyWord, lanes, labelSegments}
}
