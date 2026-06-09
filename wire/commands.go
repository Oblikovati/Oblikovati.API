// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// CommandInfo is the JSON shape of a registered command (its control-definition
// metadata) plus its current enabled state.
type CommandInfo struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"displayName"`
	Ribbon      types.RibbonKey   `json:"ribbon,omitempty"`
	Tab         string            `json:"tab,omitempty"`
	Category    string            `json:"category,omitempty"`
	Environment types.Environment `json:"environment,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Tooltip     string            `json:"tooltip,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	ButtonStyle types.ButtonStyle `json:"buttonStyle,omitempty"`
	Enabled     bool              `json:"enabled"`
}

// ListCommandsResult is the response of [MethodCommandsList].
type ListCommandsResult struct {
	Commands []CommandInfo `json:"commands"`
}

// ExecuteCommandArgs is the request of [MethodCommandsExecute]: the id of the
// command to run.
type ExecuteCommandArgs struct {
	ID string `json:"id"`
}

// SetCommandStateArgs is the request of [MethodCommandsSetState]: an add-in updating one of
// its own commands' live ribbon state. Active toggles the button's pressed/highlighted look
// (rendered in the accent color), so a stateful control like a presenter or follow toggle
// reads on/off at a glance. DisplayName, when non-empty, relabels the button (e.g.
// "Presenter" → "Presenting") — leave it empty to keep the current label. Enabled, when
// non-nil, greys the button out (false) or restores it (true) — e.g. a collaboration add-in
// disabling its presenter/follow controls until the user joins a session; nil leaves the
// enabled state unchanged.
type SetCommandStateArgs struct {
	ID          string `json:"id"`
	Active      bool   `json:"active"`
	Enabled     *bool  `json:"enabled,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// CreateCommandArgs is the request of [MethodCommandsCreate]: an add-in registering a
// ribbon button. The host creates a command with this metadata that appears in the
// ribbon; clicking it runs no host logic but fires a command-ended event the add-in
// receives via its Notify entry point and acts on. ID and DisplayName are required; the
// rest place and style the button.
//
// Ribbon picks which document ribbon the button lands on (empty ⇒ the Part ribbon);
// Environment scopes it to a context (empty/base ⇒ always shown; sketch ⇒ the contextual
// Sketch tab). Together with Tab/Category this places a button on a named panel of a named
// tab of a chosen ribbon (e.g. the Draw panel of the Sketch tab of the Part ribbon).
type CreateCommandArgs struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"displayName"`
	Ribbon      types.RibbonKey   `json:"ribbon,omitempty"`
	Tab         string            `json:"tab,omitempty"`
	Category    string            `json:"category,omitempty"`
	Environment types.Environment `json:"environment,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Tooltip     string            `json:"tooltip,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	ButtonStyle types.ButtonStyle `json:"buttonStyle,omitempty"`
}
