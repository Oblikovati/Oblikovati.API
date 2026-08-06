// SPDX-License-Identifier: Apache-2.0

package featureargs

// The kind constants are the single source of the feature-kind strings: both an arg
// type's Kind method and the host's operation registry key on them, so the wire name and
// the registered name cannot drift (#1616).
const (
	KindExtrude    = "extrude"
	KindRevolve    = "revolve"
	KindRib        = "rib"
	KindEmboss     = "emboss"
	KindCoil       = "coil"
	KindHole       = "hole"
	KindBoss       = "boss"
	KindThread     = "thread"
	KindGrill      = "grill"
	KindMesh       = "mesh"
	KindDirectEdit = "directEdit"
)

// Extrude turns a closed sketch profile into a prism (KindExtrude).
type Extrude struct {
	SketchIndex    int    `json:"sketchIndex"`
	ProfileIndex   int    `json:"profileIndex"`
	Distance       string `json:"distance"`                 // e.g. "50 mm", "5 cm"
	Operation      string `json:"operation"`                // join|cut|intersect|new (default new)
	Extent         string `json:"extent,omitempty"`         // distance|through-all|to-next|to-face|from-to|distance-from-face (default distance)
	Direction      string `json:"direction,omitempty"`      // positive|negative|symmetric (default positive)
	SecondDistance string `json:"secondDistance,omitempty"` // asymmetric two-direction depth
	Taper          string `json:"taper,omitempty"`          // draft angle, e.g. "3 deg"
	ToFace         string `json:"toFace,omitempty"`         // to-face / from-to end / distance-from-face target: a planar face key, "plane/N", or "origin/plane/xy"
	// ToFaceGeom names the to-face termination target by GEOMETRY (a planar face's centroid +
	// normal) instead of a ToFace key/plane-ref, so an external author (e.g. an exporter) can
	// terminate an extrude at a body face it did not mint a key for — the extent counterpart of
	// [Hole.PlacementFaceGeom]. The host finds the matching planar face on the current body and
	// freezes its plane as the extent's stop plane. When set it wins over ToFace. See [GeomFaceSel].
	ToFaceGeom *GeomFaceSel `json:"toFaceGeom,omitempty"`
	// FromFace is the START target of a from-to extent (extent "from-to"): a planar face key,
	// "plane/N", or "origin/plane/xy". The prism is bounded below by FromFace and above by ToFace,
	// with the sketch plane supplying only the profile. Unused by the other extents.
	FromFace string `json:"fromFace,omitempty"`
	// FromFaceGeom names the from-to START target by GEOMETRY (a planar face's centroid + normal)
	// instead of a FromFace key/plane-ref — the from-to counterpart of [Extrude.ToFaceGeom]. When
	// set it wins over FromFace.
	FromFaceGeom *GeomFaceSel `json:"fromFaceGeom,omitempty"`
	// ProfileSeeds selects the extruded region(s) by an interior seed point (sketch 2-D, cm),
	// one point per region, instead of ProfileIndex. An external author (e.g. an exporter) cannot
	// predict the host's DCEL region ordering, so it names regions by containment: the host resolves
	// each seed to the region that contains it, ON THE SOLVED SKETCH, every recompute. When present
	// it wins over ProfileIndex. This is the extrude counterpart of the geometric selectors.
	ProfileSeeds [][]float64 `json:"profileSeeds,omitempty"`
}

// Kind reports the feature kind Extrude creates.
func (Extrude) Kind() string { return KindExtrude }

