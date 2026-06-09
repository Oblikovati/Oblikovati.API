// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"

	"oblikovati.org/api/wire"
)

// Features is the feature-creation operation group for the active part.
type Features struct{ c *Client }

// Features returns the feature operation group.
func (c *Client) Features() Features { return Features{c} }

// List returns every feature operation Add can create, each with its args schema.
func (f Features) List() (wire.ListFeatureKindsResult, error) {
	var r wire.ListFeatureKindsResult
	return r, f.c.call(wire.MethodFeaturesList, nil, &r)
}

// Add applies a feature operation. The result shape is operation-specific (see the
// kind's schema from List), so it is returned as raw JSON for the caller to decode.
func (f Features) Add(args wire.AddFeatureArgs) (json.RawMessage, error) {
	var r json.RawMessage
	return r, f.c.call(wire.MethodFeaturesAdd, args, &r)
}
