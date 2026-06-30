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
	// TooltipTitle / TooltipExpanded are the progressive tooltip (M05-F09): the
	// title heads the hover tip; the expanded text appears after a longer hover.
	TooltipTitle    string `json:"tooltipTitle,omitempty"`
	TooltipExpanded string `json:"tooltipExpanded,omitempty"`
	Icon            string `json:"icon,omitempty"`
	// IconSVG is inline SVG markup the add-in supplied for this button (see
	// [CreateCommandArgs.IconSVG]); empty when the button uses a host-bundled Icon key.
	IconSVG     string            `json:"iconSvg,omitempty"`
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

// SubmitCommandLineArgs is the request of [MethodCommandLineSubmit]: one line of Command
// Window input — a command word ("EXTRUDE"), an alias ("E"), a coordinate/value/keyword for
// the active command's current step ("10,5", "25", "Close"), or "" to finish/repeat. The
// host drives the same command-line REPL the UI uses, so an add-in or MCP tool can model
// headlessly: submit "LINE", then "0,0", then "10,0".
type SubmitCommandLineArgs struct {
	Line string `json:"line"`
}

// CommandLineResult is the response of [MethodCommandLineSubmit]. Output is the scrollback
// lines this submission produced (echoes, prompts, results). Prompt is the active command's
// next step prompt, and Awaiting is true while a command is mid-interaction (more input
// expected). Error carries a command-line error (e.g. an unknown command) as a message
// rather than a transport failure, so the caller can show it inline like the UI does.
type CommandLineResult struct {
	Output   []string `json:"output,omitempty"`
	Prompt   string   `json:"prompt,omitempty"`
	Awaiting bool     `json:"awaiting"`
	Error    string   `json:"error,omitempty"`
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
// Kind picks the control's behavior (default a one-shot button). A PopupControl
// names other registered commands in Items: the button opens a menu of them and
// each item runs its own command — the CommandBarPopUp equivalent (M05-F03, #247).
type CreateCommandArgs struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"displayName"`
	Ribbon      types.RibbonKey   `json:"ribbon,omitempty"`
	Tab         string            `json:"tab,omitempty"`
	Category    string            `json:"category,omitempty"`
	Environment types.Environment `json:"environment,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Tooltip     string            `json:"tooltip,omitempty"`
	// The progressive tooltip (M05-F09): title heads the hover tip; the expanded
	// text appears after a longer hover.
	TooltipTitle    string `json:"tooltipTitle,omitempty"`
	TooltipExpanded string `json:"tooltipExpanded,omitempty"`
	Icon            string `json:"icon,omitempty"`
	// IconSVG lets an add-in ship its own button glyph as inline SVG markup instead of
	// referencing a host-bundled Icon key — so an add-in is not limited to the icons the
	// host happens to embed. When set it takes precedence over Icon. The markup should
	// follow the host's glyph conventions: a square (24×24) viewBox and the theme's sentinel
	// paints (the outline, a fill role, and an accent role), which the host recolours per
	// theme. Oversized markup is rejected by the host. (Oblikovati#671)
	IconSVG     string            `json:"iconSvg,omitempty"`
	ButtonStyle types.ButtonStyle `json:"buttonStyle,omitempty"`
	Kind        types.ControlKind `json:"kind,omitempty"`
	Items       []string          `json:"items,omitempty"`
}