// Revolve sweeps a closed sketch profile about an axis (KindRevolve).
type Revolve struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	AxisRef      string `json:"axisRef,omitempty"`
	// AboutCenterline revolves about the sketch's single centerline (Inventor's "revolve about
	// the sketch centerline") instead of an explicit AxisRef — the way to spin a profile about
	// an internal, tilted axis authored with [AddSketchEntityArgs.Centerline]. When set, AxisRef
	// is ignored; the sketch must have exactly one centerline. See PartDesigner #54.
	AboutCenterline bool   `json:"aboutCenterline,omitempty"`
	Angle           string `json:"angle"`
	Angle2          string `json:"angle2,omitempty"` // second-direction sweep (#313)
	// Direction is the side Angle sweeps to: "positive" (default, forward from the profile),
	// "negative" (the same sweep the other way) or "symmetric" (half the angle each way) — the
	// revolve counterpart of [Extrude.Direction] (#2019). Ignored when Angle2 is set, which is the
	// asymmetric mode and names both sides itself. Unobservable on a full revolution.
	Direction string `json:"direction,omitempty"`
	Operation string `json:"operation,omitempty"`
	// Extent is how the revolve TERMINATES, the revolve half of PartFeatureExtentEnum: "angle"
	// (default — sweep Angle/Angle2, Inventor's kAngleExtent, with a full turn spelled "360 deg"),
	// "to-face" (sweep until the profile reaches ToFace), "from-to" (bounded by FromFace and
	// ToFace) or "to-next" (stop at the next material the sweep meets). The geometric extents
	// terminate a turned part on its own geometry — a groove that stops on a rib wall keeps its
	// parametric link instead of freezing a hand-computed angle. Angle is unused by them.
	Extent string `json:"extent,omitempty"`
	// ToFace is the stop target of the "to-face" extent and the END of "from-to": a planar face
	// reference key, "plane/N", or "origin/plane/xy". A revolve terminator must CONTAIN the revolve
	// axis (a radial face) — only then does the swept solid meet it at one constant sweep angle.
	ToFace string `json:"toFace,omitempty"`
	// ToFaceGeom names the ToFace target by GEOMETRY (a planar face's centroid + normal) instead of
	// a key, for an author that cannot mint one — see [Extrude.ToFaceGeom]. Wins over ToFace.
	ToFaceGeom *GeomFaceSel `json:"toFaceGeom,omitempty"`
	// FromFace is the START target of a "from-to" revolve, named like ToFace. The swept wedge runs
	// backwards from the profile to FromFace and forwards to ToFace, so it always contains the
	// profile; a span that closes on itself is a full revolution.
	FromFace string `json:"fromFace,omitempty"`
	// FromFaceGeom names the "from-to" START target by GEOMETRY instead of a key. Wins over FromFace.
	FromFaceGeom *GeomFaceSel `json:"fromFaceGeom,omitempty"`
	// ProfileSeed selects the revolved region by an interior seed point (sketch 2-D, cm) instead
	// of ProfileIndex, resolved by containment on the solved sketch each recompute (see
	// [Extrude.ProfileSeeds]). When present it wins over ProfileIndex.
	ProfileSeed []float64 `json:"profileSeed,omitempty"`
}

// Kind reports the feature kind Revolve creates.
func (Revolve) Kind() string { return KindRevolve }

// Rib thickens an open sketch profile into a supporting rib/web (KindRib).
type Rib struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	Thickness    string `json:"thickness"`
	Depth        string `json:"depth,omitempty"`
	ToNext       bool   `json:"toNext,omitempty"` // extend to the existing material (#316)
	Operation    string `json:"operation,omitempty"`
}

// Kind reports the feature kind Rib creates.
func (Rib) Kind() string { return KindRib }

// Emboss raises or engraves a sketch profile (or text) on the model (KindEmboss).
type Emboss struct {
	SketchIndex    int    `json:"sketchIndex"`
	ProfileIndices []int  `json:"profileIndices,omitempty"`
	ProfileIndex   int    `json:"profileIndex"`
	TextEntity     uint64 `json:"textEntity,omitempty"`
	Depth          string `json:"depth"`
	Engrave        bool   `json:"engrave,omitempty"`
}

// Kind reports the feature kind Emboss creates.
func (Emboss) Kind() string { return KindEmboss }

// Coil sweeps a profile along a helix about an axis (KindCoil). Two of pitch/revolutions/
// height fix the helix (#316).
type Coil struct {
	SketchIndex  int    `json:"sketchIndex"`
	ProfileIndex int    `json:"profileIndex"`
	AxisRef      string `json:"axisRef,omitempty"`
	Pitch        string `json:"pitch,omitempty"`
	Revolutions  string `json:"revolutions,omitempty"`
	Height       string `json:"height,omitempty"`
	Taper        string `json:"taper,omitempty"`
	Operation    string `json:"operation,omitempty"`
	// Spring end treatment (Inventor CoilFeature.Start/EndTransitionAngle / Start/EndFlatAngle).
	// TransitionAngle winds the pitch down to zero over that sweep; FlatAngle then sweeps flat
	// (constant height) — together they ground/flatten a spring end. Angles are unit-bearing
	// expressions (e.g. "90 deg"); omit for a plain helical end. #1883.
	StartTransitionAngle string `json:"startTransitionAngle,omitempty"`
	StartFlatAngle       string `json:"startFlatAngle,omitempty"`
	EndTransitionAngle   string `json:"endTransitionAngle,omitempty"`
	EndFlatAngle         string `json:"endFlatAngle,omitempty"`
}

