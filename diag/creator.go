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
func NewCreator(statements []*dslmodel.Statement) *Creator {
	return &Creator{}
}

// Create produces the graphics model corresponding to the statements
// provided at construction time.
func (c *Creator) Create() *graphics.Model {
	return nil
}
