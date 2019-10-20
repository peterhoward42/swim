package interactions

import (
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
func (mkr *Maker) Scan(statements []*dsl.Statement) (newTidemark float64, err error) {
	for _, _ = range statements {
		/*
			switch s.Keyword {
			case umli.Dash:
				sc.interactionLabel(s)
				sc.potentiallyStartToBox(s)
				sc.interactionLine(s)
			case umli.Full:
				sc.interactionLabel(s)
				sc.potentiallyStartFromBox(s)
				sc.potentiallyStartToBox(s)
				sc.interactionLine(s)
			case umli.Self:
				sc.potentiallyStartFromBox(s)
				sc.selfInteractionLines(s)
			case umli.Stop:
				sc.endBox(s)
			}
		*/
	}
	return -1.0, nil
}
