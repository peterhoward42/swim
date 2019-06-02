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
	// Width of the diagram required - in virtual pixels.
	width int
	// Font height is used as the root for all sizing decisions.
	fontHeight float64
	// Parsed DSL script.
	allStatements []*dslmodel.Statement
	// The statements representing lifelines - isolated.
	lifelineStatements []*dslmodel.Statement
	// Keeps track of activity boxes in progress.
	allBoxStates allBoxStates
	// The output.
	graphicsModel *graphics.Model
	// Knows how to size everything.
	sizer *sizers.Sizer
	// Gradually moves down the page during creation.
	tideMark float64
}

/*
NewCreator provides a Creator ready to use.
*/
func NewCreator(width int, fontHeight float64,
	allStatements []*dslmodel.Statement) *Creator {
	lifelineStatements := isolateLifelines(allStatements)
	allBoxStates := newAllBoxStates(lifelineStatements)
	sizer := sizers.NewSizer(width, fontHeight, lifelineStatements)
	creator := &Creator{
		width:              width,
		fontHeight:         fontHeight,
		allStatements:      allStatements,
		lifelineStatements: lifelineStatements,
		allBoxStates:       allBoxStates,
		sizer:              sizer,
	}
	return creator
}

/*
Create is the main API method which work out what the diagram should look
like. It orchestrates a multi-pass creation process which accumulates
the graphics primitives required in its graphicsModel attribute and then
returns that model.
*/
func (c *Creator) Create() *graphics.Model {
	diagHeight := 0 // Set later to accomodate contents once known.
	c.graphicsModel = graphics.NewModel(
		c.width, diagHeight, c.fontHeight,
		c.sizer.DashLineDashLen, c.sizer.DashLineDashGap)
	c.createFirstPass()
	c.finalizeActivityBoxes()
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
			// Evaluate and add the graphics primitives required.
			c.graphicsForDrawingEvent(evt, statement)
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
	statement *dslmodel.Statement) {

	switch evt {
	case EndBox:
		c.endBox(statement)
	case InteractionLine:
		c.interactionLine(statement)
	case InteractionLabel:
		c.interactionLabel(statement)
	case LaneLine:
	case LaneTitleBox:
		c.laneTitleBox(statement)
	case SelfInteractionLines:
		c.selfInteractionLines(statement)
	case PotentiallyStartFromBox:
		c.potentiallyStartFromBox(statement)
	case PotentiallyStartToBox:
		c.potentiallyStartToBox(statement)
	}
}

/*
laneTitleBox generates the lines to represent the rectangular box at the top
of a lane, and calculates the tide mark corresponding to the bottom of these
boxes.
*/
func (c *Creator) laneTitleBox(
	statement *dslmodel.Statement) {
	thisLane := c.sizer.Lanes.Individual[statement]
	// First the rectangular box
	left := thisLane.TitleBoxLeft
	right := thisLane.TitleBoxRight
	top := c.sizer.DiagramPadT
	bot := c.sizer.DiagramPadT + c.sizer.Lanes.TitleBoxHeight
	c.graphicsModel.Primitives.AddRect(left, top, right, bot)
	// Now the strings
	nRows := len(statement.LabelSegments)
	for i, str := range statement.LabelSegments {
		rowOffset := float64(nRows-1-i) * c.fontHeight
		y := top + c.sizer.Lanes.TitleBoxBottomRowOfText - rowOffset
		c.graphicsModel.Primitives.AddLabel(str, c.fontHeight, thisLane.Centre,
			y, graphics.Centre, graphics.Bottom)
	}
	// In the particular case of a title box, the tide mark can
	// be set absolutely rather than advancing it by an increment.
	c.tideMark = bot + c.sizer.Lanes.TitleBoxPadB
}

