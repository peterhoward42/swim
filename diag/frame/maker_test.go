package frame

import (
	"testing"

	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

/*
Given a frame.Maker, initialised with a Sizer that you can configure to
return fixed values,
then when calling its InitFrameAndMakeTitleBox method,
it should produce a correctly sized and positioned frame title box, with a title
inside in the right place. And also return the correctly calculated new
tide mark.
*/
func TestInitializeStep(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"FontHeight":         6,
		"FramePadLR":         11,
		"FrameTitleBoxWidth": 200,
		"FrameTitleRectPadB": 2,
		"FrameTitleTextPadL": 4,
		"FrameTitleTextPadB": 7,
		"FrameTitleTextPadT": 5,
	})
	prims := graphics.NewPrimitives()

	maker := NewMaker(sizer, prims)
	frameTop := 5.0
	title := "My title"
	tideMark := maker.InitFrameAndMakeTitleBox([]string{title}, frameTop)

	// Strings?
	assert.Len(prims.Labels, 1)
	label := prims.Labels[0]
	assert.Equal(title, label.TheString)
	assert.Equal(float64(6), label.FontHeight)
	assert.Equal(float64(15), label.Anchor.X)
	assert.Equal(float64(10), label.Anchor.Y)
	assert.Equal(graphics.Left, label.HJust)
	assert.Equal(graphics.Top, label.VJust)

	// Title Rect
	assert.Len(prims.Lines, 4)
	tl := graphics.Point{X: 11, Y: 5}
	br := graphics.Point{X:200, Y:23}
	assert.True(prims.ContainsRect(tl, br))

	// Tidemark
	assert.Equal(float64(25), tideMark)
}

/*
Given a frame.Maker, initialised with a Sizer that you can configure to
return fixed values,
then calling its FinalizeFrame method,
it should produce a correctly sized and positioned rectangle. And also return
the correctly calculated new tide mark.
*/
func TestFinalizeStep(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"FrameInternalPadB": 5,
		"FramePadLR": 7,
	})
	prims := graphics.NewPrimitives()

	maker := NewMaker(sizer, prims)
	initialTideMark := 200.0
	diagWidth := 500.0

	maker.frameTop = 10 // Simulates this state having been set in earlier step.
	tideMark := maker.FinalizeFrame(initialTideMark, diagWidth)

	// Title Rect
	assert.Len(prims.Lines, 4)
	tl := graphics.Point{X: 7, Y: 10}
	br := graphics.Point{X:493, Y:205}
	assert.True(prims.ContainsRect(tl, br))

	// Tidemark
	assert.Equal(float64(205), tideMark)
}
