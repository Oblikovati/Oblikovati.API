// SPDX-License-Identifier: Apache-2.0

package wire

import "github.com/Oblikovati/api/types"

// CommandInfo is the JSON shape of a registered command (the metadata Inventor put
// on a ControlDefinition) plus its current enabled state.
type CommandInfo struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"displayName"`
	Tab         string            `json:"tab,omitempty"`
	Category    string            `json:"category,omitempty"`
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

// CreateCommandArgs is the request of [MethodCommandsCreate]: an add-in registering a
// ribbon button (Inventor's ButtonDefinition). The host creates a command with this
// metadata that appears in the ribbon; clicking it runs no host logic but fires a
// command-ended event the add-in receives via its Notify entry point and acts on. ID
// and DisplayName are required; the rest place and style the button.
type CreateCommandArgs struct {
	ID          string            `json:"id"`
	DisplayName string            `json:"displayName"`
	Tab         string            `json:"tab,omitempty"`
	Category    string            `json:"category,omitempty"`
	Alias       string            `json:"alias,omitempty"`
	Tooltip     string            `json:"tooltip,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	ButtonStyle types.ButtonStyle `json:"buttonStyle,omitempty"`
}
