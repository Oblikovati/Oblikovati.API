// SPDX-License-Identifier: Apache-2.0

package featureargs

// The freeform (T-spline) primitive feature kinds (#1709): a box, a plane, and a quadball, each a
// subdivision cage that becomes editable freeform geometry. Level is the cage subdivision level
// (0/absent ⇒ 1). The host historically shared one struct for all three; each is its own kind here.

const (
	KindFreeformBox      = "freeformBox"
	KindFreeformPlane    = "freeformPlane"
	KindFreeformQuadBall = "freeformQuadBall"
)

// FreeformBox creates a freeform (T-spline) box primitive (KindFreeformBox).
type FreeformBox struct {
	SizeX string `json:"sizeX,omitempty"`
	SizeY string `json:"sizeY,omitempty"`
	SizeZ string `json:"sizeZ,omitempty"`
	Level int    `json:"level,omitempty"`
}

// Kind reports the feature kind FreeformBox creates.
func (FreeformBox) Kind() string { return KindFreeformBox }

// FreeformPlane creates a freeform (T-spline) plane primitive (KindFreeformPlane).
type FreeformPlane struct {
	SizeX string `json:"sizeX,omitempty"`
	SizeY string `json:"sizeY,omitempty"`
	Level int    `json:"level,omitempty"`
}

// Kind reports the feature kind FreeformPlane creates.
func (FreeformPlane) Kind() string { return KindFreeformPlane }

// FreeformQuadBall creates a freeform (T-spline) quadball (sphere) primitive (KindFreeformQuadBall).
type FreeformQuadBall struct {
	Radius string `json:"radius,omitempty"`
	Level  int    `json:"level,omitempty"`
}

// Kind reports the feature kind FreeformQuadBall creates.
func (FreeformQuadBall) Kind() string { return KindFreeformQuadBall }

// freeformArgs is the freeform family's contribution to [All].
var freeformArgs = []Arg{FreeformBox{}, FreeformPlane{}, FreeformQuadBall{}}
