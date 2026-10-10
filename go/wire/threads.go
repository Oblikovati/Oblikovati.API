// SPDX-License-Identifier: Apache-2.0

package wire

// Thread table query and designation resolution (M09-F01 PBI-101,
// Oblikovati/Oblikovati#325): the one source of truth for thread data that
// the thread feature, hole tapping (#326), and drawings (M14) consume.

// ThreadTableQueryArgs is the request of [MethodThreadsTableQuery]. The query
// is progressive: with no filters the result lists the thread types; given
// ThreadType it also lists that type's nominal sizes; given NominalSize the
// designations; given Designation the classes. Internal narrows class
// listings to the internal (nut) or external (bolt) side.
type ThreadTableQueryArgs struct {
	Internal    bool   `json:"internal,omitempty"`
	ThreadType  string `json:"threadType,omitempty"`
	NominalSize string `json:"nominalSize,omitempty"`
	Designation string `json:"designation,omitempty"`
}

// ThreadTableQueryResult is the response of [MethodThreadsTableQuery]. Each
// level is filled when its prerequisite filter was given (ThreadTypes always).
type ThreadTableQueryResult struct {
	ThreadTypes  []string `json:"threadTypes"`
	NominalSizes []string `json:"nominalSizes,omitempty"`
	Designations []string `json:"designations,omitempty"`
	Classes      []string `json:"classes,omitempty"`
}

// ResolveThreadArgs is the request of [MethodThreadsResolve]: a designation
// (e.g. "M8x1.25", "1/4-20") with the optional tolerance class, handedness,
// internal/external side, and the tapered (pipe thread) flag carried per the
// reference's StandardThreadInfo/TaperedThreadInfo split.
type ResolveThreadArgs struct {
	Designation string `json:"designation"`
	Class       string `json:"class,omitempty"`
	Internal    bool   `json:"internal,omitempty"`
	LeftHanded  bool   `json:"leftHanded,omitempty"`
	Tapered     bool   `json:"tapered,omitempty"`
}

// ThreadInfoResult is the response of [MethodThreadsResolve]: the resolved
// thread data. Diameters and pitch are millimetres. PitchDiameter is the ISO
// basic pitch diameter; TapDrillDiameter the drill for a tapped hole of this
// thread (≈ the minor diameter).
type ThreadInfoResult struct {
	Designation      string  `json:"designation"`
	ThreadType       string  `json:"threadType"`
	NominalSize      string  `json:"nominalSize"`
	Class            string  `json:"class,omitempty"`
	Metric           bool    `json:"metric"`
	Internal         bool    `json:"internal"`
	RightHanded      bool    `json:"rightHanded"`
	Tapered          bool    `json:"tapered"`
	Pitch            float64 `json:"pitch"`
	MajorDiameter    float64 `json:"majorDiameter"`
	MinorDiameter    float64 `json:"minorDiameter"`
	PitchDiameter    float64 `json:"pitchDiameter"`
	TapDrillDiameter float64 `json:"tapDrillDiameter"`
}
