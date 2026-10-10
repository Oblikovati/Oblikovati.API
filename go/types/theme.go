// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"strconv"
	"strings"
)

// Rgba is a straight-alpha color with each channel in [0,1] — the shape Dear ImGui
// (ImVec4) and the head's draw overlays ([4]float32) both consume, so a theme color
// crosses to the renderer without conversion. It is plain data; helpers parse/format
// the "#RRGGBB[AA]" hex used in theme files and on the wire.
//
// This is the canonical, Apache-2.0 definition; the GPL theme package aliases it
// (theme.Rgba = types.Rgba) so the head and renderer share one color type.
type Rgba struct {
	R, G, B, A float32
}

// Array returns the color as the [4]float32 the head passes to ImGui and the renderer.
//
//	tint := theme.Color(types.TokenIconPrimary).Array() // [4]float32 for ImageButton
func (c Rgba) Array() [4]float32 { return [4]float32{c.R, c.G, c.B, c.A} }

// Hex formats the color as "#RRGGBBAA" (lower-case), the form stored in theme files.
func (c Rgba) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x%02x", to255(c.R), to255(c.G), to255(c.B), to255(c.A))
}

// ParseHex reads "#RRGGBB" (alpha defaults to opaque) or "#RRGGBBAA" into an Rgba. It
// errors, naming the offending string, on a missing "#", a bad length, or non-hex
// digits — so a hand-edited theme file fails loudly rather than rendering black.
func ParseHex(s string) (Rgba, error) {
	body, ok := strings.CutPrefix(s, "#")
	if !ok || (len(body) != 6 && len(body) != 8) {
		return Rgba{}, fmt.Errorf("types: bad hex color %q, want \"#RRGGBB\" or \"#RRGGBBAA\"", s)
	}
	a := byte(255)
	if len(body) == 8 {
		v, err := parseByte(body[6:8], s)
		if err != nil {
			return Rgba{}, err
		}
		a = v
	}
	return parseRGB(body, a, s)
}

// parseRGB reads the first three byte pairs of body into an opaque-or-given-alpha Rgba.
func parseRGB(body string, a byte, src string) (Rgba, error) {
	r, err := parseByte(body[0:2], src)
	if err != nil {
		return Rgba{}, err
	}
	g, err := parseByte(body[2:4], src)
	if err != nil {
		return Rgba{}, err
	}
	b, err := parseByte(body[4:6], src)
	if err != nil {
		return Rgba{}, err
	}
	return Rgba{from255(r), from255(g), from255(b), from255(a)}, nil
}

// parseByte parses one 2-digit hex pair, naming the whole color on failure.
func parseByte(pair, src string) (byte, error) {
	v, err := strconv.ParseUint(pair, 16, 8)
	if err != nil {
		return 0, fmt.Errorf("types: bad hex color %q: non-hex digits %q", src, pair)
	}
	return byte(v), nil
}

// to255 quantizes a [0,1] channel to a byte, clamping out-of-range inputs.
func to255(v float32) byte {
	switch {
	case v <= 0:
		return 0
	case v >= 1:
		return 255
	default:
		return byte(v*255 + 0.5)
	}
}

// from255 expands a byte channel to [0,1].
func from255(v byte) float32 { return float32(v) / 255 }
