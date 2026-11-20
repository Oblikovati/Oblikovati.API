// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestStartupSketchPlaneValid(t *testing.T) {
	for _, plane := range []StartupSketchPlane{StartupSketchNone, StartupSketchXY, StartupSketchYZ, StartupSketchXZ} {
		if !plane.IsValid() {
			t.Fatalf("plane %q: expected valid", plane)
		}
	}
	for _, plane := range []StartupSketchPlane{"", "XY", "bogus"} {
		if plane.IsValid() {
			t.Fatalf("plane %q: expected invalid", plane)
		}
	}
}