// Kind reports the feature kind Coil creates.
func (Coil) Kind() string { return KindCoil }

// Hole drills a hole into a picked face: drilled, counterbore, countersink, or tapped (KindHole).
type Hole struct {
	FaceRef         string `json:"faceRef"`
	Type            string `json:"type,omitempty"` // drilled (default) | counterbore | countersink | tapped
	Diameter        string `json:"diameter"`
	Depth           string `json:"depth,omitempty"` // omit (drilled) => through-all
	CounterDiameter string `json:"counterDiameter,omitempty"`
	CounterDepth    string `json:"counterDepth,omitempty"`
	SinkDiameter    string `json:"sinkDiameter,omitempty"`
	IncludedAngle   string `json:"includedAngle,omitempty"`
	Designation     string `json:"designation,omitempty"`
	// Center is the drill point in model-space cm (matches Inventor's InventorHole.Center);
	// nil ⇒ the picked face's centroid. Needed to place more than one hole on a face.
	Center []float64 `json:"center,omitempty"`
	// CenterExpr is the parameter-expression form of Center: each entry is the expression for
	// one coordinate (x, y, z), overriding the literal Center when present.
	CenterExpr []string `json:"centerExpr,omitempty"`
	// PlacementFaceGeom selects the placement face by GEOMETRY (centroid + normal) instead of a
	// FaceRef key, so the binding survives the hole's own recompute — the way an external author
	// (exporter) must reference a face it did not mint a key for. When set it supplies the face;
	// FaceRef becomes optional. See [GeomFaceSel].
	PlacementFaceGeom *GeomFaceSel `json:"placementFaceGeom,omitempty"`
	// DrillPoint sets a blind drilled/tapped hole's bottom (Inventor HoleDrillPointTypeEnum):
	// "flat" (default) bottoms in a flat disc; "angled" bottoms in a cone of included angle
	// TipAngle. Ignored for a through-all hole. #1863.
	DrillPoint string `json:"drillPoint,omitempty"`
	// TipAngle is the included angle of an "angled" drill point (unit-bearing, e.g. "118 deg");
	// defaults to 118 deg (the standard twist-drill point) when DrillPoint is "angled" and it is
	// omitted. #1863.
	TipAngle string `json:"tipAngle,omitempty"`
	// Tap is the hole's thread FUNCTION, orthogonal to Type (its seat) — Inventor keeps the two on
	// separate axes, so a counterbored tapped hole is an ordinary thing (#1862). "none" (default),
	// "tapped", or "taperTapped" for an NPT-style taper thread. Type "tapped" remains accepted as
	// the older spelling of a drilled hole with Tap "tapped".
	Tap string `json:"tap,omitempty"`
	// ThreadClass is the fit class the tap is cut to, e.g. "6H" (metric) or "2B" (unified).
	ThreadClass string `json:"threadClass,omitempty"`
	// LeftHanded reverses the tap's thread sense; the default is the ordinary right-hand thread.
	LeftHanded bool `json:"leftHanded,omitempty"`
	// Clearance sizes the bore from a fastener table instead of Diameter (#1862), so the FASTENER
	// stays the authored thing and the hole follows when it changes.
	Clearance *HoleClearance `json:"clearance,omitempty"`
	// Placement is the rule LOCATING the bores, Inventor's HolePlacementTypeEnum (#1861): "sketch"
	// (one bore per centre point of PlacementSketchIndex), "linear" (offsets from two edges),
	// "concentric" (on a circular edge's axis) or "point" (a work point along a work axis). Absent
	// ⇒ the single bore on FaceRef at Center, which is the face placement.
	Placement            string `json:"placement,omitempty"`
	PlacementSketchIndex int    `json:"placementSketchIndex,omitempty"`
	// PlacementFlipped drills along the sketch normal / work axis instead of into it.
	PlacementFlipped bool `json:"placementFlipped,omitempty"`
	// ConcentricRef is the circular edge whose axis a "concentric" placement centres on.
	ConcentricRef string `json:"concentricRef,omitempty"`
	// Edge1Ref/Edge2Ref and Offset1/Offset2 locate a "linear" placement: two reference edges of the
	// placement face and the unit-bearing distances measured from each, INTO the face.
	Edge1Ref string `json:"edge1Ref,omitempty"`
	Edge2Ref string `json:"edge2Ref,omitempty"`
	Offset1  string `json:"offset1,omitempty"`
	Offset2  string `json:"offset2,omitempty"`
	// PointRef and AxisRef locate a "point" placement: the work point to drill at and the work axis
	// to drill along (e.g. "point/0", "origin/axis/z").
	PointRef string `json:"pointRef,omitempty"`
	AxisRef  string `json:"axisRef,omitempty"`
	// Termination is where the bore STOPS (#1863): "distance" (default — Depth from the placement
	// face), "through-all", "to-face" (down to ToFace) or "from-to" (between FromFace and ToFace).
	// A named terminator must be square to the drill axis, since a bore bottoms at one depth.
	Termination string `json:"termination,omitempty"`
	// ToFace/FromFace and their geometric selectors name the terminators, exactly as an extrude's
	// extent does. See [Extrude.ToFace].
	ToFace       string       `json:"toFace,omitempty"`
	ToFaceGeom   *GeomFaceSel `json:"toFaceGeom,omitempty"`
	FromFace     string       `json:"fromFace,omitempty"`
	FromFaceGeom *GeomFaceSel `json:"fromFaceGeom,omitempty"`
}

