package diag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	NewLifelineSpacing(2000, 20, 3)
}

func TestCanCallCentreLineMethod(t *testing.T) {
	assert := assert.New(t)
	sp := NewLifelineSpacing(2000, 20, 3)
	x := sp.CentreLine(1)
	assert.InDelta(1000, x, 0.1)
}
