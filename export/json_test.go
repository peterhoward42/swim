package export

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peterhoward42/umlinteraction/graphics"
)

func TestJSONOutputIsAsExpected(t *testing.T) {
	assert := assert.New(t)
	mdl := graphics.NewModel(200, 3)
	mdl = append(mdl.Lines, &graphics.Line{
		X1: 0, Y1: 0, X2: 100, Y2: 100, Dashed: false, Arrow: false})
	mdl.Labels = append(mdl.Labels, &graphics.Label{
		LinesOfText: []string{"foo", "bar"},
		X:           42,
		Y:           43,
		HJust:       graphics.Left,
		VJust:       graphics.Centre,
	})
	theJSON, _ := SerializeToJSON(mdl)
	assert.JSONEq(expectedJSON, string(theJSON))
}

const expectedJSON = `{
	"Width": 200,
	"FontHeight": 3,
	"Lines": [
	  {
		"X1": 0,
		"X2": 100,
		"Y1": 0,
		"Y2": 100,
		"Arrow": false,
		"Dashed": false
	  }
	],
	"Labels": [
	  {
		"LinesOfText": [
		  "foo",
		  "bar"
		],
		"X": 42,
		"Y": 43,
		"HJust": "Left",
		"VJust": "Centre"
	  }
	]
  }
  `
