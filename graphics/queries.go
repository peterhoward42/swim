package graphics

/*
This module provides some queries that can be made against the model's
contents. For example does it contain the four lines required to represent
a given rectangle?

The queries are optimised for code simplicity (mainly for tests) at the
expense of Big-O complexity.
*/

const tol = 0.001

// ValEqualIsh checks for equality, while treating values with only
// very small differences as equal.
func ValEqualIsh(m, n float64) bool {
	if m > n+tol {
		return false
	}
	if m < n-tol {
		return false
	}
	return true
}

// EqualIsh checks for equality of p with another point, using thes
// same tolerance as ValEqualIsh.
func (p Point) EqualIsh(q Point) bool {
	if !ValEqualIsh(p.X, q.X) {
		return false
	}
	if !ValEqualIsh(p.Y, q.Y) {
		return false
	}
	return true
}

// EqualIsh checks for equality of lbl with otherLabel, using Point.EqualIsh
// for the anchor points.
func (lbl Label) EqualIsh(otherLabel Label) bool {
	if lbl.TheString != otherLabel.TheString {
		return false
	}
	if lbl.FontHeight != otherLabel.FontHeight {
		return false
	}
	if !lbl.Anchor.EqualIsh(otherLabel.Anchor) {
		return false
	}
	if lbl.HJust != otherLabel.HJust {
		return false
	}
	if lbl.VJust != otherLabel.VJust {
		return false
	}
	return true
}

// ContainsRect evaluates if the primitives contain the lines required to
// represent the rectangle defined by tl and br?
func (p *Primitives) ContainsRect(tl, br Point) bool {

	l := tl.X
	r := br.X
	t := tl.Y
	b := br.Y

	required := []Line{}

	required = append(required, Line{Point{l, t}, Point{r, t}, false})
	required = append(required, Line{Point{r, t}, Point{r, b}, false})
	required = append(required, Line{Point{r, b}, Point{l, b}, false})
	required = append(required, Line{Point{l, b}, Point{l, t}, false})

	for _, line := range required {
		if !p.ContainsLine(line) {
			return false
		}
	}
	return true
}

// ContainsLine evaluates if the primitives contain a line matching
// line. It matches also when P1 and P2 are swapped.
func (p *Primitives) ContainsLine(line Line) bool {
	for _, sampleLine := range p.Lines {
		if sampleLine.Dashed != line.Dashed {
			continue
		}
		if sampleLine.P1.EqualIsh(line.P1) && sampleLine.P2.EqualIsh(line.P2) {
			return true
		}
		if sampleLine.P1.EqualIsh(line.P2) && sampleLine.P2.EqualIsh(line.P1) {
			return true
		}
	}
	return false
}

// ContainsLabel evaluates if the primitives contain a label matching
// label.
func (p *Primitives) ContainsLabel(label Label) bool {
	for _, x := range p.Labels {
		if x.EqualIsh(label) {
			return true
		}
	}
	return false
}

// BoundingBoxOfLines provides the bounding box extents of the lines
// present in the Primitives object.
func (p *Primitives) BoundingBoxOfLines() (left, top, right, bottom float64) {
	const huge = 1e12
	left = huge
	top = huge
	right = -huge
	bottom = -huge
	for _, line := range p.Lines {
		for _, pt := range []Point{line.P1, line.P2} {
			if pt.X < left {
				left = pt.X
			}
			if pt.Y < top {
				top = pt.Y
			}
			if pt.X > right {
				right = pt.X
			}
			if pt.Y > bottom {
				bottom = pt.Y
			}
		}
	}
	return
}
