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

// SetCameraArgs is the request of [MethodViewSetCamera]: the camera frame to apply to
// the viewport. Same fields as [CameraView] (kept a distinct type so the request and
// response evolve independently, mirroring [SetDisplayModeArgs] vs [DisplayModeView]).
type SetCameraArgs struct {
	Eye    [3]float64 `json:"eye"`
	Target [3]float64 `json:"target"`
	Up     [3]float64 `json:"up"`
	FOV    float64    `json:"fov"`
}
