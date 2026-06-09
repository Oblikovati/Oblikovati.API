// SPDX-License-Identifier: Apache-2.0

package wire

// AddSketchPatternArgs is the request of [MethodSketchAddPattern] — the discriminated
// sketch-pattern constructor. Kind is a
// [oblikovati.org/api/types.SketchPatternKind]; Entities are the seed selection.
//
//   - rectangular: Count1/Count2 instances stepped by Spacing1/Spacing2 (unit-bearing
//     lengths) along Dir1/Dir2 ([x,y] directions; default [1,0] and [0,1]). The seed is
//     cell (0,0), so the copies number Count1·Count2 − 1.
//   - circular: Count instances (incl. the seed) spread over Angle (a unit-bearing angle)
//     about Center ([x,y] cm); the copies number Count − 1.
type AddSketchPatternArgs struct {
	SketchIndex int      `json:"sketchIndex"`
	Kind        string   `json:"kind"`
	Entities    []uint64 `json:"entities"`

	Count1   int       `json:"count1,omitempty"`
	Count2   int       `json:"count2,omitempty"`
	Spacing1 string    `json:"spacing1,omitempty"`
	Spacing2 string    `json:"spacing2,omitempty"`
	Dir1     []float64 `json:"dir1,omitempty"`
	Dir2     []float64 `json:"dir2,omitempty"`

	Count  int       `json:"count,omitempty"`
	Angle  string    `json:"angle,omitempty"`
	Center []float64 `json:"center,omitempty"`
}

// AddSketchPatternResult is the response of [MethodSketchAddPattern]: the ids of the
// created (non-seed) copy entities.
type AddSketchPatternResult struct {
	Created []uint64 `json:"created"`
}
