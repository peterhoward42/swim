package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/diag/interactions"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/geom"
	"github.com/stretchr/testify/assert"
)

func TestGetCorrectGapsWhenLifelineIsAffectedByAllNoGoZones(t *testing.T) {
	assert := assert.New(t)

	a := &dsl.Statement{}
	b := &dsl.Statement{}
	c := &dsl.Statement{}
	allLifelines := []*dsl.Statement{a, b, c}

	seg12 := geom.Segment{Start: 1, End: 2}
	seg56 := geom.Segment{Start: 5, End: 6}

	// Use two NoGozone(s) between lifelines a and c.
	nogozones := []interactions.NoGoZone{
		interactions.NewNoGoZone(seg12, a, c),
		interactions.NewNoGoZone(seg56, a, c),
	}

	// Lifeline b is affected by these NoGoZone(s) because b lies in between
	// a and c, so calling PopulateFromNoGoZones should produce gaps that match
	// all the NoGoZones.

	gaps := Gaps{}
	gaps.PopulateFromNoGoZones(nogozones, b, allLifelines)

	segs := gaps.Items
	assert.Len(segs, 2)
	assert.Equal(seg12, segs[0])
	assert.Equal(seg56, segs[1])
}

func TestGetZeroGapsWhenLifelineIsAffectedByNoneOfTheNoGoZones(t *testing.T) {
	assert := assert.New(t)

	a := &dsl.Statement{}
	b := &dsl.Statement{}
	c := &dsl.Statement{}
	allLifelines := []*dsl.Statement{a, b, c}

	seg12 := geom.Segment{Start: 1, End: 2}
	seg56 := geom.Segment{Start: 5, End: 6}

	// Use two NoGozone(s) between lifelines a and c.
	nogozones := []interactions.NoGoZone{
		interactions.NewNoGoZone(seg12, a, c),
		interactions.NewNoGoZone(seg56, a, c),
	}

	// Lifeline c is not affected by these NoGoZone(s) because c does not lie
	// between a and c, so calling PopulateFromNoGoZones should produce no gaps.

	gaps := Gaps{}
	gaps.PopulateFromNoGoZones(nogozones, c, allLifelines)

	segs := gaps.Items
	assert.Len(segs, 0)
}

// make into table test
// make sure get no segs when lifeline unaffected by those registered

// make sure hybrid case works
