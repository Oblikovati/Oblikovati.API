// SPDX-License-Identifier: Apache-2.0

package types

// Color3 is a 3-channel, alpha-free color in the ACEScg working space
// (https://docs.acescentral.com/specifications/acescg/) — OpenPBR Surface's default color
// space for every `color3` parameter (parametrization.md.html). Unlike [Rgba], components
// are not implicitly clamped to [0,1]: OpenPBR's emission_color is unbounded above.
//
// This is the canonical, Apache-2.0 definition; the GPL model aliases it
// (model/material.Color3 = types.Color3).
type Color3 struct {
	R, G, B float32
}

// NewColor3 builds a Color3 from its ACEScg channels.
func NewColor3(r, g, b float32) Color3 { return Color3{R: r, G: g, B: b} }

// Array returns the color as the [3]float32 the renderer's ACEScg working-space buffers
// consume.
func (c Color3) Array() [3]float32 { return [3]float32{c.R, c.G, c.B} }
