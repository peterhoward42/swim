package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This test feeds the Scanner with two Statement(s), and
// makes sure that it properly aggregates the required graphical events
// in the per-statement lookup table it builds.
func TestScannerProperlyAggregatesGraphicsEvents(t *testing.T) {
	assert := assert.New(t)

	statements := parser.MustCompileParse(`
		lane A foo
		self A msg
	`)
	lane := statements[0]
	self := statements[1]

	// Get the scanner to scan them, and scrutinise the *Events* attribute
	// it populates.
	scanner := NewScanner()
	eventsLookup := scanner.Scan([]*dslmodel.Statement{lane, self})

	assert.Len(eventsLookup[lane], 3)
	assert.Len(eventsLookup[self], 2)
}
