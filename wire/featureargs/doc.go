// SPDX-License-Identifier: Apache-2.0

// Package featureargs is the typed, per-kind argument surface for creating features
// (wire.MethodFeaturesAdd). wire.AddFeatureArgs carries the kind plus an opaque
// json.RawMessage; these structs give each kind's arguments a compile-checked shape an
// add-in can import and build, instead of hand-assembling raw JSON against a schema it
// cannot see (ADR-0018 puts request DTOs in api/wire; #1616, audit B5).
//
// Each arg type carries its own kind via [Arg.Kind], so the type IS the binding between
// the struct and the feature it creates — there is no separate lookup table to drift.
// The generic client constructor client.AddFeature marshals one of these and tags the
// envelope from Kind(); the host (addin/opregistry) decodes into the SAME type, so the
// wire shape and the host decoder cannot diverge.
//
// Not every feature kind is promoted yet: the composite kinds (loft, sweep, patterns,
// the fillet/chamfer dress-up family, model tolerance, move-body) and the add-in-
// registered dynamic kinds are still driven by the opaque path and are listed in the
// host's dynamic-kind allowlist. The wire<->host parity guard enforces that every
// registered kind is either promoted here or explicitly allowlisted.
//
//	// build a feature with a typed, compile-checked argument struct:
//	_, err := client.AddFeature(c.Features(), featureargs.Extrude{SketchIndex: 0, Distance: "50 mm"})
package featureargs

// Arg is one feature kind's typed argument struct. Kind reports the feature kind the
// struct creates — the value that fills wire.AddFeatureArgs.Kind — so a caller never
// repeats the kind string and the struct-to-kind mapping lives in exactly one place.
type Arg interface {
	Kind() string
}

// All returns a zero value of every promoted feature-arg type, in a stable order. It is
// the single registry the parity guard and the marshal round-trip test enumerate, so a
// new promoted kind is added in exactly one place.
func All() []Arg {
	args := []Arg{
		Extrude{}, Revolve{}, Rib{}, Emboss{}, Coil{},
		Hole{}, Boss{}, Thread{}, Grill{}, Mesh{}, DirectEdit{},
	}
	// Each family file (#1709) contributes its promoted kinds through a package-level slice, so a
	// new kind is registered in exactly one place next to its struct.
	args = append(args, loftArgs...)
	args = append(args, dressupArgs...)
	args = append(args, modifyArgs...)
	args = append(args, surfaceArgs...)
	args = append(args, patternArgs...)
	args = append(args, advancedArgs...)
	args = append(args, freeformArgs...)
	args = append(args, toleranceArgs...)
	args = append(args, plasticArgs...)
	args = append(args, hullArgs...)
	args = append(args, sheetMetalArgs...)
	return args
}

// Kinds returns the kind string of every promoted feature-arg type (see [All]).
func Kinds() []string {
	args := All()
	kinds := make([]string, len(args))
	for i, a := range args {
		kinds[i] = a.Kind()
	}
	return kinds
}
