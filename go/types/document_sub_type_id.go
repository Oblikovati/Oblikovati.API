// SPDX-License-Identifier: Apache-2.0

package types

// DocumentSubTypeID refines a document's base DocumentType with a flavored
// sub-type: a sheet-metal part is a part document carrying the sheet-metal
// sub-type id (M03-F11, Oblikovati/Oblikovati#612). The empty id is a plain
// document. Add-ins register their own ids over documents.registerSubType
// (M05-F15); ids under [ReservedSubTypePrefix] are built-in and cannot be
// registered by clients.
type DocumentSubTypeID string

// ReservedSubTypePrefix marks built-in sub-type ids; client registration of
// ids under this prefix is rejected.
const ReservedSubTypePrefix = "org.oblikovati."

const (
	// SubTypePlain is the absent sub-type: an unflavored document.
	SubTypePlain DocumentSubTypeID = ""
	// SubTypeSheetMetalPart flavors a part document as sheet metal — reserved
	// now so .obk files persist the discriminator from the first release; the
	// sheet-metal environment itself lands with M20.
	SubTypeSheetMetalPart DocumentSubTypeID = "org.oblikovati.part.sheetMetal"
)

// BuiltIn reports whether the id is reserved by the host (or plain).
func (id DocumentSubTypeID) BuiltIn() bool {
	return id == SubTypePlain || len(id) > len(ReservedSubTypePrefix) &&
		string(id[:len(ReservedSubTypePrefix)]) == ReservedSubTypePrefix
}
