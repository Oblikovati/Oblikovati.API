// SPDX-License-Identifier: Apache-2.0

package featureargs

// The surfacing feature kinds (#1709): build surface bodies (boundary patch, ruled surface),
// modify them (surface offset, extend, mid-surface, stitch, sculpt), and the Class-A NURBS
// surfaces (fill, bridge, network, fair, fit). Continuity spellings mirror api/types
// SurfaceContinuity (g0/g1/g2).

const (
	KindBoundaryPatch  = "boundaryPatch"
	KindRuledSurface   = "ruledSurface"
	KindSurfaceOffset  = "surfaceOffset"
	KindExtend         = "extend"
	KindMidSurface     = "midSurface"
	KindStitch         = "stitch"
	KindSculpt         = "sculpt"
	KindFillSurface    = "fillSurface"
	KindBridgeSurface  = "bridgeSurface"
	KindNetworkSurface = "networkSurface"
	KindFairSurface    = "fairSurface"
	KindFitSurface     = "fitSurface"
)

// BoundaryPatch fills a closed sketch loop with a surface patch (KindBoundaryPatch).
type BoundaryPatch struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Condition    string `json:"condition,omitempty"`
}

// Kind reports the feature kind BoundaryPatch creates.
func (BoundaryPatch) Kind() string { return KindBoundaryPatch }

// RuledSurface sweeps straight rulings off a profile into a surface (KindRuledSurface).
type RuledSurface struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Type         string `json:"type,omitempty"`
	Distance     string `json:"distance"`
}

// Kind reports the feature kind RuledSurface creates.
func (RuledSurface) Kind() string { return KindRuledSurface }

// SurfaceOffset offsets the part's surface bodies by a distance (KindSurfaceOffset).
type SurfaceOffset struct {
	Distance string `json:"distance"`
}

// Kind reports the feature kind SurfaceOffset creates.
func (SurfaceOffset) Kind() string { return KindSurfaceOffset }

// Extend extends a surface body past its boundary edges (KindExtend). EdgeRefs is the #1878
// multi-edge set (EdgeRef is the legacy single edge, read when EdgeRefs is empty). ExtentType
// distance grows by Distance; toPlane/toObject grows each edge until it reaches TargetRef (a work
// plane or planar face). ExtensionType natural|stretched selects the continuity mode (they coincide
// for the planar faces supported today; natural bites on the curved-surface extend).
type Extend struct {
	EdgeRef       string   `json:"edgeRef,omitempty"`
	EdgeRefs      []string `json:"edgeRefs,omitempty"`
	Distance      string   `json:"distance,omitempty"`
	ExtentType    string   `json:"extentType,omitempty"`
	TargetRef     string   `json:"targetRef,omitempty"`
	ExtensionType string   `json:"extensionType,omitempty"`
}

// Kind reports the feature kind Extend creates.
func (Extend) Kind() string { return KindExtend }

// FacePair is one manual mid-surface face pairing by reference key (#1885).
type FacePair struct {
	A string `json:"a"`
	B string `json:"b"`
}

// MidSurface builds mid-surfaces between thin-wall face pairs (KindMidSurface). Auto-pairing takes
// pairs whose separation is within [MinThickness, MaxThickness] on the selected BodyIndices
// (default the last body); FacePairs instead pairs the named faces explicitly (ribs/bosses where
// auto-pairing fails). Each pair's min/max thickness range is reported on the result (#1885).
type MidSurface struct {
	MaxThickness string     `json:"maxThickness,omitempty"`
	MinThickness string     `json:"minThickness,omitempty"`
	BodyIndices  []int      `json:"bodyIndices,omitempty"`
	FacePairs    []FacePair `json:"facePairs,omitempty"`
}

// Kind reports the feature kind MidSurface creates.
func (MidSurface) Kind() string { return KindMidSurface }

