// SPDX-License-Identifier: Apache-2.0

package featureargs

// The sheet-metal feature kinds (#1709). Every one requires the active part to be in the
// sheet-metal environment. Edge/edges are reference keys (get_reference_keys); sketch/line inputs
// are indices. Angle/radius default to the sheet-metal rule when omitted. Unfold and Refold take
// no arguments (zero-field arg types, like Hull).

const (
	KindSheetMetalFace          = "sheetMetalFace"
	KindSheetMetalFlange        = "sheetMetalFlange"
	KindSheetMetalHem           = "sheetMetalHem"
	KindSheetMetalBend          = "sheetMetalBend"
	KindSheetMetalFold          = "sheetMetalFold"
	KindSheetMetalCorner        = "sheetMetalCorner"
	KindSheetMetalContourFlange = "sheetMetalContourFlange"
	KindSheetMetalLoftedFlange  = "sheetMetalLoftedFlange"
	KindSheetMetalContourRoll   = "sheetMetalContourRoll"
	KindSheetMetalCornerSeam    = "sheetMetalCornerSeam"
	KindSheetMetalCut           = "sheetMetalCut"
	KindSheetMetalRip           = "sheetMetalRip"
	KindSheetMetalPunch         = "sheetMetalPunch"
	KindSheetMetalLip           = "sheetMetalLip"
	KindSheetMetalCosmeticBend  = "sheetMetalCosmeticBend"
	KindSheetMetalUnfold        = "sheetMetalUnfold"
	KindSheetMetalRefold        = "sheetMetalRefold"
)

// SheetMetalFace adds a base wall (operation new) or a secondary wall (join) from a sketch profile
// (KindSheetMetalFace).
type SheetMetalFace struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Operation    string `json:"operation"`
	Direction    string `json:"direction,omitempty"`
}

// Kind reports the feature kind SheetMetalFace creates.
func (SheetMetalFace) Kind() string { return KindSheetMetalFace }

// SheetMetalFlange adds a flange along an edge (KindSheetMetalFlange).
type SheetMetalFlange struct {
	Edge   string `json:"edge"`
	Height string `json:"height"`
	Angle  string `json:"angle,omitempty"`
	Radius string `json:"radius,omitempty"`
	Flip   bool   `json:"flip,omitempty"`
	// BendPosition is how far back from the picked edge the bend sits — Inventor's
	// BendPositionEnum (#1957). Two flanges of the same height and angle in different positions
	// are different parts, because the position decides whether the wall overhangs the edge or
	// finishes flush with it. "adjacentFace" (default) starts the bend AT the edge;
	// "outsideBaseFace" and "insideBendFace" set it back until the wall's outer or inner face
	// reaches the edge; "outerEdgeOffset" and "innerEdgeOffset" are those two plus PositionOffset.
	BendPosition string `json:"bendPosition,omitempty"`
	// PositionOffset is the explicit distance for the two edge-offset positions.
	PositionOffset string `json:"positionOffset,omitempty"`
	// HeightDatum is what Height is measured FROM — Inventor's HeightDatumTypeEnum: "tangent"
	// (default; where the bend ends), "outer" or "inner" (the sharp corner the outer/inner faces
	// would make, which is how a drawing dimensions it), or "outerOrtho"/"innerOrtho" (the same
	// corners measured perpendicular to the base face, a different number on any bend that is not
	// a right angle).
	HeightDatum string `json:"heightDatum,omitempty"`
	// Width is how much of the picked edge the wall covers (#1958). Absent ⇒ the whole edge.
	Width *FlangeWidthExtent `json:"width,omitempty"`
	// Options overrides the style's bend properties for THIS bend only (#1959). Absent ⇒ the style.
	Options *BendOptions `json:"options,omitempty"`
	// ApplyAutoMiter extends this wall and the one it corners with until they meet, then cuts
	// MiterGap between them (#1961). Two walls each stop at their own bend line, so the corner
	// between them is otherwise OPEN. Off by default, so an existing part's corners are unchanged.
	ApplyAutoMiter bool `json:"applyAutoMiter,omitempty"`
	// MiterGap is the gap left on the miter line; absent ⇒ the style's GapSize.
	MiterGap string `json:"miterGap,omitempty"`
}

