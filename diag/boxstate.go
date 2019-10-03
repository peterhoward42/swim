package diag

/*
This module keeps track of the sequence of activity boxes on one lifeline.
It takes instructions to register the Y coordinate at which a box should
be registered as starting, and similarly instructions to later terminate a
box at a given (bigger) Y coordinate.
*/

import (
	"github.com/peterhoward42/umli/dsl"
)

// boxExtent keeps track of the Y coordinates at which a box starts and
// ends. When a box has been started, but the end coordinate is not yet
// known - it is said to be inProgress.
type boxExtent struct {
	extent     *segment
	inProgress bool
}

// Type lifelineBoxes keeps track of the activity boxes (for one lifeline)
// during diagram creation.
type lifelineBoxes struct {
	boxes []*boxExtent
}

// newlifelineBoxes provides a lifelineBoxes ready to use.
func newlifelineBoxes() *lifelineBoxes {
	return &lifelineBoxes{[]*boxExtent{}}
}

// inProgress returns true when the most recently started activity box
// on the lifeline has not yet been finished.
func (llb *lifelineBoxes) inProgress() bool {
	boxExtent := llb.mostRecent()
	if boxExtent == nil {
		return false
	}
	return boxExtent.inProgress
}

func (llb *lifelineBoxes) terminateInProgressBoxAt(y float64) {
	boxExtent := llb.mostRecent()
	if boxExtent == nil {
		return
	}
	boxExtent.extent.end = y
	boxExtent.inProgress = false
}

func (llb *lifelineBoxes) startBoxAt(y float64) {
	segment := &segment{y, -1}
	boxExtent := &boxExtent{segment, true}
	llb.boxes = append(llb.boxes, boxExtent)
}

// mostRecent returns the most recently added boxExtent for this lifeline,
// (or nil when none have been added.
func (llb *lifelineBoxes) mostRecent() *boxExtent {
	i := len(llb.boxes)
	if i == 0 {
		return nil
	}
	return llb.boxes[i-1]
}

// boxExtentsAsSegments provides a list of segments that represent the vertical
// space occupied by this lifeline's activity boxes.
func (llb *lifelineBoxes) boxExtentsAsSegments() []*segment {
	segs := []*segment{}
	for _, box := range llb.boxes {
		seg := box.extent
		segs = append(segs, seg)
	}
	return segs
}

// newBoxStates provides a lifelineBoxes ready to use.
func newLifelineBoxes(lifelineStatement *dsl.Statement) *lifelineBoxes {
	return &lifelineBoxes{[]*boxExtent{}}
}
