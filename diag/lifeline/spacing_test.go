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
When calling its CentreLine method for the second of the lifelines...
Then it should produce an X coordinate faithful to the internal algorithm.
*/
func TestCentreLineCanonicalCase(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"DiagWidth":                  800,
		"FontHt":                     20,
		"IdealLifelineTitleBoxWidth": 100.0,
	})

	lifelineA := &dsl.Statement{}
	lifelineB := &dsl.Statement{}
	lifelines := []*dsl.Statement{lifelineA, lifelineB}

	spacing := NewSpacing(sizer, lifelines)
	x, err := spacing.CentreLine(lifelineB)
	assert.NoError(err)
	assert.Equal(550.0, x)
}

/*
Given a Spacing object initialised with two lifelines, and a hard-coded
sizer that asks for huge title boxes...
When calling its CentreLine method for the second of the lifelines...
Then it should produce an X coordinate faithful to the internal algorithm which
reduces the title box size to maintain the gutter widths.
*/
func TestCentreLineSquashedCase(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"DiagWidth":                  800,
		"FontHt":                     20,
		"IdealLifelineTitleBoxWidth": 99999999.0,
	})

	lifelineA := &dsl.Statement{}
	lifelineB := &dsl.Statement{}
	lifelines := []*dsl.Statement{lifelineA, lifelineB}

	spacing := NewSpacing(sizer, lifelines)
	x, err := spacing.CentreLine(lifelineB)
	assert.NoError(err)
	assert.Equal(1195.0, x)
}

/*
Given a Spacing object initialised with a single lifelines, and a hard-coded
sizer...
When calling its CentreLine method for the second of the lifelines...
Then it should produce an X coordinate in the middle of the diagram's width.
*/
func TestCentreLineCornerCaseOfOneLifeline(t *testing.T) {
	assert := assert.New(t)
	_ = assert

	sizer := sizer.NewLiteralSizer(map[string]float64{
		"DiagWidth":                  800,
		"FontHt":                     20,
		"IdealLifelineTitleBoxWidth": 100.0,
	})

	lifelineA := &dsl.Statement{}
	lifelines := []*dsl.Statement{lifelineA}

	spacing := NewSpacing(sizer, lifelines)
	x, err := spacing.CentreLine(lifelineA)
	assert.NoError(err)
	assert.Equal(400.0, x)
}