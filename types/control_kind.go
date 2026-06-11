// SPDX-License-Identifier: Apache-2.0

package types

// ControlKind is how a command behaves as a ribbon control — the
// CommandControl-type discriminator. The kind decides interaction (one-shot,
// on/off, pick-one, numeric, menu), while [ButtonStyle] decides only its look.
//
// This is the canonical, Apache-2.0 definition; the GPL implementation aliases it
// (app.ControlKind) so existing call sites are unaffected.
type ControlKind uint8

const (
	// ButtonControl runs its command on click — a one-shot action (the zero value).
	ButtonControl ControlKind = 0
	// ToggleControl is a stateful on/off control; its active state renders pressed.
	ToggleControl ControlKind = 1
	// ComboControl is one mutually-exclusive choice of its panel's commands; a
	// panel of combo controls renders as a single drop-down selection box.
	ComboControl ControlKind = 2
	// SpinnerControl is a numeric stepper.
	SpinnerControl ControlKind = 3
	// PopupControl opens a menu of other registered commands (the CommandBarPopUp
	// equivalent): the control itself runs nothing; each menu item runs its own
	// command (M05-F03, #247).
	PopupControl ControlKind = 4
)

var controlKindNames = map[ControlKind]string{
	ButtonControl: "button", ToggleControl: "toggle", ComboControl: "combo",
	SpinnerControl: "spinner", PopupControl: "popup",
}

// String returns the kind's stable name.
func (k ControlKind) String() string {
	if name, ok := controlKindNames[k]; ok {
		return name
	}
	return "controlKind(?)"
}
