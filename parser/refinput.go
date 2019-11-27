package parser

// ReferenceInput is a convenient example DSL script.
const ReferenceInput = `
title Authentication

life A  Client
life B  API Server
life C  User Store

full AB API Request | With User/Passwd
full BC Validate user/pass
dash CB Valid?
stop C
self B Handle Request
full BA API Response
`
