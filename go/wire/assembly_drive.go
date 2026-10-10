// SPDX-License-Identifier: Apache-2.0

package wire

import "oblikovati.org/api/types"

// The assembly drive surface (M12-F03, Oblikovati/Oblikovati#366): drive a joint's driven
// variable through a value range and get back the per-step occurrence placements — the
// frames of a kinematic motion study. Each step re-solves the assembly (constraints and
// joints together) with the driven variable pinned to the step value; with collision
// detection the sweep halts at the first frame where the moved component interferes.

// DriveSettingsDTO describes a sweep: which variable, the value range and (positive) step,
// repetitions, ping-pong, and the collision-stop flag. Units follow the variable — radians
// for angular, centimetres for linear.
type DriveSettingsDTO struct {
	Variable                string  `json:"variable,omitempty"` // natural (default) | angular | linear
	Start                   float64 `json:"start"`
	End                     float64 `json:"end"`
	Step                    float64 `json:"step"`
	RepetitionCount         int     `json:"repetitionCount,omitempty"`
	RepetitionStartEndStart bool    `json:"repetitionStartEndStart,omitempty"`
	CollisionDetection      bool    `json:"collisionDetection,omitempty"`
}

// DrivePlacement is one occurrence's resulting transform at a drive frame (row-major, in the
// assembly's space).
type DrivePlacement struct {
	Occurrence uint64       `json:"occurrence"`
	Transform  types.Matrix `json:"transform"`
}

// DriveFrame is one step of a drive: the driven value and the occurrence placements the
// assembly solved to there. Collided marks the frame at which interference was detected.
type DriveFrame struct {
	Value      float64          `json:"value"`
	Placements []DrivePlacement `json:"placements"`
	Collided   bool             `json:"collided,omitempty"`
}

// DriveJointArgs is the request of [MethodAssemblyDrivePreview]: sweep the joint with the
// given id through Settings and return the resulting frames.
type DriveJointArgs struct {
	Joint    uint64           `json:"joint"`
	Settings DriveSettingsDTO `json:"settings"`
}

// DriveResult is the response of a drive preview: the frames in play order, and — when
// collision detection halted the sweep — the flag and the index of the last (collided) frame.
type DriveResult struct {
	Frames             []DriveFrame `json:"frames"`
	StoppedByCollision bool         `json:"stoppedByCollision,omitempty"`
	StoppedAtStep      int          `json:"stoppedAtStep,omitempty"`
}
