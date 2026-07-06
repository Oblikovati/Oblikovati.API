// SPDX-License-Identifier: Apache-2.0

package client

import (
	"strconv"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Panel-control constructors for building a dockable-window form declaratively. They cover the
// editable control vocabulary: label, button,
// separator, text box, unit-bearing value editor, check box, dropdown, combo box, slider. Each
// editable control carries an ID the host echoes back in a [wire.PanelValueChangedEvent] when
// the user edits it.

// PanelLabel is a static text row.
func PanelLabel(id, text string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelLabel, ID: id, Text: text}
}

// PanelButton runs commandID when clicked.
func PanelButton(id, text, commandID string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelButton, ID: id, Text: text, CommandID: commandID}
}

// PanelSeparator is a horizontal rule.
func PanelSeparator() wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelSeparator}
}

// PanelTextBox is a single-line text field labelled text with the current value.
func PanelTextBox(id, text, value string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelTextBox, ID: id, Text: text, Value: value}
}

// PanelValueEditor is a unit-bearing numeric field (value is a unit expression, e.g. "46.7 mm")
// — the parametric-driver input.
func PanelValueEditor(id, text, value string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelValueEditor, ID: id, Text: text, Value: value}
}

// PanelCheckBox is a boolean toggle.
func PanelCheckBox(id, text string, checked bool) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelCheckBox, ID: id, Text: text, Value: strconv.FormatBool(checked)}
}

// PanelDropdown selects one of options (selected is the current choice).
func PanelDropdown(id, text string, options []string, selected string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelDropdown, ID: id, Text: text, Options: options, Value: selected}
}

// PanelComboBox is an editable dropdown (pick from options or type a value).
func PanelComboBox(id, text string, options []string, value string) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelComboBox, ID: id, Text: text, Options: options, Value: value}
}

// PanelSlider is a bounded numeric slider over [min,max] stepping by step.
func PanelSlider(id, text string, value, min, max, step float64) wire.PanelControlSpec {
	return wire.PanelControlSpec{
		Kind: types.PanelSlider, ID: id, Text: text,
		Value: strconv.FormatFloat(value, 'g', -1, 64), Min: min, Max: max, Step: step,
	}
}

// PanelReferenceList builds a geometry reference-list control: rows are the current picked
// refs; accepts limits which host selection kinds Add may append ("face"/"edge"/"vertex";
// empty = any). Row edits arrive as a wire.PanelReferencesChangedEvent, not the scalar value event.
func PanelReferenceList(id, text string, accepts []string, rows []wire.PanelReferenceRow) wire.PanelControlSpec {
	return wire.PanelControlSpec{Kind: types.PanelReferenceList, ID: id, Text: text, Accepts: accepts, Rows: rows}
}
