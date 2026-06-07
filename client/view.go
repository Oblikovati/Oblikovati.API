// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati/api/types"
	"oblikovati/api/wire"
)

// View is the viewport operation group: it lets an add-in read and set the display mode
// so automations can switch the model between shaded, wireframe,
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

// Camera returns the viewport's current camera as a look-at frame (eye/target/up/fov).
//
//	cam, _ := client.View().Camera()
//	// follow a presenter: lerp eye/target toward cam, then SetCamera.
func (v View) Camera() (wire.CameraView, error) {
	var r wire.CameraView
	return r, v.c.call(wire.MethodViewGetCamera, nil, &r)
}

// SetCamera moves the viewport camera to the given look-at frame and returns the
// resulting camera (the host may normalize Up or clamp FOV).
//
//	client.View().SetCamera(wire.SetCameraArgs{Eye: e, Target: t, Up: u, FOV: f})
func (v View) SetCamera(a wire.SetCameraArgs) (wire.CameraView, error) {
	var r wire.CameraView
	return r, v.c.call(wire.MethodViewSetCamera, a, &r)
}
