package dslmodel

import (
	"testing"
)

var referenceInput = `
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
`

func TestBuildWithCorrectInput(t *testing.T) {
}
