package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMsgWhenTooFewWords(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse("life")
	assert.EqualError(
		err,
		"Error on this line <life> (line: 1): must have at least 2 words")
}

func TestErrorMsgWhenKeywordIsUnrecognized(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse("foo bar")
	assert.EqualError(
		err,
		"Error on this line <foo bar> (line: 1): unrecognized keyword: foo")
}

func TestErrorMsgWhenLifelineIsNotSingleUCLetterForStopAndLifeline(t *testing.T) {
	assert := assert.New(t)

	// Few cases to look at details of error message.
	_, err := Parse("life AB")
	assert.EqualError(err,
		"Error on this line <life AB> (line: 1): Lifeline name "+
			"must be single, upper case letter")
	_, err = Parse("life a")
	assert.NotNil(err)
	assert.EqualError(err,
		"Error on this line <life a> (line: 1): Lifeline name "+
			"must be single, upper case letter")

	// Make sure it behaves the same way with the only other keywords that
	// requires a single lifeline to be specified: "stop".
	_, err = Parse("stop a")
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lifeline name must "+
			"be single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lifeline to be specified: "stop".
	_, err = Parse("stop a")
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lifeline name must "+
			"be single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lifeline to be specified: "self".
	_, err = Parse("self a")
	assert.EqualError(err,
		"Error on this line <self a> (line: 1): Lifeline name must "+
			"be single, upper case letter")
}

func TestErrorMsgForKeywordsThatExpectTwoLifelinesDontSpecifyTwoUCLetters(
	t *testing.T) {
	assert := assert.New(t)

	// A few different scenarios

	// Upper case letter but only one of them, <full> keyword
	_, err := Parse("full A")
	assert.EqualError(err,
		"Error on this line <full A> (line: 1): Lifelines specified must "+
			"be two, upper case letters")

	// Two letters but wrong case - dash keyword
	_, err = Parse("dash ab")
	assert.EqualError(err,
		"Error on this line <dash ab> (line: 1): Lifelines specified must "+
			"be two, upper case letters")

	// Two characters but one is not a letter - dash keyword
	_, err = Parse("dash A3")
	assert.EqualError(err,
		"Error on this line <dash A3> (line: 1): Lifelines specified "+
			"must be two, upper case letters")
}

func TestItIgnoresBlankLines(t *testing.T) {
	assert := assert.New(t)
	statements, err := Parse(`
		life A  SL App

		life B  Core Permissions API
    `)
	assert.Nil(err)
	assert.Len(statements, 2)
}

func TestItCapturesLabelTextWithNoLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	statements, err := Parse("life A SL App")
	assert.Nil(err)
	assert.Len(statements[0].LabelSegments, 1)
	assert.Equal("SL App", statements[0].LabelSegments[0], 1)
}

func TestItCapturesLabelTextWithLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	statements, err := Parse("life A  The quick | brown fox | etc")
	assert.Nil(err)
	assert.Len(statements[0].LabelSegments, 3)
	// Note we check not only the splitting but also that each
	// segment is trimmed of whitespace.
	assert.Equal("The quick", statements[0].LabelSegments[0])
	assert.Equal("brown fox", statements[0].LabelSegments[1])
	assert.Equal("etc", statements[0].LabelSegments[2])
}

func TestErrorMessageWhenAnUnknownLifelineIsReferenced(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse("full AB foo")
	assert.EqualError(err,
		"Error on this line <full AB foo> (line: 1): Unknown lifeline: A")
}

func TestErrorMessageWhenAStatementOmitsALabel(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse("life A")
	assert.EqualError(err,
		"Error on this line <life A> (line: 1): Label text missing")
}

func TestMakeSureEveryKeywordIsHandledWithoutError(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse(ReferenceInput)
	assert.Nil(err)
}

func TestMakeSureARepresentativeStatementOutputIsProperlyFormed(t *testing.T) {
	assert := assert.New(t)
	statements, err := Parse(ReferenceInput)
	assert.Nil(err)

	// full CB  get_user_permissions( | token)
	s := statements[5]
	assert.Equal("full", s.Keyword)
	assert.Equal("C", s.ReferencedLifelines[0].LifelineName)
	assert.Equal("B", s.ReferencedLifelines[1].LifelineName)
	assert.Equal("get_user_permissions(", s.LabelSegments[0])
	assert.Equal("token)", s.LabelSegments[1])
}

func TestNoErrorsWhenOptionalStatementsAreOmitted(t *testing.T) {
	assert := assert.New(t)
	// This DSL excludes *title* and *textsize* statements.
	// Both of which are optional.
	_, err := Parse(`
        life A foo
        self A bar
    `)
	assert.Nil(err)
}

func TestErrorsForMalformedTextSize(t *testing.T) {
	assert := assert.New(t)
	_, err := Parse("textsize garbage")
	assert.EqualError(err,
		"Error on this line <textsize garbage> (line: 1): Text size must be a number")
	_, err = Parse("textsize 21")
	assert.EqualError(err,
		"Error on this line <textsize 21> (line: 1): textsize must be between 5 and 20")
	_, err = Parse("textsize 4")
	assert.EqualError(err,
		"Error on this line <textsize 4> (line: 1): textsize must be between 5 and 20")
}

func TestWellFormedTextSizeIsParsedCorrectly(t *testing.T) {
	assert := assert.New(t)
	statements, err := Parse("textsize 10")
	assert.NoError(err)
	s := statements[0]
	expected := 10.0
	assert.Equal(expected, s.TextSize)
}
