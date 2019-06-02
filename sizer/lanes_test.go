package sizers

import (
	"github.com/peterhoward42/umli/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLanesSetsTitleBoxHeightWhenLabelHasOneLineOfText(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)
	assert.InDelta(40.0, lanes.TitleBoxHeight, 0.1)
}
func TestNewLanesSetsTitleBoxHeightWhenLabelHasThreeLinesOfText(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo | bar | baz
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)
	assert.InDelta(80.0, lanes.TitleBoxHeight, 0.1)
}
func TestNewLanesSetsTitleBoxHeightWhenLabelsHaveDifferingHeight(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A fibble
		lane B foo | bar | baz
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)
	assert.InDelta(80.0, lanes.TitleBoxHeight, 0.1)
}

func TestNewLanesSetsScalarAttributesCorrectlyForOneLane(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)

	assert.Equal(1, lanes.NumLanes)
	assert.InDelta(1333.3, lanes.TitleBoxWidth, 0.1)
	assert.InDelta(40, lanes.TitleBoxHeight, 0.1)
	assert.InDelta(1666.7, lanes.TitleBoxPitch, 0.1)

	assert.InDelta(25, lanes.TitleBoxBottomRowOfText, 0.1)
	assert.InDelta(333.3, lanes.TitleBoxPadR, 0.1)
	assert.InDelta(333.3, lanes.FirstTitleBoxPadL, 0.1)
	assert.InDelta(30.0, lanes.TitleBoxPadB, 0.1)
}

func TestNewLanesSetsScalarAttributesCorrectlyForTwoLanes(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)

	assert.Equal(2, lanes.NumLanes)
	assert.InDelta(727.2, lanes.TitleBoxWidth, 0.1)
	assert.InDelta(909.1, lanes.TitleBoxPitch, 0.1)
	assert.InDelta(181.8, lanes.TitleBoxPadR, 0.1)
	assert.InDelta(181.8, lanes.FirstTitleBoxPadL, 0.1)
}

func TestNewLanesSetsIndividualLaneAttributesCorrectlyForTwoLanes(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		lane A foo
		lane B bar
	`)
	b := statements[1]

	diagWidth := 2000
	fontHeight := 20.0
	lanes := NewLanes(diagWidth, fontHeight, statements)
	individual := lanes.Individual[b]

	assert.InDelta(1090.9, individual.TitleBoxLeft, 0.1)
	assert.InDelta(1454.5, individual.Centre, 0.1)
	assert.InDelta(1818.18, individual.TitleBoxRight, 0.1)

	assert.InDelta(1818.2, individual.SelfLoopRight, 0.1)
	assert.InDelta(1636.4, individual.SelfLoopCentre, 0.1)

	assert.InDelta(1439.5, individual.ActivityBoxLeft, 0.1)
	assert.InDelta(1469.5, individual.ActivityBoxRight, 0.1)
}
