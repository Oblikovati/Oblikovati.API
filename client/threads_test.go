// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestThreadsTableQueryMarshalsFilters(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"threadTypes":["ISO Metric profile"],"nominalSizes":["M8"]}`)}
	c := New(ft)

	got, err := c.Threads().TableQuery(wire.ThreadTableQueryArgs{ThreadType: "ISO Metric profile"})
	if err != nil {
		t.Fatalf("TableQuery: %v", err)
	}
	if ft.gotMethod != wire.MethodThreadsTableQuery {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodThreadsTableQuery)
	}
	var sent wire.ThreadTableQueryArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.ThreadType != "ISO Metric profile" {
		t.Errorf("sent threadType = %q, want ISO Metric profile", sent.ThreadType)
	}
	if len(got.NominalSizes) != 1 || got.NominalSizes[0] != "M8" {
		t.Errorf("decoded sizes = %v, want [M8]", got.NominalSizes)
	}
}

func TestThreadsResolveCarriesClassAndTapered(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"designation":"M8x1.25","threadType":"ISO Metric profile","nominalSize":"M8","metric":true,"internal":true,"rightHanded":true,"tapered":false,"pitch":1.25,"majorDiameter":8,"minorDiameter":6.647,"pitchDiameter":7.188,"tapDrillDiameter":6.8}`)}
	c := New(ft)

	got, err := c.Threads().Resolve(wire.ResolveThreadArgs{Designation: "M8x1.25", Class: "6H", Internal: true})
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if ft.gotMethod != wire.MethodThreadsResolve {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodThreadsResolve)
	}
	var sent wire.ResolveThreadArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.Class != "6H" || !sent.Internal {
		t.Errorf("sent = %+v, want class 6H internal", sent)
	}
	if got.Pitch != 1.25 || !got.Metric {
		t.Errorf("decoded = %+v, want metric pitch 1.25", got)
	}
}
