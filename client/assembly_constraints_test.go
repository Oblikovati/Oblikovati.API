// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestAssemblyConstraintsListDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"constraints":[{"id":7,"type":"mate","name":"Mate:1","a":{"occurrence":1,"entity":"fA"},"b":{"occurrence":2,"entity":"fB"},"health":"ok"}]}`)}
	c := New(ft)

	r, err := c.AssemblyConstraints().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyConstraintsList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyConstraintsList)
	}
	if len(r.Constraints) != 1 || r.Constraints[0].ID != 7 || r.Constraints[0].Type != "mate" {
		t.Fatalf("decoded = %+v, want one mate constraint id 7", r.Constraints)
	}
	if r.Constraints[0].A.Occurrence != 1 || r.Constraints[0].B.Entity != "fB" {
		t.Errorf("geometry refs = %+v, want a.occ=1 b.entity=fB", r.Constraints[0])
	}
}

func TestAssemblyConstraintsAddMateMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"constraint":{"id":1,"type":"mate","name":"Mate:1"}}`)}
	c := New(ft)

	args := wire.AddMateArgs{
		A:        wire.ConstraintGeomRef{Occurrence: 1, Entity: "fA"},
		B:        wire.ConstraintGeomRef{Occurrence: 2, Entity: "fB"},
		Offset:   0.5,
		Solution: "aligned",
	}
	if _, err := c.AssemblyConstraints().AddMate(args); err != nil {
		t.Fatalf("AddMate: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyConstraintsAddMate {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyConstraintsAddMate)
	}
	var sent wire.AddMateArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.A.Occurrence != 1 || sent.B.Entity != "fB" || sent.Offset != 0.5 || sent.Solution != "aligned" {
		t.Errorf("sent = %+v, want occ1/fB/offset0.5/aligned", sent)
	}
}

func TestAssemblyConstraintsSolveDecodesHealth(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"status":"under-constrained","constraints":2,"redundant":0,"degreesOfFreedom":3,"converged":true,"occurrences":[{"occurrence":2,"degreesOfFreedom":3}]}`)}
	c := New(ft)

	r, err := c.AssemblyConstraints().Solve()
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyConstraintsSolve {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyConstraintsSolve)
	}
	if r.DegreesOfFreedom != 3 || !r.Converged || len(r.Occurrences) != 1 || r.Occurrences[0].Occurrence != 2 {
		t.Fatalf("decoded = %+v, want 3 DOF converged with one occurrence row", r)
	}
}

// TestAssemblyConstraintsAddRouteToTheirMethods guards that every typed add/mutate method
// targets its own wire constant — a copy-paste of the wrong method name would corrupt the
// whole surface.
func TestAssemblyConstraintsAddRouteToTheirMethods(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(g AssemblyConstraints) error
	}{
		{wire.MethodAssemblyConstraintsAddFlush, func(g AssemblyConstraints) error { _, e := g.AddFlush(wire.AddFlushArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddAngle, func(g AssemblyConstraints) error { _, e := g.AddAngle(wire.AddAngleArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddTangent, func(g AssemblyConstraints) error { _, e := g.AddTangent(wire.AddTangentArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddInsert, func(g AssemblyConstraints) error { _, e := g.AddInsert(wire.AddInsertArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddSymmetry, func(g AssemblyConstraints) error { _, e := g.AddSymmetry(wire.AddSymmetryArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddRotateRotate, func(g AssemblyConstraints) error { _, e := g.AddRotateRotate(wire.AddRotateRotateArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddRotateTranslate, func(g AssemblyConstraints) error {
			_, e := g.AddRotateTranslate(wire.AddRotateTranslateArgs{})
			return e
		}},
		{wire.MethodAssemblyConstraintsAddTranslateTranslate, func(g AssemblyConstraints) error {
			_, e := g.AddTranslateTranslate(wire.AddTranslateTranslateArgs{})
			return e
		}},
		{wire.MethodAssemblyConstraintsAddTransitional, func(g AssemblyConstraints) error { _, e := g.AddTransitional(wire.AddTransitionalArgs{}); return e }},
		{wire.MethodAssemblyConstraintsAddCustom, func(g AssemblyConstraints) error { _, e := g.AddCustom(wire.AddCustomArgs{}); return e }},
		{wire.MethodAssemblyConstraintsDelete, func(g AssemblyConstraints) error { _, e := g.Delete(1); return e }},
		{wire.MethodAssemblyConstraintsSetLimits, func(g AssemblyConstraints) error { _, e := g.SetLimits(wire.SetConstraintLimitsArgs{ID: 1}); return e }},
		{wire.MethodAssemblyConstraintsHealth, func(g AssemblyConstraints) error { _, e := g.Health(); return e }},
	} {
		ft := &fakeTransport{reply: []byte(`{}`)}
		if err := tc.call(New(ft).AssemblyConstraints()); err != nil {
			t.Fatalf("%s: %v", tc.method, err)
		}
		if ft.gotMethod != tc.method {
			t.Errorf("method = %q, want %q", ft.gotMethod, tc.method)
		}
	}
}
