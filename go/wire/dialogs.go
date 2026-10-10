// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Host-provided modal dialogs (M05-F08, #615): the file open/save dialog and the
// web-view family. Like prompts (M05-F09), showing is asynchronous — a wire call
// must never block the session goroutine on user input — so the file dialog's
// choice arrives as a dialog.fileChosen push event. The reference's CancelError
// flag is subsumed by the event's Cancelled field.

// ShowFileDialogArgs is the request of [MethodDialogsShowFileDialog]: open the
// host's file dialog. ID keys the [FileDialogChosenEvent] that delivers the answer.
// Save switches to save semantics (type a new name) instead of open. Filter is a
// display-name|pattern list like "Meshes (*.stl *.obj)|*.stl;*.obj" with
// FilterIndex picking the initial entry. MultiSelect requests multiple paths
// (hosts may return a single one when their dialog cannot multi-select).
type ShowFileDialogArgs struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Save        bool   `json:"save,omitempty"`
	Filter      string `json:"filter,omitempty"`
	FilterIndex int    `json:"filterIndex,omitempty"`
	InitialDir  string `json:"initialDir,omitempty"`
	MultiSelect bool   `json:"multiSelect,omitempty"`
}

// FileDialogChosenEvent is the push event (type [EventFileDialogChosen]) that
// delivers the user's choice; Cancelled true means they dismissed the dialog.
type FileDialogChosenEvent struct {
	Type      string   `json:"type"` // always EventFileDialogChosen
	ID        string   `json:"id"`
	Paths     []string `json:"paths,omitempty"`
	Cancelled bool     `json:"cancelled,omitempty"`
}

// WebDialogSpec is one web view the host presents: a floating dialog (Modal pins
// it on top) or, with Dock set, a view docked into the chrome like an add-in
// window. The host's web rendering is behind a thin engine seam; a host without an
// embedded engine shows the URL with an open-in-browser affordance.
type WebDialogSpec struct {
	ID      string             `json:"id"`
	Title   string             `json:"title"`
	URL     string             `json:"url"`
	Modal   bool               `json:"modal,omitempty"`
	Dock    types.DockingState `json:"dock,omitempty"`
	Visible bool               `json:"visible"`
}

// ShowWebDialogArgs is the request of [MethodDialogsShowWebDialog]: create the web
// view or replace its title/URL if it exists.
type ShowWebDialogArgs struct {
	Dialog WebDialogSpec `json:"dialog"`
}

// CloseWebDialogArgs is the request of [MethodDialogsCloseWebDialog].
type CloseWebDialogArgs struct {
	ID string `json:"id"`
}

// ListWebViewsResult is the response of [MethodDialogsListWebViews], in creation
// order — the WebView(s) enumeration.
type ListWebViewsResult struct {
	Views []WebDialogSpec `json:"views"`
}

// WebDialogChangedEvent is the push event (type [EventWebDialogChanged]) fired when
// a web view's visibility changes (shown, or closed by the user/the owner).
type WebDialogChangedEvent struct {
	Type    string `json:"type"` // always EventWebDialogChanged
	ID      string `json:"id"`
	Visible bool   `json:"visible"`
}
