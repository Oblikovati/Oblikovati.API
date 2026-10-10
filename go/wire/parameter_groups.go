// SPDX-License-Identifier: Apache-2.0

package wire

// Custom parameter groups (M02-F05, Oblikovati/Oblikovati#604): named,
// add-in-attributable views over the parameter set. Membership never affects
// parameter semantics — a parameter may belong to several groups, and removing
// it from one neither deletes it nor touches its other memberships.

// ParameterGroupInfo is the JSON shape of one custom parameter group: the
// immutable internal name that identifies it, the editable display name, the
// owning client id (add-in provenance, empty for user-created groups), and the
// member parameter names in collection order.
type ParameterGroupInfo struct {
	InternalName string   `json:"internalName"`
	DisplayName  string   `json:"displayName"`
	ClientID     string   `json:"clientId,omitempty"`
	Members      []string `json:"members,omitempty"`
}

// ListParameterGroupsResult is the response of [MethodParametersGroupsList].
type ListParameterGroupsResult struct {
	Groups []ParameterGroupInfo `json:"groups,omitempty"`
}

// ParameterGroupAddArgs is the request of [MethodParametersGroupsAdd]. An empty
// DisplayName defaults to the internal name.
type ParameterGroupAddArgs struct {
	InternalName string `json:"internalName"`
	DisplayName  string `json:"displayName,omitempty"`
	ClientID     string `json:"clientId,omitempty"`
}

// ParameterGroupDeleteArgs is the request of [MethodParametersGroupsDelete].
// DeleteParameters opts into the cascade that also deletes the member
// parameters; without it only the group goes, the members stay.
type ParameterGroupDeleteArgs struct {
	InternalName     string `json:"internalName"`
	DeleteParameters bool   `json:"deleteParameters,omitempty"`
}

// ParameterGroupDisplayNameArgs is the request of
// [MethodParametersGroupsSetDisplayName]: the group's editable display name.
// The internal name can never change.
type ParameterGroupDisplayNameArgs struct {
	InternalName string `json:"internalName"`
	DisplayName  string `json:"displayName"`
}

// ParameterGroupMemberArgs is the request of [MethodParametersGroupsAddMember]
// / [MethodParametersGroupsRemoveMember]: one parameter (by name) and the group
// it joins or leaves.
type ParameterGroupMemberArgs struct {
	InternalName string `json:"internalName"`
	Parameter    string `json:"parameter"`
}
