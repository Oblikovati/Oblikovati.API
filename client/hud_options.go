// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// HUDOptions reads the in-canvas sketch input configuration: the pointer-input and
// dimension-input boxes shown while geometry is placed, and whether a typed value becomes a
// persistent driving dimension (Oblikovati/Oblikovati#2014).
//
// mcp:tool application_get_hud_options
// mcp:summary Reads the in-canvas sketch input configuration (pointer input, dimension input, and whether typed values become dimensions).
func (a Application) HUDOptions() (wire.HeadsUpDisplayOptionsView, error) {
	return call[wire.HeadsUpDisplayOptionsView](a.c, wire.MethodApplicationGetHUDOptions, nil)
}

// SetHUDOptions replaces the in-canvas sketch input configuration and returns the stored value.
// Every field is replaced, so read with [Application.HUDOptions] first to change just one.
//
// mcp:tool application_set_hud_options
// mcp:summary Replaces the in-canvas sketch input configuration (pointer input, dimension input, and whether typed values become dimensions).
func (a Application) SetHUDOptions(view wire.HeadsUpDisplayOptionsView) (wire.HeadsUpDisplayOptionsView, error) {
	return call[wire.HeadsUpDisplayOptionsView](a.c, wire.MethodApplicationSetHUDOptions, view)
}
