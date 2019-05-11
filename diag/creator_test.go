package diag

import (
	"os"
	"testing"

	"github.com/peterhoward42/umli/parser"
	"github.com/peterhoward42/umli/export"
	"github.com/stretchr/testify/assert"
)

/*
Tests to have
	o  create doesn't crash
	o  when dsl has only one lane and nothing else
	o  a lane gets a title box
*/

// When the environment variable UMLI_VISUAL_TESTS is set to "true",
// some of the tests in this module, output .png images in /tmp for visual
// inspection instead of programmatic inspection of what Creator has produced.
func visualTestMode() bool {
	return os.Getenv("UMLI_VISUAL_TESTS") == "true"
}

func TestCreateRunsWithoutCrashing(t *testing.T) {
	statements := parser.MustCompileParse(parser.ReferenceInput)
	width := 200
	fontHeight := 3.0
	creator := NewCreator(width, fontHeight, statements)
	creator.Create()
}

func TestWhenTheOnlyThingPresentIsOneLaneYouGetALaneTitleBox(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse("lane A foo")
	width := 2000
	fontHeight := 20.0
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	if visualTestMode() {
		err := export.CreatePNG("/tmp/one-lane.png", graphicsModel)
		assert.NoError(err)
	} else {
		prims := graphicsModel.Primitives
		assert.Len(prims.Lines, 4)
	}
}
