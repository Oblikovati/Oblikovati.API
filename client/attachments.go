// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// Linked external-file attachments on a document (M03-F08).

// Attachments returns a document's attachment records.
//
// mcp:tool documents_list_attachments
// mcp:summary Returns a document's attachment records.
func (d Documents) Attachments(id uint64) (wire.ListAttachmentsResult, error) {
	var r wire.ListAttachmentsResult
	return r, d.c.call(wire.MethodDocumentsListAttachments, wire.ListAttachmentsArgs{Document: id}, &r)
}

// AddAttachment attaches an external file to a document under a unique name,
// returning the resolved record.
//
//	rec, err := c.Documents().AddAttachment(wire.AddAttachmentArgs{
//	    Document: doc.ID, Name: "loads", Kind: types.AttachmentLinked, FullFileName: "/w/loads.csv",
//	})
//
// mcp:tool documents_add_attachment
// mcp:summary Attaches an external file to a document under a unique name, returning the resolved record.
func (d Documents) AddAttachment(args wire.AddAttachmentArgs) (wire.AttachmentInfo, error) {
	var r wire.AttachmentInfo
	return r, d.c.call(wire.MethodDocumentsAddAttachment, args, &r)
}

// RemoveAttachment deletes a document's named attachment record (and an
// embedded payload with it).
//
// mcp:tool documents_remove_attachment
// mcp:summary Deletes a document's named attachment record (and an embedded payload with it).
func (d Documents) RemoveAttachment(id uint64, name string) (wire.OKResult, error) {
	var r wire.OKResult
	args := wire.RemoveAttachmentArgs{Document: id, Name: name}
	return r, d.c.call(wire.MethodDocumentsRemoveAttachment, args, &r)
}
