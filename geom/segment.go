package geom

import (
	"math"
	"sort"
)

/*
Segment can be used represent a gap in a line, or a short section of
a line. It provides methods to sort and merge segments when you have
multiple Segments that belong to the same underlying line.

Note it is modelled in one-dimension only, and has no knowledge of
the underlying continuous line it relates to.
*/
type Segment struct {
	Start float64
	End   float64
}

// Length provides the length of the segment.
func (s *Segment) Length() float64 {
	return math.Abs(s.Start - s.End)
}

/*
SortSegments orders (in-situ) a slice of Segment objects by their Start
attribute, smallest-first.

Note the (confusing) pass by value semantics. It still has the effect of
sorting the caller's copy, because the copy made when the parameter is passed
in, (by definition) still refers to the same underlying array. Hence sorting
the copy is swapping the positions of items in the underlying array they share
in common.
*/
func SortSegments(segs []Segment) {
	sort.Slice(segs, func(i, j int) bool {
		return segs[i].Start < segs[j].Start
	})
}

// MergeSegments takes an *ordered* list of Segments and merges any that overlap.
func MergeSegments(segs []Segment) (newSegs []Segment) {
	// NB. this DEPENDS on the segs being pre-ordered by seg.Start.

	// The process is:

	// Always keep the first segment.
	// For each subsequent segment...
	// Preserve it, if it is entirely beyond the previous one.
	// Throw it away if it is entirely within the previous one.
	// Throw it away if it overlaps with the previous one, but lengthen
	// the previous one accordingly.
	newSegs = []Segment{}
	var tail *Segment
	for i, seg := range segs {
		if i == 0 {
			newSegs = append(newSegs, seg)
			tail = &newSegs[0]
			continue
		}
		if seg.Start >= tail.End {
			newSegs = append(newSegs, seg)
			tail = &newSegs[len(newSegs)-1]
			continue
		}
		if seg.Start >= tail.Start && seg.End <= tail.End {
			continue
		}
		if seg.Start >= tail.Start && seg.End >= tail.End {
			tail.End = seg.End
			continue
		}
	}
	return newSegs
}
