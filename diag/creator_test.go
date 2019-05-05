package diag

import (
	"testing"
	"strings"
	"bufio"

	"github.com/peterhoward42/umlinteraction/parser"
	"github.com/peterhoward42/umlinteraction/graphics"
	"github.com/stretchr/testify/assert"
)

func TestToTeaseOutAPIDuringDevelopment(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader(parser.ReferenceInput)
	p := parser.NewParser()
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)

	creator := NewCreator()
	// These widths and heights are chosen to be similar to the size
	// of A4 paper (in mm), to help think about the sizing abstractions.
	width := 200
	height := 297
    created := creator.Create(statements, width, height)

    assert.IsType(&graphics.Model{}, created)
}
