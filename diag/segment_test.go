package diag

import (
	"testing"

	"github.com/peterhoward42/umli/dslmodel"
	"github.com/stretchr/testify/assert"
)

func TestSortSegments(t *testing.T) {
	assert := assert.New(t)

	a := &segment{1, 0}
	b := &segment{2, 0}
	c := &segment{3, 0}

	// Unaltered when already sorted?
	segs := []*segment{a, b, c}
	sortSegments(segs)
	assert.Equal([]*segment{a, b, c}, segs)

	// Reversed when pre-sorted backwards?
	segs = []*segment{c, b, a}
	sortSegments(segs)
	assert.Equal([]*segment{a, b, c}, segs)
}

func TestMergeSegments(t *testing.T) {
	assert := assert.New(t)
	creator := NewCreator(2000, 40, []*dslmodel.Statement{})
	newLifelineMaker(creator)

	// Unaltered when do not overlap?
	a := &segment{1, 2}
	b := &segment{3, 4}
	c := &segment{5, 6}
	segments := []*segment{a, b, c}
	newSegments := mergeSegments(segments)
	assert.Equal([]*segment{a, b, c}, newSegments)

	// Three reduced to two when last two overlap
	a = &segment{1, 2}
	b = &segment{3, 4}
	c = &segment{3.5, 4.5}
	segments = []*segment{a, b, c}
	newSegments = mergeSegments(segments)
	assert.Len(newSegments, 2)
	assert.Equal(3.0, newSegments[1].start)
	assert.Equal(4.5, newSegments[1].end)

	// Make sure contiguous segments are treated as overlapping
	a = &segment{1, 2}
	b = &segment{3, 4}
	b = &segment{4, 5}
	segments = []*segment{a, b, c}
	newSegments = mergeSegments(segments)
	assert.Len(newSegments, 2)

	// Two reduced to one, when one swamps the other
	a = &segment{2, 3}
	b = &segment{1, 4}
	segments = []*segment{a, b}
	sortSegments(segments)
	newSegments = mergeSegments(segments)
	assert.Len(newSegments, 1)
	assert.Equal(1.0, segments[0].start)
	assert.Equal(4.0, segments[0].end)
}

func TestLength(t *testing.T) {
	assert := assert.New(t)
	a := &segment{1, 4}
	assert.InDelta(3.0, a.Length(), 0.0001)
	a = &segment{4, 1}
	assert.InDelta(3.0, a.Length(), 0.0001)
	a = &segment{1, 1}
	assert.InDelta(0.0, a.Length(), 0.0001)
}
