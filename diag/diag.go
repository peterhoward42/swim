package diag

import (
	"fmt"
	"github.com/peterhoward42/umli/dsl"
	"github.com/peterhoward42/umli/graphics"
)

/*
Creator is the type that provides the API and entry point for the diag package.
It provides the main Create method that produces a diagram.
*/
type Creator struct {
}

/*
NewCreator instantiates a Creator ready to use.
*/
func NewCreator() (*Creator, error) {
	return nil, nil
}

/*
Create is the main API method which work out what the diagram should look like.
It orchestrates a multi-pass creation process which accumulates the graphics
primitives required in its graphicsModel and then returns that model.
*/
func (c *Creator) Create(dslModel dsl.Model) (*graphics.Model, error) {
	dd := &DrivingDimensions{}
	width, fontHeight, err := dd.WidthAndFontHeight(dslModel)
	if err != nil {
		return nil, fmt.Errorf("WidthAndFontHeight: %v", err)
	}
	return nil, nil
}
