package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestSimpleOneLifelineExampleExhaustively(t *testing.T) {
	assert := assert.New(t)

	lifeline := &dsl.Statement{
		Keyword:       umli.Life,
		LifelineName:  "A",
		LabelSegments: []string{"foo"},
	}
	lifelines := []*dsl.Statement{lifeline}

	prims := graphics.NewPrimitives()

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"DiagWidth": 2000,
		"FontHt": 6,
		"IdealLifelineTitleBoxWidth": 200.0,
		"TitleBoxLabelPadB": 2,
		"TitleBoxLabelPadT": 5,
	})
	spacer := NewSpacing(sizer, lifelines)
	tideMark := 10.0

	newTideMark, err := TitleBoxes{}.Make(tideMark, lifelines, sizer, spacer, prims)
	assert.NoError(err)

	top := 10.0
	left := 900.0
	bot := 23.0
	right := 1100.0

	tl := graphics.Point{X: left, Y: top}
	br := graphics.Point{X: right, Y: bot}

	assert.True(prims.ContainsRect(tl, br))

	// contains rect of given x
	// contains string with content foo anchored at top left below top of rect
	// new tidemark is a little below bottom of box

	_ = newTideMark
}
