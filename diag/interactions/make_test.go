package interactions

import (
	"testing"

	"github.com/peterhoward42/umli/diag/lifeline"
	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/parser"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestForSingleInteractionLineItProducesCorrectGraphicsAndSideEffects(t *testing.T) {
	assert := assert.New(t)

	dslScript := `
		life A foo
		life B bar
		full AB fibble
	`
	dslModel := parser.MustCompileParse(dslScript)
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"ActivityBoxWidth":           40.0,
		"ArrowLen":                   10.0,
		"IdealLifelineTitleBoxWidth": 300.0,
		"InteractionLinePadB":        4.0,
		"InteractionLineTextPadB":    5.0,
	})
	width := 2000.0
	fontHt := 10.0
	lifelines := dslModel.LifelineStatements()
	spacer := lifeline.NewSpacing(sizer, fontHt, width, lifelines)
	dashLineDashLength := 5.0
	dashLineGapLength := 1.0
	graphicsModel := graphics.NewModel(
		width, fontHt, dashLineDashLength, dashLineGapLength)

	activityBoxes := map[*dsl.Statement]*lifeline.ActivityBoxes{}
	for _, ll := range lifelines {
		activityBoxes[ll] = lifeline.NewActivityBoxes()
	}
	noGoZones := []*nogozone.NoGoZone{}
	makerDependencies := NewMakerDependencies(
		fontHt, spacer, sizer, activityBoxes, noGoZones)
	interactionsMaker := NewMaker(makerDependencies, graphicsModel)
	tideMark := 30.0
	updatedTideMark, err := interactionsMaker.Scan(tideMark, dslModel.Statements())
	assert.NoError(err)

	// Should be one just one line, one string, and one arrow in the graphics.
	assert.Len(graphicsModel.Primitives.Labels, 1)
	assert.Len(graphicsModel.Primitives.Lines, 1)
	assert.Len(graphicsModel.Primitives.FilledPolys, 1)

	// Inspect details of these primitives...

	// Interaction line lateral dimensions are plausible.
	line := graphicsModel.Primitives.Lines[0]
	assert.True(line.P1.X > 0.1*width/10.0)
	assert.True(line.P1.X < 0.4*width)
	assert.True(line.P2.X > 0.6*width/10.0)
	assert.True(line.P2.X < 0.9*width)

	// Label anchor is at tidemark, and line sits a little below it.
	label := graphicsModel.Primitives.Labels[0]
	assert.Equal(tideMark, label.Anchor.Y)
	assert.True(line.P1.Y > label.Anchor.Y + fontHt)
	assert.True(line.P1.Y < label.Anchor.Y+)

	// Label lateral dimensions, and other properties are correct.
	assert.True(label.Anchor.X > line.P1.X)
	assert.True(label.Anchor.X < line.P2.X)
	assert.Equal("fibble", label.TheString)
	assert.Equal(10.0, label.FontHeight)
	assert.Equal(graphics.Centre, label.HJust)
	assert.Equal(graphics.Top, label.VJust)

	// Label Y is at initial tidemark, and interaction line is a little below
	// that.

	// Arrow tip is at line.P2, and other vertices are above or below it, and
	// to its left, and sensible distance away.

	// Updated tidemark is a little to the South of the interaction line.

	// Two no go zones got registered, with the correct X coordinates, and
	// heights that match those occupied by the label and line.

	// A Box was registered as starting for lifeline A, starting just above
	// the interaction line.

	// A Box was registered as starting for lifeline B, starting exactly at
	// the interaction line.

	_ = updatedTideMark
}
