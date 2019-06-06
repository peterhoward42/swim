package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const delta = 0.1

func TestHasExactlyOneVertexWithXRejectsWhenNone(t *testing.T) {
	assert := assert.New(t)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(0, 0)
	poly := &FilledPoly{[]*Point{p1, p2}}
	assert.False(poly.HasExactlyOneVertexWithX(1, delta))
}

func TestHasExactlyOneVertexWithXRejectsMoreThanOne(t *testing.T) {
	assert := assert.New(t)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(0, 0)
	poly := &FilledPoly{[]*Point{p1, p2}}
	assert.False(poly.HasExactlyOneVertexWithX(0, delta))
}

func TestHasExactlyOneVertexWithXMatchesInSimpleCase(t *testing.T) {
	assert := assert.New(t)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 0)
	poly := &FilledPoly{[]*Point{p1, p2}}
	assert.True(poly.HasExactlyOneVertexWithX(0, delta))
}

func TestHasExactlyOneVertexWithXMatchesDueToDelta(t *testing.T) {
	assert := assert.New(t)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 0)
	poly := &FilledPoly{[]*Point{p1, p2}}
	assert.True(poly.HasExactlyOneVertexWithX(0.5*delta, delta))
}

func TestHasExactlyOneVertexWithXRejectsDueToDelta(t *testing.T) {
	assert := assert.New(t)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 0)
	poly := &FilledPoly{[]*Point{p1, p2}}
	assert.False(poly.HasExactlyOneVertexWithX(1.5*delta, delta))
}
