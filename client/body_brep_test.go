// SPDX-License-Identifier: Apache-2.0

package client

import (
	"strings"
	"testing"

	"oblikovati.org/api/types"
	"oblikovati.org/api/wire"
)

// Round-trip smoke tests for the M07 body/brep method groups
// (Oblikovati/Oblikovati#293/#628/#629/#630): the request marshals the typed
// args onto the right wire method, and the reply decodes into the DTO.

func TestBodyShellsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"shells":[{"index":0,"closed":true,"faces":6,"edges":12,"volume":56,"key":"k0","transientKey":7},{"index":1,"closed":true,"void":true,"faces":6,"edges":12,"volume":8,"key":"k1","transientKey":9}]}`)}
	c := New(ft)
	r, err := c.Body().Shells(0)
	if err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodBodyShells {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodBodyShells)
	}
	if len(r.Shells) != 2 || !r.Shells[1].Void || r.Shells[1].Volume != 8 {
		t.Errorf("shells = %+v, want the void skin decoded", r.Shells)
	}
}

func TestBodyIsPointInsideParsesContainment(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"containment":"inside"}`)}
	c := New(ft)
	got, err := c.Body().IsPointInside(wire.IsPointInsideArgs{BodyIndex: 0, Point: []float64{1, 1, 1}})
	if err != nil {
		t.Fatal(err)
	}
	if got != types.InsideContainment {
		t.Errorf("containment = %v, want inside", got)
	}
	if !strings.Contains(string(ft.gotReq), `"point":[1,1,1]`) {
		t.Errorf("request %s should carry the point", ft.gotReq)
	}
}

func TestBodyConvexityEdgesSpellsCollection(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"edges":[{"key":"e1"}]}`)}
	c := New(ft)
	if _, err := c.Body().ConvexityEdges(0, types.AllConcaveEdges); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(ft.gotReq), `"collection":"allConcave"`) {
		t.Errorf("request %s should spell the collection", ft.gotReq)
	}
}

func TestBodyCalculateFacetsRoundTrip(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"vertexCount":3,"facetCount":1,"vertexCoordinates":[0,0,0,1,0,0,0,1,0],"vertexIndices":[0,1,2],"indexCountPerFace":[3]}`)}
	c := New(ft)
	r, err := c.Body().CalculateFacets(wire.CalculateFacetsArgs{BodyIndex: 0, Tolerance: 0.01})
	if err != nil {
		t.Fatal(err)
	}
	if r.FacetCount != 1 || len(r.VertexCoordinates) != 9 || r.IndexCountPerFace[0] != 3 {
		t.Errorf("facet set = %+v, want the triangle decoded", r)
	}
	if ft.gotMethod != wire.MethodBodyCalculateFacets {
		t.Errorf("method = %q", ft.gotMethod)
	}
}

func TestBrepCreatePrimitiveAndBoolean(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"handle":3,"stats":{"solid":true,"faces":6,"edges":12,"vertices":8,"shells":1,"volume":30}}`)}
	c := New(ft)
	r, err := c.TransientBRep().CreatePrimitive(wire.CreatePrimitiveArgs{
		Kind: "block", Min: []float64{0, 0, 0}, Max: []float64{3, 2, 5},
	})
	if err != nil {
		t.Fatal(err)
	}
	if r.Handle != 3 || r.Stats.Volume != 30 {
		t.Errorf("primitive = %+v", r)
	}
	if _, err := c.TransientBRep().DoBoolean(3, wire.BrepBodyRef{Handle: 4}, types.BooleanUnion); err != nil {
		t.Fatal(err)
	}
	if ft.gotMethod != wire.MethodBrepBoolean || !strings.Contains(string(ft.gotReq), `"operation":"union"`) {
		t.Errorf("boolean call = %q %s", ft.gotMethod, ft.gotReq)
	}
}

func TestBrepCreateFromDefinitionCarriesGraph(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"issues":[{"path":"edges[2]","problem":"vertex indices out of range"}]}`)}
	c := New(ft)
	def := types.BrepBodyDefinition{
		Solid:    true,
		Vertices: []types.BrepVertexDef{{Position: []float64{0, 0, 0}}},
	}
	r, err := c.TransientBRep().CreateFromDefinition(def)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Issues) != 1 || r.Issues[0].Path != "edges[2]" {
		t.Errorf("issues = %+v", r.Issues)
	}
	if !strings.Contains(string(ft.gotReq), `"vertices":[{"position":[0,0,0]}]`) {
		t.Errorf("request %s should carry the graph", ft.gotReq)
	}
}

func TestWireOffsetPlanarSpellsClosure(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"handle":5,"wires":[{"points":[0,0,0,1,0,0],"closed":false}]}`)}
	c := New(ft)
	r, err := c.Body().OffsetPlanarWire(0, 0, []float64{0, 0, 1}, 0.5, types.CircularCornerClosure)
	if err != nil {
		t.Fatal(err)
	}
	if r.Handle != 5 || len(r.Wires) != 1 {
		t.Errorf("offset = %+v", r)
	}
	if !strings.Contains(string(ft.gotReq), `"cornerClosure":"circular"`) {
		t.Errorf("request %s should spell the closure", ft.gotReq)
	}
}
