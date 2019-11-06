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

func TestLifelinesGetDrawnCorrectlyIncludingMakingTheRequiredGaps(t *testing.T) {
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
		"FrameInternalPadB":          10,
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

	minSegLen := 1.0
	lifelineF := NewFinalizer(lifelines, spacer, noGoZones, activityBoxes, sizer)
	top := 10.0
	bottom := 100.0
	primitives := graphics.NewPrimitives()
	updatedTidemark, err := lifelineF.Finalize(top, bottom, minSegLen, primitives)
	assert.NoError(err)

	// The lines created for lifelines A and C (only) should run from
	// top to bottom uninterrupted.
	for _, i := range []int{0, 2} {
		lifeCoords, err := spacer.CentreLine(lifelines[i])
		assert.NoError(err)
		expectedX := lifeCoords.Centre
		expectedLine := graphics.Line{
			P1:     graphics.Point{X: expectedX, Y: top},
			P2:     graphics.Point{X: expectedX, Y: bottom},
			Dashed: true,
		}
		assert.True(primitives.ContainsLine(expectedLine))
	}

	// The lines created for lifeline B should have gaps in.
	lifeCoords, err := spacer.CentreLine(lifelines[1])
	assert.NoError(err)
	expectedX := lifeCoords.Centre
	expectedSegments := []geom.Segment{
		geom.Segment{Start: 10, End: 50},
		geom.Segment{Start: 60, End: 80},
		geom.Segment{Start: 90, End: 100},
	}
	for _, seg := range expectedSegments {
		expectedLine := graphics.Line{
			P1:     graphics.Point{X: expectedX, Y: seg.Start},
			P2:     graphics.Point{X: expectedX, Y: seg.End},
			Dashed: true,
		}
		assert.True(primitives.ContainsLine(expectedLine))
	}

	// tidemark should be what it should be
	assert.True(graphics.ValEqualIsh(updatedTidemark, bottom+sizer.Get("FrameInternalPadB")))
}
