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
then calling its InitFrameAndMakeTitleBox method,
it should produce a correctly sized and positioned frame title box,  and
also return the correctly calculated new tide mark.
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
	_ = tideMark

	// Strings?
	assert.Len(prims.Labels, 1)
	label := prims.Labels[0]
	assert.Equal(title, label.TheString)
	assert.Equal(float64(6), label.FontHeight)
	assert.Equal(float64(15), label.Anchor.X)
	assert.Equal(float64(10), label.Anchor.Y)
	assert.Equal(graphics.Left, label.HJust)
	assert.Equal(graphics.Top, label.VJust)

	assert.Len(prims.Lines, 4)
	tl = graphics.Point(1, 1)
	br = graphics.Point(2, 2)
	assert.True(prims.ContainsRect(tl, br))

	assert.Equal(9999, tideMark)
}
