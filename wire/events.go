// SPDX-License-Identifier: Apache-2.0

package wire

import "encoding/json"

// EditCommittedEvent is the JSON shape of an [EventEditCommitted] push event: a model
// mutation that was committed on a document, expressed as the very [wire] request that
// produced it. The host delivers it to an add-in's Notify entry point (ADR-0016); there
// is no request/response method and no client call — an add-in matches on Type.
//
// It is the basis for operational replication in a collaboration session
// (oblikovati-meeting ADR-0004): a remote add-in can replay the edit by dispatching
// Method with Args through its own client, since the payload is the same frozen contract.
//
// Method is a method-name constant from this package (e.g. [MethodParametersSet]); Args
// is that method's request DTO, left raw so the receiver decodes it with the matching
// type. Document is the affected document's id.
//
// NOTE (v1 scope): only edits that arrive through the host's method router are emitted.
// Edits performed directly in the host UI are not yet captured — a known gap tracked in
// oblikovati-meeting ADR-0004.
type EditCommittedEvent struct {
	Type     string          `json:"type"` // always EventEditCommitted
	Document uint64          `json:"document"`
	Method   string          `json:"method"`
	Args     json.RawMessage `json:"args,omitempty"`
}
