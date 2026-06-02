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
