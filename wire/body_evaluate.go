// SPDX-License-Identifier: Apache-2.0

package wire

// Face-evaluation modes for [FaceEvaluateArgs.Mode]. They select which query the host runs
// over the supplied inputs and which result arrays come back. pointAtParam, normalAtParam and
// tangents take (u, v) parameter pairs; closestPoint takes (x, y, z) points and projects each
// onto the face.
const (
	FaceEvalPointAtParam  = "pointAtParam"  // (u,v) → Points
	FaceEvalNormalAtParam = "normalAtParam" // (u,v) → Points + Normals
	FaceEvalTangents      = "tangents"      // (u,v) → Points + UTangents + VTangents
	FaceEvalClosestPoint  = "closestPoint"  // (x,y,z) → Points + UVs (the projected params)
)

// FaceEvaluateArgs is the request of [MethodBodyFaceEvaluate]: a batched evaluation of one
// face's surface, addressed by reference key. Inputs is a flat array — (u, v) parameter pairs
// for the parametric modes, or (x, y, z) point triples (database units, cm) for closestPoint.
// Batching matters: callers that sample a surface densely (surface-following toolpaths,
// drop-cutter projection) issue one request, not one per point.
type FaceEvaluateArgs struct {
	BodyIndex int       `json:"bodyIndex"`
	FaceKey   string    `json:"faceKey"`
	Mode      string    `json:"mode"`
	Inputs    []float64 `json:"inputs,omitempty"`
}

// FaceEvaluateResult is the response of [MethodBodyFaceEvaluate]. Only the arrays relevant to
// the requested mode are populated; each is flat and parallel to the inputs (xyz triples for
// points/normals/tangents, uv pairs for UVs). ParamRange is always returned as
// [uMin, vMin, uMax, vMax] (infinite surface domains are clamped to a finite window). Points
// and tangents are in database units (cm); normals are unit vectors.
type FaceEvaluateResult struct {
	Points     []float64 `json:"points,omitempty"`
	Normals    []float64 `json:"normals,omitempty"`
	UTangents  []float64 `json:"uTangents,omitempty"`
	VTangents  []float64 `json:"vTangents,omitempty"`
	UVs        []float64 `json:"uvs,omitempty"`
	ParamRange []float64 `json:"paramRange"`
}
