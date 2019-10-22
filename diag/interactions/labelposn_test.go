package interactions

import (
	"testing"

	"github.com/peterhoward42/umli/graphics"
	"github.com/stretchr/testify/assert"
)

func TestAnswerForLeftToRightLabel(t *testing.T) {
	assert := assert.New(t)
	from := 20.0
	to := 100.0
	x, horizJustification := NewLabelPosn(from, to).Get()
	assert.Equal(60.0, x)
	assert.Equal(graphics.Centre, horizJustification)
}

