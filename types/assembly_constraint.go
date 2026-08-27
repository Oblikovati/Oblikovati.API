// SPDX-License-Identifier: Apache-2.0

package types

// The assembly-constraint vocabulary (M12-F01, Oblikovati/Oblikovati#358/#363):
// the kinds of relationship that position one occurrence relative to another, plus
// the solution-type discriminators a few of them carry. These are the canonical,
// Apache-2.0 definitions; the GPL solver (model/assembly) aliases them so its call
// sites are unaffected (ADR-0018).
//
// Values are STABLE ACROSS SESSIONS and must never be renumbered: an assembly
// recipe persists a constraint's kind and solution type.

// AssemblyConstraintType discriminates the assembly relationship kinds. The classic
// set (mate…symmetry) positions geometry; the motion set (rotate-rotate…
// translate-translate) couples driven values; transitional models sliding contact;
// custom is solved by an add-in.
type AssemblyConstraintType uint32

const (
	// ConstraintUnknown is the zero value: an unresolved or not-yet-typed constraint.
	ConstraintUnknown AssemblyConstraintType = 0
	// ConstraintMate makes two faces/edges/points coincident (normals opposed by default).
	ConstraintMate AssemblyConstraintType = 1
	// ConstraintFlush aligns two faces so their normals point the same way (co-planar).
	ConstraintFlush AssemblyConstraintType = 2
	// ConstraintAngle holds a fixed angle between two directions.
	ConstraintAngle AssemblyConstraintType = 3
	// ConstraintTangent keeps a face tangent to a curved face (inside or outside).
	ConstraintTangent AssemblyConstraintType = 4
	// ConstraintInsert combines an axis mate with a plane mate (a bolt into a hole).
	ConstraintInsert AssemblyConstraintType = 5
	// ConstraintSymmetry positions two occurrences symmetrically about a plane.
	ConstraintSymmetry AssemblyConstraintType = 6
	// ConstraintRotateRotate couples two rotations by a gear ratio.
	ConstraintRotateRotate AssemblyConstraintType = 7
	// ConstraintRotateTranslate couples a rotation to a translation (rack and pinion).
	ConstraintRotateTranslate AssemblyConstraintType = 8
	// ConstraintTranslateTranslate couples two translations by a ratio.
	ConstraintTranslateTranslate AssemblyConstraintType = 9
	// ConstraintTransitional keeps a face in sliding contact with a set of faces.
	ConstraintTransitional AssemblyConstraintType = 10
	// ConstraintCustom is a relationship solved by an add-in, not the built-in solver.
	ConstraintCustom AssemblyConstraintType = 11
)

// IsValid reports whether t names a real constraint kind (not unknown).
func (t AssemblyConstraintType) IsValid() bool {
	return t >= ConstraintMate && t <= ConstraintCustom
}

// String returns a stable lowercase name, used in diagnostics and the wire DTOs. The
// value, not this name, is the persisted identity.
func (t AssemblyConstraintType) String() string {
	switch t {
	case ConstraintMate:
		return "mate"
	case ConstraintFlush:
		return "flush"
	case ConstraintAngle:
		return "angle"
	case ConstraintTangent:
		return "tangent"
	case ConstraintInsert:
		return "insert"
	case ConstraintSymmetry:
		return "symmetry"
	case ConstraintRotateRotate:
		return "rotate-rotate"
	case ConstraintRotateTranslate:
		return "rotate-translate"
	case ConstraintTranslateTranslate:
		return "translate-translate"
	case ConstraintTransitional:
		return "transitional"
	case ConstraintCustom:
		return "custom"
	default:
		return "unknown"
	}
}

// MateConstraintSolutionType discriminates how a mate between two directional faces
// resolves: opposed normals (a true mate) or aligned normals (a flush). It is the
// directed sense the solver enforces on the two faces' normals.
type MateConstraintSolutionType uint32

const (
	// MateSolutionOpposed is the default mate: the two face normals point at each other.
	MateSolutionOpposed MateConstraintSolutionType = 0
	// MateSolutionAligned is the flush sense: the two face normals point the same way.
	MateSolutionAligned MateConstraintSolutionType = 1
	// MateSolutionUndirected resolves to whichever normal sense the parts already hold, so a drag
	// or drive never forces a flip — the reference CAD API's kUndirectedSolutionType (#1971).
	MateSolutionUndirected MateConstraintSolutionType = 2
	// MateSolutionNoSolution leaves the directional sense unconstrained, holding only the offset —
	// the reference CAD API's kNoSolutionType (#1971).
	MateSolutionNoSolution MateConstraintSolutionType = 3
)

// String returns a stable lowercase name.
func (s MateConstraintSolutionType) String() string {
	switch s {
	case MateSolutionAligned:
		return "aligned"
	case MateSolutionUndirected:
		return "undirected"
	case MateSolutionNoSolution:
		return "noSolution"
	default:
		return "opposed"
	}
}

// ParseMateConstraintSolutionType resolves a wire spelling to a mate solution; "" ⇒ opposed.
func ParseMateConstraintSolutionType(s string) (MateConstraintSolutionType, bool) {
	switch s {
	case "", "opposed":
		return MateSolutionOpposed, true
	case "aligned":
		return MateSolutionAligned, true
	case "undirected":
		return MateSolutionUndirected, true
	case "noSolution":
		return MateSolutionNoSolution, true
	}
	return MateSolutionOpposed, false
}

// AngleConstraintSolutionType discriminates how an angle constraint measures its
// angle: undirected (the unsigned angle), directed about an explicit reference axis
// (so the sign and full 0–360° range are meaningful), or the reference-vector form
// that names the third axis directly.
type AngleConstraintSolutionType uint32

const (
	// AngleSolutionUndirected is the default: the unsigned angle between the directions.
	AngleSolutionUndirected AngleConstraintSolutionType = 0
	// AngleSolutionDirected measures the signed angle about an implied reference axis.
	AngleSolutionDirected AngleConstraintSolutionType = 1
	// AngleSolutionReferenceVector measures the angle about an explicit reference vector.
	AngleSolutionReferenceVector AngleConstraintSolutionType = 2
)

// String returns a stable lowercase name.
func (s AngleConstraintSolutionType) String() string {
	switch s {
	case AngleSolutionDirected:
		return "directed"
	case AngleSolutionReferenceVector:
		return "reference-vector"
	default:
		return "undirected"
	}
}

// HealthStatus is the condition a constraint (or any modeling entity) reports when it
// cannot be fully evaluated — the public, Apache-2.0 form of the host's health
// vocabulary (model/health). A constraint whose geometry vanished goes Sick and is
// re-selectable; an over-constrained assembly that still solved reports Warning.
type HealthStatus uint8

const (
	// HealthOK: the entity is fully evaluated and valid.
	HealthOK HealthStatus = 0
	// HealthWarning: usable but flagged (e.g. redundant constraints that still solved).
	HealthWarning HealthStatus = 1
	// HealthSick: cannot be evaluated and needs user attention (e.g. a lost reference).
	HealthSick HealthStatus = 2
	// HealthSuppressed: intentionally excluded from evaluation by the user.
	HealthSuppressed HealthStatus = 3
)

// String returns a stable lowercase name, byte-compatible with the host's
// health.Status.String() so the two vocabularies agree on the wire.
func (h HealthStatus) String() string {
	switch h {
	case HealthWarning:
		return "warning"
	case HealthSick:
		return "sick"
	case HealthSuppressed:
		return "suppressed"
	default:
		return "ok"
	}
}
