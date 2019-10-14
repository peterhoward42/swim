package lifeline

import "github.com/peterhoward42/umli/dsl"

/*
SpanExcl works out for any two lifelines (from, to), which other lifelines lie
between them, and will thus be crossed by an interaction line between (from,to)
*/
func SpanExcl(
	from, to *dsl.Statement, allLifelines []*dsl.Statement) []*dsl.Statement {
	lifelineIndex := map[*dsl.Statement]int{}
	indexToLifeline := map[int]*dsl.Statement{}
	for i, s := range allLifelines {
		lifelineIndex[s] = i
		indexToLifeline[i] = s
	}
	fromIndex := lifelineIndex[from]
	toIndex := lifelineIndex[to]
	span := []*dsl.Statement{}
	if toIndex > fromIndex {
		for i := fromIndex + 1; i <= toIndex-1; i++ {
			span = append(span, allLifelines[i])
		}
	} else {
		for i := fromIndex - 1; i >= toIndex+1; i-- {
			span = append(span, allLifelines[i])
		}
	}
	return span
}
