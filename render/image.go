package render

/*
This module provides the ImageFileCreator type and its methods.
*/

import (
	"fmt"

	"golang.org/x/image/colornames"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"

	"github.com/peterhoward42/umli/graphics"
)

// Encoding specifies image file encoding formats
type Encoding int

// Values for type Encoding
const (
	PNG Encoding = iota
	JPG
)

// ImageFileCreator is able to render a graphics.Model into an image file.
type ImageFileCreator struct {
	mdl  *graphics.Model
	dc   *gg.Context
	font *truetype.Font
}

// NewImageFileCreator consumes a font object parameter in order to avoid
// the (presumed expensive) cost of the font-parsing operation in every
// Create() operation. (e.g. truetype.Parse(goregular.TTF)). Create() is
// supposed to be more or less instant to support the anticipated UXP.
func NewImageFileCreator(font *truetype.Font) *ImageFileCreator {
	return &ImageFileCreator{font: font}
}

// Create renders a graphics model into an image file.
func (cr *ImageFileCreator) Create(
	filePath string, encoding Encoding, mdl *graphics.Model) error {
	// Initialise the Creator's state.
	cr.mdl = mdl
	cr.dc = gg.NewContext(int(mdl.Width), int(mdl.Height))

	// Now do the rendering and saving.
	cr.paintBackground()
	cr.renderLines()
	cr.renderPolygons()
	cr.renderText()
	err := cr.save(filePath, encoding)
	if err != nil {
		return fmt.Errorf("Create(): %v", err)
	}
	return nil
}

func (cr ImageFileCreator) paintBackground() {
	cr.dc.SetColor(colornames.White)
	cr.dc.DrawRectangle(0, 0, float64(cr.mdl.Width), float64(cr.mdl.Height))
	cr.dc.Fill()
}

func (cr ImageFileCreator) renderLines() {
	cr.dc.SetColor(colornames.Black)
	for _, line := range cr.mdl.Primitives.Lines {
		cr.setDashStyle(&line)
		cr.dc.DrawLine(line.P1.X, line.P1.Y, line.P2.X, line.P2.Y)
		cr.dc.Stroke()
	}
}

func (cr ImageFileCreator) renderPolygons() {
	cr.dc.SetColor(colornames.Black)
	for _, poly := range cr.mdl.Primitives.FilledPolys {
		start := poly[0]
		cr.dc.MoveTo(start.X, start.Y)
		for _, vertex := range poly {
			cr.dc.LineTo(vertex.X, vertex.Y)
		}
		cr.dc.ClosePath()
		cr.dc.Fill()
	}
}

// gg has a text justification continuum on which zero is left and 1.0 is
// right.
var ggJustification = map[graphics.Justification]float64{
	graphics.Left:   0.0,
	graphics.Centre: 0.5,
	graphics.Right:  1.0,
	graphics.Bottom: 0.0,
	graphics.Top:    1.0,
}

func (cr ImageFileCreator) renderText() {
	cr.dc.SetColor(colornames.Black)
	for _, label := range cr.mdl.Primitives.Labels {
		// truetype.Options takes its Size parameter in point units.
		// With typical image DPI/PPI this a 1:1 correletion to the
		// size in pixels.
		face := truetype.NewFace(cr.font,
			&truetype.Options{Size: label.FontHeight})
		cr.dc.SetFontFace(face)
		cr.dc.DrawStringAnchored(
			label.TheString, label.Anchor.X, label.Anchor.Y,
			ggJustification[label.HJust], ggJustification[label.VJust])
	}
}

func (cr ImageFileCreator) setDashStyle(line *graphics.Line) {
	switch line.Dashed {
	case false:
		cr.dc.SetDash()
	case true:
		cr.dc.SetDash(cr.mdl.DashLineDashLen, cr.mdl.DashLineGapLen)
	}
}

func (cr ImageFileCreator) save(
	filePath string, encoding Encoding) error {
	var err error
	switch encoding {
	case PNG:
		err = cr.dc.SavePNG(filePath)
	case JPG:
		quality := 92
		err = cr.dc.SaveJPG(filePath, quality)
	default:
		return fmt.Errorf(
			"save(): Not implemented encoding value: %v", encoding)
	}
	if err != nil {
		return fmt.Errorf("save() of file: <%v>: %v", filePath, err)
	}
	return nil
}
