// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Interaction is the operation group for the host's current interaction status: whether
// the local user is mid-action. A collaboration add-in queries it to gate incoming remote
// edits — buffering them while the local user is busy (oblikovati-meeting ADR-0005).
type Interaction struct{ c *Client }

// Interaction returns the interaction-status operation group.
func (c *Client) Interaction() Interaction { return Interaction{c} }

// State reports whether an interactive tool/command is active or a transaction is open.
//
//	if st, _ := client.Interaction().State(); st.Busy { /* buffer remote edits */ }
//
// mcp:tool interaction_state
// mcp:summary Reports whether an interactive tool/command is active or a transaction is open.
func (i Interaction) State() (wire.InteractionState, error) {
	return call[wire.InteractionState](i.c, wire.MethodInteractionState, nil)
}

// SetNotice shows a short, transient message in the host status bar (the host clears it on
// the next user input). An add-in uses it to surface state the user can't otherwise see —
// e.g. connection progress or failure.
//
//	client.Interaction().SetNotice("Meeting: connection failed")
//
// mcp:tool interaction_set_notice
// mcp:summary Shows a short, transient message in the host status bar (the host clears it on the next user input).
func (i Interaction) SetNotice(message string) (wire.OKResult, error) {
	return call[wire.OKResult](i.c, wire.MethodInteractionSetNotice, wire.SetNoticeArgs{Message: message})
}
