// SPDX-License-Identifier: Apache-2.0

package types

// TriadSegment is the part of the move/rotate triad the user grabbed — the
// TriadSegmentEnum equivalent (M05-F13, #620).
type TriadSegment uint8

const (
	TriadOrigin  TriadSegment = 0
	TriadXAxis   TriadSegment = 1
	TriadYAxis   TriadSegment = 2
	TriadZAxis   TriadSegment = 3
	TriadXYPlane TriadSegment = 4
	TriadYZPlane TriadSegment = 5
	TriadXZPlane TriadSegment = 6
	TriadXRing   TriadSegment = 7
	TriadYRing   TriadSegment = 8
	TriadZRing   TriadSegment = 9
)

var triadSegmentNames = map[TriadSegment]string{
	TriadOrigin: "origin", TriadXAxis: "x", TriadYAxis: "y", TriadZAxis: "z",
	TriadXYPlane: "xy", TriadYZPlane: "yz", TriadXZPlane: "xz",
	TriadXRing: "rx", TriadYRing: "ry", TriadZRing: "rz",
}

// String returns the segment's stable name.
func (t TriadSegment) String() string {
	if name, ok := triadSegmentNames[t]; ok {
		return name
	}
	return "triadSegment(?)"
}

// TriadMoveType is the motion a drag produces — the TriadMoveTypeEnum equivalent.
type TriadMoveType uint8

const (
	// TriadTranslate slides along one axis.
	TriadTranslate TriadMoveType = 0
	// TriadTranslatePlanar slides in one plane.
	TriadTranslatePlanar TriadMoveType = 1
	// TriadRotate spins around one axis (a ring grab).
	TriadRotate TriadMoveType = 2
	// TriadFree moves with the view plane (the origin grab).
	TriadFree TriadMoveType = 3
)

var triadMoveTypeNames = map[TriadMoveType]string{
	TriadTranslate: "translate", TriadTranslatePlanar: "planar",
	TriadRotate: "rotate", TriadFree: "free",
}

// String returns the move type's stable name.
func (t TriadMoveType) String() string {
	if name, ok := triadMoveTypeNames[t]; ok {
		return name
	}
	return "triadMoveType(?)"
}

// PointInferenceKind is the snap a drag landed on — the PointInferenceEnum
// equivalent, narrowed to what the host infers today (work/datum points; the
// sketch-internal kinds stay in the sketch domain).
type PointInferenceKind uint8

const (
	// InferenceNone is an unsnapped point (the zero value).
	InferenceNone PointInferenceKind = 0
	// InferencePoint snapped to a datum/work point.
	InferencePoint PointInferenceKind = 1
)

var pointInferenceNames = map[PointInferenceKind]string{
	InferenceNone: "none", InferencePoint: "point",
}

// String returns the inference kind's stable name.
func (k PointInferenceKind) String() string {
	if name, ok := pointInferenceNames[k]; ok {
		return name
	}
	return "pointInference(?)"
}
