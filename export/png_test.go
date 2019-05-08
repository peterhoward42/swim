package export

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/udhos/equalfile"

	"github.com/peterhoward42/umlinteraction/graphics"
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
func TestRegressionWithEveryGraphicsType(t *testing.T) {
	assert := assert.New(t)
	width := 1000
	fontHeight := 33.0
	graphicsModel := graphics.NewModel(width, fontHeight)

	prims := graphicsModel.Primitives
	prims.AddLine(50, 50, 250, 50, false, false)
	prims.AddLine(50, 150, 250, 150, true, false)
	prims.AddLine(50, 250, 250, 50, false, true)

  basePath := "umli_test_regress.png"
  newFilePath := filepath.Join(os.TempDir(), basePath)
  //referencePath := filepath.Join(".", "test_data", basePath)
	err := CreatePNG(newFilePath, graphicsModel)
	assert.NoError(err)
  // Compare to reference file.
  /*
	cmp := equalfile.New(nil, equalfile.Options{})
	equal, err := cmp.CompareFile(newFilePath, referencePath)
	assert.NoError(err)
  assert.True(equal, "File produced is not identical to reference file")
  */
}
