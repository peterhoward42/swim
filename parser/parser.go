// Package parser provides the Parse method, which is capable of parsing lines
// of DSL text to provide a structured representation of it - in the form
// of a slice of dslmodel.Statement(s).
package parser

import (
	"bufio"
	"errors"
	"fmt"
	re "regexp"
	"strings"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
)

type lifelineStatementsByName = map[string]*dslmodel.Statement

// Parse is the parsing invocation method.
func Parse(DSLScript string) ([]*dslmodel.Statement, error) {
	reader := strings.NewReader(DSLScript)
	scanner := bufio.NewScanner(reader)
	statements := []*dslmodel.Statement{}
	knownLifelines := lifelineStatementsByName{}
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}
		statement, err := parseLine(trimmed, knownLifelines)
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
func parseLine(line string, knownLifelines lifelineStatementsByName) (
	*dslmodel.Statement, error) {
	// Fail fast when < 2 words.
	words := strings.Split(line, " ")
	if len(words) < 2 {
		return nil, errors.New("must have at least 2 words")
	}
	// Fail fast on unrecognized keyword.
	keyWord := words[0]
	if !strings.Contains(strings.Join(umli.AllKeywords, " "), keyWord) {
		return nil, fmt.Errorf("unrecognized keyword: %s", keyWord)
	}
	// Validate and reconcile the lifelines referenced in the second word.
	lifelineNamesOperand := words[1]
	lifelinesReferenced, err := parseLifelinesOperand(
		lifelineNamesOperand, keyWord, knownLifelines)
	if err != nil {
		return nil, err
	}

	// Isolate label text by stripping what we have already consumed.
	labelText := strings.Replace(line, keyWord, "", 1)
	labelText = strings.Replace(labelText, words[1], "", 1)
    // Interpret pipes (|) as line breaks.
	labelIndividualLines := isolateLabelConstituentLines(labelText)

	// Make sure labels are present on statement for which they are
	// mandatory
	if keyWord != umli.Stop && len(labelIndividualLines) == 0 {
		return nil, errors.New("Label text missing")
	}

	// Construct the Statement to return
	statement := dslmodel.NewStatement()
	statement.Keyword = keyWord
	statement.LabelSegments = labelIndividualLines
	statement.ReferencedLifelines = lifelinesReferenced

	// A few extra steps for *Life* statements
	if statement.Keyword == umli.Life {
		statement.LifelineName = lifelineNamesOperand
		knownLifelines[statement.LifelineName] = statement
	}
	return statement, nil
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
