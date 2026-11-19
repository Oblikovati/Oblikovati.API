// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestConstraintGlyphVisibilityClientRoutesTypedRequests(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"ok":true}`)}
	group := New(transport).Sketch().Constrain(2)
	if _, err := group.SetGlyphVisible(4, false); err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodSketchSetConstraintGlyphVisibility {
		t.Fatalf("method %q", transport.gotMethod)
	}
	var one wire.ConstraintGlyphVisibilityArgs
	if err := json.Unmarshal(transport.gotReq, &one); err != nil {
		t.Fatal(err)
	}
	if one.SketchIndex != 2 || one.ConstraintID != 4 || one.Visible {
		t.Fatalf("request %+v", one)
	}
	if _, err := group.SetAllGlyphsVisible(true); err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodSketchSetAllConstraintGlyphsVisibility {
		t.Fatalf("method %q", transport.gotMethod)
	}
	var all wire.AllConstraintGlyphsVisibilityArgs
	if err := json.Unmarshal(transport.gotReq, &all); err != nil {
		t.Fatal(err)
	}
	if all.SketchIndex != 2 || !all.Visible {
		t.Fatalf("request %+v", all)
	}
}

func TestConstraintGlyphVisibilityClientReadsOverrides(t *testing.T) {
	transport := &fakeTransport{reply: []byte(`{"shown":true,"relationIds":[1,2,3],"hiddenIds":[1,3]}`)}
	view, err := New(transport).Sketch().Constrain(2).GlyphVisibility()
	if err != nil {
		t.Fatal(err)
	}
	if transport.gotMethod != wire.MethodSketchConstraintGlyphVisibility || !view.Shown || len(view.RelationIDs) != 3 || len(view.HiddenIDs) != 2 || view.HiddenIDs[1] != 3 {
		t.Fatalf("method %q, view %+v", transport.gotMethod, view)
	}
}
