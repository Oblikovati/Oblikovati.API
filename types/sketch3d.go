// SPDX-License-Identifier: Apache-2.0

package types

// Sketch3DEntityKind discriminates the kind of a 3D (non-planar) sketch entity in the
// wire protocol — the value of
// [oblikovati/api/wire.AddSketch3DEntityArgs.Kind] and of each enumerated
// 3D entity's Kind. The set is the full Inventor Sketch3D geometry family; the string
// values are frozen. Members are wired in across M22 (F02 base curves, F03 conics/
// splines, F04 helix, F11 surface-derived curves).
type Sketch3DEntityKind string

const (
	Sketch3DEntityPoint              Sketch3DEntityKind = "point"
	Sketch3DEntityLine               Sketch3DEntityKind = "line"
	Sketch3DEntityCircle             Sketch3DEntityKind = "circle"
	Sketch3DEntityArc                Sketch3DEntityKind = "arc"
	Sketch3DEntityBend               Sketch3DEntityKind = "bend"
	Sketch3DEntityEllipse            Sketch3DEntityKind = "ellipse"
	Sketch3DEntityEllipticalArc      Sketch3DEntityKind = "ellipticalArc"
	Sketch3DEntitySpline             Sketch3DEntityKind = "spline"
	Sketch3DEntityControlPointSpline Sketch3DEntityKind = "controlPointSpline"
	Sketch3DEntityFixedSpline        Sketch3DEntityKind = "fixedSpline"
	Sketch3DEntityEquationCurve      Sketch3DEntityKind = "equationCurve"
	Sketch3DEntityHelical            Sketch3DEntityKind = "helical"
	Sketch3DEntityIntersection       Sketch3DEntityKind = "intersection"
	Sketch3DEntityOnFace             Sketch3DEntityKind = "onFace"
	Sketch3DEntityProjectToSurface   Sketch3DEntityKind = "projectToSurface"
	Sketch3DEntitySilhouette         Sketch3DEntityKind = "silhouette"
	Sketch3DEntityOffset             Sketch3DEntityKind = "offset"
	// Included reference geometry (Include Geometry: a model vertex/edge linked into the
	// 3D sketch, tracking its source — M22-F08).
	Sketch3DEntityIncludedPoint Sketch3DEntityKind = "includedPoint"
	Sketch3DEntityIncludedCurve Sketch3DEntityKind = "includedCurve"
	Sketch3DEntityUnknown       Sketch3DEntityKind = "unknown"
)

// Geometric3DConstraintKind discriminates a 3D geometric (non-dimensional) constraint —
// the value of [oblikovati/api/wire.AddSketch3DConstraintArgs.Kind] and of
// each enumerated 3D constraint's Kind. String values are frozen. Members are wired in
// across M22-F05.
type Geometric3DConstraintKind string

const (
	Geo3DCoincident        Geometric3DConstraintKind = "coincident"
	Geo3DCollinear         Geometric3DConstraintKind = "collinear"
	Geo3DConcentric        Geometric3DConstraintKind = "concentric"
	Geo3DEqual             Geometric3DConstraintKind = "equal"
	Geo3DParallel          Geometric3DConstraintKind = "parallel"
	Geo3DPerpendicular     Geometric3DConstraintKind = "perpendicular"
	Geo3DTangent           Geometric3DConstraintKind = "tangent"
	Geo3DSmooth            Geometric3DConstraintKind = "smooth"
	Geo3DMidpoint          Geometric3DConstraintKind = "midpoint"
	Geo3DGround            Geometric3DConstraintKind = "ground"
	Geo3DParallelToXAxis   Geometric3DConstraintKind = "parallelToXAxis"
	Geo3DParallelToYAxis   Geometric3DConstraintKind = "parallelToYAxis"
	Geo3DParallelToZAxis   Geometric3DConstraintKind = "parallelToZAxis"
	Geo3DParallelToXYPlane Geometric3DConstraintKind = "parallelToXYPlane"
	Geo3DParallelToXZPlane Geometric3DConstraintKind = "parallelToXZPlane"
	Geo3DParallelToYZPlane Geometric3DConstraintKind = "parallelToYZPlane"
	Geo3DSplineFitPoints   Geometric3DConstraintKind = "splineFitPoints"
	Geo3DBend              Geometric3DConstraintKind = "bend"
	Geo3DHelical           Geometric3DConstraintKind = "helical"
	Geo3DUnknown           Geometric3DConstraintKind = "unknown"
)

// Dimension3DConstraintKind discriminates a 3D dimensional (driving/driven) constraint —
// the value of [oblikovati/api/wire.AddSketch3DDimensionArgs.Kind] and of
// each enumerated 3D dimension's Kind. String values are frozen. Members are wired in
// across M22-F06.
type Dimension3DConstraintKind string

const (
	Dim3DDistance           Dimension3DConstraintKind = "distance"
	Dim3DLineLength         Dimension3DConstraintKind = "lineLength"
	Dim3DRadius             Dimension3DConstraintKind = "radius"
	Dim3DPointPlaneDistance Dimension3DConstraintKind = "pointPlaneDistance"
	Dim3DTwoLineAngle       Dimension3DConstraintKind = "twoLineAngle"
	Dim3DSplineLength       Dimension3DConstraintKind = "splineLength"
	Dim3DUnknown            Dimension3DConstraintKind = "unknown"
)
