package parser

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMsgWhenTooFewWords(t *testing.T) {
	assert := assert.New(t)
	p := &Parser{}
	reader := strings.NewReader(`Lane`)
	_, err := p.Parse(bufio.NewScanner(reader))
	assert.NotNil(err)
	assert.EqualError(
		err,
		"Error on this line <Lane> (line: 1): Must have at least 2 words.")
}
func TestItIgnoresBlankLines(t *testing.T) {
	assert := assert.New(t)
	p := &Parser{}
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
	p := &Parser{}
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
	p := &Parser{}
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

func TestSamplingOutputWithReferenceInput(t *testing.T) {
	assert := assert.New(t)
	p := &Parser{}
	reader := strings.NewReader(ReferenceInput)
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
	assert.Len(statements, 11)
	// This has excercised all the keywords.
	aLane := statements[1]
	assert.Equal("B", aLane.LaneName)
}


add more complete checking in parser
	kw is from known set known
	lanes spec is one char for lane and for stop, two for all others
	lanes are uc letters
	lanes addressed got seen already
	lanes have label
	full,dash,self have label

test output for reference input
*/