// HoleClearance names the fastener a clearance hole is drilled for — Inventor's HoleClearanceInfo.
// The host sizes the bore from the published table every recompute, so changing the fastener
// resizes the hole; recording the resolved diameter instead would break that link.
type HoleClearance struct {
	// Standard is the table the fastener is drawn from; "ISO 273" is the one carried today.
	Standard string `json:"standard,omitempty"`
	// Fastener is the thread designation the hole must pass, e.g. "M6".
	Fastener string `json:"fastener"`
	// Fit is "close", "medium" (default) or "free".
	Fit string `json:"fit,omitempty"`
}

// Kind reports the feature kind Hole creates.
func (Hole) Kind() string { return KindHole }

// Boss raises a cylindrical boss from a face (KindBoss).
type Boss struct {
	FaceRef  string `json:"faceRef"`
	Diameter string `json:"diameter"`
	Height   string `json:"height"`
}

// Kind reports the feature kind Boss creates.
func (Boss) Kind() string { return KindBoss }

// Thread threads a cylindrical face — cosmetic (display) or a real modeled cut (KindThread).
type Thread struct {
	FaceRef       string `json:"faceRef"`
	Designation   string `json:"designation"`
	Cut           bool   `json:"cut,omitempty"`
	Class         string `json:"class,omitempty"`
	Tapered       bool   `json:"tapered,omitempty"`
	ModelDiameter string `json:"modelDiameter,omitempty"`
	// Length is the threaded run along the cylinder axis (a distance expression), measured from
	// the face's start edge plus Offset. Empty ⇒ the thread runs the full length of the face
	// (Inventor's FullDepth). A double-ended stud threads only b1 at one end and b2 at the other,
	// each a separate Thread with its own Offset+Length on the same face.
	Length string `json:"length,omitempty"`
	// Offset is the distance (expression) from the face's start edge to where the thread begins
	// (Inventor's ThreadOffset). Empty ⇒ 0 (the thread starts at the face's start edge).
	Offset string `json:"offset,omitempty"`
}

// Kind reports the feature kind Thread creates.
func (Thread) Kind() string { return KindThread }

// Grill cuts a ventilation grill bridged by the boundary profile's rib/spar/island structure (KindGrill).
type Grill struct {
	SketchIndex int     `json:"sketchIndex"`
	Boundaries  []int   `json:"boundaries"`
	Draft       float64 `json:"draft"`
}

// Kind reports the feature kind Grill creates.
func (Grill) Kind() string { return KindGrill }

// Mesh places an ASCII STL as mesh reference geometry (KindMesh).
type Mesh struct {
	Path  string `json:"path"`
	Solid bool   `json:"solid"`
}

// Kind reports the feature kind Mesh creates.
func (Mesh) Kind() string { return KindMesh }

// DirectEdit moves/sizes/rotates/deletes picked faces, or scales the body uniformly (KindDirectEdit).
type DirectEdit struct {
	Operation   string    `json:"operation"`
	FaceRefs    []string  `json:"faceRefs,omitempty"`
	Translation []float64 `json:"translation,omitempty"`
	Direction   []float64 `json:"direction,omitempty"`
	Distance    string    `json:"distance,omitempty"`
	AxisPoint   []float64 `json:"axisPoint,omitempty"`
	AxisDir     []float64 `json:"axisDir,omitempty"`
	Angle       string    `json:"angle,omitempty"`
	Scale       float64   `json:"scale,omitempty"`
	Base        []float64 `json:"base,omitempty"`
}

// Kind reports the feature kind DirectEdit creates.
func (DirectEdit) Kind() string { return KindDirectEdit }
