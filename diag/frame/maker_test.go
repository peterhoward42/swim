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
		"FramePadLR":         11,
		"FrameTitleRectPadB": 2,
		"FrameTitleTextPadL": 4,
		"FrameTitleTextPadB": 7,
		"FrameTitleTextPadT": 5,
	})
	prims := graphics.NewPrimitives()
	fontHeight := 6.0
	diagWidth := 2000.0
	maker := NewMaker(sizer, fontHeight, diagWidth, prims)
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
	tl := graphics.NewPoint( 11, 5)
	br := graphics.NewPoint( 611, 23)
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
		"FramePadLR":        7,
	})
	prims := graphics.NewPrimitives()
	unusedFontHeight := 9999999999.0
	diagWidth := 2000.0
	maker := NewMaker(sizer, unusedFontHeight, diagWidth, prims)
	initialTideMark := 200.0

	maker.frameTop = 10 // Simulates this state having been set in earlier step.
	tideMark := maker.FinalizeFrame(initialTideMark)

	// Title Rect
	assert.Len(prims.Lines, 4)
	tl := graphics.NewPoint( 7, 10)
	br := graphics.NewPoint( 1993, 205)
	assert.True(prims.ContainsRect(tl, br))

	// Tidemark
	assert.Equal(float64(205), tideMark)
}
