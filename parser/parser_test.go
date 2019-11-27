package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Start by testing the lower-level parsing helper functions

func TestParseAddsLineNumberContextToErrors(t *testing.T) {
	// Make sure that the Parse method augments errors returned to it
	// with the line concerned and its line number.
	assert := assert.New(t)
	_, err := NewParser(`title Some Title
		life A foo
		life B bar
		nonsense line
	`).Parse()
	msg := err.Error()
	expectedBeginning := "Error on this line <nonsense line> (line: 4):"
	assert.Equal(expectedBeginning, msg[:len(expectedBeginning)])
}

func TestLabelSplittingFunction(t *testing.T) {
	// Make sure that that the utility function that splits labels into
	// individual line segments works properly.
	assert := assert.New(t)

	// Multiline delimiters present
	segments := (&Parser{}).isolateLabelConstituentLines("abc def | ghi | jkl")
	assert.Len(segments, 3)
	assert.Equal("abc def", segments[0])
	assert.Equal("ghi", segments[1])
	assert.Equal("jkl", segments[2])

	// Multiline delimiters absent
	segments = (&Parser{}).isolateLabelConstituentLines(" abc ")
	assert.Len(segments, 1)
	assert.Equal("abc", segments[0])
}

func TestStrippingOutWords(t *testing.T) {
	// Make sure that that the utility function that removes words from
	// the line works properly.
	assert := assert.New(t)

	// Remove 3 words that are present
	stripped := (&Parser{}).removeStrings(
		"the quick lazy brown fox",
		"quick", "lazy", "fox")
	assert.Equal("the   brown", stripped)

	// Make sure it works when one of the words to remove is not present
	stripped = (&Parser{}).removeStrings(
		"the quick lazy brown fox",
		"quick", "QQQQQQ", "fox")
	assert.Equal("the  lazy brown", stripped)
}

func TestErrorMsgWhenNoInput(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser(`
        

    `).Parse()
	assert.EqualError(err, "There is no input text")
}

func TestErrorMsgWhenTooFewWords(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser("life").Parse()
	assert.EqualError(
		err,
		"Error on this line <life> (line: 1): A <life> line, must have at least 3 words")
}

func TestErrorMsgWhenKeywordIsUnrecognized(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser("foo bar").Parse()
	assert.EqualError(
		err,
		"Error on this line <foo bar> (line: 1): Unrecognized keyword: foo")
}

func TestErrorWhenSingleLetterLifelineExpected(t *testing.T) {
	assert := assert.New(t)

	// Few cases to look at details of error message.
	_, err := NewParser("life AB foo").Parse()
	assert.EqualError(err,
		"Error on this line <life AB foo> (line: 1): Lifeline name (AB) must "+
			"be a single, upper case letter")
	_, err = NewParser("life a foo").Parse()
	assert.NotNil(err)
	assert.EqualError(err,
		"Error on this line <life a foo> (line: 1): Lifeline name (a) must "+
			"be a single, upper case letter")

	// Make sure it behaves the same way with the only other keywords that
	// requires a single lifeline to be specified: "stop".
	_, err = NewParser("stop a").Parse()
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lifeline name (a) must be "+
			"a single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lifeline to be specified: "stop".
	_, err = NewParser("stop a").Parse()
	assert.EqualError(err,
		"Error on this line <stop a> (line: 1): Lifeline name (a) must be "+
			"a single, upper case letter")

	// Make sure it behaves the same way with the other keywords that
	// requires a single lifeline to be specified: "self".
	_, err = NewParser("self a foo").Parse()
	assert.EqualError(err,
		"Error on this line <self a foo> (line: 1): Lifeline name (a) "+
			"must be a single, upper case letter")
}

func TestErrorWhenTwoLetterLifelinesExpected(
	t *testing.T) {
	assert := assert.New(t)

	// A few different scenarios

	// Upper case letter but only one of them, <full> keyword
	_, err := NewParser("full A foo").Parse()
	assert.EqualError(err, "Error on this line <full A foo> (line: 1): "+
		"Lifelines specified must be two, upper case letters")

	// Two letters but wrong case - dash keyword
	_, err = NewParser("dash ab foo").Parse()
	assert.EqualError(err, "Error on this line <dash ab foo> (line: 1): "+
		"Lifelines specified must be two, upper case letters")

	// Two characters but one is not a letter - dash keyword
	_, err = NewParser("dash A3 foo").Parse()
	assert.EqualError(err, "Error on this line <dash A3 foo> (line: 1): "+
		"Lifelines specified must be two, upper case letters")

	// Using the same letter twice
	_, err = NewParser(`
		life A foo
		full AA bar
	`).Parse()
	assert.EqualError(err, "Error on this line <full AA bar> (line: 3): "+
		"Lifeline letters must be different:(AA)")
}

