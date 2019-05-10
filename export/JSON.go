package export

import (
	"encoding/json"
	"fmt"

	"github.com/peterhoward42/umli/graphics"
)

// SerializeToJSON renders a graphics.Model into a
// JSON representation. See json_test.go for example output.
func SerializeToJSON(model *graphics.Model) ([]byte, error) {
	b, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Serialize: %v", err)
	}
	return b, nil
}
