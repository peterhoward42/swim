package graphics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsLine(t *testing.T) {
	assert := assert.New(t)

	type testcase struct {
		line Line
		expected bool
	}

	p1 := Point{1, 2}
	p2 := Point{9, 10}

	cases := []testcase{}
	cases = append(cases, testcase{ Line{p1, p2, true}, true})

	p := NewPrimitives()
	p.AddLine(1,2,9,10, true)

	for _, testcase := range cases {
		actual := p.ContainsLine(testcase.line)
		assert.Equal(actual, testcase.expected)
	}
}

