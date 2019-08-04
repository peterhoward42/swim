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

type Parser struct {
    inputScript string
    lifelineStatementsByName = map[string]*dslmodel.Statement
}

func NewParser(inputScript string) *Parser {
    return &Parser{
       inputScript: inputScript,
       lifelineStatementsByName: map[string]*dslmodel.Statement{} 
    }
}


// Parse is the parsing invocation method.
func (p *Parser) Parse() ([]*dslmodel.Statement, error) {
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
		statement, err := p.parseLine(trimmed, knownLifelines)
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

var singleUCLetter = re.MustCompile(`^[A-Z]$`)
var twoUCLetters = re.MustCompile(`^[A-Z][A-Z]$`)

// parseLine parses the text present in a single line of DSL, into
// the fields expected, validates them, and packages the result into a
// dslmodel.Statement.
func (p *Parser) parseLine(line string) (s *dslmodel.Statement, err error) {
	words := strings.Split(line, " ")
	keyWord := words[0]
    if !p.keyWordIsRecognized(keyWord) {
		return nil, fmt.Errorf("Unrecognized keyword: %s", keyWord)
    }
    requiredNumberOfWords := p.minWordsRequiredFor(keyWord)
    if len(words) < requiredNumberOfWords {
		return nil, fmt.Errorf(
            "Too few words in <%s> line, must be at least %d",
            requiredNumberOfWords)
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
	default:
		return nil, fmt.Errorf("unrecognized keyword: %s", keyWord)
	}
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (p *Parser) parseTitle(line string, words []string) (
    s *dslmodel.Statement, err error) {
    label := p.removeLeadingWords(line, words, 1)
    return &dslmodel.Statement{
        Keyword: umli.Title,
        LabelSegments: p.isolateLabelConstituentLines(label)
    }, nil
}

func (p *Parser) parseTextSize(line string, words []string) (
    s *dslmodel.Statement, err error) {
	if textSize, err = strconv.ParseFloat(words[1], 64); err != nil {
		return nil, errors.New("Text size must be a number")
	}
	const minTextSize = 5
	const maxTextSize = 20
	if textSize < minTextSize || textSize > maxTextSize {
		return nil, fmt.Errorf("textsize must be between %v and %v",
			minTextSize, maxTextSize)
	}
    return &dslmodel.Statement{
        Keyword: umli.TextSize,
        TextSize: textSize
    }, nil
}

func (p *Parser) parseLife(line string, words []string) (
    s *dslmodel.Statement, err error) {
    if !p.isUpperCaseLetter(words[1]) {
        return nil, fmt.Errorf(
            "Lifeline name (%s) must be a single, upper case letter", words[1])
    }
    lifelineName := words[1]
    if p.lifelineIsKnown(lifelineName) {
        return nil, fmt.Errorf(
            "Lifeline (%s) has already been used", lifelineName)
    }
    label := removeLeadingWords(line, words, 2)
    s := &dslmodel.Statement{
        Keyword: umli.Life,
        LifelineName: lifelineLetter,
        LabelSegments: p.isolateLabelConstituentLines(label)
    }
    p.lifelineStatementsByName[lifelineName] = s
    return s, nil
}

func (p *Parser) parseFullOrDash(line string, words []string) (
    s *dslmodel.Statement, err error) {
    if !p.isPairOfUpperCaseLetter(words[1]) {
        return nil, fmt.Errorf(
            "Need two lifeline letters, instead of <%s>", words[1])
    }
    lifelineLetters := strings.Split(words[1], "")
    if lifelineLetters[0] == lifelineLetters[1] {
        return nil, fmt.Errorf(
            "Lifeline letters must be different:(%s)",words[1])
    }
    lifelines := []*dslmodel.Statement{}
    for _, letter := range lifelineLetters {
        lifeline := p.lifelineStatementFor(letter)
        if lifeline == nil {
            return nil, fmt.Errorf("Unknown lifeline: %v", letter)
        }
        lifelines = append(lifelines, lifeline)
    }
    label := p.removeLeadingWords(line, words, 2)
    return &dslmodel.Statement{
        Keyword: words[0],
        ReferencedLifelines: lifelines,
        LabelSegments: p.isolateLabelConstituentLines(label)
    }, nil
}

func (p *Parser) parseStop(line string, words []string) (
    s *dslmodel.Statement, err error) {
    if !p.isUpperCaseLetter(words[1]) {
        return nil, fmt.Errorf(
            "Need one lifeline letter, instead of <%s>", words[1])
    }
    lifeline := p.lifelineStatementFor(words[1])
    if lifeline == nil {
        return nil, fmt.Errorf("Unknown lifeline: %v", letter)
    }
    return &dslmodel.Statement{
        Keyword: umli.Stop,,
        ReferencedLifelines: []*dslmodel.Statement{lifeline},
    }, nil
}

func (p *Parser) parseSelf(line string, words []string) (
    s *dslmodel.Statement, err error) {
    if !p.isUpperCaseLetter(words[1]) {
        return nil, fmt.Errorf(
            "Need one lifeline letter, instead of <%s>", words[1])
    }
    lifeline := p.lifelineStatementFor(words[1])
    if lifeline == nil {
        return nil, fmt.Errorf("Unknown lifeline: %v", letter)
    }
    label := p.removeLeadingWords(line, words, 2)
    return &dslmodel.Statement{
        Keyword: umli.Self,,
        ReferencedLifelines: []*dslmodel.Statement{lifeline},
        LabelSegments: p.isolateLabelConstituentLines(label)
    }, nil
}

// isolateLabelConstituentLines takes the label text from a DSL line and
// splits it into the constituent lines according to its author's intent.
// I.e. by splitting it at "|" delimiters. Note the removal of whitespace
// either side of any "|" present.
// E.g. From this: "edit_facilities( | payload, user_token)"
// It produces: []string{"edit_facilities(", "payload, user_token)"}
func isolateLabelConstituentLines(labelText string) []string {
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

/*
parseLifelinesOperand makes sure the lifelines that are specified in the second
word of a DSL line are properly formed. This depends on the keyword.  It also
maintains a look up table of lifeline name to corresponding Lifeline statement
in the parser.
*/
func parseLifelinesOperand(lifelineNamesOperand, keyWord string,
	knownLifelines lifelineStatementsByName) (
	[]*dslmodel.Statement, error) {

	// Fail fast on statement types that require a single lifline to be
	// specified, when this is not so.
	if keyWord == umli.Life || keyWord == umli.Stop || keyWord == umli.Self {
		if !singleUCLetter.MatchString(lifelineNamesOperand) {
			return nil,
				errors.New("Lifeline name must be single, upper case letter")
		}
	}
	// Same sort of thing where two lifelines must be specified.
	if keyWord == umli.Full || keyWord == umli.Dash {
		if !twoUCLetters.MatchString(lifelineNamesOperand) {
			return nil,
				errors.New(
					"Lifelines specified must be two, upper case letters")
		}
	}
	// Capture ptrs to the lifeline Statement being referenced by the
	// second word. (Unless this IS a lifeline statement).
	lifelinestatements := []*dslmodel.Statement{}
	if keyWord != umli.Life {
		lifelineLetters := strings.Split(lifelineNamesOperand, "")
		for _, lifelineLetter := range lifelineLetters {
			lifelinestatement, ok := knownLifelines[lifelineLetter]
			if !ok {
				return nil, fmt.Errorf("Unknown lifeline: %v", lifelineLetter)
			}
			lifelinestatements = append(lifelinestatements, lifelinestatement)
		}
	}
	return lifelinestatements, nil
}

// parseTextSize converts the textSizeOperand string into a float64
// within acceptable bounds if possible.
func parseTextSize(textSizeOperand string) (textSize float64, err error) {
	if textSize, err = strconv.ParseFloat(textSizeOperand, 64); err != nil {
		return -1, errors.New("Text size must be a number")
	}
	const minTextSize = 5
	const maxTextSize = 20
	if textSize < minTextSize || textSize > maxTextSize {
		return -1, fmt.Errorf("textsize must be between %v and %v",
			minTextSize, maxTextSize)
	}
	return textSize, nil
}
