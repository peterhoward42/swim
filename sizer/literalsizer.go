package sizer

import "fmt"

/*
LiteralSizer implements the Sizer interface by providing a trivial
implementation that provides values by looking then up from a map you
provided to NewLiteralSizer.
*/
type LiteralSizer struct {
	m map[string]float64
}

// Make sure LiteralSizer implements Sizer at compile time.
var ls LiteralSizer
var _ Sizer = ls

// NewLiteralSizer provides a LiteralSizer that has been initialised
// and is ready for use.
func NewLiteralSizer(m map[string]float64) *LiteralSizer {
	return &LiteralSizer{m}
}

// Get returns the size specified by propertyName, or panics
// if the property is not recognized.
func (s LiteralSizer) Get(propertyName string) (size float64) {
	v, ok := s.m[propertyName]
	if !ok {
		msg := fmt.Sprintf("Sizer could not look up the key: <%s>", propertyName)
		panic(msg)
	}
	return v
}
