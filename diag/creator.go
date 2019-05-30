package diag

import (
	"github.com/peterhoward42/umli"
	"github.com/peterhoward42/umli/dslmodel"
	"github.com/peterhoward42/umli/graphics"
	sizers "github.com/peterhoward42/umli/sizer"
)

/*
Creator is the top level type for the diag package and provides the API
to create diagrams.
*/
type Creator struct {
	width              int
	fontHeight         float64
	allStatements      []*dslmodel.Statement
	lifelineStatements []*dslmodel.Statement
	boxState           *BoxState
	graphicsModel      *graphics.Model
	sizer              *sizers.Sizer
	tideMark           float64 // Gradually moves down the page during creation.
}

/*
NewCreator provides a Creator ready to use.  All the layout and sizing
decisions are derived from the diagram width and fontHeight parameters (in
pixels), while the required height is calculated automatially.
*/
func NewCreator(width int, fontHeight float64,
	allStatements []*dslmodel.Statement) *Creator {
	lifelineStatements := isolateLifelines(allStatements)
	boxState := NewBoxState(lifelineStatements)
	sizer := sizers.NewSizer(width, fontHeight, lifelineStatements)
	creator := &Creator{
		width:              width,
		fontHeight:         fontHeight,
		allStatements:      allStatements,
		lifelineStatements: lifelineStatements,
		boxState:           boxState,
		sizer:              sizer,
	}
	return creator
}

/*
Create is the main API method which work out what the diagram should look
like. It orchestrates a multi-pass creation process.
*/
func (c *Creator) Create() *graphics.Model {
	diagHeight := 0 // Set later.
	c.graphicsModel = graphics.NewModel(
		c.width, diagHeight, c.fontHeight,
		c.sizer.DashLineDashLen, c.sizer.DashLineDashGap)
	c.createFirstPass()
	c.finalizeDiagramHeight()
	return c.graphicsModel
}

// isolateLifelines provides the subset of Statements from the
// given list that correspond to lifelines.
func isolateLifelines(
	allStatements []*dslmodel.Statement) []*dslmodel.Statement {
	lifelineStatements := []*dslmodel.Statement{}
	for _, statement := range allStatements {
		if statement.Keyword == umli.Lane {
			lifelineStatements = append(lifelineStatements, statement)
		}
	}
	return lifelineStatements
}

/*
createFirstPass takes each parsed statement from the DSL script in turn, to
generate the primitives required that can be determined from a first pass.
This includes for example the lane title boxes and the interaction lines and
labels. But it excludes the generation of primitives that can only be
dimensioned once the interaction line Y coordinates are known; for example
the activity boxes that sit on lanes.
*/
func (c *Creator) createFirstPass() {
	graphicalEvents := NewScanner().Scan(c.allStatements)
	// Outer loop is per DSL statement
	for _, statement := range c.allStatements {
		statementEvents := graphicalEvents[statement]
		// Inner loop is for the (multiple) graphical events called for
		// by that statement.
		for _, evt := range statementEvents {
			// Evaluate and accumulate the graphics primitives required.
			prims := c.graphicsForDrawingEvent(evt, statement)
			c.graphicsModel.Primitives.Add(prims)
		}
	}
}

/*
finalizeDiagramHeight sets the graphics model's Height attribute to just
large enough to accomodate the final tide mark.
*/
func (c *Creator) finalizeDiagramHeight() {
	c.graphicsModel.Height = int(c.tideMark + c.sizer.DiagramPadB)
}

/*
graphicsForDrawingEvent synthesizes the lines and label strings primititives
required to render a single diagram element drawing event. It also advances
c.tideMark, to accomodate the space taken up by what it generated.
*/
func (c *Creator) graphicsForDrawingEvent(evt EventType,
	statement *dslmodel.Statement) (prims *graphics.Primitives) {

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
		prims = c.selfInteractionLines(statement)
	case SelfInteractionLabel:
		prims = c.selfInteractionLabels(statement)
	case PotentiallyStartFromBox:
		prims = c.potentiallyStartFromBox(statement)
	case PotentiallyStartToBox:
		prims = c.potentiallyStartToBox(statement)
	}
	return prims
}

/*
laneTitleBox generates the lines to represent the rectangular box at the top
of a lane, and calculates the tide mark corresponding to the bottom of these
boxes.
*/
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

