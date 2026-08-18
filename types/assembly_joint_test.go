// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestAssemblyJointTypeStringAndDOF(t *testing.T) {
	cases := []struct {
		kind AssemblyJointType
		name string
		dof  int
	}{
		{JointUnknown, "unknown", 0},
		{JointRigid, "rigid", 0},
		{JointRotational, "rotational", 1},
		{JointSlider, "slider", 1},
		{JointCylindrical, "cylindrical", 2},
		{JointPlanar, "planar", 3},
		{JointBall, "ball", 3},
	}
	for _, c := range cases {
		if got := c.kind.String(); got != c.name {
			t.Errorf("AssemblyJointType(%d).String() = %q, want %q", c.kind, got, c.name)
		}
		if got := c.kind.DegreesOfFreedom(); got != c.dof {
			t.Errorf("%v.DegreesOfFreedom() = %d, want %d", c.kind, got, c.dof)
		}
	}
}

func TestAssemblyJointTypeIsValid(t *testing.T) {
	if JointUnknown.IsValid() {
		t.Error("JointUnknown.IsValid() = true, want false")
	}
	for k := JointRigid; k <= JointBall; k++ {
		if !k.IsValid() {
			t.Errorf("%v.IsValid() = false, want true", k)
		}
	}
	if (JointBall + 1).IsValid() {
		t.Error("out-of-range joint type reported valid")
	}
}

func TestJointOriginAndDSStrings(t *testing.T) {
	origin := map[AssemblyJointOriginDefinitionType]string{
		JointOriginUnknown: "unknown", JointOriginPoint: "point",
		JointOriginEdge: "edge", JointOriginPlane: "plane",
	}
	for k, want := range origin {
		if got := k.String(); got != want {
			t.Errorf("origin(%d).String() = %q, want %q", k, got, want)
		}
	}
	ds := map[DSJointType]string{
		DSJointRigid: "rigid", DSJointRotational: "rotational", DSJointPrismatic: "prismatic",
		DSJointCylindrical: "cylindrical", DSJointPlanar: "planar", DSJointSpherical: "spherical",
	}
	for k, want := range ds {
		if got := k.String(); got != want {
			t.Errorf("DSJointType(%d).String() = %q, want %q", k, got, want)
		}
	}
	imposed := map[DSDOFImposedMotionType]string{DSDOFFree: "free", DSDOFDriven: "driven", DSDOFLocked: "locked"}
	for k, want := range imposed {
		if got := k.String(); got != want {
			t.Errorf("DSDOFImposedMotionType(%d).String() = %q, want %q", k, got, want)
		}
	}
}

// TestAssemblyJointOriginModeRoundTrip pins the origin-mode spellings, the empty/zero default (infer)
// and rejection of an unknown spelling (#1973).
func TestAssemblyJointOriginModeRoundTrip(t *testing.T) {
	cases := map[AssemblyJointOriginMode]string{
		JointOriginInfer: "infer", JointOriginOffset: "offset", JointOriginBetweenTwoFaces: "betweenTwoFaces",
	}
	for m, want := range cases {
		if got := m.String(); got != want {
			t.Errorf("AssemblyJointOriginMode(%d).String() = %q, want %q", m, got, want)
		}
		if parsed, ok := ParseAssemblyJointOriginMode(want); !ok || parsed != m {
			t.Errorf("ParseAssemblyJointOriginMode(%q) = (%d, %v), want (%d, true)", want, parsed, ok, m)
		}
	}
	if parsed, ok := ParseAssemblyJointOriginMode(""); !ok || parsed != JointOriginInfer {
		t.Errorf(`ParseAssemblyJointOriginMode("") = (%d, %v), want (JointOriginInfer, true)`, parsed, ok)
	}
	if _, ok := ParseAssemblyJointOriginMode("midpoint"); ok {
		t.Error("ParseAssemblyJointOriginMode(midpoint) = ok, want rejected")
	}
}

// TestOccurrenceDOFStateRoundTrip pins the DOF-state spellings, the empty/zero default and rejection
// of an unknown spelling (#1980).
func TestOccurrenceDOFStateRoundTrip(t *testing.T) {
	cases := map[OccurrenceDOFState]string{
		NoDegreeOfFreedomState: "none", CanRetainDegreeOfFreedom: "canRetain", CanIgnoreDegreeOfFreedom: "canIgnore",
	}
	for st, want := range cases {
		if got := st.String(); got != want {
			t.Errorf("OccurrenceDOFState(%d).String() = %q, want %q", st, got, want)
		}
		if parsed, ok := ParseOccurrenceDOFState(want); !ok || parsed != st {
			t.Errorf("ParseOccurrenceDOFState(%q) = (%d, %v), want (%d, true)", want, parsed, ok, st)
		}
	}
	if parsed, ok := ParseOccurrenceDOFState(""); !ok || parsed != NoDegreeOfFreedomState {
		t.Errorf(`ParseOccurrenceDOFState("") = (%d, %v), want (NoDegreeOfFreedomState, true)`, parsed, ok)
	}
	if _, ok := ParseOccurrenceDOFState("maybe"); ok {
		t.Error("ParseOccurrenceDOFState(maybe) = ok, want rejected")
	}
}
