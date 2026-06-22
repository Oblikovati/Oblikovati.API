// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestEnvironmentString(t *testing.T) {
	cases := map[Environment]string{
		BaseEnvironment:     "base",
		SketchEnvironment:   "sketch",
		Sketch3DEnvironment: "sketch3d",
		Environment(99):     "environment(?)",
	}
	for env, want := range cases {
		if got := env.String(); got != want {
			t.Errorf("Environment(%d).String() = %q, want %q", env, got, want)
		}
	}
}

// TestEnvironmentValuesAreStable guards the wire-stable numeric values, which the host and
// add-ins both depend on (a control's environment crosses the C ABI as this number).
func TestEnvironmentValuesAreStable(t *testing.T) {
	if BaseEnvironment != 0 || SketchEnvironment != 1 || Sketch3DEnvironment != 2 {
		t.Fatalf("environment values drifted: base=%d sketch=%d sketch3d=%d",
			BaseEnvironment, SketchEnvironment, Sketch3DEnvironment)
	}
}
