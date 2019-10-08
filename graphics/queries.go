package graphics

/*
This module provides some queries that can be made against the model's
contents. For example does it contain the four lines required to represent
a given rectangle?

The queries are optimised for code simplicity (mainly for tests) at the
expense of Big-O complexity.
*/

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
	for _, x := range p.Lines {
		if x.P1 == line.P1 && x.P2 == line.P2 && x.Dashed == line.Dashed {
			return true
		}
		if x.P1 == line.P2 && x.P2 == line.P1 && x.Dashed == line.Dashed {
			return true
		}
	}
	return false
}
