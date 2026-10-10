// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// ClientApplications is the external-client registry operation group: out-of-process
// automation drivers (an MCP client, a test harness, a companion app) announce
// themselves here so the session can list who is driving it — the ClientApplications
// equivalent, distinct from in-process add-ins.
type ClientApplications struct{ c *Client }

// ClientApplications returns the external-client registry operation group.
func (c *Client) ClientApplications() ClientApplications { return ClientApplications{c} }

// Register announces an external client by display name and returns the
// session-unique id to pass to [ClientApplications.Unregister] on disconnect.
//
//	reg, _ := client.ClientApplications().Register("acme-pipeline")
//	defer client.ClientApplications().Unregister(reg.ID)
//
// mcp:tool client_apps_register
// mcp:summary Announces an external client by display name and returns the session-unique id to pass to [ClientApplications.Unregister] on disconnect.
func (g ClientApplications) Register(name string) (wire.RegisterClientApplicationResult, error) {
	args := wire.RegisterClientApplicationArgs{Name: name}
	return call[wire.RegisterClientApplicationResult](g.c, wire.MethodClientAppsRegister, args)
}

// Unregister removes a previously registered external client.
//
// mcp:tool client_apps_unregister
// mcp:summary Removes a previously registered external client.
func (g ClientApplications) Unregister(id int) (wire.OKResult, error) {
	return call[wire.OKResult](g.c, wire.MethodClientAppsUnregister, wire.UnregisterClientApplicationArgs{ID: id})
}

// List returns the registered external clients in registration order.
//
// mcp:tool client_apps_list
// mcp:summary Returns the registered external clients in registration order.
func (g ClientApplications) List() (wire.ListClientApplicationsResult, error) {
	return call[wire.ListClientApplicationsResult](g.c, wire.MethodClientAppsList, nil)
}
