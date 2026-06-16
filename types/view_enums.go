// SPDX-License-Identifier: Apache-2.0

package types

// ViewTypeEnum is the kind of view in a document's view collection: a graphics (3D) view, the
// model-browser pane, or a notebook view. Frozen ids 9217–9220.
type ViewTypeEnum int32

const (
	// UnknownViewType is an unclassified view (9217).
	UnknownViewType ViewTypeEnum = 9217
	// GraphicsViewType is a 3D graphics view (9218).
	GraphicsViewType ViewTypeEnum = 9218
	// BrowserViewType is the model-browser pane (9219).
	BrowserViewType ViewTypeEnum = 9219
	// NoteBookViewType is a notebook view (9220).
	NoteBookViewType ViewTypeEnum = 9220
)

var viewTypeNames = map[ViewTypeEnum]string{
	UnknownViewType: "Unknown", GraphicsViewType: "Graphics",
	BrowserViewType: "Browser", NoteBookViewType: "Notebook",
}

// String returns the view-type's user-facing name.
func (v ViewTypeEnum) String() string {
	if name, ok := viewTypeNames[v]; ok {
		return name
	}
	return "viewType(?)"
}

// IsValid reports whether v is a defined view type.
func (v ViewTypeEnum) IsValid() bool { _, ok := viewTypeNames[v]; return ok }

// AllViewTypes returns every defined view type.
func AllViewTypes() []ViewTypeEnum {
	return []ViewTypeEnum{GraphicsViewType, BrowserViewType, NoteBookViewType, UnknownViewType}
}

// ViewOperationTypeEnum is an interactive view manipulation: rotate (orbit), pan, or zoom.
// Frozen ids 30209–30211.
type ViewOperationTypeEnum int32

const (
	// RotateViewOperation orbits the camera (30209).
	RotateViewOperation ViewOperationTypeEnum = 30209
	// PanViewOperation pans the camera (30210).
	PanViewOperation ViewOperationTypeEnum = 30210
	// ZoomViewOperation zooms the camera (30211).
	ZoomViewOperation ViewOperationTypeEnum = 30211
)

var viewOperationNames = map[ViewOperationTypeEnum]string{
	RotateViewOperation: "Rotate", PanViewOperation: "Pan", ZoomViewOperation: "Zoom",
}

// String returns the view-operation's user-facing name.
func (v ViewOperationTypeEnum) String() string {
	if name, ok := viewOperationNames[v]; ok {
		return name
	}
	return "viewOperation(?)"
}

// IsValid reports whether v is a defined view operation.
func (v ViewOperationTypeEnum) IsValid() bool { _, ok := viewOperationNames[v]; return ok }

// AllViewOperations returns every defined view operation.
func AllViewOperations() []ViewOperationTypeEnum {
	return []ViewOperationTypeEnum{RotateViewOperation, PanViewOperation, ZoomViewOperation}
}

// OrbitTypeEnum is how the orbit gesture is constrained: a free (trackball) orbit or a
// constrained (turntable) orbit about vertical. Frozen ids 86017–86018.
type OrbitTypeEnum int32

const (
	// FreeOrbit is an unconstrained trackball orbit (86017).
	FreeOrbit OrbitTypeEnum = 86017
	// ConstrainedOrbit is a turntable orbit constrained about vertical (86018).
	ConstrainedOrbit OrbitTypeEnum = 86018
)

var orbitTypeNames = map[OrbitTypeEnum]string{
	FreeOrbit: "Free", ConstrainedOrbit: "Constrained",
}

// String returns the orbit-type's user-facing name.
func (o OrbitTypeEnum) String() string {
	if name, ok := orbitTypeNames[o]; ok {
		return name
	}
	return "orbitType(?)"
}

// IsValid reports whether o is a defined orbit type.
func (o OrbitTypeEnum) IsValid() bool { _, ok := orbitTypeNames[o]; return ok }

// AllOrbitTypes returns every defined orbit type.
func AllOrbitTypes() []OrbitTypeEnum { return []OrbitTypeEnum{FreeOrbit, ConstrainedOrbit} }

// ViewTileTypeEnum is how a document's views are tiled when more than one is shown: a single
// arrange, or a horizontal/vertical split. Frozen ids 117761–117764. (This is the typed
// counterpart of the numeric [ViewLayout] used by the views collection.)
type ViewTileTypeEnum int32

const (
	// UnknownViewTileType is an unclassified tiling (117761).
	UnknownViewTileType ViewTileTypeEnum = 117761
	// ArrangeViewTileType arranges the views automatically (117762).
	ArrangeViewTileType ViewTileTypeEnum = 117762
	// HorizontalViewTileType splits the views horizontally (117763).
	HorizontalViewTileType ViewTileTypeEnum = 117763
	// VerticalViewTileType splits the views vertically (117764).
	VerticalViewTileType ViewTileTypeEnum = 117764
)

var viewTileNames = map[ViewTileTypeEnum]string{
	UnknownViewTileType: "Unknown", ArrangeViewTileType: "Arrange",
	HorizontalViewTileType: "Horizontal", VerticalViewTileType: "Vertical",
}

// String returns the view-tile-type's user-facing name.
func (v ViewTileTypeEnum) String() string {
	if name, ok := viewTileNames[v]; ok {
		return name
	}
	return "viewTileType(?)"
}

// IsValid reports whether v is a defined view-tile type.
func (v ViewTileTypeEnum) IsValid() bool { _, ok := viewTileNames[v]; return ok }

// AllViewTileTypes returns every defined view-tile type.
func AllViewTileTypes() []ViewTileTypeEnum {
	return []ViewTileTypeEnum{ArrangeViewTileType, HorizontalViewTileType, VerticalViewTileType, UnknownViewTileType}
}
