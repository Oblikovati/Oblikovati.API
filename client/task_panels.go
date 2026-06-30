// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// TaskPanels is the modal task-panel operation group: show a FreeCAD-Task-style modal panel built
// from declarative controls, with OK/Cancel. The result arrives as a wire.TaskPanelClosedEvent;
// control edits arrive as the ordinary value/references events keyed on the panel id.
type TaskPanels struct{ c *Client }

// TaskPanels returns the modal task-panel operation group.
func (c *Client) TaskPanels() TaskPanels { return TaskPanels{c} }

// Show displays the modal task panel (asynchronous — never blocks; the accept/cancel arrives as a
// wire.TaskPanelClosedEvent).
//
// mcp:tool task_panel_show
// mcp:summary Show a modal task panel (OK/Cancel) built from declarative controls.
func (t TaskPanels) Show(p wire.TaskPanelSpec) (wire.OKResult, error) {
	var r wire.OKResult
	return r, t.c.call(wire.MethodTaskPanelShow, wire.ShowTaskPanelArgs{Panel: p}, &r)
}

// Close dismisses an open task panel programmatically.
//
// mcp:tool task_panel_close
// mcp:summary Dismiss an open task panel programmatically.
func (t TaskPanels) Close(id string) (wire.OKResult, error) {
	var r wire.OKResult
	return r, t.c.call(wire.MethodTaskPanelClose, wire.CloseTaskPanelArgs{ID: id}, &r)
}
