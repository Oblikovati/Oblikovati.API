// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// PanelControlSpec is one declarative control of a dockable window's content. A
// PanelButton executes CommandID when clicked, so the add-in observes it through
// the ordinary command-ended event — no separate click plumbing. Text is the
// label/button caption; ID names the control for later content updates.
type PanelControlSpec struct {
	Kind      types.PanelControlKind `json:"kind,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Text      string                 `json:"text,omitempty"`
	CommandID string                 `json:"commandId,omitempty"`
}

// DockableWindowSpec is one add-in dockable window (M05-F03, #247): a titled panel
// the host docks into its layout (or floats), holding declarative content the host
// renders. Setting it again replaces title/content; Dock is only the initial
// placement (the user's re-docking wins afterwards).
type DockableWindowSpec struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	Dock     types.DockingState `json:"dock,omitempty"`
	Visible  bool               `json:"visible"`
	Controls []PanelControlSpec `json:"controls,omitempty"`
}

// SetDockableWindowArgs is the request of [MethodDockableWindowsSet]: create the
// window or replace its title/content if it exists.
type SetDockableWindowArgs struct {
	Window DockableWindowSpec `json:"window"`
}

// SetDockableWindowVisibleArgs is the request of [MethodDockableWindowsSetVisible].
type SetDockableWindowVisibleArgs struct {
	ID      string `json:"id"`
	Visible bool   `json:"visible"`
}

// DeleteDockableWindowArgs is the request of [MethodDockableWindowsDelete].
type DeleteDockableWindowArgs struct {
	ID string `json:"id"`
}

// ListDockableWindowsResult is the response of [MethodDockableWindowsList], in
// creation order.
type ListDockableWindowsResult struct {
	Windows []DockableWindowSpec `json:"windows"`
}

// DockableWindowChangedEvent is the push event (type [EventDockableWindowChanged])
// fired when a dockable window's visibility changes — whether by the add-in, a
// host menu toggle, or the user closing the window (the DockableWindowsEvents
// OnShow/OnHide equivalent).
type DockableWindowChangedEvent struct {
	Type    string `json:"type"` // always EventDockableWindowChanged
	ID      string `json:"id"`
	Visible bool   `json:"visible"`
}
