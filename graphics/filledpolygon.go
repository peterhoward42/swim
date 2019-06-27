package graphics

import (
	"math"
)

// FilledPoly represents a filled polygon.
// Which can be used for an arrow head.
type FilledPoly struct {
	Vertices []*Point // Do not repeat first point as last point.
}

// NewFilledPoly provices a FilledPoly ready to use.
func NewFilledPoly(vertices []*Point) *FilledPoly {
	return &FilledPoly{vertices}
}

// HasExactlyOneVertexWithX asserts that this polygon has one, and only one
// vertex that has the given X coordinate. (within delta)
func (p *FilledPoly) HasExactlyOneVertexWithX(x, delta float64) bool {
	count := 0
	for _, v := range p.Vertices {
		dx := math.Abs(v.X - x)
		if dx < delta {
			count++
		}
	}
	return count == 1 // Must be exactly one.
}
