package sizer

import "fmt"

/*
CompleteSizer implements the Sizer interface by providing an implementation
that can provide the comprehensive set of sizing parameters required by
the diag package and its sub packages.

A widely adopted naming convention is illustrated by:

	FrameTitleTextPadT

Where

	FrameTitle: is the diagram element it refers to.
	Pad: is short for padding.
	T: chooses from {Top, Bottom, Left, Right, Centre}
*/
type CompleteSizer struct {
	fontHeight float64
}

// Make sure CompleteSizer implements Sizer at compile time.
var cs CompleteSizer
var _ Sizer = cs

// NewCompleteSizer provides a NewSizer structure that has been initialised
// and is ready for use.
func NewCompleteSizer(diagramWidth float64, fontHeight float64) *CompleteSizer {
	return &CompleteSizer{diagramWidth}
}

// Get returns the size specified by propertyName, or panics
// if the property is not recognized.
func (s CompleteSizer) Get(propertyName string) (size float64) {
	v, ok := table[propertyName]
	if !ok {
		msg := fmt.Sprintf("Sizer could not look up the key: <%s>", propertyName)
		panic(msg)
	}
	return v * s.fontHeight
}

var table = map[string]float64{

	// Whole diagram scope
	"DiagramPadT": 1.0,
	"DiagramPadB": 1.0,

	// Outer frame and diagram title
	"FramePadLR":         0.5,
	"FrameInternalPadB":  1.0,
	"FrameTitleBoxWidth": 100.0,
	"FrameTitleTextPadT": 0.5,
	"FrameTitleTextPadB": 1.0,
	"FrameTitleTextPadL": 1.0,
	"FrameTitleRectPadB": 1.0,

	// Lifeline title boxes
	"TitleBoxLabelPadT":          0.25,
	"TitleBoxLabelPadB":          1.0,
	"IdealLifelineTitleBoxWidth": 15.0,
	"TitleBoxPadB":               1.5,

	// Interaction lines
	"ArrowLen":                1.5,
	"InteractionLinePadB":     0.5,
	"InteractionLineTextPadB": 0.5,
	"SelfLoopHeight":          3.0,
	"SelfLoopWidthFactor":     0.7, // proportion of lifeline pitch

	// Dashes
	"DashLineDashLen": 0.5,
	"DashLineDashGap": 0.25,

	// Activity boxes
	"ActivityBoxWidth":           1.5,
	"ActivityBoxVerticalOverlap": 0.5,
	"FinalizedActivityBoxesPadB": 1.0,

	// Lifelines
	"MinLifelineSegLength": 0.5,
}
