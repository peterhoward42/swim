package interactions

import (
	"fmt"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
)

/*
Maker knows how to make the interaction lines.
*/
type Maker struct {
	dependencies  *MakerDependencies
	graphicsModel *graphics.Model
}

/*
MakerDependencies encapsulate the prior state of diagram creation process, and all
the things the Maker needs from the outside to do its job.
*/
type MakerDependencies struct {
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
Make goes through statements in order, and works out what graphics are
required to represent interaction lines, activitiy boxes etc. It advances
the tidemark as it goes, and returns the final resultant tidemark.
*/
func (mkr *Maker) Scan(
	tidemark float64,
	statements []*dsl.Statement) (newTidemark float64, err error) {

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
	var prevTm float64 = tidemark
	var newTm float64
	for _, action := range actions {
		newTm, err := action.fn(prevTm, action.statement)
		if err != nil {
			return -1, fmt.Errorf("actionFn: %v", err)
		}
		prevTm = newTm
	}
	return newTm, nil
}

func (mkr *Maker) interactionLabel(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	return -1, nil
}

func (mkr *Maker) startToBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	return -1, nil
}

func (mkr *Maker) startFromBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	return -1, nil
}

func (mkr *Maker) interactionLine(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	return -1, nil
}

func (mkr *Maker) endBox(
	tidemark float64, s *dsl.Statement) (newTidemark float64, err error) {
	return -1, nil
}

type actionFn func(
	tideMark float64,
	s *dsl.Statement) (newTidemark float64, err error)

type dispatch struct {
	fn        actionFn
	statement *dsl.Statement
}
