// SPDX-License-Identifier: Apache-2.0

package types

// PanelControlKind is the kind of one declarative control inside an add-in panel
// surface (a dockable window today; M05-F07 mini-toolbars extend the vocabulary).
// The set is deliberately small: a panel is declared data the host renders, not a
// widget toolkit (M05-F03, #247).
type PanelControlKind uint8

const (
	// PanelLabel is a static text row (the zero value).
	PanelLabel PanelControlKind = 0
	// PanelButton is a clickable button that executes the command it names — the
	// add-in observes the click through the ordinary command-ended event.
	PanelButton PanelControlKind = 1
	// PanelSeparator is a horizontal rule between control groups.
	PanelSeparator PanelControlKind = 2
)

var panelControlKindNames = map[PanelControlKind]string{
	PanelLabel: "label", PanelButton: "button", PanelSeparator: "separator",
}

// String returns the kind's stable name.
func (k PanelControlKind) String() string {
	if name, ok := panelControlKindNames[k]; ok {
		return name
	}
	return "panelControlKind(?)"
}
