package sizers

import (
	"testing"

	"github.com/peterhoward42/umli/parser"

	"github.com/stretchr/testify/assert"
)

func TestNewLifelinesSetsTitleBoxHeightWhenLabelHasOneLineOfText(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A foo
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)
	assert.InDelta(40.0, lifelines.TitleBoxHeight, 0.1)
}
func TestNewLifelinesSetsTitleBoxHeightWhenLabelHasThreeLinesOfText(
	t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A foo | bar | baz
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)
	assert.InDelta(80.0, lifelines.TitleBoxHeight, 0.1)
}
func TestNewLifelineLifelinesTitleBoxHeightWhenLabelsHaveDifferingHeight(
	t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A fibble
		life B foo | bar | baz
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)
	assert.InDelta(80.0, lifelines.TitleBoxHeight, 0.1)
}

func TestNewLifelinesSetsScalarAttributesCorrectlyForOneLifeline(t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A foo
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)

	assert.Equal(1, lifelines.NumLifelines)
	assert.InDelta(1333.3, lifelines.TitleBoxWidth, 0.1)
	assert.InDelta(40, lifelines.TitleBoxHeight, 0.1)
	assert.InDelta(1666.7, lifelines.TitleBoxPitch, 0.1)

	assert.InDelta(333.3, lifelines.TitleBoxPadR, 0.1)
	assert.InDelta(333.3, lifelines.FirstTitleBoxPadL, 0.1)
	assert.InDelta(30.0, lifelines.TitleBoxPadB, 0.1)
}

func TestNewLifelinesSetsScalarAttributesCorrectlyForTwoLifelines(
	t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A foo
		life B bar
	`)
	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)

	assert.Equal(2, lifelines.NumLifelines)
	assert.InDelta(727.2, lifelines.TitleBoxWidth, 0.1)
	assert.InDelta(909.1, lifelines.TitleBoxPitch, 0.1)
	assert.InDelta(181.8, lifelines.TitleBoxPadR, 0.1)
	assert.InDelta(181.8, lifelines.FirstTitleBoxPadL, 0.1)
}

func TestNewLifelinesSetsIndividualLifelineAttributesCorrectlyForTwoLifelines(
	t *testing.T) {
	assert := assert.New(t)
	statements := parser.MustCompileParse(`
		life A foo
		life B bar
	`)
	b := statements[1]

	diagWidth := 2000
	fontHeight := 20.0
	lifelines := NewLifelines(diagWidth, fontHeight, statements)
	individual := lifelines.Individual[b]

	assert.InDelta(1090.9, individual.TitleBoxLeft, 0.1)
	assert.InDelta(1454.5, individual.Centre, 0.1)
	assert.InDelta(1818.18, individual.TitleBoxRight, 0.1)

	assert.InDelta(1818.2, individual.SelfLoopRight, 0.1)
	assert.InDelta(1636.4, individual.SelfLoopCentre, 0.1)

	assert.InDelta(1439.5, individual.ActivityBoxLeft, 0.1)
	assert.InDelta(1469.5, individual.ActivityBoxRight, 0.1)
}
