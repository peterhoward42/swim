package diag

import (
	"testing"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

func TestMainOrchestrationStepsWereProperlyStrungTogether(t *testing.T) {
	assert := assert.New(t)
	dslScript := `
		life A foo
		life B bar
		full AB fibble
	`
	dslModel := parser.MustCompileParse(dslScript)
	creator, err := NewCreator()
	assert.NoError(err)
	graphicsModel, err := creator.Create(*dslModel)
	assert.NoError(err)

	prims := graphicsModel.Primitives

	// How many lines should we expect?
	// Title box: 				4
	// Lifeline title boxes:	8
	// Interaction line         1
	// Lifelines:   			4 // Each broken with activity box
	// Activity boxes: 			8
	// Frame:					4
	numLines := len(prims.Lines)
	assert.Equal(29, numLines)

	// One arrow
	assert.Equal(1, len(prims.FilledPolys))

	// How many strings?
	// Implicit title:									1
	// Lifeline titles, including implicit lifeline
	// lifeline letter and blank line spacing:			6
	// Interaction label								1
	assert.Equal(8, len(prims.Labels))

	// Plausible finalized model size?

	// At the time of writing the default virtual model width is 2000.
	// Expect the depth to be a small-ish proportion of the width, because
	// it only has one lifeline.
	assert.Equal(2000.0, graphicsModel.Width)
	assert.True(graphicsModel.Height > 0.1*graphicsModel.Width)
	assert.True(graphicsModel.Height < 0.25*graphicsModel.Width)

	// Bounding box of all graphics just inside model size?
	left, top, right, bottom := prims.BoundingBoxOfLines()
	assert.True(left > 0.0 && left < 0.10*graphicsModel.Width)
	assert.True(top > 0.0 && top < 0.10*graphicsModel.Height)
	assert.True(right < graphicsModel.Width && right > 0.90*graphicsModel.Width)
	assert.True(bottom < graphicsModel.Height && bottom > 0.90*graphicsModel.Height)
}
