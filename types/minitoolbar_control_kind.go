// SPDX-License-Identifier: Apache-2.0

package types

// MiniToolbarControlKind is the kind of one declarative control on an in-canvas
// mini-toolbar — the MiniToolbarControlTypeEnum equivalent (M05-F07, #614). Like
// the dockable-window panel controls, a mini-toolbar is declared data the host
// renders, not a widget toolkit.
type MiniToolbarControlKind uint8

const (
	// MiniToolbarButton is a clickable button (the zero value); clicks arrive as
	// change events.
	MiniToolbarButton MiniToolbarControlKind = 0
	// MiniToolbarCheckbox is an on/off toggle.
	MiniToolbarCheckbox MiniToolbarControlKind = 1
	// MiniToolbarCombo is a pick-one dropdown of its Options.
	MiniToolbarCombo MiniToolbarControlKind = 2
	// MiniToolbarSlider is a numeric slider over [Min, Max].
	MiniToolbarSlider MiniToolbarControlKind = 3
	// MiniToolbarTextBox is a single-line text input.
	MiniToolbarTextBox MiniToolbarControlKind = 4
	// MiniToolbarValueEditor is a unit-aware numeric input: its text Value is an
	// expression the host's parameter engine evaluates ("12 mm", "width/2").
	MiniToolbarValueEditor MiniToolbarControlKind = 5
	// MiniToolbarTextEditor is a multi-line text input.
	MiniToolbarTextEditor MiniToolbarControlKind = 6
)

var miniToolbarControlKindNames = map[MiniToolbarControlKind]string{
	MiniToolbarButton: "button", MiniToolbarCheckbox: "checkbox", MiniToolbarCombo: "combo",
	MiniToolbarSlider: "slider", MiniToolbarTextBox: "textbox",
	MiniToolbarValueEditor: "value-editor", MiniToolbarTextEditor: "text-editor",
}

// String returns the kind's stable name.
func (k MiniToolbarControlKind) String() string {
	if name, ok := miniToolbarControlKindNames[k]; ok {
		return name
	}
	return "miniToolbarControlKind(?)"
}
