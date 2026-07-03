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
}

// Kind reports the feature kind SheetMetalFlange creates.
func (SheetMetalFlange) Kind() string { return KindSheetMetalFlange }

// SheetMetalHem folds a hem (closed or open) along an edge (KindSheetMetalHem).
type SheetMetalHem struct {
	Edge   string `json:"edge"`
	Length string `json:"length"`
	Type   string `json:"type,omitempty"`
	Gap    string `json:"gap,omitempty"`
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
// (KindSheetMetalCorner).
type SheetMetalCorner struct {
	Edges     []string `json:"edges"`
	Treatment string   `json:"treatment"`
	Size      string   `json:"size"`
}

// Kind reports the feature kind SheetMetalCorner creates.
func (SheetMetalCorner) Kind() string { return KindSheetMetalCorner }

// SheetMetalContourFlange runs a contour flange (a profile sketch) along an edge
// (KindSheetMetalContourFlange).
type SheetMetalContourFlange struct {
	Edge          string `json:"edge"`
	ProfileSketch int    `json:"profileSketch"`
	Flip          bool   `json:"flip,omitempty"`
}

// Kind reports the feature kind SheetMetalContourFlange creates.
func (SheetMetalContourFlange) Kind() string { return KindSheetMetalContourFlange }

// SheetMetalLoftedFlange lofts a flange between two profile sketches (KindSheetMetalLoftedFlange).
type SheetMetalLoftedFlange struct {
	ProfileA  int    `json:"profileA"`
	ProfileB  int    `json:"profileB"`
	Operation string `json:"operation"`
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

// SheetMetalCornerSeam relieves the seam at picked corner edges by a gap (KindSheetMetalCornerSeam).
type SheetMetalCornerSeam struct {
	Edges []string `json:"edges"`
	Gap   string   `json:"gap"`
	Type  string   `json:"type,omitempty"`
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

// SheetMetalRip rips the wall open along a sketch line by a gap (KindSheetMetalRip).
type SheetMetalRip struct {
	SketchIndex int    `json:"sketchIndex"`
	LineIndex   int    `json:"lineIndex"`
	Gap         string `json:"gap,omitempty"`
}

// Kind reports the feature kind SheetMetalRip creates.
func (SheetMetalRip) Kind() string { return KindSheetMetalRip }

// SheetMetalPunch punches a sketch-driven tool into the wall to a depth (KindSheetMetalPunch).
type SheetMetalPunch struct {
	SketchIndex int    `json:"sketchIndex"`
	Depth       string `json:"depth,omitempty"`
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
