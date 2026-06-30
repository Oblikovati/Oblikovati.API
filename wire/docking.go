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
//
// Container kinds (grid/group/tabs, ADR-0019) nest: they own Children and are laid out
// rather than drawn. A grid declares its column tracks in Columns with ColumnGap/RowGap
// spacing, and each child may carry a Cell placement (nil = auto-flow). A group is a
// titled vertical stack (Title is its caption). Tabs treats each child as one tab whose
// Title is the caption. All these fields are omitempty, so a leaf control marshals exactly
// as before — older hosts ignore the unknown fields and degrade an unknown kind to its Text.
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

	// Container fields (grid/group/tabs only).
	Title     string             `json:"title,omitempty"`     // group/tab caption
	Children  []PanelControlSpec `json:"children,omitempty"`  // nested controls of a container
	Columns   []types.GridTrack  `json:"columns,omitempty"`   // grid: column tracks (rows are auto-height)
	ColumnGap float64            `json:"columnGap,omitempty"` // grid: px gap between columns
	RowGap    float64            `json:"rowGap,omitempty"`    // grid: px gap between rows
	Cell      *types.GridCell    `json:"cell,omitempty"`      // this control's placement in its parent grid

	Rows    []PanelReferenceRow `json:"rows,omitempty"`    // referenceList: current picked refs
	Accepts []string            `json:"accepts,omitempty"` // referenceList: allowed kinds ("face"/"edge"/"vertex"); empty = any
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

// SetDockableWindowValueArgs is the request of [MethodDockableWindowsSetValue]: it drives one
// editable control of an add-in dockable window to a value, exactly as a user edit would — the host
// updates the stored control and notifies the owning add-in (which may react, e.g. re-render the
// window). Value is the control's string form: the option text for a dropdown/combo, "true"/"false"
// for a checkbox, the number for a value editor/slider, the text for a text box.
type SetDockableWindowValueArgs struct {
	WindowId  string `json:"windowId"`
	ControlId string `json:"controlId"`
	Value     string `json:"value"`
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

// PanelReferenceRow is one row of a referenceList control: a host geometry selection
// reference plus an optional display label (the host derives one, e.g. "Face3", when empty).
type PanelReferenceRow struct {
	Ref   string `json:"ref"`
	Label string `json:"label,omitempty"`
}

// SetDockableWindowReferencesArgs is the request of [MethodDockableWindowsSetReferences]: it
// replaces a referenceList control's rows exactly as an Add-from-selection would, and notifies
// the owning add-in with a [PanelReferencesChangedEvent]. Refs is the full new set.
type SetDockableWindowReferencesArgs struct {
	WindowId  string   `json:"windowId"`
	ControlId string   `json:"controlId"`
	Refs      []string `json:"refs"`
}

// PanelReferencesChangedEvent is the push event (type [EventPanelReferencesChanged]) fired when a
// referenceList control's rows change — by the user's Add-from-selection / per-row Remove or by
// [MethodDockableWindowsSetReferences]. Refs is the FULL new set (bulk-state, matching the rest of
// the panel model); Action is "add"/"remove" for diagnostics only.
type PanelReferencesChangedEvent struct {
	Type      string   `json:"type"` // always EventPanelReferencesChanged
	WindowId  string   `json:"windowId"`
	ControlId string   `json:"controlId"`
	Refs      []string `json:"refs"`
	Action    string   `json:"action,omitempty"`
}
