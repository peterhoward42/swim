package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peterhoward42/umli/graphics"
)

/*
Given hard-coded input parameters,
When when makeArrow is called,
Then, it should produce geometrically correct vertices.
*/
func TestRightwardsArrowIsWellFormed(t *testing.T) {
	assert := assert.New(t)
	x1 := 0.0
	x2 := 10.0
	y := 5.0
	arrowLen := 2.0
	arrowHeight := 1.0
	vertices := makeArrow(x1, x2, y, arrowLen, arrowHeight)
	assert.Equal(graphics.Point{X: 8, Y: 4.5}, vertices[0])
	assert.Equal(graphics.Point{X: 10, Y: 5.0}, vertices[1])
	assert.Equal(graphics.Point{X: 8, Y: 5.5}, vertices[2])
}
