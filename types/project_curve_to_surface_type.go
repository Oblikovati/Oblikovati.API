// SPDX-License-Identifier: Apache-2.0

package types

// ProjectCurveToSurfaceType is how a source curve is projected onto a target surface (Inventor
// ProjectCurveToSurfaceTypeEnum, #1841): perpendicular to the closest point, along a fixed
// direction vector, or wrapped (arc-length) onto the surface. The values are a frozen block
// matching the reference enum. The zero value is not a member; an omitted selector parses to
// closest-point (the pre-#1841 default).
type ProjectCurveToSurfaceType int32

const (
	// ProjectAlongVector intersects, per source sample, the ray from the sample along
	// ProjectDirection with the surface (kProjectAlongVectorType).
	ProjectAlongVector ProjectCurveToSurfaceType = 117505
	// ProjectToClosestPoint drops each source sample to its perpendicular foot on the surface —
	// the default (kProjectToClosestPointType).
	ProjectToClosestPoint ProjectCurveToSurfaceType = 117506
	// ProjectWrapToSurface wraps the source curve onto the surface preserving arc length
	// (kWrapToSurfaceType) — the MapPointCurve result.
	ProjectWrapToSurface ProjectCurveToSurfaceType = 117507
)

// projectCurveToSurfaceTypeNames are the frozen wire spellings.
var projectCurveToSurfaceTypeNames = map[ProjectCurveToSurfaceType]string{
	ProjectAlongVector:    "alongVector",
	ProjectToClosestPoint: "closestPoint",
	ProjectWrapToSurface:  "wrap",
}

// String returns the projection type's wire spelling.
func (p ProjectCurveToSurfaceType) String() string {
	return enumName(projectCurveToSurfaceTypeNames, p, "enum(?)")
}

// ParseProjectCurveToSurfaceType resolves a wire spelling back to its projection type; the empty
// string maps to closest-point so an omitted selector is the default (#1841).
func ParseProjectCurveToSurfaceType(s string) (ProjectCurveToSurfaceType, bool) {
	if s == "" {
		return ProjectToClosestPoint, true
	}
	return enumFromName(projectCurveToSurfaceTypeNames, s)
}
