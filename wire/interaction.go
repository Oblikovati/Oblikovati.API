// SPDX-License-Identifier: Apache-2.0

package wire

// InteractionState is the JSON shape of the host's current interaction status — whether
// the local user is mid-action — the response of [MethodInteractionState].
//
// It exists so a collaboration add-in can gate incoming remote edits: applying a remote
// mutation while the local user is dragging a handle or running a tool risks a race in
// the host's recompute/undo state, so the add-in buffers remote operations while Busy is
// true (oblikovati-meeting ADR-0005).
//
// Busy is the single predicate add-ins should branch on (true when an interactive tool or
// command is active, or an explicit transaction is open). ActiveCommand/ActiveTool name
// what is running, for diagnostics/UI; they are empty when nothing is active.
type InteractionState struct {
	Busy          bool   `json:"busy"`
	ActiveCommand string `json:"activeCommand,omitempty"`
	ActiveTool    string `json:"activeTool,omitempty"`
	InTransaction bool   `json:"inTransaction"`
}