// BendOptions overrides the sheet-metal style's bend properties for one feature — Inventor's
// BendOptions (#1959). Every field is optional and an omitted one defers to the style, which is
// what makes this an override rather than a restatement of the whole style.
type BendOptions struct {
	// ReliefShape, ReliefWidth and ReliefDepth reshape the notch cut at THIS bend's ends: "round",
	// "straight" or "tear" (no cut), with the notch's width along the bend and depth into the
	// parent.
	ReliefShape string `json:"reliefShape,omitempty"`
	ReliefWidth string `json:"reliefWidth,omitempty"`
	ReliefDepth string `json:"reliefDepth,omitempty"`
	// MinimumRemnant is the thinnest strip of parent material a relief may leave standing. A notch
	// that would leave less takes the sliver with it, since a strip that thin tears off in
	// handling and is not what anyone drew.
	MinimumRemnant string `json:"minimumRemnant,omitempty"`
	// Transition is how the material is shaped where this bend runs into the face beside it —
	// "none", "intersection", "straightLine", "arc" or "trimToBend"; "default" defers to the style.
	Transition string `json:"transition,omitempty"`
	// TransitionArcRadius sizes the arc transition.
	TransitionArcRadius string `json:"transitionArcRadius,omitempty"`
}

// FlangeWidthExtent bounds a flange's wall to part of its edge — a bracket tab on a long edge, or
// a wall that stops short of the corners so the neighbouring flanges have somewhere to go (#1958).
// Type picks which distances apply:
//
//   - "edge" (default) spans the whole edge and takes none;
//   - "centered" takes Width;
//   - "offsets" takes Offset (from the edge's start) and Offset2 (from its end);
//   - "offsetWidth" takes Offset and Width.
//
// Inventor's fifth extent, bounded by two referenced entities rather than by distances, is not
// offered: it needs vertex/plane reference binding the host does not have yet.
type FlangeWidthExtent struct {
	Type    string `json:"type,omitempty"`
	Width   string `json:"width,omitempty"`
	Offset  string `json:"offset,omitempty"`
	Offset2 string `json:"offset2,omitempty"`
}

// Kind reports the feature kind SheetMetalFlange creates.
func (SheetMetalFlange) Kind() string { return KindSheetMetalFlange }

