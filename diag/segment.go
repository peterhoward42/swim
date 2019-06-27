package diag

/*
This module contains code that is capable of modelling a one dimensional  
extent in terms of a start and end coordinate. I.e. a *segment*. 
They are used elsewhere in the package to represent lifeline segments, or a 
gap in a lifeline. These only need a one dimensional extent because we know
they are vertical.

It additionally offers *sort* and *merge* services on a lists of segments.
*/

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

// sortSegments orders (in-situ) a list of segment objects by their start 
// attribute, smallest-first.
func sortSegments(segs []*segment) {
	sort.Slice(segs, func(i, j int) bool {
		return segs[i].start < segs[j].start
	})
}

// mergeGaps takes an *ordered* list of segments and merges any that overlap.
func mergeSegments(segs []*segment) (newSegs []*segment) {
	// NB. this DEPENDS on the segs being ordered by seg.start.

    // The process is:

	// Always keep the first segment.
	// Preserve a segment that is entirely beyond the previous one.
	// Throw away a segment that is entirely within the previous one.
	// Throw away a segment that overlaps with the previous one, but lengthen
	// the previous one accordingly.
	newSegs = []*segment{}
	var prev *segment
	for i, seg := range segs {
		if (i == 0) || (seg.start >= prev.end) {
			newSegs = append(newSegs, seg)
			prev = seg
			continue
		}
		if seg.start >= prev.start && seg.end <= prev.end {
			continue
		}
		if seg.start >= prev.start && seg.end >= prev.end {
			prev.end = seg.end
			continue
		}
	}
	return newSegs
}
