package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMsgWhenTooFewWords(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`Lane`)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(
		err,
		"Error on this line <Lane> (line: 1): must have at least 2 words")
}

func TestErrorMsgWhenKeywordIsUnrecognized(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`foo bar`)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(
		err,
		"Error on this line <foo bar> (line: 1): unrecognized keyword: foo")
}

func TestErrorMsgWhenLaneIsNotSingleUCLetterForStopAndLane(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()

	// Few cases to look at details of error message.
	reader := strings.NewReader("lane AB")
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <lane AB> (line: 1): Lane name must be single, upper case letter")
	reader = strings.NewReader("lane a")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.NotNil(err)
	assert.EqualError(err,
		"Error on this line <lane a> (line: 1): Lane name must be single, upper case letter")

	// Make sure it behaves the same way with the only other keywords that
	// requires a single lane to be specified: "stop".
	reader = strings.NewReader("stop a")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lane name must be single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lane to be specified: "stop".
	reader = strings.NewReader("stop a")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lane name must be single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lane to be specified: "self".
	reader = strings.NewReader("self a")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <self a> (line: 1): Lane name must be single, upper case letter")
}

func TestErrorMsgForKeywordsThatExpectTwoLanesDontSpecifyTwoUCLetters(
	t *testing.T) {
	assert := assert.New(t)
	p := NewParser()

	// A few different scenarios

	// Upper case letter but only one of them, <full> keyword
	reader := strings.NewReader("full A")
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <full A> (line: 1): Lanes specified must be two, upper case letters")

	// Two letters but wrong case - dash keyword
	reader = strings.NewReader("dash ab")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <dash ab> (line: 1): Lanes specified must be two, upper case letters")

	// Two characters but one is not a letter - dash keyword
	reader = strings.NewReader("dash A3")
	_, err = p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <dash A3> (line: 1): Lanes specified must be two, upper case letters")
}

func TestItIgnoresBlankLines(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`
		lane A  SL App

		lane B  Core Permissions API
	`)
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
	assert.Len(statements, 2)
}

func TestItCapturesLabelTextWithNoLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`
		lane A  SL App
	`)
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
	assert.Len(statements[0].LabelSegments, 1)
	assert.Equal("SL App", statements[0].LabelSegments[0], 1)
}

func TestItCapturesLabelTextWithLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`
		lane A  The quick | brown fox | etc
	`)
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
	assert.Len(statements[0].LabelSegments, 3)
	// Note we check not only the splitting but also that each
	// segment is trimmed of whitespace.
	assert.Equal("The quick", statements[0].LabelSegments[0])
	assert.Equal("brown fox", statements[0].LabelSegments[1])
	assert.Equal("etc", statements[0].LabelSegments[2])
}

func TestErrorMessageWhenAnUnknownLaneIsReferenced(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`full AB  foo`)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <full AB  foo> (line: 1): Unknown lane: A")
}

func TestErrorMessageWhenAStatementOmitsALabel(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(`
		lane A
	`)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.EqualError(err,
		"Error on this line <lane A> (line: 2): Label text missing")
}

func TestMakeSureEveryKeywordIsHandledWithoutError(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(ReferenceInput)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
}

func TestMakeSureARepresentativeStatementOutputIsProperlyFormed(t *testing.T) {
	assert := assert.New(t)
	p := NewParser()
	reader := strings.NewReader(ReferenceInput)
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)

	// full CB  get_user_permissions( | token)
	s := statements[4]
	assert.Equal("full", s.Keyword)
	assert.Equal("C", s.ReferencedLanes[0].LaneName)
	assert.Equal("B", s.ReferencedLanes[1].LaneName)
	assert.Equal("get_user_permissions(", s.LabelSegments[0])
	assert.Equal("token)", s.LabelSegments[1])
}
