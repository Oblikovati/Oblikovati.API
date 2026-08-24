// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// Appearance is the in-process contract for one PBR appearance — what the renderer shows
// for a surface. It is a metallic-roughness description with solid (non-textured) values;
// the GPL model/material.Appearance satisfies it (compile-time asserted there).
//
// Albedo and Emissive are [types.Rgba] (shared with theming); Metallic, Roughness and
// Opacity are in [0,1].
//
// Deprecated: use [OpenPBRAppearance] for new work — it covers the full OpenPBR Surface
// v1.1.1 lobe set (M45, ADR-0053) that this metallic-roughness subset cannot express
// (coat, fuzz, subsurface, thin-film, dispersion). Appearance keeps working: it is not
// removed by this deprecation, and removal is a separate future MAJOR-version decision.
type Appearance interface {
	// ID is the stable identity used by assignments and library lookups.
	ID() string
	// DisplayName is the label shown in the appearance browser.
	DisplayName() string
	// Source says whether this is a built-in, project, or document-embedded asset.
	Source() types.AssetSource
	// Albedo is the base (diffuse) color.
	Albedo() types.Rgba
	// Metallic is the metalness in [0,1].
	Metallic() float32
	// Roughness is the microfacet roughness in [0,1].
	Roughness() float32
	// Emissive is the self-emitted color (black = none).
	Emissive() types.Rgba
	// Opacity is the surface opacity in [0,1] (1 = fully opaque).
	Opacity() float32
}
