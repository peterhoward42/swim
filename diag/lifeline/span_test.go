package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/dsl"
	"github.com/stretchr/testify/assert"
)

func TestSpanSimplestCase(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	a := &dsl.Statement{}
	b := &dsl.Statement{}
	c := &dsl.Statement{}
	allLifelines := []*dsl.Statement{a, b, c}
	span := SpanExcl(a, c, allLifelines)
	assert.Len(span, 1)
	assert.Equal(b, span[0])
}

func TestSpanSimplestCaseReversed(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	a := &dsl.Statement{}
	b := &dsl.Statement{}
	c := &dsl.Statement{}
	allLifelines := []*dsl.Statement{a, b, c}
	span := SpanExcl(c, a, allLifelines)
	assert.Len(span, 1)
	assert.Equal(b, span[0])
}

func TestSpanCountsRight(t *testing.T) {
	assert := assert.New(t)
	_ = assert
	a := &dsl.Statement{}
	b := &dsl.Statement{}
	c := &dsl.Statement{}
	d := &dsl.Statement{}
	allLifelines := []*dsl.Statement{a, b, c, d}
	span := SpanExcl(a, d, allLifelines)
	assert.Len(span, 2)
	assert.Equal(b, span[0])
	assert.Equal(c, span[1])
}
