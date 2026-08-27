// SPDX-License-Identifier: Apache-2.0

package types

// Accuracy selects the computational accuracy of a property calculation
// (region properties, mass properties). Higher accuracy maps to denser
// sampling/quadrature and costs proportionally more (M06-F08,
// Oblikovati/Oblikovati#623; shared vocabulary with M18-F01 mass properties).
//
// The values are a frozen block matching the reference API's accuracy enum;
// never renumber them.
type Accuracy int32

const (
	AccuracyLow      Accuracy = 69377
	AccuracyMedium   Accuracy = 69378
	AccuracyHigh     Accuracy = 69379
	AccuracyVeryHigh Accuracy = 69380
)

// accuracyNames are the frozen wire spellings.
var accuracyNames = map[Accuracy]string{
	AccuracyLow:      "low",
	AccuracyMedium:   "medium",
	AccuracyHigh:     "high",
	AccuracyVeryHigh: "veryHigh",
}

// String returns the accuracy's wire spelling.
func (a Accuracy) String() string { return enumName(accuracyNames, a, "enum(?)") }

// ParseAccuracy resolves a wire spelling back to its accuracy.
func ParseAccuracy(s string) (Accuracy, bool) { return enumFromName(accuracyNames, s) }
