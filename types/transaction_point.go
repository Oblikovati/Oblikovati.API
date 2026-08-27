// SPDX-License-Identifier: Apache-2.0

package types

// TransactionPoint identifies a position on a document's transaction stream
// relative to the cursor — which step a transaction event or navigation acts
// on. A committed edit acts on the current point, undo acts on the previous
// one, redo on the next (M04-F05, Oblikovati/Oblikovati#613).
//
// The values are a frozen block matching the reference API's transaction-point
// enum; never renumber them.
type TransactionPoint int32

const (
	// TransactionPointUnknown is the sentinel for "no particular step" (e.g.
	// a whole stream being deleted on document close).
	TransactionPointUnknown TransactionPoint = 3585
	// TransactionPointNext is the step ahead of the cursor (what redo acts on).
	TransactionPointNext TransactionPoint = 3586
	// TransactionPointPrevious is the step behind the cursor (what undo acts on).
	TransactionPointPrevious TransactionPoint = 3587
	// TransactionPointCurrent is the step at the cursor (what a commit produced).
	TransactionPointCurrent TransactionPoint = 3588
	// TransactionPointUpToSpecified addresses a span of steps up to a named one
	// (multi-step undo navigation).
	TransactionPointUpToSpecified TransactionPoint = 3589
)

// transactionPointNames are the frozen wire spellings.
var transactionPointNames = map[TransactionPoint]string{
	TransactionPointUnknown:       "unknown",
	TransactionPointNext:          "next",
	TransactionPointPrevious:      "previous",
	TransactionPointCurrent:       "current",
	TransactionPointUpToSpecified: "upToSpecified",
}

// String returns the transaction point's wire spelling.
func (p TransactionPoint) String() string { return enumName(transactionPointNames, p, "enum(?)") }

// ParseTransactionPoint resolves a wire spelling back to its TransactionPoint.
func ParseTransactionPoint(s string) (TransactionPoint, bool) {
	return enumFromName(transactionPointNames, s)
}