func TestItIgnoresBlankLines(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(`
		life A  SL App

		life B  Core Permissions API
    `).Parse()
	assert.Nil(err)
	statements := model.Statements()
	assert.Len(statements, 2)
}

func TestItCapturesLabelTextWithNoLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(`
	life A SL App
	showletters false`).Parse()
	assert.Nil(err)
	statements := model.Statements()
	assert.Len(statements[0].LabelSegments, 1)
	assert.Equal("SL App", statements[0].LabelSegments[0], 1)
}

func TestItCapturesLabelTextWithLineBreaksIn(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser("life A  The quick | brown fox | etc").Parse()
	assert.Nil(err)
	statements := model.Statements()
	assert.Len(statements[0].LabelSegments, 5)
	// Note we check not only the splitting but also that each
	// segment is trimmed of whitespace.
	assert.Equal("The quick", statements[0].LabelSegments[0])
	assert.Equal("brown fox", statements[0].LabelSegments[1])
	assert.Equal("etc", statements[0].LabelSegments[2])
}

func TestErrorMessageWhenAnUnknownLifelineIsReferenced(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser("full AB foo").Parse()
	assert.EqualError(err,
		"Error on this line <full AB foo> (line: 1): Unknown lifeline: A")
}

func TestMakeSureEveryKeywordIsHandledWithoutError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser(ReferenceInput).Parse()
	assert.Nil(err)
}

func TestMakeSureARepresentativeStatementOutputIsProperlyFormed(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(ReferenceInput).Parse()
	assert.Nil(err)

	// full BC Validate user/pass
	statements := model.Statements()
	s := statements[5]
	assert.Equal("full", s.Keyword)
	assert.Equal("B", s.ReferencedLifelines[0].LifelineName)
	assert.Equal("C", s.ReferencedLifelines[1].LifelineName)
	assert.Equal("Validate user/pass", s.LabelSegments[0])
}

func TestNoErrorsWhenOptionalStatementsAreOmitted(t *testing.T) {
	assert := assert.New(t)
	// This DSL excludes {title, textsize, showletters} statements
	// Both of which are optional.
	_, err := NewParser(`
        life A foo
        self A bar
    `).Parse()
	assert.Nil(err)
}

func TestErrorsForMalformedTextSize(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser("textsize garbage").Parse()
	assert.EqualError(err,
		"Error on this line <textsize garbage> (line: 1): Text size must be a number")
	_, err = NewParser("textsize 21").Parse()
	assert.EqualError(err,
		"Error on this line <textsize 21> (line: 1): Text size must be between 5 and 20")
	_, err = NewParser("textsize 4").Parse()
	assert.EqualError(err,
		"Error on this line <textsize 4> (line: 1): Text size must be between 5 and 20")
}

func TestWellFormedTextSizeIsParsedCorrectly(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser("textsize 10").Parse()
	assert.NoError(err)
	statements := model.Statements()
	s := statements[0]
	expected := 10.0
	assert.Equal(expected, s.TextSize)
}

func TestErrorsForMalformedShowLetters(t *testing.T) {
	assert := assert.New(t)
	_, err := NewParser("showletters garbage").Parse()
	assert.EqualError(err,
		"Error on this line <showletters garbage> (line: 1): showletters expects <true> or <false>")
}

func TestWellFormedShowLettersIsParsedCorrectly(t *testing.T) {
	assert := assert.New(t)

	model, err := NewParser("showletters true").Parse()
	assert.NoError(err)
	statements := model.Statements()
	s := statements[0]
	assert.True(s.ShowLetters)

	model, err = NewParser("showletters false").Parse()
	assert.NoError(err)
	statements = model.Statements()
	s = statements[0]
	assert.False(s.ShowLetters)
}

func TestLifelineTitlesGetLettersWhenShowLettersIsTrue(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(`
		life A foo
		showletters true`).Parse()
	assert.NoError(err)
	lifeline := model.Statements()[0]
	assert.Equal("A", lifeline.LabelSegments[2])
}

func TestLifelineTitlesOmitLettersWhenShowLettersIsFalse(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(`
		life A foo
		showletters false`).Parse()
	assert.NoError(err)
	lifeline := model.Statements()[0]
	assert.Equal(lifeline.LabelSegments[0], "foo")
}

func TestLifelineTitlesGetLettersWhenShowLettersIsOmitted(t *testing.T) {
	assert := assert.New(t)
	model, err := NewParser(`
		life A foo`).Parse()
	assert.NoError(err)
	lifeline := model.Statements()[0]
	assert.Equal("A", lifeline.LabelSegments[2])
}
