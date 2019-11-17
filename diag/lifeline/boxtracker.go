package lifeline

import (
	"errors"

	"github.com/peterhoward42/umli/geom"
)

/*
BoxTracker keeps track of the vertical extents of the activity boxes
on a single lifeline. The nub of the problem it takes care of, is that you
don't know when an activity box should be closed off, at the time you have
to register where it should start.
*/
type BoxTracker struct {
	segs []geom.Segment // Used to track the start and end of each box.
}

// NewBoxTracker provides an an BoxTracker ready to use.
func NewBoxTracker() *BoxTracker {
	return &BoxTracker{}
}

// AddStartingAt registers a new box, with the Y coordinate at which it
// should start, but with the Y coordinate at which it should end - as yet
// unknown.
func (ab *BoxTracker) AddStartingAt(startY float64) error {
	if len(ab.segs) != 0 && ab.segs[len(ab.segs)-1].End == -1 {
		return errors.New("Cannot add new box when previous is not terminated")
	}
	ab.segs = append(ab.segs, geom.NewSegment(startY,-1))
	return nil
}

// TerminateAt finalises the Segment representing the most recently registered
// box, noting that it should end at endY.
func (ab *BoxTracker) TerminateAt(endY float64) error {
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
func (ab *BoxTracker) AsSegments() []geom.Segment {
	return ab.segs
}

/*
GetStartOfFinalBoxIfNotTerminated provides the y coordinate at which the
most recently added box starts at - but only when there is at least one
box, and the most recently added box has not yet been terminated. Otherwise
it returns nil.
*/
func (ab *BoxTracker) GetStartOfFinalBoxIfNotTerminated() *float64 {
	if len(ab.segs) == 0 {
		return nil
	}
	mostRecent := &ab.segs[len(ab.segs)-1]
	if mostRecent.End == -1 {
		return &mostRecent.Start
	}
	return nil
}

// HasABoxInProgress returns true if there are some boxes, and the most
// recently added one has not yet been terminated.
func (ab *BoxTracker) HasABoxInProgress() bool {
	if len(ab.segs) == 0 {
		return false
	}
	mostRecent := &ab.segs[len(ab.segs)-1]
	if mostRecent.End == -1 {
		return true
	}
	return false
}
