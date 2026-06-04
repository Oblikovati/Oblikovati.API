// SPDX-License-Identifier: Apache-2.0

package contract

// Sketch3D is the scalar view of a 3D (non-planar) sketch the host implements
// (model/sketch.Sketch3D) and add-ins consume: its identity, visibility, entity count,
// and the solver's remaining degree-of-freedom count. Like [Sketch] it exposes only the
// scalar surface — entity geometry and constraints travel as wire DTOs (see api/wire).
// The host satisfies this via a compile-time assertion.
type Sketch3D interface {
	// Name is the sketch's display name.
	Name() string
	// Visible reports whether the sketch is shown.
	Visible() bool
	// EntityCount is the number of geometry entities the sketch owns.
	EntityCount() int
	// DegreesOfFreedom is the sketch's remaining free DOF (0 when fully constrained).
	DegreesOfFreedom() int
}

// Profile3D is the scalar view of a closed, planar loop of a 3D sketch a planar-section
// feature consumes: its enclosed area and that it is closed. The host's
// model/sketch.Profile3D satisfies this.
type Profile3D interface {
	// Area is the loop's enclosed area in model cm².
	Area() float64
	// IsClosed reports whether the loop encloses a region (always true for a Profile3D).
	IsClosed() bool
}
