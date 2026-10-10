// SPDX-License-Identifier: Apache-2.0

package wire

// Helical curve definition (M06-F09, Oblikovati/Oblikovati#624): beyond the
// constant-shape modes shipped with M22-F04, a helix can carry a variable
// shape — a row table where pitch/diameter change per height/revolution
// station — plus start/end transition conditions.

// HelixShapeRow is one station of a variable-shape helix. Diameter and Pitch
// are unit-bearing lengths; exactly one of Height (unit-bearing length from
// the helix base) or Revolution (turn count from the start) positions the
// station, matching the definition's shape kind. Empty fields interpolate
// between neighboring rows.
type HelixShapeRow struct {
	Diameter   string  `json:"diameter,omitempty"`
	Pitch      string  `json:"pitch,omitempty"`
	Height     string  `json:"height,omitempty"`
	Revolution float64 `json:"revolution,omitempty"`
}

// HelixEndCondition is the treatment of one helix end. Kind is the
// [oblikovati.org/api/types.HelixEndKind] wire spelling (empty ⇒ "natural");
// TransitionAngle and FlatAngle (unit-bearing angles) shape a flat end: the
// transition sweep blends from the true helix pitch down to flat, then the
// flat sweep continues at zero pitch.
type HelixEndCondition struct {
	Kind            string `json:"kind,omitempty"`
	TransitionAngle string `json:"transitionAngle,omitempty"`
	FlatAngle       string `json:"flatAngle,omitempty"`
}

// EditHelixArgs is the request of [MethodSketch3DEditHelix]: redefine an
// existing helical curve in place. Entity is the helix's session id. Mode,
// Pitch, Height, Revolutions, Taper and Clockwise have the same meaning as in
// [AddSketch3DEntityArgs] and replace the constant-shape definition; a
// non-empty Rows replaces the definition with a variable shape instead.
// Start/End set the end conditions (nil keeps the current ones). The curve is
// regenerated and dependents resolve against the new shape.
type EditHelixArgs struct {
	SketchIndex int    `json:"sketchIndex"`
	Entity      uint64 `json:"entity"`

	Mode        string  `json:"mode,omitempty"`
	Pitch       string  `json:"pitch,omitempty"`
	Height      string  `json:"height,omitempty"`
	Revolutions float64 `json:"revolutions,omitempty"`
	Taper       string  `json:"taper,omitempty"`
	Clockwise   bool    `json:"clockwise,omitempty"`

	Rows  []HelixShapeRow    `json:"rows,omitempty"`
	Start *HelixEndCondition `json:"start,omitempty"`
	End   *HelixEndCondition `json:"end,omitempty"`
}

// HelixDefinitionView is the response of [MethodSketch3DEditHelix]: the
// definition as stored. ShapeKind is the
// [oblikovati.org/api/types.HelicalShapeDefinitionKind] wire spelling;
// Variable reports a row-table definition (Rows then carries the resolved
// stations with both height and revolution filled in, in database units —
// diameter/pitch in cm). Start/End are always present with resolved angles in
// radians.
type HelixDefinitionView struct {
	ShapeKind   string             `json:"shapeKind"`
	Variable    bool               `json:"variable"`
	Pitch       float64            `json:"pitch,omitempty"`
	Height      float64            `json:"height,omitempty"`
	Revolutions float64            `json:"revolutions,omitempty"`
	Taper       float64            `json:"taper,omitempty"`
	Clockwise   bool               `json:"clockwise,omitempty"`
	Rows        []HelixShapeRow    `json:"rows,omitempty"`
	Start       *HelixEndCondition `json:"start,omitempty"`
	End         *HelixEndCondition `json:"end,omitempty"`
}
