package sizers

import (
	umli "github.com/peterhoward42/umlinteraction"
	"github.com/peterhoward42/umlinteraction/dslmodel"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSizerComposesItselfProperly(t *testing.T) {
	assert := assert.New(t)
	s := dslmodel.NewStatement()
	s.Keyword = umli.Lane
	s.LaneName = "A"
	statements := []*dslmodel.Statement{s,}
	sizer := NewSizer(200, 3.5, statements)
	assert.InDelta(7.0, sizer.TopMargin, 0.1)
	assert.InDelta(200, sizer.Lanes.DiagramWidth, 0.1)
	assert.Equal(s, sizer.Lanes.LaneStatements[0])
}
