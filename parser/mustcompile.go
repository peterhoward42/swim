package parser

import (
	"fmt"

	"github.com/peterhoward42/umlinteraction/dslmodel"
)

// MustCompileParse is a short-form wrapper for the parser, that panics
// if errors are returned from the Parse() method.
func MustCompileParse(DSLScript string) []*dslmodel.Statement {
	statements, err := Parse(DSLScript)
	if err != nil {
		msg := fmt.Sprintf("MustCompile(): %v", err)
		panic(msg)
	}
	return statements
}
