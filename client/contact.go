// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// The assembly contact & interference operation groups (M12-F05, Oblikovati/Oblikovati#362/
// #368): manage contact sets (occurrences that resist interpenetration when dragged), toggle
// the contact solver, and run a static interference analysis (overlapping volumes between
// occurrences).

// ContactSets is the contact-set operation group.
type ContactSets struct{ c *Client }

// ContactSets returns the contact-set operation group.
func (c *Client) ContactSets() ContactSets { return ContactSets{c} }

// Create adds a new contact set named name.
//
// mcp:tool create_contact_set
// mcp:summary Create a contact set (a group of occurrences that resist interpenetration when the contact solver is on). Returns the set.
func (cs ContactSets) Create(name string) (wire.ContactSetResult, error) {
	var r wire.ContactSetResult
	return r, cs.c.call(wire.MethodContactSetsCreate, wire.CreateContactSetArgs{Name: name}, &r)
}

// List returns the assembly's contact sets.
//
// mcp:tool list_contact_sets
// mcp:summary List the assembly's contact sets: each with id, name, and member occurrence ids.
func (cs ContactSets) List() (wire.ContactSetsResult, error) {
	var r wire.ContactSetsResult
	return r, cs.c.call(wire.MethodContactSetsList, struct{}{}, &r)
}

// Delete removes a contact set and returns the remaining sets.
//
// mcp:tool delete_contact_set
// mcp:summary Delete a contact set (id). Returns the remaining sets.
func (cs ContactSets) Delete(id uint64) (wire.ContactSetsResult, error) {
	var r wire.ContactSetsResult
	return r, cs.c.call(wire.MethodContactSetsDelete, wire.ContactSetRef{ID: id}, &r)
}

// AddMember adds an occurrence to a contact set.
//
// mcp:tool add_contact_member
// mcp:summary Add an occurrence (by id) to a contact set (set id). Returns the updated set.
func (cs ContactSets) AddMember(args wire.ContactMemberArgs) (wire.ContactSetResult, error) {
	var r wire.ContactSetResult
	return r, cs.c.call(wire.MethodContactSetsAddMember, args, &r)
}

// RemoveMember removes an occurrence from a contact set.
//
// mcp:tool remove_contact_member
// mcp:summary Remove an occurrence (by id) from a contact set (set id). Returns the updated set.
func (cs ContactSets) RemoveMember(args wire.ContactMemberArgs) (wire.ContactSetResult, error) {
	var r wire.ContactSetResult
	return r, cs.c.call(wire.MethodContactSetsRemoveMember, args, &r)
}

// ContactSolver is the contact-solver toggle group.
type ContactSolver struct{ c *Client }

// ContactSolver returns the contact-solver operation group.
func (c *Client) ContactSolver() ContactSolver { return ContactSolver{c} }

// SetEnabled enables or disables contact enforcement during a drag.
//
// mcp:tool set_contact_solver
// mcp:summary Enable (enabled:true) or disable the contact solver — whether dragging a contact-set member stops at contact with another. Returns the solver state.
func (s ContactSolver) SetEnabled(enabled bool) (wire.ContactSolverResult, error) {
	var r wire.ContactSolverResult
	return r, s.c.call(wire.MethodContactSolverSetEnabled, wire.ContactSolverEnableArgs{Enabled: enabled}, &r)
}

// Status returns whether the contact solver is enabled and how many sets it governs.
//
// mcp:tool contact_solver_status
// mcp:summary Report the contact solver's state: enabled flag and the number of contact sets.
func (s ContactSolver) Status() (wire.ContactSolverResult, error) {
	var r wire.ContactSolverResult
	return r, s.c.call(wire.MethodContactSolverStatus, struct{}{}, &r)
}

// Interference is the static interference-analysis group.
type Interference struct{ c *Client }

// Interference returns the interference-analysis operation group.
func (c *Client) Interference() Interference { return Interference{c} }

// Analyze runs a static interference analysis and returns the overlapping pairs and volumes,
// e.g. Analyze(wire.AnalyzeInterferenceArgs{}) for the whole assembly.
//
// mcp:tool analyze_interference
// mcp:summary Run a static interference analysis over the active assembly (or the occurrences subset). Returns each overlapping pair (occurrence ids + overlap volume + a representative point) and the total overlap volume.
func (i Interference) Analyze(args wire.AnalyzeInterferenceArgs) (wire.InterferenceResultsResult, error) {
	var r wire.InterferenceResultsResult
	return r, i.c.call(wire.MethodInterferenceAnalyze, args, &r)
}
