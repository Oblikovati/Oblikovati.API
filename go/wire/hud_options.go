// SPDX-License-Identifier: Apache-2.0

package wire

// HeadsUpDisplayOptionsView is the in-canvas sketch input configuration: the pointer-input boxes
// that place a shape's first point, the dimension-input boxes that size it, and whether a typed
// value becomes a persistent driving dimension (Oblikovati/Oblikovati#2014).
//
// It is the result of [MethodApplicationGetHUDOptions] and the request of
// [MethodApplicationSetHUDOptions]; a set replaces every field, so read-modify-write to change
// one.
type HeadsUpDisplayOptionsView struct {
	Enabled                              bool `json:"enabled"`
	PointerInputEnabled                  bool `json:"pointerInputEnabled"`
	PointerInputInCartesianCoordinates   bool `json:"pointerInputInCartesianCoordinates"`
	DimensionInputEnabled                bool `json:"dimensionInputEnabled"`
	DimensionInputInCartesianCoordinates bool `json:"dimensionInputInCartesianCoordinates"`
	CreateDimensionsOnValueInput         bool `json:"createDimensionsOnValueInput"`
}
