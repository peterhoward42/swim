package export

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peterhoward42/umlinteraction/graphics"
)

func TestJSONOutputIsAsExpected(t *testing.T) {
	assert := assert.New(t)
	mdl := graphics.NewModel(200, 3)
	p := mdl.Primitives
	p.AddLine(0, 0, 100, 100, false, false)
	p.AddLabel([]string{"foo"}, 3, 4, graphics.Right, graphics.Centre)
	theJSON, _ := SerializeToJSON(mdl)
	//fmt.Print(string(theJSON))

	assert.JSONEq(expectedJSON, string(theJSON))
}

var expectedJSON = strings.TrimSpace(`
{
  "Width": 200,
  "Height": 2000,
  "FontHeight": 3,
  "Primitives": {
    "Lines": [
      {
        "X1": 0,
        "X2": 0,
        "Y1": 100,
        "Y2": 100,
        "Arrow": false,
        "Dashed": false
      }
    ],
    "Labels": [
      {
        "LinesOfText": [
          "foo"
        ],
        "X": 3,
        "Y": 4,
        "HJust": "Right",
        "VJust": "Centre"
      }
    ]
  }
}
`)
