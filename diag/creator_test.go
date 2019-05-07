package diag

import (
	"testing"
	"strings"
	"bufio"

	"github.com/peterhoward42/umlinteraction/parser"
	"github.com/peterhoward42/umlinteraction/graphics"
	"github.com/stretchr/testify/assert"
)


/*
Tests to have
	o  create doesn't crash
	o  when dsl has only one lane and nothing else
	o  a lane gets a title box
*/

func TestCreateRunsWithoutCrashing(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader(parser.ReferenceInput)
	p := parser.NewParser()
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)
	width := 200
	fontHeight := 3.0
	creator := NewCreator(width, fontHeight, statements)
    creator.Create()
}


func TestALaneGetsATitleBox(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader(parser.ReferenceInput)
	p := parser.NewParser()
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)

	// These widths and heights are chosen to be similar to the size
	// of A4 paper (in mm), to help think about the sizing abstractions.
	width := 200
	fontHeight := 3.0
	creator := NewCreator(width, fontHeight, statements)
    created := creator.Create()

    assert.IsType(&graphics.Model{}, created)
}
