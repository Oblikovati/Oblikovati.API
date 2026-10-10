// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// CameraView is the JSON shape of the viewport's camera as a look-at frame: the eye
// position, the target it looks at, the up vector, and the vertical field of view in
// radians. It is the response of [MethodViewGetCamera] and [MethodViewSetCamera].
//
// The look-at form is the host-native representation (lossless round-trip with the
// renderer's camera). A collaboration add-in that wants a position+orientation form
// (e.g. for slerp during presenter-follow, see the oblikovati-meeting ADR-0003)
// derives the rotation from eye→target and up on its own side.
//
// Eye and Target are positions, Up a direction — the M01-F05 geometry value
// types, whose [x, y, z] JSON form is identical to the [3]float64 encoding this
// type used before they existed.
type CameraView struct {
	Eye    types.Point  `json:"eye"`
	Target types.Point  `json:"target"`
	Up     types.Vector `json:"up"`
	FOV    float64      `json:"fov"`
	// Projection is how the view projects: orthographic (parallel), perspective, or perspective
	// with orthographic view-cube faces. It is a property of the VIEW, not of the look-at frame,
	// and was previously reachable only as the global new-window default — so a client could not
	// tell whether what it captured was foreshortened.
	Projection types.ProjectionTypeEnum `json:"projection,omitempty"`
}

// SetCameraArgs is the request of [MethodViewSetCamera]: the camera frame to apply, plus
// optional addressing of which document it applies to. Camera state is per-view (a
// document owns a Views collection, each view owns a camera); the frame is applied to the
// addressed document's active view (Document 0 ⇒ the active document). To target a
// non-active view, activate it first ([MethodViewsActivate]); every view's camera is also
// readable via [MethodViewsList]. The frame fields mirror [CameraView] (a distinct type so
// request and response evolve independently, like [SetDisplayModeArgs] vs [DisplayModeView]).
type SetCameraArgs struct {
	Document uint64 `json:"document,omitempty"` // 0 ⇒ active document; applies to that document's active view

	Eye    types.Point  `json:"eye"`
	Target types.Point  `json:"target"`
	Up     types.Vector `json:"up"`
	FOV    float64      `json:"fov"`
	// Projection changes how the view projects. Omit (0) to leave it as it is — a caller moving
	// the camera should not have to restate the projection to keep it.
	Projection types.ProjectionTypeEnum `json:"projection,omitempty"`
}

// GetCameraArgs is the request of [MethodViewGetCamera]: which document's active-view
// camera to read. Document is optional and defaults to the active document (so a no-arg
// call reads the active document's active view, as before camera became per-view).
type GetCameraArgs struct {
	Document uint64 `json:"document,omitempty"` // 0 ⇒ active document; reads that document's active view
}
