package lifeline

import (
	"errors"

	"github.com/peterhoward42/umli/geom"
)

/*
ActivityBoxes keeps track of the vertical extents of the activity boxes
on a single lifeline. The nub of the problem it takes care of, is that you
don't know when an activity box should be closed off, at the time you have
to register where it should start.
*/
type ActivityBoxes struct {
	segs []geom.Segment // Used to track the box start and end Y coordinate.
}

// NewActivityBoxes provides an an ActivityBoxes ready to use.
func NewActivityBoxes() *ActivityBoxes {
	return &ActivityBoxes{}
}

// AddStartingAt registers a new box, with the Y coordinate at which it
// should start, but with the Y coordinate at which it should end - as yet
// unknown.
func (ab *ActivityBoxes) AddStartingAt(startY float64) error {
	if len(ab.segs) != 0 && ab.segs[len(ab.segs) -1].End == -1 {
		return errors.New("Cannot add new box when previous is not terminated")
	}
	ab.segs = append(ab.segs, geom.Segment{Start: startY, End:-1})
	return nil
}

// TerminateAt finalises the Segment representing the most recently registered
// box, noting that it should end at endY.
func (ab *ActivityBoxes) TerminateAt(endY float64) error {
	if len(ab.segs) == 0 {
		return errors.New("There is no box to terminate")
	}
	mostRecent := &ab.segs[len(ab.segs)-1]
	if mostRecent.End != -1 {
		return errors.New("Cannot terminate an already-terminated box")
	}
	mostRecent.End = endY
	return nil
}

// AsSegments provides all the boxes that have been registered.
func (ab *ActivityBoxes) AsSegments() []geom.Segment {
	return ab.segs
}
