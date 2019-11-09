package lifeline

import (
	"github.com/peterhoward42/umli/graphics"
)

// ActivityBoxDrawer knows how to draw the lines required to represent the
// activity boxes on a lifeline - as specified by an ActivityBoxes object.
type ActivityBoxDrawer struct {
	activityBoxes ActivityBoxes
	centreX       float64
	boxWidth      float64
}

// NewActivityBoxDrawer provides an ActivityBoxDrawer ready to use.
func NewActivityBoxDrawer(activityBoxes ActivityBoxes, centreX float64,
	boxWidth float64) *ActivityBoxDrawer {
	return &ActivityBoxDrawer{
		activityBoxes: activityBoxes,
		centreX:       centreX,
		boxWidth:      boxWidth,
	}
}

// Draw creates the lines required, and add them to prims.
func (abc *ActivityBoxDrawer) Draw(prims *graphics.Primitives) {
	dx := 0.5 * abc.boxWidth
	for _, segment := range abc.activityBoxes.AsSegments() {
		left := abc.centreX - dx
		right := abc.centreX + dx
		top := segment.Start
		bottom := segment.End
		prims.AddRect(left, top, right, bottom)
	}
}
