// SPDX-License-Identifier: Apache-2.0

package wire

// FeatureInfo is the JSON shape of one feature in the model tree. Health is empty
// when the feature is healthy.
type FeatureInfo struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Suppressed bool   `json:"suppressed"`
	Health     string `json:"health,omitempty"`
}

// ModelTreeResult is the response of [MethodModelTree]: a read-only snapshot of the
// active part's structure (parameter names, sketch count, feature program, body
// count).
type ModelTreeResult struct {
	Document   string        `json:"document"`
	Parameters []string      `json:"parameters"`
	Sketches   int           `json:"sketches"`
	Features   []FeatureInfo `json:"features"`
	Bodies     int           `json:"bodies"`
}

// SelectionResult is the response of [MethodModelSelection]: how many entities are
// selected, their selection kinds, and — parallel to Kinds — each entity's work-feature
// reference for entities that have one, empty otherwise. A reference is one of:
//   - a datum plane/axis/point key (a work-feature key);
//   - "face/<url-base64(key)>"   — a picked B-rep face;
//   - "vertex/<url-base64(key)>" — a picked B-rep vertex;
//   - "body/<url-base64(key)>"   — a picked whole body, the same recompute-stable key as
//     [BodyInfo.Key] (#1492), so an add-in can tell which body was selected.
//
// A client reads Refs to feed a selected face/point/plane into [MethodWorkPlanesCreate], or a
// selected body into a body-scoped operation; all forms round-trip through [MethodModelSelect].
type SelectionResult struct {
	Count int      `json:"count"`
	Kinds []int    `json:"kinds"`
	Refs  []string `json:"refs"`
}

// SelectArgs is the request of [MethodModelSelect] (#157): select the entities named by Refs —
// reference strings as returned in a [SelectionResult] (face/vertex/body, #1492). Mode "add"
// extends the current selection; any other value (or empty) replaces it. The reply is the new
// selection.
type SelectArgs struct {
	Refs []string `json:"refs"`
	Mode string   `json:"mode,omitempty"`
}

// DeselectArgs is the request of [MethodModelDeselect] (#157): remove the named entities from the
// current selection (references not currently selected are ignored). The reply is the new selection.
type DeselectArgs struct {
	Refs []string `json:"refs"`
}

// TopologyRef identifies one piece of part topology by its persistent reference key, with a
// representative point [x,y,z] (a face's range-box centre, an edge's midpoint, a vertex's
// position) so a caller can recognise which entity it is. Kind is the geometry classification
// for faces — "plane" | "cylinder" | "cone" | "sphere" | "torus" | "spline" — so a caller can
// pick, say, the cylindrical face to thread (empty for edges/vertices).
type TopologyRef struct {
	Key   string    `json:"key"`
	Point []float64 `json:"point"`
	Kind  string    `json:"kind,omitempty"`
}

// BodyTopology groups one body's faces/edges/vertices by reference key.
type BodyTopology struct {
	Faces    []TopologyRef `json:"faces"`
	Edges    []TopologyRef `json:"edges"`
	Vertices []TopologyRef `json:"vertices"`
}

// ReferenceKeysResult is the response of [MethodModelReferenceKeys]: the active part's
// topology per body, each entity carrying the persistent reference key consumed by the key
// consumers (include, addSurfaceCurve, project geometry, attributes). It is how an add-in
// obtains a face/edge/vertex key without a viewport pick.
type ReferenceKeysResult struct {
	Bodies []BodyTopology `json:"bodies"`
}
