package diag

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peterhoward42/umli/dslmodel"
)

func TestSortGaps(t *testing.T) {
	assert := assert.New(t)
	creator := NewCreator(2000, 40, []*dslmodel.Statement{})
	lifelines := NewLifelines(creator)

	a := &gap{1, 0}
	b := &gap{2, 0}
	c := &gap{3, 0}

	// Unaltered when already sorted?
	gaps := []*gap{a,b,c}
	lifelines.sortGaps(gaps)
    assert.Equal([]*gap{a,b,c}, gaps)

	// Reversed when pre-sorted backwards?
	gaps = []*gap{c,b,a}
	lifelines.sortGaps(gaps)
    assert.Equal([]*gap{a,b,c}, gaps)
}

func TestMergeGaps(t *testing.T) {
	assert := assert.New(t)
	creator := NewCreator(2000, 40, []*dslmodel.Statement{})
	lifelines := NewLifelines(creator)

	// Unaltered when do not overlap?
	a := &gap{1, 2}
	b := &gap{3, 4}
	c := &gap{5, 6}
	gaps := []*gap{a,b,c}
	newGaps := lifelines.mergeGaps(gaps)
    assert.Equal([]*gap{a,b,c}, newGaps)

	// Three reduced to two when last two overlap
	a = &gap{1, 2}
	b = &gap{3, 4}
	c = &gap{3.5, 4.5}
	gaps = []*gap{a,b,c}
	newGaps = lifelines.mergeGaps(gaps)
    assert.Len(newGaps, 2)
    assert.Equal(3.0, newGaps[1].topY)
    assert.Equal(4.5, newGaps[1].botY)

	// Make sure contiguous gaps are treated as overlapping
	a = &gap{1, 2}
	b = &gap{3, 4}
	b = &gap{4, 5}
	gaps = []*gap{a,b,c}
	newGaps = lifelines.mergeGaps(gaps)
    assert.Len(newGaps, 2)
}
