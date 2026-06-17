// SPDX-License-Identifier: Apache-2.0

package types

// UnitsType is the dimension category of a measured value (parity:
// ValueUnitsTypeEnum) — the kind of quantity a unit name belongs to, as opposed
// to the unit name itself ("mm", "in"). It is carried on the wire by its
// spelling (UnitsType.String()); the numeric ids are stable and must never be
// renumbered.
//
// This is the canonical, Apache-2.0 definition. The host's parameter engine has
// its own richer internal category type (which also models non-arithmetic flag/
// text values); it maps to and from these spellings at the wire boundary.
type UnitsType int32

const (
	UnitsUnitless UnitsType = 0
	UnitsLength   UnitsType = 1
	UnitsAngle    UnitsType = 2
	UnitsArea     UnitsType = 3
	UnitsVolume   UnitsType = 4
	UnitsMass     UnitsType = 5
	UnitsTime     UnitsType = 6
)

// unitsTypeNames are the wire spellings, identical to the host category
// spellings so a value crosses the boundary without translation tables.
var unitsTypeNames = map[UnitsType]string{
	UnitsUnitless: "unitless",
	UnitsLength:   "length",
	UnitsAngle:    "angle",
	UnitsArea:     "area",
	UnitsVolume:   "volume",
	UnitsMass:     "mass",
	UnitsTime:     "time",
}

// String returns the category's wire spelling (e.g. "length").
func (u UnitsType) String() string { return enumName(unitsTypeNames, u) }

// ParseUnitsType resolves a wire spelling back to its UnitsType.
func ParseUnitsType(s string) (UnitsType, bool) { return enumFromName(unitsTypeNames, s) }
