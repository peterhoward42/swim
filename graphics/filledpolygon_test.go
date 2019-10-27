package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIncludesThisVertex(t *testing.T) {
	assert := assert.New(t)

	poly := FilledPoly{Point{2,3}}

	assert.True(poly.IncludesThisVertex(Point{2,3}))
	assert.False(poly.IncludesThisVertex(Point{2,3.1}))
	assert.False(poly.IncludesThisVertex(Point{2.1,3}))
}
