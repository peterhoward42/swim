// Package parser provides the Parser type, which is capable of parsing lines
// of DSL text to provide a structured representation of the text fields.
package parser

import (
	"bufio"
	"strings"

	umli "github.com/peterhoward42/umlinteraction"
)

// Parser is capable of parsing lines of DSL text to provide a more
// convenient, model based representation of the text. It does not
// thoroughly validate what it finds - leaving that to become self
// evident when it is interpreted.
type Parser struct{}

// Parse is the parsing invocation method.
func (p *Parser) Parse(input *bufio.Scanner) ([]*dslmodel.Statement, error) {
	lines := []*dslmodel.Statement{}
	lineNo := 0
	for input.Scan() {
		line := input.Text()
		lineNo += 1
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parsedLine, err := p.parseLine(line, lineNo)
		if err != nil {
			return nil, err
		}
		lines = append(lines, parsedLine)
	}
	if err := input.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func (p *Parser) parseLine(line string, lineNo int) (*ParsedLine, error) {
	words := strings.Split(line, " ")
	if len(words < 2) {
		return nil, p.error(line, lineNo, "Must have at least 2 words.")
	}
	keyWord := words[0]
	lanesReferenced := strings.Split(words[1], "")
	labelSegments := p.isolateLabelsConstituentLines(words[2:])
	return dslmodel.NewStatement(line, lineNo, keyWord, lanesReferenced, labelSegments)
}
