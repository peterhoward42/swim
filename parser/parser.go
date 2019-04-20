package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	umli "github.com/peterhoward42/umlinteraction"
)

// Parser is capable of parsing lines of DSL text to provide a more
// conventient, model based representation of the text.
type Parser struct{}

// Parse is the parsing invocation method.
func (p *Parser) Parse(input *bufio.Scanner) ([]*ParsedLine, error) {
	lines := []*ParsedLine{}
	for input.Scan() {
		line := input.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parsedLine, err := p.parseLine(line)
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

// Example input: "dash BC  foo bar | baz"

// Regex for choice of keywords - e.g. <dash>
var kwRe = "(" + strings.Join(umli.AllKeywords, "|") + ")"

// Regex for (in this case) <BC>. Allows one or two letters.
const lanesOperandRe = `([A-Z][A-Z]?)`

// Everything that's left over.
const theRestRe = `(.*$)`

// Put it all together expecting whitespace between <dash> and <BC>
var lineRe = regexp.MustCompile(kwRe + `\s+` + lanesOperandRe + `\s*` + theRestRe)

func (p *Parser) parseLine(line string) (*ParsedLine, error) {
	segments := lineRe.FindStringSubmatch(line)
	if len(segments) == 0 {
		return nil, fmt.Errorf("parseLine() input line malformed: %v", line)
	}
	kw := segments[1]
	lanes := strings.Split(segments[2], "") // Splits into constituent letters.
	labelSegments := []string{}             // Split label section at pipe characters.
	for _, seg := range strings.Split(segments[3], "|") {
		seg := strings.TrimSpace(seg)
		if len(seg) != 0 {
			labelSegments = append(labelSegments, seg)
		}
	}

	pl := NewParsedLine(line, kw, lanes, labelSegments)
	return pl, nil
}
