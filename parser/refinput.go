package parser

// ReferenceInput is a convenient example DSL script for testing.
const ReferenceInput = `
title This is the | Working Title
textsize 10
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
`
