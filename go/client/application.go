// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Application is the host-application info operation group: read-only facts an add-in
// can query about the running host. It is the ThisApplication equivalent for the
// version surface — distinct from the add-in registry (see [Client.AddIns]).
type Application struct{ c *Client }

// Application returns the host-application info operation group.
func (c *Client) Application() Application { return Application{c} }

// ApiVersion returns the semantic version of the api contract the running host
// implements. The load-time handshake already guarantees the major matches, so an
// add-in uses this only to adapt to minor/patch differences within that major.
//
//	v, _ := client.Application().ApiVersion()
//	if v.Major == 0 { /* pre-1.0: surface MAY change between minors */ }
//
// mcp:tool application_api_version
// mcp:summary Returns the semantic version (version string + major) of the Oblikovati API contract the running host implements.
func (g Application) ApiVersion() (wire.ApplicationApiVersionResult, error) {
	return call[wire.ApplicationApiVersionResult](g.c, wire.MethodApplicationApiVersion, nil)
}
