// SPDX-License-Identifier: Apache-2.0

package wire

// CommandInfo is the JSON shape of a registered command (the metadata Inventor put
// on a ControlDefinition) plus its current enabled state.
type CommandInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Tab         string `json:"tab,omitempty"`
	Category    string `json:"category,omitempty"`
	Alias       string `json:"alias,omitempty"`
	Tooltip     string `json:"tooltip,omitempty"`
	Enabled     bool   `json:"enabled"`
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
