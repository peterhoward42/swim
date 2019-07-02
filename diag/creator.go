package diag

/*
This module contains the *Creator* type which provides the public interface
to clients of the *diag* package. Clients should construct a Creator with
NewCreator() and then call its Create() method.

The module then provides the high-level implementation for Create() and
expresses the essential creation algorith - delegating much of its work to
code in other modules in the package.
*/

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
	// In charge of making the outer frame and title.
	frameMaker *frameMaker
	// Keeps track of activity box top and bottom coordinates.
	activityBoxes map[*dslmodel.Statement]*lifelineBoxes
	// Keeps track of the space taken up by interaction lines.
	ilZones *InteractionLineZones
	// A delegate to make the lifelines dashed line segments.
	lifelineMaker *lifelines
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
	activityBoxes := map[*dslmodel.Statement]*lifelineBoxes{}
	for _, s := range lifelineStatements {
		activityBoxes[s] = newlifelineBoxes()
	}
	sizer := sizers.NewSizer(width, fontHeight, lifelineStatements)
	creator := &Creator{
		width:              width,
		fontHeight:         fontHeight,
		allStatements:      allStatements,
		lifelineStatements: lifelineStatements,
		activityBoxes:      activityBoxes,
		sizer:              sizer,
	}
	creator.frameMaker = newFrameMaker(creator)
	creator.ilZones = NewInteractionLineZones(creator)
	creator.lifelineMaker = newLifelines(creator)
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
	c.finalizeLifelines()
	c.finalizeDiagramHeight()
	return c.graphicsModel
}

// isolateLifelines provides the subset of Statements from the
// given list that correspond to lifelines.
func isolateLifelines(
	allStatements []*dslmodel.Statement) []*dslmodel.Statement {
	lifelineStatements := []*dslmodel.Statement{}
	for _, statement := range allStatements {
		if statement.Keyword == umli.Life {
			lifelineStatements = append(lifelineStatements, statement)
		}
	}
	return lifelineStatements
}

