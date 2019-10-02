/*
Package sizer is concerned with sizing diagram elements that are sensibly
defined in proportion to the diagrams font height. (Which is the principal
driver of sizing).

Not only how big things are, but also how far apart they should be.

E.g. the coordinates for each lifeline title box, the mark-space settings for
dashed lines, arrow sizing, the margins or padding required for each thing etc.

It is encapsulated in this dedicated package, to remove this responsibility
from the umli.diag package, so that umli.diag need only deal with the
algorithmic part of diagram creation.
*/
package sizer

// Sizer defines the contract for a thing that can provide sizes.
type Sizer interface {

	// Get returns the size specified by propertyName, or zero if that
	// size if the property is not recognized.
	Get(propertyName string) (size float64)
}
