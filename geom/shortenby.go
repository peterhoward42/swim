package geom

/*
ShortenLineBy modifies the horizontal line between x1 and x2 by shortening it by shortenBy
at both ends. It does not matter if the line is going from left to right or
right to left.
*/
func ShortenLineBy(shortenBy float64, x1 *float64, x2 *float64) {
	if *x2 > *x1 {
		*x1 += shortenBy
		*x2 -= shortenBy
		return
	}
	*x1 -= shortenBy
	*x2 += shortenBy
}
