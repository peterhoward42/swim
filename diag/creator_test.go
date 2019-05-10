package diag

import (
	"testing"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

/*
Tests to have
	o  create doesn't crash
	o  when dsl has only one lane and nothing else
	o  a lane gets a title box
*/

func TestCreateRunsWithoutCrashing(t *testing.T) {
	statements := parser.MustCompileParse(parser.ReferenceInput)
	width := 200
	fontHeight := 3.0
	creator := NewCreator(width, fontHeight, statements)
	creator.Create()
}

func TestForLaneTitleBoxWhenOnlyThingPresentIsOneLane(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse("lane A foo")
	width := 200
	fontHeight := 3.5
	creator := NewCreator(width, fontHeight, statements)
	graphicsModel := creator.Create()
	prims := graphicsModel.Primitives
	assert.Len(prims.Lines, 4)
}
