package diag

import (
	"os"
	"testing"

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

// When the environment variable UMLI_VISUAL_TESTS is set to "true",
// some of the tests in this module, output .png images in /tmp for visual
// inspection instead of programmatic inspection of what Creator has produced.
func visualTestMode() bool {
	return os.Getenv("UMLI_VISUAL_TESTS") == "true"
}

func TestOneLaneOnlyVisuals(t *testing.T) {
	if visualTestMode() != true {
		t.Skipf("Fibble")
	}
	assert := assert.New(t)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	statements := parser.MustCompileParse("lane A foo")
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	err = imagefile.NewCreator(font).Create(
		"/tmp/one-lane.png", imagefile.PNG, graphicsModel)
	assert.NoError(err)
}

func TestOneLaneOnlyRegression(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse("lane A foo")
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	assert.Len(graphicsModel.Primitives.Lines, 4)
}
