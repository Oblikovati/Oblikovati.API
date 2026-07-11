// SPDX-License-Identifier: Apache-2.0

package featureargs

// The direct-edit / modify feature kinds (#1709): combine, thicken, trim, the face direct edits
// (move/offset/delete/split), simplify, and unwrap. Body inputs are indices (model.tree body
// order); face inputs are reference keys (get_reference_keys).

const (
	KindCombine    = "combine"
	KindThicken    = "thicken"
	KindTrim       = "trim"
	KindMoveFace   = "moveFace"
	KindFaceOffset = "faceOffset"
	KindDeleteFace = "deleteFace"
	KindSplit      = "split"
	KindSimplify   = "simplify"
	KindUnwrap     = "unwrap"
)

// Combine booleans two solid bodies (KindCombine).
type Combine struct {
	TargetIndex int    `json:"targetIndex"`
	ToolIndex   int    `json:"toolIndex"`
	Operation   string `json:"operation"`
}

// Kind reports the feature kind Combine creates.
func (Combine) Kind() string { return KindCombine }

// Thicken thickens a surface body into a solid — or, with Operation surface, offsets it as a
// surface (KindThicken). Direction (positive|negative|symmetric, default positive per Inventor)
// picks the offset side(s); Operation (join|cut|intersect|surface) picks the output; FaceRefs
// thickens a subset (empty = whole body); CreateVerticalSurfaces (default true) closes subset
// boundaries with side walls. AutomaticFaceChain / AutomaticBlending are accepted for parity but,
// since selection is explicit, are not geometrically applied. Approximation is accepted for #331
// parity; the kernel computes the exact offset (#1876).
type Thicken struct {
	Thickness              string   `json:"thickness"`
	Approximation          string   `json:"approximation,omitempty"`
	Direction              string   `json:"direction,omitempty"`
	Operation              string   `json:"operation,omitempty"`
	FaceRefs               []string `json:"faceRefs,omitempty"`
	CreateVerticalSurfaces *bool    `json:"createVerticalSurfaces,omitempty"`
	AutomaticFaceChain     bool     `json:"automaticFaceChain,omitempty"`
	AutomaticBlending      bool     `json:"automaticBlending,omitempty"`
}

// Kind reports the feature kind Thicken creates.
func (Thicken) Kind() string { return KindThicken }

// Trim trims a surface body with a cutting tool, keeping one side (KindTrim). The tool is one of:
// an explicit plane (Origin+Normal), a work plane / planar face (ToolRef), a planar surface body
// (ToolBodyIndex), or a straight sketch line (ToolSketchIndex+ToolLineIndex — the line swept along
// its sketch normal). KeepPositive selects the kept side. Curved tool surfaces / curved sketch
// curves and multi-region selection are phase C (#1880).
type Trim struct {
	Origin          []float64 `json:"origin,omitempty"`
	Normal          []float64 `json:"normal,omitempty"`
	KeepPositive    bool      `json:"keepPositive,omitempty"`
	ToolRef         string    `json:"toolRef,omitempty"`
	ToolBodyIndex   *int      `json:"toolBodyIndex,omitempty"`
	ToolSketchIndex *int      `json:"toolSketchIndex,omitempty"`
	ToolLineIndex   int       `json:"toolLineIndex,omitempty"`
}

// Kind reports the feature kind Trim creates.
func (Trim) Kind() string { return KindTrim }

// MoveFace moves picked faces by a translation, or rotates them about an axis (#331) (KindMoveFace).
type MoveFace struct {
	FaceRefs    []string  `json:"faceRefs"`
	Translation []float64 `json:"translation,omitempty"`
	AxisPoint   []float64 `json:"axisPoint,omitempty"`
	AxisDir     []float64 `json:"axisDir,omitempty"`
	Angle       string    `json:"angle,omitempty"`
}

// Kind reports the feature kind MoveFace creates.
func (MoveFace) Kind() string { return KindMoveFace }

// FaceOffset offsets picked faces by a distance (KindFaceOffset). Approximation is accepted for
// #331 parity; the kernel computes the exact offset.
type FaceOffset struct {
	FaceRefs      []string `json:"faceRefs"`
	Distance      string   `json:"distance,omitempty"`
	Approximation string   `json:"approximation,omitempty"`
}

// Kind reports the feature kind FaceOffset creates.
func (FaceOffset) Kind() string { return KindFaceOffset }

// DeleteFace deletes picked faces (KindDeleteFace). Heal (default false, matching Inventor's
// DeleteFaceFeatures.Add) extends the neighbouring faces to close the opening; when false the
// faces are removed leaving an open surface body. Selecting the faces of an internal void shell
// instead removes that void and restores mass (the FaceShell arm) (#1884).
type DeleteFace struct {
	FaceRefs []string `json:"faceRefs"`
	Heal     bool     `json:"heal,omitempty"`
}

// Kind reports the feature kind DeleteFace creates.
func (DeleteFace) Kind() string { return KindDeleteFace }

// Split splits picked faces along their intersections (KindSplit).
type Split struct {
	FaceRefs []string `json:"faceRefs"`
}

// Kind reports the feature kind Split creates.
func (Split) Kind() string { return KindSplit }

// Simplify removes and heals selected faces and/or fills internal voids (KindSimplify).
type Simplify struct {
	FaceRefs  []string `json:"faceRefs,omitempty"`
	FillVoids bool     `json:"fillVoids,omitempty"`
}

// Kind reports the feature kind Simplify creates.
func (Simplify) Kind() string { return KindSimplify }

// Unwrap flattens a cylindrical face into a flat sheet patch (KindUnwrap).
type Unwrap struct {
	FaceRef string `json:"faceRef"`
}

// Kind reports the feature kind Unwrap creates.
func (Unwrap) Kind() string { return KindUnwrap }

// modifyArgs is the modify family's contribution to [All].
var modifyArgs = []Arg{
	Combine{}, Thicken{}, Trim{}, MoveFace{}, FaceOffset{},
	DeleteFace{}, Split{}, Simplify{}, Unwrap{},
}
