// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestAssemblyConstraintTypeString(t *testing.T) {
	cases := map[AssemblyConstraintType]string{
		ConstraintUnknown:            "unknown",
		ConstraintMate:               "mate",
		ConstraintFlush:              "flush",
		ConstraintAngle:              "angle",
		ConstraintTangent:            "tangent",
		ConstraintInsert:             "insert",
		ConstraintSymmetry:           "symmetry",
		ConstraintRotateRotate:       "rotate-rotate",
		ConstraintRotateTranslate:    "rotate-translate",
		ConstraintTranslateTranslate: "translate-translate",
		ConstraintTransitional:       "transitional",
		ConstraintCustom:             "custom",
	}
	for kind, want := range cases {
		if got := kind.String(); got != want {
			t.Errorf("AssemblyConstraintType(%d).String() = %q, want %q", kind, got, want)
		}
	}
}

func TestAssemblyConstraintTypeIsValid(t *testing.T) {
	if ConstraintUnknown.IsValid() {
		t.Error("ConstraintUnknown.IsValid() = true, want false")
	}
	for kind := ConstraintMate; kind <= ConstraintCustom; kind++ {
		if !kind.IsValid() {
			t.Errorf("%v.IsValid() = false, want true", kind)
		}
	}
	if (ConstraintCustom + 1).IsValid() {
		t.Error("out-of-range constraint type reported valid")
	}
}

func TestMateConstraintSolutionTypeString(t *testing.T) {
	cases := map[MateConstraintSolutionType]string{
		MateSolutionOpposed:    "opposed",
		MateSolutionAligned:    "aligned",
		MateSolutionUndirected: "undirected",
		MateSolutionNoSolution: "noSolution",
	}
	for sol, want := range cases {
		if got := sol.String(); got != want {
			t.Errorf("MateConstraintSolutionType(%d).String() = %q, want %q", sol, got, want)
		}
		if got, ok := ParseMateConstraintSolutionType(want); !ok || got != sol {
			t.Errorf("ParseMateConstraintSolutionType(%q) = (%d, %v), want (%d, true)", want, got, ok, sol)
		}
	}
	if got, ok := ParseMateConstraintSolutionType(""); !ok || got != MateSolutionOpposed {
		t.Errorf(`ParseMateConstraintSolutionType("") = (%d, %v), want (opposed, true)`, got, ok)
	}
	if _, ok := ParseMateConstraintSolutionType("sideways"); ok {
		t.Error("an unknown mate solution should not resolve")
	}
}

func TestAngleConstraintSolutionTypeString(t *testing.T) {
	cases := map[AngleConstraintSolutionType]string{
		AngleSolutionUndirected:      "undirected",
		AngleSolutionDirected:        "directed",
		AngleSolutionReferenceVector: "reference-vector",
	}
	for sol, want := range cases {
		if got := sol.String(); got != want {
			t.Errorf("AngleConstraintSolutionType(%d).String() = %q, want %q", sol, got, want)
		}
	}
}

// TestHealthStatusStringMatchesHostVocabulary guards that the public health names
// stay byte-identical to the host's model/health.Status.String() values, so the
// wire's health strings agree across the contract boundary.
func TestHealthStatusStringMatchesHostVocabulary(t *testing.T) {
	cases := map[HealthStatus]string{
		HealthOK:         "ok",
		HealthWarning:    "warning",
		HealthSick:       "sick",
		HealthSuppressed: "suppressed",
	}
	for status, want := range cases {
		if got := status.String(); got != want {
			t.Errorf("HealthStatus(%d).String() = %q, want %q", status, got, want)
		}
	}
}
