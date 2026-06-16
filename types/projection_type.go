// SPDX-License-Identifier: Apache-2.0

package types

// ProjectionTypeEnum is a view's camera projection: orthographic (parallel), perspective, or
// perspective with the faces of the view cube drawn orthographically. The numeric ids are
// stable, frozen values (86273–86275).
//
// First consumed by the display-settings new-window default; the views surface reuses it. This
// is the canonical Apache-2.0 definition; the GPL implementation aliases it.
type ProjectionTypeEnum int32

const (
	// OrthographicProjection is a parallel (non-foreshortened) projection (86273).
	OrthographicProjection ProjectionTypeEnum = 86273
	// PerspectiveProjection is a perspective (foreshortened) projection (86274).
	PerspectiveProjection ProjectionTypeEnum = 86274
	// PerspectiveWithOrthoFacesProjection is perspective with ortho view-cube faces (86275).
	PerspectiveWithOrthoFacesProjection ProjectionTypeEnum = 86275
)

var projectionTypeNames = map[ProjectionTypeEnum]string{
	OrthographicProjection:              "Orthographic",
	PerspectiveProjection:               "Perspective",
	PerspectiveWithOrthoFacesProjection: "Perspective with Ortho Faces",
}

// String returns the projection type's user-facing name.
func (p ProjectionTypeEnum) String() string {
	if name, ok := projectionTypeNames[p]; ok {
		return name
	}
	return "projectionType(?)"
}

// IsValid reports whether p is a defined projection type.
func (p ProjectionTypeEnum) IsValid() bool {
	_, ok := projectionTypeNames[p]
	return ok
}

// AllProjectionTypes returns every defined projection type.
func AllProjectionTypes() []ProjectionTypeEnum {
	return []ProjectionTypeEnum{OrthographicProjection, PerspectiveProjection, PerspectiveWithOrthoFacesProjection}
}
