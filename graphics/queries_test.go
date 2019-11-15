package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const lessThanTol = 0.000001

func TestValEqualIsh(t *testing.T) {
	assert := assert.New(t)

	const lessThanTol = 0.000001

	assert.True(ValEqualIsh(1.3, 1.3))
	assert.True(ValEqualIsh(1.3, 1.3+lessThanTol))
	assert.True(ValEqualIsh(1.3, 1.3-lessThanTol))
	assert.True(ValEqualIsh(1.3, 1.3+lessThanTol))
	assert.True(ValEqualIsh(1.3, 1.3-lessThanTol))

}
func TestPointEqualIsh(t *testing.T) {
	assert := assert.New(t)

	assert.True(Point{1, 2}.EqualIsh(Point{1, 2}))
	assert.True(Point{1, 2}.EqualIsh(Point{1, 2 + lessThanTol}))
	assert.True(Point{1, 2}.EqualIsh(Point{1, 2 - lessThanTol}))
	assert.False(Point{1, 2}.EqualIsh(Point{1, 2 + 0.3}))
	assert.False(Point{1, 2}.EqualIsh(Point{1, 2 - 0.3}))
}

func TestLabelEqualIsh(t *testing.T) {
	assert := assert.New(t)

	referenceLabel := Label{
		TheString:  "foo",
		FontHeight: 10.0,
		Anchor:     Point{2, 3},
		HJust:      Left,
		VJust:      Top,
	}
	assert.True(referenceLabel.EqualIsh(
		Label{
			TheString:  "foo",
			FontHeight: 10.0,
			Anchor:     Point{2, 3},
			HJust:      Left,
			VJust:      Top,
		}))
	assert.True(referenceLabel.EqualIsh(
			Label{
				TheString:  "foo",
				FontHeight: 10.0,
				Anchor:     Point{2, 3 + lessThanTol},
				HJust:      Left,
				VJust:      Top,
			}))
}

/*
Given a Primitives containing only one Line: <X>,
Then when calling ContainsLine with various test Line objects,
It should return true only when the test Line object has the same
(or equivalent reversed) geometry as X, and has the same dashed property as X.
*/
func TestContainsLine(t *testing.T) {
	assert := assert.New(t)

	type testcase struct {
		line     Line
		expected bool
	}

	p1 := Point{1, 2}
	p2 := Point{9, 10}
	p3 := Point{9, 11}

	cases := []testcase{}
	cases = append(cases, testcase{Line{p1, p2, true}, true})   // canonical
	cases = append(cases, testcase{Line{p2, p1, true}, true})   // reversed
	cases = append(cases, testcase{Line{p1, p3, true}, false})  // different geom
	cases = append(cases, testcase{Line{p1, p2, false}, false}) // not dashed

	p := NewPrimitives()
	p.AddLine(1, 2, 9, 10, true)

	for _, testcase := range cases {
		actual := p.ContainsLine(testcase.line)
		assert.Equal(actual, testcase.expected)
	}
}

/*
Given a Primitives containing only 4 lines that make a rectangle <X>,
Then when calling ContainsRect with various test rectangles,
It should return true only when the test rectangle has the same geometry
as X.
*/
func TestContainsRect(t *testing.T) {
	assert := assert.New(t)

	top := 1.0
	left := 2.0
	bot := 3.0
	right := 4.0

	p := NewPrimitives()
	p.AddRect(left, top, right, bot)

	// Canonical case.
	present := p.ContainsRect(Point{left, top}, Point{right, bot})
	assert.True(present)

	// Specify corners opposite way round.
	present = p.ContainsRect(Point{right, bot}, Point{left, top})
	assert.True(present)

	// Fractionally different geom.
	present = p.ContainsRect(Point{left + 0.1, top}, Point{right, bot})
	assert.False(present)
}

func TestContainsLabel(t *testing.T) {
	assert := assert.New(t)

	p := NewPrimitives()
	p.AddLabel("foo", 3.0, 600.0, 200.0, Left, Top)

	label := Label{
		TheString:  "foo",
		FontHeight: 3.0,
		Anchor:     Point{X: 600, Y: 200},
		HJust:      Left,
		VJust:      Top,
	}
	assert.True(p.ContainsLabel(label))
	label.Anchor.X += 1
	assert.False(p.ContainsLabel(label))
}

func TestContentsBoundingBox(t *testing.T) {
	assert := assert.New(t)
	p := NewPrimitives()
	dashed := true
	p.AddLine(10, 11, 100, 101, dashed)
	p.AddLine(20, 3, 150, 80, dashed)
	left, top, right, bottom := p.BoundingBoxOfLines()
	assert.Equal(10.0, left)
	assert.Equal(3.0, top)
	assert.Equal(150.0, right)
	assert.Equal(101.0, bottom)
}
