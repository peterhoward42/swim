package sizers

import (
	"github.com/peterhoward42/umli/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSizerComposesItselfProperly(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse("lane A foo")
	sizer := NewSizer(200, 3.5, statements)
	assert.InDelta(7.0, sizer.TopMargin, 0.1)
	assert.InDelta(200, sizer.Lanes.DiagramWidth, 0.1)
	assert.Equal(statements[0], sizer.Lanes.LaneStatements[0])
}
