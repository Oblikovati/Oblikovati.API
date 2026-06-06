// SPDX-License-Identifier: Apache-2.0

package types

// AssetSource says where an appearance or material asset comes from, which fixes its
// edit policy and resolution priority. A document-embedded copy is authoritative for
// portability, then the project library (the shared catalog), then the shipped
// built-ins (read-only). This follows the standard library/document asset distinction.
//
// This is the canonical, Apache-2.0 definition; the GPL model aliases it
// (model/material.Source = types.AssetSource).
type AssetSource string

const (
	// AssetBuiltin is a shipped, read-only catalog asset.
	AssetBuiltin AssetSource = "builtin"
	// AssetProject is a user asset in the active project's shared library.
	AssetProject AssetSource = "project"
	// AssetDocument is a copy embedded in a document (so the .obk is self-contained).
	AssetDocument AssetSource = "document"
)

// Editable reports whether the user may modify an asset of this source (everything but
// the read-only built-ins).
func (s AssetSource) Editable() bool { return s != AssetBuiltin }
