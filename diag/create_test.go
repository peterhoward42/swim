package diag

import (
	"testing"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

func TestCreateRunsWithoutCrashing(t *testing.T) {
	assert := assert.New(t)
	creator, err := NewCreator()
	assert.NoError(err)
	var dslModel dsl.Model
	graphicsModel, err := creator.Create(dslModel)
	_ = graphicsModel
	assert.NoError(err)
}

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
	// Lifelines:   			4 // broken with one activity box apiece.
	// Activity boxes: 			8
	// Frame:					4

	numLines := len(prims.Lines)
	assert.Equal(28, numLines)

	// How many arrows?

	// How many labels?

	// Plausible model size?

	// Bounding box of all graphics just inside the model size?
}
