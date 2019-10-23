package geom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenX2IsGreater(t *testing.T) {
	assert := assert.New(t)
	x1 := 0.0
	x2 := 100.0
	shortenBy := 5.0
	ShortenLineBy(shortenBy, &x1, &x2)
	assert.Equal(x1, 5.0)
	assert.Equal(x2, 95.0)
}

func TestWhenX1IsGreater(t *testing.T) {
	assert := assert.New(t)
	x1 := 100.0
	x2 := 0.0
	shortenBy := 5.0
	ShortenLineBy(shortenBy, &x1, &x2)
	assert.Equal(x1, 95.0)
	assert.Equal(x2, 5.0)
}
