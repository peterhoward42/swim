/*
Package graphics encapsulates the types that are needed to capture a diagram
in terms of low level primites (like lines, arrow heads,  and strings).

The aim is that these be kept very simple, so that they are easily 
renderable into diverse graphics formats, or serialized into JSON or YAML.

The origin is top left.
Positive is right and down.
The size units used are pixels.
The size values are float64 for three reasons:

    1. It makes it easier for all the calling code to do calculations
       in the floating point world, for accuracy and to avoid integer
       division truncation coding mistakes.
    2. It enables vector-graphics renderers (like SVG) to exploit
       floating point native rendering.
    3. It enables image renderers (which are bound to disrete pixels),
       to exploit anit-aliasing to simulate a finer resolution than is
       really present.
*/
package graphics
