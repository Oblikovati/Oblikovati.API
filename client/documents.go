// SPDX-License-Identifier: Apache-2.0

package client

import "github.com/Oblikovati/api/wire"

// Documents is the document-management operation group.
type Documents struct{ c *Client }

// Documents returns the document operation group.
func (c *Client) Documents() Documents { return Documents{c} }

// List returns every open document and which one is active.
func (d Documents) List() (wire.ListDocumentsResult, error) {
	var r wire.ListDocumentsResult
	return r, d.c.call(wire.MethodDocumentsList, nil, &r)
}

// Create makes a new document of the given kind active and returns it.
func (d Documents) Create(args wire.CreateDocumentArgs) (wire.DocumentInfo, error) {
	var r wire.DocumentInfo
	return r, d.c.call(wire.MethodDocumentsCreate, args, &r)
}

// Activate makes the document with the given session id active.
func (d Documents) Activate(id uint64) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDocumentsActivate, wire.ActivateDocumentArgs{ID: id}, &r)
}
