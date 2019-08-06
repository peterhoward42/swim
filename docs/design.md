# Design Documentation

## Conceptual

Diagram creation has three phases orchestrated by three packages:

- `parser` parses the DSL
- `diag` composes the diagram into a graphics model
- `render` renders the graphics model - into, for example, an image file

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
type has a field for this called `ReferencedLifelines`. Rather than storing the
lifeline letters - which would require another downstream look up - it stores 
direct pointers to the `dslmodel.Statement`s that created the referenced 
lifelines.

The DSL allows a shorthand notation for defining a multiline string label Like
this `full AB first line | second line | third`. The `Statement` however holds
each (trimmed) line segment for a label separately in its `LabelSegments`
field: a slice of strings.

## Package: graphics

Having described the input to the `diag` package, let us next describe
its output - the `graphics.Model` type. This is in preparation for describing 
how the `diag` produces the required outputs from its inputs.

The `Model` type is a container for individual lines, arrow heads and strings,
with precisely defined coordinates.

The set of primitives available in the model are delibarately minimal to make
it straightforward and less potentially ambiguous to implement faithful and
diverse renderers, and to reduce the scope of testing required for renderers
since they are completely de-coupled from the rest of the system:

- Lines
- Strings
- Filled Polygons

Significant conceptual characteristics of this model are:

- Is completely decoupled from the context of any particular diagram type; it
  doesn't know what its graphical entities are for
- Assumes a top-left origin coordinate system
- Has fields for `Width` and `Height`
- Promises that all the graphics entities it holds lie within the rectangle
  formed by the origin and width/height
- `float64` coordinates - for lossless representation
- Is oblivious to the coordinate systems and scale that renderers might favour;
  renderers should assume that the coordinates must be transformed
- Has no concept of colour or grayscale.

### The Line Primitive

- End points only
- No concept of line stroke width or weight
- May be **dashed**; carried by a boolean property
- All dashed lines have the same mark/space lengths - defined at model scope

### Text Primitive

- Simple, single-line string
- No rotation. Horizontal only.
- All strings have the same font height - defined at model scope, using
  model length units (not point-size or similar)
- Position defined with origin point and a horizontal and vertical 
  justification `{left, right, centre, top, bottom}`.
- No concept of typeface; renderers can choose, but must respect font height

### Arrow Heads

- The model does not know about arrow heads, nor line termination styles.
- However it can convey arrow heads using its filled-polygon primitive. These
  are defined only as a series of X/Y vertices.
- An implication is that renderers cannot discriminate arrow heads from any
  other filled polygons, and need not (cannot) be concerned with sizing them.
  This design ensures that the umli system can control completely the
  suitability of arrow head size and aspect ratio in relation to everything 
  else, deterministically.

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

In the `Sizer` code - there are many statments like this:

`sizer.FrameTitleTextPadT = 0.5 * fh`

This follows a pattern:

- `fh` is font height
- `Pad` is shorthand for padding
- `T` is shorthand for top

Thus the statement can be read as:

> The text inside the frame's title box should have some blank space reserved 
> above it, to make it look right, and using half a font height is the amount
> to use.

## Package: diag

The `diag` package is the engine room of `umli`; where the logic lives to
decide what the diagram needs to have in it, and what it should look like.

There are quite a few modules in the `diag` package, with `creator.go`
providing the API and orchestration of the method. The other modules each
encapsulate specific delegated responsibilities.

The diagram synthesis algorithm has these conceptual steps:

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
  broken to avoid clashing with activity boxes and interaction lines.

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
  the intermediary, and holds the information required.
- `lifelinegeomH`. Short for *lifeline geometry horizontal*. Multiple decisions
  must be made about how to position and space graphics **horizontally** across
  the diagram. Mostly deciding the X coordinates for each lifeline, and many
  other decisions predicated on that. This includes some trade-off policies
  depending on how dense the diagram is with lifelines. This module takes care
  of all that.
- `lifelinemaker.go` is willing to draw lifelines near the end of the diagram
  creation process - once the information required is available. For example
  the resultant height of the diagram, how tall the lifeline title boxes, 
  had to be to fit the label with the most lines, and where lifelines must be
  broken to not clash with something else. 
- `scanner.go` iterates over the `dslmodel.Statement`s sequentially,
   packaging up and emitting the required drawing event mandates - as advised
   by `events.go`.
- `segment.go` It was found during implementation that several operations
  required the concept of a one-dimensionsal extent. I.e. a thing with a
  `start` and an `end` coordinate value, and additionally that these be
  *sortable* and *mergeable*. That's what this module encapsulates.

### The `diag` coordinate system

Previous sections have explained that the `diag` package's output is a graphics
model, and have defined the expectations and contract for the model's
coordinate system. In particular, that renderers should make no assumptions
that it is scaled conveniently for their purposes.

Hence the `diag` package is free to adopt a scale of its choice. It chooses to
always work on the basis that the diagram is 2000 units wide. This is an
internal private decision to the package, and is not part of its contract or
API. The design intent is that this could be changed without requiring any
changes elsewhere in the system.

It establishes the text height to use from the DSL's `textheight` statement if
present. This use of 20 in this text height statement produces very large text:

    `textheight 20`

It specifies that the text height should be the diagram's width multiplied by 
0.020. I.e a ratio of 1/50. In the absence of a `textheight` statement, it 
defaults to an ideal value for typical diagrams of 1/100, as if the DSL had:

    `textheight 10`

The smallest advisable text is produced by:

    `textheight 5`

The DSL parser does not allow values lower than 5 or greater than 20.

The internal diagram width of 2000 is convenient for development and debugging,
being a human-relatable size when the units are considered to be pixels. It
also supports early-stage naive pixel-based renderers that don't bother to 
scale the models.

## Package: render

The `render` package is responsible for consuming a graphics model and
producing something that can be visualised. At the time of writing, it is
capable only of producing *.png* and *.jpg* image files, but is intended to
be expanded to multiple other formats.

These will include rendering operations that only make sense in the context of
a web page - when the system has been compiled as a Web Assembly component.
For example rendering directly into an HTML canvas or SVG object. 

It will also encompass rendering that is more in the style of serialization.
For example to JSON or YAML formats. Thus opening out additional rendering,
or architectural choices to third party code.
  

## Test strategy

The test design strategy is unorthodox. Where conventional unit tests fit
easily they are of course used. For example the parser is thoroughly unit
tested.

But the condundrum is that a human being looking at rendered diagrams that
exhibit representative and telling use-cases, is too valuable and efficient 
to be ignored. Thus the diag package is evolving a suite of visual tests that
produce image files. During development these are used for human inspection 
of the diagrams created, where errors or the need for improvements is
immediately apparent.

It is inconveniently time consuming and fragile to create unit tests to check
for the same things.  Instead, once the visual unit test suite is complete, the
golden-reference image files will be commited to the repo, and automated tests
will perform regression tests against these. This idea is gratefully copied
from [JEST's snapshot feature](https://jestjs.io/docs/en/snapshot-testing).
There is already an image comparison package ready (not yet in this repo),
which is capable of comparing images pixel for pixel programmatically. This is
necessary because a byte for byte comparison of the files produces false
negatives in the presence of metadata inside the image files - for example
timestamps.


## Build, Test, Release and Deploy

todo