// Stitch stitches surface bodies into a quilt (or solid) within Tolerance (KindStitch).
type Stitch struct {
	Tolerance         string `json:"tolerance,omitempty"`
	MaintainAsSurface bool   `json:"maintainAsSurface,omitempty"`
}

// Kind reports the feature kind Stitch creates.
func (Stitch) Kind() string { return KindStitch }

// Sculpt combines surfaces and solids into a sculpted body (KindSculpt).
// SculptSurface is one bounding surface of a sculpt: the surface body (by index) and the direction
// (positive|negative) naming which side of it faces the enclosed volume (#1881).
type SculptSurface struct {
	BodyIndex int    `json:"bodyIndex"`
	Direction string `json:"direction,omitempty"`
}

// Sculpt fills the volume bounded by surface bodies into a solid (KindSculpt). Without Surfaces it
// welds all running bounding surfaces (a closed quilt); with Surfaces it uses the listed bounding
// surfaces and their per-surface directions to close open surfaces (#1881). AffectedBodyIndex is the
// join/cut target (default the last solid).
type Sculpt struct {
	Operation         string          `json:"operation,omitempty"`
	Tolerance         string          `json:"tolerance,omitempty"`
	Surfaces          []SculptSurface `json:"surfaces,omitempty"`
	AffectedBodyIndex *int            `json:"affectedBodyIndex,omitempty"`
}

// Kind reports the feature kind Sculpt creates.
func (Sculpt) Kind() string { return KindSculpt }

// FillSurface closes an N-sided opening bounded by the last N surface bodies with a single NURBS
// (KindFillSurface). Sides omitted/0 or 4 = four-sided; 3, 5, 6… = N-sided (#1300).
type FillSurface struct {
	Continuity string `json:"continuity,omitempty"`
	Sides      int    `json:"sides,omitempty"`
}

// Kind reports the feature kind FillSurface creates.
func (FillSurface) Kind() string { return KindFillSurface }

// BridgeSurface connects the last two surface bodies with a clean NURBS transition, one continuity
// per side (KindBridgeSurface).
type BridgeSurface struct {
	ContinuityA string `json:"continuityA,omitempty"`
	ContinuityB string `json:"continuityB,omitempty"`
}

// Kind reports the feature kind BridgeSurface creates.
func (BridgeSurface) Kind() string { return KindBridgeSurface }

// NetworkSurface interpolates a grid of intersecting U/V curves with a single NURBS (Gordon
// network surface) (KindNetworkSurface). Each curve is a list of [x,y,z] model-space points.
type NetworkSurface struct {
	UCurves [][][]float64 `json:"uCurves"`
	VCurves [][][]float64 `json:"vCurves"`
}

// Kind reports the feature kind NetworkSurface creates.
func (NetworkSurface) Kind() string { return KindNetworkSurface }

// FairSurface smooths curvature wrinkles out of the running surface, holding its boundary
// continuity (KindFairSurface).
type FairSurface struct {
	Continuity string  `json:"continuity,omitempty"`
	Strength   float64 `json:"strength,omitempty"`
	Iterations int     `json:"iterations,omitempty"`
}

// Kind reports the feature kind FairSurface creates.
func (FairSurface) Kind() string { return KindFairSurface }

// FitSurface fits a clean Class-A NURBS surface to a named point cloud's cropped region
// (KindFitSurface).
type FitSurface struct {
	Cloud  string `json:"cloud"`
	Degree int    `json:"degree,omitempty"`
	NU     int    `json:"nu,omitempty"`
	NV     int    `json:"nv,omitempty"`
}

// Kind reports the feature kind FitSurface creates.
func (FitSurface) Kind() string { return KindFitSurface }

// surfaceArgs is the surfacing family's contribution to [All].
var surfaceArgs = []Arg{
	BoundaryPatch{}, RuledSurface{}, SurfaceOffset{}, Extend{}, MidSurface{}, Stitch{},
	Sculpt{}, FillSurface{}, BridgeSurface{}, NetworkSurface{}, FairSurface{}, FitSurface{},
}
