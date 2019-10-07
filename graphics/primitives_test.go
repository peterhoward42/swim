package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRectDecomposesCorrectly(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	p.AddRect(0, 0, 4, 3)
	assert.Len(p.Lines, 4)
	assert.Equal(Line{Point{0, 0}, Point{4, 0}, false}, *p.Lines[0])
	assert.Equal(Line{Point{4, 0}, Point{4, 3}, false}, *p.Lines[1])
	assert.Equal(Line{Point{4, 3}, Point{0, 3}, false}, *p.Lines[2])
	assert.Equal(Line{Point{0, 3}, Point{0, 0}, false}, *p.Lines[3])

	sampleLine := p.Lines[3]
	start := sampleLine.P1
	end := sampleLine.P2
	assert.Equal(Point{0, 3}, start)
	assert.Equal(Point{0, 0}, end)
}

func TestRowOfStringsDecomposesCorrectly(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	x := 5.0
	topRow := 20.0
	fontHeight := 3.0
	p.RowOfStrings(x, topRow, fontHeight, Right, []string{"foo", "bar"})
	assert.Len(p.Labels, 2)
	// Inspect the second of the two to make sure it has the geometry it
	// should (including the offset below the first row)
	expected := Label{
		TheString:  "bar",
		FontHeight: fontHeight,
		Anchor:     Point{x, topRow + fontHeight},
		HJust:      Right,
		VJust:      Top,
	}
	assert.Equal(expected, *p.Labels[1])
}

func TestAddAccumulatesProperly(t *testing.T) {
	assert := assert.New(t)

	a := NewPrimitives()
	a.AddLine(0, 0, 3, 3, false)
	a.AddFilledPoly([]Point{})
	a.AddLabel("", 0, 0, 0, Left, Top)

	b := NewPrimitives()
	b.AddLine(0, 0, 3, 3, false)
	b.AddFilledPoly([]Point{})
	b.AddLabel("", 0, 0, 0, Left, Top)

	a.Add(b)
	assert.Len(a.Lines, 2)
	assert.Len(a.FilledPolys, 2)
	assert.Len(a.Labels, 2)
}
