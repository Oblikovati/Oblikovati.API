// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati/api/types"

// MeshTranslator is the in-process contract for a foreign-mesh-format translator: it
// reports which formats it can read/write so the host (or an add-in registering its
// own translator) can route an import/export request to the right backend. The GPL
// implementation (model/exchange) satisfies it (compile-time asserted there).
//
// The actual byte-level translation crosses the wire as documents.import /
// documents.export ([oblikovati/api/wire]); this contract is the capability query that
// lets the host pick a translator without hard-coding the format list.
//
// Example:
//
//	if t.CanImport(types.FormatSTL) { /* route an .stl through t */ }
type MeshTranslator interface {
	// Formats lists every format this translator handles (in either direction).
	Formats() []types.ExchangeFormat
	// CanImport reports whether this translator can read the given format.
	CanImport(format types.ExchangeFormat) bool
	// CanExport reports whether this translator can write the given format.
	CanExport(format types.ExchangeFormat) bool
}
