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
func (f Files) Get(fullFileName string) (wire.FileInfo, error) {
	var r wire.FileInfo
	return r, f.c.call(wire.MethodFilesGet, wire.GetFileArgs{FullFileName: fullFileName}, &r)
}

// References returns the file's persisted file-to-file reference records.
func (f Files) References(fullFileName string) (wire.ListFileReferencesResult, error) {
	var r wire.ListFileReferencesResult
	return r, f.c.call(wire.MethodFilesListReferences, wire.GetFileArgs{FullFileName: fullFileName}, &r)
}

// ReplaceReference re-points one reference of a file at a new target (the
// broken-reference repair), returning the updated record.
func (f Files) ReplaceReference(args wire.ReplaceFileReferenceArgs) (wire.FileReferenceInfo, error) {
	var r wire.FileReferenceInfo
	return r, f.c.call(wire.MethodFilesReplaceReference, args, &r)
}
