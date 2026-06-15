// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Assembly derive/shrinkwrap requests (M11-F06, Oblikovati/Oblikovati#631/#716):
// derive a source assembly into the active part as a base body, or simplify it into a
// lightweight shrinkwrap body. Each create returns the new feature as a
// [FeatureDetailResult] (id, kind, health, history index), so the caller can address
// it later — e.g. to break its link. Break-link takes that id.

// DeriveCreateArgs is the request of [MethodAssemblyDeriveCreate]: derive the open
// assembly document Source into the active part as a base body, merging every
// component (the include-all derive). Source is a document session id from
// documents.list; it must name an open assembly document.
type DeriveCreateArgs struct {
	Source uint64 `json:"source"`
}

// ShrinkwrapCreateArgs is the request of [MethodAssemblyShrinkwrapCreate]: derive the
// open assembly document Source into the active part as a simplified, lightweight base
// body. RemoveStyle/MinPartVolume drop parts before merging (MinPartVolume in document
// units³ applies to RemoveSmallParts); EnvelopeStyle replaces the kept parts with
// bounding-box proxies; PatchHoles fills internal voids first; MaxHoleDiameter (when > 0,
// document units) caps surface-opening through-holes/pockets no wider than it, closing them
// flush while keeping the real outer geometry. The zero options reduce to a plain include-all
// derive.
type ShrinkwrapCreateArgs struct {
	Source          uint64                        `json:"source"`
	RemoveStyle     types.ShrinkwrapRemoveStyle   `json:"removeStyle,omitempty"`
	MinPartVolume   float64                       `json:"minPartVolume,omitempty"`
	EnvelopeStyle   types.ShrinkwrapEnvelopeStyle `json:"envelopeStyle,omitempty"`
	PatchHoles      bool                          `json:"patchHoles,omitempty"`
	MaxHoleDiameter float64                       `json:"maxHoleDiameter,omitempty"`
}

// DeriveBreakLinkArgs is the request of [MethodAssemblyDeriveBreakLink]: freeze and
// sever the source link of the derived-assembly or shrinkwrap feature with this id
// (from model.tree), so the part keeps the current geometry without further updates.
type DeriveBreakLinkArgs struct {
	ID uint64 `json:"id"`
}

// DeriveStatusArgs is the request of [MethodAssemblyDeriveStatus] and
// [MethodAssemblyDeriveUpdate]: address a derive-family feature (derived-assembly,
// derived-part, or shrinkwrap) by its id (from model.tree).
type DeriveStatusArgs struct {
	ID uint64 `json:"id"`
}

// DeriveStatusResult reports whether a derived component is out of date relative to its
// source document — the reference API's drive state. OutOfDate is true when the source's
// current recipe revision (CurrentRevision) differs from the one captured when the derive
// was last created/updated (SavedRevision); it is the basis for the "out of date" badge.
// Linked is false after a break-link (the source is frozen). SourceDocument is the source's
// full document name. CurrentRevision is empty when the source is not currently resolvable.
type DeriveStatusResult struct {
	OutOfDate       bool   `json:"outOfDate"`
	Linked          bool   `json:"linked"`
	SourceDocument  string `json:"sourceDocument"`
	SavedRevision   string `json:"savedRevision,omitempty"`
	CurrentRevision string `json:"currentRevision,omitempty"`
}
