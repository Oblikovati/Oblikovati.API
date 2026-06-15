// SPDX-License-Identifier: Apache-2.0

package client

import "oblikovati.org/api/wire"

// AssemblyDrive is the assembly drive operation group (M12-F03, Oblikovati/Oblikovati#366):
// sweep a joint's driven variable through a range and read back the per-step occurrence
// placements — the frames of a kinematic motion study. Each step re-solves the assembly with
// the driven variable pinned; with collision detection the sweep halts at the first
// interfering frame.
type AssemblyDrive struct{ c *Client }

// AssemblyDrive returns the assembly drive operation group.
func (c *Client) AssemblyDrive() AssemblyDrive { return AssemblyDrive{c} }

// Preview drives a joint through the given settings and returns the resulting frames, e.g.
// Preview(wire.DriveJointArgs{Joint: id, Settings: wire.DriveSettingsDTO{Start: 0, End: math.Pi, Step: math.Pi / 18}}).
//
// mcp:tool drive_joint
// mcp:summary Drive a joint's variable through a range (settings: start, end, step in radians for angular / cm for linear; optional variable angular|linear, repetitionCount, repetitionStartEndStart ping-pong, collisionDetection). Re-solves each step; returns the frames (driven value + occurrence transforms), halting at the first interfering frame when collision detection is on.
func (a AssemblyDrive) Preview(args wire.DriveJointArgs) (wire.DriveResult, error) {
	var r wire.DriveResult
	return r, a.c.call(wire.MethodAssemblyDrivePreview, args, &r)
}
