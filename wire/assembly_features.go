// SPDX-License-Identifier: Apache-2.0

package wire

// Assembly feature program surface (M11-F08, Oblikovati/Oblikovati#633/#725): the
// features authored in the assembly that machine placed component geometry in place.
// Features are addressed by a stable id (like the part features.* methods); each
// participates on a set of component occurrences, addressed by their session ids (the
// same ids the occurrence push events carry, #723). The end-of-features marker rolls
// the program back, suppressing the trailing features.

// AssemblyFeatureInfo renders one assembly feature: its stable id, type, display name,
// suppression, health (empty when healthy), and the session ids of the occurrences it
// machines.
type AssemblyFeatureInfo struct {
	ID           uint64   `json:"id"`
	Kind         string   `json:"kind"`
	Name         string   `json:"name"`
	Suppressed   bool     `json:"suppressed,omitempty"`
	Health       string   `json:"health,omitempty"`
	Participants []uint64 `json:"participants,omitempty"`
	// ParticipantPaths is the feature's nested-path restriction (each path a sequence of
	// occurrence instance names, root first), present only when the feature is restricted
	// to specific placements of a sub-assembly placed more than once. Empty means it
	// machines every path through a participating leaf occurrence (the default).
	ParticipantPaths [][]string `json:"participantPaths,omitempty"`
}

// AssemblyFeaturesResult is the reply of [MethodAssemblyFeaturesList]: the feature
// program in order plus the end-of-features marker state.
type AssemblyFeaturesResult struct {
	Features      []AssemblyFeatureInfo `json:"features"`
	EndOfFeatures int                   `json:"endOfFeatures"`
	RolledBack    bool                  `json:"rolledBack"`
}

// AssemblyFeatureResult is the reply of the single-feature mutators (add,
// setParticipants): the affected feature's refreshed info.
type AssemblyFeatureResult struct {
	Feature AssemblyFeatureInfo `json:"feature"`
}

// AddAssemblyFeatureArgs is the request of [MethodAssemblyFeaturesAdd]: add a cut
// feature whose tool is the axis-aligned box [ToolMin,ToolMax] in the assembly's
// space, machined into every participating occurrence with Operation (a
// [types.BooleanType] spelling: "difference" cuts, "union" adds, "intersect" keeps the
// common volume). The new feature defaults to participating on every component present.
type AddAssemblyFeatureArgs struct {
	ToolMin   [3]float64 `json:"toolMin"`
	ToolMax   [3]float64 `json:"toolMax"`
	Operation string     `json:"operation"`
}

// SetAssemblyParticipantsArgs is the request of [MethodAssemblyFeaturesSetParticipants]:
// replace the participation set of feature ID with the occurrences named by their
// session ids. An unknown occurrence id is rejected.
type SetAssemblyParticipantsArgs struct {
	ID           uint64   `json:"id"`
	Participants []uint64 `json:"participants"`
}

// AddAssemblyExtrudeArgs is the request of [MethodAssemblyFeaturesAddExtrude]: extrude
// the ProfileIndex-th closed region of the active assembly's SketchIndex-th sketch
// (authored on an assembly work plane) by Distance (document units) into every
// participant, applying Operation (a [types.BooleanType] spelling: "difference" cuts a
// pocket, "union" adds a boss). The assembly sketching subsystem (#739) supplies the
// profile; Distance must be positive.
type AddAssemblyExtrudeArgs struct {
	SketchIndex  int     `json:"sketchIndex"`
	ProfileIndex int     `json:"profileIndex"`
	Distance     float64 `json:"distance"`
	Operation    string  `json:"operation"`
}

// AddAssemblyHoleArgs is the request of [MethodAssemblyFeaturesAddHole]: drill a hole
// of Diameter and Depth (document units) from Center along Axis (a direction in the
// assembly's space) through every participating occurrence — a parametric assembly
// feature kind that needs no sketch (M11-F08 kind set, #735). Diameter and Depth must
// be positive and Axis non-zero.
type AddAssemblyHoleArgs struct {
	Center   [3]float64 `json:"center"`
	Axis     [3]float64 `json:"axis"`
	Diameter float64    `json:"diameter"`
	Depth    float64    `json:"depth"`
}

// AddProxyCutFeatureArgs is the request of [MethodAssemblyFeaturesAddProxyCut]: add a
// feature whose tool is supplied as an occurrence-context proxy — the geometry of the
// Source occurrence (by session id), resolved into assembly space and re-resolved on
// every rebuild, so the machining follows the source as it moves or changes (M11-F08
// proxy inputs, #734). Operation is a [types.BooleanType] spelling ("difference" cuts).
// The source occurrence is excluded from the new feature's default participation (a
// component does not machine itself).
type AddProxyCutFeatureArgs struct {
	Source    uint64 `json:"source"`
	Operation string `json:"operation"`
}

// SetAssemblyParticipantPathsArgs is the request of
// [MethodAssemblyFeaturesSetParticipantPaths]: restrict feature ID to the given nested
// occurrence paths (each a sequence of instance names, root first), disambiguating a
// sub-assembly placed more than once. Passing no paths clears the restriction so the
// feature again machines every path through a participating leaf occurrence.
type SetAssemblyParticipantPathsArgs struct {
	ID    uint64     `json:"id"`
	Paths [][]string `json:"paths"`
}

// SetAssemblyFeaturesSuppressedArgs is the request of
// [MethodAssemblyFeaturesSetSuppressed]: suppress or unsuppress every named feature in
// one batch.
type SetAssemblyFeaturesSuppressedArgs struct {
	IDs        []uint64 `json:"ids"`
	Suppressed bool     `json:"suppressed"`
}

// EndOfFeaturesResult is the reply of [MethodAssemblyGetEndOfFeatures]: the marker
// position (-1 at the end) and whether the assembly is currently rolled back.
type EndOfFeaturesResult struct {
	Position   int  `json:"position"`
	RolledBack bool `json:"rolledBack"`
}

// SetEndOfFeaturesArgs is the request of [MethodAssemblySetEndOfFeatures]: move the
// marker to Position (negative restores it to the end, re-including every feature).
type SetEndOfFeaturesArgs struct {
	Position int `json:"position"`
}
