// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
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

// Capture writes the viewport framebuffer to a PNG and returns its path and pixel size. The host
// writes the file on the next rendered frame, so read the returned Path after a short delay.
//
//	r, _ := client.View().Capture(wire.CaptureViewportArgs{Path: "/tmp/shot.png"})
//	// poll r.Path until it exists, then load the image.
func (v View) Capture(a wire.CaptureViewportArgs) (wire.CaptureViewportResult, error) {
	var r wire.CaptureViewportResult
	return r, v.c.call(wire.MethodViewportCapture, a, &r)
}

// CaptureWindow writes the WHOLE application window — the chrome (ribbon, browser, open dialogs)
// composited with the 3D viewport — to a PNG and returns its path and pixel size. Unlike Capture
// (the 3D framebuffer alone) it shows UI state. The host writes the file once the next frame
// composites, so read the returned Path after a short delay.
//
//	r, _ := client.View().CaptureWindow(wire.CaptureWindowArgs{Path: "/tmp/window.png"})
func (v View) CaptureWindow(a wire.CaptureWindowArgs) (wire.CaptureWindowResult, error) {
	var r wire.CaptureWindowResult
	return r, v.c.call(wire.MethodViewportCaptureWindow, a, &r)
}

// SetNormalDebug turns the viewport's normal-debug render on/off (front-facing green, back-facing red)
// so a capture reveals winding/flipped-normal defects.
//
//	client.View().SetNormalDebug(wire.SetNormalDebugArgs{On: true})
func (v View) SetNormalDebug(a wire.SetNormalDebugArgs) (wire.NormalDebugResult, error) {
	var r wire.NormalDebugResult
	return r, v.c.call(wire.MethodViewportSetNormalDebug, a, &r)
}

// SetMeshColors turns the mesh-debug-colors render on/off (each B-rep face — or each triangle when
// PerTriangle — a distinct color), so a capture maps a region back to a primitive index in the mesh.
//
//	client.View().SetMeshColors(wire.SetMeshColorsArgs{On: true, PerTriangle: true})
func (v View) SetMeshColors(a wire.SetMeshColorsArgs) (wire.MeshColorsResult, error) {
	var r wire.MeshColorsResult
	return r, v.c.call(wire.MethodViewportSetMeshColors, a, &r)
}
