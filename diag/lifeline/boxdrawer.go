package lifeline

import (
	"github.com/peterhoward42/umli/graphics"
)

// BoxDrawer knows how to draw the lines required to represent the
// activity boxes on a lifeline - as specified by an BoxTracker object.
type BoxDrawer struct {
	boxes BoxTracker
	centreX       float64
	boxWidth      float64
}

// NewBoxDrawer provides an BoxDrawer ready to use.
func NewBoxDrawer(boxes BoxTracker, centreX float64,
	boxWidth float64) *BoxDrawer {
	return &BoxDrawer{
		boxes: boxes,
		centreX:       centreX,
		boxWidth:      boxWidth,
	}
}

// Draw creates the lines required, and add them to prims.
func (abc *BoxDrawer) Draw(prims *graphics.Primitives) {
	dx := 0.5 * abc.boxWidth
	for _, segment := range abc.boxes.AsSegments() {
		left := abc.centreX - dx
		right := abc.centreX + dx
		top := segment.Start
		bottom := segment.End
		prims.AddRect(left, top, right, bottom)
	}
}
