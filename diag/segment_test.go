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
	NewLifelines(creator)

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
}
