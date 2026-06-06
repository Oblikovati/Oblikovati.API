// SPDX-License-Identifier: Apache-2.0

package client

import (
	"oblikovati/api/types"
	"oblikovati/api/wire"
)

// Pattern is the sketch-pattern group, reached via [Sketch.Pattern].
type Pattern struct {
	c     *Client
	index int
}

// Pattern returns the sketch-pattern group for the sketch at index.
func (s Sketch) Pattern(index int) Pattern { return Pattern{s.c, index} }

// Rectangular duplicates the seed entities on a count1×count2 grid stepped by the
// unit-bearing spacings along the default axes; returns the created copy ids.
func (p Pattern) Rectangular(entities []uint64, count1, count2 int, spacing1, spacing2 string) (wire.AddSketchPatternResult, error) {
	var r wire.AddSketchPatternResult
	args := wire.AddSketchPatternArgs{
		SketchIndex: p.index, Kind: string(types.SketchPatternRectangular), Entities: entities,
		Count1: count1, Count2: count2, Spacing1: spacing1, Spacing2: spacing2,
	}
	return r, p.c.call(wire.MethodSketchAddPattern, args, &r)
}

// Circular duplicates the seed entities into count instances spread over the unit-bearing
// angle about center ([x,y] cm); returns the created copy ids.
func (p Pattern) Circular(entities []uint64, center []float64, count int, angle string) (wire.AddSketchPatternResult, error) {
	var r wire.AddSketchPatternResult
	args := wire.AddSketchPatternArgs{
		SketchIndex: p.index, Kind: string(types.SketchPatternCircular), Entities: entities,
		Count: count, Angle: angle, Center: center,
	}
	return r, p.c.call(wire.MethodSketchAddPattern, args, &r)
}
