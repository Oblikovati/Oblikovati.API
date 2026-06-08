// SPDX-License-Identifier: Apache-2.0

package wire

// This file is the JSON contract for capturing the live viewport framebuffer (the rendered 3D image)
// to a PNG file — the headless way to SEE what the renderer drew (import diagnostics, normal-debug).
// The method-name constant lives in methods.go; these are the request/response DTOs.

// CaptureViewportArgs is the request of [MethodViewportCapture]: write the active document's viewport
// framebuffer to a PNG at Path (empty ⇒ a default temp location). The capture reflects the NEXT
// rendered frame, so the host writes the file asynchronously and the caller reads it once written.
type CaptureViewportArgs struct {
	Path string `json:"path,omitempty"`
}

// CaptureViewportResult is the reply of [MethodViewportCapture]: the PNG path the host will write and
// the viewport's pixel size. The file appears within a frame of the call (poll Path / its mtime).
type CaptureViewportResult struct {
	Path   string `json:"path"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// SetNormalDebugArgs is the request of [MethodViewportSetNormalDebug]: turn the viewport's normal-debug
// render On or off — shaded triangles draw front-facing GREEN and back-facing RED, so winding /
// flipped-normal defects (hidden by normal two-sided shading) are obvious in a capture.
type SetNormalDebugArgs struct {
	On bool `json:"on"`
}

// NormalDebugResult is the reply of [MethodViewportSetNormalDebug]: the resulting state.
type NormalDebugResult struct {
	On bool `json:"on"`
}

// SetMeshColorsArgs is the request of [MethodViewportSetMeshColors]: turn the mesh-debug-colors render
// On or off — every B-rep face (or every TRIANGLE when PerTriangle) is painted a distinct, index-
// derived color, so a captured region maps back to a face/triangle index in the mesh data.
type SetMeshColorsArgs struct {
	On          bool `json:"on"`
	PerTriangle bool `json:"perTriangle,omitempty"`
}

// MeshColorsResult is the reply of [MethodViewportSetMeshColors]: the resulting state.
type MeshColorsResult struct {
	On          bool `json:"on"`
	PerTriangle bool `json:"perTriangle"`
}
