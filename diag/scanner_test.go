package diag

import (
	"testing"

	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/parser"
	"github.com/stretchr/testify/assert"
)

// This test feeds the Scanner with two Statement(s), and
// makes sure that it properly aggregates the required graphical events
// in the per-statement lookup table it builds.
func TestScannerProperlyAggregatesGraphicsEvents(t *testing.T) {
	assert := assert.New(t)

	statements := parser.MustCompileParse(`
		life A foo
		self A msg
	`)
	Lifeline := statements[0]
	self := statements[1]

	// Get the scanner to scan them, and scrutinise the *Events* attribute
	// it populates.
	scanner := newScanner()
	eventsLookup := scanner.Scan([]*dslmodel.Statement{Lifeline, self})

	assert.Len(eventsLookup[Lifeline], 1)
	assert.Len(eventsLookup[self], 2)
}
