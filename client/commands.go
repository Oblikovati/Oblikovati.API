// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

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
