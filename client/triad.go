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
func (t Triad) Show(spec wire.TriadSpec) (wire.OKResult, error) {
	var r wire.OKResult
	spec.Visible = true
	return r, t.c.call(wire.MethodTriadShow, wire.ShowTriadArgs{Triad: spec}, &r)
}

// Update repositions/reorients the visible triad.
func (t Triad) Update(spec wire.TriadSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, t.c.call(wire.MethodTriadUpdate, wire.ShowTriadArgs{Triad: spec}, &r)
}

// Hide dismisses the triad.
func (t Triad) Hide() (wire.OKResult, error) {
	var r wire.OKResult
	return r, t.c.call(wire.MethodTriadHide, nil, &r)
}

// Get returns the current triad spec (visible or not).
func (t Triad) Get() (wire.TriadSpec, error) {
	var r wire.TriadSpec
	return r, t.c.call(wire.MethodTriadGet, nil, &r)
}

// Manipulators is the custom-gizmo operation group (M05-F13): declare drag
// handles (world hotspots over your client graphics); gestures stream back as
// manipulator.drag push events.
type Manipulators struct{ c *Client }

// Manipulators returns the custom-gizmo operation group.
func (c *Client) Manipulators() Manipulators { return Manipulators{c} }

// Set replaces one gizmo's handle set.
func (m Manipulators) Set(id string, handles []wire.ManipulatorHandleSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodManipulatorsSet, wire.SetManipulatorsArgs{ID: id, Handles: handles}, &r)
}

// Remove dismisses a gizmo's handles.
func (m Manipulators) Remove(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, m.c.call(wire.MethodManipulatorsRemove, wire.RemoveManipulatorsArgs{ID: id}, &r)
}
