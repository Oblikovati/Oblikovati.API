// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Help is the help-routing operation group (M05-F14): register your add-in's help
// source and open topics through the host, so add-in help behaves like product
// help. Language reports the host locale.
type Help struct{ c *Client }

// Help returns the help-routing operation group.
func (c *Client) Help() Help { return Help{c} }

// RegisterContext declares a help source: a URL prefix or local directory topics
// resolve against.
//
//	client.Help().RegisterContext("com.x.sim", "https://docs.example.org/sim/")
//
// mcp:tool help_register_context
// mcp:summary Declares a help source: a URL prefix or local directory topics resolve against.
func (h Help) RegisterContext(source, base string) (wire.OKResult, error) {
	args := wire.RegisterHelpContextArgs{Source: source, Base: base}
	return call[wire.OKResult](h.c, wire.MethodHelpRegisterContext, args)
}

// Display opens a topic of a registered source ("" ⇒ the host's documentation).
//
// mcp:tool help_display
// mcp:summary Opens a topic of a registered source ("" ⇒ the host's documentation).
func (h Help) Display(source, topic string) (wire.OKResult, error) {
	return call[wire.OKResult](h.c, wire.MethodHelpDisplay, wire.DisplayHelpArgs{Source: source, Topic: topic})
}

// Path returns a source's registered base.
//
// mcp:tool help_path
// mcp:summary Returns a source's registered base.
func (h Help) Path(source string) (wire.HelpPathResult, error) {
	return call[wire.HelpPathResult](h.c, wire.MethodHelpPath, wire.DisplayHelpArgs{Source: source})
}

// LanguageInfo returns the host's locale as a BCP-47 tag.
//
// mcp:tool language_info
// mcp:summary Returns the host's locale as a BCP-47 tag.
func (h Help) LanguageInfo() (wire.LanguageInfoResult, error) {
	return call[wire.LanguageInfoResult](h.c, wire.MethodLanguageInfo, nil)
}
