// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
)

func TestSketchConstraintSelectionsRoundTrip(t *testing.T) {
	original := DefaultSketchSettings()
	original.CustomInferenceTypes = true
	original.InferenceTypes = SketchInferHorizontal | SketchInferTangent
	original.CustomRelaxTypes = true
	original.RelaxTypes = SketchRelaxParallel | SketchRelaxFix
	encoded, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var restored SketchSettings
	if err := json.Unmarshal(encoded, &restored); err != nil {
		t.Fatal(err)
	}
	if restored != original {
		t.Fatalf("round trip = %+v, want %+v", restored, original)
	}
}

func TestSketchConstraintSelectionsRemainOptional(t *testing.T) {
	var legacy SketchSettings
	if err := json.Unmarshal([]byte(`{"inferConstraints":true}`), &legacy); err != nil {
		t.Fatal(err)
	}
	if legacy.CustomInferenceTypes || legacy.CustomRelaxTypes {
		t.Fatalf("legacy choices = %+v / %+v, want legacy selections", legacy.InferenceTypes, legacy.RelaxTypes)
	}
	var cleared SketchSettings
	if err := json.Unmarshal([]byte(`{"customInferenceTypes":true,"customRelaxTypes":true}`), &cleared); err != nil {
		t.Fatal(err)
	}
	if !cleared.CustomInferenceTypes || !cleared.CustomRelaxTypes || cleared.InferenceTypes != 0 || cleared.RelaxTypes != 0 {
		t.Fatalf("cleared choices = %+v / %+v, want explicit empty selections", cleared.InferenceTypes, cleared.RelaxTypes)
	}
}
