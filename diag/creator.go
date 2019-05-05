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
	hSizer     *sizers.Horizontal
	vSizer     *sizers.Vertical
	tideMark   int
}

// NewCreator creates a Creator ready to use.
func NewCreator(width int, fontHeight float64,
	statements []*dslmodel.Statement) *Creator {
	hSizer := sizers.NewHorizontal(width, fontHeight, statements)
	vSizer := sizers.NewVertical(width, fontHeight, statements)
	creator := &Creator{width, fontHeight, statements, hSizer, vSizer, 0}
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
			graphicsModel.Append(graphicPrimitives)
		}
	}
	return graphicsModel
}

// graphicsForDrawingEvent synthesizes the lines and label strings required
// to render the given single diagram element drawing event. In so doing it
// also advances the tide mark.
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
	return nil
}

// laneTitleBox generates the lines to represent the
// rectangular box at the top of a lane, and sets the tide mark
// to the bottom of the box.
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) *graphics.Primitives {
	left := c.hSizer.LaneTitleBoxes[statement].Left
	right := c.hSizer.LaneTitleBoxes[statement].Right
	topMargin := c.vSizer.TopMargin
	height := c.vSizer.TitleBoxHeight
	top := topMargin
	bot := top + height
	prims := graphics.NewPrimitives()
	prims.AddLines(graphics.Rect(top, left, right, bottom))
	c.tideMark = bot
	return prims
}
