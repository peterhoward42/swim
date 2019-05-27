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

func TestOneLane(t *testing.T) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse("lane A foo")
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, "one-lane.png")
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}

func TestThreeLanes(t *testing.T) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
		lane C The | quick | brown | fox | jumps over
	`)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, "three-lanes.png")
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}

func TestInteractionLines(t *testing.T) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
		full AB [a guard] then | a multiline | label
		dash BA to show tidemark advancement
	`)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, "interaction-lines.png")
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}
func TestSelfLoop(t *testing.T) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
		self B go fetch some | info
	`)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, "self-lines.png")
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}

func TestReferenceModel(t *testing.T) {
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse(parser.ReferenceInput)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	fPath := filepath.Join(testResultsDir, "reference-model.png")
	err = imagefile.NewCreator(font).Create(
		fPath, imagefile.PNG, graphicsModel)
	assert.NoError(err)
}
