// SPDX-License-Identifier: Apache-2.0

package contract

// ConstraintGlyphVisibility is a sketch's geometric-constraint display state.
// Relation IDs are stable within their owning sketch.
// Example: `shown := visibility.ConstraintGlyphsShown()`.
type ConstraintGlyphVisibility interface {
	// ConstraintGlyphsShown reports whether the global display switch is on.
	ConstraintGlyphsShown() bool
	// ConstraintGlyphVisible reports whether a current relation's glyph is visible.
	ConstraintGlyphVisible(id uint64) (bool, error)
}
