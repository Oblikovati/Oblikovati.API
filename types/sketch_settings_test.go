// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestDefaultSketchSettings pins the out-of-the-box configuration: the pre-#1877 inference toggles
// plus the Inventor-aligned grid/snap and constraint-display defaults (#1877).
func TestDefaultSketchSettings(t *testing.T) {
	s := DefaultSketchSettings()
	if !s.InferConstraints || !s.AutoApplyConstraints || s.ConstraintPriority != PriorityHorizontalVertical {
		t.Errorf("inference defaults = %+v, want inference/auto-apply on + horizontalVertical", s)
	}
	if s.XSnapSpacing != 0.1 || s.YSnapSpacing != 0.1 || s.SnapsPerMinorGrid != 1 || s.MinorLinesPerMajorGridLine != 10 {
		t.Errorf("grid defaults = %+v, want a 1 mm snap grid with 10 minor lines per major", s)
	}
	if !s.PersistInferredConstraints || !s.EditDimensionsWhenCreated || s.DisplayConstraintsOnCreation {
		t.Errorf("constraint-display defaults = %+v, want persist on / edit-dims on / display-on-create off", s)
	}
	if s.OverConstrainedBehavior != OverConstrainedApplyDriven {
		t.Errorf("default over-constrained behaviour = %v, want applyDriven", s.OverConstrainedBehavior)
	}
	if s.EnableRelaxMode || !s.KeepDimensionsWithEquationInRelaxMode {
		t.Errorf("relax defaults = %+v, want relax off / keep-equation-dims on", s)
	}
}

// TestSketchSettingsJSONRoundTrips asserts the wire shape survives marshal/unmarshal unchanged and
// carries the #1877 keys.
func TestSketchSettingsJSONRoundTrips(t *testing.T) {
	in := SketchSettings{
		InferConstraints: true, AutoApplyConstraints: false, ConstraintPriority: PriorityNone,
		XSnapSpacing: 0.25, YSnapSpacing: 0.5, SnapsPerMinorGrid: 2, MinorLinesPerMajorGridLine: 5,
		PersistInferredConstraints: true, DisplayConstraintsOnCreation: true, EditDimensionsWhenCreated: false,
		OverConstrainedBehavior: OverConstrainedPrompt,
		EnableRelaxMode:         true, KeepDimensionsWithEquationInRelaxMode: false,
	}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, key := range []string{"xSnapSpacing", "snapsPerMinorGrid", "persistInferredConstraints", "overConstrainedBehavior", "enableRelaxMode"} {
		if !strings.Contains(string(b), key) {
			t.Errorf("JSON %s is missing key %q", b, key)
		}
	}
	var out SketchSettings
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if out != in {
		t.Errorf("round-trip = %+v, want %+v", out, in)
	}
}
