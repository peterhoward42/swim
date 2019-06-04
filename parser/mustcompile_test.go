package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsedOutputIsCorrectWhenInputIsProper(t *testing.T) {
	assert := assert.New(t)
	statements := MustCompileParse(`
		life A foo
		life B bar
	`)
	assert.Len(statements, 2)
}

func TestPanicsWhenImportIsMalformed(t *testing.T) {
	assert := assert.New(t)
	malformed := `
		life A foo
		liXe B bar
	`
	assert.Panics(func() { MustCompileParse(malformed) })
}
