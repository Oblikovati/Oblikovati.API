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
//
// mcp:tool get_display_mode
// mcp:summary Read the viewport's current display mode (visual style: shaded, wireframe, …).
func (v View) DisplayMode() (wire.DisplayModeView, error) {
	return call[wire.DisplayModeView](v.c, wire.MethodViewGetDisplayMode, nil)
}

// SetDisplayMode switches the viewport to mode, returning the resulting mode and label.
//
//	client.View().SetDisplayMode(types.WireframeWithHiddenEdgesRendering)
//
// mcp:tool set_display_mode
// mcp:summary Set the viewport's display mode (visual style) by id; see list_display_modes.
func (v View) SetDisplayMode(mode types.DisplayModeEnum) (wire.DisplayModeView, error) {
	return call[wire.DisplayModeView](v.c, wire.MethodViewSetDisplayMode, wire.SetDisplayModeArgs{Mode: mode})
}

// ListDisplayModes returns every selectable display mode, flagging the active one.
//
// mcp:tool list_display_modes
// mcp:summary List the available viewport display modes (the values set_display_mode accepts).
func (v View) ListDisplayModes() (wire.ListDisplayModesResult, error) {
	return call[wire.ListDisplayModesResult](v.c, wire.MethodViewListDisplayModes, nil)
}

// Camera returns the viewport's current camera as a look-at frame (eye/target/up/fov).
//
//	cam, _ := client.View().Camera()
//	// follow a presenter: lerp eye/target toward cam, then SetCamera.
//
// mcp:tool get_camera
// mcp:summary Read a document's active-view camera as a look-at frame (eye, target, up, fov); document 0 = active.
func (v View) Camera() (wire.CameraView, error) {
	return call[wire.CameraView](v.c, wire.MethodViewGetCamera, nil)
}

// SetCamera moves the viewport camera to the given look-at frame and returns the
// resulting camera (the host may normalize Up or clamp FOV).
//
//	client.View().SetCamera(wire.SetCameraArgs{Eye: e, Target: t, Up: u, FOV: f})
//
// mcp:tool set_camera
// mcp:summary Move a document's active-view camera to a look-at frame (eye, target, up in model units; fov radians); document 0 = active. Returns the resulting camera.
// mcp:input setCameraArg
func (v View) SetCamera(a wire.SetCameraArgs) (wire.CameraView, error) {
	return call[wire.CameraView](v.c, wire.MethodViewSetCamera, a)
}

// SetOrientation jumps the active view to a standard orientation (front/top/iso…), optionally
// fitting the model to the view, and returns the resulting camera.
//
//	client.View().SetOrientation(wire.SetOrientationArgs{Orientation: types.IsoTopRightViewOrientation, Fit: true})
//
// mcp:tool set_view_orientation
// mcp:summary Jump the active view to a standard orientation (front/top/iso…) by id; set fit to frame the model. Returns the resulting camera.
func (v View) SetOrientation(a wire.SetOrientationArgs) (wire.CameraView, error) {
	return call[wire.CameraView](v.c, wire.MethodViewSetOrientation, a)
}

// Capture writes the viewport framebuffer to a PNG and returns its path and pixel size. The host
// writes the file on the next rendered frame, so read the returned Path after a short delay.
//
//	r, _ := client.View().Capture(wire.CaptureViewportArgs{Path: "/tmp/shot.png"})
//	// poll r.Path until it exists, then load the image.
//
// mcp:tool capture_viewport
// mcp:summary Capture the live 3D viewport framebuffer and return it as an IMAGE so you can SEE exactly what the renderer drew — import results, shading, Normal-Debug (green=outward, red=back-facing). Optional path writes the PNG to a host file; otherwise a temp file is used.
// mcp:image
func (v View) Capture(a wire.CaptureViewportArgs) (wire.CaptureViewportResult, error) {
	return call[wire.CaptureViewportResult](v.c, wire.MethodViewportCapture, a)
}

// CaptureWindow writes the WHOLE application window — the chrome (ribbon, browser, open dialogs)
// composited with the 3D viewport — to a PNG and returns its path and pixel size. Unlike Capture
// (the 3D framebuffer alone) it shows UI state. The host writes the file once the next frame
// composites, so read the returned Path after a short delay.
//
//	r, _ := client.View().CaptureWindow(wire.CaptureWindowArgs{Path: "/tmp/window.png"})
//
// mcp:tool capture_window
// mcp:summary Capture the WHOLE application window — the ribbon, browser, any open dialog, and the 3D viewport, exactly as the user sees it — and return it as an IMAGE. Use this to SEE UI state (e.g. whether a dialog is open, what a panel shows); use capture_viewport for the 3D render alone. Optional path writes the PNG to a host file; otherwise a temp file is used.
// mcp:image
func (v View) CaptureWindow(a wire.CaptureWindowArgs) (wire.CaptureWindowResult, error) {
	return call[wire.CaptureWindowResult](v.c, wire.MethodViewportCaptureWindow, a)
}

