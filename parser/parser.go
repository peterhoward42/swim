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

type laneStatementsByName = map[string]*dslmodel.Statement

// Parse is the parsing invocation method.
func Parse(DSLScript string) ([]*dslmodel.Statement, error) {
	reader := strings.NewReader(DSLScript)
	scanner := bufio.NewScanner(reader)
	statements := []*dslmodel.Statement{}
	knownLanes := laneStatementsByName{}
	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}
		statement, err := parseLine(trimmed, knownLanes)
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
func parseLine(line string, knownLanes laneStatementsByName) (
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
	// Validate and reconcile the lanes referenced in the second word.
	laneNamesOperand := words[1]
	lanesReferenced, err := parseLanesOperand(
		laneNamesOperand, keyWord, knownLanes)
	if err != nil {
		return nil, err
	}

	// Isolate label text by stripping what we have already consumed.
	labelText := strings.Replace(line, keyWord, "", 1)
	labelText = strings.Replace(labelText, words[1], "", 1)
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
	statement.ReferencedLanes = lanesReferenced

	// A few extra steps for *Lane* statements
	if statement.Keyword == umli.Lane {
		statement.LaneName = laneNamesOperand
		knownLanes[statement.LaneName] = statement
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

// parseLanesOperand makes sure the lanes that are specified in the second
// word of a DSL line are properly formed. This depends on the keyword.
// It also maintains a look up table of lane name to corresponding Lane
// statement in the parser.
func parseLanesOperand(
	laneNamesOperand, keyWord string, knownLanes laneStatementsByName) (
	[]*dslmodel.Statement, error) {

	// Fail fast on statement types that require a single lane to be specified,
	// when this is not so.
	if keyWord == umli.Lane || keyWord == umli.Stop || keyWord == umli.Self {
		if !singleUCLetter.MatchString(laneNamesOperand) {
			return nil,
				errors.New("Lane name must be single, upper case letter")
		}
	}
	// Same sort of thing where two lanes must be specified.
	if keyWord == umli.Full || keyWord == umli.Dash {
		if !twoUCLetters.MatchString(laneNamesOperand) {
			return nil,
				errors.New("Lanes specified must be two, upper case letters")
		}
	}
	// Capture ptrs to the lane Statement being referenced by the second word.
	// (Unless this IS a lane statement).
	laneStatements := []*dslmodel.Statement{}
	if keyWord != umli.Lane {
		laneLetters := strings.Split(laneNamesOperand, "")
		for _, laneLetter := range laneLetters {
			laneStatement, ok := knownLanes[laneLetter]
			if !ok {
				return nil, fmt.Errorf("Unknown lane: %v", laneLetter)
			}
			laneStatements = append(laneStatements, laneStatement)
		}
	}
	return laneStatements, nil
}
