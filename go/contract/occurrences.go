// SPDX-License-Identifier: Apache-2.0

package contract

// ComponentOccurrence is the scalar read surface of one placement of a component in an
// assembly (M11-F01/F02, Oblikovati/Oblikovati#728): its session id, instance name, and
// per-instance state. The placement transform and every mutation travel over api/wire
// (the assembly.* methods); this is what an in-proc consumer reads directly. The host
// implementation lives in /source (occurrence.Occurrence).
type ComponentOccurrence interface {
	// ID is the occurrence's session id, unique within its owning collection.
	ID() uint64
	// Name is the instance name, e.g. "pin:1".
	Name() string
	// Suppressed reports whether the occurrence is excluded from the model.
	Suppressed() bool
	// Grounded reports whether the occurrence is fixed in the assembly's space.
	Grounded() bool
	// Adaptive reports whether the occurrence's geometry may flex to satisfy constraints.
	Adaptive() bool
	// IsSubstitute reports whether this is a substitute (simplified) representation.
	IsSubstitute() bool
}

// ComponentOccurrences is the scalar read surface of an assembly's occurrence collection —
// the components placed directly in it (host: occurrence.Occurrences). The tree itself is
// read over api/wire (assembly.occurrences).
type ComponentOccurrences interface {
	// Count returns the number of occurrences in the collection.
	Count() int
}

// AssemblyComponentDefinition is the scalar read surface of an assembly document's model: a
// version string that advances whenever the occurrence structure or any placement changes,
// so a consumer can tell when to re-read the occurrence tree (over assembly.occurrences).
// Host: compdef.AssemblyComponentDefinition.
type AssemblyComponentDefinition interface {
	// ModelGeometryVersion advances on every structural or placement change to the assembly.
	ModelGeometryVersion() string
}
