// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Custom parameter groups (M02-F05, Oblikovati/Oblikovati#604), on the
// Parameters operation group.

// ListGroups returns the active document's custom parameter groups (part or assembly) with their
// members, in creation order.
//
// mcp:tool parameters_groups_list
// mcp:summary Returns the active document's custom parameter groups (part or assembly) with their members, in creation order.
func (p Parameters) ListGroups() (wire.ListParameterGroupsResult, error) {
	return call[wire.ListParameterGroupsResult](p.c, wire.MethodParametersGroupsList, nil)
}

// AddGroup creates an empty custom group keyed by an immutable internal name;
// an empty display name defaults to it.
//
// mcp:tool parameters_groups_add
// mcp:summary Creates an empty custom group keyed by an immutable internal name; an empty display name defaults to it.
func (p Parameters) AddGroup(args wire.ParameterGroupAddArgs) (wire.ParameterGroupInfo, error) {
	return call[wire.ParameterGroupInfo](p.c, wire.MethodParametersGroupsAdd, args)
}

// DeleteGroup removes a group; args.DeleteParameters opts into also deleting
// the member parameters (otherwise the members stay, only the group goes).
//
// mcp:tool parameters_groups_delete
// mcp:summary Removes a group; args.DeleteParameters opts into also deleting the member parameters (otherwise the members stay, only the group goes).
func (p Parameters) DeleteGroup(args wire.ParameterGroupDeleteArgs) error {
	return p.c.call(wire.MethodParametersGroupsDelete, args, nil)
}

// SetGroupDisplayName edits a group's display name (the internal name can
// never change) and returns the updated group.
//
// mcp:tool parameters_groups_set_display_name
// mcp:summary Edits a group's display name (the internal name can never change) and returns the updated group.
func (p Parameters) SetGroupDisplayName(args wire.ParameterGroupDisplayNameArgs) (wire.ParameterGroupInfo, error) {
	return call[wire.ParameterGroupInfo](p.c, wire.MethodParametersGroupsSetDisplayName, args)
}

// AddGroupMember adds a parameter to a group (membership in other groups is
// untouched) and returns the updated group.
//
// mcp:tool parameters_groups_add_member
// mcp:summary Adds a parameter to a group (membership in other groups is untouched) and returns the updated group.
func (p Parameters) AddGroupMember(args wire.ParameterGroupMemberArgs) (wire.ParameterGroupInfo, error) {
	return call[wire.ParameterGroupInfo](p.c, wire.MethodParametersGroupsAddMember, args)
}

// RemoveGroupMember detaches a parameter from a group — the parameter itself
// is kept — and returns the updated group.
//
// mcp:tool parameters_groups_remove_member
// mcp:summary Detaches a parameter from a group — the parameter itself is kept — and returns the updated group.
func (p Parameters) RemoveGroupMember(args wire.ParameterGroupMemberArgs) (wire.ParameterGroupInfo, error) {
	return call[wire.ParameterGroupInfo](p.c, wire.MethodParametersGroupsRemoveMember, args)
}
