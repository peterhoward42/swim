package diag

import (
	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	"github.com/peterhoward42/umli/sizer"
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
	creator := &Creator{
		width:      width,
		fontHeight: fontHeight,
		statements: statements,
		sizer:      sizer,
	}
	return creator
}

// Create works out what the diagram should look like by analysing the
// DSL Statement(s) provided. It is responsible for gradually moving
// creator.tideMark down the page as each diagram element is produced.
// All sizing and spacing decisions are
// based on the diagram width, and font height (in pixels) parameters.
func (c *Creator) Create() *graphics.Model {
	graphicalEvents := NewScanner().Scan(c.statements)
	initialHeightAssumption := int(0.33 * float64(c.width)) // Overriden later.
	graphicsModel := graphics.NewModel(
		c.width, initialHeightAssumption, c.fontHeight,
		c.sizer.DashLineDashLen, c.sizer.DashLineDashGap)
	// First pass is per-statement.
	for _, statement := range c.statements {
		statementEvents := graphicalEvents[statement]
		for _, evt := range statementEvents {
			// Inner loop is for each graphics event called for by the statement.
			prims := c.graphicsForDrawingEvent(evt, statement)
			graphicsModel.Primitives.Add(prims)
		}
	}
	return graphicsModel
}

// graphicsForDrawingEvent synthesizes the lines and label strings primititives
// required to render a single diagram element drawing event. It also advances
// c.tideMark, to accomodate the space taken up by the new primitives.
func (c *Creator) graphicsForDrawingEvent(
	evt EventType, statement *dslmodel.Statement) (prims *graphics.Primitives) {

	prims = graphics.NewPrimitives()

	switch evt {
	case EndBox:
	case InteractionLine:
		prims = c.interactionLine(statement)
	case InteractionLabel:
		prims = c.interactionLabel(statement)
	case LaneLine:
	case LaneTitleBox:
		prims = c.laneTitleBox(statement)
	case SelfInteractionLines:
	case SelfInteractionLabel:
	case PotentiallyStartFromBox:
	case PotentiallyStartToBox:
	}
	return prims
}

// laneTitleBox generates the lines to represent the
// rectangular box at the top of a lane, and calculates the tide mark
// corresponding to the bottom of these boxes.
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	thisLane := c.sizer.Lanes.Individual[statement]
	// First the rectangular box
	left := thisLane.TitleBoxLeft
	right := thisLane.TitleBoxRight
	top := c.sizer.DiagramPadT
	bot := c.sizer.DiagramPadT + c.sizer.Lanes.TitleBoxHeight
	prims.AddRect(left, top, right, bot)
	// Now the strings
	nRows := len(statement.LabelSegments)
	for i, str := range statement.LabelSegments {
		rowOffset := float64(nRows-1-i) * c.fontHeight
		y := top + c.sizer.Lanes.TitleBoxBottomRowOfText - rowOffset
		prims.AddLabel(str, c.fontHeight, thisLane.Centre, y,
			graphics.Centre, graphics.Bottom)
	}
	// In the particular case of a title box, the tide mark can
	// be set absolutely rather than advancing it by an increment.
	c.tideMark = bot + c.sizer.Lanes.TitleBoxPadB
	return prims
}

// interactionLabel generates the labels that sit above one of the
// horizontal interaction arrows. It then claims the vertical space
// it has consumed for itself by advancing the tide mark.
func (c *Creator) interactionLabel(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	leftLane := statement.ReferencedLanes[0]
	rightLane := statement.ReferencedLanes[1]
	centreX := 0.5 * (c.sizer.Lanes.Individual[leftLane].Centre +
		c.sizer.Lanes.Individual[rightLane].Centre)
	for i, labelSeg := range statement.LabelSegments {
		y := c.tideMark + float64(i)*c.fontHeight
		prims.AddLabel(labelSeg, c.fontHeight, centreX, y,
			graphics.Centre, graphics.Top)
	}
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
	return prims
}

// interactionLine generates the horizontal line and arrow head.
// It then claims the vertical space
// it claims for itself by advancing the tide mark.
func (c *Creator) interactionLine(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	leftLane := statement.ReferencedLanes[0]
	rightLane := statement.ReferencedLanes[1]
	x1 := c.sizer.Lanes.Individual[leftLane].Centre
	x2 := c.sizer.Lanes.Individual[rightLane].Centre
	y := c.tideMark + 0.5*c.sizer.ArrowHeight
	prims.AddLine(x1, y, x2, y, statement.Keyword == umli.Dash)
	arrowVertices := makeArrow(x1, x2, y, c.sizer.ArrowLen, c.sizer.ArrowHeight)
	prims.AddFilledPoly(arrowVertices)
	c.tideMark += c.sizer.ArrowHeight + c.sizer.InteractionLinePadB
	return prims
}
