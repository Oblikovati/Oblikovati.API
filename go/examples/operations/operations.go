// SPDX-License-Identifier: Apache-2.0

// Package operations demonstrates typed requests through an injected host.
package operations

import (
	"encoding/json"
	"oblikovati.org/api/client"
	"oblikovati.org/api/wire"
	"oblikovati.org/api/wire/featureargs"
)

// CreatePart opens a part document through the host. Example: CreatePart(host, "Bracket").
func CreatePart(host client.Caller, name string) (wire.DocumentInfo, error) {
	return client.New(host).Documents().Create(wire.CreateDocumentArgs{Type: "part", Name: name})
}

// ExtrudeProfile extrudes a closed, fully constrained sketch profile in the active part.
// Example: ExtrudeProfile(host, 0, "50 mm"). Create the sketch before this call.
func ExtrudeProfile(host client.Caller, sketchIndex int, distance string) (json.RawMessage, error) {
	args := featureargs.Extrude{SketchIndex: sketchIndex, Distance: distance, Operation: "new"}
	return client.AddFeature(client.New(host).Features(), args)
}
