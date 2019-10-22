package lifeline

import (
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
	"github.com/peterhoward42/umli/diag/nogozone"
)

/*
Gaps represents the gaps that are necessary when drawing a lifeline
to avoid drawing over the top of something else.
*/
type Gaps struct {
	Items []geom.Segment
}

/*
PopulateFromNoGoZones assembles the gaps that are needed in lifeline, to
respect noGoZones.
*/
func (gaps *Gaps) PopulateFromNoGoZones(noGoZones []nogozone.NoGoZone,
	lifeline *dsl.Statement, allLifelines []*dsl.Statement) {
	gaps.Items = []geom.Segment{}
	for _, noGoZone := range noGoZones {
		// Does this noGoZone affect lifeline?
		affectedLifelines := SpanExcl(noGoZone.OneEndLifeline,
			noGoZone.OtherEndLifeline,
			allLifelines)
		if gaps.lifelineIsAmong(affectedLifelines, lifeline) {
			gaps.Items = append(gaps.Items, noGoZone.Height)
		}
	}
}

// lifelineIsAmong returns true if aLifeline is in someLifelines.
func (gaps *Gaps) lifelineIsAmong(
	someLifelines []*dsl.Statement, aLifeline *dsl.Statement) bool {
	for _, q := range someLifelines {
		if q == aLifeline {
			return true
		}
	}
	return false
}
