// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// AddIns is the add-in registry operation group — the ApplicationAddIns equivalent:
// enumerate what is installed, drive the activate/deactivate lifecycle, persist load
// behavior, and reach another add-in's automation surface.
type AddIns struct{ c *Client }

// AddIns returns the add-in registry operation group.
func (c *Client) AddIns() AddIns { return AddIns{c} }

// List returns every registered add-in with its manifest identity and runtime state.
//
//	for _, a := range mustList(client.AddIns().List()).AddIns { fmt.Println(a.ID, a.Activated) }
//
// mcp:tool addins_list
// mcp:summary Returns every registered add-in with its manifest identity and runtime state.
func (a AddIns) List() (wire.ListAddInsResult, error) {
	return call[wire.ListAddInsResult](a.c, wire.MethodAddInsList, nil)
}

// Get returns one registry entry by add-in id.
//
// mcp:tool addins_get
// mcp:summary Returns one registry entry by add-in id.
func (a AddIns) Get(id string) (wire.AddInInfo, error) {
	return call[wire.AddInInfo](a.c, wire.MethodAddInsGet, wire.AddInRefArgs{ID: id})
}

// Activate runs the add-in's activation (a no-op if it is already active). It fails
// for an add-in whose load behavior is LoadDisabled.
//
// mcp:tool addins_activate
// mcp:summary Runs the add-in's activation (a no-op if it is already active).
func (a AddIns) Activate(id string) (wire.OKResult, error) {
	return call[wire.OKResult](a.c, wire.MethodAddInsActivate, wire.AddInRefArgs{ID: id})
}

// Deactivate runs the add-in's shutdown (a no-op if it is not active).
//
// mcp:tool addins_deactivate
// mcp:summary Runs the add-in's shutdown (a no-op if it is not active).
func (a AddIns) Deactivate(id string) (wire.OKResult, error) {
	return call[wire.OKResult](a.c, wire.MethodAddInsDeactivate, wire.AddInRefArgs{ID: id})
}

// SetLoadBehavior persists when the host activates the add-in on future startups.
//
// mcp:tool addins_set_load_behavior
// mcp:summary Persists when the host activates the add-in on future startups.
func (a AddIns) SetLoadBehavior(id string, b types.AddInLoadBehavior) (wire.OKResult, error) {
	args := wire.SetAddInLoadBehaviorArgs{ID: id, LoadBehavior: b}
	return call[wire.OKResult](a.c, wire.MethodAddInsSetLoadBehavior, args)
}

// CallAutomation invokes a method on another add-in's automation surface
// (ApplicationAddIn.Automation) and returns its opaque JSON reply.
//
//	out, _ := client.AddIns().CallAutomation("com.example.solver", "solve", json.RawMessage(`{"n":3}`))
//
// mcp:tool addins_call_automation
// mcp:summary Invokes a method on another add-in's automation surface (ApplicationAddIn.Automation) and returns its opaque JSON reply.
func (a AddIns) CallAutomation(id, method string, args json.RawMessage) (json.RawMessage, error) {
	req := wire.CallAddInAutomationArgs{ID: id, Method: method, Args: args}
	r, err := call[wire.CallAddInAutomationResult](a.c, wire.MethodAddInsCallAutomation, req)
	if err != nil {
		return nil, err
	}
	return r.Result, nil
}
