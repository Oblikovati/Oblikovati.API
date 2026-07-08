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
	Extent         string `json:"extent,omitempty"`         // distance|through-all|to-next|to-face (default distance)
	Direction      string `json:"direction,omitempty"`      // positive|negative|symmetric (default positive)
	SecondDistance string `json:"secondDistance,omitempty"` // asymmetric two-direction depth
	Taper          string `json:"taper,omitempty"`          // draft angle, e.g. "3 deg"
	ToFace         string `json:"toFace,omitempty"`         // to-face target: a planar face key, "plane/N", or "origin/plane/xy"
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
	Angle        string `json:"angle"`
	Angle2       string `json:"angle2,omitempty"` // second-direction sweep (#313)
	Operation    string `json:"operation,omitempty"`
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
