# Design Documentation

## Conceptual

Diagram creation has three phases handled by three separate packages:

- `parser` parses the DSL
- `diag` composes the diagram into a graphics model
- 'render' renders the graphics model - into, for example, an image file

# Package Separation of Concerns

## Package:parser

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

## Package: dslmodel

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

## Package: graphics

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

## Package: sizer

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

## Package: diag

The `diag` package is the engine room of `umli`; where the logic lives to
decide what the diagram needs to have in it, and what it should look like.

There are quite a few modules in the `diag` package, with `creator.go`
providing the API and orchestration of the method. The other modules each
encapsulate specific delegated responsibilities.

Diagram synthesis has a few conceptual steps:

- Composing and initializing the specialised helper objects.
- Producing the graphics (model) elements that can be thought of as being 
  anchored to the top of the diagram.
- Producing the graphics elements that arise from each line in the DSL. For
  example interaction lines and their labels, and inferring where lifeline 
  activity boxes must spring into being. Crucially this phase conceptually
  moves down the page as it goes - advancing a **tide mark**.
- Producing the graphical elements that can only be determined once the
  sequential phase above has completed. For example drawing the lifelines, now
  that we know where the bottom of the diagram is, and where they must be
  broken to avoid activity boxes and interaction lines.

### Catalogue of diag helper objects


-  `arrow.go`: knows how to work out the geometry for a polygon to
   represent an interaction line arrow head.
- `boxstate.go`: offers to keep track of when activity boxes are started
  on lifelines, and offers to draw the full box once an external party has
  decided they should be closed off.
- `events.go` advises what types of graphical elements are (or may be) 
  required in response to each type of DSL statement.
- `framemaker.go` offers to start drawing the diagram frame and its title at
  the beginning, and then finishing it off once the diagram's depth is known.
- `ilzones.go`. Short for *interaction line zones*. There is cross talk
  required between the drawing of interaction lines and their labels, and the
  drawing of lifelines, because the latter must be interrupted where there
  are potentially-clashing interaction lines crossing them. This module acts as
  the go-between and holds the information required.
- `lifelinegeomH`. Short for *lifeline geometry horizontal*. Multiple decisions
  must be made about how to position and space stuff, horizontally across the
  diagram. Mostly deciding the Y coordinates for each lifeline, and many
  other decision predicated on that. This includes some trade-off policies
  depending on how dense the diagram is with lifelines. This module takes care
  of all that.
- `lifelinemaker.go` is willing to draw lifelines near the end of the diagram
  creation process - once the information required is available. For example
  how deep were the lifeline title boxes, and where must gaps be made in them.
- `scanner.go` iterates over the `dslmodel.Statement`s sequentially,
   packaging up and emitting the required drawing event mandates - as advised
   by `events.go`.
- `segment.go` It was found during implementation that several operations
  required the concept of a one-dimensionsal extent. I.e. a thing with a
  `start`
  and an `end` coordinate value, and additionally that these be *sortable*
  and *mergeable*. That's what this module encapsulates.
  

## Test strategy

## Build, Test and Deploy


