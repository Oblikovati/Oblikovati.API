// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Fonts is the font operation group: it lists the faces a text/emboss can use — the host's
// bundled faces plus the OS-installed fonts (ADR-0031).
type Fonts struct{ c *Client }

// Fonts returns the font operation group.
func (c *Client) Fonts() Fonts { return Fonts{c} }

// List returns every selectable face (bundled "embedded" + host "system"); a system face
// carries the file path whose bytes are embedded into the document when it is chosen.
//
//	faces, _ := client.Fonts().List()
func (f Fonts) List() (wire.ListFontsResult, error) {
	var r wire.ListFontsResult
	return r, f.c.call(wire.MethodFontsList, nil, &r)
}
