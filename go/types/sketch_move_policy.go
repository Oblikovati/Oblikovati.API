// SPDX-License-Identifier: Apache-2.0

package types

// SketchMoveDimensionPolicy controls whether a planned move may change dimensions.
// A plan request must choose a supported nonempty value; there is no wire default.
type SketchMoveDimensionPolicy string

const (
	// SketchMoveDimensionNever keeps all dimensional constraints unchanged.
	SketchMoveDimensionNever SketchMoveDimensionPolicy = "never"
	// SketchMoveDimensionIfNoEquation permits eligible dimensions to change only when
	// they do not depend on another dimension; unknown dependencies are refused.
	SketchMoveDimensionIfNoEquation SketchMoveDimensionPolicy = "ifNoEquation"
	// SketchMoveDimensionAlways permits dimensions between selected and unselected
	// geometry to be recomputed. Dimensions entirely inside the selection stay fixed.
	SketchMoveDimensionAlways SketchMoveDimensionPolicy = "always"
	// SketchMoveDimensionPrompt requires explicit approval for proposed dimension changes.
	SketchMoveDimensionPrompt SketchMoveDimensionPolicy = "prompt"
)

// SketchMoveGeometricPolicy controls whether a planned move may remove relations.
// A plan request must choose a supported nonempty value; there is no wire default.
type SketchMoveGeometricPolicy string

const (
	// SketchMoveGeometricNever keeps all geometric relations unchanged.
	SketchMoveGeometricNever SketchMoveGeometricPolicy = "never"
	// SketchMoveGeometricAlways permits removal of fixed constraints on selected
	// geometry and relations across the selection boundary, except parallel,
	// perpendicular and equal relations. Relations within the selection remain.
	SketchMoveGeometricAlways SketchMoveGeometricPolicy = "always"
	// SketchMoveGeometricPrompt requires explicit approval for proposed relation removals.
	SketchMoveGeometricPrompt SketchMoveGeometricPolicy = "prompt"
)
