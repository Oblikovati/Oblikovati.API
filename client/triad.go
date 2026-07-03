// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Triad is the move/rotate gizmo operation group (M05-F13): place the triad at a
// position/orientation and stream the user's drags as triad.drag push events
// (delta transform + drag context), with hover changes as triad.segment events.
type Triad struct{ c *Client }

// Triad returns the triad operation group.
func (c *Client) Triad() Triad { return Triad{c} }

// Show places (and shows) the triad.
//
//	client.Triad().Show(wire.TriadSpec{Position: [3]float64{0, 0, 5}, Visible: true})
//
// mcp:tool triad_show
// mcp:summary Places (and shows) the triad.
func (t Triad) Show(spec wire.TriadSpec) (wire.OKResult, error) {
	spec.Visible = true
	return call[wire.OKResult](t.c, wire.MethodTriadShow, wire.ShowTriadArgs{Triad: spec})
}

// Update repositions/reorients the visible triad.
//
// mcp:tool triad_update
// mcp:summary Repositions/reorients the visible triad.
func (t Triad) Update(spec wire.TriadSpec) (wire.OKResult, error) {
	return call[wire.OKResult](t.c, wire.MethodTriadUpdate, wire.ShowTriadArgs{Triad: spec})
}

// Hide dismisses the triad.
//
// mcp:tool triad_hide
// mcp:summary Dismisses the triad.
func (t Triad) Hide() (wire.OKResult, error) {
	return call[wire.OKResult](t.c, wire.MethodTriadHide, nil)
}

// Get returns the current triad spec (visible or not).
//
// mcp:tool triad_get
// mcp:summary Returns the current triad spec (visible or not).
func (t Triad) Get() (wire.TriadSpec, error) {
	return call[wire.TriadSpec](t.c, wire.MethodTriadGet, nil)
}

// Manipulators is the custom-gizmo operation group (M05-F13): declare drag
// handles (world hotspots over your client graphics); gestures stream back as
// manipulator.drag push events.
type Manipulators struct{ c *Client }

// Manipulators returns the custom-gizmo operation group.
func (c *Client) Manipulators() Manipulators { return Manipulators{c} }

// Set replaces one gizmo's handle set.
//
// mcp:tool manipulators_set
// mcp:summary Replaces one gizmo's handle set.
func (m Manipulators) Set(id string, handles []wire.ManipulatorHandleSpec) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodManipulatorsSet, wire.SetManipulatorsArgs{ID: id, Handles: handles})
}

// Remove dismisses a gizmo's handles.
//
// mcp:tool manipulators_remove
// mcp:summary Dismisses a gizmo's handles.
func (m Manipulators) Remove(id string) (wire.OKResult, error) {
	return call[wire.OKResult](m.c, wire.MethodManipulatorsRemove, wire.RemoveManipulatorsArgs{ID: id})
}
