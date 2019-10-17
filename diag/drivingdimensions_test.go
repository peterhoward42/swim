package diag

import (
	"testing"

	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dsl"
	"github.com/stretchr/testify/assert"
)

func TestCorrectResultsWhenNoTextSizeStatement(t *testing.T) {
	assert := assert.New(t)
	mdl := dsl.Model{}
	diagWidth, fontHeight := DrivingDimensions{}.WidthAndFontHeight(mdl)
	assert.Equal(diagWidth, 2000.0)
	assert.Equal(fontHeight, 20.0)
}

func TestCorrectResultsWhenTextSizeStatementPresent(t *testing.T) {
	assert := assert.New(t)
	mdl := dsl.Model{}
	s := dsl.NewStatement()
	s.Keyword = umli.TextSize
	s.TextSize = 20.0
	mdl.Append(s)
	diagWidth, fontHeight := DrivingDimensions{}.WidthAndFontHeight(mdl)
	assert.Equal(diagWidth, 2000.0)
	assert.Equal(fontHeight, 40.0)
}
