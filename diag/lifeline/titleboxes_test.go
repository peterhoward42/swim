package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestSmallestViableExampleExhaustively(t *testing.T) {
	assert := assert.New(t)

	lifeline := &dsl.Statement{
		Keyword:       umli.Life,
		LifelineName:  "A",
		LabelSegments: []string{"foo"},
	}
	lifelines := []*dsl.Statement{lifeline}

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"IdealLifelineTitleBoxWidth": 200.0,
		"TitleBoxLabelPadB":          2,
		"TitleBoxLabelPadT":          5,
		"TitleBoxPadB":               3,
	})
	fontHeight := 6.0
	diagWidth := 2000.0
	spacer := NewSpacing(sizer, fontHeight, diagWidth, lifelines)
	tideMark := 10.0
	prims := graphics.NewPrimitives()
	titleBoxes := NewTitleBoxes(sizer, spacer, lifelines, fontHeight)
	newTideMark, bottomOfBoxes, err := titleBoxes.Make(tideMark, prims)
	assert.NoError(err)

	// Correct title box rectangle present?
	top := 10.0
	left := 900.0
	bot := 23.0
	right := 1100.0
	tl := graphics.NewPoint( left, top)
	br := graphics.NewPoint( right, bot)
	assert.True(prims.ContainsRect(tl, br))

	// Correct title string present?
	expectedLabel := graphics.Label{
		TheString:  "foo",
		FontHeight: 6.0,
		Anchor:     graphics.NewPoint( 1000, 15),
		HJust:      graphics.Centre,
		VJust:      graphics.Top,
	}
	assert.True(prims.ContainsLabel(expectedLabel))

	assert.Equal(23.0, bottomOfBoxes)
	assert.Equal(26.0, newTideMark)
}

func TestIsConsumingMultipleLifelinesProperly(t *testing.T) {
	assert := assert.New(t)

	// Run it with a single lifeline, only to capture some metrics.
	lifelineA := &dsl.Statement{
		Keyword:       umli.Life,
		LifelineName:  "A",
		LabelSegments: []string{"foo"},
	}
	lifelines := []*dsl.Statement{lifelineA}
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"IdealLifelineTitleBoxWidth": 200.0,
		"TitleBoxLabelPadB":          2,
		"TitleBoxLabelPadT":          5,
		"TitleBoxPadB":               3,
	})
	fontHeight := 6.0
	diagWidth := 2000.0
	spacer := NewSpacing(sizer, fontHeight, diagWidth, lifelines)
	tideMark := 10.0
	prims := graphics.NewPrimitives()
	titleBoxes := NewTitleBoxes(sizer, spacer, lifelines, fontHeight)
	newTideMark, _, err := titleBoxes.Make(tideMark, prims)
	assert.NoError(err)

	// Capture metrics
	firstRunTidemark := newTideMark
	linesProduced := len(prims.Lines)

	// Run it again with two lifelines present, and make sure there are more
	// lines produced, but the tidemark returned does not change.
	lifelineB := &dsl.Statement{
		Keyword:       umli.Life,
		LifelineName:  "A",
		LabelSegments: []string{"foo"},
	}
	lifelines = []*dsl.Statement{lifelineA, lifelineB}
	spacer = NewSpacing(sizer, fontHeight, diagWidth, lifelines)
	tideMark = 10.0
	prims = graphics.NewPrimitives()
	titleBoxes = NewTitleBoxes(sizer, spacer, lifelines, fontHeight)
	newTideMark, _, err = titleBoxes.Make(tideMark, prims)

	newLinesProduced := len(prims.Lines)
	assert.True(newLinesProduced > linesProduced)
	assert.Equal(firstRunTidemark, newTideMark)
}
