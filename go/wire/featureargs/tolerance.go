// SPDX-License-Identifier: Apache-2.0

package featureargs

// The model GD&T tolerance carrier kind (#1709, M20-F13 #866): a metadata feature that annotates
// model geometry with feature-control frames and datum labels. It changes no geometry — the
// records survive recompute and the .obk round trip.

const KindModelTolerance = "modelTolerance"

// ToleranceFrame is one feature-control frame: a geometric Characteristic + zone Value applied to
// the geometry named by Geometry (a reference key), optionally referencing Datums.
type ToleranceFrame struct {
	Geometry       string   `json:"geometry"`
	Characteristic string   `json:"characteristic"`
	Value          string   `json:"value"`
	Datums         []string `json:"datums"`
}

// DatumLabel tags model geometry (a reference key) with a datum Label.
type DatumLabel struct {
	Geometry string `json:"geometry"`
	Label    string `json:"label"`
}

// ModelTolerance annotates model geometry with GD&T frames and datums, changing no geometry
// (KindModelTolerance).
type ModelTolerance struct {
	Frames []ToleranceFrame `json:"frames"`
	Datums []DatumLabel     `json:"datums"`
}

// Kind reports the feature kind ModelTolerance creates.
func (ModelTolerance) Kind() string { return KindModelTolerance }

// toleranceArgs is the tolerance family's contribution to [All].
var toleranceArgs = []Arg{ModelTolerance{}}
