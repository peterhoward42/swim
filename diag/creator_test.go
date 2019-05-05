package diag

import (
	"testing"
	"strings"
	"bufio"

	"github.com/peterhoward42/umlinteraction/parser"
	"github.com/peterhoward42/umlinteraction/graphics"
	"github.com/stretchr/testify/assert"
)

func TestToTeaseOutAPI(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader(parser.ReferenceInput)
	p := parser.NewParser()
	statements, err := p.Parse(bufio.NewScanner(reader))
	assert.Nil(err)

    creator := NewCreator(statements)
    created := creator.Create()

    assert.IsType(&graphics.Model{}, created)
}
