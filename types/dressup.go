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
