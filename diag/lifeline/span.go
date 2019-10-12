package lifeline

import "github.com/peterhoward42/umli/dsl"

/*
SpanExcl is concerned with the lifelines that cross an interaction
line. Specifically, when you provide the lifelines from which an interaction
line goes to and from, it provides the lifelines that lie between them.
*/
func SpanExcl(
	from, to *dsl.Statement, allLifelines []*dsl.Statement) []*dsl.Statement {
	sToI := map[*dsl.Statement]int{}
	iToS := map[int]*dsl.Statement{}
	for i, s := range allLifelines {
		sToI[s] = i
		iToS[i] = s
	}
	fromI := sToI[from]
	toI := sToI[to]
	span := []*dsl.Statement{}
	if toI > fromI {
		for i := fromI + 1; i <= toI-1; i++ {
			span = append(span, allLifelines[i])
		}
	} else {
		for i := fromI - 1; i >= toI+1; i-- {
			span = append(span, allLifelines[i])
		}
	}
	return span
}
