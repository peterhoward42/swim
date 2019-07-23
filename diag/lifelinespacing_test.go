package diag

import (
	"testing"

	"github.com/peterhoward42/umli/dslmodel"
)

func TestCanConstruct(t *testing.T) {
	NewLifelineSpacing(2000, 20, []*dslmodel.Statement{}, 100, 40)
}
