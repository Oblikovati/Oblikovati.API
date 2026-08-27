// SPDX-License-Identifier: Apache-2.0

package types

import "fmt"

// ColorSourceTypeEnum is where an object's color comes from: an explicit per-object
// override, an automatic (by-object-class) color, the object's layer, or its sheet. The
// numeric ids are stable, frozen values (79105–79108).
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.ColorSourceTypeEnum).
type ColorSourceTypeEnum int32

const (
	// OverrideColorSource: the color is an explicit per-object override (79105).
	OverrideColorSource ColorSourceTypeEnum = 79105
	// AutomaticColorSource: the color is the automatic (by-class) color (79106).
	AutomaticColorSource ColorSourceTypeEnum = 79106
	// LayerColorSource: the color is inherited from the object's layer (79107).
	LayerColorSource ColorSourceTypeEnum = 79107
	// SheetColorSource: the color is inherited from the object's sheet (79108).
	SheetColorSource ColorSourceTypeEnum = 79108
)

var colorSourceNames = map[ColorSourceTypeEnum]string{
	OverrideColorSource:  "Override",
	AutomaticColorSource: "Automatic",
	LayerColorSource:     "Layer",
	SheetColorSource:     "Sheet",
}

// String returns the color-source's user-facing name.
func (c ColorSourceTypeEnum) String() string {
	return enumName(colorSourceNames, c, "colorSource(?)")
}

// IsValid reports whether c is a defined color source.
func (c ColorSourceTypeEnum) IsValid() bool {
	return enumValid(colorSourceNames, c)
}

// AllColorSources returns every defined color source, in picker order.
func AllColorSources() []ColorSourceTypeEnum {
	return []ColorSourceTypeEnum{OverrideColorSource, AutomaticColorSource, LayerColorSource, SheetColorSource}
}

// Color is the universal color value object: the currency of the whole visual surface
// (render styles, lights, highlight sets, client-graphics color sets, display settings all
// traffic in it). It is plain data — 8-bit R/G/B components, a 0–1 Opacity (0 fully
// transparent, 1 fully opaque), and a Source discriminating an override from an inherited
// color. It mirrors the reference Color object's Red/Green/Blue/Opacity/ColorSourceType.
//
// This is the canonical Apache-2.0 definition; the GPL implementation aliases it
// (app.Color = types.Color) so the head, renderer, and model share one color type.
//
//	c := types.NewColor(255, 0, 0)            // opaque red, override source
//	c.Opacity = 0.5                           // half-transparent
type Color struct {
	R       uint8               `json:"r"`
	G       uint8               `json:"g"`
	B       uint8               `json:"b"`
	Opacity float64             `json:"opacity"`
	Source  ColorSourceTypeEnum `json:"source"`
}

// NewColor builds an opaque override Color from 8-bit R/G/B components — the common case.
func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b, Opacity: 1, Source: OverrideColorSource}
}

// IsOverride reports whether the colour is an explicit per-object override rather than an
// inherited one (automatic, layer or sheet). Callers storing optional colour overrides use it as
// the "is this set?" test, so the meaning of automatic lives in one place — and so the zero Color,
// whose Source is 0 and therefore not a member of the enum, never reads as set.
func (c Color) IsOverride() bool { return c.Source == OverrideColorSource }

// Rgba converts to the renderer's [Rgba] (float32 channels in [0,1], Opacity → alpha), so a
// Color crosses to the draw path without a bespoke conversion at every call site.
func (c Color) Rgba() Rgba {
	return Rgba{R: float32(c.R) / 255, G: float32(c.G) / 255, B: float32(c.B) / 255, A: float32(c.Opacity)}
}

// Hex formats the color as "#RRGGBB" (lower-case), ignoring opacity — the form a scheme
// file stores for an opaque swatch.
func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
