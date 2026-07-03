// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Display is the display options/settings operation group: it lets an add-in read and tune the
// application display options and a document's per-document display settings (background, edge
// color, ground plane, shadows) — the objects that parameterize the display modes.
type Display struct{ c *Client }

// Display returns the display options/settings operation group.
func (c *Client) Display() Display { return Display{c} }

// Options returns the application-level display options.
//
//	o, _ := client.Display().Options()
//
// mcp:tool get_display_options
// mcp:summary Read the application display options (quality, transition time, edges, ray tracing).
func (d Display) Options() (wire.DisplayModeOptionsView, error) {
	return call[wire.DisplayModeOptionsView](d.c, wire.MethodDisplayGetOptions, nil)
}

// SetOptions applies the application-level display options, returning the stored result.
//
// mcp:tool set_display_options
// mcp:summary Set the application display options; see get_display_options for the shape.
func (d Display) SetOptions(v wire.DisplayModeOptionsView) (wire.DisplayModeOptionsView, error) {
	return call[wire.DisplayModeOptionsView](d.c, wire.MethodDisplaySetOptions, v)
}

// Settings returns the active document's per-document display settings.
//
//	s, _ := client.Display().Settings()
//
// mcp:tool get_display_settings
// mcp:summary Read a document's display settings (background, edge color, ground plane, shadows).
func (d Display) Settings() (wire.DisplaySettingsView, error) {
	return call[wire.DisplaySettingsView](d.c, wire.MethodDocumentGetDisplaySettings, wire.GetDisplaySettingsArgs{})
}

// SetSettings applies the active document's per-document display settings.
//
// mcp:tool set_display_settings
// mcp:summary Set a document's display settings; see get_display_settings for the shape.
func (d Display) SetSettings(v wire.DisplaySettingsView) (wire.DisplaySettingsView, error) {
	return call[wire.DisplaySettingsView](d.c, wire.MethodDocumentSetDisplaySettings, wire.SetDisplaySettingsArgs{Settings: v})
}
