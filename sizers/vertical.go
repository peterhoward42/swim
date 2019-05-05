package sizers

import (
	"github.com/peterhoward42/umlinteraction/dslmodel"
)

// Vertical makes decisions about horizontal sizing and spacing.
type Vertical struct {
	diagWidth  int
	fontHeight float64
	statements []*dslmodel.Statement
}

// NewVertical creates a Vertical ready to use.
func NewVertical(diagWidth int, fontHeight float64,
	statements []*dslmodel.Statement) *Vertical {
	return &Vertical{diagWidth, fontHeight, statements}
}
