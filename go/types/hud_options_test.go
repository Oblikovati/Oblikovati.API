// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
)

// TestDefaultHeadsUpDisplayOptions pins the out-of-the-box configuration: a shape's first point
// is entered in Cartesian coordinates while the shape itself is sized in polar length and angle,
// and a typed value becomes a persistent dimension (Oblikovati/Oblikovati#2014).
func TestDefaultHeadsUpDisplayOptions(t *testing.T) {
	o := DefaultHeadsUpDisplayOptions()
	if !o.Enabled || !o.PointerInputEnabled || !o.DimensionInputEnabled {
		t.Errorf("defaults = %+v, want the display and both input surfaces on", o)
	}
	if !o.PointerInputInCartesianCoordinates {
		t.Error("a shape's first point defaults to Cartesian X/Y entry")
	}
	if o.DimensionInputInCartesianCoordinates {
		t.Error("the shape being placed defaults to polar length/angle entry")
	}
	if !o.CreateDimensionsOnValueInput {
		t.Error("a typed value defaults to becoming a persistent dimension")
	}
}

func TestHeadsUpDisplayOptionsJSONRoundTrips(t *testing.T) {
	in := HeadsUpDisplayOptions{
		Enabled: true, PointerInputEnabled: false, PointerInputInCartesianCoordinates: true,
		DimensionInputEnabled: true, DimensionInputInCartesianCoordinates: true,
		CreateDimensionsOnValueInput: false,
	}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var out HeadsUpDisplayOptions
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out != in {
		t.Errorf("round trip = %+v, want %+v", out, in)
	}
}
