package interactions

import (
	"github.com/peterhoward42/umli/graphics"
)

/*
LabelPosn knows how to work out the position for an interaction line's label
given the X coordinates (from, to) between which the interaction line travels.
The interface is designed to be able to cope with the answer being different
for left-to-right interaction lines vs right-to-left ones.
*/
type LabelPosn struct {
	from float64
	to   float64
}

// NewLabelPosn provides a LabelPosn ready to use.
func NewLabelPosn(from, to float64) *LabelPosn {
	return &LabelPosn{from, to}
}

func (lp *LabelPosn) Get() (x float64, horizontalJustification graphics.Justification) {
	// So far it looks better if it's centred regardless of direction.
	return 0.5 * (lp.from + lp.to), graphics.Centre
}
