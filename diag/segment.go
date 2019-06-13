package diag

import (
	"sort"
)

/*
Type segment encapsulates and manipulate *segment*s - ie
one-dimensional extents with a *start* and an *end*. These could be used
to model a gap, or a line segment for example.
*/
type segment struct {
	start float64
	end   float64
}

// sortSegments orders (in-situ) a list of gap objects by their start attribute,
// least-first.
func sortSegments(segs []*segment) {
	sort.Slice(segs, func(i, j int) bool {
		return segs[i].start < segs[j].start
	})
}

// mergeGaps takes an *ordered* list of gaps and merges any that overlap.
func mergeSegments(segs []*segment) (newSegs []*segment) {
	newSegs = []*segment{}
	for i, g := range segs {
		if i == 0 || (g.start > newSegs[i-1].end) {
			newSegs = append(newSegs, g)
		} else {
			newSegs[i-1].end = g.end
		}
	}
	return newSegs
}
