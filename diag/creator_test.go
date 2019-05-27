package diag

import (
	"path/filepath"
	"testing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/peterhoward42/umli-export/imagefile"

	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

/*
Tests to have
	o  create doesn't crash
	o  when dsl has only one lane and nothing else
	o  a lane gets a title box
	o  to be continued...
*/

var testResultsDir = filepath.Join(".", "testresults", "new")

func TestOneLaneOnlyVisuals(t *testing.T) {
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

func TestThreeLanesOnlyVisuals(t *testing.T) {
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

func TestInteractionLineVisuals(t *testing.T) {
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
func TestSelfLoopVisuals(t *testing.T) {
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

func TestInteractionLineQuantitatively(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
		full AB two line | label
	`)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	// Inspect the position and content of the interaction line
	// label strings (which will be the last two added)
	n := len(graphicsModel.Primitives.Labels)
	labels := graphicsModel.Primitives.Labels[n-2 : n]

	firstL := labels[0]
	assert.Equal(graphics.NewPoint(1000, 70), firstL.Anchor)
	assert.Equal("two line", firstL.TheString)
	assert.Equal(graphics.Centre, firstL.HJust)
	assert.Equal(graphics.Top, firstL.VJust)

	secondL := labels[1]
	assert.Equal(graphics.NewPoint(1000, 90), secondL.Anchor)

	// Inspect the end points for the interaction line (which will be
	// the last one added.
}

func TestInteractionLineGetsArrowAtRightEndFacingRightWay(t *testing.T) {
	// Make sure the arrows created for interaction lines are
	// at the right end, and point the right way.
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
		full AB first label
		full BA second label
	`)
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	n := len(graphicsModel.Primitives.Lines)
	rightToLeftLine := graphicsModel.Primitives.Lines[n-1]
	xLeft := rightToLeftLine.P2.X
	xRight := rightToLeftLine.P1.X
	leftRightArrow := graphicsModel.Primitives.FilledPolys[0]
	rightLeftArrow := graphicsModel.Primitives.FilledPolys[1]
	delta := 0.1
	assert.True(leftRightArrow.HasExactlyOneVertexWithX(xRight, delta),
		"Wrong position or direction")
	assert.True(rightLeftArrow.HasExactlyOneVertexWithX(xLeft, delta),
		"Wrong position or direction")
}

func TestReferenceModelVisuals(t *testing.T) {
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
