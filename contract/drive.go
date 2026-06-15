// SPDX-License-Identifier: Apache-2.0

package contract

import "oblikovati.org/api/types"

// The scalar read surface of a drive (M12-F03, Oblikovati/Oblikovati#366): the settings that
// describe sweeping a constraint's or joint's driven variable through a range, re-solving at
// each step to animate the assembly. An in-proc consumer reads the settings directly; the
// drive itself is computed over api/wire (assemblyDrive.*). The host implementation lives in
// /source (model/assembly).

// DriveSettings describes one drive: which variable to sweep, the value range and step, how
// many times to repeat, whether to ping-pong (end→start as well as start→end), and whether a
// collision halts the sweep. Values are in the variable's units (radians for angular,
// centimetres for linear).
type DriveSettings interface {
	// Variable selects which driven variable the drive sweeps (natural/angular/linear).
	Variable() types.DriveVariable
	// Range returns the start and end values of the sweep (variable units).
	Range() (start, end float64)
	// Step returns the increment between consecutive frames (variable units, always positive).
	Step() float64
	// RepetitionCount returns how many times the sweep repeats (1 = a single pass).
	RepetitionCount() int
	// RepetitionStartEndStart reports whether each repetition also plays back end→start
	// (a ping-pong), versus restarting at the start value.
	RepetitionStartEndStart() bool
	// CollisionDetection reports whether the drive halts when the moved component interferes
	// with another.
	CollisionDetection() bool
}
