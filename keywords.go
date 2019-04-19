package umlinteraction

// These constants represent the keywords in the DSL language.
// They are in the top level package so that they can be used
// in both the parser and the dslmodel respecably.
const (
	Lane = "lane"
	Full = "full"
	Dash = "dash"
	Self = "self"
	Stop = "stop"
)

// AllKeywords provides the keywords as a list.
var AllKeywords = []string{Lane, Full, Dash, Self, Stop}
