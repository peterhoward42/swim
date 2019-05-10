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
	tideMark   float64
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
// based on the diagram width and font height (in pixels) parameters.
func (c *Creator) Create() *graphics.Model {
	graphicalEvents := NewScanner().Scan(c.statements)
	graphicsModel := graphics.NewModel(c.width, c.fontHeight)
	for _, statement := range c.statements {
		statementEvents := graphicalEvents[statement]
		for _, evt := range statementEvents {
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
	case LaneTitleLabel:
	case SelfInteractionLines:
	case SelfInteractionLabel:
	case PotentiallyStartFromBox:
	case PotentiallyStartToBox:
	}
	return prims, tideMark
}

// laneTitleBox generates the lines to represent the
// rectangular box at the top of a lane, and sets the tide mark
// to the bottom of these boxes.
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) (
	prims *graphics.Primitives, tideMark float64) {
	prims = graphics.NewPrimitives()
	thisLane := c.sizer.Lanes.Individual[statement]
	left := thisLane.TitleBoxLeft
	right := thisLane.TitleBoxRight
	top := c.sizer.TopMargin
	bot := c.sizer.TopMargin + c.sizer.Lanes.TitleBoxHeight
	prims.AddRect(left, top, right, bot)
	tideMark = bot
	return prims, tideMark
}
