// Package parser provides the Parse function, which is capable of parsing lines
// of DSL text to provide a structured representation of it - in the form
// of a slice of dslmodel.Statement(s).
package parser

import (
	"bufio"
	"errors"
	"fmt"
	re "regexp"
	"strconv"
	"strings"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
)

/*
Parser provides the service of parsing the DSL into dslmodel.Statements.
*/
type Parser struct {
	inputScript              string
	lifelineStatementsByName map[string]*dslmodel.Statement
}

/*
NewParser returns a Parser ready to use.
*/
func NewParser(inputScript string) *Parser {
	return &Parser{
		inputScript:              inputScript,
		lifelineStatementsByName: map[string]*dslmodel.Statement{},
	}
}

// Parse is the parsing invocation method.
func (p *Parser) Parse() ([]*dslmodel.Statement, error) {
	if len(strings.TrimSpace(p.inputScript)) == 0 {
		return nil, errors.New("There is no input text")
	}
	reader := strings.NewReader(p.inputScript)
	scanner := bufio.NewScanner(reader)
	statements := []*dslmodel.Statement{}
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}
		statement, err := p.parseLine(trimmed)
		if err != nil {
			return nil, umli.DSLError(trimmed, lineNo, err.Error())
		}
		statements = append(statements, statement)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return statements, nil
}

// parseLine parses the text present in a single line of DSL, into
// the fields expected, validates them, and packages the result into a
// dslmodel.Statement.
func (p *Parser) parseLine(line string) (s *dslmodel.Statement, err error) {
	words := strings.Split(line, " ")
	keyWord := words[0]
	if !umli.KnownKeyword(keyWord) {
		return nil, fmt.Errorf("Unrecognized keyword: %s", keyWord)
	}
	requiredNumberOfWords := p.minWordsRequiredFor(keyWord)
	if len(words) < requiredNumberOfWords {
		return nil, fmt.Errorf(
			"A <%s> line, must have at least %d words",
			keyWord, requiredNumberOfWords)
	}
	switch keyWord {
	case umli.Title:
		s, err = p.parseTitle(line, words)
	case umli.TextSize:
		s, err = p.parseTextSize(line, words)
	case umli.Life:
		s, err = p.parseLife(line, words)
	case umli.Full, umli.Dash:
		s, err = p.parseFullOrDash(line, words)
	case umli.Self:
		s, err = p.parseSelf(line, words)
	case umli.Stop:
		s, err = p.parseStop(line, words)
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *Parser) parseTitle(line string, words []string) (
	s *dslmodel.Statement, err error) {
	label := p.removeWords(line, umli.Title)
	return &dslmodel.Statement{
		Keyword:       umli.Title,
		LabelSegments: p.isolateLabelConstituentLines(label),
	}, nil
}

func (p *Parser) parseTextSize(line string, words []string) (
	s *dslmodel.Statement, err error) {
	var textSize float64
	if textSize, err = strconv.ParseFloat(words[1], 64); err != nil {
		return nil, errors.New("Text size must be a number")
	}
	const minTextSize = 5
	const maxTextSize = 20
	if textSize < minTextSize || textSize > maxTextSize {
		return nil, fmt.Errorf("Text size must be between %v and %v",
			minTextSize, maxTextSize)
	}
	return &dslmodel.Statement{
		Keyword:  umli.TextSize,
		TextSize: textSize,
	}, nil
}

func (p *Parser) parseLife(line string, words []string) (
	s *dslmodel.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifelineName := words[1]
	if p.lifelineIsKnown(lifelineName) {
		return nil, fmt.Errorf(
			"Lifeline (%s) has already been used", lifelineName)
	}
	label := p.removeWords(line, umli.Life, lifelineName)
	s = &dslmodel.Statement{
		Keyword:       umli.Life,
		LifelineName:  lifelineName,
		LabelSegments: p.isolateLabelConstituentLines(label),
	}
	p.lifelineStatementsByName[lifelineName] = s
	return s, nil
}

func (p *Parser) parseFullOrDash(line string, words []string) (
	s *dslmodel.Statement, err error) {
	if !twoUCLetters.MatchString(words[1]) {
		return nil, errors.New(
			"Lifelines specified must be two, upper case letters")
	}
	lifelineLetters := strings.Split(words[1], "")
	if lifelineLetters[0] == lifelineLetters[1] {
		return nil, fmt.Errorf(
			"Lifeline letters must be different:(%s)", words[1])
	}
	lifelines := []*dslmodel.Statement{}
	for _, letter := range lifelineLetters {
		lifeline, ok := p.lifelineStatementsByName[letter]
		if !ok {
			return nil, fmt.Errorf("Unknown lifeline: %s", letter)
		}
		lifelines = append(lifelines, lifeline)
	}
	label := p.removeWords(line, words[0], words[1])
	return &dslmodel.Statement{
		Keyword:             words[0],
		ReferencedLifelines: lifelines,
		LabelSegments:       p.isolateLabelConstituentLines(label),
	}, nil
}

func (p *Parser) parseStop(line string, words []string) (
	s *dslmodel.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifeline, ok := p.lifelineStatementsByName[words[1]]
	if !ok {
		return nil, fmt.Errorf("Unknown lifeline: %s", words[1])
	}
	return &dslmodel.Statement{
		Keyword:             umli.Stop,
		ReferencedLifelines: []*dslmodel.Statement{lifeline},
	}, nil
}

func (p *Parser) parseSelf(line string, words []string) (
	s *dslmodel.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifeline, ok := p.lifelineStatementsByName[words[1]]
	if !ok {
		return nil, fmt.Errorf("Unknown lifeline: %s", words[1])
	}
	label := p.removeWords(line, umli.Self, words[1])
	return &dslmodel.Statement{
		Keyword:             umli.Self,
		ReferencedLifelines: []*dslmodel.Statement{lifeline},
		LabelSegments:       p.isolateLabelConstituentLines(label),
	}, nil
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

// minWordsRequiredFor provides the minimum number of words is required
// of a line starting with the given keyword.
func (p *Parser) minWordsRequiredFor(keyWord string) int {
	switch keyWord {
	case umli.Title, umli.TextSize, umli.Stop:
		return 2
	case umli.Life, umli.Full, umli.Dash, umli.Self:
		return 3
	default:
		return 999
	}
}

/*
removeWords returns the given line, from which any instances of the
stringsToRemove have been removed. It also trims whitespace from the ends.
*/
func (p *Parser) removeWords(line string, stringsToRemove ...string) string {
	for _, s := range stringsToRemove {
		line = strings.Replace(line, s, "", 1)
	}
	return strings.TrimSpace(line)
}

// lifelineIsKnown returns true if the parser has previously encountered
// a lifeline statement for the given name.
func (p *Parser) lifelineIsKnown(lifelineName string) bool {
	_, ok := p.lifelineStatementsByName[lifelineName]
	return ok
}

var singleUCLetter = re.MustCompile(`^[A-Z]$`)
var twoUCLetters = re.MustCompile(`^[A-Z][A-Z]$`)
