package diag

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
	"github.com/peterhoward42/umlinteraction/graphics"
)

// Creator is the top level entry point orchestrator for the diag package.
// It is capable of consuming a sequence of dslmodel.Statement(s), and 
// producing the corresponding diagram definition in terms of its low-level
// primitives - like lists of line segments and strings to render.
type Creator struct {
}

// NewCreator creates a Creator ready to use.
func NewCreator() *Creator {
	return &Creator{}
}

// Create works out what the diagram should look like by analysing the
// DSL Statement(s) provided. Provide the required width and height (in pixels). 
func (c *Creator) Create(
	statements []*dslmodel.Statement, width int, height int) *graphics.Model {
	scanner := NewScanner()
	// Capture the graphical events required for all statements
	eventsPerStatement := scanner.Scan(statements)
	// Iterate over them in statement-order
	for _, statement := range statements {
		statementEvents := eventsPerStatement[statement]
		// Iterate over the set of graphical events for an individual statement
		// processing them to produce the required line segments and strings to
		// be rendered. Doing so in the context of a Y-coordinate tide-mark that
		// progresses down the page, as each event *claims* the vertical
		// room it needs.
		var tideMark int // Lint doesn't like explicit init to zero.
		graphicsModel := graphics.NewModel(width, height)
		for _, drawingEvent := range(statementEvents) {
			c.addGraphicsForDrawingEvent(
				drawingEvent, &tideMark, graphicsModel)
		}
		return graphicsModel
	}
	return nil
}

// addGraphicsForDrawingEvent synthesizes the lines and label strings required
// to render the given single diagram element drawing event, and adds them to
// the graphics model provided. In so doing it advances the tide mark that
// represents the diagram gradually spreading down the page.
func (c *Creator) addGraphicsForDrawingEvent(event EventType, tideMark *int,
	graphicsModel *graphics.Model) {
}
