package parser

// ParsedLine is a model for a single line in the DSL input text.
type ParsedLine struct {
	keyWord       string
	lanes         []string
	labelSegments []string
}

// NewParsedLine constructs a ParsedLine ready to use.
func NewParsedLine(keyWord string, lanes []string, labelSegments []string) *ParsedLine {
	return &ParsedLine{keyWord, lanes, labelSegments}
}
