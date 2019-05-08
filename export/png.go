package export

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/peterhoward42/umlinteraction/graphics"
)

// CreatePNG renders a graphics model into a .png image.
func CreatePNG(filePath string, graphicsModel *graphics.Model) error {
	mdl := graphicsModel
	dc := gg.NewContext(mdl.Width, mdl.Height)
	dc.SetRGB(100,255,255)
	dc.DrawRectangle(0, 0, float64(mdl.Width), float64(mdl.Height))
	dc.Fill()
	dc.SetRGB(0,0,0)
	for _, line := range mdl.Primitives.Lines {
		dc.DrawLine(line.X1, line.Y1, line.X2, line.Y2)
		dc.Stroke()
	}
	err := dc.SavePNG(filePath)
	if err != nil {
		return fmt.Errorf("CreatePNG: %v", err)
	}
	return nil
}