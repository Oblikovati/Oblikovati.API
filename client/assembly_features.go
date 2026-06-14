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

// AddExtrude extrudes a closed sketch profile (authored on an assembly work plane) into
// every participant — a profiled pocket ("difference") or boss ("union"). E.g.
// AddExtrude(wire.AddAssemblyExtrudeArgs{SketchIndex: 0, ProfileIndex: 0, Distance: 6, Operation: "difference"}).
func (a AssemblyFeatures) AddExtrude(args wire.AddAssemblyExtrudeArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddExtrude, args, &r)
}

// AddRevolve revolves a closed sketch profile (authored on an assembly work plane) about
// the axis line (origin + direction) into every participant — a turned groove
// ("difference") or boss ("union"). Angle is radians in (0,2π] (2π is a full turn). E.g.
// AddRevolve(wire.AddAssemblyRevolveArgs{SketchIndex: 0, ProfileIndex: 0, Origin: [3]float64{0, 0, 0}, Axis: [3]float64{0, 1, 0}, Angle: math.Pi, Operation: "difference"}).
func (a AssemblyFeatures) AddRevolve(args wire.AddAssemblyRevolveArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddRevolve, args, &r)
}

// AddHole drills a hole of the given diameter and depth from center along axis through
// the active assembly's participants — a parametric kind needing no sketch. E.g.
// AddHole(wire.AddAssemblyHoleArgs{Center: [3]float64{5, 5, 0}, Axis: [3]float64{0, 0, 1}, Diameter: 6, Depth: 20}).
func (a AssemblyFeatures) AddHole(args wire.AddAssemblyHoleArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddHole, args, &r)
}

// AddProxyCut adds a feature whose tool is the geometry of the source occurrence,
// supplied as an occurrence-context proxy and re-resolved each rebuild (associative), so
// the machining follows the source. E.g. AddProxyCut(sourceOccurrenceID, "difference").
func (a AssemblyFeatures) AddProxyCut(source uint64, operation string) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	args := wire.AddProxyCutFeatureArgs{Source: source, Operation: operation}
	return r, a.c.call(wire.MethodAssemblyFeaturesAddProxyCut, args, &r)
}

// SetParticipants replaces a feature's participation set with the occurrences named by
// their session ids, e.g. SetParticipants(featureID, []uint64{7, 8}).
func (a AssemblyFeatures) SetParticipants(id uint64, participants []uint64) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	args := wire.SetAssemblyParticipantsArgs{ID: id, Participants: participants}
	return r, a.c.call(wire.MethodAssemblyFeaturesSetParticipants, args, &r)
}

// SetParticipantPaths restricts a feature to specific nested occurrence paths (each a
// sequence of instance names, root first), disambiguating a sub-assembly placed more
// than once; passing no paths clears the restriction. E.g.
// SetParticipantPaths(featureID, [][]string{{"gearbox:1", "bolt:3"}}).
func (a AssemblyFeatures) SetParticipantPaths(id uint64, paths [][]string) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	args := wire.SetAssemblyParticipantPathsArgs{ID: id, Paths: paths}
	return r, a.c.call(wire.MethodAssemblyFeaturesSetParticipantPaths, args, &r)
}

// SetSuppressed suppresses or unsuppresses the named features in one batch and returns
// the refreshed program, e.g. SetSuppressed([]uint64{3}, true).
func (a AssemblyFeatures) SetSuppressed(ids []uint64, suppressed bool) (wire.AssemblyFeaturesResult, error) {
	var r wire.AssemblyFeaturesResult
	args := wire.SetAssemblyFeaturesSuppressedArgs{IDs: ids, Suppressed: suppressed}
	return r, a.c.call(wire.MethodAssemblyFeaturesSetSuppressed, args, &r)
}

// AddChamfer chamfers the given component edges by distance on every participant, e.g.
// AddChamfer(wire.AddAssemblyChamferArgs{Edges: []wire.AssemblyEdgeRef{{Occurrence: o, Edge: key}}, Distance: 2}).
func (a AssemblyFeatures) AddChamfer(args wire.AddAssemblyChamferArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddChamfer, args, &r)
}

// AddFillet rounds the given component edges to radius on every participant, e.g.
// AddFillet(wire.AddAssemblyFilletArgs{Edges: []wire.AssemblyEdgeRef{{Occurrence: o, Edge: key}}, Radius: 1}).
func (a AssemblyFeatures) AddFillet(args wire.AddAssemblyFilletArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddFillet, args, &r)
}

// AddMoveFace translates the given component faces by the vector on every participant, e.g.
// AddMoveFace(wire.AddAssemblyMoveFaceArgs{Faces: []wire.AssemblyFaceRef{{Occurrence: o, Face: key}}, Translation: [3]float64{0, 0, 1}}).
func (a AssemblyFeatures) AddMoveFace(args wire.AddAssemblyMoveFaceArgs) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	return r, a.c.call(wire.MethodAssemblyFeaturesAddMoveFace, args, &r)
}

// Edit sets editable scalars of assembly feature id in place and returns the refreshed
// feature, e.g. Edit(3, []wire.ScalarEdit{{Index: 0, Value: "8 mm"}}) to deepen a pocket.
// Scalar indices come from the feature's Scalars; the whole batch is validated before any
// is applied.
func (a AssemblyFeatures) Edit(id uint64, scalars []wire.ScalarEdit) (wire.AssemblyFeatureResult, error) {
	var r wire.AssemblyFeatureResult
	args := wire.EditAssemblyFeatureArgs{ID: id, Scalars: scalars}
	return r, a.c.call(wire.MethodAssemblyFeaturesEdit, args, &r)
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
