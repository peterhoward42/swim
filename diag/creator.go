package diag

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
	"github.com/peterhoward42/umlinteraction/graphics"
	"github.com/peterhoward42/umlinteraction/sizers"
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
			graphicPrimitives := c.graphicsForDrawingEvent(evt, statement)
			graphicsModel.Primitives.Append(graphicPrimitives)
		}
	}
	return graphicsModel
}

// graphicsForDrawingEvent synthesizes the lines and label strings required
// to render a single diagram element drawing event. In so doing it
// also advances the creators tide mark.
func (c *Creator) graphicsForDrawingEvent(
	evt EventType, statement *dslmodel.Statement) *graphics.Primitives {
	switch evt {
	case EndBox:
	case InteractionLine:
	case InteractionLabel:
	case LaneLine:
	case LaneTitleBox:
		return c.laneTitleBox(statement)
	case LaneTitleLabel:
	case SelfInteractionLines:
	case SelfInteractionLabel:
	case PotentiallyStartFromBox:
	case PotentiallyStartToBox:
	}
	return graphics.NewPrimitives()
}

// laneTitleBox generates the lines to represent the
// rectangular box at the top of a lane, and sets the tide mark
// to the bottom of these boxes.
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) *graphics.Primitives {
	prims := graphics.NewPrimitives()
	thisLane := c.sizer.Lanes.Individual[statement]
	left := thisLane.TitleBoxLeft
	right := thisLane.TitleBoxRight
	top := c.sizer.TopMargin
	bot := c.sizer.TopMargin + c.sizer.Lanes.TitleBoxHeight
	prims.AddRect(left, top, right, bot)
	c.tideMark = bot
	return prims
}
