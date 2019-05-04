/*
Package umlinteraction is the topmost package in the uml interaction system.
It is concerned with global topics of the system, that cannot sensibly be
hidden away in more specialized packages.
*/
package umlinteraction

// These constants represent the keywords in the DSL language.
// They are in the top level package so that they can be used
// in both the parser and the dslmodel respecably.
const (
	Lane = "lane"
	Dash = "dash"
	Full = "full"
	Self = "self"
	Stop = "stop"
)

// AllKeywords provides the keywords as a list.
var AllKeywords = []string{Lane, Full, Dash, Self, Stop}
