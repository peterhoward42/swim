package lifeline

import (
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/sizer"
)

/*
TitleBox knows how to calculate the height of a lifeline's title box.
*/
type TitleBox struct {
	lifeline *dsl.Statement
}

// NewTitleBox creates a TitleBox ready to use.
func NewTitleBox(lifeline *dsl.Statement) *TitleBox {
	return &TitleBox{
		lifeline: lifeline,
	}
}

// Height provides the height of the title box.
func (tb *TitleBox) Height(sizer sizer.Sizer) float64 {
	n := len(tb.lifeline.LabelSegments)
	topMargin := sizer.Get("TitleBoxLabelPadT")
	botMargin := sizer.Get("TitleBoxLabelPadB")
	ht := topMargin + botMargin + float64(n)*sizer.Get("FontHeight")
	return ht
}
