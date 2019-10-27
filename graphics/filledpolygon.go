package graphics

import (
)

// FilledPoly represents a filled polygon.
// Which can be used for an arrow head.
type FilledPoly []Point // Do not repeat first point as last point.

// IncludesThisVertex asserts that this polygon has one, and only one
// vertex matching p.
func (p *FilledPoly) IncludesThisVertex(q Point) bool {
	count := 0
	for _, v := range *p {
		if v.EqualIsh(q) {
			count++
		}
	}
	return count == 1 // Must be exactly one.
}
