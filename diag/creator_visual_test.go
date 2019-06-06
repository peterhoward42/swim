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

	"github.com/peterhoward42/umli-export/imagefile"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

var testResultsDir = filepath.Join(".", "testresults", "new")

// TestScratch is a wip test to isolate something during dev
func TestScratch(t *testing.T) {
	width := 3000
	fontHeight := 30.0
	script := `
		life A hbuild | store
		life B dp_graph | neptune | hierarchy
		life C dp_graph | neptune | query
		life D dp_graph | neptune.go
		life E gremgo | neptune | pool
		life F gremgo | neptune | client
		life G graphson
		full AB CountNodes
		full BC CountHierarchyNodes
		dash CB "g.V().hasLabel ..."
		stop C
		full BD GetNumber(query)
		full DE GetCount
		self D retry
		full EF GetCount
		full FG Deserialize
		dash GF number json
		stop G
		
	`
	genericCreateHelper(t, script, width, fontHeight, "scratch.png")
}

/*
self E retry
		full EF Deserialize
		dash FE count json
		stop F
*/

// TestReferenceModel uses the reference DSL script and a typical
// diagram size and font size.
func TestReferenceModel(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := parser.ReferenceInput
	genericCreateHelper(t, script, width, fontHeight, "canonical.png")
}

/*
TestStopStartBox *stops* a lifeline activity box explicitly with
a *stop* line in the DSL, and then sends a message to that lifeline
later in the script, to check that a new activity box gets started.
*/
func TestStopStartBox(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := `
        life A foo
        life B bar
        full AB apple
        dash BA orange
        stop B
        full AB banana
    `
	genericCreateHelper(t, script, width, fontHeight, "stopstartbox.png")
}

/*
TestIgnoresRedundantStop uses a script that tries to *stop* a lifeline
activity box, when that lifeline doesn't have a box in progress. It provides
visual confirmation that the statement is silently ignored.
*/
func TestIgnoresRedundantStop(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := `
		life A foo
		life B bar
		full AB baz
		stop B
		self A henrietta:w
        stop B
	`
	genericCreateHelper(t, script, width, fontHeight, "redundantstop.png")
}

// TestOneLifeline makes a diagram with just one lifeline - to help reveal
// corner cases.
func TestOneLifeline(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := `
		life A Foo
		self A Bar
		self A Baz
		self A A long | label over | multiple lines
	`
	genericCreateHelper(t, script, width, fontHeight, "onelifeline.png")
}

// TestLargeNumberOfLifelines illustrates the sizing and composition logic
// when there are a large number of Lifelines.
func TestLargeNumberOfLifelines(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := `
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
	genericCreateHelper(t, script, width, fontHeight, "manylifelines.png")
}

// TestLargeNumberOfInteractions illustrates the sizing and composition
// logic when there are a large number of interactions.
func TestLargeNumberOfInteraction(t *testing.T) {
	width := 2000
	fontHeight := 20.0
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
	genericCreateHelper(t, script, width, fontHeight, "manyinteractions.png")
}

// TestLargeFont illustrates the sizing and composition
// logic when font is relatively large.
func TestLargeFont(t *testing.T) {
	width := 2000
	fontHeight := 40.0
	script := parser.ReferenceInput
	genericCreateHelper(t, script, width, fontHeight, "largfont.png")
}

// TestSmallFont illustrates the sizing and composition
// logic when font is relatively small.
func TestSmallFont(t *testing.T) {
	width := 2000
	fontHeight := 10.0
	script := parser.ReferenceInput
	genericCreateHelper(t, script, width, fontHeight, "smallfont.png")
}

// TestSmallDiagram illustrates the sizing and composition
// logic when the diagram size (width) is specified to be rather
// small.
func TestSmallDiagram(t *testing.T) {
	width := 1000
	fontHeight := 10.0
	script := parser.ReferenceInput
	genericCreateHelper(t, script, width, fontHeight, "smalldiag.png")
}

// Helper functions (DRY)

// genericCreateHelper makes a diagram from the given DSL script and
// saves it in the ./testresults/new directory.
func genericCreateHelper(t *testing.T, script string, width int,
	fontHeight float64, imageBaseName string) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(script)
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, imageBaseName)
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}
