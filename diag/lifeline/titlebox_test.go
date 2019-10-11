package lifeline

import (
	"testing"

	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/sizer"
	"github.com/stretchr/testify/assert"
)

func TestHeightCalculation(t *testing.T) {
	assert := assert.New(t)
	sizer := sizer.NewLiteralSizer(map[string]float64{
		"FontHeight": 10.0,
		"TitleBoxLabelPadB": 5.0,
		"TitleBoxLabelPadT": 3.0,
	})
	lifeline := &dsl.Statement{
		LabelSegments: []string{"foo", "bar"},
	}
	titleBox := NewTitleBox(lifeline)
	height := titleBox.Height(sizer)
	assert.Equal(28.0, height)
}
