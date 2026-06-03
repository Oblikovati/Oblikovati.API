// SPDX-License-Identifier: Apache-2.0

// Package contract's sketch interfaces are the in-process Go view of a 2D sketch the
// host implements (model/sketch.Sketch) and add-ins consume. They expose only the
// scalar surface — no GPL model types leak across the boundary; richer access (entity
// geometry, constraints) travels as wire DTOs (see api/wire).
package contract

// Sketch is the scalar view of a 2D planar sketch: its identity, visibility, and the
// solver's degree-of-freedom count. The host's model/sketch.Sketch satisfies this via
// a compile-time assertion.
type Sketch interface {
	// Name is the sketch's display name.
	Name() string
	// Visible reports whether the sketch is shown.
	Visible() bool
	// EntityCount is the number of geometry entities the sketch owns.
	EntityCount() int
	// DegreesOfFreedom is the sketch's remaining free DOF (0 when fully constrained).
	DegreesOfFreedom() int
}

// Profile is the scalar view of a sketch region a feature consumes: its enclosed area and
// whether it is closed (extrudable into a solid). The host's model/sketch.Profile
// satisfies this.
type Profile interface {
	// Area is the profile's enclosed area in sketch-plane cm² (holes subtracted).
	Area() float64
	// IsClosed reports whether the profile encloses a region.
	IsClosed() bool
}
