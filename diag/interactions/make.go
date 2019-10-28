package interactions

import (
	"fmt"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/diag/lifeline"
	"github.com/peterhoward42/umli/diag/nogozone"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
)

/*
Maker knows how to make the interaction lines and when to start/stop
their activity boxes.
*/
type Maker struct {
	dependencies  *MakerDependencies
	graphicsModel *graphics.Model
}

/*
MakerDependencies encapsulates the prior state of the diagram creation
process at the time that the Make method is called. And includes all
the things the Maker needs from the outside to do its job.
*/
type MakerDependencies struct {
	activityBoxes map[*dsl.Statement]*lifeline.ActivityBoxes
	fontHt        float64
	noGoZones     []*nogozone.NoGoZone
	sizer         sizer.Sizer
	spacer        *lifeline.Spacing
}

// NewMakerDependencies makes a MakerDependencies ready to use.
func NewMakerDependencies(fontHt float64, spacer *lifeline.Spacing,
	sizer sizer.Sizer,
	activityBoxes map[*dsl.Statement]*lifeline.ActivityBoxes,
	noGoZones []*nogozone.NoGoZone) *MakerDependencies {
	return &MakerDependencies{
		activityBoxes: activityBoxes,
		fontHt:        fontHt,
		noGoZones:     noGoZones,
		sizer:         sizer,
		spacer:        spacer,
	}
}

/*
NewMaker initialises a Maker ready to use.
*/
func NewMaker(d *MakerDependencies, gm *graphics.Model) *Maker {
	return &Maker{
		dependencies:  d,
		graphicsModel: gm,
	}
}

/*
ScanInteractionStatements goes through the DSL statements in order, and
works out what graphics are required to represent interaction lines, and
activitiy boxes etc. It advances the tidemark as it goes, and returns the
final resultant tidemark.
*/
func (mkr *Maker) ScanInteractionStatements(
	tidemark float64,
	statements []*dsl.Statement) (newTidemark float64, err error) {

	// Build a list of actions to execute depending on the statement
	// keyword.
	actions := []dispatch{}
	for _, s := range statements {
		switch s.Keyword {
		case umli.Dash:
			actions = append(actions, dispatch{mkr.interactionLabel, s})
			actions = append(actions, dispatch{mkr.startToBox, s})
			actions = append(actions, dispatch{mkr.interactionLine, s})
		case umli.Full:
			actions = append(actions, dispatch{mkr.interactionLabel, s})
			actions = append(actions, dispatch{mkr.startFromBox, s})
			actions = append(actions, dispatch{mkr.startToBox, s})
			actions = append(actions, dispatch{mkr.interactionLine, s})
		case umli.Self:
			actions = append(actions, dispatch{mkr.startFromBox, s})
			actions = append(actions, dispatch{mkr.interactionLine, s})
		case umli.Stop:
			actions = append(actions, dispatch{mkr.endBox, s})
		}
	}
	var prevTidemark float64 = tidemark
	var updatedTidemark float64
	for _, action := range actions {
		updatedTidemark, err = action.fn(prevTidemark, action.statement)
		if err != nil {
			return -1, fmt.Errorf("actionFn: %v", err)
		}
		prevTidemark = updatedTidemark
	}
	return updatedTidemark, nil
}

// interactionLabel creates the graphics label that belongs to an interaction
// line.
func (mkr *Maker) interactionLabel(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	dep := mkr.dependencies
	sourceLifeline := s.ReferencedLifelines[0]
	destLifeline := s.ReferencedLifelines[1]
	fromX, toX, err := mkr.LifelineCentres(sourceLifeline, destLifeline)
	labelX, horizJustification := NewLabelPosn(fromX, toX).Get()
	mkr.graphicsModel.Primitives.RowOfStrings(
		labelX, tidemark, dep.fontHt, horizJustification, s.LabelSegments)
	newTidemark = tidemark + float64(len(s.LabelSegments))*
		dep.fontHt + dep.sizer.Get("InteractionLineTextPadB")
	noGoZone := nogozone.NewNoGoZone(
		geom.Segment{Start: tidemark, End: newTidemark},
		sourceLifeline, destLifeline)
	dep.noGoZones = append(dep.noGoZones, &noGoZone)
	return newTidemark, nil
}

