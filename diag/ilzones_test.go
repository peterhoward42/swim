package diag

import (
	"testing"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

func TestInteractionLineZones(t *testing.T) {
	assert := assert.New(t)
	width := 3000
	fontHeight := 30.0
	script := `
		life A foo
		life B bar
        full AB two | lines
	`
	statements := parser.MustCompileParse(script)
	creator := NewCreator(width, fontHeight, statements)
	creator.Create()

	// There should be two claims - one for the label and one (contiguously)
	// for the line.
	claims := creator.ilZones.claims
	assert.Len(claims, 2)

	assert.Equal(135.0, claims[0].extent.start)
	assert.Equal(210.0, claims[0].extent.end)
	assert.Equal(210.0, claims[1].extent.start)
	assert.Equal(234.0, claims[1].extent.end)
}
