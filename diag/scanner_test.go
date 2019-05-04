package diag

import (
	"testing"
	umli "github.com/peterhoward42/umlinteraction"
	"github.com/peterhoward42/umlinteraction/dslmodel"
	"github.com/stretchr/testify/assert"
)

// This test feeds the Scanner with two Statement(s), and
// makes sure that it properly aggregates the required graphical events
// in the per-statement lookup table it builds.
func TestScannerProperlyAggregatesGraphicsEvents(t *testing.T) {
	assert := assert.New(t)

	// Prepare two Statements
	lane := dslmodel.NewStatement()
	lane.Keyword = umli.Lane
	lane.LaneName = "A"
	self := dslmodel.NewStatement()
	self.Keyword = umli.Self
	self.ReferencedLanes = []*dslmodel.Statement{lane,}

	// Get the scanner to scan them, and scrutinise the *Events* attribute
	// it populates.
	scanner := NewScanner()
	scanner.Scan([]*dslmodel.Statement{lane, self})

	assert.Len(scanner.Events[lane], 3)
	assert.Len(scanner.Events[self], 2)
}
