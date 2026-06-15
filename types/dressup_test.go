// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

// The M09 dress-up enum blocks are FROZEN to the reference ids and wire
// spellings (Oblikovati/Oblikovati#323 #325 #330 #331 #332). SplitType's
// "trimSolid" carries 32769, the id the reference declares twice
// (kSplitPart/kTrimSolid).

func TestFilletTypeFrozenBlock(t *testing.T) {
	want := map[FilletType]string{61697: "edge", 61698: "face", 61699: "fullRound"}
	assertFrozenBlock(t, "FilletType", want, filletTypeNames, ParseFilletType)
}

func TestFeatureApproximationTypeFrozenBlock(t *testing.T) {
	want := map[FeatureApproximationType]string{
		60673: "none", 60674: "mean", 60675: "neverTooThick", 60676: "neverTooThin",
	}
	assertFrozenBlock(t, "FeatureApproximationType", want, featureApproximationTypeNames,
		ParseFeatureApproximationType)
}

func TestModelDiameterFromThreadFrozenBlock(t *testing.T) {
	want := map[ModelDiameterFromThread]string{
		21761: "major", 21762: "minor", 21763: "pitch", 21764: "tapDrill",
	}
	assertFrozenBlock(t, "ModelDiameterFromThread", want, modelDiameterFromThreadNames,
		ParseModelDiameterFromThread)
}

func TestDirectEditOperationTypeFrozenBlock(t *testing.T) {
	want := map[DirectEditOperationType]string{
		105729: "move", 105730: "size", 105731: "rotate",
		105732: "delete", 105733: "unknown", 105734: "scale",
	}
	assertFrozenBlock(t, "DirectEditOperationType", want, directEditOperationTypeNames,
		ParseDirectEditOperationType)
}

func TestSplitTypeFrozenBlock(t *testing.T) {
	want := map[SplitType]string{32769: "trimSolid", 32770: "splitFaces", 32771: "splitBody"}
	assertFrozenBlock(t, "SplitType", want, splitTypeNames, ParseSplitType)
}

func TestChamferTypeFrozenBlock(t *testing.T) {
	want := map[ChamferType]string{26881: "distance", 26882: "distanceAndAngle", 26883: "twoDistances"}
	assertFrozenBlock(t, "ChamferType", want, chamferTypeNames, ParseChamferType)
}
