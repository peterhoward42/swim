package diag

import (
	"testing"
	"path/filepath"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/peterhoward42/umli-export/imagefile"

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
