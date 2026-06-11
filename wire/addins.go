// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"

	"oblikovati.org/api/types"
)

// AddInManifest is the JSON document a shared-library add-in embeds and returns from
// ObkAddInManifest (see include/oblikovati_addin.h) — the add-in's self-description.
// The host parses it into the registry entry it serves from addins.list, so the
// manifest is the single place an add-in declares its identity.
type AddInManifest struct {
	ID           string          `json:"id"`
	DisplayName  string          `json:"displayName"`
	Version      string          `json:"version"`
	Description  string          `json:"description,omitempty"`
	Kind         types.AddInKind `json:"kind,omitempty"`
	Capabilities []string        `json:"capabilities,omitempty"`
}

// AddInInfo is one registry entry of [MethodAddInsList] / [MethodAddInsGet]: the
// manifest identity plus the host-side runtime state (the ApplicationAddIn
// equivalent). Location is the shared-library path the add-in was loaded from —
// empty for a first-party in-process add-in. HasAutomation reports whether
// [MethodAddInsCallAutomation] can target this add-in.
type AddInInfo struct {
	ID            string                  `json:"id"`
	DisplayName   string                  `json:"displayName,omitempty"`
	Version       string                  `json:"version,omitempty"`
	Description   string                  `json:"description,omitempty"`
	Kind          types.AddInKind         `json:"kind,omitempty"`
	LoadBehavior  types.AddInLoadBehavior `json:"loadBehavior,omitempty"`
	Activated     bool                    `json:"activated"`
	Location      string                  `json:"location,omitempty"`
	Capabilities  []string                `json:"capabilities,omitempty"`
	HasAutomation bool                    `json:"hasAutomation,omitempty"`
}

// ListAddInsResult is the response of [MethodAddInsList], in registration order.
type ListAddInsResult struct {
	AddIns []AddInInfo `json:"addIns"`
}

// AddInRefArgs names the add-in a registry method targets — the request of
// [MethodAddInsGet], [MethodAddInsActivate] and [MethodAddInsDeactivate].
type AddInRefArgs struct {
	ID string `json:"id"`
}

// SetAddInLoadBehaviorArgs is the request of [MethodAddInsSetLoadBehavior]: persist
// when the host should activate the add-in on future startups. Setting LoadDisabled
// does not deactivate a running add-in — call [MethodAddInsDeactivate] for that.
type SetAddInLoadBehaviorArgs struct {
	ID           string                  `json:"id"`
	LoadBehavior types.AddInLoadBehavior `json:"loadBehavior"`
}

// CallAddInAutomationArgs is the request of [MethodAddInsCallAutomation]: invoke a
// method on another add-in's automation surface (ApplicationAddIn.Automation). Args
// is the target's own request shape, passed through opaquely — the host routes, it
// does not interpret.
type CallAddInAutomationArgs struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	Args   json.RawMessage `json:"args,omitempty"`
}

// CallAddInAutomationResult is the response of [MethodAddInsCallAutomation]: the
// target add-in's reply, passed through opaquely.
type CallAddInAutomationResult struct {
	Result json.RawMessage `json:"result,omitempty"`
}

// ClientApplicationInfo is one entry of [MethodClientAppsList]: an external client
// application (an out-of-process automation driver, e.g. an MCP client connected
// through the bridge) registered for the session — the ClientApplications
// equivalent, distinct from in-process add-ins.
type ClientApplicationInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RegisterClientApplicationArgs is the request of [MethodClientAppsRegister].
type RegisterClientApplicationArgs struct {
	Name string `json:"name"`
}

// RegisterClientApplicationResult is the response of [MethodClientAppsRegister]: the
// session-unique id the client passes to [MethodClientAppsUnregister] when it
// disconnects.
type RegisterClientApplicationResult struct {
	ID int `json:"id"`
}

// UnregisterClientApplicationArgs is the request of [MethodClientAppsUnregister].
type UnregisterClientApplicationArgs struct {
	ID int `json:"id"`
}

// ListClientApplicationsResult is the response of [MethodClientAppsList], in
// registration order.
type ListClientApplicationsResult struct {
	Clients []ClientApplicationInfo `json:"clients"`
}
