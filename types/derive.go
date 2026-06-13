// SPDX-License-Identifier: Apache-2.0

package types

// Derive and shrinkwrap value types for the assembly→part derive surface (M11-F06,
// Oblikovati/Oblikovati#631/#716). A derived assembly pulls a source assembly's bodies
// into a part as one base body; a shrinkwrap additionally simplifies that result by
// removing parts and replacing the rest with envelopes.

// DeriveStyle controls how one source occurrence contributes to a derived assembly's
// base body — the reference API's per-occurrence derive style. The zero value is
// DeriveInclude, so an unstyled occurrence is merged in.
type DeriveStyle int32

const (
	// DeriveInclude merges the occurrence's bodies into the derived base.
	DeriveInclude DeriveStyle = iota
	// DeriveExclude omits the occurrence entirely.
	DeriveExclude
	// DeriveSubtract cuts the occurrence's bodies from the merged base.
	DeriveSubtract
)

var deriveStyleNames = map[DeriveStyle]string{
	DeriveInclude:  "include",
	DeriveExclude:  "exclude",
	DeriveSubtract: "subtract",
}

// String returns the derive style's wire spelling.
func (s DeriveStyle) String() string { return enumName(deriveStyleNames, s) }

// ParseDeriveStyle resolves a wire spelling back to its DeriveStyle.
func ParseDeriveStyle(s string) (DeriveStyle, bool) { return enumFromName(deriveStyleNames, s) }

// ShrinkwrapRemoveStyle selects which source parts a shrinkwrap drops before merging —
// the reference API's shrinkwrap remove style. Removal makes the result lighter by
// discarding parts a viewer never sees or that are too small to matter.
type ShrinkwrapRemoveStyle int32

const (
	// RemoveNone keeps every part (the full, unsimplified set).
	RemoveNone ShrinkwrapRemoveStyle = iota
	// RemoveSmallParts drops parts whose body volume is below a threshold.
	RemoveSmallParts
	// RemoveInternalParts drops parts fully enclosed by other parts.
	RemoveInternalParts
)

var shrinkwrapRemoveStyleNames = map[ShrinkwrapRemoveStyle]string{
	RemoveNone:          "none",
	RemoveSmallParts:    "smallParts",
	RemoveInternalParts: "internalParts",
}

// String returns the remove style's wire spelling.
func (s ShrinkwrapRemoveStyle) String() string { return enumName(shrinkwrapRemoveStyleNames, s) }

// ParseShrinkwrapRemoveStyle resolves a wire spelling back to its ShrinkwrapRemoveStyle.
func ParseShrinkwrapRemoveStyle(s string) (ShrinkwrapRemoveStyle, bool) {
	return enumFromName(shrinkwrapRemoveStyleNames, s)
}

// ShrinkwrapEnvelopeStyle selects how kept parts are replaced by simpler proxy
// geometry — the reference API's envelopes-replace style. Envelopes erase internal
// detail (and, by construction, holes) so the result is a lightweight closed solid.
type ShrinkwrapEnvelopeStyle int32

const (
	// EnvelopeNone keeps each kept part's real geometry.
	EnvelopeNone ShrinkwrapEnvelopeStyle = iota
	// EnvelopePerPart replaces each kept part with its axis-aligned bounding box.
	EnvelopePerPart
	// EnvelopeWhole replaces the entire kept set with one axis-aligned bounding box.
	EnvelopeWhole
)

var shrinkwrapEnvelopeStyleNames = map[ShrinkwrapEnvelopeStyle]string{
	EnvelopeNone:    "none",
	EnvelopePerPart: "perPart",
	EnvelopeWhole:   "whole",
}

// String returns the envelope style's wire spelling.
func (s ShrinkwrapEnvelopeStyle) String() string { return enumName(shrinkwrapEnvelopeStyleNames, s) }

// ParseShrinkwrapEnvelopeStyle resolves a wire spelling back to its ShrinkwrapEnvelopeStyle.
func ParseShrinkwrapEnvelopeStyle(s string) (ShrinkwrapEnvelopeStyle, bool) {
	return enumFromName(shrinkwrapEnvelopeStyleNames, s)
}
