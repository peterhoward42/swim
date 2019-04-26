// Package parser provides the Parser type, which is capable of parsing lines
// of DSL text to provide a structured representation of it - in the form
// of a slice of dslmodel.Statement(s).
package parser

import (
	"bufio"
	"strings"

	umli "github.com/peterhoward42/umlinteraction"
	 "github.com/peterhoward42/umlinteraction/dslmodel"
)

// Parser is capable of parsing lines of DSL text to provide a more
// convenient, model based representation of the text. It does not
// thoroughly validate what it finds - leaving that to become self
// evident when it is interpreted.
type Parser struct{}

// Parse is the parsing invocation method.
func (p *Parser) Parse(input *bufio.Scanner) ([]*dslmodel.Statement, error) {
	statements := []*dslmodel.Statement{}
	lineNo := 0
	for input.Scan() {
		line := input.Text()
		lineNo++
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		statement, err := p.parseLine(line, lineNo)
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}
	if err := input.Err(); err != nil {
		return nil, err
	}
	return statements, nil
}

// parseLine parses the text present in a single line of DSL, into
// the fields expected, and packages the result into a dslmodel.Statement.
func (p *Parser) parseLine(line string, lineNo int) (*dslmodel.Statement, error) {
	words := strings.Split(line, " ")
	if len(words) < 2 {
		return nil, umli.DSLError(line, lineNo, "Must have at least 2 words.")
	}
	keyWord := words[0]
	lanesReferenced := strings.Split(words[1], "")
	// Isolate label text by stripping what we have already consumed.
	labelText := strings.Replace(line, keyWord, "", 1)
	labelText = strings.Replace  (labelText, words[1], "", 1)
	labelIndividualLines := p.isolateLabelConstituentLines(labelText)
	return dslmodel.NewStatement(
		line, lineNo, keyWord, lanesReferenced, labelIndividualLines), nil
}

// isolateLabelConstituentLines takes the label text from a DSL line and
// splits it into the constituent lines according to its author's intent.
// I.e. by splitting it at "|" delimiters. Note the removal of whitespace
// either side of any "|" present.
// E.g. From this: "edit_facilities( | payload, user_token)"
// It produces: []string{"edit_facilities(", "payload, user_token)"}
func (p *Parser) isolateLabelConstituentLines(labelText string) []string {
	segments := strings.Split(labelText, "|")
	constituentLines := []string{}
	for _, seg := range segments {
		seg := strings.TrimSpace(seg)
		if len(seg) != 0 {
			constituentLines = append(constituentLines, seg)
		}
	}	
	return constituentLines
}
