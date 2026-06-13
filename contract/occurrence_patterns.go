// SPDX-License-Identifier: Apache-2.0

package contract

// OccurrencePatternElement is the scalar read surface of one replicated instance in an
// assembly component pattern (M11-F04, Oblikovati/Oblikovati#729): whether it is
// suppressed and whether its placement has been individually overridden. The element's
// transform is read with the occurrence tree over api/wire. Host:
// occurrence.OccurrencePatternElement.
type OccurrencePatternElement interface {
	// Suppressed reports whether this element is excluded from the model.
	Suppressed() bool
	// Repositioned reports whether this element's placement was individually overridden.
	Repositioned() bool
}

// OccurrencePattern is the scalar read surface of an assembly component pattern — a seed
// component replicated across an arrangement (host: occurrence.OccurrencePattern). The
// generated placements are created/read as occurrences over api/wire
// (assembly.patternCreate, assembly.occurrences).
type OccurrencePattern interface {
	// Count returns the number of elements the arrangement generates (incl. the seed).
	Count() int
}
