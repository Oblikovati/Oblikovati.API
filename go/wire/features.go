// SPDX-License-Identifier: Apache-2.0

package wire

import "encoding/json"

// FeatureKindInfo is the self-describing entry for one feature operation: its name,
// summary, and JSON Schema for arguments — enough for a client (or an LLM) to call
// [MethodFeaturesAdd].
type FeatureKindInfo struct {
	Kind    string          `json:"kind"`
	Summary string          `json:"summary"`
	Schema  json.RawMessage `json:"schema"`
}

// ListFeatureKindsResult is the response of [MethodFeaturesList].
type ListFeatureKindsResult struct {
	Kinds []FeatureKindInfo `json:"kinds"`
}

// AddFeatureArgs is the request of [MethodFeaturesAdd]: a feature kind and its
// operation-specific arguments (an opaque JSON object validated by the kind's
// schema).
type AddFeatureArgs struct {
	Kind string          `json:"kind"`
	Args json.RawMessage `json:"args"`
}

// FeatureRefArgs is the request of [MethodFeaturesGet] and [MethodFeaturesDelete]:
// one placed feature, addressed by the stable id from [FeatureInfo] (model.tree).
// The id survives rename and reorder; an index does not, so the wire never uses one
// to address a feature.
type FeatureRefArgs struct {
	ID uint64 `json:"id"`
}

// FeatureScalar describes one editable scalar of a placed feature — a distance,
// radius, angle, or pattern count — mirroring [WorkPlaneScalar]: its slot Index,
// Label, the Unit its value is shown in ("mm", "deg", …), and the current Value in
// that unit. Integer marks a whole-number input (a pattern count). An edit sets it
// via [ScalarEdit] keyed on Index.
type FeatureScalar struct {
	Index   int     `json:"index"`
	Label   string  `json:"label"`
	Unit    string  `json:"unit,omitempty"`
	Value   float64 `json:"value"`
	Integer bool    `json:"integer,omitempty"`
}

// FeatureSlot describes one re-pickable geometric input of a placed feature — the edges
// of a fillet, the removed faces of a shell, a hole's face, an extrude's profile, a
// mirror's plane. Index addresses it in a [FeatureRepick]; Kind ("edges"|"faces"|"face"|
// "profile"|"plane") says what reference it accepts; Multi marks a slot that accumulates
// several references (edges/faces) and can be cleared; Count is how many it currently holds.
type FeatureSlot struct {
	Index int    `json:"index"`
	Label string `json:"label"`
	Kind  string `json:"kind"`
	Multi bool   `json:"multi,omitempty"`
	Count int    `json:"count"`
}

// FeatureDetail is one placed feature with its history position and the editable inputs
// [MethodFeaturesEdit] accepts: scalar fields ([FeatureScalar]) and re-pickable geometry
// slots ([FeatureSlot]). Both are empty for features whose definition exposes nothing of
// that kind (e.g. a cosmetic feature has neither).
type FeatureDetail struct {
	FeatureInfo
	Index   int             `json:"index"`
	Scalars []FeatureScalar `json:"scalars,omitempty"`
	Slots   []FeatureSlot   `json:"slots,omitempty"`
}

// FeatureDetailResult is the response of [MethodFeaturesGet], [MethodFeaturesEdit],
// [MethodFeaturesRename], [MethodFeaturesSetSuppressed], and [MethodFeaturesReorder]:
// the feature's refreshed state after the call (post-recompute health included).
type FeatureDetailResult struct {
	Feature FeatureDetail `json:"feature"`
}

// EditFeatureArgs is the request of [MethodFeaturesEdit]: edit the feature in place — set
// editable scalars AND/OR re-pick its geometric references. Every edit (scalar parse, slot
// resolution) is validated before ANY is applied, so a failed batch (bad value, unbindable
// key, wrong slot kind) leaves the definition untouched; then the part recomputes once.
// Scalar indices come from [FeatureDetail.Scalars]; Repick slot indices from
// [FeatureDetail.Slots].
type EditFeatureArgs struct {
	ID      uint64          `json:"id"`
	Scalars []ScalarEdit    `json:"scalars,omitempty"`
	Repick  []FeatureRepick `json:"repick,omitempty"`
}

// FeatureRepick re-points one geometric Slot (a [FeatureSlot] index) of a placed feature.
// The fields read by the slot's Kind:
//   - edges/faces/face: Ref is a topology reference key from [MethodModelReferenceKeys].
//   - profile: SketchIndex + ProfileIndex name a sketch region (from model.tree / sketch.profiles).
//   - plane: Ref is a planar-face key, a work-plane ref ("plane/N"), or an origin plane ("origin/plane/xy").
//
// Clear empties a clearable multi-slot (edges/faces) and ignores the other fields.
type FeatureRepick struct {
	Slot         int    `json:"slot"`
	Ref          string `json:"ref,omitempty"`
	SketchIndex  int    `json:"sketchIndex,omitempty"`
	ProfileIndex int    `json:"profileIndex,omitempty"`
	Clear        bool   `json:"clear,omitempty"`
}

// DeleteFeatureResult is the response of [MethodFeaturesDelete].
type DeleteFeatureResult struct {
	ID      uint64 `json:"id"`
	Deleted bool   `json:"deleted"`
}

// RenameFeatureArgs is the request of [MethodFeaturesRename]. The id is stable
// across renames; the new name must be non-empty and unique within the part.
type RenameFeatureArgs struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// SetFeatureSuppressedArgs is the request of [MethodFeaturesSetSuppressed]: set
// (not toggle) explicit suppression, so the call is idempotent for replication.
type SetFeatureSuppressedArgs struct {
	ID         uint64 `json:"id"`
	Suppressed bool   `json:"suppressed"`
}

// ReorderFeatureArgs is the request of [MethodFeaturesReorder]: move the feature to
// NewIndex in history order (0-based, from model.tree). A move that would place a
// feature before one it depends on is rejected.
type ReorderFeatureArgs struct {
	ID       uint64 `json:"id"`
	NewIndex int    `json:"newIndex"`
}
