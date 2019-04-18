package parser

import (
	"testing"
)

/*
Get correct num lines back.
Line types are as expected.
Sample has correct args.
Then from imple.
*/

func TestGetOffTheGround(t *testing.T) {
    p := &Parser{}
    parsed_lines, err := p.Parse("fibble")
    if err != nil {
        t.Errorf("Parse: %v", err)
    }
    if len(parsed_lines) != 14 {
        t.Errorf("Returned no lines: %v", err)
    }
}
