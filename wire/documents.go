// SPDX-License-Identifier: Apache-2.0

package wire

// DocumentInfo is the JSON shape of an open document.
type DocumentInfo struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Dirty   bool   `json:"dirty"`
	Visible bool   `json:"visible"`
	Active  bool   `json:"active"`
}

// ListDocumentsResult is the response of [MethodDocumentsList]: every open document
// and (via each DocumentInfo.Active flag) which one is active.
type ListDocumentsResult struct {
	Documents []DocumentInfo `json:"documents"`
}

// CreateDocumentArgs is the request of [MethodDocumentsCreate]: the kind
// (part|assembly|drawing|presentation) and the document name.
type CreateDocumentArgs struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// ActivateDocumentArgs is the request of [MethodDocumentsActivate]: the session id
// of the document to make active.
type ActivateDocumentArgs struct {
	ID uint64 `json:"id"`
}
