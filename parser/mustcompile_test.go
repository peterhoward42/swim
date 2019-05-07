package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsedOutputIsCorrectWhenInputIsProper(t *testing.T) {
	assert := assert.New(t)
	statements := MustCompileParse(`
		lane A foo
		lane B bar
	`)
	assert.Len(statements, 2)
}

func TestPanicsWhenImportIsMalformed(t *testing.T) {
	assert := assert.New(t)
	malformed := `
		lane A foo
		laXe B bar
	`
	assert.Panics(func() { MustCompileParse(malformed) })
}
