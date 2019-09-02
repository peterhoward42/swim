// Package parser provides the Parse function, which is capable of parsing lines
// of DSL text to provide a structured representation of it - in the form
// of a dsl.Model.
package parser

import (
	"bufio"
	"errors"
	"fmt"
	re "regexp"
	"strconv"
	"strings"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dsl"
)

// Parser is capable of parsing the DSL script to produce a dsl.Model.
type Parser struct {
    inputScript string
	model dsl.Model
}

func NewParser(inputScript string) *Parser {
    return &Parser{
        inputScript: inputScript,
        }
}

// Parse is the parsing invocation method.
func (p *Parser) Parse() (*dsl.Model, error) {
	if len(strings.TrimSpace(p.inputScript)) == 0 {
		return nil, errors.New("There is no input text")
	}
	reader := strings.NewReader(p.inputScript)
	scanner := bufio.NewScanner(reader)
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
		p.model.Append(statement)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &p.model, nil
}

// parseLine parses the text present in a single line of DSL, into
// the fields expected, validates them, and packages the result into a
// dsl.Statement.
func (p *Parser) parseLine(line string) (s *dsl.Statement, err error) {
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
	case umli.ShowLetters:
		s, err = p.parseShowLetters(line, words)
	case umli.Life:
		s, err = p.parseLife(line, words)
	case umli.Full, umli.Dash:
		s, err = p.parseFullOrDash(line, words)
	case umli.Self:
		s, err = p.parseSelf(line, words)
	case umli.Stop:
		s, err = p.parseStop(line, words)
	default:
		panic(fmt.Sprintf(
			"Developer has registered keyword <%s> but forgotten to call handler",
			keyWord))
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *Parser) parseTitle(line string, words []string) (
	s *dsl.Statement, err error) {
	label := p.removeWords(line, umli.Title)
	return &dsl.Statement{
		Keyword:       umli.Title,
		LabelSegments: p.isolateLabelConstituentLines(label),
	}, nil
}

func (p *Parser) parseTextSize(line string, words []string) (
	s *dsl.Statement, err error) {
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
	return &dsl.Statement{
		Keyword:  umli.TextSize,
		TextSize: textSize,
	}, nil
}

func (p *Parser) parseShowLetters(line string, words []string) (
	s *dsl.Statement, err error) {
	var show bool
	switch words[1] {
	case "true":
		show = true
	case "false":
		show = false
	default:
		return nil, errors.New("showletters expects <true> or <false>")
	}
	return &dsl.Statement{
		Keyword:     umli.ShowLetters,
		ShowLetters: show,
	}, nil
}

func (p *Parser) parseLife(line string, words []string) (
	s *dsl.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifelineName := words[1]
	if p.model.LifelineIsKnown(lifelineName) {
		return nil, fmt.Errorf(
			"Lifeline (%s) has already been used", lifelineName)
	}
	label := p.removeWords(line, umli.Life, lifelineName)
	s = &dsl.Statement{
		Keyword:       umli.Life,
		LifelineName:  lifelineName,
		LabelSegments: p.isolateLabelConstituentLines(label),
	}
	return s, nil
}

func (p *Parser) parseFullOrDash(line string, words []string) (
	s *dsl.Statement, err error) {
	if !twoUCLetters.MatchString(words[1]) {
		return nil, errors.New(
			"Lifelines specified must be two, upper case letters")
	}
	lifelineLetters := strings.Split(words[1], "")
	if lifelineLetters[0] == lifelineLetters[1] {
		return nil, fmt.Errorf(
			"Lifeline letters must be different:(%s)", words[1])
	}
	lifelines := []*dsl.Statement{}
	for _, letter := range lifelineLetters {
		lifeline, ok := p.model.LifelineStatementByName(letter)
		if !ok {
			return nil, fmt.Errorf("Unknown lifeline: %s", letter)
		}
		lifelines = append(lifelines, lifeline)
	}
	label := p.removeWords(line, words[0], words[1])
	return &dsl.Statement{
		Keyword:             words[0],
		ReferencedLifelines: lifelines,
		LabelSegments:       p.isolateLabelConstituentLines(label),
	}, nil
}

func (p *Parser) parseStop(line string, words []string) (
	s *dsl.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifeline, ok := p.model.LifelineStatementByName(words[1])
	if !ok {
		return nil, fmt.Errorf("Unknown lifeline: %s", words[1])
	}
	return &dsl.Statement{
		Keyword:             umli.Stop,
		ReferencedLifelines: []*dsl.Statement{lifeline},
	}, nil
}

func (p *Parser) parseSelf(line string, words []string) (
	s *dsl.Statement, err error) {
	if !singleUCLetter.MatchString(words[1]) {
		return nil, fmt.Errorf(
			"Lifeline name (%s) must be a single, upper case letter", words[1])
	}
	lifeline, ok := p.model.LifelineStatementByName(words[1])
	if !ok {
		return nil, fmt.Errorf("Unknown lifeline: %s", words[1])
	}
	label := p.removeWords(line, umli.Self, words[1])
	return &dsl.Statement{
		Keyword:             umli.Self,
		ReferencedLifelines: []*dsl.Statement{lifeline},
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
	case umli.Title, umli.TextSize, umli.ShowLetters, umli.Stop:
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

var singleUCLetter = re.MustCompile(`^[A-Z]$`)
var twoUCLetters = re.MustCompile(`^[A-Z][A-Z]$`)
