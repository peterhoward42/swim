package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/geom"
	"github.com/stretchr/testify/assert"
)

/*
Given an ActivityBoxes instance...
When, its methods are used to create, and terminate two boxes...
Then, its AsSegments method should return two segments corresponding to
the values used for the AddStartingAt and TerminatAt methods calls.
*/
func TestHappyPath(t *testing.T) {
	assert := assert.New(t)

	boxes := NewActivityBoxes()
	err := boxes.AddStartingAt(25);
	assert.NoError(err)
	err = boxes.TerminateAt(60)
	assert.NoError(err)
	err = boxes.AddStartingAt(90)
	assert.NoError(err)
	err = boxes.TerminateAt(105)
	assert.NoError(err)

	segments := boxes.AsSegments()
	assert.Equal(geom.Segment{Start:25, End:60}, segments[0])
	assert.Equal(geom.Segment{Start:90, End:105}, segments[1])
}

/*
Given an ActivityBoxes instance with no boxes registered...
When, its TerminateAt method is called...
It should produce an error suitable to alert the developer.
*/
func TestErrorWhenThereIsNoBoxToTerminate(t *testing.T) {
	assert := assert.New(t)
	boxes := NewActivityBoxes()
	err := boxes.TerminateAt(105)
	assert.EqualError(err, "There is no box to terminate")
}

/*
Given an ActivityBoxes instance with one box registered that has already been
terminated...
When, its TerminateAt method is called...
Then it should produce an error suitable to alert the developer.
*/
func TestErrorWhenABoxIsTerminatedTwice(t *testing.T) {
	assert := assert.New(t)
	boxes := NewActivityBoxes()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.TerminateAt(60)
	assert.NoError(err)
	err = boxes.TerminateAt(105)
	assert.EqualError(err, "Cannot terminate an already-terminated box")
}

/*
Given an ActivityBoxes instance with one box registered that has not been
terminated...
When, its AddStartingAt method is called a second time...
Then it should produce an error suitable to alert the developer.
*/
func TestErrorWhenYouAddABoxWithoutTerminatingThePreviousOne(t *testing.T) {
	assert := assert.New(t)
	boxes := NewActivityBoxes()
	err := boxes.AddStartingAt(25)
	assert.NoError(err)
	err = boxes.AddStartingAt(50)
	assert.EqualError(err, "Cannot add new box when previous is not terminated")
}