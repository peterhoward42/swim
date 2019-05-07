package sizers

import (
	umli "github.com/peterhoward42/umlinteraction"
	"github.com/peterhoward42/umlinteraction/dslmodel"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLanesSetsScalarAttributesCorrectly(t *testing.T) {
	assert := assert.New(t)

	a := dslmodel.NewStatement()
	a.Keyword = umli.Lane
	a.LaneName = "A"

	b := dslmodel.NewStatement()
	b.Keyword = umli.Lane
	b.LaneName = "B"

	statements := []*dslmodel.Statement{a, b}

	diagWidth := 200
	fontHeight := 3.5
	lanes := NewLanes(diagWidth, fontHeight, statements)

	assert.Equal(2, lanes.NumLanes)
	assert.InDelta(72.7, lanes.TitleBoxWidth, 0.1)
	assert.InDelta(90.9, lanes.TitleBoxPitch, 0.1)
	assert.InDelta(18.2, lanes.TitleBoxHorizGap, 0.1)
	assert.InDelta(18.2, lanes.TitleBoxLeftMargin, 0.1)
}

func TestNewLanesSetsIndividualLaneAttributesCorrectly(t *testing.T) {
	assert := assert.New(t)

	a := dslmodel.NewStatement()
	a.Keyword = umli.Lane
	a.LaneName = "A"

	b := dslmodel.NewStatement()
	b.Keyword = umli.Lane
	b.LaneName = "B"

	statements := []*dslmodel.Statement{a, b}

	diagWidth := 200
	fontHeight := 3.5
	lanes := NewLanes(diagWidth, fontHeight, statements)
	individual := lanes.Individual[b]

	assert.InDelta(72.7, individual.TitleBoxLeft, 0.1)
	assert.InDelta(109.1, individual.Centre, 0.1)
	assert.InDelta(145.5, individual.TitleBoxRight, 0.1)
}
