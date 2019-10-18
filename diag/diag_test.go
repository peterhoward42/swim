package diag

import (
	"testing"

	"github.com/peterhoward42/umli/dsl"
	"github.com/stretchr/testify/assert"
)

func TestCreateRunsWithoutCrashing(t *testing.T) {
	assert := assert.New(t)
	cr, err := NewCreator()
	assert.NoError(err)
	var dslModel dsl.Model
	graphicsModel, err := cr.Create(dslModel)
	_ = graphicsModel
	assert.NoError(err)
}
