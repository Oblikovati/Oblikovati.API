// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"fmt"

	"oblikovati.org/api/wire"
	"oblikovati.org/api/wire/featureargs"
)

// AddFeature creates a feature from a typed, compile-checked argument struct instead of a
// hand-assembled raw-JSON blob (the house rule "never reach the host with raw JSON"; ADR-0018,
// #1616). It marshals the struct, tags the envelope from its own [featureargs.Arg.Kind], and
// sends it through [Features.Add]; the result shape is operation-specific, so it is returned as
// raw JSON for the caller to decode (as with Add).
//
// A method cannot carry its own type parameter in Go, so this is a free function over the
// Features group rather than a method — one generic constructor keyed on the arg type, not one
// hand-written method per kind (which would re-create the duplication this removes):
//
//	_, err := client.AddFeature(c.Features(), featureargs.Extrude{SketchIndex: 0, Distance: "50 mm"})
func AddFeature[A featureargs.Arg](f Features, a A) (json.RawMessage, error) {
	raw, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("client: marshal %s feature args: %w", a.Kind(), err)
	}
	return f.Add(wire.AddFeatureArgs{Kind: a.Kind(), Args: raw})
}
