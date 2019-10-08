package seg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSegments(t *testing.T) {
	assert := assert.New(t)

	a := Segment{1, 0}
	b := Segment{2, 0}
	c := Segment{3, 0}

	// Unaltered when already sorted?
	segs := []Segment{a, b, c}
	SortSegments(segs)
	assert.Equal([]Segment{a, b, c}, segs)

	// Reversed when pre-sorted backwards?
	segs = []Segment{c, b, a}
	SortSegments(segs)
	assert.Equal([]Segment{a, b, c}, segs)
}

func TestMergeSegments(t *testing.T) {
	assert := assert.New(t)

	// Unaltered when do not overlap?
	a := Segment{1, 2}
	b := Segment{3, 4}
	c := Segment{5, 6}
	segments := []Segment{a, b, c}
	newSegments := MergeSegments(segments)
	assert.Equal(segments, newSegments)

	// Three reduced to two when last two overlap
	a = Segment{1, 2}
	b = Segment{3, 4}
	c = Segment{3.5, 4.5}
	segments = []Segment{a, b, c}
	newSegments = MergeSegments(segments)
	assert.Len(newSegments, 2)
	assert.Equal(3.0, newSegments[1].Start)
	assert.Equal(4.5, newSegments[1].End)

	// Make sure contiguous Segments are treated as overlapping
	a = Segment{1, 2}
	b = Segment{3, 4}
	b = Segment{4, 5}
	segments = []Segment{a, b, c}
	newSegments = MergeSegments(segments)
	assert.Len(newSegments, 2)

	// Two reduced to one, when one swamps the other
	a = Segment{2, 3}
	b = Segment{1, 4}
	segments = []Segment{a, b}
	SortSegments(segments)
	newSegments = MergeSegments(segments)
	assert.Len(newSegments, 1)
	assert.Equal(1.0, segments[0].Start)
	assert.Equal(4.0, segments[0].End)
}

func TestLength(t *testing.T) {
	assert := assert.New(t)
	a := Segment{1, 4}
	assert.InDelta(3.0, a.Length(), 0.0001)
	a = Segment{4, 1}
	assert.InDelta(3.0, a.Length(), 0.0001)
	a = Segment{1, 1}
	assert.InDelta(0.0, a.Length(), 0.0001)
}
