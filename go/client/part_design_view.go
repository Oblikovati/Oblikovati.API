// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// PartDesignViews provides the active part's design representations.
// Example: `views := client.PartDesignViews()`.
type PartDesignViews struct{ c *Client }

// PartDesignViews returns the part design-view method group.
// Example: `views := client.PartDesignViews()`.
func (c *Client) PartDesignViews() PartDesignViews { return PartDesignViews{c} }

// List reads the active part's design views in creation order.
// Example: `result, err := client.PartDesignViews().List()`.
//
// mcp:tool list_part_design_views
// mcp:summary List the active part's design views in creation order.
func (p PartDesignViews) List() (wire.PartDesignViewsResult, error) {
	return call[wire.PartDesignViewsResult](p.c, wire.MethodPartDesignViewsList, struct{}{})
}

// Copy creates an inactive copy of a part view.
// Example: `result, err := client.PartDesignViews().Copy(wire.CopyPartDesignViewArgs{SourceID: 1})`.
//
// mcp:tool copy_part_design_view
// mcp:summary Create an inactive copy of a part design view.
func (p PartDesignViews) Copy(args wire.CopyPartDesignViewArgs) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsCopy, args)
}

// Rename changes an authored part view's name without changing its identity.
// Example: `result, err := client.PartDesignViews().Rename(wire.RenamePartDesignViewArgs{ID: 2, Name: "Review"})`.
//
// mcp:tool rename_part_design_view
// mcp:summary Rename an authored part design view by ID.
func (p PartDesignViews) Rename(args wire.RenamePartDesignViewArgs) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsRename, args)
}

// Activate applies the stored display state of a part view.
// Example: `result, err := client.PartDesignViews().Activate(2)`.
//
// mcp:tool activate_part_design_view
// mcp:summary Activate a part design view and apply its stored display state.
func (p PartDesignViews) Activate(id uint64) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsActivate, wire.PartDesignViewRef{ID: id})
}

// Delete removes an authored view and returns the remaining collection.
// Example: `result, err := client.PartDesignViews().Delete(2)`.
//
// mcp:tool delete_part_design_view
// mcp:summary Delete an authored part design view and return the remaining views.
func (p PartDesignViews) Delete(id uint64) (wire.PartDesignViewsResult, error) {
	return call[wire.PartDesignViewsResult](p.c, wire.MethodPartDesignViewsDelete, wire.PartDesignViewRef{ID: id})
}

// SetOptions updates the supplied view options atomically.
// Example: `result, err := client.PartDesignViews().SetOptions(wire.SetPartDesignViewOptionsArgs{ID: 2})`.
//
// mcp:tool set_part_design_view_options
// mcp:summary Update the supplied display options of a part design view.
func (p PartDesignViews) SetOptions(args wire.SetPartDesignViewOptionsArgs) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsSetOptions, args)
}

// CaptureCamera stores an explicit camera for a part view.
// Example: `result, err := client.PartDesignViews().CaptureCamera(args)`.
//
// mcp:tool capture_part_design_view_camera
// mcp:summary Store the supplied camera in an unlocked part design view.
func (p PartDesignViews) CaptureCamera(args wire.CapturePartDesignViewCameraArgs) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsCaptureCamera, args)
}

// Camera reads the stored camera; an uncaptured camera returns an error.
// Example: `camera, err := client.PartDesignViews().Camera(2)`.
//
// mcp:tool get_part_design_view_camera
// mcp:summary Read the camera stored for a part design view.
func (p PartDesignViews) Camera(id uint64) (types.PartDesignViewCamera, error) {
	return call[types.PartDesignViewCamera](p.c, wire.MethodPartDesignViewsGetCamera, wire.PartDesignViewRef{ID: id})
}

// SetBodyVisibility applies body visibility to the active view.
// Example: `result, err := client.PartDesignViews().SetBodyVisibility(args)`.
//
// mcp:tool set_part_design_view_body_visibility
// mcp:summary Set body visibility in the active part design view.
func (p PartDesignViews) SetBodyVisibility(args wire.SetPartDesignViewBodyVisibilityArgs) (wire.PartDesignViewResult, error) {
	return call[wire.PartDesignViewResult](p.c, wire.MethodPartDesignViewsSetBodyVisibility, args)
}
