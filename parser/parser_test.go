package parser

import (
	"bufio"
	"strings"
	"testing"
)

var referenceInput = `
lane A  SL App
lane B  Core Permissions API
lane C  SL Admin API | edit_facilities | endpoint

full AC  edit_facilities( | payload, user_token)
full CB  get_user_permissions( | token)
dash BC  permissions_list
stop B
self C   [has EDIT_FACILITIES permission] | store changes etc
dash CA  status_ok, payload
self C   [no permission]
dash CA  status_not_authorized
`

func TestWithCorrectInput(t *testing.T) {
	p := &Parser{}
	reader := strings.NewReader(referenceInput)
	parsedLines, err := p.Parse(bufio.NewScanner(reader))
	if err != nil {
		t.Errorf("Parse(): %v", err)
	}
	if len(parsedLines) != 11 {
		t.Errorf("Parse() returns %v lines, should be 11", len(parsedLines))
	}

	// Sample this line
	//      full CB  get_user_permissions( | token)
	// Because it has pipe characters in the text, and two
	// lanes cited in the lanes operand.

	pl := parsedLines[4]
	if kw := pl.KeyWord; kw != "full" {
		t.Errorf("Wrong keyword %v, should be <full>", kw)
	}
	if lane := pl.Lanes[0]; lane != "C" {
		t.Errorf("Wrong first lane %v, should be <C>", lane)
	}
	if lane := pl.Lanes[1]; lane != "B" {
		t.Errorf("Wrong second lane %v, should be <B>", lane)
	}
	if label := pl.LabelSegments[0]; label != "get_user_permissions(" {
		t.Errorf("Wrong label %v, should be <get_user_permissions(>", label)
	}
	if label := pl.LabelSegments[1]; label != "token)" {
		t.Errorf("Wrong label %v, should be <token)", label)
	}
}

func TestProducesWellFormedError(t *testing.T) {
	p := &Parser{}
	reader := strings.NewReader("garbage input")
	_, err := p.Parse(bufio.NewScanner(reader))
	if err == nil {
		t.Errorf("Should have produced error")
	}
	expected := "parseLine() input line malformed: garbage input"
	if msg := err.Error(); msg != expected {
		t.Errorf("Wrong error message %v, should be <%v>)", msg, expected)
	}
}
