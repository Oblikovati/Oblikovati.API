// SPDX-License-Identifier: Apache-2.0

package types

// AddInKind classifies an installed add-in — the ApplicationAddInTypeEnum
// equivalent. The host treats kinds identically today; the classification exists so
// UI (an add-in manager dialog) and tooling can group entries, and so translator
// add-ins can be routed to the exchange layer when that seam opens (M05-F01, #245).
type AddInKind uint8

const (
	// StandardAddIn is a general-purpose extension (the zero value — every add-in
	// without a declared kind).
	StandardAddIn AddInKind = 0
	// TranslatorAddIn contributes import/export format support (a
	// [oblikovati.org/api/contract.MeshTranslator] behind the wire boundary).
	TranslatorAddIn AddInKind = 1
)

var addInKindNames = map[AddInKind]string{
	StandardAddIn: "standard", TranslatorAddIn: "translator",
}

// String returns the kind's stable name ("standard", "translator").
func (k AddInKind) String() string {
	if name, ok := addInKindNames[k]; ok {
		return name
	}
	return "addInKind(?)"
}
