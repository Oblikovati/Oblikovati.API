// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

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

// Close closes the document with the given session id; force discards unsaved
// changes instead of saving them first.
//
//	closed, err := c.Documents().Close(doc.ID, false)
func (d Documents) Close(id uint64, force bool) (wire.CloseDocumentsResult, error) {
	var r wire.CloseDocumentsResult
	return r, d.c.call(wire.MethodDocumentsClose, wire.CloseDocumentArgs{ID: id, Force: force}, &r)
}

// CloseAll closes every open document; force discards unsaved changes (the usual
// choice to reset to a clean session).
//
//	closed, err := c.Documents().CloseAll(true)
func (d Documents) CloseAll(force bool) (wire.CloseDocumentsResult, error) {
	var r wire.CloseDocumentsResult
	return r, d.c.call(wire.MethodDocumentsCloseAll, wire.CloseAllDocumentsArgs{Force: force}, &r)
}

// RegisterSubType declares a flavored document subtype over a base type; the
// flavor's lifecycle reaches the owner as client.operation push events (M05-F15).
//
//	client.Documents().RegisterSubType(wire.RegisterDocumentSubTypeArgs{
//	    ID: "com.x.sim.study", BaseType: "part", DisplayName: "Simulation Study",
//	})
func (d Documents) RegisterSubType(args wire.RegisterDocumentSubTypeArgs) (wire.OKResult, error) {
	var r wire.OKResult
	return r, d.c.call(wire.MethodDocumentsRegisterSubType, args, &r)
}

// SubTypes returns the registered flavored subtypes.
func (d Documents) SubTypes() (wire.ListDocumentSubTypesResult, error) {
	var r wire.ListDocumentSubTypesResult
	return r, d.c.call(wire.MethodDocumentsListSubTypes, nil, &r)
}
