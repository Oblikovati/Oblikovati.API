// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Model is the read-only model-inspection operation group.
type Model struct{ c *Client }

// Model returns the model-inspection operation group.
func (c *Client) Model() Model { return Model{c} }

// Tree returns a read-only snapshot of the active part's structure.
func (m Model) Tree() (wire.ModelTreeResult, error) {
	var r wire.ModelTreeResult
	return r, m.c.call(wire.MethodModelTree, nil, &r)
}

// Selection returns the current selection summary.
func (m Model) Selection() (wire.SelectionResult, error) {
	var r wire.SelectionResult
	return r, m.c.call(wire.MethodModelSelection, nil, &r)
}
