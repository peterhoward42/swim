package frame

import (
	"testing"

	"github.com/peterhoward42/umli/diag/frame"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestInitializeStep(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.TestSizer(map[string]float64{
		"foo": 45,
	})
	prims := graphics.NewPrimitives()

	maker := frame.NewMaker(sizer, prims)
	frameTop := 5.0
	tideMark := maker.InitFrameAndMakeTitleBox([]string{"title"}, frameTop)
	_ = tideMark

	// row of strings faithfull content and sizing
	// ditto rect
	// returned tidemark right
}
