# umli - UML Interaction (diagram)

This repository will contain code that can produce a UMLI interaction diagram
from a simple script in a domain specific language.

For example making this diagram:

<https://www.tutorialspoint.com/uml/images/uml_sequence_diagram.jpg>

From this script:

    lane A  SL App
    lane B  Core Permissions API
    lane C  SL Admin API | edit_facilities | endpoint

    full AC  edit_facilities( | payload, user_token)
    full CB  get_user_permissions( | token)
    dash BC  permissions_list
    stop B
    self C   [has EDIT_FACILITIES permission] | store changes etc
    dash CA  status_ok, payload
    self C   [no permission]
    dash CA  status_not_authorized

> todo: change the diagram example to be one that matches the script.

The output is a very simple graphics model **umli/graphics/model.Model** which
is little more than a list of lines and strings, along with the coordinates
at which they should be rendered to produce the diagram.

A separate repository provides renderers for the graphics model into
a variety of image and serialization formats. See:
<https://github.com/peterhoward42/umli-export>

The code aims to be consumable in several different ways:

- As a Go package
- As a command line program
- As a REST API (serving images or JSON/YAML serialized graphics model)
- As a WebAssembly component that can render into an HTML Canvas
- As a WebAssembly component that can render into an SVG object
- As a WebAssembly component that can render into JSON
