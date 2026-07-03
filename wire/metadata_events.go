// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// Metadata-mutation push events (Oblikovati/Oblikovati#1644): the mutation class add-ins previously
// could not observe. A rename or a property change (suppression, a sketch setting) is announced so
// an add-in mirroring the document — a browser extension, an exporter caching names, the MCP bridge
// showing state to an agent — stays consistent without polling. Like the other push events these
// are delivered to an add-in's Notify entry point (ADR-0016) with no request/response; an add-in
// matches on Type. They fire for BOTH UI-driven and add-in-driven mutations, since the host emits at
// its model/session seams (mirroring EmitFeatureLifecycle), not per router method.

// ObjectRenamedEvent is the JSON shape of the [EventObjectRenamed] push event: a document object's
// name changed. Kind is the object kind (body/sketch/feature/occurrence/document); Key is the
// object's stable reference key (empty for the document itself); OldName/NewName are the prior and
// new names.
type ObjectRenamedEvent struct {
	Type     string           `json:"type"` // always EventObjectRenamed
	Document uint64           `json:"document"`
	Kind     types.ObjectKind `json:"kind"`
	Key      string           `json:"key,omitempty"`
	OldName  string           `json:"oldName,omitempty"`
	NewName  string           `json:"newName"`
}

// PropertyChangedEvent is the JSON shape of the [EventPropertyChanged] push event: a document
// object's property changed (e.g. suppression toggled, a sketch setting edited). Kind and Key
// identify the object as in [ObjectRenamedEvent]; Property names the property (e.g. "suppressed");
// OldValue/NewValue are its prior and new values rendered as strings (a bool as "true"/"false"), so
// one generic event carries any property without a bespoke DTO per setting.
type PropertyChangedEvent struct {
	Type     string           `json:"type"` // always EventPropertyChanged
	Document uint64           `json:"document"`
	Kind     types.ObjectKind `json:"kind"`
	Key      string           `json:"key,omitempty"`
	Property string           `json:"property"`
	OldValue string           `json:"oldValue,omitempty"`
	NewValue string           `json:"newValue,omitempty"`
}
