// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Files is the file-surface operation group (M03-F07): file identity, the
// persisted file-to-file reference records, and reference repair.
type Files struct{ c *Client }

// Files returns the file-surface operation group.
func (c *Client) Files() Files { return Files{c} }

// Get returns one open file's identity and load state.
//
//	info, err := client.Files().Get("/work/bracket.obk")
//
// mcp:tool files_get
// mcp:summary Returns one open file's identity and load state.
func (f Files) Get(fullFileName string) (wire.FileInfo, error) {
	return call[wire.FileInfo](f.c, wire.MethodFilesGet, wire.GetFileArgs{FullFileName: fullFileName})
}

// References returns the file's persisted file-to-file reference records.
//
// mcp:tool files_list_references
// mcp:summary Returns the file's persisted file-to-file reference records.
func (f Files) References(fullFileName string) (wire.ListFileReferencesResult, error) {
	return call[wire.ListFileReferencesResult](f.c, wire.MethodFilesListReferences, wire.GetFileArgs{FullFileName: fullFileName})
}

// ReplaceReference re-points one reference of a file at a new target (the
// broken-reference repair), returning the updated record.
//
// mcp:tool files_replace_reference
// mcp:summary Re-points one reference of a file at a new target (the broken-reference repair), returning the updated record.
func (f Files) ReplaceReference(args wire.ReplaceFileReferenceArgs) (wire.FileReferenceInfo, error) {
	return call[wire.FileReferenceInfo](f.c, wire.MethodFilesReplaceReference, args)
}
