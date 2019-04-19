package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

// Parser is capable of parsing lines of DSL text to provide a more
// conventient, model based representation.
type Parser struct{}

// Parse is the parsing invocation method.
func (p *Parser) Parse(input *bufio.Scanner) (parseLines []ParsedLine, err error) {
	lines := []*ParsedLine{}
	for input.Scan() {
		line := input.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parsedLine, err := p.parseLine(line)
		if err != nil {
			s := err.Error()
			fmt.Printf("XXXX: %v", s)
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

const kwRe = `(lane|full|dash|stop|self)`
const lanesOperandRe = `([A-Z][A-Z]?)` // ? means zero or one
const theRestRe = `(.*$)`

var lineRe = regexp.MustCompile(kwRe + `\s+` + lanesOperandRe + `\s*` + theRestRe)

func (p *Parser) parseLine(line string) (*ParsedLine, error) {
	// Example input: "dash BC  foo bar | baz"
	segments := lineRe.FindStringSubmatch(line)
	if len(segments) == 0 {
		return nil, fmt.Errorf("parseLine() input line malformed: %v", line)
	}
	kw := segments[1]
	lanes := strings.Split(segments[2], "")
	labelSegments := []string{}
	for _, seg := range strings.Split(segments[3], "|") {
		seg := strings.TrimSpace(seg)
		if len(seg) != 0 {
			labelSegments = append(labelSegments, seg)
		}
	}

	pl := NewParsedLine(kw, lanes, labelSegments)
	return pl, nil
}