// interactionLine makes an interaction line (and its arrow)
func (mkr *Maker) interactionLine(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	dep := mkr.dependencies
	sourceLifeline := s.ReferencedLifelines[0]
	destLifeline := s.ReferencedLifelines[1]
	fromX, toX, err := mkr.LifelineCentres(sourceLifeline, destLifeline)
	halfActivityBoxWidth := 0.5 * dep.sizer.Get("ActivityBoxWidth")
	geom.ShortenLineBy(halfActivityBoxWidth, &fromX, &toX)
	y := tidemark
	dashed := s.Keyword == umli.Dash
	mkr.graphicsModel.Primitives.AddLine(fromX, y, toX, y, dashed)
	arrowLen := dep.sizer.Get("ArrowLen")
	arrowWidth := dep.sizer.Get("ArrowWidth")
	arrow := geom.MakeArrow(fromX, toX, y, arrowLen, arrowWidth)
	mkr.graphicsModel.Primitives.AddFilledPoly(arrow)
	newTidemark = tidemark + dep.sizer.Get("InteractionLinePadB")
	noGoZone := nogozone.NewNoGoZone(
		geom.Segment{Start: tidemark, End: newTidemark},
		sourceLifeline, destLifeline)
	dep.noGoZones = append(dep.noGoZones, &noGoZone)
	return newTidemark, nil
}

// startToBox registers with a lifeline.ActivityBoxes that an activity box
// on a lifeline should be started ready for an interaction line to arrive at
// the top of it. (If a box is not already in progress for this lifeline.)
func (mkr *Maker) startToBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	dep := mkr.dependencies
	toLifeline := s.ReferencedLifelines[1]
	activityBoxes := dep.activityBoxes[toLifeline]
	if activityBoxes.HasABoxInProgress() {
		return tidemark, nil
	}
	activityBoxes.AddStartingAt(tidemark)
	// Return an unchanged tidemark.
	return tidemark, nil
}

// starFromBox registers that an activity box on a lifeline
// should be started (unless it is already) for an activity line to emenate from.
func (mkr *Maker) startFromBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	dep := mkr.dependencies
	fromLifeline := s.ReferencedLifelines[0]
	activityBoxes := dep.activityBoxes[fromLifeline]
	if activityBoxes.HasABoxInProgress() {
		return tidemark, nil
	}
	// The activity box should start just a tiny bit before the first
	// interaction line leaving from it. This need not claim any vertical
	// space of its own however, because the space already claimed by the interaction
	// line label is sufficient.
	backTrackToStart := dep.sizer.Get("ActivityBoxVerticalOverlap")
	activityBoxes.AddStartingAt(tidemark - backTrackToStart)
	// Return an unchanged tidemark.
	return tidemark, nil
}

// endBox processes an explicit "stop" statement.
func (mkr *Maker) endBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	dep := mkr.dependencies
	fromLifeline := s.ReferencedLifelines[0]
	activityBoxes := dep.activityBoxes[fromLifeline]
	tidemark += dep.sizer.Get("wontbeoverlapatendofbox")
	err = activityBoxes.TerminateAt(tidemark)
	if err != nil {
		return -1, fmt.Errorf("activityBoxes.TerminateAt: %v", err)
	}
	tidemark += dep.sizer.Get("wontbeoverendboxdaylight")
	return tidemark, nil
}

// actionFn describes a function that can be called to draw something
// related to statement s. Such a function receives the current tidemark,
// calculates how it should be advanced, and return the updated value.
type actionFn func(
	tideMark float64,
	s *dsl.Statement) (newTidemark float64, err error)

// dispatch is a simple container to hold a binding between  an actionFn and
// the statement to which it refers.
type dispatch struct {
	fn        actionFn
	statement *dsl.Statement
}

/*
LifelineCentres evaluates the X coordinates for the lifelines between which
an interaction line travels.
*/
func (mkr *Maker) LifelineCentres(
	sourceLifeline, destLifeline *dsl.Statement) (fromX, toX float64, err error) {
	fromCoords, err := mkr.dependencies.spacer.CentreLine(sourceLifeline)
	if err != nil {
		return -1.0, -1.0, fmt.Errorf("space.CentreLine: %v", err)
	}
	toCoords, err := mkr.dependencies.spacer.CentreLine(destLifeline)
	if err != nil {
		return -1.0, -1.0, fmt.Errorf("space.CentreLine: %v", err)
	}
	return fromCoords.Centre, toCoords.Centre, nil
}
