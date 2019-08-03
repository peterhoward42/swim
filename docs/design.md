# Design Documentation

## Conceptual

Diagram creation has three phases handled by three separate packages:

- `parser` parses the DSL
- `diag` composes the diagram into a graphics model
- 'render' renders the graphics model - into, for example, an image file

# Package Separation of Concerns

## parser

The input to the parser is the DSL script - in the form of a string.
If the DSL script is well-formed and fit for purpose, the output is a 
sequence of `dlsmodel.Statement` structures.  These form a validated, 
structured, unambiguous, and machine-readable representation of the DSL 
script. Otherwise, an `error` is produced which is intended to be sufficient
and friendly enough to help a user as they write their script.

The parser's responsibilities are to:

- Produce the validated, machine-readable representation of the DSL script,
  in order to de-couple downstream processes from the original text form.
- Validate the DSL thoroughly, so that downstream processing can assume
  it is receiving fit-for-purpose input.
- Produce error messages that are convenient and meaningful to the human
  writing the DSL script.

## dslmodel

The type provided by `dslmodel` is `Statement`. This a structure with fields
that can represent any of the DSL line forms explicitly. For example, it has a
`Keyword` field which might contain for example the DSL `full` keyword. Nb. the
top-level `umli` package defines symbolic constants for the keywords to avoid
the proliferation of magic strings. In this case `umli.Full`.

Most of the keywords require either one or two lifelines to be referenced. For
example `stop A` requires one, while `full AB` requires two. The `Statement`
type has a field for this called `ReferencedLifelines`, which is a slice of
`*dslmodel.Statement`, i.e. pointers to the `Statement`s that define lifelines
*A* and *B*.

The DSL allows a shorthand notation for defining a multiline string label Like
this `full AB first line | second line | third`. The `Statement` however holds
each (trimmed) line segment for a label separately in its `LabelSegments`
field: a slice of strings.

## graphics

Having described the input to the `diag` package, let us next describe
it's output - the `graphics.Model` type. This is preparation for describing 
how the `diag` produces the required outputs from its inputs.

The `Model` type is a container for individual lines, arrow heads and strings
to be rendered. Each having precisely defined coordinates.

The set of primitives available in the model are delibarately minimal to make
it straightforward and less potentially ambiguous to implement faithful and
diverse renderers, and to reduce the scope of testing required for renderers
since they are completely de-coupled from the rest of the system:

- Lines
- Strings
- Filled Polygons

Important conceptual characteristics of this model are:

- Is completely decoupled from the context of any particular diagram type; it
  doesn't know what its graphical entities are for
- Assumes a top-left origin coordinate system
- Has fields for `Width` and `Height`
- Promises that all the graphics entities it holds lie within the rectangle
  formed by the origin and width/height
- `float64` coordinates - for lossless representation
- Makes no other promises about the scale or bounds of its coordinates; that is
  for renderers to worry about.
- Has no concept of colour or grayscale.

### The Line Primitive

- End points only
- No concept of line stroke width or weight
- Dashed property (boolean)
- All dashed lines have the same mark/space lengths - defined at model scope

### Text Primitive

- Simple, single-line string
- No rotation. Horizontal only.
- All strings have the same font height - defined at model scope, using
  model coordinate system units (not point size or similar)
- Position defined with origin point and a horizontal and vertical 
  justification
- No concept of typeface

### Arrow Heads

- The model does not know about arrow heads, nor line termination styles.
- However it can convey arrow heads using its filled-polygon primitive. These
  are defined only as a series of X/Y vertices.
- An implication is that renderers cannot discriminate arrow heads from any
  other filled polygons, and need not (cannot) be concerned with sizing them.

## sizer

The `sizer` package is another auxilliary package that is worth explaining
before we tackle the more complex `diag` package. It exists to minimize the
scope of the `diag` package, allowing `diag` to express only the algorithmic 
logic for creating diagrams. The `sizer` on the other hand is the single source
of truth for how big to make things, and how far to space things apart etc.

It's essentially a big lookup table for sizing psuedo-constants.

The `diag` package *will* incorporate these sizing decisions into the 
graphics it generates - but delegates the decision making of the actual sizes
to `sizer`.

The reason I described the values it holds as psuedo-constants, is that the
construction of a `Sizer` requires the caller to provide an intended diagram
width and an intended font height, and then initializes all its values as
**proportions** of those.

If you view the `Sizer` code - most of what you will see is statments like
this:

`sizer.FrameTitleTextPadT = 0.5 * fh`

This follows a pattern:

- `fh` is font height
- `Pad` is shorthand for padding
- `T` is shorthand for top

Thus `FrameTitleTextPadT` can be read as:

> The text inside the frame's title box should have some blank space reserved 
> about it, to make it look right, and using half a font height is the amount
> to use.




## diag

The `diag` package is the engine room of `umli`; where the logic lives to
decide what the diagram needs to have in it, and what it should look like.



## Diagram Synthesis Logic

## Graphics Model Responsibility

## Rendering Responsibility

## Sizer

## Test strategy

## Build, Test and Deploy


