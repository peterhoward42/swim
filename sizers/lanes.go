package sizers

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
	umli "github.com/peterhoward42/umlinteraction"
)

type Lanes struct {
	NumLanes int
	TitleBoxWidth int
	TitleBoxHeight int
	TitleBoxHorizGap int
	InfoPerLane InfoPerLane
}

type InfoPerLane map[dslmodel.Statement]*LaneInfo

type LaneInfo struct {
	Left int
	Centre int
	Right int
}
