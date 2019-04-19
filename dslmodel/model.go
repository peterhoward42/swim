package dslmodel

import (
	"fmt"
    "github.com/peterhoward42/umlinteraction/parser"
)

// Model encapsulates a programmatic model of some text written in the
// DSL. It aims to provide something more convenient for the diagram building
// stages to consume as input than the text form, and to decouple the diagram 
// builder from the Parser.
type Model struct {
	// Every type of input line can be represented in a Statement object.
	// The Statements attribute here is what models their sequence.
	Statements []Statement
}

// NewModel constructs a usable Model.
func NewModel() *Model {
	return &Model{
		Statements: []Statement{},
	}
}

// Build populates a Model by interpreting the list of ParsedLine structures
// provided.
func (m *Model) Build(inputLines [] *parser.ParsedLine) error {
    for _, parsedLine := range(inputLines) {
        switch parsedLine.KeyWord {
		case "Lane":
			m.lane(parsedLine)
		case "full":
			m.interaction(parsedLine)
		case "dash":
			m.interaction(parsedLine)
		case "self":
			m.interaction(parsedLine)
		case "stop":
			m.stop(parsedLine)
         default:
			return fmt.Errorf("Build(): Unknown keyword <%v>", parsedLine.KeyWord)
		 }
    }
    return nil
}

func (m *Model) lane(line *parser.ParsedLine) error {
	statement := Statement{}
	statement.Keyword = line.KeyWord
	statement.LaneName = line.Lanes[0]
	statement.LabelLines = line.LabelSegments

	m.Statements = append(m.Statements, statement)
}
func (m *Model) interaction(line *parser.ParsedLine) error {
	statement := Statement{}
	statement.Keyword = line.KeyWord
	fromStatement, toStatement := m.findLaneStatements(line.Lanes)
	if fromStatement == nil {
		return fmt.Errorf("Line refers to unknown lanes: <%v>", line.Lanes)
	}
	statement.FirstLane = fromStatement
	statement.SecondLane = toStatement=
	statement.LabelLines = line.LabelSegments

	m.Statements = append(m.Statements, statement)
}

func (m *Model) findLaneStatements(laneLetters []string) (from, to Statement) {
	fromLetter := laneLetters[0]
	fromLaneStatement, ok := m.LaneStatements[fromLetter]
	if ok == false {
		return nil, nil
	}
	toLetter := laneLetters[1]
	toLaneStatement, ok := m.LaneStatements[toLetter]
	if ok == false {
		return nil, nil
	}
	return fromLaneStatement, toLaneStatement
}
