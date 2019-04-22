/*
Package dslmodel encapsulates a domain-specic-language (DSL) script.
*/
package dslmodel

import (
	"bufio"
	umli "github.com/peterhoward42/umlinteraction"
	"strings"
	"testing"

	"github.com/peterhoward42/umlinteraction/parser"
)

func TestBuildWithCorrectInput(t *testing.T) {
	// Use the Parser as shortcut to get hold of list of ParsedLines.
	p := &parser.Parser{}
	reader := strings.NewReader(parser.ReferenceInput)
	parsedLines, err := p.Parse(bufio.NewScanner(reader))
	if err != nil {
		t.Errorf("Parse(): %v", err)
	}
	// Now we can exercise the DSL model builder
	model := NewModel()
	err = model.Build(parsedLines)
	if err != nil {
		t.Errorf("Failed to Build() with error: %v", err)
	}

	// A few sanity checks
	if len(model.Statements) != 11 {
		t.Errorf("Wrong number Statements built")
	}

	// Sniff around the statement corresponding to:
	// 		dash CA  status_ok, payload
	stmt := model.Statements[8]
	if stmt.Keyword != umli.Dash {
		t.Errorf("Wrong keyword")
	}
	if stmt.ReferencedLanes[0].LaneName != "C" {
		t.Errorf("Wrong lane referenced")
	}
	if stmt.ReferencedLanes[1].LaneName != "A" {
		t.Errorf("Wrong lane referenced")
	}
}
