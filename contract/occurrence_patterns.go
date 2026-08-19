// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

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

// OccurrencePattern is the scalar read surface of a PERSISTENT assembly component pattern — a
// seed component replicated across an arrangement, re-readable and editable after creation
// (#1976, host: occurrence.OccurrencePattern). The generated placements are read as occurrences
// over api/wire (assembly.patternCreate, assembly.occurrences).
type OccurrencePattern interface {
	// ID returns the pattern's session id, stable for the document lifetime.
	ID() uint64
	// Name returns the pattern's display name.
	Name() string
	// Count returns the number of elements the arrangement generates (incl. the seed).
	Count() int
	// Suppression reports whether none, some, or all of the elements are suppressed.
	Suppression() types.OccurrencePatternSuppression
}

// OccurrencePatterns is the read surface of an assembly's persistent occurrence patterns
// (#1976, host: occurrence.OccurrencePatternSet).
type OccurrencePatterns interface {
	// Count returns how many patterns the assembly holds.
	Count() int
	// Item returns the i-th pattern in creation order.
	Item(i int) OccurrencePattern
}
