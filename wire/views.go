// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati/api/types"

// A document owns a collection of views (the document's view collection); each view has its
// own camera, and the layout decides how many render at once. These DTOs are the wire
// shape of that collection and its mutations. Where a request carries a Document field,
// 0 means the active document.

// ViewInfo is the JSON shape of one view: its index in the document's collection, its
// name, whether it is the active view, and its camera frame.
type ViewInfo struct {
	Index  int        `json:"index"`
	Name   string     `json:"name"`
	Active bool       `json:"active"`
	Camera CameraView `json:"camera"`
}

// ListViewsResult is the response of [MethodViewsList]: every view of a document, which
// one is active, and the current tiling layout.
type ListViewsResult struct {
	Views       []ViewInfo       `json:"views"`
	ActiveIndex int              `json:"activeIndex"`
	Layout      types.ViewLayout `json:"layout"`
}

// ListViewsArgs is the request of [MethodViewsList]: which document to enumerate.
type ListViewsArgs struct {
	Document uint64 `json:"document,omitempty"` // 0 ⇒ active document
}

// AddViewArgs is the request of [MethodViewsAdd]: create a new view of a document
// (adds to the document's view collection). When CopyActiveCamera is set the new view starts at the active
// view's camera; otherwise it gets a default framed camera. The new view becomes active.
type AddViewArgs struct {
	Document         uint64 `json:"document,omitempty"`
	Name             string `json:"name,omitempty"`
	CopyActiveCamera bool   `json:"copyActiveCamera,omitempty"`
}

// ActivateViewArgs is the request of [MethodViewsActivate]: make the indexed view active.
type ActivateViewArgs struct {
	Document uint64 `json:"document,omitempty"`
	Index    int    `json:"index"`
}

// CloseViewArgs is the request of [MethodViewsClose]: remove the indexed view. Closing
// the last remaining view of a document is refused (a document always has ≥1 view).
type CloseViewArgs struct {
	Document uint64 `json:"document,omitempty"`
	Index    int    `json:"index"`
}

// RenameViewArgs is the request of [MethodViewsRename]: set the indexed view's name.
type RenameViewArgs struct {
	Document uint64 `json:"document,omitempty"`
	Index    int    `json:"index"`
	Name     string `json:"name"`
}

// LayoutResult is the response of [MethodViewsGetLayout] / [MethodViewsSetLayout]: the
// document's current tiling layout.
type LayoutResult struct {
	Document uint64           `json:"document,omitempty"`
	Layout   types.ViewLayout `json:"layout"`
}

// SetLayoutArgs is the request of [MethodViewsSetLayout]: how to tile the document's views.
type SetLayoutArgs struct {
	Document uint64           `json:"document,omitempty"`
	Layout   types.ViewLayout `json:"layout"`
}
