/*
Package render is capable of rendering the graphics.Model(s) produced
by parser.Parser in various ways. At the time of writing it only supports
creating .png or .jpg image files. But the plans are to include rendering the
model into a JSON format, and possibly rendering into an in-memory SVG frame
or canvas, in the context of the package being compiled into a WebAssembly
componenent in a web page.
It uses the github.com/fogleman/gg 2D graphics package.
*/
package render
