/*
Package umli is the topmost package in the uml interaction system.
It is concerned with global topics of the system, that cannot sensibly be
hidden away in more specialized packages.
*/
package umli

// These constants represent the keywords in the DSL language.
// They are in the top level package so that they can be used
// in both the parser and the dslmodel.
const (
	Title = "title"
    TextSize = "textsize"
	Life  = "life"
	Dash  = "dash"
	Full  = "full"
	Self  = "self"
	Stop  = "stop"
)

// AllKeywords provides the keywords as a list.
var AllKeywords = []string{Title, Life, Full, Dash, Self, Stop, TextSize}

// KnownKeyword returns true if the given keyword is a recognized one.
func KnownKeyword(keyWord string) bool  {
    for _,  known := range AllKeywords {
        if keyWord == known  {
            return true
        }
    }
    return false
}
