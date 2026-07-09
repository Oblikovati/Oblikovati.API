// SPDX-License-Identifier: Apache-2.0

package featureargs

// The dress-up (edge/face-referencing) feature kinds — fillet, chamfer, rule fillet, full-round
// fillet, draft, shell, lip — promoted from the host's opaque path (#1709). Each references an
// existing body's edges or faces by reference key (get_reference_keys).

const (
	KindFillet          = "fillet"
	KindChamfer         = "chamfer"
	KindRuleFillet      = "ruleFillet"
	KindFullRoundFillet = "fullRoundFillet"
	KindDraft           = "draft"
	KindShell           = "shell"
	KindLip             = "lip"
)

// FilletRadiusPoint is one intermediate radius stop on a variable fillet edge (#695).
type FilletRadiusPoint struct {
	T      float64 `json:"t"`
	Radius string  `json:"radius"`
}

// FilletEdgeSet is one fillet edge set: constant (Radius) or variable (StartRadius+EndRadius over
// exactly one edge, plus optional intermediate RadiusPoints, #695).
type FilletEdgeSet struct {
	EdgeRefs     []string            `json:"edgeRefs"`
	Radius       string              `json:"radius,omitempty"`
	StartRadius  string              `json:"startRadius,omitempty"`
	EndRadius    string              `json:"endRadius,omitempty"`
	RadiusPoints []FilletRadiusPoint `json:"radiusPoints,omitempty"`
}

// Fillet rounds picked edges of a body (KindFillet). EdgeSets is the edge-set form (#323), taking
// precedence over the flat EdgeRefs+Radius pair; FaceRefsA/B is the face-fillet form (#694).
type Fillet struct {
	EdgeRefs        []string        `json:"edgeRefs"`
	Radius          string          `json:"radius,omitempty"`
	EdgeSets        []FilletEdgeSet `json:"edgeSets,omitempty"`
	FaceRefsA       []string        `json:"faceRefsA,omitempty"`
	FaceRefsB       []string        `json:"faceRefsB,omitempty"`
	CornerType      string          `json:"cornerType,omitempty"`      // shared-corner treatment (default miter)
	CrossSection    string          `json:"crossSection,omitempty"`    // blend cross-section (default arc; #1284)
	Rho             float64         `json:"rho,omitempty"`             // conic fullness (0<ρ<1, 0.5=parabola)
	ConcaveStrategy string          `json:"concaveStrategy,omitempty"` // concave-edge handling (default outward)
	// EdgesGeom selects the rounded edges by GEOMETRY (midpoint + direction) instead of EdgeRefs
	// keys, so the binding survives recompute — the way an external author references edges it did
	// not mint keys for. When set it supplies the edges; EdgeRefs becomes optional. See [GeomEdgeSel].
	EdgesGeom []GeomEdgeSel `json:"edgesGeom,omitempty"`
}

// Kind reports the feature kind Fillet creates.
func (Fillet) Kind() string { return KindFillet }

// Chamfer bevels picked edges of a body by a setback distance (KindChamfer).
type Chamfer struct {
	EdgeRefs        []string `json:"edgeRefs"`
	Distance        string   `json:"distance,omitempty"`
	ChamferType     string   `json:"chamferType,omitempty"`     // mode (default distance)
	Distance2       string   `json:"distance2,omitempty"`       // twoDistances second face
	Angle           string   `json:"angle,omitempty"`           // distanceAndAngle
	ConcaveStrategy string   `json:"concaveStrategy,omitempty"` // concave-edge handling (default outward)
	// EdgesGeom selects the bevelled edges by GEOMETRY (midpoint + direction) instead of EdgeRefs
	// keys, so the binding survives recompute (see [Fillet.EdgesGeom]). When set, EdgeRefs is optional.
	EdgesGeom []GeomEdgeSel `json:"edgesGeom,omitempty"`
}

// Kind reports the feature kind Chamfer creates.
func (Chamfer) Kind() string { return KindChamfer }

// RuleFillet rounds a whole dihedral class of a body's edges in one feature (KindRuleFillet).
type RuleFillet struct {
	Rule   string `json:"rule"`
	Radius string `json:"radius"`
}

// Kind reports the feature kind RuleFillet creates.
func (RuleFillet) Kind() string { return KindRuleFillet }

// FullRoundFillet replaces a center face with a full round tangent to two parallel side faces
// (KindFullRoundFillet).
type FullRoundFillet struct {
	Side1Ref  string `json:"side1Ref"`
	CenterRef string `json:"centerRef"`
	Side2Ref  string `json:"side2Ref"`
}

// Kind reports the feature kind FullRoundFillet creates.
func (FullRoundFillet) Kind() string { return KindFullRoundFillet }

// Draft tapers picked faces by a draft angle (KindDraft).
type Draft struct {
	FaceRefs []string `json:"faceRefs"`
	Angle    string   `json:"angle,omitempty"`
	// PullDirection is the explicit pull/parting direction as a unit vector (matches
	// InventorDraft.Pull); nil ⇒ the host infers it from the neutral faces (current behavior).
	PullDirection []float64 `json:"pullDirection,omitempty"`
	// FacesGeom selects the drafted faces by GEOMETRY (centroid + normal) instead of FaceRefs keys,
	// so the binding survives recompute (see [Fillet.EdgesGeom]). When set, FaceRefs is optional.
	FacesGeom []GeomFaceSel `json:"facesGeom,omitempty"`
	// NeutralPlane names the fixed (neutral) plane for a fixed-plane face draft — Inventor's
	// kFixedPlaneFaceDraftDefinitionType. Each drafted face pivots on the line where it meets this
	// plane, so dimensions in the plane are preserved. Value is a planar face reference key, a work
	// plane ("plane/N"), or an origin plane ("origin/plane/xy"). When set and PullDirection is
	// empty, the pull defaults to the neutral plane's normal. #1866.
	NeutralPlane string `json:"neutralPlane,omitempty"`
}

// Kind reports the feature kind Draft creates.
func (Draft) Kind() string { return KindDraft }

// Shell hollows a body to a wall thickness, removing the picked faces (KindShell).
type Shell struct {
	FaceRefs  []string `json:"faceRefs"`
	Thickness string   `json:"thickness,omitempty"`
	// FacesGeom selects the removed faces by GEOMETRY (centroid + normal) instead of FaceRefs keys,
	// so the binding survives recompute (see [Fillet.EdgesGeom]). When set, FaceRefs is optional.
	FacesGeom []GeomFaceSel `json:"facesGeom,omitempty"`
}

// Kind reports the feature kind Shell creates.
func (Shell) Kind() string { return KindShell }

// Lip runs a raised lip (or recessed groove, Groove=true) bead along picked edges (KindLip).
type Lip struct {
	EdgeRefs []string `json:"edgeRefs"`
	Width    string   `json:"width"`
	Height   string   `json:"height"`
	Groove   bool     `json:"groove,omitempty"`
}

// Kind reports the feature kind Lip creates.
func (Lip) Kind() string { return KindLip }

// dressupArgs is the dress-up family's contribution to [All].
var dressupArgs = []Arg{
	Fillet{}, Chamfer{}, RuleFillet{}, FullRoundFillet{}, Draft{}, Shell{}, Lip{},
}
