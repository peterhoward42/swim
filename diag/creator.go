package diag

import (
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizers"
)

// Creator is the top level entry point for the diag package.
// It is capable of consuming a sequence of dslmodel.Statement(s), and
// producing the corresponding diagram definition in terms of its low-level
// primitives - i.e. lists of line segments and strings to render.
type Creator struct {
	width      int
	fontHeight float64
	statements []*dslmodel.Statement
	sizer      *sizers.Sizer
	tideMark   float64 // Gradually moves down the page during creation.
}

// NewCreator creates a Creator ready to use.
func NewCreator(width int, fontHeight float64,
	statements []*dslmodel.Statement) *Creator {
	sizer := sizers.NewSizer(width, fontHeight, statements)
	tideMark := 0.0
	creator := &Creator{width, fontHeight, statements, sizer, tideMark}
	return creator
}

// Create works out what the diagram should look like by analysing the
// DSL Statement(s) provided. All sizing and spacing decisions are
// based on the diagram width, and font height (in pixels) parameters.
func (c *Creator) Create() *graphics.Model {
	graphicalEvents := NewScanner().Scan(c.statements)
	initialHeightAssumption := int(0.33 * float64(c.width)) // Overriden later.
	graphicsModel := graphics.NewModel(
		c.width, initialHeightAssumption, c.fontHeight)
	// First pass is per-statement.
	for _, statement := range c.statements {
		statementEvents := graphicalEvents[statement]
		for _, evt := range statementEvents {
			// Inner loop is for each graphics event called for by the statement.
			prims, newTideMark := c.graphicsForDrawingEvent(evt, statement)
			c.tideMark = newTideMark
			graphicsModel.Primitives.Add(prims)
		}
	}
	return graphicsModel
}

// graphicsForDrawingEvent synthesizes the lines and label strings primititives
// required to render a single diagram element drawing event. It also returns
// and tide mark value, suitably advanced to accomodate the space taken
// up by the new primitives.
func (c *Creator) graphicsForDrawingEvent(
	evt EventType, statement *dslmodel.Statement) (
	prims *graphics.Primitives, tideMark float64) {

	prims = graphics.NewPrimitives()
	tideMark = 0

	switch evt {
	case EndBox:
	case InteractionLine:
	case InteractionLabel:
	case LaneLine:
	case LaneTitleBox:
		prims, tideMark = c.laneTitleBox(statement)
	case SelfInteractionLines:
	case SelfInteractionLabel:
	case PotentiallyStartFromBox:
	case PotentiallyStartToBox:
	}
	return prims, tideMark
}

// laneTitleBox generates the lines to represent the
// rectangular box at the top of a lane, and calculates the tide mark
// corresponding to the bottom of these boxes.
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) (
	prims *graphics.Primitives, tideMark float64) {
	prims = graphics.NewPrimitives()
	thisLane := c.sizer.Lanes.Individual[statement]
	// First the rectangular box
	left := thisLane.TitleBoxLeft
	right := thisLane.TitleBoxRight
	top := c.sizer.TopMargin
	bot := c.sizer.TopMargin + c.sizer.Lanes.TitleBoxHeight
	prims.AddRect(left, top, right, bot)
	tideMark = bot
	// Now the strings
	for i, str := range statement.LabelSegments {
		prims.AddLabel(str, c.fontHeight, thisLane.Centre,
			top+thisLane.TitleBoxFirstRowOfText+float64(i+1)*c.fontHeight,
			graphics.Centre, graphics.Bottom)
	}
	return prims, tideMark
}
