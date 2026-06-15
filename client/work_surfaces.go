// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// WorkSurfaces is the construction-surface group for the active part: List enumerates the
// surface bodies surface-output features produced (each a named, visibility-controlled
// WorkSurface), and the helpers get / show-or-hide / rename one. Work surfaces are created
// as a side effect of surface features (there is no create here); a surface's Ref from
// List can be passed wherever a feature accepts a surface input.
type WorkSurfaces struct{ c *Client }

// WorkSurfaces returns the construction-surface group.
func (c *Client) WorkSurfaces() WorkSurfaces { return WorkSurfaces{c} }

// List returns the active part's work surfaces (construction surfaces produced by
// surface-output features), in creation order.
//
// mcp:tool list_work_surfaces
// mcp:summary List the work surfaces (construction surfaces) of the active part. Each reports its name, ref, visibility, translucency, the count of surface bodies it wraps, and the feature that produced it.
func (w WorkSurfaces) List() (wire.ListWorkSurfacesResult, error) {
	var r wire.ListWorkSurfacesResult
	return r, w.c.call(wire.MethodWorkSurfacesList, nil, &r)
}

// Get returns one work surface's state by its index (from List), e.g. Get(0).
//
// mcp:tool work_surface_get
// mcp:summary Returns one work surface's state (name, ref, visibility, body count, source feature) by its index from list_work_surfaces.
func (w WorkSurfaces) Get(index int) (wire.WorkSurfaceDetailResult, error) {
	var r wire.WorkSurfaceDetailResult
	return r, w.c.call(wire.MethodWorkSurfacesGet, wire.WorkSurfaceRefArgs{Index: index}, &r)
}

// SetVisible shows or hides the work surface at index and returns its refreshed state,
// e.g. SetVisible(0, false).
//
// mcp:tool work_surface_set_visible
// mcp:summary Show or hide a work surface by its index (from list_work_surfaces). Returns the surface's refreshed state.
func (w WorkSurfaces) SetVisible(index int, visible bool) (wire.WorkSurfaceDetailResult, error) {
	var r wire.WorkSurfaceDetailResult
	args := wire.SetWorkSurfaceVisibleArgs{Index: index, Visible: visible}
	return r, w.c.call(wire.MethodWorkSurfacesSetVisible, args, &r)
}

// Rename sets the work surface's display name (must be non-empty and unique), e.g.
// Rename(0, "Parting Surface").
//
// mcp:tool work_surface_rename
// mcp:summary Sets a work surface's display name by its index (from list_work_surfaces); the name must be non-empty and unique within the part.
func (w WorkSurfaces) Rename(index int, name string) (wire.WorkSurfaceDetailResult, error) {
	var r wire.WorkSurfaceDetailResult
	args := wire.RenameWorkSurfaceArgs{Index: index, Name: name}
	return r, w.c.call(wire.MethodWorkSurfacesRename, args, &r)
}
