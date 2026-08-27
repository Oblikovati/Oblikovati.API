// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The scalar read surface of the assembly constraint set (M12-F01,
// Oblikovati/Oblikovati#358/#363). An in-proc consumer reads a constraint's kind,
// driven value, limits, and health directly through these interfaces; every mutation
// (add/delete/set-limits/solve) travels over api/wire (the assemblyConstraints.*
// methods). The host implementations live in /source (model/assembly).

// AssemblyConstraint is the common read surface of every assembly relationship: its
// session id, kind, name, driven value, limits, health, and suppression.
type AssemblyConstraint interface {
	// ID is the constraint's session id, unique within its assembly's set.
	ID() uint64
	// Type is the relationship kind.
	Type() types.AssemblyConstraintType
	// Name is the constraint's display name (e.g. "Mate:1").
	Name() string
	// Value is the driven value: an offset (cm), angle (radians), or coupling ratio,
	// depending on the kind. Constraints with no driven value report 0.
	Value() float64
	// Health reports whether the constraint is fully evaluated, or sick (lost geometry).
	Health() types.HealthStatus
	// Suppressed reports whether the constraint is excluded from the solve.
	Suppressed() bool
	// Limits returns the constraint's driven-value bounds, or nil when unbounded.
	Limits() ConstraintLimits
}

// MateConstraint makes two geometries coincident with a directional solution (opposed
// faces, the default, or aligned). It is the read surface of a mate.
type MateConstraint interface {
	AssemblyConstraint
	// SolutionType is the directed sense the solver enforces on the two normals.
	SolutionType() types.MateConstraintSolutionType
}

// FlushConstraint makes two faces co-planar with aligned normals.
type FlushConstraint interface{ AssemblyConstraint }

// AngleConstraint holds an angle between two directions.
type AngleConstraint interface {
	AssemblyConstraint
	// SolutionType is how the angle is measured (undirected/directed/reference-vector).
	SolutionType() types.AngleConstraintSolutionType
}

// TangentConstraint keeps a face tangent to a curved face.
type TangentConstraint interface {
	AssemblyConstraint
	// Inside reports inside tangency (the curved face wraps the other); false is outside.
	Inside() bool
}

// InsertConstraint combines an axis mate with a plane mate (a bolt into a hole).
type InsertConstraint interface {
	AssemblyConstraint
	// Aligned reports the aligned plane sense; false is the default opposed sense.
	Aligned() bool
}

// AssemblySymmetryConstraint positions two geometries symmetrically about a plane.
type AssemblySymmetryConstraint interface{ AssemblyConstraint }

// RotateRotateConstraint couples two rotations by a gear ratio.
type RotateRotateConstraint interface {
	AssemblyConstraint
	// Ratio is revolutions of the second axis per revolution of the first.
	Ratio() float64
}

// RotateTranslateConstraint couples a rotation to a translation (rack and pinion).
type RotateTranslateConstraint interface {
	AssemblyConstraint
	// Distance is the translation moved per revolution (cm).
	Distance() float64
}

// TranslateTranslateConstraint couples two translations by a ratio.
type TranslateTranslateConstraint interface {
	AssemblyConstraint
	// Ratio is the distance of the second axis per unit distance of the first.
	Ratio() float64
}

// TransitionalConstraint keeps a face in sliding contact with a transition face.
type TransitionalConstraint interface{ AssemblyConstraint }

// CustomConstraint is a relationship solved by an add-in, not the built-in solver.
type CustomConstraint interface {
	AssemblyConstraint
	// Kind names the add-in relationship the host dispatches to.
	Kind() string
	// Params are the driving values passed to the add-in solver.
	Params() []float64
}

// ConstraintLimits bounds a constraint's driven value. Each bound is independently
// optional; an absent bound does not clamp.
type ConstraintLimits interface {
	// Minimum returns the lower bound and whether it is set.
	Minimum() (float64, bool)
	// Maximum returns the upper bound and whether it is set.
	Maximum() (float64, bool)
	// Resting returns the value a drive returns to when released and whether it is set.
	Resting() (float64, bool)
}

// AssemblyConstraints is the assembly's constraint collection — the relationships
// authored directly in it, in creation order (host: assembly.ConstraintSet).
type AssemblyConstraints = Enumerable[AssemblyConstraint]

// AssemblyConstraintsEnumerator is the per-occurrence view of the constraints that
// reference one occurrence (the reference API's per-component Constraints collection).
type AssemblyConstraintsEnumerator = Enumerable[AssemblyConstraint]
