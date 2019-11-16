package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/geom"
	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	assert := assert.New(t)

	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(60)
	assert.NoError(err)
	err = boxes.AddStartingAt(90)
	assert.NoError(err)
	err = boxes.TerminateAt(105)
	assert.NoError(err)

	segments := boxes.AsSegments()
	assert.Equal(geom.Segment{Start: 25, End: 60}, segments[0])
	assert.Equal(geom.Segment{Start: 90, End: 105}, segments[1])
}

func TestErrorWhenThereIsNoBoxToTerminate(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.TerminateAt(105)
	assert.EqualError(err, "There is no box to terminate")
}

func TestCannotTerminateAnAlreadyTerminatedBox(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(60)
	assert.NoError(err)
	err = boxes.TerminateAt(105)
	assert.EqualError(err, "Cannot terminate an already-terminated box")
}

func TestCannotAddNewBoxWhenPreviousIsNotTerminated(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.AddStartingAt(50)
	assert.EqualError(err, "Cannot add new box when previous is not terminated")
}

func TestGetStartOfFinalBoxIfNotTerminatedHappyPath(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(30)
	assert.NoError(err)
	err = boxes.AddStartingAt(50)
	assert.NoError(err)
	finalBoxStart := boxes.GetStartOfFinalBoxIfNotTerminated()
	assert.Equal(float64(50), *finalBoxStart)
}

func TestHasABoxInProgressWhenTrue(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	assert.True(boxes.HasABoxInProgress())
}

func TestHasABoxInProgressWhenThereAreNoBoxes(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	assert.False(boxes.HasABoxInProgress())
}

func TestHasABoxInProgressWhenFinalBoxIsTerminated(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(30)
	assert.NoError(err)
	assert.False(boxes.HasABoxInProgress())
}

func TestGetStartOfFinalBoxIfNotTerminatedWhenIsAlreadyTerminated(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(30)
	assert.NoError(err)
	finalBoxStart := boxes.GetStartOfFinalBoxIfNotTerminated()
	assert.Nil(finalBoxStart)
}

func TestGetStartOfFinalBoxIfNotTerminatedWhenThereAreNoBoxes(t *testing.T) {
	assert := assert.New(t)
	boxes := NewBoxTracker()
	finalBoxStart := boxes.GetStartOfFinalBoxIfNotTerminated()
	assert.Nil(finalBoxStart)
}
