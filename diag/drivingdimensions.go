package diag

import "github.com/peterhoward42/umli/dsl"

/*
DrivingDimensions knows how to calculate the diagram width and the font height
that the diagram creation process will use as the fundamental sizes
from which all other sizing and spacing is derived.
*/
type DrivingDimensions struct{}

// WidthAndFontHeight provides the diagram width and font height.
func (dd DrivingDimensions) WidthAndFontHeight(dslModel dsl.Model) (
	diagWidth, fontHeight float64) {
	// The diagram width is in a sense arbitrary, because the contract of
	// of the diag package is that it will choose an arbitrary diagram
	// width that is convenient to itself, and produce a graphics.Model
	// accordingly. Renderers of a graphics.Model are obliged to scale the
	// coordinates to suit their rendering needs.

	// We choose 2000 because its easy to reason about during debugging if
	// you think of it as pixels.
	width := 2000.0

	const defaultTextHeightRatio = 1.0 / 100.0 // Works  well empirically.
	textHeightRatio := defaultTextHeightRatio
	sizeValue, ok := dslModel.SizeFromTextStatement()
	if ok {
		if sizeValue == 0 {
			panic("Developer error, this must not be allowed to happen")
		}	
		// 5  -> 0.005
		// 10 -> 0.010
		// 20 -> 0.020
		textHeightRatio = sizeValue / 1000.0
	}
	fontHeight = width * textHeightRatio
	return width, fontHeight
}
