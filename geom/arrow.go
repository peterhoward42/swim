package geom

import (
	"github.com/peterhoward42/umli/graphics"
)

// makeArrow assembles the vertices required to make a horizontal!
// arrow head polygon that has its tip at (x2, y), and points in
// the x1->x2 direction.
func MakeArrow(x1 float64, x2 float64, y float64,
	arrowLen float64, arrowHeight float64) []graphics.Point {
	dx := arrowLen
	dy := 0.5 * arrowHeight
	var p1, p2, p3 graphics.Point
	if x2 > x1 {
		p1 = graphics.Point{X: x2 - dx, Y: y - dy}
		p2 = graphics.Point{X: x2, Y: y}
		p3 = graphics.Point{X: x2 - dx, Y: y + dy}
	} else {
		p1 = graphics.Point{X: x2 + dx, Y: y - dy}
		p2 = graphics.Point{X: x2, Y: y}
		p3 = graphics.Point{X: x2 + dx, Y: y + dy}
	}
	return []graphics.Point{p1, p2, p3}
}
