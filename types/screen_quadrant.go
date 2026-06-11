// SPDX-License-Identifier: Apache-2.0

package types

// ScreenQuadrant is one slot of a radial marking menu — the ScreenQuadrantEnum
// equivalent (M05-F12, #619): eight directions around the cursor, clockwise from
// north.
type ScreenQuadrant uint8

const (
	QuadrantNorth     ScreenQuadrant = 0
	QuadrantNorthEast ScreenQuadrant = 1
	QuadrantEast      ScreenQuadrant = 2
	QuadrantSouthEast ScreenQuadrant = 3
	QuadrantSouth     ScreenQuadrant = 4
	QuadrantSouthWest ScreenQuadrant = 5
	QuadrantWest      ScreenQuadrant = 6
	QuadrantNorthWest ScreenQuadrant = 7
)

var screenQuadrantNames = map[ScreenQuadrant]string{
	QuadrantNorth: "n", QuadrantNorthEast: "ne", QuadrantEast: "e", QuadrantSouthEast: "se",
	QuadrantSouth: "s", QuadrantSouthWest: "sw", QuadrantWest: "w", QuadrantNorthWest: "nw",
}

// String returns the quadrant's stable compass name.
func (q ScreenQuadrant) String() string {
	if name, ok := screenQuadrantNames[q]; ok {
		return name
	}
	return "screenQuadrant(?)"
}
