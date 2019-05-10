package export

import (
	"os"
	"path/filepath"
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/imgeq"
)

func TestErrorMsgIsCorrectWhenCantOpenFile(t *testing.T) {
	assert := assert.New(t)
	width := 200
	fontHeight := 3.5
	graphicsModel := graphics.NewModel(width, fontHeight)
	graphicsModel.Primitives.AddRect(0, 0, 100, 100)
	invalidFilePath := ""
	err := CreatePNG(invalidFilePath, graphicsModel)
	assert.EqualError(err, "CreatePNG: open : no such file or directory")
}
func TestDotPNGOutputProducesAFile(t *testing.T) {
	assert := assert.New(t)
	width := 200
	fontHeight := 3.5
	graphicsModel := graphics.NewModel(width, fontHeight)
	graphicsModel.Primitives.AddRect(0, 0, 100, 100)
	filePath := filepath.Join(os.TempDir(), "umli.png")
	err := CreatePNG(filePath, graphicsModel)
	assert.NoError(err)
	assert.FileExists(filePath)
}

// Make a .png file for a graphics model containing one of each
// primitive and makes sure it is idential to a stored golden reference
// image file.
func TestRegressionWithEveryGraphicsType(t *testing.T) {
	assert := assert.New(t)
	width := 1000
	fontHeight := 33.0
	graphicsModel := graphics.NewModel(width, fontHeight)

	prims := graphicsModel.Primitives
	prims.AddLine(50, 50, 250, 50, false, false)
	prims.AddLine(50, 150, 250, 150, true, false)
	prims.AddLine(50, 250, 250, 250, false, true)

  basePath := "umli_test_regress.png"
	generated := filepath.Join(os.TempDir(), basePath)
	//fmt.Printf("Generated file: %s", generated)
  goldenRef := filepath.Join(".", "test_data", basePath)
	err := CreatePNG(generated, graphicsModel)
  assert.NoError(err)
  
  areEqual, err := imgeq.AreEqual(goldenRef, generated)
	assert.NoError(err)
  assert.True(areEqual, "File produced is not identical to reference file")
}
