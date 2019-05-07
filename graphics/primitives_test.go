package graphics

import (
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
func TestAddRectDecomposesCorrectly(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddRect(0, 0, 4, 3)
	assert.Len(p.Lines, 4)
	assert.Equal(Line{0, 0, 4, 0, false, false}, *p.Lines[0])
	assert.Equal(Line{4, 0, 4, 3, false, false}, *p.Lines[1])
	assert.Equal(Line{4, 3, 0, 3, false, false}, *p.Lines[2])
	assert.Equal(Line{0, 3, 0, 0, false, false}, *p.Lines[3])
}

func TestAppendAccumulatesProperly(t *testing.T) {
	assert := assert.New(t)

	a := NewPrimitives()
	a.AddLine(0, 0, 3, 3, false, false)
	a.AddLabel(nil, 0, 0, Left, Top)

	b := NewPrimitives()
	b.AddLine(0, 0, 3, 3, false, false)
	b.AddLabel(nil, 0, 0, Left, Top)

	a.Append(b)
	assert.Len(a.Lines, 2)
	assert.Len(a.Labels, 2)
}
