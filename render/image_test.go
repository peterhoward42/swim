package render

import (
	"path/filepath"
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/peterhoward42/umli/graphics"
)

var testResultsDir = filepath.Join(".", "testresults", "new")

// fullCoverageModel is a DRY test helper that makes a graphics model
// with one example of every graphics primitive in.
func fullCoverageModel() *graphics.Model {
	width := 2000.0
	height := 1000.0
	fh := 45.0
	dashLineDashLen := 45.0
	dashLineGapLen := 20.0
	graphicsModel := graphics.NewModel(width, height, fh,
		dashLineDashLen, dashLineGapLen)
	prims := graphicsModel.Primitives

	// We draw two solid horizontal lines separated by the font height.
	// Plus dashed cross-hairs to show the horiz/vert mid point.
	// Then render 3 strings, using that font height that should position
	// the strings snugly between the two solid lines.
	// The 3 strings are left, centre, right justified w.r.t. the left, centre
	// and right of the lines, and should thus line up with the line ends and
	// centre horizontally.
	// Each uses a different vertical justification (and anchor) which should
	// leave them all ostensibly fitting snugly between the lines.

	// Notably they are identically positioned in Y in the output,
	// and all look identically biased towards the bottom line. Suggesting a
	// font height (by definition) includes some padding above the glyphs.
	// Note also some characters like 'p' have descenders below the bottom line.

	left := 100.0
	right := 1000.0
	top := 100.0
	bot := top + fh
	midX := 0.5 * (left + right)
	midY := 0.5 * (top + bot)

	prims.AddLine(left, top, right, top, false)
	prims.AddLine(left, bot, right, bot, false)
	prims.AddLine(left, midY, right, midY, true)
	prims.AddLine(midX, top, midX, bot, true)

	polyRight := right + fh
	prims.AddFilledPoly([]*graphics.Point{
		&graphics.Point{X: right, Y: top},
		&graphics.Point{X: polyRight, Y: bot},
		&graphics.Point{X: right, Y: bot},
	})

	prims.AddLabel("LeftBot", fh, left, bot, graphics.Left, graphics.Bottom)
	prims.AddLabel("CtrCtr", fh, midX, midY, graphics.Centre, graphics.Centre)
	prims.AddLabel("RightTop", fh, right, top, graphics.Right, graphics.Top)

	return graphicsModel
}

var fileExtensions = map[Encoding]string{
	PNG: ".png",
	JPG: "jpg",
}

// createAndsaveImageFileForExampleModel is a DRY test helper that first
// builds an example graphics model, and then saves it according the
// parameters passed in.
func createAndsaveImageFileForExampleModel(t *testing.T, encoding Encoding,
	fileBaseName string) (savedFilePath string) {
	assert := assert.New(t)
	graphicsModel := fullCoverageModel()
	fileExtension := fileExtensions[encoding]
	saveAs := filepath.Join(testResultsDir, fileBaseName+fileExtension)
	font, err := truetype.Parse(goregular.TTF)
	assert.NoError(err)
	err = NewImageFileCreator(font).Create(saveAs, encoding, graphicsModel)
	assert.NoError(err)
	return saveAs
}

func TestThatSavesExampleModelAsPNGForVisualInspection(t *testing.T) {
	createAndsaveImageFileForExampleModel(t, PNG, "example")
}

func TestThatSavesExampleModelAsJPGForVisualInspection(t *testing.T) {
	createAndsaveImageFileForExampleModel(t, JPG, "example")
}
