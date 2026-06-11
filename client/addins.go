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
func (a AddIns) List() (wire.ListAddInsResult, error) {
	var r wire.ListAddInsResult
	return r, a.c.call(wire.MethodAddInsList, nil, &r)
}

// Get returns one registry entry by add-in id.
func (a AddIns) Get(id string) (wire.AddInInfo, error) {
	var r wire.AddInInfo
	return r, a.c.call(wire.MethodAddInsGet, wire.AddInRefArgs{ID: id}, &r)
}

// Activate runs the add-in's activation (a no-op if it is already active). It fails
// for an add-in whose load behavior is LoadDisabled.
func (a AddIns) Activate(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, a.c.call(wire.MethodAddInsActivate, wire.AddInRefArgs{ID: id}, &r)
}

// Deactivate runs the add-in's shutdown (a no-op if it is not active).
func (a AddIns) Deactivate(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, a.c.call(wire.MethodAddInsDeactivate, wire.AddInRefArgs{ID: id}, &r)
}

// SetLoadBehavior persists when the host activates the add-in on future startups.
func (a AddIns) SetLoadBehavior(id string, b types.AddInLoadBehavior) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.SetAddInLoadBehaviorArgs{ID: id, LoadBehavior: b}
	return r, a.c.call(wire.MethodAddInsSetLoadBehavior, args, &r)
}

// CallAutomation invokes a method on another add-in's automation surface
// (ApplicationAddIn.Automation) and returns its opaque JSON reply.
//
//	out, _ := client.AddIns().CallAutomation("com.example.solver", "solve", json.RawMessage(`{"n":3}`))
func (a AddIns) CallAutomation(id, method string, args json.RawMessage) (json.RawMessage, error) {
	var r wire.CallAddInAutomationResult
	req := wire.CallAddInAutomationArgs{ID: id, Method: method, Args: args}
	if err := a.c.call(wire.MethodAddInsCallAutomation, req, &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}
