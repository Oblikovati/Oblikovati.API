// SPDX-License-Identifier: Apache-2.0

package wire

// WorkSurface DTOs (M20·F16). A work surface is a named, visibility-controlled
// construction surface produced by a surface-output feature (a surface extrude/revolve/
// loft/sweep, a boundary patch, a knit/stitch, a ruled or trim/extend result). It wraps
// one or more surface bodies and is addressable as a feature input (split tool,
// replace-face target, to-face extent) and as a sketch/work-feature reference. Modeled
// like a work plane (object + collection + proxy), not as a feature triangle: instances
// are produced by features, so there is no create method — only list/get/setVisible/rename.

// WorkSurfaceInfo is one row of [MethodWorkSurfacesList] and the payload of
// [MethodWorkSurfacesGet]: a construction surface's identity, display state, the count of
// surface bodies it wraps, the feature that produced it, and a stable reference for
// consuming it as a feature input.
type WorkSurfaceInfo struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Ref         string `json:"ref"`
	Visible     bool   `json:"visible"`
	Translucent bool   `json:"translucent"`
	Bodies      int    `json:"bodies"`           // number of surface bodies wrapped
	Source      string `json:"source,omitempty"` // name of the feature that produced it
}

// ListWorkSurfacesResult is the response of [MethodWorkSurfacesList].
type ListWorkSurfacesResult struct {
	Surfaces []WorkSurfaceInfo `json:"surfaces"`
}

// WorkSurfaceRefArgs is the request of [MethodWorkSurfacesGet]: one work surface by its
// Index in the collection (from [MethodWorkSurfacesList]).
type WorkSurfaceRefArgs struct {
	Index int `json:"index"`
}

// WorkSurfaceDetailResult is the response of [MethodWorkSurfacesGet],
// [MethodWorkSurfacesSetVisible], and [MethodWorkSurfacesRename]: the surface's refreshed
// state after the call.
type WorkSurfaceDetailResult struct {
	Surface WorkSurfaceInfo `json:"surface"`
}

// SetWorkSurfaceVisibleArgs is the request of [MethodWorkSurfacesSetVisible]: set (not
// toggle) the surface's visibility, so the call is idempotent for replication.
type SetWorkSurfaceVisibleArgs struct {
	Index   int  `json:"index"`
	Visible bool `json:"visible"`
}

// RenameWorkSurfaceArgs is the request of [MethodWorkSurfacesRename]. The new name must
// be non-empty and unique within the part.
type RenameWorkSurfaceArgs struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}
