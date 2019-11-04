package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/parser"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestGapLogicWithSimpleCase(t *testing.T) {
	assert := assert.New(t)

	// Our test case will have 3 lifelines from left to right A,B,C,
	// There will be one NoGoZone and one ActivityBox interrupting
	// lifeline B.

	dslScript := `
		life A foo
		life B bar
		life C baz
	`
	dslModel := parser.MustCompileParse(dslScript)
	width := 2000.0
	fontHt := 10.0
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"IdealLifelineTitleBoxWidth": 300.0,
	})
	lifelines := dslModel.LifelineStatements()
	spacer := NewSpacing(sizer, fontHt, width, lifelines)
	noGoSeg := geom.Segment{Start: 50, End: 60}
	zone := nogozone.NewNoGoZone(noGoSeg, lifelines[0], lifelines[2])
	noGoZones := []nogozone.NoGoZone{zone}
	activityBoxes := map[*dsl.Statement]*ActivityBoxes{}
	for _, ll := range lifelines {
		activityBoxes[ll] = NewActivityBoxes()
	}
	boxesForLifeB := activityBoxes[lifelines[1]]
	err := boxesForLifeB.AddStartingAt(80)
	assert.NoError(err)
	err = boxesForLifeB.TerminateAt(90)
	assert.NoError(err)

	lifelineF := NewFinalizer(lifelines, noGoZones, activityBoxes)
	top := 10.0
	bottom := 100.0
	primitives := graphics.NewPrimitives()
	err = lifelineF.Finalize(top, bottom, primitives)
	assert.NoError(err)

	// The lines created for lifelines A and C should run from top to bottom
	// uninterrupted.
	lifeACoords, err := spacer.CentreLine(lifelines[0])
	assert.NoError(err)
	expectedX := lifeACoords.Centre
	expectedLineA := graphics.Line{
		P1:     graphics.Point{X: expectedX, Y: top},
		P2:     graphics.Point{X: expectedX, Y: bottom},
		Dashed: true,
	}
	assert.True(primitives.ContainsLine(expectedLineA))

	// The lines created for lifeline B should have gaps in.

	// segs should be various top to bottom coords

	// tidemark should be what it should be
}
