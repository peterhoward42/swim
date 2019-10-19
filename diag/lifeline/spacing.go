package lifeline

import (
	"errors"
	"fmt"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/sizer"
)

/*
Spacing holds the knowledge about the horizontal pitch and geometry
of lifelines. For example, how to space them out across the page.

It makes all the title boxes the same width, and distributes these equally
across the width of the diagram. It uses the same gap (gutter) between these
boxes and as margins at the left and right edge of the diagram.
*/
type Spacing struct {
	sizer     sizer.Sizer
	lifelines []*dsl.Statement
	drivingValues drivingValues
}

// NewSpacing  provides a Spacing  ready to use.
func NewSpacing(sizer sizer.Sizer, lifelines []*dsl.Statement) *Spacing {
	spacer := &Spacing{
		sizer:     sizer,
		lifelines: lifelines,
	}
	spacer.setDrivingValues()
	return spacer
}

type TitleBoxXCoords struct {
	Left float64
	Centre float64
	Right float64
}

/*
CentreLine provides the X coordinate for the centreline of lifeline.
*/
func (s Spacing) CentreLine(lifeline *dsl.Statement) (*TitleBoxXCoords, error) {
	num, err := s.lifelineNumber(lifeline)
	if err != nil {
		return nil, fmt.Errorf("lifelineNumber: %v", err)
	}
	dv := s.drivingValues
	centre :=(float64(num)+1)*dv.titleBoxGutter + (float64(num)+0.5)*dv.titleBoxWidth
	delta := dv.titleBoxWidth / 2.0
	return &TitleBoxXCoords{centre - delta, centre, centre + delta}, nil
}

/*
The spacing algorithm is dependent on a few key driving values.
For example the chosen width of the lifeline title boxes, and the gap between
them. This function decides what they should be.
*/
type drivingValues struct {
	titleBoxWidth  float64
	titleBoxGutter float64
}

/*
setDrivingValues calculates the values that other spacing decisions are derived
from. They include trying to use an optimal looking width for lifeline title
boxes, but backtracking when this would make the gutter between the title boxes
too small and reducing the size of the title boxes such that a minimum gutter
of one font height is preserved.
*/
func (s *Spacing) setDrivingValues() {
	s.drivingValues.titleBoxWidth = s.sizer.Get("IdealLifelineTitleBoxWidth")
	n := len(s.lifelines)
	diagWidth := s.sizer.Get("DiagWidth")
	spaceAvail := diagWidth - s.drivingValues.titleBoxWidth*float64(n)
	nGuttersRequired := n + 1
	s.drivingValues.titleBoxGutter = spaceAvail / float64(nGuttersRequired)

	// But if that has that made the gutter too small, or even negative,
	// make the boxes less wide to preserve a minimum gutter equal to
	// one font height.
	fontHt := s.sizer.Get("FontHt")
	if s.drivingValues.titleBoxGutter < fontHt {
		s.drivingValues.titleBoxGutter = fontHt
		s.drivingValues.titleBoxWidth = diagWidth - float64(n+1)*s.drivingValues.titleBoxGutter/float64(n)
	}
}

func (s *Spacing) lifelineNumber(lifeline *dsl.Statement) (int, error) {
	for num, registered := range s.lifelines {
		if registered == lifeline {
			return num, nil
		}
	}
	return -1, errors.New("lifeline is not registered")
}
