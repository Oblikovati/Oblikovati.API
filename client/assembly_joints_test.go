// SPDX-License-Identifier: Apache-2.0

package client

import (
	"encoding/json"
	"testing"

	"oblikovati.org/api/wire"
)

func TestAssemblyJointsListDecodes(t *testing.T) {
	ft := &fakeTransport{reply: []byte(
		`{"joints":[{"id":4,"type":"rotational","name":"Rotational:1","a":{"occurrence":1,"entity":"eA"},"b":{"occurrence":2,"entity":"eB"},"degreesOfFreedom":1}]}`)}
	c := New(ft)

	r, err := c.AssemblyJoints().List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyJointsList {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyJointsList)
	}
	if len(r.Joints) != 1 || r.Joints[0].Type != "rotational" || r.Joints[0].DegreesOfFreedom != 1 {
		t.Fatalf("decoded = %+v, want one rotational joint with 1 DOF", r.Joints)
	}
}

func TestAssemblyJointsAddRotationalMarshalsArgs(t *testing.T) {
	ft := &fakeTransport{reply: []byte(`{"joint":{"id":1,"type":"rotational","name":"Rotational:1","degreesOfFreedom":1}}`)}
	c := New(ft)

	args := wire.AddJointArgs{A: wire.ConstraintGeomRef{Occurrence: 1, Entity: "eA"}, B: wire.ConstraintGeomRef{Occurrence: 2, Entity: "eB"}, Flip: true}
	if _, err := c.AssemblyJoints().AddRotational(args); err != nil {
		t.Fatalf("AddRotational: %v", err)
	}
	if ft.gotMethod != wire.MethodAssemblyJointsAddRotational {
		t.Errorf("method = %q, want %q", ft.gotMethod, wire.MethodAssemblyJointsAddRotational)
	}
	var sent wire.AddJointArgs
	if err := json.Unmarshal(ft.gotReq, &sent); err != nil {
		t.Fatalf("request not valid JSON: %v", err)
	}
	if sent.A.Occurrence != 1 || sent.B.Entity != "eB" || !sent.Flip {
		t.Errorf("sent = %+v, want occ1/eB/flip", sent)
	}
}

// TestAssemblyJointAndDSMethodsRoute guards that every typed joint method targets its own
// wire constant.
func TestAssemblyJointAndDSMethodsRoute(t *testing.T) {
	for _, tc := range []struct {
		method string
		call   func(c *Client) error
	}{
		{wire.MethodAssemblyJointsAddRigid, func(c *Client) error { _, e := c.AssemblyJoints().AddRigid(wire.AddJointArgs{}); return e }},
		{wire.MethodAssemblyJointsAddSlider, func(c *Client) error { _, e := c.AssemblyJoints().AddSlider(wire.AddJointArgs{}); return e }},
		{wire.MethodAssemblyJointsAddCylindrical, func(c *Client) error { _, e := c.AssemblyJoints().AddCylindrical(wire.AddJointArgs{}); return e }},
		{wire.MethodAssemblyJointsAddPlanar, func(c *Client) error { _, e := c.AssemblyJoints().AddPlanar(wire.AddJointArgs{}); return e }},
		{wire.MethodAssemblyJointsAddBall, func(c *Client) error { _, e := c.AssemblyJoints().AddBall(wire.AddJointArgs{}); return e }},
		{wire.MethodAssemblyJointsDelete, func(c *Client) error { _, e := c.AssemblyJoints().Delete(1); return e }},
		{wire.MethodAssemblyJointsSetLimits, func(c *Client) error { _, e := c.AssemblyJoints().SetLimits(wire.SetJointLimitsArgs{ID: 1}); return e }},
		{wire.MethodAssemblyJointsSetFlip, func(c *Client) error { _, e := c.AssemblyJoints().SetFlip(1, true); return e }},
		{wire.MethodDSJointsList, func(c *Client) error { _, e := c.DSJoints().List(); return e }},
		{wire.MethodDSJointsAdd, func(c *Client) error { _, e := c.DSJoints().Add(wire.AddDSJointArgs{Type: "rotational"}); return e }},
		{wire.MethodDSJointsSetImposedMotion, func(c *Client) error {
			_, e := c.DSJoints().SetImposedMotion(wire.SetImposedMotionArgs{ID: 1})
			return e
		}},
		{wire.MethodDSJointsDelete, func(c *Client) error { _, e := c.DSJoints().Delete(1); return e }},
	} {
		ft := &fakeTransport{reply: []byte(`{}`)}
		if err := tc.call(New(ft)); err != nil {
			t.Fatalf("%s: %v", tc.method, err)
		}
		if ft.gotMethod != tc.method {
			t.Errorf("method = %q, want %q", ft.gotMethod, tc.method)
		}
	}
}
