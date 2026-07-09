// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"
)

// TestCreateWorkPointArgsPositionOmitsRefsKind guards the JSON shape of the default
// position request: it carries At and no empty kind/refs (omitempty), so the original
// position-only request is byte-for-byte unchanged.
func TestCreateWorkPointArgsPositionOmitsRefsKind(t *testing.T) {
	b, _ := json.Marshal(CreateWorkPointArgs{At: []float64{1, 2, 3}})
	got := string(b)
	if contains(got, "refs") || contains(got, "kind") {
		t.Errorf("position point leaked kind/refs: %s", got)
	}
	if !contains(got, `"at":[1,2,3]`) {
		t.Errorf("position point missing at: %s", got)
	}
}

// TestCreateWorkPointArgsIntersectionRoundTrip checks a plane-axis-intersection point
// survives a JSON round-trip with its kind and refs [plane, axis] intact and no at.
func TestCreateWorkPointArgsIntersectionRoundTrip(t *testing.T) {
	args := CreateWorkPointArgs{Kind: "plane-axis-intersection", Refs: []string{"origin/plane/xy", "axis/0"}}
	var back CreateWorkPointArgs
	b, _ := json.Marshal(args)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Kind != "plane-axis-intersection" || len(back.Refs) != 2 || back.Refs[0] != "origin/plane/xy" {
		t.Errorf("round-trip lost data: %+v", back)
	}
	if back.At != nil {
		t.Errorf("intersection point should carry no at: %+v", back)
	}
}

// TestCreateWorkPointResultHealthRoundTrip checks the result's new health fields survive
// a round-trip and Reason is omitted when healthy.
func TestCreateWorkPointResultHealthRoundTrip(t *testing.T) {
	b, _ := json.Marshal(CreateWorkPointResult{Index: 2, Ref: "point/2", Name: "Work Point1", Healthy: true})
	if contains(string(b), "reason") {
		t.Errorf("healthy result leaked reason: %s", b)
	}
	var back CreateWorkPointResult
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Ref != "point/2" || !back.Healthy {
		t.Errorf("result round-trip lost data: %+v", back)
	}
}
