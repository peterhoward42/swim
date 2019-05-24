package diag

import (
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizers"
)

// makeArrow assembles the vertices required to make a horizontal! 
// arrow head polygon that has its tip at (x2, y), and points in 
// the x1->x2 direction.
func makeArrow(x1 float64, x2 float64, y float64,
	sizer *sizers.Sizer) []*graphics.Point {
	dx := sizer.ArrowLen
	dy := 0.5 * sizer.ArrowHeight
	var p1, p2, p3 *graphics.Point
	if x2 > x1 {
		p1 = graphics.NewPoint(x2 - dx, y - dy)
		p2 = graphics.NewPoint(x2, y)
		p3 = graphics.NewPoint(x2 - dx, y + dy)
	} else {
		p1 = graphics.NewPoint(x2 + dx, y - dy)
		p2 = graphics.NewPoint(x2, y)
		p3 = graphics.NewPoint(x2 + dx, y + dy)
	}
	return []*graphics.Point{p1, p2, p3}
}
