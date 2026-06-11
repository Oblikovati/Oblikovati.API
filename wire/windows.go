// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The window-management surface of M05-F10 (#617): the view frame (the host's
// top-level window) and the tab strip of open documents. The host is a
// single-frame application today, so listFrames reports exactly one frame;
// tear-off frames and tab groups wait for a multi-window head (#160 ADR records
// the decline). View TILING inside the frame is the views.* group's layout
// surface (views.getLayout/setLayout), which this group deliberately does not
// duplicate.

// ViewFrameInfo is one top-level frame: its caption, window state and pixel size.
type ViewFrameInfo struct {
	Caption string            `json:"caption"`
	State   types.WindowState `json:"state,omitempty"`
	Width   int               `json:"width,omitempty"`
	Height  int               `json:"height,omitempty"`
}

// ListViewFramesResult is the response of [MethodWindowsListFrames].
type ListViewFramesResult struct {
	Frames []ViewFrameInfo `json:"frames"`
}

// ViewTabInfo is one document tab of the frame's tab strip.
type ViewTabInfo struct {
	Document uint64 `json:"document"`
	Title    string `json:"title"`
	Active   bool   `json:"active"`
	Dirty    bool   `json:"dirty,omitempty"`
}

// ListViewTabsResult is the response of [MethodWindowsListTabs], in strip order.
type ListViewTabsResult struct {
	Tabs []ViewTabInfo `json:"tabs"`
}

// ActivateViewTabArgs is the request of [MethodWindowsActivateTab].
type ActivateViewTabArgs struct {
	Document uint64 `json:"document"`
}

// CloseViewTabArgs is the request of [MethodWindowsCloseTab]; Force discards
// unsaved changes (otherwise the document is saved first).
type CloseViewTabArgs struct {
	Document uint64 `json:"document"`
	Force    bool   `json:"force,omitempty"`
}