// SheetMetalHem folds a hem along an edge (KindSheetMetalHem) — Inventor's four HemTypeEnum
// shapes. Two of them are driven by a curl rather than a leg, so which dimensions apply depends
// on Type (#1956):
//
//   - "single" (default) and "double" take Length + Gap;
//   - "rolled" and "teardrop" take Radius + Angle, and a teardrop derives its closing tail from
//     them, which is why it takes no length.
type SheetMetalHem struct {
	Edge string `json:"edge"`
	// Length is how far the folded-back leg runs; single and double hems only.
	Length string `json:"length,omitempty"`
	// Type is the hem shape: "single", "double", "rolled" or "teardrop". "closed" and "open" are
	// the spellings this feature shipped with and both mean a single hem — what set them apart was
	// the gap, which is still what says it.
	Type string `json:"type,omitempty"`
	// Gap is the clear distance between the folded-back leg and the parent (the fold's inside
	// radius is half of it). Absent ⇒ the hem folds tight at half the material thickness.
	Gap string `json:"gap,omitempty"`
	// Radius and Angle drive the curled types: the roll's inside radius and how far it sweeps. A
	// teardrop must sweep more than a half-turn and less than a full one for its tail to close.
	Radius string `json:"radius,omitempty"`
	Angle  string `json:"angle,omitempty"`
	Flip   bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind SheetMetalHem creates.
func (SheetMetalHem) Kind() string { return KindSheetMetalHem }

// SheetMetalBend adds a bend across a sketch line (KindSheetMetalBend).
type SheetMetalBend struct {
	SketchIndex int    `json:"sketchIndex"`
	LineIndex   int    `json:"lineIndex"`
	Angle       string `json:"angle,omitempty"`
	Radius      string `json:"radius,omitempty"`
	Flip        bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind SheetMetalBend creates.
func (SheetMetalBend) Kind() string { return KindSheetMetalBend }

// SheetMetalFold folds the part along a sketch line (KindSheetMetalFold).
type SheetMetalFold struct {
	SketchIndex int    `json:"sketchIndex"`
	LineIndex   int    `json:"lineIndex"`
	Angle       string `json:"angle,omitempty"`
	Radius      string `json:"radius,omitempty"`
	Location    string `json:"location,omitempty"`
	Flip        bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind SheetMetalFold creates.
func (SheetMetalFold) Kind() string { return KindSheetMetalFold }

// SheetMetalCorner applies a corner treatment (of the given Size) to picked edges
// (KindSheetMetalCorner, #1967). A chamfer is a single Size by default, or — with ChamferType —
// a distance-and-angle or two-distance bevel; a round carries one Size, or several EdgeSets each
// with its own radius in one feature.
type SheetMetalCorner struct {
	Edges     []string `json:"edges"`
	Treatment string   `json:"treatment"`
	Size      string   `json:"size,omitempty"`
	// ChamferType picks the chamfer's setback shape (ChamferType): "distance" (default, equal both
	// faces), "distanceAndAngle" (Size on one face, Angle to it) or "twoDistances" (Size + DistanceTwo).
	ChamferType string `json:"chamferType,omitempty"`
	// DistanceTwo is the second setback for the two-distance chamfer; Angle is the bevel angle for
	// the distance-and-angle chamfer; FaceKey names the face Size is measured on (both variants).
	DistanceTwo string `json:"distanceTwo,omitempty"`
	Angle       string `json:"angle,omitempty"`
	FaceKey     string `json:"faceKey,omitempty"`
	// EdgeSets rounds several corner-edge groups, each with its own Radius, in one feature (a corner
	// round with multiple radii). When present it supersedes Edges/Size for the round treatment.
	EdgeSets []CornerRoundEdgeSet `json:"edgeSets,omitempty"`
}

// CornerRoundEdgeSet is one radius group of a multi-set corner round (#1967).
type CornerRoundEdgeSet struct {
	Edges  []string `json:"edges"`
	Radius string   `json:"radius"`
}

// Kind reports the feature kind SheetMetalCorner creates.
func (SheetMetalCorner) Kind() string { return KindSheetMetalCorner }

// SheetMetalContourFlange runs a contour flange (a profile sketch) along an edge
// (KindSheetMetalContourFlange).
type SheetMetalContourFlange struct {
	Edge          string `json:"edge"`
	ProfileSketch int    `json:"profileSketch"`
	Flip          bool   `json:"flip,omitempty"`
	// Width bounds the swept wall to part of the edge (#1958); absent ⇒ the whole edge.
	Width *FlangeWidthExtent `json:"width,omitempty"`
	// Operation is how the wall joins the model (#1961): "join" (default) unions it onto the
	// running sheet, "new" starts a body of its own.
	Operation string `json:"operation,omitempty"`
	// Radius rounds the profile's corners into bends; absent ⇒ the rule's BendRadius. A contour
	// flange's corners ARE bends, so a sharp one is geometry no press brake can make.
	Radius string `json:"radius,omitempty"`
}

// Kind reports the feature kind SheetMetalContourFlange creates.
func (SheetMetalContourFlange) Kind() string { return KindSheetMetalContourFlange }

// SheetMetalLoftedFlange lofts a flange between two profile sketches (KindSheetMetalLoftedFlange,
// #1966).
type SheetMetalLoftedFlange struct {
	ProfileA  int    `json:"profileA"`
	ProfileB  int    `json:"profileB"`
	Operation string `json:"operation"`
	// OutputType is how the transition is calculated (LoftedFlangeOutputType): "dieFormed"
	// (default, smooth) or one of the "pressBrake…" faceted modes. The press-brake modes read
	// FacetTolerance, each measuring it differently (chord deviation / facet angle / facet width).
	OutputType string `json:"outputType,omitempty"`
	// FacetTolerance bounds the facets for a press-brake output — a length for the chord/distance
	// modes ("0.5 mm") or an angle for the facet-angle mode ("5 deg"). Ignored for die-formed.
	FacetTolerance string `json:"facetTolerance,omitempty"`
	// Converge merges the wall's corners to a point rather than carrying them straight through.
	Converge bool `json:"converge,omitempty"`
	// Radius rounds the transition's end bends; absent ⇒ the rule's BendRadius.
	Radius string `json:"radius,omitempty"`
}

// Kind reports the feature kind SheetMetalLoftedFlange creates.
func (SheetMetalLoftedFlange) Kind() string { return KindSheetMetalLoftedFlange }

// SheetMetalContourRoll rolls a profile about an axis line (KindSheetMetalContourRoll).
type SheetMetalContourRoll struct {
	ProfileSketch int    `json:"profileSketch"`
	AxisLine      int    `json:"axisLine"`
	Angle         string `json:"angle,omitempty"`
	Operation     string `json:"operation"`
}

// Kind reports the feature kind SheetMetalContourRoll creates.
func (SheetMetalContourRoll) Kind() string { return KindSheetMetalContourRoll }

// SheetMetalCornerSeam finishes the seam at picked corner edges where two flange walls meet
// (KindSheetMetalCornerSeam, #1964).
type SheetMetalCornerSeam struct {
	Edges []string `json:"edges"`
	// Gap is the relief left between the two walls (the gap seam) or the seam clearance for the
	// other types — Inventor's GapWidth.
	Gap string `json:"gap"`
	// Type is the seam finish: "gap" (default), "overlap", "reverseOverlap" or "noOverlap"
	// (CornerSeamType). The last three lap or butt the walls rather than gapping them.
	Type string `json:"type,omitempty"`
	// Overlap is how far one wall laps over the other, as a percentage 0–100 — Inventor's
	// PercentOverlap. It applies to the overlap and reverseOverlap types; ignored for gap/noOverlap.
	Overlap float64 `json:"overlap,omitempty"`
	// ReliefShape and ReliefSize cut a relief at the seam root — the CornerReliefShape spelling
	// ("round", "square", "tear", …) and its size as a length. Absent ⇒ no seam-root relief.
	ReliefShape string `json:"reliefShape,omitempty"`
	ReliefSize  string `json:"reliefSize,omitempty"`
	// DefinitionType is how the gap is measured — "maxDistance" (default) or "faceEdgeDistance"
	// (CornerSeamDefinitionType). They agree on a square miter and differ on an oblique corner.
	DefinitionType string `json:"definitionType,omitempty"`
}

// Kind reports the feature kind SheetMetalCornerSeam creates.
func (SheetMetalCornerSeam) Kind() string { return KindSheetMetalCornerSeam }

// SheetMetalCut cuts a sketch profile through the wall (or across a bend) (KindSheetMetalCut).
type SheetMetalCut struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Direction    string `json:"direction,omitempty"`
	Distance     string `json:"distance,omitempty"`
	AcrossBend   bool   `json:"acrossBend,omitempty"`
}

// Kind reports the feature kind SheetMetalCut creates.
func (SheetMetalCut) Kind() string { return KindSheetMetalCut }

// SheetMetalRip rips a wall open by a gap (KindSheetMetalRip, #1965). A rip has three forms —
// Inventor's RipTypeEnum — and they take different inputs:
//
//   - "pointToPoint" (default) rips between two points. Given a SketchIndex/LineIndex it rips
//     along that sketch line (the long-standing form); given FaceKey + Point + PointTwo it rips
//     between two vertices of that face.
//   - "singlePoint" rips the full extent of FaceKey through the one vertex Point.
//   - "faceExtents" rips the full extent of FaceKey and needs no point.
//
// GapSide places the removed material: "symmetric" (default) straddles the rip line, "positive"
// / "negative" take it wholly to one side (Inventor's PartFeatureExtentDirectionEnum).
type SheetMetalRip struct {
	SketchIndex int    `json:"sketchIndex,omitempty"`
	LineIndex   int    `json:"lineIndex,omitempty"`
	Gap         string `json:"gap,omitempty"`
	// Type is the rip form: "pointToPoint" (default), "singlePoint" or "faceExtents" (RipType).
	Type string `json:"type,omitempty"`
	// FaceKey is the reference key of the face being ripped (the RipFace); required for the
	// singlePoint and faceExtents forms, and for a point-to-point rip defined by two face vertices.
	FaceKey string `json:"faceKey,omitempty"`
	// Point / PointTwo are reference keys of the face vertices the rip runs through — Point for
	// singlePoint, both for a vertex-defined point-to-point rip.
	Point    string `json:"point,omitempty"`
	PointTwo string `json:"pointTwo,omitempty"`
	// GapSide is which side of the rip line the gap sits: "positive", "negative" or "symmetric"
	// (default).
	GapSide string `json:"gapSide,omitempty"`
}

// Kind reports the feature kind SheetMetalRip creates.
func (SheetMetalRip) Kind() string { return KindSheetMetalRip }

// SheetMetalPunch punches a sketch-driven tool into the wall to a depth (KindSheetMetalPunch,
// #1968).
type SheetMetalPunch struct {
	SketchIndex int    `json:"sketchIndex"`
	Depth       string `json:"depth,omitempty"`
	// Angle rotates the punched profiles about their centroid (e.g. "30 deg") — the die's rotation.
	Angle string `json:"angle,omitempty"`
	// AcrossBends lets the punch span a bent region; UnfoldInFlat controls whether it develops into
	// the flat pattern. RepresentationType and ToolID are die metadata carried into the flat/drawing.
	AcrossBends        bool   `json:"acrossBends,omitempty"`
	UnfoldInFlat       bool   `json:"unfoldInFlat,omitempty"`
	RepresentationType string `json:"representationType,omitempty"`
	ToolID             string `json:"toolId,omitempty"`
}

// Kind reports the feature kind SheetMetalPunch creates.
func (SheetMetalPunch) Kind() string { return KindSheetMetalPunch }

// SheetMetalLip runs a lip along an edge (KindSheetMetalLip).
type SheetMetalLip struct {
	Edge         string `json:"edge"`
	Height       string `json:"height"`
	ReturnLength string `json:"returnLength,omitempty"`
	Angle        string `json:"angle,omitempty"`
	Radius       string `json:"radius,omitempty"`
	Flip         bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind SheetMetalLip creates.
func (SheetMetalLip) Kind() string { return KindSheetMetalLip }

// SheetMetalCosmeticBend adds a cosmetic (display-only) bend across a sketch line
// (KindSheetMetalCosmeticBend).
type SheetMetalCosmeticBend struct {
	SketchIndex int    `json:"sketchIndex"`
	LineIndex   int    `json:"lineIndex"`
	Angle       string `json:"angle,omitempty"`
	Radius      string `json:"radius,omitempty"`
}

// Kind reports the feature kind SheetMetalCosmeticBend creates.
func (SheetMetalCosmeticBend) Kind() string { return KindSheetMetalCosmeticBend }

// SheetMetalUnfold flattens every bend of the active sheet-metal part (KindSheetMetalUnfold). It
// has no fields — it acts on the part's recorded bends.
type SheetMetalUnfold struct{}

// Kind reports the feature kind SheetMetalUnfold creates.
func (SheetMetalUnfold) Kind() string { return KindSheetMetalUnfold }

// SheetMetalRefold re-folds the bends an earlier unfold flattened (KindSheetMetalRefold). It has
// no fields.
type SheetMetalRefold struct{}

// Kind reports the feature kind SheetMetalRefold creates.
func (SheetMetalRefold) Kind() string { return KindSheetMetalRefold }

// sheetMetalArgs is the sheet-metal family's contribution to [All].
var sheetMetalArgs = []Arg{
	SheetMetalFace{}, SheetMetalFlange{}, SheetMetalHem{}, SheetMetalBend{}, SheetMetalFold{},
	SheetMetalCorner{}, SheetMetalContourFlange{}, SheetMetalLoftedFlange{}, SheetMetalContourRoll{},
	SheetMetalCornerSeam{}, SheetMetalCut{}, SheetMetalRip{}, SheetMetalPunch{}, SheetMetalLip{},
	SheetMetalCosmeticBend{}, SheetMetalUnfold{}, SheetMetalRefold{},
}
