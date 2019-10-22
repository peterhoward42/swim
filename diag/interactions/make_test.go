package interactions

import (
	"testing"

	"github.com/peterhoward42/umli/diag/lifeline"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/parser"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestSingleInteractionLine(t *testing.T) {
	assert := assert.New(t)

	dslScript := `
		life A foo
		life B bar
		full AB fibble
	`
	dslModel := parser.MustCompileParse(dslScript)

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"IdealLifelineTitleBoxWidth": 300.0,
	})
	width := 2000.0
	fontHt := 10.0
	lifelines := dslModel.LifelineStatements()
	spacer := lifeline.NewSpacing(sizer, fontHt, width, lifelines)
	makerDependencies := NewMakerDependencies(fontHt, spacer)

	dashLineDashLength := 5.0
	dashLineGapLength := 1.0
	graphicsModel := graphics.NewModel(
		width, fontHt, dashLineDashLength, dashLineGapLength)

	interactionsMaker := NewMaker(makerDependencies, graphicsModel)
	tideMark := 1.0
	updatedTideMark, err := interactionsMaker.Scan(tideMark, dslModel.Statements())
	assert.NoError(err)

	// Should be one just one line, one string, and one arrow in the graphics.
	
	assert.Len(graphicsModel.Primitives.Labels, 1)
	assert.Len(graphicsModel.Primitives.Lines, 1)
	assert.Len(graphicsModel.Primitives.FilledPolys, 1)

	// inspect details of these prims

	_ = updatedTideMark
}
