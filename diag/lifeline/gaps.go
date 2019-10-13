package lifeline

import (
	"github.com/peterhoward42/umli/diag/interactions"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
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
func (gaps *Gaps) PopulateFromNoGoZones(noGoZones []interactions.NoGoZone,
	lifeline *dsl.Statement, allLifelines []*dsl.Statement) {
	gaps.Items = []geom.Segment{}
	for _, noGoZone := range noGoZones {
		// Does this noGoZone affect lifeline?
		affectedLifelines := SpanExcl(noGoZone.OneEnd, noGoZone.TheOtherEnd, 
		allLifelines)
		if gaps.lifelineIsAmong(affectedLifelines, lifeline) {
			gaps.Items = append(gaps.Items, noGoZone.Height)
		}
	}
}

/*
PopulateFrom

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

