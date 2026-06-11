// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Dialogs is the host-dialog operation group (M05-F08): the file open/save dialog
// and the web-view family, so add-ins never ship their own toolkit. File choices
// arrive as dialog.fileChosen push events (asynchronous, like prompts).
type Dialogs struct{ c *Client }

// Dialogs returns the host-dialog operation group.
func (c *Client) Dialogs() Dialogs { return Dialogs{c} }

// ShowFileDialog opens the host's file dialog; the user's choice arrives as a
// [wire.FileDialogChosenEvent] keyed by args.ID.
//
//	client.Dialogs().ShowFileDialog(wire.ShowFileDialogArgs{
//	    ID: "sim.report", Save: true, Title: "Save report", Filter: "HTML (*.html)|*.html",
//	})
func (d Dialogs) ShowFileDialog(args wire.ShowFileDialogArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDialogsShowFileDialog, args, &r)
}

// ShowWebDialog presents a web view: floating (modal or not), or docked when the
// spec carries a docking state.
func (d Dialogs) ShowWebDialog(spec wire.WebDialogSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDialogsShowWebDialog, wire.ShowWebDialogArgs{Dialog: spec}, &r)
}

// CloseWebDialog dismisses a web view.
func (d Dialogs) CloseWebDialog(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDialogsCloseWebDialog, wire.CloseWebDialogArgs{ID: id}, &r)
}

// ListWebViews returns the presented web views in creation order.
func (d Dialogs) ListWebViews() (wire.ListWebViewsResult, error) {
	var r wire.ListWebViewsResult
	return r, d.c.call(wire.MethodDialogsListWebViews, nil, &r)
}
