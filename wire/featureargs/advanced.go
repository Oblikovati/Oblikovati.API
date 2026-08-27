// SPDX-License-Identifier: Apache-2.0

package featureargs

// The feature kinds that need a custom host resolver (#1709): sweep (a profile along a path),
// move body, bend part, replace face, core/cavity tooling, and split solid. Enum spellings mirror
// the frozen api/types values (sweep definition/orientation/scaling, bend type, split type).

const (
	KindSweep       = "sweep"
	KindMoveBody    = "moveBody"
	KindBendPart    = "bendPart"
	KindReplaceFace = "replaceFace"
	KindCoreCavity  = "coreCavity"
	KindSplitSolid  = "splitSolid"
)

// SweepTwistStation is one pathAndSectionTwists row: a twist Angle at normalized arclength T.
type SweepTwistStation struct {
	T     float64 `json:"t"`
	Angle string  `json:"angle"`
}

// Sweep sweeps a profile along an open path (rail) into a solid (KindSweep). DefinitionType
// discriminates the union (path | pathAndGuideRail | pathAndGuideSurface | pathAndSectionTwists |
// solid), mirroring the frozen api/types sweep spellings (#314).
type Sweep struct {
	SketchIndex     int                 `json:"sketchIndex"`
	ProfileIndex    int                 `json:"profileIndex"`
	PathSketchIndex int                 `json:"pathSketchIndex"`
	PathIndex       int                 `json:"pathIndex"`
	Twist           string              `json:"twist,omitempty"`
	Operation       string              `json:"operation,omitempty"`
	DefinitionType  string              `json:"definitionType,omitempty"`
	Orientation     string              `json:"orientation,omitempty"`
	AlignVector     []float64           `json:"alignVector,omitempty"`
	Taper           string              `json:"taper,omitempty"`
	TwistStations   []SweepTwistStation `json:"twistStations,omitempty"`
	RailSketchIndex int                 `json:"railSketchIndex,omitempty"`
	RailIndex       int                 `json:"railIndex,omitempty"`
	Scaling         string              `json:"scaling,omitempty"`
	GuideFaceKey    string              `json:"guideFaceKey,omitempty"`
	ToolBodyIndex   *int                `json:"toolBodyIndex,omitempty"`
	// PathPoints is an explicit 3D polyline path in cm (each entry an [x,y,z] point), mirroring
	// how LoftRail.Points overrides its sketch path; nil ⇒ use the PathSketchIndex/PathIndex arg.
	PathPoints [][]float64 `json:"pathPoints,omitempty"`
}

// Kind reports the feature kind Sweep creates.
func (Sweep) Kind() string { return KindSweep }

// MoveBodyOp is one entry of an ordered move-operation list (M20-F20): a typed, independently
// parametric step (freeDrag: x/y/z; alongRay: dir+dist; rotateAboutLine: point+dir+angle).
type MoveBodyOp struct {
	Type  string    `json:"type"`
	X     string    `json:"x,omitempty"`
	Y     string    `json:"y,omitempty"`
	Z     string    `json:"z,omitempty"`
	Dir   []float64 `json:"dir,omitempty"`
	Dist  string    `json:"dist,omitempty"`
	Point []float64 `json:"point,omitempty"`
	Angle string    `json:"angle,omitempty"`
}

// MoveBody moves a solid body by a translation or an ordered list of parametric operations
// (KindMoveBody). Operations, when given, supersedes Translation.
type MoveBody struct {
	BodyIndex   int          `json:"bodyIndex"`
	Translation []float64    `json:"translation,omitempty"`
	Operations  []MoveBodyOp `json:"operations,omitempty"`
}

// Kind reports the feature kind MoveBody creates.
func (MoveBody) Kind() string { return KindMoveBody }

// BendPart bends a solid body around a sketch bend line (KindBendPart). BendType picks which two
// of radius/angle/arcLength drive the bend (the third is derived).
type BendPart struct {
	SketchIndex int    `json:"sketchIndex"`
	LineIndex   int    `json:"lineIndex,omitempty"`
	BendType    string `json:"bendType,omitempty"`
	Radius      string `json:"radius,omitempty"`
	Angle       string `json:"angle,omitempty"`
	ArcLength   string `json:"arcLength,omitempty"`
	Flip        bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind BendPart creates.
func (BendPart) Kind() string { return KindBendPart }

// ReplaceFace replaces picked faces with new face geometry (KindReplaceFace). NewFaceRefs is the
// #1886 replacement set — each entry is a planar face key, a work plane ("plane/N",
// "origin/plane/xy"), and the faces may come from different bodies; each picked face is retrimmed
// onto its nearest new face. TargetRef is the legacy single same-body target (kept associative);
// NewFaceRefs takes precedence when present.
type ReplaceFace struct {
	FaceRefs    []string `json:"faceRefs"`
	TargetRef   string   `json:"targetRef,omitempty"`
	NewFaceRefs []string `json:"newFaceRefs,omitempty"`
}

// Kind reports the feature kind ReplaceFace creates.
func (ReplaceFace) Kind() string { return KindReplaceFace }

// CoreCavity splits the body at a parting plane into core and cavity tooling (KindCoreCavity).
type CoreCavity struct {
	Axis      string  `json:"axis,omitempty"`
	Position  string  `json:"position"`
	Shrinkage float64 `json:"shrinkage,omitempty"`
}

// Kind reports the feature kind CoreCavity creates.
func (CoreCavity) Kind() string { return KindCoreCavity }

// SplitSolid splits the solid along a work plane: trim one side, split into bodies, or split faces
// only (KindSplitSolid). Type mirrors the frozen api/types SplitType (#330).
type SplitSolid struct {
	WorkPlaneIndex int    `json:"workPlaneIndex"`
	Keep           string `json:"keep,omitempty"`
	Type           string `json:"type,omitempty"`
	// Tool is the GEOMETRY the split cuts with — the reference CAD API's SplitToolTypeEnum (#1891):
	// "workPlane" (default, uses WorkPlaneIndex), "workSurface" or "surfaceBody" (both use
	// ToolIndex), or "path" (a 2D sketch path projected onto a face). Type says what the split
	// DOES; Tool says what it cuts with, and the two are independent.
	Tool string `json:"tool,omitempty"`
	// ToolIndex addresses the surface tool: a position in the part's work surfaces for
	// "workSurface", or a body index for "surfaceBody". Ignored for "workPlane".
	ToolIndex int `json:"toolIndex,omitempty"`
	// SketchIndex and ProfileIndex address the sketch profile the "path" tool projects onto the
	// faces and scores them with (#2068). Ignored by every other tool. They follow the same
	// sketch/profile addressing every profile-based feature uses (e.g. Extrude).
	SketchIndex  int `json:"sketchIndex,omitempty"`
	ProfileIndex int `json:"profileIndex,omitempty"`
}

// Kind reports the feature kind SplitSolid creates.
func (SplitSolid) Kind() string { return KindSplitSolid }

// advancedArgs is the custom-resolver family's contribution to [All].
var advancedArgs = []Arg{
	Sweep{}, MoveBody{}, BendPart{}, ReplaceFace{}, CoreCavity{}, SplitSolid{},
}
