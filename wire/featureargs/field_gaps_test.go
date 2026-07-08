// SPDX-License-Identifier: Apache-2.0

package featureargs

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestHoleCenterRoundTrip checks the multi-hole placement fields survive a JSON round-trip
// and are omitted when nil, so a plate with several holes on one face stays expressible.
func TestHoleCenterRoundTrip(t *testing.T) {
	h := Hole{FaceRef: "face/0", Diameter: "5 mm", Center: []float64{1, 2, 0}, CenterExpr: []string{"d0", "d1", "0"}}
	var back Hole
	b, _ := json.Marshal(h)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(h, back) {
		t.Errorf("hole center did not round-trip: sent %#v, got %#v (json %s)", h, back, b)
	}

	bare, _ := json.Marshal(Hole{FaceRef: "face/0", Diameter: "5 mm"})
	if got := string(bare); contains(got, "center") || contains(got, "centerExpr") {
		t.Errorf("nil hole center leaked into json: %s", got)
	}
}

// TestDraftPullDirectionRoundTrip checks the explicit pull direction survives a JSON
// round-trip and is omitted when nil (host infers from the neutral faces).
func TestDraftPullDirectionRoundTrip(t *testing.T) {
	d := Draft{FaceRefs: []string{"face/1"}, Angle: "3 deg", PullDirection: []float64{0, 0, 1}}
	var back Draft
	b, _ := json.Marshal(d)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(d, back) {
		t.Errorf("draft pull direction did not round-trip: sent %#v, got %#v (json %s)", d, back, b)
	}

	bare, _ := json.Marshal(Draft{FaceRefs: []string{"face/1"}, Angle: "3 deg"})
	if got := string(bare); contains(got, "pullDirection") {
		t.Errorf("nil pull direction leaked into json: %s", got)
	}
}

// TestSweepPathPointsRoundTrip checks the explicit polyline path survives a JSON round-trip
// and is omitted when nil (fall back to the sketch-path arg).
func TestSweepPathPointsRoundTrip(t *testing.T) {
	s := Sweep{SketchIndex: 0, ProfileIndex: 0, PathPoints: [][]float64{{0, 0, 0}, {1, 0, 0}, {1, 1, 0}}}
	var back Sweep
	b, _ := json.Marshal(s)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, back) {
		t.Errorf("sweep path points did not round-trip: sent %#v, got %#v (json %s)", s, back, b)
	}

	bare, _ := json.Marshal(Sweep{SketchIndex: 0, ProfileIndex: 0})
	if got := string(bare); contains(got, "pathPoints") {
		t.Errorf("nil path points leaked into json: %s", got)
	}
}

// TestThreadSpanRoundTrip checks the partial-length fields (Inventor's ThreadDepth/ThreadOffset)
// survive a JSON round-trip and are omitted when empty, so a full-length cosmetic thread stays
// the terse default while a double-ended stud can thread only its ends.
func TestThreadSpanRoundTrip(t *testing.T) {
	th := Thread{FaceRef: "face/0", Designation: "M12x1.75", Offset: "length - nut_thread_length", Length: "nut_thread_length"}
	var back Thread
	b, _ := json.Marshal(th)
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(th, back) {
		t.Errorf("thread span did not round-trip: sent %#v, got %#v (json %s)", th, back, b)
	}

	bare, _ := json.Marshal(Thread{FaceRef: "face/0", Designation: "M12x1.75"})
	if got := string(bare); contains(got, "length") || contains(got, "offset") {
		t.Errorf("empty thread span leaked into json: %s", got)
	}
}

// contains reports whether s holds sub — a tiny local helper so these tests avoid a
// strings import, matching the package's terse test style.
func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
