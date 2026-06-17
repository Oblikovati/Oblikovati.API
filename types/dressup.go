// SPDX-License-Identifier: Apache-2.0

package types

// Dress-up / modify vocabulary (M09, Oblikovati/Oblikovati#323 #325 #330 #331
// #332). All numeric blocks are FROZEN to the reference enums (FilletTypeEnum /
// FeatureApproximationTypeEnum / ModelDiameterFromThreadEnum /
// DirectEditOperationTypeEnum / SplitTypeEnum); never renumber.

// FilletType discriminates the fillet definition: edge sets, a face-pair
// fillet, or a full round across three face sets.
type FilletType int32

const (
	EdgeFillet      FilletType = 61697
	FaceFillet      FilletType = 61698
	FullRoundFillet FilletType = 61699
)

var filletTypeNames = map[FilletType]string{
	EdgeFillet:      "edge",
	FaceFillet:      "face",
	FullRoundFillet: "fullRound",
}

// String returns the fillet type's wire spelling.
func (t FilletType) String() string { return enumName(filletTypeNames, t) }

// ParseFilletType resolves a wire spelling back to its type.
func ParseFilletType(s string) (FilletType, bool) { return enumFromName(filletTypeNames, s) }

// FilletCornerType selects how an edge fillet treats a corner where two filleted edges meet at a
// vertex whose third edge stays sharp (the two rolling-ball cylinders cannot be joined by a single
// sphere unless that third edge is also rounded). It is an Oblikovati extension with no reference-API
// equivalent, so the numeric block below is OURS (chosen clear of the frozen reference blocks above)
// — but, once shipped, it is equally frozen: never renumber.
//
//   - FilletCornerMiter — the two cylinders mutually trim along their intersection seam (a crease);
//     the geometrically exact rolling-ball result for rounding only two of a corner's three edges.
//   - FilletCornerSetback — a spherical corner patch tangent to the three faces, set back from the
//     sharp edge by a small planar setback face (a smoothed corner; not the exact rolling-ball form).
//   - FilletCornerRound — round the corner fully into a single sphere octant by also rolling over the
//     third edge (the true 3-edge blend; the corner becomes G1-smooth all round).
type FilletCornerType int32

const (
	// FilletCornerMiter is the default: a crease seam where the two cylinders mutually trim.
	FilletCornerMiter FilletCornerType = 200001
	// FilletCornerSetback is a spherical patch set back from the sharp edge by a planar face.
	FilletCornerSetback FilletCornerType = 200002
	// FilletCornerRound rounds the corner fully into a sphere octant (rolls over the third edge).
	FilletCornerRound FilletCornerType = 200003
)

var filletCornerTypeNames = map[FilletCornerType]string{
	FilletCornerMiter:   "miter",
	FilletCornerSetback: "setback",
	FilletCornerRound:   "round",
}

// String returns the corner type's wire spelling.
func (t FilletCornerType) String() string { return enumName(filletCornerTypeNames, t) }

// ParseFilletCornerType resolves a wire spelling back to its corner type.
func ParseFilletCornerType(s string) (FilletCornerType, bool) {
	return enumFromName(filletCornerTypeNames, s)
}

// FeatureApproximationType is the approximation a thicken / face-offset
// feature may accept when the exact offset is not computable (#331 parity).
// An exact result satisfies every bound below, so a kernel computing the
// exact offset honours any requested approximation.
type FeatureApproximationType int32

const (
	NoApproximation            FeatureApproximationType = 60673
	MeanApproximation          FeatureApproximationType = 60674
	NeverTooThickApproximation FeatureApproximationType = 60675
	NeverTooThinApproximation  FeatureApproximationType = 60676
)

var featureApproximationTypeNames = map[FeatureApproximationType]string{
	NoApproximation:            "none",
	MeanApproximation:          "mean",
	NeverTooThickApproximation: "neverTooThick",
	NeverTooThinApproximation:  "neverTooThin",
}

// String returns the approximation's wire spelling.
func (t FeatureApproximationType) String() string {
	return enumName(featureApproximationTypeNames, t)
}

// ParseFeatureApproximationType resolves a wire spelling back to its type.
func ParseFeatureApproximationType(s string) (FeatureApproximationType, bool) {
	return enumFromName(featureApproximationTypeNames, s)
}

// ModelDiameterFromThread selects which diameter of a thread definition the
// modeled cylindrical face represents (#325 parity).
type ModelDiameterFromThread int32

const (
	ThreadMajorDiameter    ModelDiameterFromThread = 21761
	ThreadMinorDiameter    ModelDiameterFromThread = 21762
	ThreadPitchDiameter    ModelDiameterFromThread = 21763
	ThreadTapDrillDiameter ModelDiameterFromThread = 21764
)

var modelDiameterFromThreadNames = map[ModelDiameterFromThread]string{
	ThreadMajorDiameter:    "major",
	ThreadMinorDiameter:    "minor",
	ThreadPitchDiameter:    "pitch",
	ThreadTapDrillDiameter: "tapDrill",
}

// String returns the model diameter's wire spelling.
func (t ModelDiameterFromThread) String() string { return enumName(modelDiameterFromThreadNames, t) }

// ParseModelDiameterFromThread resolves a wire spelling back to its value.
func ParseModelDiameterFromThread(s string) (ModelDiameterFromThread, bool) {
	return enumFromName(modelDiameterFromThreadNames, s)
}

