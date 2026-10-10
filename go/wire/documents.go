// SPDX-License-Identifier: Apache-2.0

package wire

// DocumentInfo is the JSON shape of an open document.
type DocumentInfo struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	SubType string `json:"subType,omitempty"` // add-in flavored subtype id (M05-F15)
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
// SubType, when set, stamps the new document with a REGISTERED flavored subtype
// (documents.registerSubType) whose base type must match Type (M05-F15).
type CreateDocumentArgs struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	SubType string `json:"subType,omitempty"`
}

// ActivateDocumentArgs is the request of [MethodDocumentsActivate]: the session id
// of the document to make active.
type ActivateDocumentArgs struct {
	ID uint64 `json:"id"`
}

// CloseDocumentArgs is the request of [MethodDocumentsClose]: the session id of the
// document to close. Force discards unsaved changes instead of saving them first.
type CloseDocumentArgs struct {
	ID    uint64 `json:"id"`
	Force bool   `json:"force"`
}

// CloseAllDocumentsArgs is the request of [MethodDocumentsCloseAll]: close every
// open document. Force discards unsaved changes (the usual choice to start a clean
// session).
type CloseAllDocumentsArgs struct {
	Force bool `json:"force"`
}

// CloseDocumentsResult is the response of [MethodDocumentsClose] /
// [MethodDocumentsCloseAll]: how many documents were closed.
type CloseDocumentsResult struct {
	Closed int `json:"closed"`
}

// RegisterDocumentSubTypeArgs is the request of [MethodDocumentsRegisterSubType]:
// an add-in declaring a flavored document type over a base type ("part", …). The
// flavor's lifecycle reaches the add-in as client.operation push events
// (M05-F15, the DocumentSubTypeHandler + ClientOperationEvents equivalent).
type RegisterDocumentSubTypeArgs struct {
	ID          string `json:"id"`
	BaseType    string `json:"baseType"`
	DisplayName string `json:"displayName,omitempty"`
}

// DocumentSubTypeInfo is one entry of [MethodDocumentsListSubTypes].
type DocumentSubTypeInfo struct {
	ID          string `json:"id"`
	BaseType    string `json:"baseType"`
	DisplayName string `json:"displayName,omitempty"`
}

// ListDocumentSubTypesResult is the response of [MethodDocumentsListSubTypes].
type ListDocumentSubTypesResult struct {
	SubTypes []DocumentSubTypeInfo `json:"subTypes"`
}

// ClientOperationEvent is the push event (type [EventClientOperation]) telling a
// subtype's owner its flavored document needs servicing: Operation is "created",
// "opened", "activated" or "saved".
type ClientOperationEvent struct {
	Type      string `json:"type"` // always EventClientOperation
	Document  uint64 `json:"document"`
	SubType   string `json:"subType"`
	Operation string `json:"operation"`
}
