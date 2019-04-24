/*
Package export provides code to render a graphics.Model instance into various formats
including image files, JSON and possibly others.
*/
package export

import (
	"encoding/json"
	"fmt"

	"github.com/peterhoward42/umlinteraction/graphics"
)

// SerializeToJSON renders a given instance of a graphics.Model into a
// JSON representation. See json_test.go for example output.
func SerializeToJSON(model *graphics.Model) ([]byte, error) {
	b, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Serialize: %v", err)
	}
	return b, nil
}
