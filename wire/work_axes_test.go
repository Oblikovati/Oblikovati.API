// SPDX-License-Identifier: Apache-2.0

package wire

import (
	"encoding/json"
	"testing"
)

// TestCreateWorkAxisArgsLineOmitsRefs guards the JSON shape: a grounded "line" axis
// carries origin/direction and no empty refs field (omitempty), so the discriminator
// stays clean for the host router.
func TestCreateWorkAxisArgsLineOmitsRefs(t *testing.T) {
	b, _ := json.Marshal(CreateWorkAxisArgs{Kind: "line", Origin: []float64{0, 0, 0}, Direction: []float64{0, 0, 1}})
	got := string(b)
	if contains(got, "refs") {
		t.Errorf("line axis leaked refs: %s", got)
	}
	for _, want := range []string{"origin", "direction", `"kind":"line"`} {
		if !contains(got, want) {
			t.Errorf("line axis missing %q: %s", want, got)
		}
	}
}

// TestCreateWorkAxisArgsRefsRoundTrip checks a reference-model axis (two-points /
// plane-intersection) survives a JSON round-trip with its kind and refs intact.
func TestCreateWorkAxisArgsRefsRoundTrip(t *testing.T) {
	args := CreateWorkAxisArgs{Kind: "two-points", Refs: []string{"point/0", "point/1"}}
	var back CreateWorkAxisArgs
	b, _ := json.Marshal(args)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.Kind != "two-points" || len(back.Refs) != 2 || back.Refs[1] != "point/1" {
		t.Errorf("round-trip lost data: %+v", back)
	}
	if back.Origin != nil || back.Direction != nil {
		t.Errorf("refs-kind axis should carry no origin/direction: %+v", back)
	}
}

// TestWorkAxisInfoRoundTrip checks a list row survives a JSON round-trip with its
// geometry and health intact.
func TestWorkAxisInfoRoundTrip(t *testing.T) {
	res := ListWorkAxesResult{Axes: []WorkAxisInfo{{
		Index: 3, Name: "Work Axis1", Ref: "axis/3", Kind: "plane-intersection",
		Origin: []float64{1, 2, 3}, Direction: []float64{0, 1, 0}, Healthy: true,
	}}}
	var back ListWorkAxesResult
	b, _ := json.Marshal(res)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Axes) != 1 {
		t.Fatalf("axes lost: %+v", back)
	}
	ax := back.Axes[0]
	if ax.Ref != "axis/3" || ax.Kind != "plane-intersection" || ax.IsOrigin || !ax.Healthy {
		t.Errorf("axis info lost: %+v", ax)
	}
	if len(ax.Origin) != 3 || ax.Direction[1] != 1 {
		t.Errorf("axis geometry lost: %+v", ax)
	}
}
