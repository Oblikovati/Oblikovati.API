// SPDX-License-Identifier: Apache-2.0

package wire

// FeatureLifecycleEvent is the JSON shape of the [EventFeatureAdded] / [EventFeatureEdited] /
// [EventFeatureDeleted] push events (#148): a feature was created, modified, or removed on a
// document. Delivered to an add-in's Notify entry point (ADR-0016) with no request/response —
// an add-in matches on Type and reads the affected feature's identity from the payload (so it
// need not diff model.tree). Kind is the feature's operation kind (e.g. "extrude", "circular-
// pattern"); Name is its tree name. The Feature id is stable across rename/reorder.
//
// NOTE (v1 scope): like [EditCommittedEvent], only feature mutations that arrive through the host
// method router (features.add / features.edit / features.delete) are emitted; the batched
// [ModelChangedEvent] remains the coarse signal that also covers other paths.
type FeatureLifecycleEvent struct {
	Type     string `json:"type"` // EventFeatureAdded, EventFeatureEdited, or EventFeatureDeleted
	Document uint64 `json:"document"`
	Feature  uint64 `json:"feature"`
	Name     string `json:"name,omitempty"`
	Kind     string `json:"kind,omitempty"`
}

// SketchEditEvent is the JSON shape of the [EventSketchEditEntered] / [EventSketchEditExited] push
// events (#148): the host entered or left the edit mode of a sketch — the SketchEvents surface. It
// is delivered to an add-in's Notify entry point (ADR-0016); an add-in matches on Type and reads
// the sketch identity from the payload. On exit the Sketch id is the sketch just left. Fires for
// both UI-driven and add-in-driven sketch-edit transitions.
type SketchEditEvent struct {
	Type     string `json:"type"` // EventSketchEditEntered or EventSketchEditExited
	Document uint64 `json:"document"`
	Sketch   uint64 `json:"sketch"`
	Name     string `json:"name,omitempty"`
}