/*
interactionLabel generates the labels that sit above one of the horizontal
interaction lines. It then claims the vertical space it has consumed for
itself by advancing the tide mark.
*/
func (c *Creator) interactionLabel(
	statement *dslmodel.Statement) {
	sourceLane := statement.ReferencedLanes[0]
	destLane := statement.ReferencedLanes[1]
	x, horizJustification := c.sizer.Lanes.InteractionLabelPosition(
		sourceLane, destLane, c.sizer.InteractionLineLabelIndent)
	firstRowY := c.tideMark
	c.rowOfLabels(x, firstRowY, horizJustification, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
}

/*
selfInteractionLabels generates the labels that sit above one of the *self*
interaction loops. It then claims the vertical space it has consumed for
itself by advancing the tide mark.
*/
/*
func (c *Creator) selfInteractionLabels(
	statement *dslmodel.Statement) {
	theLane := statement.ReferencedLanes[0]
	labelX := c.sizer.Lanes.Individual[theLane].ActivityBoxRight +
		c.sizer.InteractionLineLabelIndent
	firstRowY := c.tideMark
	c.rowOfLabels(labelX, firstRowY, graphics.Left, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
}
*/

/*
rowOfLabels is a (DRY) helper function to make the graphics.Primitives
objects for the set of strings in a label. It hard-codes top vertical
justification.
*/
func (c *Creator) rowOfLabels(centreX float64, firstRowY float64,
	horizJustification graphics.Justification, labelSegments []string) {
	for i, labelSeg := range labelSegments {
		y := firstRowY + float64(i)*c.fontHeight
		c.graphicsModel.Primitives.AddLabel(labelSeg, c.fontHeight,
			centreX, y, horizJustification, graphics.Top)
	}
}

/*
interactionLine generates the horizontal line and arrow head.  It then claims
the vertical space it claims for itself by advancing the tide mark.
*/
func (c *Creator) interactionLine(
	statement *dslmodel.Statement) {
	sourceLane := statement.ReferencedLanes[0]
	destLane := statement.ReferencedLanes[1]
	x1, x2 := c.sizer.Lanes.InteractionLineEndPoints(sourceLane, destLane)
	y := c.tideMark
	c.graphicsModel.Primitives.AddLine(x1, y, x2, y,
		statement.Keyword == umli.Dash)
	arrowVertices := makeArrow(x1, x2, y, c.sizer.ArrowLen,
		c.sizer.ArrowHeight)
	c.graphicsModel.Primitives.AddFilledPoly(arrowVertices)
	c.tideMark += 0.5*c.sizer.ArrowHeight + c.sizer.InteractionLinePadB
}

/*
selfInteractionLines generates the three lines and arrow head for a *self*
interaction loop.  It then claims the vertical space it claims for itself by
advancing the tide mark.
*/
func (c *Creator) selfInteractionLines(
	statement *dslmodel.Statement) {
	theLane := statement.ReferencedLanes[0]
	left := c.sizer.Lanes.Individual[theLane].ActivityBoxRight
	right := c.sizer.Lanes.Individual[theLane].SelfLoopRight
	top := c.tideMark
	bot := c.tideMark + c.sizer.SelfLoopHeight

	prims := c.graphicsModel.Primitives
	prims.AddLine(left, top, right, top, false)
	prims.AddLine(right, top, right, bot, false)
	prims.AddLine(right, bot, left, bot, false)
	arrowVertices := makeArrow(right, left, bot,
		c.sizer.ArrowLen, c.sizer.ArrowHeight)
	prims.AddFilledPoly(arrowVertices)

	// Labels go inside the self-loop.
	labelX := left + c.sizer.InteractionLineLabelIndent
	n := len(statement.LabelSegments)
	firstRowY := bot - float64(n)*c.fontHeight - c.sizer.InteractionLineTextPadB
	c.rowOfLabels(labelX, firstRowY, graphics.Left, statement.LabelSegments)

	c.tideMark = bot + c.sizer.InteractionLinePadB
}

/*
potentiallyStartToBox works out if the Creator has already started a
lifeline activity box for the lifeline that this interaction line is
going to, and if it hasn't does so by drawing the top edge.
*/
func (c *Creator) potentiallyStartToBox(
	statement *dslmodel.Statement) {
	behindTidemarkDelta := 0.0
	c.potentiallyStartActivityBox(statement.ReferencedLanes[1],
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
	statement *dslmodel.Statement) {
	behindTidemarkDelta := c.sizer.ActivityBoxVerticalOverlap
	c.potentiallyStartActivityBox(statement.ReferencedLanes[0],
		behindTidemarkDelta)
}

// potentiallyStartActivityBox is a DRY helper to (potentially) note the
// top edge of a lifeline's activity box in c.allBoxStates.
func (c *Creator) potentiallyStartActivityBox(
	lifeline *dslmodel.Statement, behindTidemarkDelta float64) {
	// Already a box in progress?
	if c.allBoxStates[lifeline].inProgress {
		return
	}
	y := c.tideMark - behindTidemarkDelta
	c.allBoxStates[lifeline].inProgress = true
	c.allBoxStates[lifeline].topY = y
}

/*
endBox finishes off a lifeline activity box in response to an
explicit *stop* instruction in the DSL. It then advances the
tide mark to a little beyond the bottom of the box.
*/
func (c *Creator) endBox(
	endBoxStatement *dslmodel.Statement) {
	lifeline := endBoxStatement.ReferencedLanes[0]
	bottom := c.tideMark
	c.finalizeActivityBox(lifeline, bottom)
}

// finalizeActivityBoxes identifies lifeline activity boxes that
// have been started, but not *stopped*, and draws them now that
// their size and position can be determined.
func (c *Creator) finalizeActivityBoxes() {
	bottom := c.tideMark + c.sizer.ActivityBoxVerticalOverlap
	for lifeline, boxState := range c.allBoxStates {
		if !boxState.inProgress {
			continue
		}
		c.finalizeActivityBox(lifeline, bottom)
	}
}

/*
finalizeActivityBox is a DRY helper that draws a single lifeline activity box -
based on the top Y coordinate stored in c.allBoxStates and the given bottom Y
coordinate. It then advances the tide mark to the bottom value provided.
*/
func (c *Creator) finalizeActivityBox(
	lifeline *dslmodel.Statement, bottom float64) {
	// Silently ignore this mandate if the lifeline does not have
	// a box in progress. (The parser cannot trap this because it knows
	// nothing about activity box inference).
	if c.allBoxStates[lifeline].inProgress == false {
		return
	}
	top := c.allBoxStates[lifeline].topY
	left := c.sizer.Lanes.Individual[lifeline].ActivityBoxLeft
	right := c.sizer.Lanes.Individual[lifeline].ActivityBoxRight
	c.graphicsModel.Primitives.AddRect(left, top, right, bottom)
	c.tideMark = bottom
	c.allBoxStates[lifeline].inProgress = false
}
