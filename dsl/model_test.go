package dsl

import (
	"testing"

	"github.com/peterhoward42/umli"
	"github.com/stretchr/testify/assert"
)

func TestLifelineStatements(t *testing.T) {
	assert := assert.New(t)

	a := makeStatement("unused", umli.Life)
	b := makeStatement("unused", umli.Title)
	c := makeStatement("unused", umli.Life)

	mdl := Model{}
	mdl.Append(a)
	mdl.Append(b)
	mdl.Append(c)

	s := mdl.LifelineStatements()
	assert.Len(s, 2)
	assert.Equal(a, mdl.LifelineStatements()[0])
	assert.Equal(c, mdl.LifelineStatements()[1])
}

func TestLifelineStatementByName(t *testing.T) {
	assert := assert.New(t)

	a := makeStatement("foo", umli.Life)
	b := makeStatement("bar", umli.Title) // Right name, wrong type.
	c := makeStatement("bar", umli.Life)  // Should find this one.

	mdl := Model{}
	mdl.Append(a)
	mdl.Append(b)
	mdl.Append(c)

	s, ok := mdl.LifelineStatementByName("bar")
	assert.True(ok)
	assert.Equal(c, s)

	s, ok = mdl.LifelineStatementByName("nosuch")
	assert.False(ok)
}

func TestFirstStatementOfType(t *testing.T) {
	assert := assert.New(t)

	a := makeStatement("unused", umli.Life)
	b := makeStatement("unused", umli.Title)

	mdl := Model{}
	mdl.Append(a)
	mdl.Append(b)

	s, ok := mdl.FirstStatementOfType(umli.Title)
	assert.True(ok)
	assert.Equal(b, s)

	s, ok = mdl.FirstStatementOfType(umli.Full)
	assert.False(ok)
}

func TestLifelineIsKnown(t *testing.T) {
	assert := assert.New(t)

	a := makeStatement("foo", umli.Life)
	b := makeStatement("bar", umli.Life)

	mdl := Model{}
	mdl.Append(a)
	mdl.Append(b)

	known := mdl.LifelineIsKnown("foo")
	assert.True(known)

	known = mdl.LifelineIsKnown("unknown")
	assert.False(known)
}

func TestTextSizeWhenPresent(t *testing.T) {
	assert := assert.New(t)

	a := makeStatement("unused", umli.TextSize)
	a.TextSize = 99.0
	mdl := Model{}
	mdl.Append(a)

	textSize, ok := mdl.SizeFromTextStatement()
	assert.True(ok)
	assert.Equal(99.0, textSize)
}
func TestTextSizeWhenNotPresent(t *testing.T) {
	assert := assert.New(t)
	mdl := Model{}
	_, ok := mdl.SizeFromTextStatement()
	assert.False(ok)
}

func makeStatement(name string, statementType string) *Statement {
	s := NewStatement()
	s.Keyword = statementType
	s.LifelineName = name
	return s
}
