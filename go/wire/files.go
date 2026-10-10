// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The file surface (M03-F07, Oblikovati/Oblikovati#608): the file as an object
// distinct from the documents it contains, with its own identity and the
// persisted "as-saved" records of its file-to-file references. Today one .obk
// file hosts one document; the shapes already allow several (model states and
// library members later).

// FileInfo is the response of [MethodFilesGet]: one file's identity and load
// state. InternalName is the stable GUID minted at creation; RevisionID stamps
// the file's content as of its last save, DatabaseRevisionID stamps only the
// model (geometry/reference) content, so an unchanged model keeps its database
// revision across property-only saves.
type FileInfo struct {
	FullFileName       string   `json:"fullFileName"`
	InternalName       string   `json:"internalName"`
	RevisionID         string   `json:"revisionId"`
	DatabaseRevisionID string   `json:"databaseRevisionId"`
	SaveCounter        int      `json:"saveCounter"`
	VersionCreated     string   `json:"versionCreated,omitempty"`
	VersionSaved       string   `json:"versionSaved,omitempty"`
	Loaded             bool     `json:"loaded"`
	Referenced         bool     `json:"referenced"`
	Documents          []uint64 `json:"documents,omitempty"`
}

// GetFileArgs is the request of [MethodFilesGet] and [MethodFilesListReferences]:
// the full file name of an open file.
type GetFileArgs struct {
	FullFileName string `json:"fullFileName"`
}

// FileReferenceInfo is one persisted file-to-file reference record of
// [MethodFilesListReferences] — the as-saved logical name (relative + library),
// where it resolved this session, and the broken-reference detail flags that
// explain a non-upToDate status.
type FileReferenceInfo struct {
	FullFileName           string                 `json:"fullFileName"`
	RelativeFileName       string                 `json:"relativeFileName,omitempty"`
	LibraryName            string                 `json:"libraryName,omitempty"`
	LocationType           types.FileLocationType `json:"locationType"`
	ResolvedFileName       string                 `json:"resolvedFileName,omitempty"`
	ReferencedInternalName string                 `json:"referencedInternalName,omitempty"`
	SaveCounter            int                    `json:"saveCounter,omitempty"`
	Status                 types.ReferenceStatus  `json:"status"`
	Missing                bool                   `json:"missing,omitempty"`
	Replaced               bool                   `json:"replaced,omitempty"`
	LocationDifferent      bool                   `json:"locationDifferent,omitempty"`
	InternalNameDifferent  bool                   `json:"internalNameDifferent,omitempty"`
}

// ListFileReferencesResult is the response of [MethodFilesListReferences].
type ListFileReferencesResult struct {
	References []FileReferenceInfo `json:"references"`
}

// ReplaceFileReferenceArgs is the request of [MethodFilesReplaceReference]: the
// repair operation re-pointing one reference of FullFileName — the record whose
// as-saved name is RequestedName — at NewFileName. The owning file reports the
// reference as replaced until it is saved.
type ReplaceFileReferenceArgs struct {
	FullFileName  string `json:"fullFileName"`
	RequestedName string `json:"requestedName"`
	NewFileName   string `json:"newFileName"`
}

// DocumentFileReferenceInfo is one entry of [MethodDocumentsListFileReferences]:
// the document-side view of a file reference, bridging to the resolved document
// when one was found.
type DocumentFileReferenceInfo struct {
	DisplayName       string                `json:"displayName,omitempty"`
	FullFileName      string                `json:"fullFileName"`
	Status            types.ReferenceStatus `json:"status"`
	DocumentFound     bool                  `json:"documentFound"`
	DifferentDocument bool                  `json:"differentDocument,omitempty"`
}

// ListDocumentFileReferencesArgs is the request of
// [MethodDocumentsListFileReferences]: the session id of the document whose
// file references to enumerate.
type ListDocumentFileReferencesArgs struct {
	Document uint64 `json:"document"`
}

// ListDocumentFileReferencesResult is the response of
// [MethodDocumentsListFileReferences].
type ListDocumentFileReferencesResult struct {
	References []DocumentFileReferenceInfo `json:"references"`
}
