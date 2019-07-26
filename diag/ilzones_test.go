package diag

import (
	"testing"

	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

func TestClaimsCreatedForSingleInteractionLine(t *testing.T) {
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

	assert.Equal(142.5, claims[0].extent.start)
	assert.Equal(217.5, claims[0].extent.end)
	assert.Equal(217.5, claims[1].extent.start)
	assert.Equal(241.5, claims[1].extent.end)
}

func TestCrossesBehaviour(t *testing.T) {
	assert := assert.New(t)
	width := 3000
	fontHeight := 30.0

    // Single interaction line across 3 lifelinesthat crosses middle
    // one from left to right
	script := `
		life A foo
		life B bar
		life C baz
        full AC message
	`
	statements := parser.MustCompileParse(script)
	creator := NewCreator(width, fontHeight, statements)
	creator.Create()

    left := statements[0]
    middle := statements[1]
    right := statements[2]
    crosses := creator.ilZones.crosses(middle, left, right)
	assert.True(crosses)

    // Single interaction line across 3 lifelinesthat crosses middle
    // one from right to left
	script = `
		life A foo
		life B bar
		life C baz
        full CA message
	`
	statements = parser.MustCompileParse(script)
	creator = NewCreator(width, fontHeight, statements)
	creator.Create()

    left = statements[0]
    middle = statements[1]
    right = statements[2]
    crosses = creator.ilZones.crosses(middle, left, right)
	assert.True(crosses)

    // Three lifelines with interaction line between first two does
    // not cross the third.
	script = `
		life A foo
		life B bar
		life C baz
        full AB message
	`
	statements = parser.MustCompileParse(script)
	creator = NewCreator(width, fontHeight, statements)
	creator.Create()

    left = statements[0]
    middle = statements[1]
    right = statements[2]
    crosses = creator.ilZones.crosses(right, left, middle)
	assert.False(crosses)

    // Two lifelines with lifeline between them, does not cross either
    // the leftmost one.
	script = `
		life A foo
		life B bar
        full AB message
	`
	statements = parser.MustCompileParse(script)
	creator = NewCreator(width, fontHeight, statements)
	creator.Create()

    left = statements[0]
    right = statements[1]
    crosses = creator.ilZones.crosses(left, left, right)
	assert.False(crosses)
    crosses = creator.ilZones.crosses(right, left, right)
	assert.False(crosses)
}
