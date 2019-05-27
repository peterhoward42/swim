/*
This module provides regression tests for a themed set of diagrams.  It
exercises the Creator type to produce a set of diagrams and to store them as
.png images in the ./testresults/new directory. It then goes on to check that
these are identical to the golden reference set of .png images in the
./testresults/goldenref directory.

The idea is that when the software is changed, a human can judge the fitness
for purpose of the new images produced visually, and when happy to copy them
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

// TestReferenceModel uses the reference DSL script and a typical
// diagram size and font size.
func TestReferenceModel(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := parser.ReferenceInput
	genericCreateHelper(t, script, width, fontHeight, "canonical.png")
}

// TestOneLane illustrates the sizing and composition logic in the context
// of a diagram with just one lane in.
func TestOneLane(t *testing.T) {
	width := 2000
	fontHeight := 20.0
	script := `
		lane A Foo
		self A Bar
		self A Baz
		self A A long | label over | multiple lines
	`
	genericCreateHelper(t, script, width, fontHeight, "onelane.png")
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
