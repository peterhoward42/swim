# umli - UML Interaction (diagram)

This repository contains code that auto-generates a UML interaction
diagram from a very simple and human-readable text script.

For example making a diagram like this:

<https://www.tutorialspoint.com/uml/images/uml_sequence_diagram.jpg>

It is intended to create an instant-feedback, iterative and repeatable 
thinking, creation and editing process for the people creating them. 

It is hoped this makes it easy to make it up as you go along, and change your 
mind often.

Using concise plain text as input is intended to make the management of this
single source of truth for the diagrams the same as managing code (in a repo).

The input is written in a Domain-Specific-Language (DSL) like this:

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

To engage with the code organisation and design, and how it works see the
design documentation [here](docs/design.md)

The code aims to be consumable in several different ways:

- As a Go package
- As a command line program
- As a REST API (serving images or JSON/YAML serialized graphics model)
- As a WebAssembly component that can render into an HTML Canvas
- As a WebAssembly component that can render into an SVG object
- As a WebAssembly component that can render into JSON
