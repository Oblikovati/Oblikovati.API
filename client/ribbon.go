// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Ribbon is the ribbon operation group — discovery of the active ribbon's structure, so an
// add-in knows the tab/panel internal names to place its controls into. Placement itself is
// done with Commands().Create (the Ribbon/Tab/Category/Environment fields).
type Ribbon struct{ c *Client }

// Ribbon returns the ribbon operation group.
func (c *Client) Ribbon() Ribbon { return Ribbon{c} }

// List returns the ribbon currently shown for the active document (ZeroDoc when none is open),
// with its tabs, panels, and controls — the discovery surface for inserting add-in buttons.
//
// mcp:tool get_ribbon
// mcp:summary Read the ribbon shown for the active document (ZeroDoc when none open): its tabs, panels, and controls — the names to place add-in buttons into.
func (rb Ribbon) List() (wire.ListRibbonResult, error) {
	var r wire.ListRibbonResult
	return r, rb.c.call(wire.MethodRibbonList, nil, &r)
}

// Environments returns the UI environments the command framework scopes by (base,
// sketch, …), flagging the active one — so an add-in placing contextual commands
// knows which contexts exist (add-in-created environments: Oblikovati#667).
//
// mcp:tool ui_list_environments
// mcp:summary Returns the UI environments the command framework scopes by (base, sketch, …), flagging the active one — so an add-in placing contextual commands knows which contexts exist (add-in-created environments: Oblikovati#667).
func (rb Ribbon) Environments() (wire.ListEnvironmentsResult, error) {
	var r wire.ListEnvironmentsResult
	return r, rb.c.call(wire.MethodUIListEnvironments, nil, &r)
}
