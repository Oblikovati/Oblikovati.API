// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Commands is the command operation group.
type Commands struct{ c *Client }

// Commands returns the command operation group.
func (c *Client) Commands() Commands { return Commands{c} }

// List returns every registered command and whether it can run now.
func (cm Commands) List() (wire.ListCommandsResult, error) {
	var r wire.ListCommandsResult
	return r, cm.c.call(wire.MethodCommandsList, nil, &r)
}

// Execute runs the command with the given id (the same path a ribbon click takes).
func (cm Commands) Execute(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, cm.c.call(wire.MethodCommandsExecute, wire.ExecuteCommandArgs{ID: id}, &r)
}

// Create registers a new ribbon button so an add-in can extend the UI. The button appears in the ribbon immediately; when the user clicks it
// the host fires a command-ended event the add-in receives via its Notify entry point,
// where it runs the button's action (typically further client calls).
func (cm Commands) Create(args wire.CreateCommandArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, cm.c.call(wire.MethodCommandsCreate, args, &r)
}

// SetState updates a command's live ribbon state: Active toggles its highlighted (accent)
// look, and a non-empty DisplayName relabels it. Use it for stateful add-in controls.
//
//	client.Commands().SetState(wire.SetCommandStateArgs{ID: id, Active: true, DisplayName: "Presenting"})
func (cm Commands) SetState(args wire.SetCommandStateArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, cm.c.call(wire.MethodCommandsSetState, args, &r)
}
