package parser

import (
	"fmt"

	"github.com/peterhoward42/umli/dsl"
)

// MustCompileParse is a short-form wrapper for the parser, that panics
// if errors are returned from the Parse() method. It is thus good for
// reducing the code in tests, but is not suitable for apps and services.
func MustCompileParse(DSLScript string) *dsl.Model {
	parser := NewParser(DSLScript)
	model, err := parser.Parse()
	if err != nil {
		msg := fmt.Sprintf("MustCompile(): %v", err)
		panic(msg)
	}
	return model
}
