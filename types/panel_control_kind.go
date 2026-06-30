// SPDX-License-Identifier: Apache-2.0

package types

// PanelControlKind is the kind of one declarative control inside an add-in panel
// surface (a dockable window; also mini-toolbars, M05-F07). A panel is declared data the
// host renders, not a widget toolkit (M05-F03, #247). The editable kinds mirror Autodesk
// Inventor's MiniToolbarControlTypeEnum (kCheckBox/kComboBox/kDropdown/kValueEditor/kSlider/
// kTextBox) so add-ins get a familiar form vocabulary. Each editable control carries an ID;
// when the user changes it the host pushes a [PanelValueChangedEvent] to the add-in.
type PanelControlKind uint8

const (
	// PanelLabel is a static text row (the zero value).
	PanelLabel PanelControlKind = 0
	// PanelButton is a clickable button that executes the command it names — the
	// add-in observes the click through the ordinary command-ended event.
	PanelButton PanelControlKind = 1
	// PanelSeparator is a horizontal rule between control groups.
	PanelSeparator PanelControlKind = 2
	// PanelTextBox is a single-line free-text input (Value = current text).
	PanelTextBox PanelControlKind = 3
	// PanelValueEditor is a unit-bearing numeric value editor for a dimension/parameter
	// (Value = a unit expression, e.g. "46.7 mm"); the natural input for parametric drivers.
	PanelValueEditor PanelControlKind = 4
	// PanelCheckBox is a boolean toggle (Value = "true"/"false").
	PanelCheckBox PanelControlKind = 5
	// PanelDropdown selects exactly one of Options (Value = the selected item).
	PanelDropdown PanelControlKind = 6
	// PanelComboBox is an editable dropdown: pick from Options or type a value (Value = text).
	PanelComboBox PanelControlKind = 7
	// PanelSlider is a bounded numeric slider (Value = number, bounded by Min/Max, step Step).
	PanelSlider PanelControlKind = 8
	// PanelGrid is a container that lays its Children out in a CSS-grid-like grid: Columns
	// declares the column tracks, each child may carry a Cell placement, otherwise children
	// auto-flow left-to-right wrapping at the column count (ADR-0019). Rows are content-height.
	PanelGrid PanelControlKind = 9
	// PanelGroup is a titled box (Title is the caption) that stacks its Children vertically —
	// the QGroupBox of this vocabulary.
	PanelGroup PanelControlKind = 10
	// PanelTabs is a tab strip: each direct child is one tab whose Title is the tab caption and
	// whose own content (typically a grid or group) is the pane.
	PanelTabs PanelControlKind = 11
)

var panelControlKindNames = map[PanelControlKind]string{
	PanelLabel: "label", PanelButton: "button", PanelSeparator: "separator",
	PanelTextBox: "textBox", PanelValueEditor: "valueEditor", PanelCheckBox: "checkBox",
	PanelDropdown: "dropdown", PanelComboBox: "comboBox", PanelSlider: "slider",
	PanelGrid: "grid", PanelGroup: "group", PanelTabs: "tabs",
}

// String returns the kind's stable name.
func (k PanelControlKind) String() string {
	if name, ok := panelControlKindNames[k]; ok {
		return name
	}
	return "panelControlKind(?)"
}
