// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Theme is the UI-theme operation group: it lets an add-in read the host's active
// theme so its own panels can match the host's colors.
type Theme struct{ c *Client }

// Theme returns the UI-theme operation group.
func (c *Client) Theme() Theme { return Theme{c} }

// Active returns the host's currently selected theme with all its colors.
//
//	t, _ := client.Theme().Active()
//	bg := t.Colors["chrome.window_bg"] // "#1e2127ff"
func (t Theme) Active() (wire.ThemeView, error) {
	var r wire.ThemeView
	return r, t.c.call(wire.MethodThemeActive, nil, &r)
}

// List returns a summary of every available theme (built-in and custom), flagging the
// active one.
func (t Theme) List() (wire.ListThemesResult, error) {
	var r wire.ListThemesResult
	return r, t.c.call(wire.MethodThemeList, nil, &r)
}
