// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// A document keeps a set of named views — saved camera frames a user (or add-in) can restore
// exactly. These DTOs are the wire shape of capturing, listing, restoring, and deleting them,
// plus jumping to a standard orientation. Document 0 means the active document.

// NamedViewInfo is the JSON shape of one saved named view: its name and the camera frame it
// restores. It is the element of [NamedViewsResult] and the response of
// [MethodViewsCaptureNamed].
type NamedViewInfo struct {
	Name   string     `json:"name"`
	Camera CameraView `json:"camera"`
}

// CaptureNamedViewArgs is the request of [MethodViewsCaptureNamed]: save the active view's
// current camera under Name (replacing any existing named view of that name).
type CaptureNamedViewArgs struct {
	Document uint64 `json:"document,omitempty"`
	Name     string `json:"name"`
}

// NamedViewsResult is the response of [MethodViewsListNamed]: every saved named view of a
// document.
type NamedViewsResult struct {
	Views []NamedViewInfo `json:"views"`
}

// NamedViewRefArgs is the request of [MethodViewsRestoreNamed] / [MethodViewsDeleteNamed]:
// the named view to restore (its camera is applied to the active view, animated) or delete.
type NamedViewRefArgs struct {
	Document uint64 `json:"document,omitempty"`
	Name     string `json:"name"`
}

// SetOrientationArgs is the request of [MethodViewSetOrientation]: jump the active view to a
// standard orientation (front/top/iso…), optionally fitting the model to the view.
type SetOrientationArgs struct {
	Document    uint64                        `json:"document,omitempty"`
	Orientation types.ViewOrientationTypeEnum `json:"orientation"`
	Fit         bool                          `json:"fit,omitempty"`
}

// CameraChangedEvent is the payload of the [EventCameraChanged] push event: the document whose
// active-view camera moved and the new frame.
type CameraChangedEvent struct {
	Document uint64     `json:"document"`
	Camera   CameraView `json:"camera"`
}
