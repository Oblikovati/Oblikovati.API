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
// selected and their selection kinds.
type SelectionResult struct {
	Count int   `json:"count"`
	Kinds []int `json:"kinds"`
}
