// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The transient B-rep factory contract (M07-F05, Oblikovati/Oblikovati#628):
// creating and manipulating ownerless (non-document) bodies. In-proc form;
// add-ins out of process use the brep.* wire methods via api/client.

// TransientBody is one ownerless body the factory manages.
type TransientBody interface {
	IsSolid() bool
	FaceCount() int
	EdgeCount() int
	VertexCount() int
	ShellCount() int
	WireCount() int
	// Volume is the enclosed volume (cm³); 0 for an open surface body.
	Volume() float64
}

// TransientBRep creates and combines transient bodies. Transform matrices
// are restricted to rotation/translation/reflection/uniform scale — scale or
// shear that would break an analytic surface is rejected with the offending
// matrix in the error.
type TransientBRep interface {
	CreateSolidBlock(min, max types.Point) (TransientBody, error)
	// CreateSolidCylinderCone: equal radii = cylinder, one zero = full cone.
	CreateSolidCylinderCone(bottom, top types.Point, bottomRadius, topRadius float64) (TransientBody, error)
	CreateSolidSphere(center types.Point, radius float64) (TransientBody, error)
	CreateSolidTorus(center types.Point, axis types.Vector, majorRadius, minorRadius float64) (TransientBody, error)
	Copy(body TransientBody) (TransientBody, error)
	Transform(body TransientBody, m types.Matrix) error
	// DoBoolean modifies blank in place.
	DoBoolean(blank, tool TransientBody, op types.BooleanType) error
	// CreateIntersectionWithPlane returns the section curves as wires on a
	// new body.
	CreateIntersectionWithPlane(body TransientBody, planeOrigin types.Point, planeNormal types.Vector) (TransientBody, error)
	// CreateFromDefinition compiles the bottom-up graph; a non-empty issue
	// list means the graph was rejected (no body).
	CreateFromDefinition(def types.BrepBodyDefinition) (TransientBody, []types.BrepDefinitionIssue, error)
}
