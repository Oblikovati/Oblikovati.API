// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati/api/types"
	"oblikovati/api/wire"
)

// View is the viewport operation group: it lets an add-in read and set the display mode
// (Inventor's View.DisplayMode) so automations can switch the model between shaded, wireframe,
// realistic, hidden-edge and NPR presentations.
type View struct{ c *Client }

// View returns the viewport operation group.
func (c *Client) View() View { return View{c} }

// DisplayMode returns the viewport's current display mode and its label.
//
//	v, _ := client.View().DisplayMode()
//	if v.Mode == types.RealisticRendering { /* … */ }
func (v View) DisplayMode() (wire.DisplayModeView, error) {
	var r wire.DisplayModeView
	return r, v.c.call(wire.MethodViewGetDisplayMode, nil, &r)
}

// SetDisplayMode switches the viewport to mode, returning the resulting mode and label.
//
//	client.View().SetDisplayMode(types.WireframeWithHiddenEdgesRendering)
func (v View) SetDisplayMode(mode types.DisplayModeEnum) (wire.DisplayModeView, error) {
	var r wire.DisplayModeView
	return r, v.c.call(wire.MethodViewSetDisplayMode, wire.SetDisplayModeArgs{Mode: mode}, &r)
}

// ListDisplayModes returns every selectable display mode, flagging the active one.
func (v View) ListDisplayModes() (wire.ListDisplayModesResult, error) {
	var r wire.ListDisplayModesResult
	return r, v.c.call(wire.MethodViewListDisplayModes, nil, &r)
}
