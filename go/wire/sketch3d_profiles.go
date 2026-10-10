// SPDX-License-Identifier: Apache-2.0

package wire

// Profile3DInfo is one enumerated profile from [MethodSketch3DProfiles]: a closed, planar
// loop of the 3D sketch. Index feeds a planar-section feature; Area is the enclosed area
// (model cm²); Normal is the loop plane's unit normal [x,y,z]; Vertices is the number of
// loop vertices.
type Profile3DInfo struct {
	Index    int       `json:"index"`
	Area     float64   `json:"area"`
	Normal   []float64 `json:"normal"`
	Vertices int       `json:"vertices"`
}

// ListProfiles3DResult is the response of [MethodSketch3DProfiles].
type ListProfiles3DResult struct {
	Profiles []Profile3DInfo `json:"profiles"`
}

// Path3DInfo is one enumerated path from [MethodSketch3DPaths]: a maximal connected chain
// of line/arc segments — a sweep/loft rail. Index identifies it; Closed reports whether
// it forms a loop; Points is the number of ordered vertices.
type Path3DInfo struct {
	Index  int  `json:"index"`
	Closed bool `json:"closed"`
	Points int  `json:"points"`
}

// ListPaths3DResult is the response of [MethodSketch3DPaths].
type ListPaths3DResult struct {
	Paths []Path3DInfo `json:"paths"`
}
