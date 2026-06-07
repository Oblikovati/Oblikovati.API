// SPDX-License-Identifier: Apache-2.0

package wire

// CameraView is the JSON shape of the viewport's camera as a look-at frame: the eye
// position, the target it looks at, the up vector, and the vertical field of view in
// radians. It is the response of [MethodViewGetCamera] and [MethodViewSetCamera].
//
// The look-at form is the host-native representation (lossless round-trip with the
// renderer's camera). A collaboration add-in that wants a position+orientation form
// (e.g. for slerp during presenter-follow, see the oblikovati-meeting ADR-0003)
// derives the rotation from eye→target and up on its own side.
//
// Eye, Target and Up are [x, y, z] in document/model units.
type CameraView struct {
	Eye    [3]float64 `json:"eye"`
	Target [3]float64 `json:"target"`
	Up     [3]float64 `json:"up"`
	FOV    float64    `json:"fov"`
}

// SetCameraArgs is the request of [MethodViewSetCamera]: the camera frame to apply, with
// optional addressing of which view it applies to. Camera state is per-view (a document
// owns a Views collection, each view owns a camera), so Document/View select the target;
// both default to the active document's active view, keeping the common call a bare frame.
// The frame fields mirror [CameraView] (a distinct type so request and response evolve
// independently, like [SetDisplayModeArgs] vs [DisplayModeView]).
type SetCameraArgs struct {
	Document uint64 `json:"document,omitempty"` // 0 ⇒ active document
	View     int    `json:"view,omitempty"`     // index into the document's views; 0 ⇒ active view

	Eye    [3]float64 `json:"eye"`
	Target [3]float64 `json:"target"`
	Up     [3]float64 `json:"up"`
	FOV    float64    `json:"fov"`
}

// GetCameraArgs is the request of [MethodViewGetCamera]: which view's camera to read.
// Both fields are optional and default to the active document's active view (so a no-arg
// call reads the active view, as before camera became per-view).
type GetCameraArgs struct {
	Document uint64 `json:"document,omitempty"` // 0 ⇒ active document
	View     int    `json:"view,omitempty"`     // 0 ⇒ active view
}
