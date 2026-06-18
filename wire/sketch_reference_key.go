// SPDX-License-Identifier: Apache-2.0

package wire

// Persistent sketch reference keys (Oblikovati/Oblikovati#153). A sketch and each of its
// entities carry a document-scoped UUID that is stable across save/load and edits — unlike
// the transient session id. An add-in stores a key to refer to geometry durably (robust
// feature-editing automations) and rebinds it later with [MethodSketchResolveReference].

// SketchReferenceKeyResult is the response of [MethodSketchReferenceKey]: the persistent
// reference key of the sketch addressed by [SketchArgs.SketchIndex].
type SketchReferenceKeyResult struct {
	ReferenceKey string `json:"referenceKey"`
}

// ResolveSketchReferenceArgs is the request of [MethodSketchResolveReference]: a previously
// obtained persistent key (a sketch's key or a sketch-entity's key).
type ResolveSketchReferenceArgs struct {
	ReferenceKey string `json:"referenceKey"`
}

// ResolveSketchReferenceResult is the response of [MethodSketchResolveReference]. Found is
// false when no current sketch/entity matches the key (a legitimate, non-fatal outcome — the
// referent was deleted). Kind is "sketch" or "sketchEntity"; SketchIndex locates the owning
// sketch; EntityID is the matched entity's current session id (0 when Kind is "sketch").
type ResolveSketchReferenceResult struct {
	Found       bool   `json:"found"`
	Kind        string `json:"kind,omitempty"`
	SketchIndex int    `json:"sketchIndex"`
	EntityID    uint64 `json:"entityId,omitempty"`
}
