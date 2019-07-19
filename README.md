# umli - UML Interaction (diagram)

This repository contains code that auto-generates a UML interaction
diagram from a very simple text script.

It is intended to create a fast and iterative thinking and creation process
for the people creating them.

For example making this diagram:

<https://www.tutorialspoint.com/uml/images/uml_sequence_diagram.jpg>

The script is written in a Domain-Specific-Language (DSL) like this:

    life A  SL App
    life B  Core Permissions API
    life C  SL Admin API | edit_facilities | endpoint

    full AC  edit_facilities( | payload, user_token)
    full CB  get_user_permissions( | token)
    dash BC  permissions_list
    stop B
    self C   [has EDIT_FACILITIES permission] | store changes etc
    dash CA  status_ok, payload
    self C   [no permission]
    dash CA  status_not_authorized

> todo: change the diagram example to be one that matches the script.

Creating the diagram has two conceptual phases. The first being to parse
the DSL and build a `graphics.Model` that is little more than a list of lines, 
strings and arrow heads, along with their coordinates, which when rendered, 
will  produce the diagram.

The second phase is to render this model. The `render` package provides
renderes for:

- .png image file
- .jpg image file
- a JSON representation
- a YAML representation
- an HTML canvas (See Note ***)

The code aims to be consumable in several different ways:

- As a Go package
- As a command line program
- As a REST API (serving images or JSON/YAML serialized graphics model)
- As a WebAssembly component that can render into an HTML Canvas
- As a WebAssembly component that can render into an SVG object
- As a WebAssembly component that can render into JSON