// Click clicks in the viewport, exactly as a user's mouse does: the running command consumes it,
// or it becomes a selection pick. Give the position as viewport pixels, or as a model-space Point
// the host projects for you — the latter needs no knowledge of the camera.
//
// This is the only way to reach behaviour that lives in the INPUT path rather than the model:
// which constraints a tool infers from where you clicked, what it previews between clicks, how a
// multi-click chain builds. Creating the same geometry with sketch.addEntity bypasses all of it.
//
//	// Draw a rectangle from an existing corner, so it picks up the coincidence.
//	client.View().Click(wire.ClickViewportArgs{Point: &types.Point{X: 6}})
//	client.View().Click(wire.ClickViewportArgs{Point: &types.Point{X: 11, Y: 4}})
//
// mcp:tool viewport_click
// mcp:input clickViewportArg
// mcp:summary Click in the 3D viewport like a user, driving the ACTIVE command (start it first with execute_command) or making a selection. Give point [x,y,z] in model space (easiest — the host projects it) or x/y in viewport pixels. Returns the pixel clicked and the command still running ("" once it finished). Use this to exercise interactive behaviour — constraint inference, previews, multi-click chains — that creating geometry directly cannot reach.
func (v View) Click(a wire.ClickViewportArgs) (wire.ClickViewportResult, error) {
	return call[wire.ClickViewportResult](v.c, wire.MethodViewportClick, a)
}

// Scroll delivers a mouse-wheel tick to the viewport (#1822): DY zooms toward the cursor as the
// vertical wheel always has; DX is the horizontal axis a two-axis wheel reports, carried but unbound.
//
//	client.View().Scroll(wire.ScrollViewportArgs{DY: 1}) // one notch toward the cursor
//
// mcp:tool viewport_scroll
// mcp:input scrollViewportArg
// mcp:summary Scroll the mouse wheel in the 3D viewport. dy zooms toward the cursor (positive = in), dx is the horizontal wheel axis (delivered but unbound). x/y locate the cursor (viewport centre if omitted). Returns the command still running ("" if none).
func (v View) Scroll(a wire.ScrollViewportArgs) (wire.ScrollViewportResult, error) {
	return call[wire.ScrollViewportResult](v.c, wire.MethodViewportScroll, a)
}

// PressKey delivers a key to the running command — "Escape" and "Enter" being the two that end a
// variable-length one (a continuous line chain, a spline) and keep what it has drawn.
//
//	client.View().PressKey(wire.PressKeyArgs{Key: "Escape"}) // finish the chain
//
// mcp:tool viewport_key
// mcp:summary Press a key in the viewport, delivered to the ACTIVE command. "Escape" or "Enter" finishes a variable-length command (a continuous line chain, a spline), keeping what it drew; "Delete" removes the selection. Returns the command still running ("" once the key ended it).
func (v View) PressKey(a wire.PressKeyArgs) (wire.PressKeyResult, error) {
	return call[wire.PressKeyResult](v.c, wire.MethodViewportKey, a)
}

// SetNormalDebug turns the viewport's normal-debug render on/off (front-facing green, back-facing red)
// so a capture reveals winding/flipped-normal defects.
//
//	client.View().SetNormalDebug(wire.SetNormalDebugArgs{On: true})
//
// mcp:tool set_normal_debug
// mcp:summary Turn the viewport's normal-debug render on/off: shaded triangles draw front-facing GREEN and back-facing RED, so capture_viewport reveals winding / flipped-normal defects.
func (v View) SetNormalDebug(a wire.SetNormalDebugArgs) (wire.NormalDebugResult, error) {
	return call[wire.NormalDebugResult](v.c, wire.MethodViewportSetNormalDebug, a)
}

// SetMeshColors turns the mesh-debug-colors render on/off (each B-rep face — or each triangle when
// PerTriangle — a distinct color), so a capture maps a region back to a primitive index in the mesh.
//
//	client.View().SetMeshColors(wire.SetMeshColorsArgs{On: true, PerTriangle: true})
//
// mcp:tool set_mesh_colors
// mcp:summary Turn the mesh-debug-colors render on/off: every B-rep face — or every TRIANGLE when perTriangle:true — is painted a distinct color, so capture_viewport lets you map a region back to a face/triangle index in the mesh data.
func (v View) SetMeshColors(a wire.SetMeshColorsArgs) (wire.MeshColorsResult, error) {
	return call[wire.MeshColorsResult](v.c, wire.MethodViewportSetMeshColors, a)
}
