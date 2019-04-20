package dslmodel

import (
	"fmt"

	umli "github.com/peterhoward42/umlinteraction"
	"github.com/peterhoward42/umlinteraction/parser"
)

// Model encapsulates a programmatic, higher-level model of a DSL script.
// It aims to provide something more convenient for the diagram building
// stages to consume as input than the text form, and to decouple the diagram
// building system from the DSL Parser.
type Model struct {
	// Every type of input line can be represented in a Statement object.
	// The Statements attribute here is to model their *sequence*.
	Statements []*Statement
	// This map provides a lookup for the Statements that define lanes. Keyed on
	// the lane letter.
	LaneLookup map[string]*Statement
}

// NewModel constructs a usable Model.
func NewModel() *Model {
	return &Model{
		Statements: []*Statement{},
		LaneLookup: map[string]*Statement{},
	}
}

// Build populates a Model by interpreting the list of ParsedLine structures
// provided.
func (m *Model) Build(inputLines []*parser.ParsedLine) error {
	for _, parsedLine := range inputLines {
		switch parsedLine.KeyWord {
		case umli.Lane:
			m.lane(parsedLine)
		case umli.Full:
			m.interaction(parsedLine)
		case umli.Dash:
			m.interaction(parsedLine)
		case umli.Self:
			m.self(parsedLine)
		case umli.Stop:
			m.stop(parsedLine)
		default:
			return fmt.Errorf("Build(): Unknown keyword <%v>", parsedLine.KeyWord)
		}
	}
	return nil
}

func (m *Model) lane(line *parser.ParsedLine) error {
	statement := &Statement{}
	statement.Keyword = line.KeyWord
	statement.LaneName = line.Lanes[0]
	statement.LabelSegments = line.LabelSegments

	m.Statements = append(m.Statements, statement)
	m.LaneLookup[statement.LaneName] = statement
	return nil
}

// interactions deals both with <full> and <dash> statements.
func (m *Model) interaction(line *parser.ParsedLine) error {
	statement := &Statement{}
	statement.Keyword = line.KeyWord
	laneStatements, err := m.findLaneStatements(line.Lanes)
	if err != nil {
		return fmt.Errorf("This line: <%v> refers to an unknown lane", line.FullText)
	}
	if len(laneStatements) != 2 {
		return fmt.Errorf("this line: <%v> should have two lanes specified", line.FullText)
	}
	statement.ReferencedLanes = laneStatements
	statement.LabelSegments = line.LabelSegments

	m.Statements = append(m.Statements, statement)
	return nil
}

func (m *Model) self(line *parser.ParsedLine) error {
	statement := &Statement{}
	statement.Keyword = line.KeyWord
	laneStatements, err := m.findLaneStatements(line.Lanes)
	if err != nil {
		return fmt.Errorf("this line: <%v> refers to an unknown lane", line.FullText)
	}
	if len(laneStatements) != 1 {
		return fmt.Errorf("this line: <%v> should have only one lane specified", line.FullText)
	}
	statement.ReferencedLanes = laneStatements
	statement.LabelSegments = line.LabelSegments

	m.Statements = append(m.Statements, statement)
	return nil
}

func (m *Model) stop(line *parser.ParsedLine) error {
	statement := &Statement{}
	statement.Keyword = line.KeyWord
	laneStatements, err := m.findLaneStatements(line.Lanes)
	if err != nil {
		return fmt.Errorf("This line: <%v> refers to an unknown lane", line.FullText)
	}
	if len(laneStatements) != 1 {
		return fmt.Errorf("this line: <%v> should have only one lane specified", line.FullText)
	}
	statement.ReferencedLanes = laneStatements
	statement.LabelSegments = line.LabelSegments

	m.Statements = append(m.Statements, statement)
	return nil
}

// findLaneStatements searches for previously captured *Lane* Statements.
func (m *Model) findLaneStatements(laneLetters []string) ([]*Statement, error) {
	statements := []*Statement{}
	for _, letter := range laneLetters {
		statement, ok := m.LaneLookup[letter]
		if ok == false {
			return nil, fmt.Errorf("Line refers to unknown lane: <%v>", letter)
		}
		statements = append(statements, statement)
	}
	return statements, nil
}
