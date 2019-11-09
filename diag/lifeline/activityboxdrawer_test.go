package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/graphics"
	"github.com/stretchr/testify/assert"
)

func TestDrawsTheRightLines(t *testing.T) {
	assert := assert.New(t)

	// We need only one activity box on the lifeline to check correct
	// operation,
	boxes := NewActivityBoxes()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(60)
	assert.NoError(err)

	
	centreX := 100.0
	boxWidth := 10.0
	drawer := NewActivityBoxDrawer(*boxes, centreX, boxWidth)
	prims := graphics.NewPrimitives()
	drawer.Draw(prims)

	assert.Equal(4, len(prims.Lines))
	topLeft := graphics.Point{X: 95, Y:25}
	bottomRight := graphics.Point{X: 105, Y:60}
	assert.True(prims.ContainsRect(topLeft, bottomRight))
}