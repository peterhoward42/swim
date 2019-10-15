/*
This module provides regression tests for a themed set of diagrams.  It
exercises the Creator type to produce a set of diagrams and to store them as
.png images in the ./testresults/new directory. It then goes on to check that
these are identical to the golden reference set of .png images in the
./testresults/goldenref directory.

The idea is that when the software is changed, a human can judge the fitness
for purpose of the new images produced visually, and when happy, to copy them
in to the the golden reference directory as the new standard.
*/
package diag

import (
	"path/filepath"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/peterhoward42/umli/render"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

var testResultsDir = filepath.Join(".", "testresults", "new")

// TestReferenceModel uses the reference DSL script and a typical
// diagram size and font size.
func TestReferenceModel(t *testing.T) {
	script := parser.ReferenceInput
	genericCreateHelper(t, script, "canonical.png")
}

/*
TestStopStartBox *stops* a lifeline activity box explicitly with
a *stop* line in the DSL, and then sends a message to that lifeline
later in the script, to check that a new activity box gets started.
*/
func TestStopStartBox(t *testing.T) {
	script := `
        life A foo
        life B bar
        full AB apple
        stop B
        full AB banana
    `
	genericCreateHelper(t, script, "stopstartbox.png")
}

/*
TestIgnoresRedundantStop uses a script that tries to *stop* a lifeline
activity box, when that lifeline doesn't have a box in progress. It provides
visual confirmation that the statement is silently ignored.
*/
func TestIgnoresRedundantStop(t *testing.T) {
	script := `
		life A foo
		life B bar
		full AB baz
		stop B
		self A henrietta:w
        stop B
	`
	genericCreateHelper(t, script, "redundantstop.png")
}

// TestOneLifeline makes a diagram with just one lifeline - to help reveal
// corner cases.
func TestOneLifeline(t *testing.T) {
	script := `
		life A Foo
		self A Bar
		self A Baz
		self A A long | label over | multiple lines
	`
	genericCreateHelper(t, script, "onelifeline.png")
}

// TestLargeNumberOfLifelines illustrates the sizing and composition logic
// when there are a large number of Lifelines. It helps itself by citing a
// small text size.
func TestLargeNumberOfLifelines(t *testing.T) {
	script := `
		title This Is The | Title
		textsize 10
		life A two word | title
		life B two word | title
		life C two word | title
		life D two word | title
		life E two word | title
		life F two word | title
		life G two word | title
		life H two word | title
		life I two word | title
		life J two word | title

        full AB three word | message
        full BC three word | message
        full CD three word | message
        full DA three word | message
        dash AD three word | message
        full DI three word | message
        self I three word | message
        self I three word | message
	`
	genericCreateHelper(t, script, "manylifelines.png")
}

// TestLargeNumberOfInteractions illustrates the sizing and composition
// logic when there are a large number of interactions.
func TestLargeNumberOfInteraction(t *testing.T) {
	script := `
		life A two word | title
		life B two word | title
		life C two word | title

        full AB three word | message
        full BC three word | message
        full CA three word | message
        full AB three word | message
        full BC three word | message
        full CA three word | message
        full AB three word | message
        full BC three word | message
        full CA three word | message
        full AB three word | message
        full BC three word | message
        full CA three word | message
        full AB three word | message
        full BC three word | message
        full CA three word | message
        full AB three word | message
	`
	genericCreateHelper(t, script, "manyinteractions.png")
}

// TestLargeFont illustrates the sizing and composition
// logic when font is relatively large.
func TestLargeFont(t *testing.T) {
	script := `
		title Large Font
		textsize 20
		life A Foo
		life B Bar
		full AB a message
	`
	genericCreateHelper(t, script, "largfont.png")
}

// TestSmallFont illustrates the sizing and composition
// logic when font is relatively small.
func TestSmallFont(t *testing.T) {
	script := `
		title Small Font
		textsize 5
		life A Foo
		life B Bar
		full AB a message
	`
	genericCreateHelper(t, script, "smallfont.png")
}

// TestWhenAllOptionalStatementsAreAbsent makes sure default text height and
// show letters gets used when explicit statements for these are not in the script.
func TestWhenAllOptionalStatementsAreAbsent(t *testing.T) {
	script := `
		life A Foo
		life B Bar
		full AB a message
	`
	genericCreateHelper(t, script, "optionsmissing.png")
}

// TestExplicitTurningOffLetters makes sure you can turn the lifeline letters
// off from the script
func TestExplicitTurningOffLetters(t *testing.T) {
	script := `
		showletters false
		life A Foo
		life B Bar
		full AB a message
	`
	genericCreateHelper(t, script, "noletters.png")
}

// Helper functions (DRY)

// genericCreateHelper makes a diagram from the given DSL script and
// saves it in the ./testresults/new directory.
func genericCreateHelper(t *testing.T, script string, imageBaseName string) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(script)
	// todo why not literal one liner?
	creator := &Creator{}
	graphicsModel := creator.Create(statements)
	fPath := filepath.Join(testResultsDir, imageBaseName)
	err = render.NewImageFileCreator(font).Create(
		fPath, render.PNG, graphicsModel)
	assert.NoError(err)
}
