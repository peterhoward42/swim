package sizers

import (
	"github.com/peterhoward42/umli/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSizerComposesItselfProperly(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse("lane A foo")
	sizer := NewSizer(2000, 20.0, statements)
	// Check some top-level attributes.
	assert.InDelta(20.0, sizer.DiagPadT, 0.1)
	assert.InDelta(2000.0, sizer.Lanes.DiagramWidth, 0.1)
	assert.Equal(statements[0], sizer.Lanes.LaneStatements[0])
	// Have the embedded ndividual lane sizing data structures been
	// installed?
	assert.Equal(1, len(sizer.Lanes.Individual), "Should be one Lane")
}
