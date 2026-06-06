// SPDX-License-Identifier: Apache-2.0

package types

// LoftType is the loft mode — a plain blend through the sections, or one additionally guided by
// rails / a centerline / area-graph sections. It mirrors the established loft-type set.
//
// This is the canonical, Apache-2.0 definition; the GPL model derives it
// (model/feature.LoftDefinition.LoftType()).
type LoftType string

const (
	// RegularLoft blends through the sections with no extra guides.
	RegularLoft LoftType = "regular"
	// LoftWithRails additionally constrains the surface to follow guide rails.
	LoftWithRails LoftType = "rails"
	// LoftWithCenterline sweeps the sections along a centerline (reserved).
	LoftWithCenterline LoftType = "centerline"
	// LoftWithAreaGraphSections places sections by an area graph along a centerline (reserved).
	LoftWithAreaGraphSections LoftType = "area-graph"
)

// LoftCondition selects how a loft surface leaves the starting section (or arrives at the
// ending section) — the boundary tangency control that lets a loft curve away from a flat
// ruled blend. It mirrors the established loft-condition set: a free (natural) end, an
// angled takeoff measured against the section's sketch plane, tangency/curvature continuity
// with adjacent faces, and the point-section cases.
//
// This is the canonical, Apache-2.0 definition; the GPL model aliases it
// (model/feature.LoftCondition = types.LoftCondition).
type LoftCondition string

const (
	// LoftFree imposes no tangency: the surface ends naturally, so a two-section Free loft is
	// ruled (a straight blend). Angle and impact are ignored.
	LoftFree LoftCondition = "free"
	// LoftAngle makes the surface leave the section at a fixed angle to the section's sketch
	// plane, weighted by an impact (takeoff weight). This is the control that curves a
	// two-section loft (e.g. a flared or necked transition).
	LoftAngle LoftCondition = "angle"
	// LoftDirection is an alias of LoftAngle (the angle/direction takeoff is one condition with
	// two exposed names).
	LoftDirection LoftCondition = "direction"
	// LoftTangent makes the surface tangent to the faces adjacent to the section (requires the
	// section to coincide with an existing body's edge).
	LoftTangent LoftCondition = "tangent"
	// LoftSmooth imposes curvature (G2) continuity with the adjacent faces.
	LoftSmooth LoftCondition = "smooth"
	// LoftSharpPoint ends the loft in a sharp point (a point section).
	LoftSharpPoint LoftCondition = "sharp"
	// LoftTangentToPlane makes the surface tangent to a plane at a point section.
	LoftTangentToPlane LoftCondition = "tangent-to-plane"
)

// CurvesViaAngle reports whether the condition is the angle/direction takeoff — the case
// driven by an angle-to-plane and impact (the others need adjacent faces or point sections).
func (c LoftCondition) CurvesViaAngle() bool { return c == LoftAngle || c == LoftDirection }

// IsTangentToPlane reports the tangent-to-plane point-section condition (a domed apex).
func (c LoftCondition) IsTangentToPlane() bool { return c == LoftTangentToPlane }

// IsSharp reports the sharp-point point-section condition (a straight cone apex).
func (c LoftCondition) IsSharp() bool { return c == LoftSharpPoint }

// IsPointCondition reports whether the condition applies to a point (apex) section.
func (c LoftCondition) IsPointCondition() bool { return c.IsTangentToPlane() || c.IsSharp() }

// IsFaceContinuity reports the conditions that continue an adjacent face's surface across the
// section edge: Tangent (G1) and Smooth (G2). They require the section to be a body face.
func (c LoftCondition) IsFaceContinuity() bool { return c == LoftTangent || c == LoftSmooth }

// IsFree reports whether the condition leaves the end natural (the zero value "" is treated
// as Free, so an unset condition keeps the ruled/natural blend).
func (c LoftCondition) IsFree() bool { return c == "" || c == LoftFree }
