package umli

import "fmt"

// DSLError builds an error that describes a fault in some input DSL,
// and which includes the original faulty line and its line number.
func DSLError(line string, lineNo int, msg string) error {
	return fmt.Errorf("Error on this line <%s> (line: %v): %s",
		line, lineNo, msg)
}
