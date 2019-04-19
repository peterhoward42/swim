package parser

import (
	"bufio"
)

// Parser is capable of parsing lines of DSL text to provide a more
// conventient, model based representation.
type Parser struct{}

// Parse is the parsing invocation method.
func (p *Parser) Parse(input *bufio.Scanner) (parseLines []ParsedLine, err error) {
	lines := []*ParsedLine{}
	for input.Scan() {
		line := input.Text()
		parsedLine, err := p.parseLine(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, parsedLine)
	}
	if err := input.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}

//----------------------------------------------------------------------------
// Private Below
//----------------------------------------------------------------------------

func (p *Parser) parseLine(line string) (*ParsedLine, error) {
	return nil, nil
}
