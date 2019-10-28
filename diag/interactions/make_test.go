package interactions

import (
	"testing"

	"github.com/peterhoward42/umli/diag/lifeline"
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
	width := 2000.0
	fontHt := 10.0
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"ActivityBoxVerticalOverlap": 5.0,
		"ActivityBoxWidth":           40.0,
		"ArrowLen":                   10.0,
		"ArrowWidth":                 4.0,
		"IdealLifelineTitleBoxWidth": 300.0,
		"InteractionLinePadB":        4.0,
		"InteractionLineTextPadB":    5.0,
	})
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
	makerDependencies := NewMakerDependencies(
		fontHt, spacer, sizer, activityBoxes)
	interactionsMaker := NewMaker(makerDependencies, graphicsModel)
	tideMark := 30.0
	updatedTideMark, noGoZones, err := interactionsMaker.ScanInteractionStatements(
		tideMark, dslModel.Statements())
	assert.NoError(err)

	// Should have generated one line, one string, and one arrow in the graphics.
	assert.Len(graphicsModel.Primitives.Labels, 1)
	assert.Len(graphicsModel.Primitives.Lines, 1)
	assert.Len(graphicsModel.Primitives.FilledPolys, 1)

	// Inspect details of these primitives...

	// X coords of interaction line end points is plausible for a diagram
	// with two lifelines only.
	line := graphicsModel.Primitives.Lines[0]
	assert.True(line.P1.X > 0.1*width/10.0)
	assert.True(line.P1.X < 0.4*width)
	assert.True(line.P2.X > 0.6*width/10.0)
	assert.True(line.P2.X < 0.9*width)

	// Expect the label's anchor to sit exactly half way along the
	// interaction line, which is also exactly half the width of the diagram,
	// and with the anchor Y at the initial tidemark.
	expectedLabel := graphics.Label{
		TheString:  "fibble",
		FontHeight: fontHt,
		Anchor:     graphics.Point{X: 1000, Y: 30},
		HJust:      graphics.Centre,
		VJust:      graphics.Top,
	}
	assert.True(graphicsModel.Primitives.ContainsLabel(expectedLabel))

	// The interaction line Y should below the label by the space
	// taken up by the label text rows, plus the padding demanded by
	// the label.
	assert.InDelta(tideMark+1.0*fontHt+
		sizer.Get("InteractionLineTextPadB"), line.P1.Y, tolerance)
	assert.Equal(line.P1.Y, line.P2.Y)

	// The polygon representing the arrow should have a tip vertex
	// at the same as the interaction line's P2, and two others a little
	// way to the left, and distributed above and below the tip.
	arrow := graphicsModel.Primitives.FilledPolys[0]
	assert.True(arrow.IncludesThisVertex(line.P2))
	upperTail := graphics.Point{
		X: line.P2.X - sizer.Get("ArrowLen"),
		Y: line.P2.Y - 0.5*sizer.Get("ArrowWidth")}
	assert.True(arrow.IncludesThisVertex(upperTail))
	lowerTail := graphics.Point{
		X: line.P2.X - sizer.Get("ArrowLen"),
		Y: line.P2.Y + 0.5*sizer.Get("ArrowWidth")}
	assert.True(arrow.IncludesThisVertex(lowerTail))

	// The updated tide mark should be just below the interaction line,
	// by the amount of a padding demanded by an interaction line.
	assert.True(graphics.ValEqualIsh(updatedTideMark, line.P2.Y+sizer.Get(
		"InteractionLinePadB")))

	// Two no go zones should have been registered, with the correct X
	// coordinates, and heights that match those occupied by the label and line.
	assert.Len(noGoZones, 2)
	zn := noGoZones[0]
	assert.True(graphics.ValEqualIsh(zn.Height.Start, tideMark))
	assert.True(graphics.ValEqualIsh(zn.Height.End, line.P1.Y))
	zn = noGoZones[1]
	assert.True(graphics.ValEqualIsh(zn.Height.Start, line.P1.Y))
	assert.True(graphics.ValEqualIsh(zn.Height.End, updatedTideMark))

	// An activity box should have been registered as starting for lifeline A,
	// starting just abov the interaction line, and not yet terminated.
	lifeA := lifelines[0]
	boxSegs := activityBoxes[lifeA].AsSegments()
	assert.Len(boxSegs, 1)
	assert.True(graphics.ValEqualIsh(boxSegs[0].Start,
		line.P1.Y-sizer.Get("ActivityBoxVerticalOverlap")))
	assert.Equal(-1.0, boxSegs[0].End)

	// An activity box should have been registered as starting for lifeline B,
	// starting exactly at the interaction line, and not yet terminated.
	lifeB := lifelines[1]
	boxSegs = activityBoxes[lifeB].AsSegments()
	assert.Len(boxSegs, 1)
	assert.True(graphics.ValEqualIsh(boxSegs[0].Start, line.P1.Y))
	assert.Equal(-1.0, boxSegs[0].End)
}

const tolerance = 0.001
