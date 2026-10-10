// SPDX-License-Identifier: Apache-2.0

package wire

// Assembly feature-program push event (M11-F08, Oblikovati/Oblikovati#725): the host
// announces every re-evaluation of an assembly's machining-feature program — after an
// add, a participation or suppression edit, or a rollback-marker move — so an add-in
// can keep external state consistent without polling. Like the occurrence events
// (#723), it is push-only: the host delivers it to an add-in's Notify entry point and
// the add-in matches on Type, then re-reads detail with assemblyFeatures.list.

// AssemblyFeatureHealth is one feature's post-recompute state in an
// [AssemblyFeaturesChangedEvent]: its id, suppression, and health reason ("" healthy).
type AssemblyFeatureHealth struct {
	ID         uint64 `json:"id"`
	Suppressed bool   `json:"suppressed,omitempty"`
	Health     string `json:"health,omitempty"`
}

// AssemblyFeaturesChangedEvent is the JSON shape of [EventAssemblyFeaturesChanged]:
// the assembly document's id and each feature's resulting health, in program order.
type AssemblyFeaturesChangedEvent struct {
	Type     string                  `json:"type"`
	Document uint64                  `json:"document"`
	Features []AssemblyFeatureHealth `json:"features,omitempty"`
}