/*
createFirstPass takes each parsed statement from the DSL script in turn, to
generate the primitives required that can be determined from a first pass.
This includes for example the lifeline title boxes and the interaction lines and
labels. But it excludes the generation of primitives that can only be
dimensioned once the interaction line Y coordinates are known; for example
the activity boxes that sit on lifelines.
*/
func (c *Creator) createFirstPass() {
	graphicalEvents := newScanner().Scan(c.allStatements)
	c.tideMark = c.sizer.DiagramPadT
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
func (c *Creator) graphicsForDrawingEvent(evt eventType,
	statement *dslmodel.Statement) {

	switch evt {
	case EndBox:
		c.endBox(statement)
	case InteractionLine:
		c.interactionLine(statement)
	case InteractionLabel:
		c.interactionLabel(statement)
	case LifelineTitleBox:
		c.lifelineTitleBox(statement)
	case SelfInteractionLines:
		c.selfInteractionLines(statement)
	case PotentiallyStartFromBox:
		c.potentiallyStartFromBox(statement)
	case PotentiallyStartToBox:
		c.potentiallyStartToBox(statement)
	case Frame:
		c.frame(statement)
	}
}

func (c *Creator) frame(statement *dslmodel.Statement) {
	// Quite complex - so delegate.
	c.frameMaker.initFrameAndMakeTitleBox(statement)
}

/*
lifelineTitleBox generates the lines to represent the rectangular box at the top
of a lifeline, and advances the tide mark corresponding to the depth they
occupy.
*/
func (c *Creator) lifelineTitleBox(
	statement *dslmodel.Statement) {
	thisLifeline := c.sizer.Lifelines.Individual[statement]
	// First make the rectangular box
	left := thisLifeline.TitleBoxLeft
	right := thisLifeline.TitleBoxRight
	var top float64
	var bot float64
	// For the first title box we encounter, we evaluate the top and bottom
	// coordinate for it and all other title boxes based on the tidemark.
	// And remember these coordinates, and advance the tidemark. For all the
	// the others, we use the saved coordinates and leave the tidemark alone.
	if c.lifelineMaker.titleBoxTopAndBottom == nil {
		top = c.tideMark
		bot = top + c.sizer.Lifelines.TitleBoxHeight
		c.lifelineMaker.titleBoxTopAndBottom = &segment{top, bot}
		c.tideMark += c.sizer.Lifelines.TitleBoxHeight
		c.tideMark += c.sizer.Lifelines.TitleBoxPadB
	} else {
		top = c.lifelineMaker.titleBoxTopAndBottom.start
		bot = c.lifelineMaker.titleBoxTopAndBottom.end
	}
	c.graphicsModel.Primitives.AddRect(left, top, right, bot)

	// Make the Label
	n := len(statement.LabelSegments)
	firstRowY := bot - float64(n)*c.fontHeight - c.sizer.Lifelines.TitleBoxLabelPadB
	c.rowOfLabels(thisLifeline.Centre, firstRowY, graphics.Centre,
		statement.LabelSegments)
}

/*
interactionLabel generates the labels that sit above one of the horizontal
interaction lines. It then claims the vertical space it has consumed for
itself by advancing the tide mark. And registers this space claim with
the creator's InteractionLineZones component.
*/
func (c *Creator) interactionLabel(
	statement *dslmodel.Statement) {
	sourceLifeline := statement.ReferencedLifelines[0]
	destLifeline := statement.ReferencedLifelines[1]
	x, horizJustification := c.sizer.Lifelines.InteractionLabelPosition(
		sourceLifeline, destLifeline, c.sizer.InteractionLineLabelIndent)
	firstRowY := c.tideMark
	c.rowOfLabels(x, firstRowY, horizJustification, statement.LabelSegments)
	c.tideMark += float64(len(statement.LabelSegments))*
		c.fontHeight + c.sizer.InteractionLineTextPadB
	c.ilZones.RegisterSpaceClaim(
		sourceLifeline, destLifeline, firstRowY, c.tideMark)
}

/*
rowOfLabels is a (DRY) helper function to make the graphics.Primitives
objects for the set of strings in a label. It hard-codes top vertical
justification. It does not advance the tideMark.
*/
func (c *Creator) rowOfLabels(x float64, firstRowY float64,
	horizJustification graphics.Justification, labelSegments []string) {
	for i, labelSeg := range labelSegments {
		y := firstRowY + float64(i)*c.fontHeight
		c.graphicsModel.Primitives.AddLabel(labelSeg, c.fontHeight,
			x, y, horizJustification, graphics.Top)
	}
}

/*
interactionLine generates the horizontal line and arrow head.  It then claims
the vertical space it claims for itself by advancing the tide mark.  And
registers this space claim with the creator's InteractionLineZones component.
*/
func (c *Creator) interactionLine(
	statement *dslmodel.Statement) {
	sourceLifeline := statement.ReferencedLifelines[0]
	destLifeline := statement.ReferencedLifelines[1]
	x1, x2 := c.sizer.Lifelines.InteractionLineEndPoints(sourceLifeline, destLifeline)
	y := c.tideMark
	c.graphicsModel.Primitives.AddLine(x1, y, x2, y,
		statement.Keyword == umli.Dash)
	arrowVertices := makeArrow(x1, x2, y, c.sizer.ArrowLen,
		c.sizer.ArrowHeight)
	c.graphicsModel.Primitives.AddFilledPoly(arrowVertices)
	c.tideMark += 0.5*c.sizer.ArrowHeight + c.sizer.InteractionLinePadB
	c.ilZones.RegisterSpaceClaim(
		sourceLifeline, destLifeline, y, c.tideMark)
}

/*
selfInteractionLines generates a set of *self* interaction lines (i.e. a loop),
including the arrow head and labels. It then claims the vertical space it it
has occupied by advancing the tide mark.
*/
func (c *Creator) selfInteractionLines(
	statement *dslmodel.Statement) {
	theLifeline := statement.ReferencedLifelines[0]
	left := c.sizer.Lifelines.Individual[theLifeline].ActivityBoxRight
	right := c.sizer.Lifelines.Individual[theLifeline].SelfLoopRight
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
going to, and if it hasn't it registers where it should start.
*/
func (c *Creator) potentiallyStartToBox(
	statement *dslmodel.Statement) {
	behindTidemarkDelta := 0.0
	c.potentiallyStartActivityBox(statement.ReferencedLifelines[1],
		behindTidemarkDelta)
}

/*
potentiallyStartFromBox works out if the Creator has already started a
lifeline activity box when an interactionline is about to be emitted
from it. If it hasn't it registers where it should start.
Note it is atypical because it renders behind the tidemark, to position the
start of the box a little before the interaction line, but then leaves the
tidemark unchanged, so that the interaction line that follows, stays in contact
with its label (which has already been emitted).
*/
func (c *Creator) potentiallyStartFromBox(
	statement *dslmodel.Statement) {
	behindTidemarkDelta := c.sizer.ActivityBoxVerticalOverlap
	c.potentiallyStartActivityBox(statement.ReferencedLifelines[0],
		behindTidemarkDelta)
}

// potentiallyStartActivityBox is a DRY helper to (potentially) note the
// top edge of a lifeline's activity box in c.activityBoxes.
func (c *Creator) potentiallyStartActivityBox(
	lifeline *dslmodel.Statement, behindTidemarkDelta float64) {
	// Already a box in progress?
	if c.activityBoxes[lifeline].inProgress() {
		return
	}
	y := c.tideMark - behindTidemarkDelta
	c.activityBoxes[lifeline].startBoxAt(y)
}

/*
endBox finishes off a lifeline activity box in response to an
explicit *stop* instruction in the DSL. It then advances the
tide mark to a little beyond the bottom of the box.
*/
func (c *Creator) endBox(
	endBoxStatement *dslmodel.Statement) {
	lifeline := endBoxStatement.ReferencedLifelines[0]
	bottom := c.tideMark
	c.finalizeActivityBox(lifeline, bottom)
}

// finalizeActivityBoxes identifies lifeline activity boxes that
// have been started, but not *stopped*, and draws them now that
// their size and position can be determined.
func (c *Creator) finalizeActivityBoxes() {
	bottom := c.tideMark + c.sizer.ActivityBoxVerticalOverlap
	for lifeline := range c.activityBoxes {
		c.finalizeActivityBox(lifeline, bottom)
	}
	c.tideMark = bottom + c.sizer.FinalizedActivityBoxesPadB
}

/*
finalizeActivityBox is a DRY helper that draws a single lifeline activity box -
based on the top Y coordinate stored in c.activityBoxes and the given bottom Y
coordinate. It then advances the tide mark to the bottom value provided.
*/
func (c *Creator) finalizeActivityBox(
	lifeline *dslmodel.Statement, bottom float64) {
	// Skip those that have been stopped earlier explicitly with a *stop*
	// statement.
	if !c.activityBoxes[lifeline].inProgress() {
		return
	}
	top := c.activityBoxes[lifeline].mostRecent().extent.start
	left := c.sizer.Lifelines.Individual[lifeline].ActivityBoxLeft
	right := c.sizer.Lifelines.Individual[lifeline].ActivityBoxRight
	c.graphicsModel.Primitives.AddRect(left, top, right, bottom)
	c.tideMark = bottom
	c.activityBoxes[lifeline].terminateInProgressBoxAt(bottom)
}

// todo
func (c *Creator) finalizeLifelines() {
	// Quite complex - so delegate.
	c.lifelineMaker.produceLifelines()
}
