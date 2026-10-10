// SPDX-License-Identifier: Apache-2.0

package types

// The transient-geometry kind discriminators (M01-F05, #602). Numeric values are
// FROZEN to the reference enums (CurveTypeEnum / Curve2dTypeEnum /
// SurfaceTypeEnum / CurveGeometryFormEnum / SurfaceGeometryFormEnum) — clients
// and saved automations depend on them. HelixCurve is an Oblikovati extension
// (the reference has no helix transient curve) continuing the 3D block.

// CurveType identifies a 3D transient curve's concrete kind.
type CurveType int32

const (
	UnknownCurve       CurveType = 5121
	LineCurve          CurveType = 5122
	LineSegmentCurve   CurveType = 5123
	CircleCurve        CurveType = 5124
	CircularArcCurve   CurveType = 5125
	EllipseFullCurve   CurveType = 5126
	EllipticalArcCurve CurveType = 5127
	BSplineCurveKind   CurveType = 5128
	PolylineCurve      CurveType = 5129
	// HelixCurve is the Oblikovati extension for kernel/geom's Helix3d.
	HelixCurve CurveType = 5130
)

var curveTypeNames = map[CurveType]string{
	UnknownCurve: "unknown", LineCurve: "line", LineSegmentCurve: "lineSegment",
	CircleCurve: "circle", CircularArcCurve: "arc", EllipseFullCurve: "ellipse",
	EllipticalArcCurve: "ellipticalArc", BSplineCurveKind: "bspline",
	PolylineCurve: "polyline", HelixCurve: "helix",
}

// String returns the kind's stable name.
func (t CurveType) String() string {
	return enumName(curveTypeNames, t, "curveType(?)")
}

// Curve2dType identifies a 2D transient curve's concrete kind.
type Curve2dType int32

const (
	UnknownCurve2d       Curve2dType = 5249
	LineCurve2d          Curve2dType = 5250
	LineSegmentCurve2d   Curve2dType = 5251
	CircleCurve2d        Curve2dType = 5252
	CircularArcCurve2d   Curve2dType = 5253
	EllipseFullCurve2d   Curve2dType = 5254
	EllipticalArcCurve2d Curve2dType = 5255
	BSplineCurve2dKind   Curve2dType = 5256
	PolylineCurve2d      Curve2dType = 5257
)

var curve2dTypeNames = map[Curve2dType]string{
	UnknownCurve2d: "unknown", LineCurve2d: "line", LineSegmentCurve2d: "lineSegment",
	CircleCurve2d: "circle", CircularArcCurve2d: "arc", EllipseFullCurve2d: "ellipse",
	EllipticalArcCurve2d: "ellipticalArc", BSplineCurve2dKind: "bspline",
	PolylineCurve2d: "polyline",
}

// String returns the kind's stable name.
func (t Curve2dType) String() string {
	return enumName(curve2dTypeNames, t, "curve2dType(?)")
}

// SurfaceType identifies a transient surface's concrete kind.
type SurfaceType int32

const (
	UnknownSurface            SurfaceType = 5889
	PlaneSurface              SurfaceType = 5890
	CylinderSurface           SurfaceType = 5891
	EllipticalCylinderSurface SurfaceType = 5892
	ConeSurface               SurfaceType = 5893
	EllipticalConeSurface     SurfaceType = 5894
	TorusSurface              SurfaceType = 5895
	SphereSurface             SurfaceType = 5896
	BSplineSurfaceKind        SurfaceType = 5897
)

var surfaceTypeNames = map[SurfaceType]string{
	UnknownSurface: "unknown", PlaneSurface: "plane", CylinderSurface: "cylinder",
	EllipticalCylinderSurface: "ellipticalCylinder", ConeSurface: "cone",
	EllipticalConeSurface: "ellipticalCone", TorusSurface: "torus",
	SphereSurface: "sphere", BSplineSurfaceKind: "bspline",
}

// String returns the kind's stable name.
func (t SurfaceType) String() string {
	return enumName(surfaceTypeNames, t, "surfaceType(?)")
}

// CurveGeometryForm is whether a curve is exactly representable as NURBS data.
type CurveGeometryForm int32

const (
	CurveFormNURBS    CurveGeometryForm = 1
	CurveFormNotNURBS CurveGeometryForm = 2
)

// SurfaceGeometryForm is the surface representation classification (flag values,
// frozen to the reference).
type SurfaceGeometryForm int32

const (
	SurfaceFormClosedUVLoops     SurfaceGeometryForm = 1
	SurfaceFormNotClosedUVLoops  SurfaceGeometryForm = 2
	SurfaceFormNURBS             SurfaceGeometryForm = 4
	SurfaceFormNotNURBS          SurfaceGeometryForm = 8
	SurfaceFormProceduralToNURBS SurfaceGeometryForm = 16
)
