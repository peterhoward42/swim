package graphics

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLinesAccumulatesThem(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddLine(0, 0, 3, 3, false, false)
	p.AddLine(3, 3, 3, 3, false, false)
	assert.Len(p.Lines, 2)
}

func TestArrowHeadAngleArithmetic(t *testing.T) {
	// Arrow for due East line should have angle zero.
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddLine(0, 0, 3, 0, false, true)
	assert.InDelta(p.ArrowHeads[0].DirectionAngle, 0, 0.1)

	// Arrow for due North line should have angle PI / 2
	p = NewPrimitives()
	p.AddLine(0, 0, 0, 3, false, true)
	assert.InDelta(p.ArrowHeads[0].DirectionAngle, math.Pi / 2.0, 0.1)
}

func TestOnlyLinesWithArrowsGenerateArrowHeads(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddLine(0, 0, 0, 3, false, true)
	p.AddLine(0, 0, 0, 4, false, false)
	assert.Len(p.ArrowHeads, 1)
}

func TestAddRectDecomposesCorrectly(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddRect(0, 0, 4, 3)
	assert.Len(p.Lines, 4)
	assert.Equal(Line{0, 0, 4, 0, false}, *p.Lines[0])
	assert.Equal(Line{4, 0, 4, 3, false}, *p.Lines[1])
	assert.Equal(Line{4, 3, 0, 3, false}, *p.Lines[2])
	assert.Equal(Line{0, 3, 0, 0, false}, *p.Lines[3])
}

func TestAddAccumulatesProperly(t *testing.T) {
	assert := assert.New(t)

	a := NewPrimitives()
	a.AddLine(0, 0, 3, 3, false, false)
	a.AddLabel(nil, 0, 0, Left, Top)

	b := NewPrimitives()
	b.AddLine(0, 0, 3, 3, false, false)
	b.AddLabel(nil, 0, 0, Left, Top)

	a.Add(b)
	assert.Len(a.Lines, 2)
	assert.Len(a.Labels, 2)
}
