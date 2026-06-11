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
func (g ClientApplications) Register(name string) (wire.RegisterClientApplicationResult, error) {
	var r wire.RegisterClientApplicationResult
	args := wire.RegisterClientApplicationArgs{Name: name}
	return r, g.c.call(wire.MethodClientAppsRegister, args, &r)
}

// Unregister removes a previously registered external client.
func (g ClientApplications) Unregister(id int) (wire.OKResult, error) {
	var r wire.OKResult
	return r, g.c.call(wire.MethodClientAppsUnregister, wire.UnregisterClientApplicationArgs{ID: id}, &r)
}

// List returns the registered external clients in registration order.
func (g ClientApplications) List() (wire.ListClientApplicationsResult, error) {
	var r wire.ListClientApplicationsResult
	return r, g.c.call(wire.MethodClientAppsList, nil, &r)
}
