// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// AssemblyFeatures is the assembly feature-program operation group (M11-F08,
// Oblikovati/Oblikovati#633/#725): list the machining features authored in the active
// assembly, add one, edit which occurrences a feature participates on, suppress them in
// batch, and move the end-of-features rollback marker.
type AssemblyFeatures struct{ c *Client }

// AssemblyFeatures returns the assembly feature-program operation group.
func (c *Client) AssemblyFeatures() AssemblyFeatures { return AssemblyFeatures{c} }

// List returns the active assembly's feature program and rollback-marker state.
func (a AssemblyFeatures) List() (wire.AssemblyFeaturesResult, error) {
	var r wire.AssemblyFeaturesResult
	return r, a.c.call(wire.MethodAssemblyFeaturesList, struct{}{}, &r)
}

// Add adds a box-tool cut feature to the active assembly, e.g.
// Add(wire.AddAssemblyFeatureArgs{ToolMin: [3]float64{0, 0, 0.5}, ToolMax: [3]float64{1, 1, 2}, Operation: "difference"}).
func (a AssemblyFeatures) Add(args wire.AddAssemblyFeatureArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAdd, args, &r)
}

// SetParticipants replaces a feature's participation set with the occurrences named by
// their session ids, e.g. SetParticipants(featureID, []uint64{7, 8}).
func (a AssemblyFeatures) SetParticipants(id uint64, participants []uint64) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	args := wire.SetAssemblyParticipantsArgs{ID: id, Participants: participants}
	return r, a.c.call(wire.MethodAssemblyFeaturesSetParticipants, args, &r)
}

// SetSuppressed suppresses or unsuppresses the named features in one batch and returns
// the refreshed program, e.g. SetSuppressed([]uint64{3}, true).
func (a AssemblyFeatures) SetSuppressed(ids []uint64, suppressed bool) (wire.AssemblyFeaturesResult, error) {
	var r wire.AssemblyFeaturesResult
	args := wire.SetAssemblyFeaturesSuppressedArgs{IDs: ids, Suppressed: suppressed}
	return r, a.c.call(wire.MethodAssemblyFeaturesSetSuppressed, args, &r)
}

// GetEndOfFeatures returns the active assembly's end-of-features marker state.
func (a AssemblyFeatures) GetEndOfFeatures() (wire.EndOfFeaturesResult, error) {
	var r wire.EndOfFeaturesResult
	return r, a.c.call(wire.MethodAssemblyGetEndOfFeatures, struct{}{}, &r)
}

// SetEndOfFeatures moves the marker to position (negative restores it to the end) and
// returns the refreshed program, e.g. SetEndOfFeatures(1).
func (a AssemblyFeatures) SetEndOfFeatures(position int) (wire.AssemblyFeaturesResult, error) {
	var r wire.AssemblyFeaturesResult
	return r, a.c.call(wire.MethodAssemblySetEndOfFeatures, wire.SetEndOfFeaturesArgs{Position: position}, &r)
}
