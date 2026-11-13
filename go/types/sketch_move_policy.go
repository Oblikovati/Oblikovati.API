// SPDX-License-Identifier: Apache-2.0

package types

// SketchMoveDimensionPolicy controls changes to driving dimensions during a sketch move.
// The empty value means Never, preserving the behavior of requests without this field.
type SketchMoveDimensionPolicy string

const (
	// SketchMoveDimensionNever keeps all driving dimensions unchanged.
	SketchMoveDimensionNever SketchMoveDimensionPolicy = "never"
	// SketchMoveDimensionIfNoEquation permits changed literal dimensions to follow the move.
	SketchMoveDimensionIfNoEquation SketchMoveDimensionPolicy = "ifNoEquation"
	// SketchMoveDimensionAlways permits changed driving dimensions to follow the move.
	SketchMoveDimensionAlways SketchMoveDimensionPolicy = "always"
	// SketchMoveDimensionPrompt requests a decision if a dimension must change.
	SketchMoveDimensionPrompt SketchMoveDimensionPolicy = "prompt"
)

// SketchMoveGeometricPolicy controls removal of geometric relations during a sketch move.
// The empty value means Never, preserving the behavior of requests without this field.
type SketchMoveGeometricPolicy string

const (
	// SketchMoveGeometricNever keeps all geometric relations.
	SketchMoveGeometricNever SketchMoveGeometricPolicy = "never"
	// SketchMoveGeometricAlways permits incompatible geometric relations to be removed.
	SketchMoveGeometricAlways SketchMoveGeometricPolicy = "always"
	// SketchMoveGeometricPrompt requests a decision if a geometric relation must be removed.
	SketchMoveGeometricPrompt SketchMoveGeometricPolicy = "prompt"
)
