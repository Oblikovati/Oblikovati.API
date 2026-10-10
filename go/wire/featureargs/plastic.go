// SPDX-License-Identifier: Apache-2.0

package featureargs

// The plastic-part feature kinds (#1709, M20-F10): a cantilever snap-fit hook and a raised/recessed
// rest pad bounded by a closed sketch profile.

const (
	KindSnapFit = "snapFit"
	KindRest    = "rest"
)

// SnapFit adds a cantilever snap-fit hook (a beam with a catch lip) to the part (KindSnapFit).
type SnapFit struct {
	Length      string `json:"length"`
	Width       string `json:"width"`
	Thickness   string `json:"thickness"`
	CatchLength string `json:"catchLength"`
	CatchHeight string `json:"catchHeight"`
}

// Kind reports the feature kind SnapFit creates.
func (SnapFit) Kind() string { return KindSnapFit }

// Rest adds a raised (or recessed) rest pad bounded by a closed sketch profile (KindRest).
type Rest struct {
	SketchIndex    int    `json:"sketchIndex"`
	ProfileIndices []int  `json:"profileIndices,omitempty"`
	ProfileIndex   int    `json:"profileIndex"`
	Depth          string `json:"depth"`
	Recessed       bool   `json:"recessed,omitempty"`
}

// Kind reports the feature kind Rest creates.
func (Rest) Kind() string { return KindRest }

// plasticArgs is the plastic-features family's contribution to [All].
var plasticArgs = []Arg{SnapFit{}, Rest{}}
