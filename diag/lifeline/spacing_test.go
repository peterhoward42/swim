package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

/*
Given a Spacing object initialised with two lifelines, and a hard-coded
sizer...
Then when calling its CentreLine object for the second of the lifelines...
It should produce an X coordinate at roughly 2/3 of the diagram width.
*/
func TestCentreLine(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	diagWidth := 800
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"DiagWidth": 800,
		"FontHt": 20,
		"IdealLifelineTitleBoxWidth": 100.0,
	})

	lifelineA := &dsl.Statement{}
	lifelineB := &dsl.Statement{}
	lifelines := []*dsl.Statement{lifelineA, lifelineB}

	spacing := NewSpacing(sizer, lifelines)
	x, err := spacing.CentreLine(lifelineB)
	assert.NoError(err)
	expected := diagWidth * 2.0 / 3.0
	assert.InDelta(expected, x, 20)
}

// convert to accurate test
// add edge case test
// add when crowded test to show special case kicks in
