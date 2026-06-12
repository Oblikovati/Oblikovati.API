// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The in-canvas mini-toolbar surface of M05-F07 (#614): a floating toolbar an
// interactive command (or add-in) declares; the host renders it in the viewport and
// streams the user's edits back as miniToolbar.changed events. One miniToolbar.set
// declares the whole toolbar (the clientGraphics bulk-group precedent).

// MiniToolbarControlSpec is one declarative control. ID names the control in change
// events and updates. The value fields are per kind: Checked (checkbox), Number +
// Min/Max (slider), Value (textbox / value-editor expression / text-editor),
// Options + Selected (combo).
type MiniToolbarControlSpec struct {
	Kind     types.MiniToolbarControlKind `json:"kind,omitempty"`
	ID       string                       `json:"id"`
	Label    string                       `json:"label,omitempty"`
	Tooltip  string                       `json:"tooltip,omitempty"`
	Value    string                       `json:"value,omitempty"`
	Checked  bool                         `json:"checked,omitempty"`
	Number   float64                      `json:"number,omitempty"`
	Min      float64                      `json:"min,omitempty"`
	Max      float64                      `json:"max,omitempty"`
	Options  []string                     `json:"options,omitempty"`
	Selected int                          `json:"selected,omitempty"`
}

// MiniToolbarSpec is one mini-toolbar. Anchor places it: a 3D model point the host
// projects each frame (Anchor non-nil), else viewport-relative pixels (ScreenX/Y).
// Command, when set, ties the toolbar's lifetime to that command: the host removes
// it when the active tool commits or cancels — the interaction-graphics lifecycle.
// OK/Apply/Cancel render when their Show flags are set; the user's choice arrives
// as a miniToolbar.committed event (ok and cancel also dismiss the toolbar).
type MiniToolbarSpec struct {
	ID          string                   `json:"id"`
	Command     string                   `json:"command,omitempty"`
	Anchor      *types.Point             `json:"anchor,omitempty"`
	ScreenX     float64                  `json:"screenX,omitempty"`
	ScreenY     float64                  `json:"screenY,omitempty"`
	Visible     bool                     `json:"visible"`
	HeadsUpText string                   `json:"headsUpText,omitempty"`
	ShowOK      bool                     `json:"showOK,omitempty"`
	ShowApply   bool                     `json:"showApply,omitempty"`
	ShowCancel  bool                     `json:"showCancel,omitempty"`
	Controls    []MiniToolbarControlSpec `json:"controls,omitempty"`
}

// SetMiniToolbarArgs is the request of [MethodMiniToolbarSet]: create the toolbar
// or replace it entirely.
type SetMiniToolbarArgs struct {
	Toolbar MiniToolbarSpec `json:"toolbar"`
}

// UpdateMiniToolbarArgs is the request of [MethodMiniToolbarUpdate]: merge the
// given controls' values into the toolbar by control id (cheaper than a full
// re-set while a command streams state).
type UpdateMiniToolbarArgs struct {
	ID       string                   `json:"id"`
	Controls []MiniToolbarControlSpec `json:"controls"`
}

// RemoveMiniToolbarArgs is the request of [MethodMiniToolbarRemove].
type RemoveMiniToolbarArgs struct {
	ID string `json:"id"`
}

// ListMiniToolbarsResult is the response of [MethodMiniToolbarList], in creation
// order.
type ListMiniToolbarsResult struct {
	Toolbars []MiniToolbarSpec `json:"toolbars"`
}

// MiniToolbarChangedEvent is the push event (type [EventMiniToolbarChanged]) fired
// when the user edits one control: the control's current value fields, per kind
// (a button click carries just the ids).
type MiniToolbarChangedEvent struct {
	Type     string  `json:"type"` // always EventMiniToolbarChanged
	Toolbar  string  `json:"toolbar"`
	Control  string  `json:"control"`
	Value    string  `json:"value,omitempty"`
	Checked  bool    `json:"checked,omitempty"`
	Number   float64 `json:"number,omitempty"`
	Selected int     `json:"selected,omitempty"`
}

// MiniToolbarCommittedEvent is the push event (type [EventMiniToolbarCommitted])
// fired when the user presses OK, Apply or Cancel; Gesture is "ok", "apply" or
// "cancel" (ok and cancel also dismiss the toolbar).
type MiniToolbarCommittedEvent struct {
	Type    string `json:"type"` // always EventMiniToolbarCommitted
	Toolbar string `json:"toolbar"`
	Gesture string `json:"gesture"`
}
