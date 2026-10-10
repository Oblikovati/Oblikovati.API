// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"fmt"
)

// marshalFloats renders the values as a fixed JSON array — the geometry wire
// encoding ([x,y,z] points/vectors, 9/16-cell matrices).
func marshalFloats(values ...float64) ([]byte, error) {
	return json.Marshal(values)
}

// unmarshalFloats decodes a fixed-length JSON array into the given fields,
// naming the type and the offending payload on a length mismatch.
func unmarshalFloats(b []byte, typeName string, fields ...*float64) error {
	var values []float64
	if err := json.Unmarshal(b, &values); err != nil {
		return fmt.Errorf("types: %s expects a JSON number array, got %s: %w", typeName, string(b), err)
	}
	if len(values) != len(fields) {
		return fmt.Errorf("types: %s expects %d numbers, got %d in %s", typeName, len(fields), len(values), string(b))
	}
	for i, f := range fields {
		*f = values[i]
	}
	return nil
}