/*
interactionLabel generates the labels that sit above one of the horizontal
interaction lines. It then claims the vertical space it has consumed for
itself by advancing the tide mark.
*/
func (c *Creator) interactionLabel(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	leftLane := statement.ReferencedLanes[0]
	rightLane := statement.ReferencedLanes[1]
	centreX := 0.5 * (c.sizer.Lanes.Individual[leftLane].Centre +
		c.sizer.Lanes.Individual[rightLane].Centre)
	firstRowY := c.tideMark
	prims = c.rowOfLabels(centreX, firstRowY, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
	return prims
}

/*
selfInteractionLabels generates the labels that sit above one of the *self*
interaction loops. It then claims the vertical space it has consumed for
itself by advancing the tide mark.
*/
func (c *Creator) selfInteractionLabels(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	theLane := statement.ReferencedLanes[0]
	labelCentreX := c.sizer.Lanes.Individual[theLane].SelfLoopCentre
	firstRowY := c.tideMark
	prims = c.rowOfLabels(labelCentreX, firstRowY, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
	return prims
}

/*
rowOfLabels is a (DRY) helper function to make the graphics.Primitives
objects for the set of strings in a label. It hard-codes centred horizontal
justification and top vertical justification.
*/
func (c *Creator) rowOfLabels(centreX float64, firstRowY float64,
	labelSegments []string) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	for i, labelSeg := range labelSegments {
		y := firstRowY + float64(i)*c.fontHeight
		prims.AddLabel(labelSeg, c.fontHeight, centreX, y,
			graphics.Centre, graphics.Top)
	}
	return prims
}

/*
interactionLine generates the horizontal line and arrow head.  It then claims
the vertical space it claims for itself by advancing the tide mark.
*/
func (c *Creator) interactionLine(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	fromLane := statement.ReferencedLanes[0]
	toLane := statement.ReferencedLanes[1]
	x1, x2 := c.sizer.Lanes.InteractionLineEndPoints(fromLane, toLane)
	y := c.tideMark
	prims.AddLine(x1, y, x2, y, statement.Keyword == umli.Dash)
	arrowVertices := makeArrow(x1, x2, y, c.sizer.ArrowLen,
		c.sizer.ArrowHeight)
	prims.AddFilledPoly(arrowVertices)
	c.tideMark += 0.5*c.sizer.ArrowHeight + c.sizer.InteractionLinePadB
	return prims
}

/*
selfInteractionLines generates the three lines and arrow head for a *self*
interaction loop.  It then claims the vertical space it claims for itself by
advancing the tide mark.
*/
func (c *Creator) selfInteractionLines(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	theLane := statement.ReferencedLanes[0]
	left := c.sizer.Lanes.Individual[theLane].ActivityBoxRight
	right := c.sizer.Lanes.Individual[theLane].SelfLoopRight
	top := c.tideMark
	bot := c.tideMark + c.sizer.SelfLoopHeight

	prims.AddLine(left, top, right, top, false)
	prims.AddLine(right, top, right, bot, false)
	prims.AddLine(right, bot, left, bot, false)
	arrowVertices := makeArrow(right, left, bot,
		c.sizer.ArrowLen, c.sizer.ArrowHeight)
	prims.AddFilledPoly(arrowVertices)
	c.tideMark = bot + c.sizer.InteractionLinePadB
	return prims
}

/*
potentiallyStartToBox works out if the Creator has already started a
lifeline activity box for the lifeline that this interaction line is
going to, and if it hasn't does so by drawing the top edge.
*/
func (c *Creator) potentiallyStartToBox(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	behindTidemarkDelta := 0.0
	return c.potentiallyStartActivityBox(statement.ReferencedLanes[1],
		behindTidemarkDelta)
}

/*
potentiallyStartFromBox works out if the Creator has already started a
lifeline activity box for the lifeline that this interaction line is
being emited from, and if it hasn't does so by drawing the top edge.
Note it is atypical because it renders behind the tidemark, to position the
start of the box a little before the interaction line, but then leaves the
tidemark unchanged, so that the interaction line that follows, stays in contact
with its label (which has already been emitted).
*/
func (c *Creator) potentiallyStartFromBox(
	statement *dslmodel.Statement) (prims *graphics.Primitives) {
	behindTidemarkDelta := c.sizer.ActivityBoxTopPadB
	return c.potentiallyStartActivityBox(statement.ReferencedLanes[0],
		behindTidemarkDelta)
}

// potentiallyStartActivityBox is a DRY helper to (potentially) draw the
// top edge of a lifeline's activity box.
func (c *Creator) potentiallyStartActivityBox(
	lifeline *dslmodel.Statement, behindTidemarkDelta float64) (
	prims *graphics.Primitives) {
	prims = graphics.NewPrimitives()
	// Already a box in progress?
	if c.boxState.boxIsInProgress(lifeline) {
		return prims
	}
	left := c.sizer.Lanes.Individual[lifeline].ActivityBoxLeft
	right := c.sizer.Lanes.Individual[lifeline].ActivityBoxRight
	// Render potentially **behind** the tidemark.
	y := c.tideMark - behindTidemarkDelta
	prims.AddLine(left, y, right, y, false)
	c.boxState.boxesInProgress[lifeline] = true
	return prims
}
