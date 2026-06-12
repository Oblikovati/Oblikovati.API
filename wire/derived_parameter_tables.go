// SPDX-License-Identifier: Apache-2.0

package wire

// Derived parameter tables (M02-F06, Oblikovati/Oblikovati#605): the
// connection objects that link parameters from another document into this
// one. A table records the source document, the chosen linked subset, and
// owns the resulting read-only derived parameters; when the source changes,
// the linked values follow, and a removed source parameter turns its derived
// counterpart sick rather than silently vanishing.

// DerivedParameterTableInfo is the JSON shape of one table: its stable id,
// the source document (full document name), the linked source-parameter
// names, the source parameters available to link (the candidates), and the
// table's health — "" when every link resolves, otherwise a reason (e.g. the
// source document is missing or a linked parameter was removed at the source).
type DerivedParameterTableInfo struct {
	ID             int      `json:"id"`
	SourceDocument string   `json:"sourceDocument"`
	Linked         []string `json:"linked,omitempty"`
	Available      []string `json:"available,omitempty"`
	Health         string   `json:"health,omitempty"`
}

// ListDerivedParameterTablesResult is the response of
// [MethodParametersDerivedTablesList].
type ListDerivedParameterTablesResult struct {
	Tables []DerivedParameterTableInfo `json:"tables,omitempty"`
}

// DerivedParameterTableAddArgs is the request of
// [MethodParametersDerivedTablesAdd]: the source document and the source
// parameter names to link. An empty Linked links nothing yet — the table
// still records the connection and lists candidates.
type DerivedParameterTableAddArgs struct {
	SourceDocument string   `json:"sourceDocument"`
	Linked         []string `json:"linked,omitempty"`
}

// DerivedParameterTableSetLinkedArgs is the request of
// [MethodParametersDerivedTablesSetLinked]: it replaces the table's linked
// subset — newly linked names gain derived parameters, unlinked ones lose
// theirs.
type DerivedParameterTableSetLinkedArgs struct {
	ID     int      `json:"id"`
	Linked []string `json:"linked,omitempty"`
}

// DerivedParameterTableDeleteArgs is the request of
// [MethodParametersDerivedTablesDelete]. Deleting a table deletes its derived
// parameters; a table owned by a derived component cannot be deleted directly
// (delete the component instead).
type DerivedParameterTableDeleteArgs struct {
	ID int `json:"id"`
}