// DirectEditOperationType discriminates the consolidated direct-edit
// feature's operation (#332).
type DirectEditOperationType int32

const (
	DirectEditMoveOperation   DirectEditOperationType = 105729
	DirectEditSizeOperation   DirectEditOperationType = 105730
	DirectEditRotateOperation DirectEditOperationType = 105731
	DirectEditDeleteOperation DirectEditOperationType = 105732
	// DirectEditUnknownOperation is the reference's placeholder for an
	// operation the API version doesn't know; it is never a valid input.
	DirectEditUnknownOperation DirectEditOperationType = 105733
	DirectEditScaleOperation   DirectEditOperationType = 105734
)

var directEditOperationTypeNames = map[DirectEditOperationType]string{
	DirectEditMoveOperation:    "move",
	DirectEditSizeOperation:    "size",
	DirectEditRotateOperation:  "rotate",
	DirectEditDeleteOperation:  "delete",
	DirectEditUnknownOperation: "unknown",
	DirectEditScaleOperation:   "scale",
}

// String returns the operation's wire spelling.
func (t DirectEditOperationType) String() string { return enumName(directEditOperationTypeNames, t) }

// ParseDirectEditOperationType resolves a wire spelling back to its operation.
func ParseDirectEditOperationType(s string) (DirectEditOperationType, bool) {
	return enumFromName(directEditOperationTypeNames, s)
}

// SplitType is the kind of split a split feature performed: trim a solid to
// one side of the tool, split its faces only (imprint), or split the body
// into separate solids. The reference declares kSplitPart sharing id 32769
// with kTrimSolid; "trimSolid" is the canonical spelling here.
type SplitType int32

const (
	TrimSolidSplit  SplitType = 32769
	SplitFacesSplit SplitType = 32770
	SplitBodySplit  SplitType = 32771
)

var splitTypeNames = map[SplitType]string{
	TrimSolidSplit:  "trimSolid",
	SplitFacesSplit: "splitFaces",
	SplitBodySplit:  "splitBody",
}

// String returns the split type's wire spelling.
func (t SplitType) String() string { return enumName(splitTypeNames, t) }

// ParseSplitType resolves a wire spelling back to its type.
func ParseSplitType(s string) (SplitType, bool) { return enumFromName(splitTypeNames, s) }

// ChamferType discriminates how an edge chamfer's setback is specified (parity:
// ChamferDefinitionType): an equal distance on both faces, a distance on one face plus the
// chamfer-face angle, or two independent distances (asymmetric). The numeric values are
// frozen at the reference API's ids and must never be renumbered.
type ChamferType int32

const (
	// ChamferDistance — equal setback distance along both adjacent faces.
	ChamferDistance ChamferType = 26881
	// ChamferDistanceAndAngle — a distance on the first face and the chamfer-face angle.
	ChamferDistanceAndAngle ChamferType = 26882
	// ChamferTwoDistances — independent setback distances on each adjacent face (asymmetric).
	ChamferTwoDistances ChamferType = 26883
)

var chamferTypeNames = map[ChamferType]string{
	ChamferDistance:         "distance",
	ChamferDistanceAndAngle: "distanceAndAngle",
	ChamferTwoDistances:     "twoDistances",
}

// String returns the chamfer type's wire spelling.
func (t ChamferType) String() string { return enumName(chamferTypeNames, t) }

// ParseChamferType resolves a wire spelling back to its type.
func ParseChamferType(s string) (ChamferType, bool) { return enumFromName(chamferTypeNames, s) }

// ChamferConcaveStrategy selects how an edge chamfer treats a CONCAVE (internal) edge — one
// where the two faces fold over the material so the dihedral exceeds π (e.g. the inside corner
// where a rib meets a plate). A convex edge always cuts its corner and ignores this. An
// Oblikovati extension with no reference-API equivalent, so the numeric block below is OURS
// (chosen clear of the frozen reference blocks) — but, once shipped, equally frozen: never
// renumber.
//
//   - ChamferConcaveOutward — fill the inside corner with material: a flat 45° gusset that
//     bridges the two faces (Boolean union). The DEFAULT; the zero value resolves to it.
//   - ChamferConcaveInward — instead cut a recessed groove into the corner, relieving it
//     (Boolean cut on the material side).
type ChamferConcaveStrategy int32

const (
	// ChamferConcaveOutward fills the inside corner with material (the default, also the zero value).
	ChamferConcaveOutward ChamferConcaveStrategy = 200101
	// ChamferConcaveInward cuts a recessed relief groove into the inside corner instead.
	ChamferConcaveInward ChamferConcaveStrategy = 200102
)

var chamferConcaveStrategyNames = map[ChamferConcaveStrategy]string{
	ChamferConcaveOutward: "outward",
	ChamferConcaveInward:  "inward",
}

// String returns the concave strategy's wire spelling.
func (t ChamferConcaveStrategy) String() string { return enumName(chamferConcaveStrategyNames, t) }

// ParseChamferConcaveStrategy resolves a wire spelling back to its strategy.
func ParseChamferConcaveStrategy(s string) (ChamferConcaveStrategy, bool) {
	return enumFromName(chamferConcaveStrategyNames, s)
}
