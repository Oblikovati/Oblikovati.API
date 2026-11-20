// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

const (
	// MethodPartDesignViewsList returns the active part’s views.
	MethodPartDesignViewsList = "partDesignViews.list"
	// MethodPartDesignViewsCopy copies one view without activating the copy.
	MethodPartDesignViewsCopy = "partDesignViews.copy"
	// MethodPartDesignViewsRename changes an authored view name.
	MethodPartDesignViewsRename = "partDesignViews.rename"
	// MethodPartDesignViewsActivate switches the active view.
	MethodPartDesignViewsActivate = "partDesignViews.activate"
	// MethodPartDesignViewsDelete removes an authored view.
	// Deleting an active authored view activates Primary. Hosts decline requests
	// for built-in views until their deletion behavior is established.
	MethodPartDesignViewsDelete = "partDesignViews.delete"
	// MethodPartDesignViewsSetOptions changes supplied view options.
	MethodPartDesignViewsSetOptions = "partDesignViews.setOptions"
	// MethodPartDesignViewsCaptureCamera stores a view camera.
	MethodPartDesignViewsCaptureCamera = "partDesignViews.captureCamera"
	// MethodPartDesignViewsGetCamera reads a stored camera.
	MethodPartDesignViewsGetCamera = "partDesignViews.getCamera"
	// MethodPartDesignViewsSetBodyVisibility changes active-view body visibility.
	MethodPartDesignViewsSetBodyVisibility = "partDesignViews.setBodyVisibility"
)

// PartDesignViewInfo describes one persisted part view. IDs are nonzero, unique
// within the part and stable across save and reopen. Primary is the built-in view
// and remains present. A new part lists Primary, Isometric, Front, Top and Right
// in that order. BuiltIn marks those initial entries. Exactly one view is active.
// Names are nonempty, contain no NUL, and are unique by exact spelling within the part; case differences remain
// distinct. Camera is absent until one is captured. Body visibility is addressed
// by sorted, distinct persistent body keys; an empty set is encoded as an empty
// array. Origin datum visibility is not a body override.
// StandardOrientation identifies a standard orientation when no camera has been
// captured; zero means no standard orientation. A captured camera takes precedence.
// Activating a view applies its stored camera and body visibility together; an
// unresolved reference is an error and does not change the active view.
// Example: `view := wire.PartDesignViewInfo{ID: 1, Name: "Primary", Primary: true, Active: true}`.
type PartDesignViewInfo struct {
	ID                  uint64                        `json:"id"`
	Name                string                        `json:"name"`
	BuiltIn             bool                          `json:"builtIn"`
	Primary             bool                          `json:"primary"`
	Active              bool                          `json:"active"`
	Locked              bool                          `json:"locked"`
	Published           bool                          `json:"published"`
	AutosaveCamera      bool                          `json:"autosaveCamera"`
	AnnotationScale     types.PartAnnotationScale     `json:"annotationScale"`
	StandardOrientation types.ViewOrientationTypeEnum `json:"standardOrientation,omitempty"`
	Camera              *types.PartDesignViewCamera   `json:"camera,omitempty"`
	HiddenBodyKeys      []string                      `json:"hiddenBodyKeys"`
}

// PartDesignViewsResult lists the views in creation order.
// Example: `result := wire.PartDesignViewsResult{Views: []wire.PartDesignViewInfo{}}`.
type PartDesignViewsResult struct {
	Views []PartDesignViewInfo `json:"views"`
}

// PartDesignViewResult returns one view after an operation.
// Example: `result := wire.PartDesignViewResult{View: wire.PartDesignViewInfo{ID: 1}}`.
type PartDesignViewResult struct {
	View PartDesignViewInfo `json:"view"`
}

// PartDesignViewRef selects one view in the active part document.
// Example: `ref := wire.PartDesignViewRef{ID: 1}`.
type PartDesignViewRef struct {
	ID uint64 `json:"id"`
}

// CopyPartDesignViewArgs copies all authored settings of SourceID into a new,
// inactive view. Empty Name requests a unique generated name; an explicit name
// must be unique. The operation does not change the active view or viewport.
// Example: `args := wire.CopyPartDesignViewArgs{SourceID: 1, Name: "Review"}`.
type CopyPartDesignViewArgs struct {
	SourceID uint64 `json:"sourceId"`
	Name     string `json:"name,omitempty"`
}

// RenamePartDesignViewArgs changes an authored view's name while retaining ID.
// Hosts decline renames of built-in views until their behavior is established.
// A duplicate or invalid name leaves state unchanged.
// Example: `args := wire.RenamePartDesignViewArgs{ID: 2, Name: "Review"}`.
type RenamePartDesignViewArgs struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// SetPartDesignViewOptionsArgs changes only the supplied options atomically.
// A locked view retains its stored camera and body visibility when the active
// display is changed; reactivation restores the stored state. Unlocking does not
// silently replace that stored state. Annotation scale
// applies to the active view; changing an inactive view's scale is refused. Invalid options leave all state
// unchanged. Omitted fields preserve their existing values.
// Example: `args := wire.SetPartDesignViewOptionsArgs{ID: 2, Locked: new(bool)}`.
type SetPartDesignViewOptionsArgs struct {
	ID              uint64                     `json:"id"`
	Locked          *bool                      `json:"locked,omitempty"`
	Published       *bool                      `json:"published,omitempty"`
	AutosaveCamera  *bool                      `json:"autosaveCamera,omitempty"`
	AnnotationScale *types.PartAnnotationScale `json:"annotationScale,omitempty"`
}

// CapturePartDesignViewCameraArgs stores a camera for an unlocked view.
// The camera must be finite and usable. A locked or unknown view is refused.
// Example: `args := wire.CapturePartDesignViewCameraArgs{ID: 2, Camera: camera}`.
type CapturePartDesignViewCameraArgs struct {
	ID     uint64                     `json:"id"`
	Camera types.PartDesignViewCamera `json:"camera"`
}

// SetPartDesignViewBodyVisibilityArgs changes current body visibility in the
// active view. BodyKeys are distinct, nonempty hexadecimal persistent body
// reference keys. An empty list means all current bodies; no origin datum is
// included. Unknown or ambiguous keys leave state unchanged. If the view is
// locked, the current display changes transiently but the stored view is not
// changed; otherwise the stored visibility changes with the display.
// Example: `args := wire.SetPartDesignViewBodyVisibilityArgs{ID: 2, Visible: false}`.
type SetPartDesignViewBodyVisibilityArgs struct {
	ID       uint64   `json:"id"`
	BodyKeys []string `json:"bodyKeys"`
	Visible  bool     `json:"visible"`
}
