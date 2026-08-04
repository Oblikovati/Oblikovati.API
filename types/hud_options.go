// SPDX-License-Identifier: Apache-2.0

package types

// HeadsUpDisplayOptions is the application's in-canvas input configuration while sketching: the
// pointer-input boxes that place a shape's first point by coordinate, the dimension-input boxes
// that size the shape being placed, and whether a typed value becomes a persistent dimension
// (Oblikovati/Oblikovati#2014).
//
// Pointer input and dimension input are separate on purpose: they answer different questions
// ("where does this start?" versus "how big is it?") and a user may want one without the other.
type HeadsUpDisplayOptions struct {
	// Enabled switches the whole in-canvas input off, leaving only the command line.
	Enabled bool `json:"enabled"`

	// PointerInputEnabled shows the coordinate boxes for a shape's first point, and
	// PointerInputInCartesianCoordinates shows them as X/Y rather than length and angle.
	PointerInputEnabled                bool `json:"pointerInputEnabled"`
	PointerInputInCartesianCoordinates bool `json:"pointerInputInCartesianCoordinates"`

	// DimensionInputEnabled shows the boxes that size the shape being placed, and
	// DimensionInputInCartesianCoordinates shows them as width/height rather than length and
	// angle where the shape offers both readings.
	DimensionInputEnabled                bool `json:"dimensionInputEnabled"`
	DimensionInputInCartesianCoordinates bool `json:"dimensionInputInCartesianCoordinates"`

	// CreateDimensionsOnValueInput makes a value the user types into a dimension box become a
	// persistent driving dimension when the shape commits. With it off the typed value still
	// sizes the shape, but states nothing afterwards.
	CreateDimensionsOnValueInput bool `json:"createDimensionsOnValueInput"`
}

// DefaultHeadsUpDisplayOptions is the out-of-the-box configuration: everything on, a shape's
// first point entered in Cartesian coordinates, the shape itself sized in polar length and
// angle, and typed values persisted as dimensions.
func DefaultHeadsUpDisplayOptions() HeadsUpDisplayOptions {
	return HeadsUpDisplayOptions{
		Enabled:                              true,
		PointerInputEnabled:                  true,
		PointerInputInCartesianCoordinates:   true,
		DimensionInputEnabled:                true,
		DimensionInputInCartesianCoordinates: false,
		CreateDimensionsOnValueInput:         true,
	}
}
