// SPDX-License-Identifier: Apache-2.0

package types

// StartupSketchPlane chooses the initial sketch for a newly created part.
// For example, StartupSketchXY starts a sketch on the XY origin plane.
// Hosts must reject values outside these constants.
type StartupSketchPlane string

const (
	StartupSketchNone StartupSketchPlane = "none"
	StartupSketchXY   StartupSketchPlane = "xy"
	StartupSketchYZ   StartupSketchPlane = "yz"
	StartupSketchXZ   StartupSketchPlane = "xz"
)

// IsValid reports whether the plane is a supported startup choice.
// For example, StartupSketchXY.IsValid() is true.
func (p StartupSketchPlane) IsValid() bool {
	switch p {
	case StartupSketchNone, StartupSketchXY, StartupSketchYZ, StartupSketchXZ:
		return true
	default:
		return false
	}
}
