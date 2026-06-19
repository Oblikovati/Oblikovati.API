// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// PanelControlSpec is one declarative control of a dockable window's content. A
// PanelButton executes CommandID when clicked, so the add-in observes it through
// the ordinary command-ended event — no separate click plumbing. Text is the
// label/button caption (or the field label for editable controls); ID names the
// control. Editable controls (textBox/valueEditor/checkBox/dropdown/comboBox/slider)
// carry their current Value; dropdown/comboBox list their choices in Options; slider
// and valueEditor bound the input with Min/Max/Step. When the user edits an editable
// control the host pushes a [PanelValueChangedEvent] carrying the control's ID + new Value.
type PanelControlSpec struct {
	Kind      types.PanelControlKind `json:"kind,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Text      string                 `json:"text,omitempty"`
	CommandID string                 `json:"commandId,omitempty"`
	Value     string                 `json:"value,omitempty"`   // current value of an editable control
	Options   []string               `json:"options,omitempty"` // choices for dropdown/comboBox
	Min       float64                `json:"min,omitempty"`     // slider/valueEditor lower bound
	Max       float64                `json:"max,omitempty"`     // slider/valueEditor upper bound
	Step      float64                `json:"step,omitempty"`    // slider/valueEditor increment
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
